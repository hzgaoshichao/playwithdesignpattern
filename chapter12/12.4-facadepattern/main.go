package main

import f "facadepattern/facade"

func main() {
	facade := f.NewFacade()
	facade.MethodA()
	facade.MethodB()
}

// 参考文档: https://learnku.com/docs/go-patterns/1.0.0/wai-guan-mo-shi-facade-pattern/14751
