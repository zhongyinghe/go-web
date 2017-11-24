package gotest

import (
	"fmt"
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}

func Test_Division_3(t *testing.T) {
	var tests = []struct {
		x float64
		y float64
	}{
		{1.0, 3.0},
		{6.0, 2.0},
		{13.0, 5.0},
		{4.0, 0.0},
	}

	for _, tt := range tests {
		if _, e := Division(tt.x, tt.y); e != nil {
			t.Error("x为:" + fmt.Sprintf("v%", tt.x) + ",y为:" + fmt.Sprintf("v%", tt.y) + " 测试不通过!")
		} else {
			t.Log("x为:" + fmt.Sprintf("v%", tt.x) + ",y为:" + fmt.Sprintf("v%", tt.y) + " 测试ok!")
		}
	}
}
