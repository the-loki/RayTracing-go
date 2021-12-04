package material

import (
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type Lambertian struct {
	albedo mathplus.Vector3
}

func NewLambertian(albedo mathplus.Vector3) *Lambertian {
	return &Lambertian{albedo: albedo}
}

func (l *Lambertian) Scatter(inRay *mathplus.Ray, hitRecord geometry.HitInfo, attenuation *mathplus.Vector3, scatteredRay *mathplus.Ray) bool {
	scatterDirection := hitRecord.Normal.Add(mathplus.RandomUnitVector3())
	*scatteredRay = *mathplus.NewRay(hitRecord.Point, scatterDirection.Normalize())
	*attenuation = l.albedo
	return true
}
