package cases

import (
	"bytes"
)

func RenderVar(_buffer bytes.Buffer, totalMessage int) {

}

func Var(totalMessage int) string {
	var _buffer bytes.Buffer
	RenderVar(_buffer, totalMessage)
	return _buffer.String()
}
