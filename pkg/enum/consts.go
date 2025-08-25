package enum

import (
	"github.com/dave/jennifer/jen"
	"github.com/khuong02/kd-gen/config"
)

func (e *Enum) Consts(name, enumType string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	f.Const().DefsFunc(func(g *jen.Group) {
		switch enumType {
		case "string":
			for _, v := range values {
				display := displayValue(v.Display, v.Name)
				g.Id(v.Name).Id(name).Op("=").Lit(display)
			}

		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64":
			for i, v := range values {
				if i == 0 {
					// use iota for auto-increment
					g.Id(v.Name).Id(name).Op("=").Iota()
				} else {
					g.Id(v.Name)
				}
			}

		case "float32", "float64":
			for i, v := range values {
				// Floats cannot use iota in the same way, so just assign index as float
				g.Id(v.Name).Id(name).Op("=").Lit(float64(i))
			}

		default:
			panic("unsupported enum type: " + enumType)
		}
	})
}
