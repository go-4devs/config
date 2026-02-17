package render

import (
	"bytes"
	"database/sql"
	"encoding"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"time"
)

func Value(name, val string, data ViewData) string {
	rnd := renderType(data)

	res, err := rnd(ValueData{ValName: name, Value: val, ViewData: data})
	if err != nil {
		return fmt.Sprintf("render value:%v", err)
	}

	return res
}

func Type(data ViewData) string {
	dt := data.View.Type()
	rtype := reflect.TypeOf(dt)

	slice := ""
	if rtype.Kind() == reflect.Slice {
		slice = "[]"
		rtype = rtype.Elem()
	}

	short := rtype.Name()

	if rtype.PkgPath() != "" {
		var err error

		short, err = data.AddType(rtype.PkgPath() + "." + rtype.Name())
		if err != nil {
			return err.Error()
		}
	}

	return slice + short
}

func Data(val any, name string, view ViewData) string {
	fn := renderData(view)

	data, err := fn(val, ValueData{ValName: name, Value: "", ViewData: view})
	if err != nil {
		return fmt.Sprintf("render dara:%v", err)
	}

	return data
}

func renderDataTime(val any, _ ValueData) (string, error) {
	data, _ := val.(time.Time)

	return fmt.Sprintf("time.Parse(%q,time.RFC3339Nano)", data.Format(time.RFC3339Nano)), nil
}

func renderDataDuration(val any, _ ValueData) (string, error) {
	data, _ := val.(time.Duration)

	return fmt.Sprintf("time.ParseDuration(%q)", data), nil
}

func renderDataUnmarhal(val any, view ValueData) (string, error) {
	res, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("render data unmarshal:%w", err)
	}

	return fmt.Sprintf("return {{.%[1]s}}, {{.%[1]s}}.UnmarshalJSON(%q)", view.ValName, res), nil
}

func renderDataUnmarhalText(val any, view ValueData) (string, error) {
	res, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("render data unmarshal:%w", err)
	}

	return fmt.Sprintf("return {{.%[1]s}}, {{.%[1]s}}.UnmarshalText(%s)", view.ValName, res), nil
}

func renderDataFlag(val any, view ValueData) (string, error) {
	return fmt.Sprintf("return {{.%[1]s}}, {{.%[1]s}}.Set(%[2]q)", view.ValName, val), nil
}

func renderType(view ViewData) func(data ValueData) (string, error) {
	return dataRender(view).Type
}

func renderData(view ViewData) func(in any, data ValueData) (string, error) {
	return dataRender(view).Value
}

func dataRender(view ViewData) DataRender {
	data := view.View.Type()
	vtype := reflect.TypeOf(data)

	if vtype.Kind() == reflect.Slice {
		return NewDataRender(sliceType, nil)
	}

	if h, ok := render[vtype]; ok {
		return h
	}

	if vtype.Kind() != reflect.Interface && vtype.Kind() != reflect.Pointer {
		vtype = reflect.PointerTo(vtype)
	}

	for extypes := range render {
		if extypes == nil || extypes.Kind() != reflect.Interface {
			continue
		}

		if vtype.Implements(extypes) {
			return render[extypes]
		}
	}

	return render[reflect.TypeOf((any)(nil))]
}

//nolint:gochecknoglobals
var render = map[reflect.Type]DataRender{
	reflect.TypeFor[encoding.TextUnmarshaler](): NewDataRender(unmarshalTextType, renderDataUnmarhalText),
	reflect.TypeFor[json.Unmarshaler]():         NewDataRender(unmarshalType, renderDataUnmarhal),
	reflect.TypeFor[flag.Value]():               NewDataRender(flagType, renderDataFlag),
	reflect.TypeFor[sql.Scanner]():              NewDataRender(scanType, nil),
	reflect.TypeFor[int]():                      NewDataRender(internalType, nil),
	reflect.TypeFor[int64]():                    NewDataRender(internalType, anyValue),
	reflect.TypeFor[bool]():                     NewDataRender(internalType, anyValue),
	reflect.TypeFor[string]():                   NewDataRender(internalType, anyValue),
	reflect.TypeFor[float64]():                  NewDataRender(internalType, anyValue),
	reflect.TypeFor[uint]():                     NewDataRender(internalType, anyValue),
	reflect.TypeFor[int64]():                    NewDataRender(internalType, anyValue),
	reflect.TypeFor[time.Duration]():            NewDataRender(durationType, renderDataDuration),
	reflect.TypeFor[time.Time]():                NewDataRender(timeType, renderDataTime),
	reflect.TypeOf((any)(nil)):                  NewDataRender(anyType, anyValue),
}

func timeType(data ValueData) (string, error) {
	return fmt.Sprintf("return %s.ParseTime()", data.ValName), nil
}

func durationType(data ValueData) (string, error) {
	return fmt.Sprintf("return %s.ParseDuration()", data.ValName), nil
}

func scanType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.Lookup("scan_value.go.tpl").Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("execute scan value:%w", err)
	}

	return b.String(), nil
}

func flagType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.Lookup("flag_value.go.tpl").Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("execute flag value:%w", err)
	}

	return b.String(), nil
}

func anyType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.ExecuteTemplate(&b, "any.go.tpl", data)
	if err != nil {
		return "", fmt.Errorf("unmarshal execute any.go.tpl:%w", err)
	}

	return b.String(), nil
}

func anyValue(data any, _ ValueData) (string, error) {
	return fmt.Sprintf("return %#v, nil", data), nil
}

func unmarshalType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.ExecuteTemplate(&b, "unmarshal_json.go.tpl", data)
	if err != nil {
		return "", fmt.Errorf("unmarshal execute unmarshal_json.go.tpl:%w", err)
	}

	return b.String(), nil
}

func unmarshalTextType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.Lookup("unmarshal_text.go.tpl").Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("execute unmarshal text:%w", err)
	}

	return b.String(), nil
}

func internalType(data ValueData) (string, error) {
	var b bytes.Buffer

	err := parceTpls.Lookup("parse.go.tpl").Execute(&b, data)
	if err != nil {
		return "", fmt.Errorf("internal execute parce.go.tpl:%w", err)
	}

	return b.String(), nil
}

func sliceType(data ValueData) (string, error) {
	return fmt.Sprintf("return %[2]s, %[1]s.Unmarshal(&%[2]s)", data.ValName, data.Value), nil
}

type ValueData struct {
	ViewData

	ValName string
	Value   string
}

func (v ValueData) FuncType() string {
	name := v.Type()

	return v.Rendering.FuncName(name)
}

type DataRender struct {
	renderType  func(data ValueData) (string, error)
	renderValue func(data any, view ValueData) (string, error)
}

func (d DataRender) Type(data ValueData) (string, error) {
	return d.renderType(data)
}

func (d DataRender) Value(data any, view ValueData) (string, error) {
	return d.renderValue(data, view)
}

func NewDataRender(rendeType func(data ValueData) (string, error), renderValue func(data any, view ValueData) (string, error)) DataRender {
	if rendeType == nil {
		rendeType = anyType
	}

	if renderValue == nil {
		renderValue = anyValue
	}

	return DataRender{
		renderType:  rendeType,
		renderValue: renderValue,
	}
}
