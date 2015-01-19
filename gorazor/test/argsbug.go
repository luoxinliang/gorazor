package cases

import (
	"bytes"
	"cases/layout"
	"github.com/sipin/gorazor/gorazor"
	. "kp/models"
	"tpl/helper"
)

func RenderArgsbug(_buffer bytes.Buffer, totalMessage int, u *User) {
	body := func() string {
		var _buffer bytes.Buffer

		messages := []string{}

		_buffer.WriteString("\n\n<p>")
		_buffer.WriteString(gorazor.HTMLEscape(gorazor.Itoa(args(messages...))))
		_buffer.WriteString("</p>")

		return _buffer.String()
	}

	layout.RenderArgs(_buffer, body())
}

func Argsbug(totalMessage int, u *User) string {
	var _buffer bytes.Buffer
	RenderArgsbug(_buffer, totalMessage, u)
	return _buffer.String()
}
