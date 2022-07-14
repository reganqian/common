package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
	// "fmt"
	// "math"
	"strconv"
	"fmt"
)

func AddTwoFloat(f1, f2 float64) float64 {
	df1  := decimal.NewFromFloat(f1)
	df2  := decimal.NewFromFloat(f2)
	res := df1.Add(df2)
	r1 , _ := res.Float64()
	return Float2Float(r1)
}

func SubTwoFloat(f1, f2 float64) float64 {
	df1  := decimal.NewFromFloat(f1)
	df2  := decimal.NewFromFloat(f2)
	res := df1.Sub(df2)
	r1 , _ := res.Float64()
	return Float2Float(r1)
}

func DivTwoFloat(f1, f2 float64) float64 {
	bigF1 := new(big.Float).SetFloat64(f1)
	bigF2 := new(big.Float).SetFloat64(f2)
	mul := new(big.Float).Quo(bigF1, bigF2)
	r1, _ := mul.Float64()
	return Float2Float(r1)
}


func MulTwoFloat(f1, f2 float64) float64 {
	bigF1 := new(big.Float).SetFloat64(f1)
	bigF2 := new(big.Float).SetFloat64(f2)
	mul := new(big.Float).Mul(bigF1, bigF2)
	r1, _ := mul.Float64()
	return Float2Float(r1)
}

func Float2Float(num float64) float64 {
    floatNum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
    return floatNum
}



func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}