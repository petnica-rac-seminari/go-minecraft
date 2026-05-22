package navigation

type JumpInput struct {
	lastSpacePress float64
}

func (j *JumpInput) TryDoubleTap(now float64, spacePressed, canJump bool) bool {
	if !spacePressed {
		return false
	}
	if j.lastSpacePress > 0 && now-j.lastSpacePress <= DoubleTapWindowSec && canJump {
		j.lastSpacePress = 0
		return true
	}
	j.lastSpacePress = now
	return false
}

func TryDoubleTapJump(j *JumpInput, now float64, spacePressed, canJump bool) bool {
	return j.TryDoubleTap(now, spacePressed, canJump)
}
