package cases

import (
	"bytes"
)

func RenderDouble_quote(_buffer bytes.Buffer) {
	_buffer.WriteString("<meta charset=\"utf-8\" />")

}

func Double_quote() string {
	var _buffer bytes.Buffer
	RenderDouble_quote(_buffer)
	return _buffer.String()
}
