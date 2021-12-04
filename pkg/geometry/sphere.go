package geometry

import (
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
	"math"
)

type Sphere struct {
	radius   float64
	center   mathplus.Vector3
	material interface{}
}

func NewSphere(radius float64, center mathplus.Vector3, material interface{}) *Sphere {
	return &Sphere{radius: radius, center: center, material: material}
}

func (s *Sphere) Hit(ray *mathplus.Ray, dMin float64, dMax float64, hitInfo *HitInfo) bool {
	co := ray.Origin().Sub(s.center)
	a := ray.Direction().LengthSqr()
	hb := ray.Direction().Dot(co)
	c := co.LengthSqr() - s.radius*s.radius
	discriminant := hb*hb - a*c

	setHitRecord := func(distance float64) {
		*hitInfo = HitInfo{}
		hitInfo.Distance = distance
		hitInfo.Point = ray.At(distance)
		normal := hitInfo.Point.Sub(s.center).Mul(1.0 / s.radius)
		hitInfo.SetFaceNormal(ray, normal)
		hitInfo.Material = s.material
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
