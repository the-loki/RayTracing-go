package material

import (
	"cmd/ray-tracing/main.go/pkg/hittable"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
)

type Metal struct {
	fuzz   float64
	albedo mathplus.Vector3
}

func NewMetal(fuzz float64, albedo mathplus.Vector3) *Metal {
	return &Metal{fuzz: fuzz, albedo: albedo}
}

func (m Metal) scatter(inRay *ray.Ray, hitRecord *hittable.HitRecord, attenuation *mathplus.Vector3, scatteredRay *ray.Ray) bool {
	reflectDirection := mathplus.Reflect(inRay.Direction().Normalize(), hitRecord.Normal)
	scatteredRay = ray.NewRay(hitRecord.Point, reflectDirection.Add(mathplus.RandomUnitVector3()).Mul(m.fuzz).Normalize())
	attenuation = &m.albedo

	return scatteredRay.Direction().Dot(hitRecord.Normal) > 0
}
