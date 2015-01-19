package cases

import (
	"bytes"
)

func RenderHeader(_buffer bytes.Buffer) {
	_buffer.WriteString("<div>Page Header</div>")

}

func Header() string {
	var _buffer bytes.Buffer
	RenderHeader(_buffer)
	return _buffer.String()
}
