package main

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	// 匿名成员
	Circle
	Spokes int
}

func main()  {
	var w Wheel
	w.X = 100
	w.Y = 200
	w.Radius = 50
	w.Spokes = 34
}