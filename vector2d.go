package vector

import (
	"fmt"
	"math"
)

// Vector2d is a two dimensional vector. It has an X and Y component, and several methods
// for manipulating the vector.
type Vector2d struct {
	X, Y float64
}

// Vector addition
func (vec1 Vector2d) AddVector(vec2 Vector2d) Vector2d {
	return Vector2d{X: vec1.X + vec2.X, Y: vec1.Y + vec2.Y}
}

// Vector subtraction
func (vec1 Vector2d) SubVector(vec2 Vector2d) Vector2d {
	return Vector2d{X: vec1.X - vec2.X, Y: vec1.Y - vec2.Y}
}

// Vector dot product
func (vec1 Vector2d) Dot(vec2 Vector2d) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}

// Scalar multiplication with vector
func (vec Vector2d) Scale(scalar float64) Vector2d {
	return Vector2d{X: vec.X * scalar, Y: vec.Y * scalar}
}

// Add scalar to vector
func (vec Vector2d) AddScalar(scalar float64) Vector2d {
	return Vector2d{X: vec.X + scalar, Y: vec.Y + scalar}
}

// Subract scalar from vector
func (vec Vector2d) SubScalar(scalar float64) Vector2d {
	return Vector2d{X: vec.X - scalar, Y: vec.Y - scalar}
}

// Subtract scalar from vector, where scalar is to the right of the vector
func (vec Vector2d) SubVectorRight(scalar float64) Vector2d {
	return Vector2d{X: scalar - vec.X, Y: scalar - vec.Y}
}

// Divide a vector by a scalar
func (vec Vector2d) Divide(scalar float64) Vector2d {
	return Vector2d{X: vec.X / scalar, Y: vec.Y / scalar}
}

// Find the magnitude of a vector. This is the same as the length of the vector.
func (vec Vector2d) Magnitude() float64 {
	return math.Sqrt(vec.Dot(vec))
}

// Get the unit vector of a vector. This is basically the direction of the vector.
func (vec Vector2d) Unit() Vector2d {
	return vec.Divide(vec.Magnitude())
}

// Rotate a vector by an angle in radians
func (vec Vector2d) Rotate(angle float64) Vector2d {
	return Vector2d{
		X: vec.X*math.Cos(angle) - vec.Y*math.Sin(angle),
		Y: vec.X*math.Sin(angle) + vec.Y*math.Cos(angle),
	}
}

// Get the angle of a vector in radians
func (vec Vector2d) Angle() float64 {
	return math.Atan2(vec.Y, vec.X)
}

// Find the maximum component of a vector
func (vec Vector2d) MaxComponent() float64 {
	if vec.X > vec.Y {
		return vec.X
	}
	return vec.Y
}

// Find the minimum component of a vector.
func (vec Vector2d) MinComponent() float64 {
	if vec.X < vec.Y {
		return vec.X
	}
	return vec.Y
}

// Get the absolute value of a vector.
func (vec Vector2d) Abs() Vector2d {
	return Vector2d{X: math.Abs(vec.X), Y: math.Abs(vec.Y)}
}

// Negate a vector.
func (vec Vector2d) Negate() Vector2d {
	return Vector2d{X: -vec.X, Y: -vec.Y}
}

// Project a vector onto another vector.
func (vec Vector2d) Project(vec2 Vector2d) Vector2d {
	return vec2.Scale(vec.Dot(vec2) / vec2.Dot(vec2))
}

// Get the string representation of a vector.
func (vec Vector2d) Repr() string {
	components := []interface{}{vec.X, vec.Y}
	res := fmt.Sprintf("Vector2d{X: %v, Y:%v}", components...)
	return res
}
