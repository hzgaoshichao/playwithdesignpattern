package testpaper

type TestPaperA struct {
	testPaperCommon
}

func (t *TestPaperA) answer1() string {
	return "b"
}

func (t *TestPaperA) answer2() string {
	return "a"
}

func (t *TestPaperA) answer3() string {
	return "c"
}
