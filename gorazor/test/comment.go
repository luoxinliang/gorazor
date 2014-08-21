package cases

import (
	"bytes"
	"cases/layout"
)

func Comment(content string, err string) string {
	var _buffer bytes.Buffer
	if _, ok := ctx.GetSecureCookie("userid"); ok {

		_buffer.WriteString("<!-- #section:basics/sidebar -->")
	}

	return layout.Base(_buffer.String(), "", "")
}
