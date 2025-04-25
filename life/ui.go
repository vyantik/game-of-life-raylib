package life

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (l *life) drawHint() {
	currentScreenWidth := int32(rl.GetScreenWidth())
	currentScreenHeight := int32(rl.GetScreenHeight())

	hintText := "Press 'F' to toggle fullscreen"
	fontSize := max(int32(currentScreenHeight/40), 10)
	textPosX := int32(currentScreenWidth / 50)
	textPosY := int32(currentScreenHeight - fontSize - currentScreenHeight/50)

	textColor := rl.White
	rl.DrawText(hintText, textPosX, textPosY, fontSize, textColor)
}

func (l *life) drawUI() {
	uiText := fmt.Sprintf("Generation: %d\nAlive cells: %d", l.generation, l.liveCells)

	currentScreenWidth := int32(rl.GetScreenWidth())
	currentScreenHeight := int32(rl.GetScreenHeight())

	fontSize := max(int32(currentScreenHeight/20), 20)

	textPosX := int32(currentScreenWidth / 50)
	textPosY := int32(currentScreenHeight / 50)

	textColor := rl.White

	textWidth := rl.MeasureText(uiText, fontSize)

	padding := max(int32(fontSize/2), 10)

	bgPosX := textPosX - padding
	bgPosY := textPosY - padding
	bgWidth := textWidth + 2*padding
	bgHeight := fontSize*2 + padding*2
	bgColor := rl.NewColor(0, 0, 0, 150)

	if bgPosX < 0 {
		bgPosX = 0
	}
	if bgPosY < 0 {
		bgPosY = 0
	}
	if bgPosX+bgWidth > currentScreenWidth {
		bgWidth = currentScreenWidth - bgPosX
	}
	if bgPosY+bgHeight > currentScreenHeight {
		bgHeight = currentScreenHeight - bgPosY
	}

	rl.DrawRectangle(bgPosX, bgPosY, bgWidth, bgHeight, bgColor)
	rl.DrawText(uiText, textPosX, textPosY, fontSize, textColor)
}
