package main

import w "workstatedemo/workstate"

func main() {
	emergencyProjects := w.NewWork()
	emergencyProjects.SetHour(9)
	emergencyProjects.WriteProgram()

	emergencyProjects.SetHour(10)
	emergencyProjects.WriteProgram()

	emergencyProjects.SetHour(12)
	emergencyProjects.WriteProgram()

	emergencyProjects.SetHour(13)
	emergencyProjects.WriteProgram()

	emergencyProjects.SetHour(14)
	emergencyProjects.WriteProgram()

	emergencyProjects.SetHour(17)
	emergencyProjects.SetWorkFinished(false)
	emergencyProjects.WriteProgram()
	emergencyProjects.SetHour(19)
	emergencyProjects.WriteProgram()
	emergencyProjects.SetHour(22)
	emergencyProjects.WriteProgram()
}

// 参考文档: https://learnku.com/docs/go-patterns/1.0.0/zhuang-tai-mo-shi-state-pattern/14755
