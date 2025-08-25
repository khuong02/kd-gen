package enum

import (
	"github.com/dave/jennifer/jen"
	"github.com/khuong02/kd-gen/config"
)

func (e *Enum) String(name string, enumType string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	// --------------------
	// string function
	// This function returns the string representation of the enum value.
	// If the value is not found, it returns "UNKNOWN".
	// --------------------
	f.Func().
		Params(jen.Id("x").Id(name)).
		Id("String").
		Params().
		String().
		Block(
			jen.Switch(jen.Id("x")).BlockFunc(func(g *jen.Group) {
				for _, v := range values {
					g.Case(jen.Id(v.Name)).Return(jen.Lit(displayValue(v.Display, v.Name)))
				}
			}),
			jen.Return(jen.Lit("UNKNOWN")),
		)
}

func (e *Enum) Parse(name, enumType string, values []config.EnumValue) {
	f := e.f
	defer f.Line()
	// --------------------
	// ParseXXX function
	// --------------------
	f.Func().
		Id("Parse"+name).
		Params(jen.Id("s").String()).
		Params(jen.Id(name), jen.Error()).
		Block(
			jen.Switch(jen.Id("s")).BlockFunc(func(g *jen.Group) {
				for _, v := range values {
					g.Case(jen.Lit(displayValue(v.Display, v.Name))).Return(jen.Id(v.Name), jen.Nil())
				}
			}),
			jen.Return(
				jen.Lit(emptyValue(enumType)),
				jen.Qual("fmt", "Errorf").Call(
					jen.Lit("invalid "+name+": %s"), jen.Id("s"),
				),
			),
		)
}

func (e *Enum) Normalize(name string) {
	f := e.f
	defer f.Line()
	// --------------------
	// Normalize function
	// --------------------
	f.Func().
		Params(jen.Id("x").Id(name)).
		Id("Normalize").
		Params().
		Id(name).
		Block(
			jen.Return(
				jen.Id(name).Call(
					jen.Qual("strings", "ToLower").Call(
						jen.String().Parens(jen.Id("x")),
					),
				),
			),
		)
}

func (e *Enum) JSON(name string) {
	f := e.f
	defer f.Line()

	// MarshalJSON
	f.Func().
		Params(jen.Id("x").Id(name)).
		Id("MarshalJSON").
		Params().
		Params(jen.Index().Byte(), jen.Error()).
		Block(
			// Chỉ marshal string value
			jen.Return(jen.Qual("encoding/json", "Marshal").Call(jen.Id("x"))),
		)

	// UnmarshalJSON
	f.Func().
		Params(jen.Id("x").Op("*").Id(name)).
		Id("UnmarshalJSON").
		Params(jen.Id("data").Index().Byte()).
		Error().
		Block(
			jen.Var().Id("s").String(),
			jen.If(
				jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").
					Call(jen.Id("data"), jen.Op("&").Id("s")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.Return(jen.Nil()),
		)
}

func (e *Enum) SQLValueScan(name string) {
	f := e.f
	defer f.Line()

	// Value() (driver.Value, error)
	f.Func().
		Params(jen.Id("x").Id(name)).
		Id("Value").
		Params().
		Params(jen.Qual("database/sql/driver", "Value"), jen.Error()).
		Block(
			jen.Return(
				jen.Qual("encoding/json", "Marshal").Call(jen.Id("x")),
			),
		)

	// Scan(value interface{}) error
	f.Func().
		Params(jen.Id("x").Op("*").Id(name)).
		Id("Scan").
		Params(jen.Id("value").Interface()).
		Error().
		Block(
			jen.Var().Id("v").Id(name),
			jen.If(
				jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").
					Call(jen.Id("value").Op(".([]byte)"), jen.Op("&").Id("v")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.Return(jen.Nil()),
		)
}
