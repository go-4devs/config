package render_test

import (
	"fmt"
	"strconv"
	"testing"

	"gitoa.ru/go-4devs/config/definition/generate/render"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/option"
)

type flagValue int

func (f flagValue) String() string {
	return strconv.Itoa(int(f))
}

func (f *flagValue) Set(in string) error {
	data, err := strconv.Atoi(in)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	*f = flagValue(data)

	return nil
}

func TestValue_FlagType(t *testing.T) {
	t.Parallel()

	const ex = `pval, perr := val.ParseString()
    if perr != nil {
        return v, fmt.Errorf("read [%v]:%w",[]string{"flagValue"}, perr)
    }

    return v, v.Set(pval)`

	viewData := render.NewViewData(nil, view.NewView(option.New("flag_value", "flag desc", flagValue(0)), nil))
	result := render.Value("val", "v", viewData)

	if result != ex {
		t.Errorf("failed render flag type ex:%s, res:%s", ex, result)
	}
}

func TestData_Flag(t *testing.T) {
	t.Parallel()

	const ex = `return {{.val}}, {{.val}}.Set("42")`

	viewData := render.NewViewData(nil, view.NewView(option.New("flag_value", "flag desc", flagValue(0)), nil))
	result := render.Data(flagValue(42), "val", viewData)

	if result != ex {
		t.Errorf("failed render flag value ex:%s, res:%s", ex, result)
	}
}

type scanValue int

func (s *scanValue) Scan(src any) error {
	res, _ := src.(string)

	data, err := strconv.Atoi(res)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	*s = scanValue(data)

	return nil
}

func TestValue_Scan(t *testing.T) {
	t.Parallel()

	const ex = `return v, v.Scan(val.Any())`

	viewData := render.NewViewData(nil, view.NewView(option.New("scan_value", "scan desc", scanValue(42)), nil))
	result := render.Value("val", "v", viewData)

	if result != ex {
		t.Errorf("failed render flag value ex:%s, res:%s", ex, result)
	}
}

type textData string

func (j *textData) UnmarshalText(in []byte) error {
	val := string(in)

	*j = textData(val)

	return nil
}

func TestData_UnmarshalText(t *testing.T) {
	t.Parallel()

	const ex = `return {{.val}}, {{.val}}.UnmarshalText("4devs")`

	data := textData("4devs")
	viewData := render.NewViewData(nil, view.NewView(option.New("tvalue", "unmarshal text desc", textData("")), nil))
	result := render.Data(data, "val", viewData)

	if result != ex {
		t.Errorf("failed render flag value ex:%s, res:%s", ex, result)
	}
}

func TestValue_UnmarshalText(t *testing.T) {
	t.Parallel()

	const ex = `pval, perr := val.ParseString()
    if perr != nil {
        return v, fmt.Errorf("read [%v]:%w", []string{"tvalue"}, perr)
    }

    return v, v.UnmarshalText([]byte(pval))`

	viewData := render.NewViewData(nil, view.NewView(option.New("tvalue", "unmarshal text desc", textData("")), nil))
	result := render.Value("val", "v", viewData)

	if result != ex {
		t.Errorf("failed render flag value ex:%s, res:%s", ex, result)
	}
}
