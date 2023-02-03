package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	i := GenernalRandNum(15000, 20000)
	fmt.Println(i)
}

//生成指定范围的随机数
func GenernalRandNum(min, max int32) int32 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	//创建随机种子
	rand.Seed(time.Now().Unix())
	//每次加上最小数
	return rand.Int31n(max-min) + min
}
