package optional

import "sync"

var Default Data = NewSpec()

func GetInt() int {
	return Default.GetInt()
}

func AddInt(value int) {
	Default.AddInt(value)
}

type Data interface {
	GetInt() int
	AddInt(value int)
}

type specificData struct {
	l    *sync.RWMutex
	spec int
}

func NewSpec() *specificData {
	return &specificData{l: &sync.RWMutex{}, spec: 0}
}

func (sd *specificData) GetInt() int {
	sd.l.RLock()
	defer sd.l.RUnlock()
	return sd.spec
}

func (sd *specificData) AddInt(value int) {
	sd.l.Lock()
	defer sd.l.Unlock()
	sd.spec += value
}
