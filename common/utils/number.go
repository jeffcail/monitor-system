package utils

import (
	"github.com/shopspring/decimal"
)

func StringToFloat(s string) float64 {
	temp, _ := decimal.NewFromString(s)
	result, _ := temp.Float64()
	return result
}

func StringToInt(s string) int64 {
	temp, _ := decimal.NewFromString(s)
	result := temp.IntPart()
	return result

}

func Sub(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Sub(db).String()
}

func SubNotFix(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Sub(db).String()
}

func Div(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Div(db).StringFixed(2)
}

func Div2(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Div(db).StringFixed(4)
}

func Add(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Add(db).String()
}

func Mul(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).String()
}

func MulFix(a string, b string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).StringFixed(fix)
}

func GE(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.GreaterThanOrEqual(db)
}

func LessThen(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.LessThan(db)
}

func LessThenAndEqual(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.LessThanOrEqual(db)
}

func Greater(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.GreaterThan(db)
}

func Equal(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Equal(db)
}

func MulUnFixed(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).String()
}

func Fix(a string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	return da.StringFixed(fix)
}

// In
func In(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}
