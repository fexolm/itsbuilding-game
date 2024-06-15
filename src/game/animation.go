package game

import graphics "github.com/quasilyte/ebitengine-graphics"

type AnimationMode uint8

const (
	AnimationForward AnimationMode = iota
	AnimationBackward
)

type Animation struct {
	sprite *graphics.Sprite

	repeated bool
	mode     AnimationMode

	frameWidth    uint16
	numFrames     uint16
	numAnimations int
	currentFrame  uint16

	frameTicker float64

	deltaPerFrame float64

	offsetY int
}

func (a *Animation) Rewind() {
	a.frameTicker = 0
	a.currentFrame = 0
	if a.mode == AnimationForward {
		a.sprite.SetFrameOffsetX(0)
	} else {
		a.sprite.SetFrameOffsetX(int(a.frameWidth) * int(a.numFrames-1))
	}
}

func (a *Animation) GetSprite() *graphics.Sprite {
	return a.sprite
}

func (a *Animation) SetSprite(s *graphics.Sprite, numFrames int) {
	a.sprite = s

	frameWidth := s.GetFrameWidth()
	if frameWidth == 0 {
		// This would work for a simple spritesheet
		// where frames have square shape and there is
		// only one "row" inside the spritesheet.
		frameWidth = s.ImageHeight()
	}

	if numFrames < 0 {
		// Auto-detection.
		numFrames = int(s.ImageWidth() / frameWidth)
	}

	a.frameWidth = uint16(frameWidth)
	a.numFrames = uint16(numFrames)
}

func (a *Animation) SetFPS(framesPerSecond float64) {
	a.deltaPerFrame = 1.0 / framesPerSecond
}

func (a *Animation) SetOffsetY(y int) {
	if a.offsetY != y {
		a.offsetY = y
		a.currentFrame = 0

		a.sprite.SetFrameOffsetY(y * a.sprite.ImageHeight() / a.numAnimations)

	}
}

func (a *Animation) Tick(delta float64) bool {
	if !a.repeated {
		if a.currentFrame > a.numFrames {
			return true
		}
	}

	finished := false
	a.frameTicker += delta
	framesAdvanced := int(a.frameTicker / a.deltaPerFrame)
	if framesAdvanced > 0 {
		a.frameTicker -= float64(framesAdvanced) * a.deltaPerFrame
		a.currentFrame += uint16(framesAdvanced)

		if a.currentFrame >= a.numFrames {
			if a.repeated {
				a.currentFrame = 0
			} else {
				finished = true
			}
		}

		frame := a.currentFrame
		if a.mode == AnimationBackward {
			frame = (a.numFrames - 1) - frame
		}

		a.sprite.SetFrameOffsetX(int(a.frameWidth) * int(frame))
	}

	return finished
}
