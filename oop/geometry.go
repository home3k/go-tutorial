package main

import (
	"math"
	"fmt"
)

type Point struct {
	X, Y float64
}

func Distance(q, p Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 方法的接收器 receiver, 向一个对象发送消息
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{10, 20}
	q := Point{4, 6}
	fmt.Println(Distance(q, p))
	fmt.Println(p.Distance(q))
}

