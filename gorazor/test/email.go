package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func RenderEmail(_buffer bytes.Buffer) {
	_buffer.WriteString("<span>rememberingsteve@apple.com ")
	_buffer.WriteString(gorazor.HTMLEscape(username))
	_buffer.WriteString("</span>")

}

func Email() string {
	var _buffer bytes.Buffer
	RenderEmail(_buffer)
	return _buffer.String()
}
