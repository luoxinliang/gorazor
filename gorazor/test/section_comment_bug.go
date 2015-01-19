package cases

import (
	"bytes"
)

func RenderSection_comment_bug(_buffer bytes.Buffer) {
	_buffer.WriteString("\n\n<a>\n    <!-- comment -->\n</a>")

	side := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString("<!-- comment -->\n    plain text")
		return _buffer.String()
	}

}

func Section_comment_bug() string {
	var _buffer bytes.Buffer
	RenderSection_comment_bug(_buffer)
	return _buffer.String()
}
