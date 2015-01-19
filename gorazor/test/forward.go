package cases

import (
	"bytes"
	"cases/layout"
)

func RenderForward(_buffer bytes.Buffer, content string, err string) {
	body := func() string {
		var _buffer bytes.Buffer

		//hello word
		/* hello this */

		return _buffer.String()
	}

	layout.RenderBase(_buffer, body(), "", "")
}

func Forward(content string, err string) string {
	var _buffer bytes.Buffer
	RenderForward(_buffer, content, err)
	return _buffer.String()
}
