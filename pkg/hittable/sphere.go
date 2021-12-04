package hittable

import (
	"cmd/ray-tracing/main.go/pkg/material"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
	"math"
)

type Sphere struct {
	radius   float64
	center   mathplus.Vector3
	material *material.Material
}

func (s Sphere) Hit(ray *ray.Ray, dMin float64, dMax float64, hitRecord *HitRecord) bool {
	co := ray.Origin().Sub(s.center)
	a := ray.Direction().LengthSqr()
	hb := ray.Direction().Dot(co)
	c := co.LengthSqr() - s.radius*s.radius
	discriminant := hb - hb - a*c

	setHitRecord := func(distance float64) {
		hitRecord.Distance = distance
		hitRecord.Point = ray.At(distance)
		normal := hitRecord.Point.Sub(s.center).Mul(1.0 / s.radius)
		hitRecord.SetFaceNormal(ray, normal)
		hitRecord.Material = s.material
	}

	if discriminant <= 0 {
		return false
	}

	sqrD := math.Sqrt(discriminant)
	distance := (-hb - sqrD) / a
	if distance > dMin && distance < dMax {
		setHitRecord(distance)
		return true
	}

	distance = (-hb + sqrD) / a
	if distance > dMin && distance < dMax {
		setHitRecord(distance)
		return true
	}

	return false
}
