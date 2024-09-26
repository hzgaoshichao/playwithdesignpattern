package main

import pb "personbuilderdemo/personbuilder"

func main() {
	person := pb.NewPersonBuilder().
		WithName("李鸿章").
		WithAge(20).
		WithGender("Male").
		WithAddress("中国南京").
		Build()

	person.Speak("精忠报国")
	person.Sleep()
}

// 参考文档：
// 这个例子更加的容易理解, 更加的优雅： https://github.com/ruanrunxue/Practice-Design-Pattern--Go-Implementation/blob/main/docs/go_practice_design_pattern__builder.md
