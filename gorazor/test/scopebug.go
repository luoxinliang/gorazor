package cases

import (
	"bytes"
	"dm"
	"zfw/models"
	. "zfw/tplhelper"
)

func RenderScopebug(_buffer bytes.Buffer, obj *models.Widget) {

	if 1 == 2 {
	} else {
		values := []int{}
		for _, v := range values {
			if v, ok := v.(type); ok {

				_buffer.WriteString("<a>\n					")
				for _, v := range values {
				}
				_buffer.WriteString("\n				</a>")

			} else {

			}
		}
	}

}

func Scopebug(obj *models.Widget) string {
	var _buffer bytes.Buffer
	RenderScopebug(_buffer, obj)
	return _buffer.String()
}
