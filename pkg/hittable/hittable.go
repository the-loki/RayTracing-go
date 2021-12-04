package hittable

import "cmd/ray-tracing/main.go/pkg/ray"

type Hittable interface {
	Hit(ray *ray.Ray, dMin float64, dMax float64, hitRecord *HitRecord) bool
}
