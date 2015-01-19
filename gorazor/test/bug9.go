package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func RenderBug9(_buffer bytes.Buffer, l *Locale) {
	_buffer.WriteString("\n<span>")
	_buffer.WriteString(gorazor.HTMLEscape(l.T(`for`)))
	_buffer.WriteString("</span>")

}

func Bug9(l *Locale) string {
	var _buffer bytes.Buffer
	RenderBug9(_buffer, l)
	return _buffer.String()
}
