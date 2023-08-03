package add

import "testing"

// 测试函数
func TestMyadd(t *testing.T) {
	ret := Myadd(1, 2)
	want := 3
	if ret != want {
		t.Errorf("error!\n")
	}
}

// 性能基准函数
func BenchmarkMyadd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Myadd(1, 2)
	}
}
