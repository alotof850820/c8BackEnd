package validate

import "time"

// bool
// 數字
// string
// slice
// map
// time

func SingleFieldValidate() {
	v := validate
	var err error

	var b bool
	err = v.Var(b, "boolean") // b必須為布林
	outRes("boolean", &err)

	var i = "100"
	err = v.Var(i, "number") // i必須為數字
	outRes("number", &err)

	var f = "100.123"
	err = v.Var(f, "numeric") // f必須為浮點型 需使用numeric 除非為非字串
	outRes("numeric", &err)

	var s = "abc"
	err = v.Var(s, "alpha") // s必須為字串
	outRes("alpha", &err)

	var sl = []int{1, 2, 3}
	err = v.Var(sl, "max=15,min=2") // sl長度必須小於15 大於2
	outRes("slice", &err)

	var mp = map[int]int{1: 1, 2: 2, 3: 3}
	err = v.Var(mp, "max=15,min=2") // mp長度必須小於15 大於2
	outRes("map", &err)

	var t = time.Now().Format("2006-01-02 15:04:05")
	err = v.Var(t, "datetime=2006-01-02 15:04:05") // t必須符合2006-01-02 15:04:05格式
	outRes("time", &err)

	s1 := "asd"
	s2 := "asd"
	err = v.VarWithValue(s1, s2, "eqfield") // s1必須等於s2
	outRes("eqfield", &err)

	i1 := 30
	i2 := 20
	err = v.VarWithValue(i1, i2, "gtfield") // i1必須大於i2 gtfield大於 ltfield小於 eqfield等於
	outRes("gtfield", &err)
}
