package builder

import "fmt"

type Monitor string
type SSD string
type Proc string

//Builder pattern separates the construction of a complex object
//from its representation so that the same construction process can create different representations.

const (
	simpleMonitor Monitor = "LG 19'"
	proMonitor    Monitor = "LG 27'"

	proSSD    SSD = "1TB Samsung"
	simpleSSD SSD = "128 NoName"

	proProc    Proc = "IntelCore i9"
	simpleProc Proc = "NoName Proc"
)

type Computer interface {
	TurnOn() string
	TurnOff() string
}

type computer struct {
	output Monitor
	memory SSD
	core   Proc
}

func (c *computer) TurnOn() string {
	return fmt.Sprintf("Turn on a %s monitor, a %s processor, a %s memory\n", c.output, c.core, c.memory)
}

func (c *computer) TurnOff() string {
	return fmt.Sprintf("Turn off a %s monitor, a %s processor, a %s memory\n", c.output, c.core, c.memory)
}

type ComputerBuilder interface {
	Build() *computer
	Output(monitor Monitor) *computerBuilder
	Memory(ssd SSD) *computerBuilder
	Core(proc Proc) *computerBuilder
}

type computerBuilder struct {
	output Monitor
	memory SSD
	core   Proc
}

func NewBuilder() *computerBuilder {
	return &computerBuilder{}
}

func (cb *computerBuilder) Build() *computer {
	return &computer{cb.output, cb.memory, cb.core}
}

func (cb *computerBuilder) Output(monitor Monitor) *computerBuilder {
	cb.output = monitor
	return cb
}

func (cb *computerBuilder) Memory(ssd SSD) *computerBuilder {
	cb.memory = ssd
	return cb
}

func (cb *computerBuilder) Core(proc Proc) *computerBuilder {
	cb.core = proc
	return cb
}

func builder() {
	builder := NewBuilder()
	computerPro := builder.Core(proProc).Memory(proSSD).Output(proMonitor).Build()
	computerSimple := builder.Core(simpleProc).Memory(simpleSSD).Output(simpleMonitor).Build()

	_, _ = computerPro, computerSimple
}
