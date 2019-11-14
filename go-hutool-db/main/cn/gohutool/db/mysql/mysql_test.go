package mysql

import (
	"github.com/huangbosbos/go-hutool/go-hutool-log/main/cn/gohutool/log"
	"testing"
)

// test case 的入参为 t *testing.T
// 如果是 test bench, 入参为 t *testing.B
func Test_findByPk(t *testing.T) {
	num := findByPk(1)
	t.Log(num)

	log.INFO.Print("should not panic")
}

//运行 b.N次
func Benchmark_findByPk(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		findByPk(1)
	}
}

// T.Errorf为打印错误信息，并且当前test case 会被跳过
// t.SkipNow()为跳过当前test，并且直接按PASS处理继续下一个test
