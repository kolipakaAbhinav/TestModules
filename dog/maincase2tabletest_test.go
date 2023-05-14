package dog

import "testing"

func Test2ars(t *testing.T) {

	type test struct {
		data int
		res  int
	}

	cases := []test{
		{1, 10},
		{-1, -10},
		{2, 20},
		{100, 1000},
		{2, 20},
		{11, 110},
		{12, 120},
	}

	for _, v := range cases {
		op := Years(v.data)
		if op != v.res {
			t.Error("Expected : ", v.res, " got ", op)
		}
	}

}
