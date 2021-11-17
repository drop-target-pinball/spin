package spin

import (
	"fmt"
	"log"
	"sync"
)

const (
	Player = "player"
	System = "system"
)

type Vars struct {
	vars  map[string]interface{}
	mutex sync.RWMutex
}

func NewVars() *Vars {
	return &Vars{vars: make(map[string]interface{})}
}

func (v *Vars) getInt(id string) int {
	val, ok := v.vars[id]
	if !ok {
		Warn("no such variable: %v", id)
		return 0
	}
	iVal, ok := val.(int)
	if !ok {
		log.Panicf("variable is not an integer: %v", id)
	}
	return iVal
}

func (v *Vars) Int(id string) int {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	return v.getInt(id)
}

func (v *Vars) SetInt(id string, val int) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.vars[id] = val
}

func (v *Vars) AddInt(id string, val int) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.vars[id] = val + v.getInt(id)
}

func (v *Vars) String(id string) string {
	v.mutex.RLock()
	defer v.mutex.RUnlock()

	val, ok := v.vars[id]
	if !ok {
		Warn("no such variable: %v", id)
		return ""
	}
	return fmt.Sprintf("%v", val)
}

func (v *Vars) SetString(id string, str string) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.vars[id] = str
}

type Namespaces struct {
	spaces map[string]*Vars
	mutex  sync.RWMutex
}

func NewNamespaces() *Namespaces {
	return &Namespaces{spaces: make(map[string]*Vars)}
}

func (n *Namespaces) Create(id string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.spaces[id] = NewVars()
}

func (n *Namespaces) Get(id string) *Vars {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	vars, ok := n.spaces[id]
	if !ok {
		log.Panicf("no such namespace: %v", id)
	}
	return vars
}

func (n *Namespaces) Delete(id string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.spaces, id)
}

func (n *Namespaces) Alias(source string, dest string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	vars := n.Get(source)
	n.spaces[dest] = vars
}
