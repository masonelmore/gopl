package main

func main() {}

// dupAppend removes adjacent duplicates in a string slice
func dupAppend(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	for i := 0; i < len(strings)-1; {
		if strings[i] == strings[i+1] {
			strings = append(strings[:i], strings[i+1:]...)
			continue
		}
		i++
	}
	return strings
}

// dupCopy removes adjacent duplicates in a string slice
func dupCopy(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	for i := 0; i < len(strings)-1; {
		if strings[i] == strings[i+1] {
			copy(strings[i:], strings[i+1:])
			strings = strings[:len(strings)-1]
			continue
		}
		i++
	}
	return strings
}

// dupRunAppend removes adjacent duplicates in a string slice
func dupRunAppend(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	run := 0
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			run++
		} else if run > 0 {
			strings = append(strings[:i-run], strings[i:]...)
			i = i - run
			run = 0
		}
	}

	// The for loop will terminate before removing a run at the end of the
	// slice, so we need to remove it here.
	if run > 0 {
		strings = strings[:len(strings)-run]
	}

	return strings
}

// dupRunCopy removes adjacent duplicates in a string slice
func dupRunCopy(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	run := 0
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			run++
		} else if run > 0 {
			copy(strings[i-run:], strings[i:])
			strings = strings[:len(strings)-run]
			i = i - run
			run = 0
		}
	}

	// The for loop will terminate before removing a run at the end of the
	// slice, so we need to remove it here.
	if run > 0 {
		strings = strings[:len(strings)-run]
	}

	return strings
}
