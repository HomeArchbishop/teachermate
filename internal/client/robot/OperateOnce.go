package robot

import (
	"github.com/go-vgo/robotgo"
)

func OperateOnce(
	linkText string,
	textInputPos Pos,
	sendBtnPos Pos,
	searchLinkClickableFromPos Pos,
) {
	robotgo.Move(textInputPos.x, textInputPos.y)
	robotgo.Click("left", false)
	robotgo.TypeStr(linkText + "?______________________________________________________")
	robotgo.Move(sendBtnPos.x, sendBtnPos.y)
	robotgo.Click("left", false)

	for y := searchLinkClickableFromPos.y; y > 0; y -= 10 {
		color := robotgo.GetPixelColor(int(float64(searchLinkClickableFromPos.x)*dpi), int(float64(y)*dpi))
		if color == "95ec69" {
			for ; y > 0; y-- {
				color := robotgo.GetPixelColor(int(float64(searchLinkClickableFromPos.x)*dpi), int(float64(y)*dpi))
				if color != "95ec69" {
					robotgo.Move(searchLinkClickableFromPos.x, int(y))
					robotgo.Click("left", false)
					break
				}
			}
			break
		}
	}
}
