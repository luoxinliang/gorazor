package cases

import (
	"bytes"
)

func RenderCodeblock(_buffer bytes.Buffer) {

}

func Codeblock() string {
	var _buffer bytes.Buffer
	RenderCodeblock(_buffer)
	return _buffer.String()
}
