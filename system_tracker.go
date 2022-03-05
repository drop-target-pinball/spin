package spin

import (
	"github.com/drop-target-pinball/coroutine"
)

/*
AddBall N=2
SwitchEvent ID=jd.SwitchRightShooterLane
SwitchEvent ID=jd.SwitchRightShooterLane Released=true
*/

type trackerSystem struct {
	eng        *Engine
	active     int
	queued     int
	draining   int
	processing bool
}

func RegisterTrackerSystem(eng *Engine) {
	sys := &trackerSystem{
		eng: eng,
	}
	eng.RegisterActionHandler(sys)
	//eng.NewCoroutine(sys.watchDrain)
	eng.NewCoroutine(sys.watchOutLanes)
	eng.NewCoroutine(sys.watchShooterLane)
	eng.NewCoroutine(sys.watchTrough)
	eng.NewCoroutine(sys.watchPlayfield)
}

func (s *trackerSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case AddBall:
		s.addBall(act)
	}
}

func (s *trackerSystem) addBall(act AddBall) {
	n := act.N
	if n == 0 {
		n = 1
	}
	s.queued += n
	if s.queued > s.eng.Config.NumBalls {
		Warn("add ball request capped to max")
		s.queued = s.eng.Config.NumBalls
	}
	//game.BallsInPlay = s.queued + s.active
	if !s.processing {
		s.processing = true
		s.eng.NewCoroutine(s.launchBall)
	}
}

func (s *trackerSystem) launchBall(e *ScriptEnv) {
	switches := GetResourceVars(e).Switches

	// e.NewCoroutine(func(e *ScriptEnv) {
	// 	for {
	// 		if done := WaitForBallArrivalLoop(e, s.eng.Config.SwitchTroughJam, 1000); done {
	// 			return
	// 		}
	// 		s.eng.Do(DriverPulse{ID: s.eng.Config.CoilTrough})
	// 	}
	// })

	for s.queued > 0 {
		Log("launching ball")
		if !switches[s.eng.Config.SwitchShooterLane].Active {
			retries := 0
			for {
				s.eng.Do(DriverPulse{ID: s.eng.Config.CoilTrough})
				evt, done := e.WaitForUntil(2_000, BallLaunchReady{})
				if done {
					return
				}
				// Ball is ready
				if evt != nil {
					break
				}
				retries += 1
				if retries > 5 {
					Log("(*) ball eject failed")
					break
				}
				Log("(!) retry ball eject")
			}
			if retries > 5 {
				break
			}
		}

		Log("ball ready for launch")
		if done := WaitForBallDepartureLoop(e, s.eng.Config.SwitchShooterLane, 250); done {
			return
		}
		s.active += 1
		s.queued -= 1
		e.Post(BallAddedEvent{BallsInPlay: s.active, Queued: s.queued})
	}
	s.processing = false
}

// func (s *trackerSystem) watchDrain(e *ScriptEnv) {
// 	game := GetGameVars(e)

// 	for {
// 		if _, done := e.WaitFor(SwitchEvent{ID: s.eng.Config.SwitchDrain}); done {
// 			return
// 		}
// 		if s.active == 0 {
// 			log.Print("(!) no balls were active when drained")
// 			continue
// 		}
// 		s.active -= 1
// 		game.BallsInPlay -= 1

// 		if s.draining > 0 {
// 			s.draining -= 1
// 		} else if game.BallSave {
// 			s.addBall(AddBall{})
// 		}
// 		e.Post(BallDrainEvent{BallsInPlay: game.BallsInPlay})
// 	}
// }

func (s *trackerSystem) watchOutLanes(e *ScriptEnv) {
	//game := GetGameVars(e)
	outlanes := make([]coroutine.Event, len(s.eng.Config.SwitchWillDrain))
	for i, id := range s.eng.Config.SwitchWillDrain {
		outlanes[i] = SwitchEvent{ID: id}
	}

	for {
		if _, done := e.WaitFor(outlanes...); done {
			return
		}
		s.draining += 1
		s.eng.Post(BallWillDrainEvent{})
		// if game.BallSave {
		// 	s.addBall(AddBall{})
		// }
	}
}

