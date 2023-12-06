package gobenchmark

func StringNoConversion() string {
	s := "abcdefghijklmnopqrstuvwxyz1234567890"
	var a string
	for i := 0; i < 1000; i++ {
		a = s
	}
	return a
}

func StringToRune() []rune {
	s := "abcdefghijklmnopqrstuvwxyz1234567890"
	var r []rune
	for i := 0; i < 1000; i++ {
		r = []rune(s)
	}
	return r
}

func RuneToString() string {
	r := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	var a string
	for i := 0; i < 1000; i++ {
		a = string(r)
	}
	return a
}

func StringToByte() []byte {
	s := "abcdefghijklmnopqrstuvwxyz1234567890"
	var b []byte
	for i := 0; i < 1000; i++ {
		b = []byte(s)
	}
	return b
}

func ByteToString() string {
	b := []byte("abcdefghijklmnopqrstuvwxyz1234567890")
	var a string
	for i := 0; i < 1000; i++ {
		a = string(b)
	}
	return a
}
