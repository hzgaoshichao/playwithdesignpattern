package calculator

type Factory interface {
	CreateOperate(oper string) Operation
}

type FactoryBasic struct {
}

func (f *FactoryBasic) CreateOperate(opertype string) Operation {
	var oper Operation
	switch opertype {
	case "+":
		oper = &Add{}
	case "-":
		oper = &Sub{}
	case "*":
		oper = &Mul{}
	case "/":
		oper = &Div{}

	}
	return oper
}

type FactoryAdvanced struct {
}

func (f *FactoryAdvanced) CreateOperate(opertype string) Operation {
	var oper Operation
	switch opertype {
	case "pow":
		oper = new(Pow)
	case "log":
		oper = new(Log)

	}
	return oper
}

type OperationFactory struct {
}

func (f *OperationFactory) CreateOperate(opertype string) Operation {
	var factory Factory
	switch opertype {
	case "+":
		fallthrough
	case "-":
		fallthrough
	case "*":
		fallthrough
	case "/":
		factory = new(FactoryBasic)
	case "pow":
		fallthrough
	case "log":
		factory = new(FactoryAdvanced)

	}

	oper := factory.CreateOperate(opertype)
	return oper
}
