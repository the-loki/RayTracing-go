package geometry

import (
	"github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type Hittable interface {
	Hit(ray *mathplus.Ray, dMin float64, dMax float64, hitRecord *HitInfo) bool
}
