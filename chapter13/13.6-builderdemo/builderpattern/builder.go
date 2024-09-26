package builderpattern

type Builder interface {
	buildPartA()         //建造部件A
	buildPartB()         //建造部件B
	GetResult() *product //得到产品
}
