package scene

import (
	"cmd/ray-tracing/main.go/pkg/hittable"
	"cmd/ray-tracing/main.go/pkg/ray"
)

type Scene struct {
	objects []*hittable.Hittable
}

func (s Scene) Hit(ray *ray.Ray, dMin float64, dMax float64, hitRecord *hittable.HitRecord) bool {
	hitAnything := false
	closetDistMax := dMax
	record := hittable.HitRecord{}

	for _, o := range s.objects {
		if (*o).Hit(ray, dMin, closetDistMax, &record) {
			hitAnything = true
			hitRecord = &record
			closetDistMax = record.Distance
		}
	}

	return hitAnything
}

func (s Scene) Clear() {
	s.objects = []*hittable.Hittable{}
}

func (s Scene) Add(object *hittable.Hittable) {
	s.objects = append(s.objects, object)
}
