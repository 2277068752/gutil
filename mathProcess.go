/**
随机数处理
create by gloomy 2017-03-29 22:11:23
*/
package gutil

import (
	"github.com/gobwas/glob"
	"math"
)

/**
四舍五入取舍
create by gloomy 2017-03-29 22:11:18
需要取舍的浮点数 取舍几位
*/
func Rounding(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

/**
四舍五入取舍
create by gloomy 2017-03-29 22:18:46
除数 被除数 取舍几位
*/
func RoundingByInt(number, subNumber, n int) float64 {
	f := float64(number) / float64(subNumber)
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

/**
四舍五入取舍 百分比
create by gloomy 2017-03-29 22:18:46
除数 被除数 取舍几位
*/
func RoundingPercentageByInt(number, subNumber, n int) float64 {
	f := float64(number) / float64(subNumber)
	n += 2
	pow10_n := math.Pow10(n)
	percentage := math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
	if math.IsInf(percentage, 0) || math.IsNaN(percentage) {
		percentage = 0
	}
	return percentage * 100
}

//匹配规则是否存在
//create by gloomy 2017-09-07 21:13:30
func MustCompileMatch(matchStr, str string) bool {
	return glob.MustCompile(matchStr).Match(str)
}
