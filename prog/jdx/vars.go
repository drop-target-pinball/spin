package jdx

import "github.com/drop-target-pinball/spin"

type Prog struct {
	ManualRightPopper bool
}

func ProgVars(store spin.Store) *Prog {
	v, ok := store.Vars("jdx")
	var vars *Prog
	if ok {
		vars = v.(*Prog)
	} else {
		vars = &Prog{
			ManualRightPopper: false,
		}
		store.RegisterVars("jdx", vars)
	}
	return vars
}
