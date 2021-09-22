package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{} //新建一个Image结构体

func (i Image) ColorModel() color.Model { //实现Image包中颜色模式的方法
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle { //实现Image包中生成图片边界的方法
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color { //实现Image包中生成图像某个点的方法
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {
	m := Image{}
	pic.ShowImage(m) //调用
}
