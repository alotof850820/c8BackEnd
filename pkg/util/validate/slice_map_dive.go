package validate

type sliceStruct struct {
	OpCode int    `v:"eq=1|eq=2"`
	Op     string `v:"required"`
}

func SliceValidate() {
	v := validate
	slice1 := []string{"123456", "456789", "789456123"}
	err := v.Var(slice1, "gte=3,dive,required,gte=5,lte=10,number") // slice1長度必須大於等於3 且每個元素必須為數字
	outRes("slice1", &err)

	slice2 := [][]string{
		{"123456", "456789", "789456123"},
		{"12345", "45678", "78945612"},
		{"12345", "45675", "7894561"},
	}
	// 2個dive意味着数据会被深入处理两次
	err = v.Var(slice2, "gte=3,dive,gte=3,dive,required,gte=5,lte=10,number") // slice2長度必須大於等於3 且每個元素必須為數字
	outRes("slice2", &err)

	slice3 := []*sliceStruct{
		{OpCode: 1, Op: "切片操作"},
		{OpCode: 2, Op: "切片操作"},
		{OpCode: 3, Op: "切片操作"},
	}
	err = v.Var(slice3, "gte=2,dive") // slice3長度必須大於等於2
	outRes("slice3", &err)
}

func MapValidate() {
	var err error
	var v = validate
	mp1 := map[string]string{
		"A": "12345",
		"B": "12345",
		"C": "12345",
	}

	// endkeys 对keys的验证结束
	err = v.Var(mp1, "gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number") // mp1長度必須大於等於3
	outRes("mp1", &err)

	mp2 := map[string]map[string]string{
		"A": {
			"A": "12345",
			"B": "12345",
			"C": "12345",
		},
		"B": {
			"A": "12345",
			"B": "12345",
			"C": "12345",
		},
		"C": {
			"A": "12345",
			"B": "12345",
			"C": "12345",
		},
	}
	err = v.Var(mp2, "gte=1,dive,keys,len=1,alpha,endkeys,required,gte=3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number") // mp2長度必須大於等於3
	outRes("mp2", &err)
}
