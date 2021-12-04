package material

import (
	"cmd/ray-tracing/main.go/pkg/hittable"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
)

type Lambertian struct {
	albedo mathplus.Vector3
}

func (l Lambertian) scatter(inRay *ray.Ray, hitRecord *hittable.HitRecord, attenuation *mathplus.Vector3, scatteredRay *ray.Ray) bool {
	scatterDirection := hitRecord.Normal.Add(mathplus.RandomUnitVector3())
	scatteredRay = ray.NewRay(hitRecord.Point, scatterDirection.Normalize())
	attenuation = &l.albedo
	return true
}
