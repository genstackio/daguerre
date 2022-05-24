package services

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/genstackio/daguerre/commons"
)

func connectCtxEntries(d *diagram.Diagram, a *commons.CtxEntry, b *commons.CtxEntry, mode string, label string) {

	if nil == a || nil == b || (a.Dnode == nil && a.Dgroup == nil && b.Dnode == nil && b.Dgroup == nil) {
		return
	}
	var n1 *diagram.Node
	var n2 *diagram.Node

	if nil != a.Dnode {
		n1 = a.Dnode
		if nil != b.Dnode {
			n2 = b.Dnode
		} else if nil != b.Dgroup && len(b.Dgroup.Nodes()) == 1 {
			n2 = b.Dgroup.Nodes()[0]
		}
	} else if nil != a.Dgroup && len(a.Dgroup.Nodes()) == 1 {
		n1 = a.Dgroup.Nodes()[0]
		if nil != b.Dnode {
			n2 = b.Dnode
		} else if nil != b.Dgroup && len(b.Dgroup.Nodes()) == 1 {
			n2 = b.Dgroup.Nodes()[0]
		}
	}

	if nil != n1 && nil != n2 {
		if len(label) > 0 {
			if "-" == mode {
				d.Connect(n1, n2, noDirections, labellize(label))
			} else if ">" == mode {
				d.Connect(n1, n2, diagram.Forward(), labellize(label))
			} else if "<" == mode {
				d.Connect(n1, n2, diagram.Reverse(), labellize(label))
			} else if "=" == mode {
				d.Connect(n1, n2, diagram.Bidirectional(), labellize(label))
			} else {
				d.Connect(n1, n2, noDirections, labellize(label))
			}
		} else {
			if "-" == mode {
				d.Connect(n1, n2, noDirections)
			} else if ">" == mode {
				d.Connect(n1, n2, diagram.Forward())
			} else if "<" == mode {
				d.Connect(n1, n2, diagram.Reverse())
			} else if "=" == mode {
				d.Connect(n1, n2, diagram.Bidirectional())
			} else {
				d.Connect(n1, n2, noDirections)
			}
		}
	}
}

func labellize(label string) func(options *diagram.EdgeOptions) {
	return func(o *diagram.EdgeOptions) {
		o.Label = label
		o.Attributes["label"] = label
	}
}

func noDirections(o *diagram.EdgeOptions) {
	o.Forward = false
	o.Reverse = false
}
