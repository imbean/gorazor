package cases

import (
	"bytes"
	"dm"
	"zfw/models"
	. "zfw/tplhelper"
)

func Scopebug(obj *models.Widget) string {
	var _buffer bytes.Buffer

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

	return _buffer.String()
}
