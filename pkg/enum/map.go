package enum

import (
	"github.com/dave/jennifer/jen"
	"github.com/khuong02/kd-gen/config"
)

func (e *Enum) Map(name, enumType string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	// --------------------
	// Optional: Map if Code is set
	// --------------------
	hasCode := false
	for _, v := range values {
		if v.Code != nil {
			hasCode = true
			break
		}
	}

	if hasCode {
		f.Var().Id(name + "Map").
			Op("=").
			Map(jen.Id(name)).Interface().
			ValuesFunc(func(g *jen.Group) {
				for _, v := range values {
					if v.Code == nil {
						continue
					}
					isLast := v == values[len(values)-1]
					switch c := v.Code.(type) {
					case string:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Lit(c)
					case int:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Lit(c)
					case bool:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Lit(c)
					case []string:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Index().String().ValuesFunc(func(arr *jen.Group) {
							for _, s := range c {
								arr.Lit(s)
							}
						})
					case []float64:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Index().Float64().ValuesFunc(func(arr *jen.Group) {
							for _, f := range c {
								arr.Lit(f)
							}
						})
					case []int64:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Index().Int64().ValuesFunc(func(arr *jen.Group) {
							for _, n := range c {
								arr.Lit(n)
							}
						})
					case []int:
						if isLast {
							g.Line().Id(v.Name).Op(":").Lit(c).Op(",").Line()
							continue
						}
						g.Line().Id(v.Name).Op(":").Index().Int().ValuesFunc(func(arr *jen.Group) {
							for _, n := range c {
								arr.Lit(n)
							}
						})
					}
				}
			})
	}
}
