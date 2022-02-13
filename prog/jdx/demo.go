package jdx

import (
	"os"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func demoScript(e *spin.ScriptEnv) {
	os.Setenv("NO_RANDOM", "true")

	demoIntro(e)
	demoPlunge(e, false)
	demoPursuit(e, false)
	demoBlackout(e, false)
	demoSniper(e, false)
	demoDrain1(e, false)
	demoBattleTank(e, false)
	demoBadImpersonator(e, false)
	demoMeltdown(e, false)
	demoDrain2(e, false)
	demoSafecracker(e, false)
	demoManhunt(e, false)
	demoStakeout(e, false)
	demoDrain3(e, false)

	//demoBlackout(e, true)
}

func demoIntro(e *spin.ScriptEnv) {
	s := spin.NewSequencer(e)

	s.Do(spin.PlayScript{ID: "ScriptInit"})
	s.Sleep(24_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchStartButton})
	s.Sleep(5_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchStartButton})

	s.Run()
}

func demoPlunge(e *spin.ScriptEnv, single bool) {
	if single {
		e.Do(spin.PlayScript{ID: ScriptGame})
	}

	s := spin.NewSequencer(e)
	s.Sleep(13_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	s.Sleep(1_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftSling})
	s.Sleep(500)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightSling})
	s.Sleep(750)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftSling})
	s.Sleep(8_000)
	s.Run()

	vars := GetVars(e)
	vars.SelectedMode = ModePursuit
}

func demoPursuit(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
	s.Sleep(4_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
	s.Sleep(4_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
	s.Sleep(10_000)

	s.Run()
}

func demoBlackout(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(8_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit})
	s.Sleep(8_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit})
	s.Sleep(8_000)
	s.Post(spin.BallDrainEvent{BallsInPlay: 1})
	s.Sleep(10_000)

	s.Run()
}

func demoSniper(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightPopper})
	s.Sleep(15_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightPopper})
	s.Sleep(10_000)

	s.Run()
}

func demoDrain1(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
	}

	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftOutlane})
	s.Sleep(2_000)
	s.Post(spin.BallDrainEvent{BallsInPlay: 0})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftFireButton})
	s.Sleep(9_000)

	s.Run()
}

func demoBattleTank(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(6_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchOuterLoopLeft})
	s.Sleep(6_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit})
	s.Sleep(6_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchBankTargets})
	s.Sleep(10_000)

	s.Run()
}

func demoBadImpersonator(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(17_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchDropTargetE})
	s.Sleep(9_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchDropTargetE})
	s.Sleep(16_000)

	s.Run()
}

func demoMeltdown(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(7_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchCaptiveBall1})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchCaptiveBall2})
	s.Sleep(12_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchCaptiveBall3})
	s.Sleep(10_000)

	s.Run()
}

func demoDrain2(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
	}

	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftOutlane})
	s.Sleep(2_000)
	s.Post(spin.BallDrainEvent{BallsInPlay: 0})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftFireButton})
	s.Sleep(9_000)

	s.Run()
}

func demoSafecracker(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
	s.Sleep(10_000)

	s.Run()
}

func demoManhunt(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(8_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
	s.Sleep(12_000)

	s.Run()
}

func demoStakeout(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
	}

	s.Post(spin.Message{ID: MessageStartChainMode})
	s.Sleep(8_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
	s.Sleep(10_000)
	s.Post(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
	s.Sleep(22_000)

	s.Run()
}

func demoDrain3(e *spin.ScriptEnv, single bool) {
	s := spin.NewSequencer(e)

	if single {
		s.Do(spin.PlayScript{ID: ScriptGame})
		s.Sleep(250)
		s.Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton})
		s.Sleep(250)
	}

	s.Post(spin.SwitchEvent{ID: jd.SwitchLeftOutlane})
	s.Sleep(2_000)
	s.Post(spin.BallDrainEvent{BallsInPlay: 0})

	s.Run()
}
