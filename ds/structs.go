package ds

type Rectangle struct {
	Width  float64
	Height float64
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}
