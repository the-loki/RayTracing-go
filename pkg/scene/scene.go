package scene

import (
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	"github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type Scene struct {
	objects []geometry.Hittable
}

func NewScene() *Scene {
	scene := &Scene{}
	scene.objects = make([]geometry.Hittable, 0)
	return scene
}

func (s *Scene) Hit(ray *mathplus.Ray, dMin float64, dMax float64, hitInfo *geometry.HitInfo) bool {
	hitAnything := false
	closetDistMax := dMax
	localHitInfo := geometry.HitInfo{}

	for _, o := range s.objects {
		if o.Hit(ray, dMin, closetDistMax, &localHitInfo) {
			hitAnything = true
			*hitInfo = localHitInfo
			closetDistMax = localHitInfo.Distance
		}
	}

	return hitAnything
}

func (s *Scene) Clear() {
	s.objects = []geometry.Hittable{}
}

func (s *Scene) Add(object geometry.Hittable) {
	s.objects = append(s.objects, object)
}
