package testpaper

type TestPaperB struct {
	testPaperCommon
}

func (t *TestPaperB) answer1() string {
	return "d"
}

func (t *TestPaperB) answer2() string {
	return "b"
}

func (t *TestPaperB) answer3() string {
	return "a"
}
