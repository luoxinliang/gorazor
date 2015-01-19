package cases

import (
	"bytes"
	"strconv"
	"zfw/models"
)

func RenderSlashbug(_buffer bytes.Buffer, objs ...*models.Widget) {

	size := strconv.Itoa(12 / len(objs))

}

func Slashbug(objs ...*models.Widget) string {
	var _buffer bytes.Buffer
	RenderSlashbug(_buffer, objs...)
	return _buffer.String()
}
