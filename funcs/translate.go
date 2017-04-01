package main

type point struct {
	x, y float64
}

type Path []point

func (p point) Add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func (p point) Sub(q point) point {
	return point{p.x - q.x, p.y - q.y}
}

func (path Path) Translate(offset point, add bool) {
	var op func(p, q point) point
	if add {
		op = point.Add
	} else {
		op = point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
