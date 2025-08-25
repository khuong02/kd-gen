package enum

import (
	"github.com/dave/jennifer/jen"
	"github.com/khuong02/kd-gen/config"
)

func (e *Enum) Slice(name, enumType string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	// --------------------
	// Slice: all values
	// --------------------
	f.Var().Id("All" + name + "s").
		Op("=").
		Index().Id(name).
		ValuesFunc(func(g *jen.Group) {
			for idx, v := range values {
				if idx == len(values)-1 {
					g.Line().Id(v.Name).Op(",").Line()
					continue
				}
				g.Line().Id(v.Name)
			}
		})
}
