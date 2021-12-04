package ray

import (
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
)

type Ray struct {
	origin    mathplus.Vector3
	direction mathplus.Vector3
}

func NewRay(origin mathplus.Vector3, direction mathplus.Vector3) *Ray {
	return &Ray{origin: origin, direction: direction}
}

func (r Ray) Origin() mathplus.Vector3 {
	return r.origin
}

func (r Ray) Direction() mathplus.Vector3 {
	return r.direction
}

func (r Ray) At(length float64) mathplus.Vector3 {
	return r.origin.Add(r.direction.Mul(length))
}
