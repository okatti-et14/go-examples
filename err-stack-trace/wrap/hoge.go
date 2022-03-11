package wrap

func testCallWrap(err error) error {
	return Wrap(err)
}

func testCallWrap2nd(err error) error {
	return Wrap(err)
}
