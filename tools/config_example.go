package tools

var yamlExamples = map[string]string{
	"string": `enums:
  - name: Color
    type: string
	methods:
	- String|Parse|Normalize|JSON|SQL
    values:
      - name: Red
        display: "Red"
        code: red
      - name: Blue
        display: "Blue"
        code: blue`,

	"int": `enums:
  - name: Status
    type: int
	methods:
	- String|Parse|Normalize|JSON|SQL
    values:
      - name: Pending
        code: 0
      - name: Active
        code: 1
      - name: Inactive
        code: 2`,

	"uint64": `enums:
  - name: Permission
    type: uint64
	methods:
	- String|Parse|Normalize|JSON|SQL
    values:
      - name: Read
        code: 1
      - name: Write
        code: 2
      - name: Execute
        code: 4`,

	"float64": `enums:
  - name: Ratio
    type: float64
	methods:
	- String|Parse|Normalize|JSON|SQL
    values:
      - name: Low
        code: 0.1
      - name: Medium
        code: 0.5
      - name: High
        code: 0.9`,

	"bool": `enums:
  - name: Switch
    type: bool
	methods:
	- String|Parse|Normalize|JSON|SQL
    values:
      - name: Off
        code: false
      - name: On
        code: true`,
}

// Small int-family types → reuse "int" example
var intFamily = map[string]bool{
	"int8":   true,
	"int16":  true,
	"int32":  true,
	"int64":  true,
	"uint":   true,
	"uint8":  true,
	"uint16": true,
	"uint32": true,
}
