package cases

import (
	"bytes"
)

func RenderFooter(_buffer bytes.Buffer) {
	_buffer.WriteString("<div>copyright 2014</div>")

}

func Footer() string {
	var _buffer bytes.Buffer
	RenderFooter(_buffer)
	return _buffer.String()
}
