package spin

import "github.com/drop-target-pinball/coroutine"

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
	eng.NewCoroutine(sys.watchDrain)
	eng.NewCoroutine(sys.watchOutLanes)
}

func (s *trackerSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case AddBall:
		s.addBall(act)
	}
}

func (s *trackerSystem) addBall(act AddBall) {
	game := GetGameVars(s.eng)

	n := act.N
	if n == 0 {
		n = 1
	}
	s.queued += n
	if s.queued > s.eng.Config.NumBalls {
		Warn("add ball request capped to max")
		s.queued = s.eng.Config.NumBalls
	}
	game.BallsInPlay = s.queued + s.active
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
		if !switches[s.eng.Config.SwitchShooterLane].Active {
			s.eng.Do(DriverPulse{ID: s.eng.Config.CoilTrough})
			if done := WaitForBallArrivalLoop(e, s.eng.Config.SwitchShooterLane, 500); done {
				return
			}
		}
		if done := WaitForBallDepartureLoop(e, s.eng.Config.SwitchShooterLane, 1000); done {
			return
		}
		s.active += 1
		s.queued -= 1
		e.Post(BallAddedEvent{BallsInPlay: s.active})
	}
	s.processing = false
}

func (s *trackerSystem) watchDrain(e *ScriptEnv) {
	game := GetGameVars(e)

	for {
		if _, done := e.WaitFor(SwitchEvent{ID: s.eng.Config.SwitchDrain}); done {
			return
		}
		if s.active == 0 {
			continue
		}
		s.active -= 1
		game.BallsInPlay -= 1

		if s.draining > 0 {
			s.draining -= 1
		} else if game.BallSave {
			Log("***** BALL SAVE")
			s.addBall(AddBall{})
		}
		e.Post(BallDrainEvent{BallsInPlay: game.BallsInPlay})
	}
}

func (s *trackerSystem) watchOutLanes(e *ScriptEnv) {
	game := GetGameVars(e)
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
		if game.BallSave {
			s.addBall(AddBall{})
		}
	}
}
