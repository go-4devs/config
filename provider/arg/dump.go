package arg

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
)

const (
	defaultSpace     = 2
	flagNameRepeat   = 2
	shortColumnWidth = 4
)

// ResolveStyle returns the dump style stored on params.
func ResolveStyle(params param.Params) ViewStyle {
	var vs ViewStyle

	data, ok := params.Param(paramDumpReferenceView)
	if ok {
		vs, _ = data.(ViewStyle)
	}

	return vs
}

// WithStyle stores the comment and info styles used by a dump.
func WithStyle(comment, info Style) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, paramDumpReferenceView, ViewStyle{
			Comment: comment,
			Info:    info,
			MLen:    0,
		})
	}
}

type ViewStyle struct {
	Comment Style
	Info    Style
	MLen    int
}

// ILen returns the info style length plus the padding.
func (v ViewStyle) ILen() int {
	return v.Info.Len() + v.MLen
}

type Style struct {
	Start string
	End   string
}

// Len returns the length of the style markers.
func (s Style) Len() int {
	return len(s.End) + len(s.Start)
}

// NewDump returns a dump that writes flag help.
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

// Reference writes arguments and options to w.
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

func (d Dump) keyMaxLen(views ...View) int {
	var maxLen int

	for _, vi := range views {
		vlen := len(vi.Name(d.sep)) + d.space

		if !vi.IsArgument() {
			if !vi.IsBool() {
				vlen = vlen*flagNameRepeat + 1
			}

			if def := vi.Default(); def != "" {
				vlen += d.space
			}

			vlen += shortColumnWidth + d.space
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

func (d Dump) writeOptions(w io.Writer, style ViewStyle, opts ...View) error {
	_, err := fmt.Fprintf(w, "\n%sOptions:%s\n",
		style.Comment.Start,
		style.Comment.End,
	)

	for _, opt := range opts {
		if opt.IsHidden() {
			continue
		}

		if oerr := d.writeOption(w, style, opt); oerr != nil {
			err = errors.Join(err, oerr)
		}
	}

	if err != nil {
		return fmt.Errorf("write options:%w", err)
	}

	return nil
}

func (d Dump) writeOption(w io.Writer, style ViewStyle, opt View) error {
	label := d.optionLabel(opt, style)

	olen, err := io.WriteString(w, label)
	if err != nil {
		return fmt.Errorf("option label:%w", err)
	}

	_, err = fmt.Fprintf(w, "%s%s",
		strings.Repeat(" ", style.ILen()+d.space-olen),
		opt.Description(),
	)
	if err != nil {
		return fmt.Errorf("option description:%w", err)
	}

	if def := opt.Default(); def != "" {
		if derr := d.writeDefault(w, style, def); derr != nil {
			return derr
		}
	}

	if opt.IsSlice() {
		_, err = fmt.Fprintf(w, "%s  (multiple values allowed)%s", style.Comment.Start, style.Comment.End)
		if err != nil {
			return fmt.Errorf("option slice:%w", err)
		}
	}

	_, err = fmt.Fprint(w, "\n")
	if err != nil {
		return fmt.Errorf("option end:%w", err)
	}

	return nil
}

func (d Dump) optionLabel(opt View, style ViewStyle) string {
	var label strings.Builder

	label.WriteString(strings.Repeat(" ", d.space))
	label.WriteString(style.Info.Start)

	if short := opt.Short(); short != "" {
		label.WriteString("-")
		label.WriteString(short)
		label.WriteString(", ")
	} else {
		label.WriteString(strings.Repeat(" ", shortColumnWidth))
	}

	label.WriteString("--")

	name := opt.Name(d.sep)
	label.WriteString(name)
	label.WriteString(optionValue(opt, name))
	label.WriteString(style.Info.End)

	return label.String()
}

func optionValue(opt View, name string) string {
	if opt.IsBool() {
		return ""
	}

	value := "=" + strings.ToUpper(name)
	if opt.IsRequired() {
		return value
	}

	return "[" + value + "]"
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

// NewViews returns dump views for opts, including options inside groups.
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

// Arguments returns positional views ordered by position.
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

// Options returns flag views ordered by position.
func (v Views) Options() []View {
	opts := make([]View, 0, len(v))

	for _, view := range v {
		if !view.IsArgument() {
			opts = append(opts, view)
		}
	}

	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Pos() < opts[j].Pos()
	})

	return opts
}

// NewView returns a dump view for one option.
func NewView(params config.Option, parent *View) View {
	pos, ok := ParamArgument(params)

	keys := make([]string, 0)
	if parent != nil {
		keys = append(keys, parent.Keys()...)
	}

	if !ok {
		pos = param.Position(params)
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

// Name joins the path with delimiter.
func (v View) Name(delimiter string) string {
	return strings.Join(v.keys, delimiter)
}

// IsArgument reports whether the view is a positional argument.
func (v View) IsArgument() bool {
	return v.isArgument
}

// Keys returns the option path.
func (v View) Keys() []string {
	return v.keys
}

// Pos returns the position.
func (v View) Pos() uint64 {
	return v.pos
}

// Default returns the default value as text.
func (v View) Default() string {
	data, ok := param.Default(v.Params)
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

// Description returns the option description.
func (v View) Description() string {
	return param.Description(v.Params)
}

// IsHidden reports whether the option is hidden.
func (v View) IsHidden() bool {
	return option.IsHidden(v.Params)
}

// Short returns the one-letter flag name.
func (v View) Short() string {
	short, _ := option.ParamShort(v.Params)

	return short
}

// IsRequired reports whether the option is hidden.
func (v View) IsRequired() bool {
	return option.IsHidden(v.Params)
}

// IsBool reports whether the option is a bool flag.
func (v View) IsBool() bool {
	return option.IsBool(v.Params)
}

// IsSlice reports whether the option is a slice.
func (v View) IsSlice() bool {
	return option.IsSlice(v.Params)
}
