package cases

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"kp/models"
	"tpl/admin/layout"
)

func RenderSectionbug(_buffer bytes.Buffer) {
	body := func() string {
		var _buffer bytes.Buffer

		return _buffer.String()
	}

	js := func() string {
		var _buffer bytes.Buffer
		for _, jsFile := range ctx.GetJS() {

			_buffer.WriteString("<script src=\"")
			_buffer.WriteString(gorazor.HTMLEscape(jsFile))
			_buffer.WriteString("\"></script>")

		}
		return _buffer.String()
	}

	layout.RenderBase(_buffer, body(), js())
}

func Sectionbug() string {
	var _buffer bytes.Buffer
	RenderSectionbug(_buffer)
	return _buffer.String()
}
