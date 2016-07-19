package badge

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

const (
	title     = "BearyChat"
	color     = "#85c158"
	pad   int = 8
	sep   int = 4
)

const tpl = `<svg xmlns="http://www.w3.org/2000/svg" width="{{ .Tw }}" height="20">
  <rect rx="3" width="{{ .Tw }}" height="20" fill="#555"></rect>
  <rect rx="3" x="{{ .Lw }}" width="{{ .Lw }}" height="20" fill="{{ .Color }}"></rect>
  <path d="M{{ .Lw }} 0h{{ .Sep }}v20h-{{ .Sep }}z" fill="{{ .Color }}"></path>
  <g text-anchor="middle" font-family="Verdana" font-size="11">
    {{text .Title .TitleX 14}}
    {{text .Value .ValueX 14}}
  </g>
</svg>`

func text(str string, x, y int) interface{} {
	t := fmt.Sprintf(
		`<text fill="#010101" fill-opacity=".3" x="%d" y="%d">%s</text>
    <text fill="#fff" x="%d" y="%d">%s</text>`,
		x,
		y+1,
		str,
		x,
		y,
		str,
	)

	return t
}

type badgeSetter func(*Badge) error

type Badge struct {
	Total, Active int

	Lw, Rw, Tw int

	Color    string
	Sep, Pad int

	Title  string
	TitleX int

	Value  string
	ValueX int
}

func NewBadge(total, active int, setters ...badgeSetter) *Badge {
	var value string
	if active > 0 {
		value = fmt.Sprintf("%d/%d", active, total)
	} else {
		value = fmt.Sprintf("%d", total)
	}

	badge := &Badge{
		Total:  total,
		Active: active,
		Color:  color,
		Sep:    sep,
		Pad:    pad,
		Title:  title,
		Value:  value,
	}

	for _, setter := range setters {
		if err := setter(badge); err != nil {
			log.Fatal(err)
		}
	}

	return badge
}

func (b *Badge) Calculate() {
	b.Lw = b.Pad + width(b.Title) + b.Sep
	b.Rw = b.Sep + width(b.Value) + b.Pad
	b.Tw = b.Lw + b.Rw
	b.TitleX = b.Lw / 2
	b.ValueX = b.Lw + b.Rw/2
}

func (b *Badge) String() string {
	b.Calculate()

	t, err := template.New("badge").
		Funcs(template.FuncMap{"text": text}).
		Parse(tpl)
	if err != nil {
		return err.Error()
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, b); err != nil {
		return err.Error()
	}
	return buf.String()
}

func width(str string) int {
	return 7 * len(str)
}

func BadgeTitle(title string) badgeSetter {
	return func(b *Badge) error {
		b.Title = fmt.Sprintf("BearyChat: %s", title)

		return nil
	}
}

func BadgeColor(color string) badgeSetter {
	return func(b *Badge) error {
		b.Color = color

		return nil
	}
}
