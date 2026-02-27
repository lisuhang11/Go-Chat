package random

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 生成一个指定位数 len 的随机整数
func GetRandomInt(len int) int {
	return rand.Intn(9*int(math.Pow(10, float64(len-1)))) + int(math.Pow(10, float64(len-1)))
}

// 返回一个字符串，由当前日期（格式 20060102，如 20250226）和 len 位的随机整数拼接而成
func GetNowAndLenRandomString(len int) string {
	return time.Now().Format("20060102") + strconv.Itoa(GetRandomInt(len))
}
