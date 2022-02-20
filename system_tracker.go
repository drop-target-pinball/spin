package spin

import (
	"log"

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
	eng.NewCoroutine(sys.watchDrain)
	eng.NewCoroutine(sys.watchOutLanes)
	eng.NewCoroutine(sys.watchShooterLane)
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
		log.Printf("launching ball")
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
					log.Printf("(*) ball eject failed")
					break
				}
				log.Printf("(!) retry ball eject")
			}
			if retries > 5 {
				break
			}
		}

		log.Printf("ball ready for launch")
		if done := WaitForBallDepartureLoop(e, s.eng.Config.SwitchShooterLane, 250); done {
			return
		}
		s.active += 1
		s.queued -= 1
		e.Post(BallAddedEvent{BallsInPlay: s.active, Queued: s.queued})
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
			log.Print("(!) no balls were active when drained")
			continue
		}
		s.active -= 1
		game.BallsInPlay -= 1

		if s.draining > 0 {
			s.draining -= 1
		} else if game.BallSave {
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
			log.Printf("ball departed")
		} else {
			if done := WaitForBallArrivalLoop(e, e.Config.SwitchShooterLane, 150); done {
				return
			}
			game.BallLaunchReady = true
			e.Post(BallLaunchReady{})
			log.Printf("ball arrived")
		}
	}
}
