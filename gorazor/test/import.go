package cases

import (
	"bytes"
	"hello"
	"huhu"
	"now"
	"strconv"
	"this"
)

func RenderImport(_buffer bytes.Buffer) {
	_buffer.WriteString("\n\n\n<p>hello</p>")

}

func Import() string {
	var _buffer bytes.Buffer
	RenderImport(_buffer)
	return _buffer.String()
}
