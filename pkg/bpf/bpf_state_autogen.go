// automatically generated by stateify.

package bpf

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (p *Program) StateTypeName() string {
	return "pkg/bpf.Program"
}

func (p *Program) StateFields() []string {
	return []string{
		"instructions",
	}
}

func (p *Program) beforeSave() {}

// +checklocksignore
func (p *Program) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.instructions)
}

func (p *Program) afterLoad() {}

// +checklocksignore
func (p *Program) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.instructions)
}

func init() {
	state.Register((*Program)(nil))
}
