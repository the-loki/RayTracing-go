package material

import (
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type Material interface {
	Scatter(inRay *mathplus.Ray, hitRecord geometry.HitInfo, attenuation *mathplus.Vector3, scatteredRay *mathplus.Ray) bool
}
