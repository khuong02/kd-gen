package enum_test

import (
	"fmt"
	"testing"

	"github.com/khuong02/kd-gen/config"
	"github.com/khuong02/kd-gen/pkg/enum"
	"github.com/stretchr/testify/assert"
)

func TestConsts_StringEnum(t *testing.T) {
	e := enum.New("test")

	values := []config.EnumValue{
		{Name: "Red", Display: "red"},
		{Name: "Blue", Display: "blue"},
	}

	e.Consts("Color", "string", values)

	code := e.JenFile().GoString()
	assert.Contains(t, code, `Red  Color = "red"`) // with 2 spaces
	assert.Contains(t, code, `Blue Color = "blue"`)

}

func TestConsts_IntEnum(t *testing.T) {
	e := enum.New("test")
	values := []config.EnumValue{
		{Name: "One"},
		{Name: "Two"},
	}
	e.JenFile().Type().Id("Number").Id("int")
	e.Consts("Number", "int", values)

	code := e.JenFile().GoString()
	assert.Contains(t, code, `One Number = iota`) // with 2 spaces
	assert.Contains(t, code, `Two`)
}

func TestConsts_FloatEnum(t *testing.T) {
	e := enum.New("test")
	values := []config.EnumValue{
		{Name: "Pi"},
		{Name: "Euler"},
	}
	e.JenFile().Type().Id("Math").Id("float64")
	e.Consts("Math", "float64", values)

	code := e.JenFile().GoString()
	fmt.Println(code)
	assert.Contains(t, code, `Pi    Math = 0.0`)
	assert.Contains(t, code, `Euler Math = 1.0`)
}

func TestConsts_UnsupportedEnum(t *testing.T) {
	e := enum.New("test")
	values := []config.EnumValue{
		{Name: "Test"},
	}

	assert.Panics(t, func() {
		e.Consts("Weird", "complex128", values)
	})
}
