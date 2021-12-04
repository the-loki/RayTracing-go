package material

import (
	"cmd/ray-tracing/main.go/pkg/hittable"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
)

type Material interface {
	scatter(inRay *ray.Ray, hitRecord *hittable.HitRecord, attenuation *mathplus.Vector3, scatteredRay *ray.Ray) bool
}
