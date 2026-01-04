package arg

import (
	"bytes"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

const (
	defaultSpace = 2
)

func ResolveStyle(params param.Params) ViewStyle {
	var vs ViewStyle

	data, ok := params.Param(paramDumpReferenceView)
	if ok {
		vs, _ = data.(ViewStyle)
	}

	return vs
}

type ViewStyle struct {
	Comment Style
	Info    Style
	MLen    int
}

func (v ViewStyle) ILen() int {
	return v.Info.Len() + v.MLen
}

type Style struct {
	Start string
	End   string
}

func (s Style) Len() int {
	return len(s.End) + len(s.Start)
}

func NewDump() Dump {
	return Dump{
		sep:   dash,
		space: defaultSpace,
	}
}

type Dump struct {
	sep   string
	space int
}

func (d Dump) Reference(w io.Writer, opt config.Options) error {
	views := NewViews(opt, nil)
	style := ResolveStyle(opt)
	style.MLen = d.keyMaxLen(views...)

	if args := views.Arguments(); len(args) > 0 {
		if aerr := d.writeArguments(w, style, args...); aerr != nil {
			return fmt.Errorf("write arguments:%w", aerr)
		}
	}

	if opts := views.Options(); len(opts) > 0 {
		if oerr := d.writeOptions(w, style, opts...); oerr != nil {
			return fmt.Errorf("write option:%w", oerr)
		}
	}

	return nil
}

//nolint:mnd
func (d Dump) keyMaxLen(views ...View) int {
	var maxLen int

	for _, vi := range views {
		vlen := len(vi.Name(d.sep)) + 6

		if !vi.IsBool() {
			vlen = vlen*2 + 1
		}

		if def := vi.Default(); def != "" {
			vlen += 2
		}

		if vlen > maxLen {
			maxLen = vlen
		}
	}

	return maxLen
}

func (d Dump) writeArguments(w io.Writer, style ViewStyle, args ...View) error {
	_, err := fmt.Fprintf(w, "\n%sArguments:%s\n",
		style.Comment.Start,
		style.Comment.End,
	)

	for _, arg := range args {
		alen, ierr := fmt.Fprintf(w, "%s%s%s%s",
			strings.Repeat(" ", d.space),
			style.Info.Start,
			arg.Name(d.sep),
			style.Info.End,
		)
		if ierr != nil {
			err = errors.Join(err, ierr)
		}

		_, ierr = fmt.Fprint(w, strings.Repeat(" ", style.ILen()+d.space-alen))
		if ierr != nil {
			err = errors.Join(err, ierr)
		}

		_, ierr = fmt.Fprint(w, arg.Description())
		if ierr != nil {
			err = errors.Join(err, ierr)
		}

		if def := arg.Default(); def != "" {
			ierr := d.writeDefault(w, style, def)
			if ierr != nil {
				err = errors.Join(err, ierr)
			}
		}

		_, ierr = fmt.Fprint(w, "\n")
		if ierr != nil {
			err = errors.Join(err, ierr)
		}
	}

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

//nolint:gocognit,gocyclo,cyclop
func (d Dump) writeOptions(w io.Writer, style ViewStyle, opts ...View) error {
	_, err := fmt.Fprintf(w, "\n%sOptions:%s\n",
		style.Comment.Start,
		style.Comment.End,
	)

	for _, opt := range opts {
		if opt.IsHidden() {
			continue
		}

		var op bytes.Buffer

		_, oerr := fmt.Fprintf(&op, "%s%s", strings.Repeat(" ", d.space), style.Info.Start)
		if oerr != nil {
			err = errors.Join(err, oerr)
		}

		if short := opt.Short(); short != "" {
			op.WriteString("-")
			op.WriteString(short)
			op.WriteString(", ")
		} else {
			op.WriteString("    ")
		}

		op.WriteString("--")
		op.WriteString(opt.Name(d.sep))

		if !opt.IsBool() {
			if !opt.IsRequired() {
				op.WriteString("[")
			}

			op.WriteString("=")
			op.WriteString(strings.ToUpper(opt.Name(d.sep)))

			if !opt.IsRequired() {
				op.WriteString("]")
			}
		}

		_, oerr = fmt.Fprintf(&op, "%s", style.Info.End)
		if oerr != nil {
			err = errors.Join(err, oerr)
		}

		olen, oerr := w.Write(op.Bytes())
		if oerr != nil {
			err = errors.Join(err, oerr)
		}

		_, oerr = fmt.Fprintf(w, "%s%s",
			strings.Repeat(" ", style.ILen()+d.space-olen),
			opt.Description(),
		)
		if oerr != nil {
			err = errors.Join(err, oerr)
		}

		if def := opt.Default(); def != "" {
			oerr = d.writeDefault(w, style, def)
			if oerr != nil {
				err = errors.Join(err, oerr)
			}
		}

		if opt.IsSlice() {
			_, oerr = fmt.Fprintf(w, "%s  (multiple values allowed)%s", style.Comment.Start, style.Comment.End)
			if oerr != nil {
				err = errors.Join(err, oerr)
			}
		}

		_, oerr = fmt.Fprint(w, "\n")
		if oerr != nil {
			err = errors.Join(err, oerr)
		}
	}

	if err != nil {
		return fmt.Errorf("write options:%w", err)
	}

	return nil
}

func (d Dump) writeDefault(w io.Writer, style ViewStyle, data string) error {
	_, err := fmt.Fprintf(w, " %s[default:%s]%s",
		style.Comment.Start,
		data,
		style.Comment.End,
	)
	if err != nil {
		return fmt.Errorf("default:%w", err)
	}

	return nil
}

func NewViews(opts config.Options, parent *View) Views {
	views := make(Views, 0, len(opts.Options()))
	for _, opt := range opts.Options() {
		views = append(views, newViews(opt, parent)...)
	}

	return views
}

func newViews(opt config.Option, parent *View) []View {
	view := NewView(opt, parent)
	switch one := opt.(type) {
	case config.Group:
		return NewViews(one, &view)
	default:
		return []View{view}
	}
}

type Views []View

func (v Views) Arguments() []View {
	args := make([]View, 0, len(v))

	for _, view := range v {
		if view.IsArgument() {
			args = append(args, view)
		}
	}

	sort.Slice(args, func(i, j int) bool {
		return args[i].Pos() < args[j].Pos()
	})

	return args
}

func (v Views) Options() []View {
	opts := make([]View, 0, len(v))

	for _, view := range v {
		if !view.IsArgument() {
			opts = append(opts, view)
		}
	}

	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Name(dash) < opts[j].Name(dash)
	})

	return opts
}

