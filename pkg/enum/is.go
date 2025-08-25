package enum

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/khuong02/kd-gen/config"
)

func (e *Enum) Is(name string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	// --------------------
	// IsX checks whether the enum equals a specific value
	// --------------------
	for _, v := range values {
		f.Func().
			Params(jen.Id("x").Id(name)).
			Id(fmt.Sprintf("Is%s", v.Name)).
			Params().
			Id("bool").
			Block(
				jen.Return(
					jen.Id("x").Op("==").Id(v.Name),
				),
			)
	}
}
