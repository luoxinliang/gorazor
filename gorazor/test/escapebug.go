package cases

import (
	"bytes"
)

func RenderEscapebug(_buffer bytes.Buffer) {
	_buffer.WriteString("<script type=\"text/javascript\">console.log(\"\\n\");</script>")

}

func Escapebug() string {
	var _buffer bytes.Buffer
	RenderEscapebug(_buffer)
	return _buffer.String()
}
