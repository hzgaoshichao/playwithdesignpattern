package main

import (
	"fmt"
	tp "templatemethod/testpaper"
)

func main() {
	fmt.Println("学生甲抄的试卷：")
	studentA := &tp.TestPaperA{}
	studentA.TemplateMethod(studentA)

	fmt.Println("学生乙抄的试卷：")
	studentB := &tp.TestPaperB{}
	studentB.TemplateMethod(studentB)
}

// 参考:
// https://learnku.com/articles/33704
// https://learnku.com/docs/the-way-to-go/104-tag-structure/3642
