package calculator

import (
	"fmt"
	"math"
)

type Operation interface {
	GetResult(numberA float64, numberB float64) float64
}

type Add struct {
}

func (a *Add) GetResult(numberA float64, numberB float64) float64 {
	return numberA + numberB

}

type Sub struct {
}

func (a *Sub) GetResult(numberA float64, numberB float64) float64 {
	return numberA - numberB

}

type Mul struct {
}

func (a *Mul) GetResult(numberA float64, numberB float64) float64 {
	return numberA * numberB

}

type Div struct {
}

func (a *Div) GetResult(numberA float64, numberB float64) float64 {
	if numberB == 0 {
		fmt.Println("除数不能为0")
		panic("除数不能为0")
	}
	return numberA / numberB

}

// 指数运算, 求以 numberA 的 numberB 的次方
type Pow struct {
}

func (a *Pow) GetResult(numberA float64, numberB float64) float64 {
	// 此处省略两个参数的有效性检测
	return math.Pow(numberA, numberB)

}

// 对数运算, 求以 numberA 为底的 numberB 的对数
type Log struct {
}

func (a *Log) GetResult(numberA float64, numberB float64) float64 {
	// 此处省略两个参数的有效性检测
	return math.Log(numberB) / math.Log(numberA)

}
