package main

type data struct {

}

func (data) string() string  {
	return "data string"
}

type node struct {
	data interface{ // 匿名接口类型
		string() string
	}
}

func main() {
	var t interface{
		string() string
	} = data{}

	n := node{
		data: t,
	}

	println(n.data.string())
}