package jd

import (
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
)

const (
	CoilDiverter           = "jd.CoilDiverter"
	CoilDropTargetReset    = "jd.CoilDropTargetReset"
	CoilDropTargetTrip     = "jd.CoilDropTargetTrip"
	CoilLeftPopper         = "jd.CoilLeftPopper"
	CoilLeftShooterLane    = "jd.CoilLeftShooterLane"
	CoilLeftSling          = "jd.CoilLeftSling"
	CoilKnocker            = "jd.CoilKnocker"
	CoilTrough             = "jd.CoilTrough"
	CoilRightPopper        = "jd.CoilRightPopper"
	CoilRightShooterLane   = "jd.CoilRightShooterLane"
	CoilRightSling         = "jd.CoilRightSling"
	FlipperLeftHold        = "jd.FlipperLeftHold"
	FlipperLeftPower       = "jd.FlipperLeftPower"
	FlipperRightHold       = "jd.FlipperRightHold"
	FlipperRightPower      = "jd.FlipperRightPower"
	FlipperUpperLeftHold   = "jd.FlipperUpperLeftHold"
	FlipperUpperLeftPower  = "jd.FlipperUpperLeftPower"
	FlipperUpperRightHold  = "jd.FlipperUpperRightHold"
	FlipperUpperRightPower = "jd.FlipperUpperRightPower"
	MagnetGlobe            = "jd.MagnetGlobe"
	MotorGlobeArm          = "jd.MotorGlobeArm"
	MotorGlobe             = "jd.MotorGlobe"
)

func RegisterCoils(eng *spin.Engine) {
	eng.Do(spin.RegisterCoil{Address: wpc.C11, ID: CoilDiverter})
	eng.Do(spin.RegisterCoil{Address: wpc.C05, ID: CoilDropTargetReset})
	eng.Do(spin.RegisterCoil{Address: wpc.C10, ID: CoilDropTargetTrip})
	eng.Do(spin.RegisterCoil{Address: wpc.C02, ID: CoilLeftPopper})
	eng.Do(spin.RegisterCoil{Address: wpc.C09, ID: CoilLeftShooterLane})
	eng.Do(spin.RegisterCoil{Address: wpc.C15, ID: CoilLeftSling})
	eng.Do(spin.RegisterCoil{Address: wpc.C07, ID: CoilKnocker})
	eng.Do(spin.RegisterCoil{Address: wpc.C13, ID: CoilTrough})
	eng.Do(spin.RegisterCoil{Address: wpc.C03, ID: CoilRightPopper})
	eng.Do(spin.RegisterCoil{Address: wpc.C08, ID: CoilRightShooterLane})
	eng.Do(spin.RegisterCoil{Address: wpc.C16, ID: CoilRightSling})

	eng.Do(spin.RegisterCoil{Address: wpc.FLLH, ID: FlipperLeftHold})
	eng.Do(spin.RegisterCoil{Address: wpc.FLLM, ID: FlipperLeftPower})
	eng.Do(spin.RegisterCoil{Address: wpc.FLRH, ID: FlipperRightHold})
	eng.Do(spin.RegisterCoil{Address: wpc.FLRM, ID: FlipperRightPower})
	eng.Do(spin.RegisterCoil{Address: wpc.FULH, ID: FlipperUpperLeftHold})
	eng.Do(spin.RegisterCoil{Address: wpc.FULM, ID: FlipperUpperLeftPower})
	eng.Do(spin.RegisterCoil{Address: wpc.FURH, ID: FlipperUpperRightHold})
	eng.Do(spin.RegisterCoil{Address: wpc.FURM, ID: FlipperUpperRightPower})
}

func RegisterMagnets(eng *spin.Engine) {
	eng.Do(spin.RegisterMagnet{Address: wpc.C01, ID: MagnetGlobe})
}

func RegisterMotors(eng *spin.Engine) {
	eng.Do(spin.RegisterMotor{Address: wpc.C04, ID: MotorGlobeArm})
	eng.Do(spin.RegisterMotor{Address: wpc.C06, ID: MotorGlobe})
}
