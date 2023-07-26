// automatically generated by stateify.

package inet

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (t *TCPBufferSize) StateTypeName() string {
	return "pkg/sentry/inet.TCPBufferSize"
}

func (t *TCPBufferSize) StateFields() []string {
	return []string{
		"Min",
		"Default",
		"Max",
	}
}

func (t *TCPBufferSize) beforeSave() {}

// +checklocksignore
func (t *TCPBufferSize) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.Min)
	stateSinkObject.Save(1, &t.Default)
	stateSinkObject.Save(2, &t.Max)
}

func (t *TCPBufferSize) afterLoad() {}

// +checklocksignore
func (t *TCPBufferSize) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.Min)
	stateSourceObject.Load(1, &t.Default)
	stateSourceObject.Load(2, &t.Max)
}

func (n *Namespace) StateTypeName() string {
	return "pkg/sentry/inet.Namespace"
}

func (n *Namespace) StateFields() []string {
	return []string{
		"inode",
		"creator",
		"isRoot",
		"userNS",
	}
}

func (n *Namespace) beforeSave() {}

// +checklocksignore
func (n *Namespace) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.inode)
	stateSinkObject.Save(1, &n.creator)
	stateSinkObject.Save(2, &n.isRoot)
	stateSinkObject.Save(3, &n.userNS)
}

// +checklocksignore
func (n *Namespace) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.inode)
	stateSourceObject.LoadWait(1, &n.creator)
	stateSourceObject.Load(2, &n.isRoot)
	stateSourceObject.Load(3, &n.userNS)
	stateSourceObject.AfterLoad(n.afterLoad)
}

func (r *namespaceRefs) StateTypeName() string {
	return "pkg/sentry/inet.namespaceRefs"
}

func (r *namespaceRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *namespaceRefs) beforeSave() {}

// +checklocksignore
func (r *namespaceRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *namespaceRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*TCPBufferSize)(nil))
	state.Register((*Namespace)(nil))
	state.Register((*namespaceRefs)(nil))
}