func (s *trackerSystem) watchShooterLane(e *ScriptEnv) {
	game := GetGameVars(e)
	switches := GetResourceVars(e).Switches
	for {
		// FIXME: The problem here is that the coroutine may start before
		// the switches have been registered. In that case, the switches
		// lookup will fail. If not yet configured, wait a bit and try again.
		// This should be fixed so that all switches are registered before
		// this starts.
		sw, ok := switches[e.Config.SwitchShooterLane]
		if !ok {
			if done := e.Sleep(100); done {
				return
			}
			continue
		}
		if sw.Active {
			if done := WaitForBallDepartureLoop(e, e.Config.SwitchShooterLane, 150); done {
				return
			}
			game.BallLaunchReady = false
			Log("ball departed")
		} else {
			if done := WaitForBallArrivalLoop(e, e.Config.SwitchShooterLane, 150); done {
				return
			}
			game.BallLaunchReady = true
			e.Post(BallLaunchReady{})
			Log("ball arrived")
		}
	}
}

func (s *trackerSystem) countBallsInTrough(e *ScriptEnv) (int, bool) {

	switches := GetResourceVars(e).Switches
	balls := 0
	for _, name := range s.eng.Config.SwitchTrough {
		sw, ok := switches[name]
		if !ok {
			// FIXME: same issue as above
			if done := e.Sleep(100); done {
				return 0, false
			}

		} else {
			if sw.Active {
				balls += 1
			}
		}
	}
	jam := switches[e.Config.SwitchTroughJam].Active
	return balls, jam
}

func (s *trackerSystem) watchTrough(e *ScriptEnv) {
	game := GetGameVars(e)
	balls, jam := s.countBallsInTrough(e)

	// FIXME: Is there a better way?
	troughEvents := make([]coroutine.Event, 0)
	for _, name := range e.Config.SwitchTrough {
		troughEvents = append(troughEvents, SwitchEvent{ID: name})
	}
	troughEvents = append(troughEvents, SwitchEvent{ID: e.Config.SwitchTroughJam})

	for {
		// Wait for initial change in the trough
		if _, done := e.WaitFor(troughEvents...); done {
			return
		}
		// Now wait for things to settle a bit. If we get another trough
		// switch change in this time, restart the timer.
		for {
			evt, done := e.WaitForUntil(150, troughEvents...)
			if done {
				return
			}
			// If we timed out, check the status of the trough
			if evt == nil {
				break
			}
			// If we saw another switch event, the balls are still bouncing
			// around in the trough
		}
		newBalls, newJam := s.countBallsInTrough(e)
		if balls != newBalls || jam != newJam {
			change := newBalls - balls
			game.BallsInPlay = e.Config.NumBalls - newBalls
			e.Post(TroughEvent{Balls: newBalls, Jam: newJam, Change: change})
			balls, jam = newBalls, newJam
		}
	}
}

func (s *trackerSystem) watchPlayfield(e *ScriptEnv) {
	game := GetGameVars(e)

	for {
		if !game.PlayfieldActive {
			for {
				if _, done := e.WaitFor(e.Config.PlayfieldSwitches...); done {
					return
				}
				count, _ := s.countBallsInTrough(e)
				if count != e.Config.NumBalls {
					game.PlayfieldActive = true
					break
				}
			}
		} else { // playfield is active
			for {
				evt, done := e.WaitFor(TroughEvent{})
				if done {
					return
				}
				event := evt.(TroughEvent)
				if event.Change > 0 {
					e.Post(BallDrainEvent{BallsInPlay: game.BallsInPlay})
					if game.BallSave {
						e.Do(AddBall{})
					}
				}

				if event.Balls == e.Config.NumBalls {
					game.PlayfieldActive = false
					break
				}
			}
		}
	}
}
