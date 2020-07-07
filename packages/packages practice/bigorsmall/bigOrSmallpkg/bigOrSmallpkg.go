package bigOrSmall

func Big(x int) bool {
	var response bool
	if x > 50 {
		response = true
	}

	return response
}

func Small(x int) bool { //if func is lowercase, its a private func
	var response bool
	if x < 50 {
		response = true
	}

	return response
}
