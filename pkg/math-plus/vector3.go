package math_plus

import "math"

type Vector3 struct {
	v1, v2, v3 float64
}

func NewVector3(v1 float64, v2 float64, v3 float64) *Vector3 {
	return &Vector3{v1: v1, v2: v2, v3: v3}
}

func (v Vector3) X() float64 {
	return v.v1
}

func (v Vector3) Y() float64 {
	return v.v2
}

func (v Vector3) Z() float64 {
	return v.v3
}

func (v Vector3) R() float64 {
	return v.v1
}

func (v Vector3) G() float64 {
	return v.v2
}

func (v Vector3) B() float64 {
	return v.v3
}

func (v Vector3) MulVector(u Vector3) Vector3 {
	return Vector3{v.v1 * u.v1, v.v2 * u.v2, v.v3 * u.v3}
}

func (v Vector3) Mul(u float64) Vector3 {
	return Vector3{v.v1 * u, v.v2 * u, v.v3 * u}
}

func (v Vector3) Add(u Vector3) Vector3 {
	return Vector3{v.v1 + u.v1, v.v2 + u.v2, v.v3 + u.v3}
}

func (v Vector3) Sub(u Vector3) Vector3 {
	return v.Add(u.Mul(-1.0))
}

func (v Vector3) Div(u float64) Vector3 {
	return v.Mul(1.0 / u)
}

func (v Vector3) Dot(u Vector3) float64 {
	return v.v1*u.v1 + v.v2*u.v2 + v.v3*u.v3
}

func (v Vector3) Cross(u Vector3) Vector3 {
	return Vector3{v.v2*u.v3 - v.v3*u.v2, v.v3*u.v1 - v.v1*u.v3, v.v1*u.v2 - v.v2*u.v1}
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.LengthSqr())
}

func (v Vector3) LengthSqr() float64 {
	return v.v1*v.v1 + v.v2*v.v2 + v.v3*v.v3
}

func (v Vector3) Normalize() Vector3 {
	return v.Div(v.Length())
}
