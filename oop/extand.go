package main

import "image/color"

type point struct {
	X,Y float64
}

type ColoredPoint struct {
	point
	X float64
	Color color.RGBA
}




