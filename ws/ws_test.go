package ws

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	count,total, last :=  int32(10000),int32(20000000), int32(20000000)
	resp, used := make(map[int32]int32), make(map[int32]int32)

	for i := int32(0); i < count; i++ {
		num := rand.Int31n(total - i) //随机出来的下标
		if restore, ok := used[num]; ok { //下标被使用过，取出存入的值
			resp[restore] = restore
		} else { //下标没有被用过，下标和值相等
			resp[num] = num
		}
		if num == last { //取值刚好为最后一个，找倒数第二个
			last--
		}
		used[num] = last //把最后一个值和已经用过的下标存起来
		last--//缩小范围
	}
	fmt.Println(len(resp))
}
