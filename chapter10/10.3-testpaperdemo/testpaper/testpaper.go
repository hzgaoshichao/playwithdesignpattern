package testpaper

import "fmt"

type testPaper interface {
	testQuestion1(p testPaper)
	answer1() string
	testQuestion2(p testPaper)
	answer2() string
	testQuestion3(p testPaper)
	answer3() string
}

type testPaperCommon struct {
}

func (t *testPaperCommon) testQuestion1(p testPaper) {
	fmt.Printf("Question 1:杨过得到，后来给了郭靖，炼成倚天剑、屠龙刀的玄铁可能是[ ] : a.球磨铸铁 b.马口铁 c.高速合金钢 d.碳素纤维 \n")
	fmt.Printf("答案: %v \n", p.answer1())
}

func (t *testPaperCommon) testQuestion2(p testPaper) {
	fmt.Printf("Question 2:杨过、程英、陆无双铲除了情花，造成[ ] : a.使这种植物不再害人 b.使一种珍稀物种灭绝 c.破坏了那个生物圈的生态平衡 d.造成该地区沙漠化 \n")
	fmt.Printf("答案: %v \n", p.answer2())
}

func (t *testPaperCommon) testQuestion3(p testPaper) {
	fmt.Printf("Question 3:蓝凤凰致使华山师徒、桃谷六仙呕吐不止,如果你是大夫,会给他们开什么药[ ] : a.阿司匹林 b.牛黄解毒片 c.氟哌酸 d.让他们喝大量的生牛奶 e.以上全不对 \n")
	fmt.Printf("答案: %v \n", p.answer3())
}

func TemplateMethod(t testPaper) {
	t.testQuestion1(t)
	t.testQuestion2(t)
	t.testQuestion3(t)
}
