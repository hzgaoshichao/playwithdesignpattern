package main

import rc "deepcopy/resumecopy"

func main() {
	resume1 := rc.NewResume("大鸟")
	resume1.SetPersonalInfo("男", "29")
	resume1.SetWorkExperience("1998-2000", "xx公司")

	resume2 := resume1.Clone().(*rc.Resume) // 这里使用了GoLang类型断言, 详细可参考: https://juejin.cn/post/7136359125738815524
	resume2.SetWorkExperience("2000-2003", "yy集团")

	resume3 := resume1.Clone().(*rc.Resume)
	resume3.SetWorkExperience("2003-2006", "zz公司")

	resume1.Display()
	resume2.Display()
	resume3.Display()

	// resume4 := resume1
	// resume4.SetWorkExperience("2007-2008", "躺平")
	// resume4.Display()

}
