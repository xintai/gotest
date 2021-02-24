package geometry 

import (
	"fmt"
//	"geometry/rectangle" // 导入自定义包
)

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v.Welcome!", name)
    return message
	//var rectLen, rectWidth float64 = 6, 7
	/*Area function of rectangle package used*/
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	//fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
