package assert

// Assert checks for a valid test, if not a panic is thrown.
func Assert(test bool, message string) {
	if !test {
		panic(message)
	}
}
