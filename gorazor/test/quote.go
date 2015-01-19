package cases

import (
	"bytes"
)

func RenderQuote(_buffer bytes.Buffer) {
	_buffer.WriteString("<html>'text'</html>")

}

func Quote() string {
	var _buffer bytes.Buffer
	RenderQuote(_buffer)
	return _buffer.String()
}
