package rectangle

// Define a struct type
type Rectangle struct {
	Width  float64
	Height float64
}

// Embedding for composition
type ColoredRectangle struct {
    Rectangle  // Embedding Rectangle
    Color string
}

// Method on Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func New(width, height float64) Rectangle {
	return Rectangle{Width: width, Height: height}
}


func NewColoredRectangle(width, height float64, color string) ColoredRectangle {
	return ColoredRectangle{
		Rectangle: Rectangle{Width: width, Height: height},
        Color:    color,
	}
}