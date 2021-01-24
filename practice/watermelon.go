package kata

func Divide(weight int) bool {
	if weight == 2 {
		return false
	}

	if weight%2 == 0 {
		return true
	}

	return false
}
