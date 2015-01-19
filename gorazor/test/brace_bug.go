package cases

import (
	"bytes"
)

func RenderBrace_bug(_buffer bytes.Buffer) {

	isActive := func(name string) {
		if active == name {

			_buffer.WriteString("<li class=\"active\">\n        ")
		} else {

			_buffer.WriteString("<li>\n        ")
		}
	}

}

func Brace_bug() string {
	var _buffer bytes.Buffer
	RenderBrace_bug(_buffer)
	return _buffer.String()
}
