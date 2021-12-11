package jd

import (
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
)

const (
	AutoSlingLeft              = "jd.AutoSlingLeft"
	AutoSlingRight             = "jd.AutoSlingRight"
	CoilDiverter               = "jd.CoilDiverter"
	CoilDropTargetReset        = "jd.CoilDropTargetReset"
	CoilDropTargetTrip         = "jd.CoilDropTargetTrip"
	CoilLeftPopper             = "jd.CoilLeftPopper"
	CoilLeftShooterLane        = "jd.CoilLeftShooterLane"
	CoilLeftSling              = "jd.CoilLeftSling"
	CoilKnocker                = "jd.CoilKnocker"
	CoilTrough                 = "jd.CoilTrough"
	CoilRightPopper            = "jd.CoilRightPopper"
	CoilRightShooterLane       = "jd.CoilRightShooterLane"
	CoilRightSling             = "jd.CoilRightSling"
	CoilFlipperLeftHold        = "jd.CoilFlipperLeftHold"
	CoilFlipperLeftPower       = "jd.CoilFlipperLeftPower"
	CoilFlipperRightHold       = "jd.CoilFlipperRightHold"
	CoilFlipperRightPower      = "jd.CoilFlipperRightPower"
	CoilFlipperUpperLeftHold   = "jd.CoilFlipperUpperLeftHold"
	CoilFlipperUpperLeftPower  = "jd.CoilFlipperUpperLeftPower"
	CoilFlipperUpperRightHold  = "jd.CoilFlipperUpperRightHold"
	CoilFlipperUpperRightPower = "jd.CoilFlipperUpperRightPower"
	FlipperLeft                = "jd.FlipperLeft"
	FlipperRight               = "jd.FlipperRight"
	FlipperUpperLeft           = "jd.FlipperUpperLeft"
	FlipperUpperRight          = "jd.FlipperUpperRight"
	MagnetGlobe                = "jd.MagnetGlobe"
	MotorGlobeArm              = "jd.MotorGlobeArm"
	MotorGlobe                 = "jd.MotorGlobe"
)

func RegisterCoils(eng *spin.Engine) {
	eng.Do(spin.RegisterCoil{Addr: wpc.C11, ID: CoilDiverter})
	eng.Do(spin.RegisterCoil{Addr: wpc.C05, ID: CoilDropTargetReset})
	eng.Do(spin.RegisterCoil{Addr: wpc.C10, ID: CoilDropTargetTrip})
	eng.Do(spin.RegisterCoil{Addr: wpc.C02, ID: CoilLeftPopper})
	eng.Do(spin.RegisterCoil{Addr: wpc.C09, ID: CoilLeftShooterLane})
	eng.Do(spin.RegisterCoil{Addr: wpc.C15, ID: CoilLeftSling})
	eng.Do(spin.RegisterCoil{Addr: wpc.C07, ID: CoilKnocker})
	eng.Do(spin.RegisterCoil{Addr: wpc.C13, ID: CoilTrough})
	eng.Do(spin.RegisterCoil{Addr: wpc.C03, ID: CoilRightPopper})
	eng.Do(spin.RegisterCoil{Addr: wpc.C08, ID: CoilRightShooterLane})
	eng.Do(spin.RegisterCoil{Addr: wpc.C16, ID: CoilRightSling})

	eng.Do(spin.RegisterCoil{Addr: wpc.FLLH, ID: CoilFlipperLeftHold})
	eng.Do(spin.RegisterCoil{Addr: wpc.FLLM, ID: CoilFlipperLeftPower})
	eng.Do(spin.RegisterCoil{Addr: wpc.FLRH, ID: CoilFlipperRightHold})
	eng.Do(spin.RegisterCoil{Addr: wpc.FLRM, ID: CoilFlipperRightPower})
	eng.Do(spin.RegisterCoil{Addr: wpc.FULH, ID: CoilFlipperUpperLeftHold})
	eng.Do(spin.RegisterCoil{Addr: wpc.FULM, ID: CoilFlipperUpperLeftPower})
	eng.Do(spin.RegisterCoil{Addr: wpc.FURH, ID: CoilFlipperUpperRightHold})
	eng.Do(spin.RegisterCoil{Addr: wpc.FURM, ID: CoilFlipperUpperRightPower})
}

func RegisterAuto(eng *spin.Engine) {
	eng.Do(spin.RegisterFlipper{
		ID:            FlipperLeft,
		SwitchAddr:    wpc.SF4,
		PowerCoilAddr: wpc.FLLM,
		HoldCoilAddr:  wpc.FLLH,
	})
	eng.Do(spin.RegisterFlipper{
		ID:            FlipperRight,
		SwitchAddr:    wpc.SF2,
		PowerCoilAddr: wpc.FLRM,
		HoldCoilAddr:  wpc.FLRH,
	})
	eng.Do(spin.RegisterFlipper{
		ID:            FlipperUpperLeft,
		SwitchAddr:    wpc.SF8,
		PowerCoilAddr: wpc.FULM,
		HoldCoilAddr:  wpc.FULH,
	})
	eng.Do(spin.RegisterFlipper{
		ID:            FlipperUpperRight,
		SwitchAddr:    wpc.SF6,
		PowerCoilAddr: wpc.FURM,
		HoldCoilAddr:  wpc.FURH,
	})

	eng.Do(spin.RegisterAutoPulse{
		ID:         AutoSlingLeft,
		SwitchAddr: wpc.S51,
		CoilAddr:   wpc.C15,
	})
	eng.Do(spin.RegisterAutoPulse{
		ID:         AutoSlingRight,
		SwitchAddr: wpc.S52,
		CoilAddr:   wpc.C16,
	})
}

func RegisterMagnets(eng *spin.Engine) {
	eng.Do(spin.RegisterMagnet{Addr: wpc.C01, ID: MagnetGlobe})
}

func RegisterMotors(eng *spin.Engine) {
	eng.Do(spin.RegisterMotor{Addr: wpc.C04, ID: MotorGlobeArm})
	eng.Do(spin.RegisterMotor{Addr: wpc.C06, ID: MotorGlobe})
}
