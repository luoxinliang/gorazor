package cases

import (
	"bytes"
)

func RenderComment(_buffer bytes.Buffer) {
	_buffer.WriteString("\n\n\n\n<p>hello </p>")

	hello

}

func Comment() string {
	var _buffer bytes.Buffer
	RenderComment(_buffer)
	return _buffer.String()
}
