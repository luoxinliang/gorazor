package cases

import (
	"bytes"
)

func RenderBlk(_buffer bytes.Buffer) {

}

func Blk() string {
	var _buffer bytes.Buffer
	RenderBlk(_buffer)
	return _buffer.String()
}