func NewView(params config.Option, parent *View) View {
	pos, ok := ParamArgument(params)

	keys := make([]string, 0)
	if parent != nil {
		keys = append(keys, parent.Keys()...)
	}

	if name := params.Name(); name != "" {
		keys = append(keys, name)
	}

	return View{
		pos:        pos,
		isArgument: ok,
		keys:       keys,
		parent:     parent,
		Params:     params,
	}
}

type View struct {
	param.Params

	keys       []string
	pos        uint64
	isArgument bool
	parent     *View
}

func (v View) Name(delimiter string) string {
	return strings.Join(v.keys, delimiter)
}

func (v View) IsArgument() bool {
	return v.isArgument
}

func (v View) Keys() []string {
	return v.keys
}

func (v View) Pos() uint64 {
	return v.pos
}

func (v View) Default() string {
	data, ok := option.DataDefaut(v.Params)
	if !ok {
		return ""
	}

	switch dt := data.(type) {
	case time.Time:
		return dt.Format(time.RFC3339)
	case encoding.TextMarshaler:
		if res, err := dt.MarshalText(); err == nil {
			return string(res)
		}
	case json.Marshaler:
		if res, err := dt.MarshalJSON(); err == nil {
			return string(res)
		}
	}

	return fmt.Sprintf("%v", data)
}

func (v View) Description() string {
	return param.Description(v.Params)
}

func (v View) IsHidden() bool {
	return option.IsHidden(v.Params)
}

func (v View) Short() string {
	short, _ := option.ParamShort(v.Params)

	return short
}

func (v View) IsRequired() bool {
	return option.IsHidden(v.Params)
}

func (v View) IsBool() bool {
	return option.IsBool(v.Params)
}

func (v View) IsSlice() bool {
	return option.IsSlice(v.Params)
}
