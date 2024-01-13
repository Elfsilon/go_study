package basics

import (
	"fmt"
	"math"
)

type Vertex2D struct {
	X float64
	Y float64
}

// Method with value receiver
// Here v is the copy of the original vertex, so it can't be modified
//
// If it calls like v.Abs(...) - it's automatically interpreted by Golang
// as (*v).Abs(...) unlike function variant
func (v Vertex2D) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method with pointer receiver
// Here v is the pointer to the original vertex, so it can be modified
//
// Opposite to value receivers:
// If it calls like v.Scale(...) - it's automatically interpreted by Golang
// as (&v).Scale(...) unlike function variant
//
// This type of receiver is preferred:
// 1. Struct doesn't copying - it's more efficient
// 2. Value can be modified
func (v *Vertex2D) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// The same is above but a function:
// This func can be called as Scale(&v, 5)
// If it calls without pointer: Scale(v, 5) - compile error        (!)
func Scale(v *Vertex2D, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex2D) String() string {
	return fmt.Sprintf("Vertex2D(x: %f, y: %f)", v.X, v.Y)
}

// Composition instead of Inheritance
type Vertex3D struct {
	Z     float64
	Point *Vertex2D
}

func NewVertex3D(x float64, y float64, z float64) Vertex3D {
	return Vertex3D{z, &Vertex2D{x, y}}
}

func (v Vertex3D) Abs() float64 {
	return math.Sqrt(v.Point.X*v.Point.X + v.Point.Y*v.Point.Y + v.Z*v.Z)
}

func (v *Vertex3D) Scale(f float64) {
	v.Point.X = v.Point.X * f
	v.Point.Y = v.Point.Y * f
	v.Z = v.Z * f
}

func (v Vertex3D) String() string {
	return fmt.Sprintf("Vertex3D(x: %f, y: %f, z: %f)", v.Point.X, v.Point.Y, v.Z)
}

func Main() {
	fmt.Printf("----- STRUCTS AND METHODS -----\n")

	v := Vertex2D{1, 2}
	v.X = -2
	fmt.Println(v)

	// Access via pointer
	// In this case we should have written (*p).X, but Go let us get is easier
	p := &v
	fmt.Printf("Vertex.X via pointer is: %v\n", p.X)

	abs := v.Abs()
	fmt.Printf("Abs of the vertex is: %f\n", abs)

	// The same is (&v).Scale(5)
	v.Scale(5)
	fmt.Println(v)

	v3 := NewVertex3D(1, 2, 3)
	fmt.Println(v3)

	v3.Scale(5)
	fmt.Println(v3)

	v3.Point.Scale(5)
	fmt.Println(v3)
}
