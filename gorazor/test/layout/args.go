package layout

import (
	"bytes"
	"strconv"
	"zfw/models"
)

func RenderArgs(_buffer bytes.Buffer, objs ...*models.Widget) {

	size := strconv.Itoa(12 / len(objs))

}

func Args(objs ...*models.Widget) string {
	var _buffer bytes.Buffer
	RenderArgs(_buffer, objs...)
	return _buffer.String()
}
