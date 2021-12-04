package material

import (
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type Metal struct {
	fuzz   float64
	albedo mathplus.Vector3
}

func NewMetal(fuzz float64, albedo mathplus.Vector3) *Metal {
	fuzz = mathplus.Clamp(fuzz, fuzz, 1)
	return &Metal{fuzz: fuzz, albedo: albedo}
}

func (m *Metal) Scatter(inRay *mathplus.Ray, hitInfo geometry.HitInfo, attenuation *mathplus.Vector3, scatteredRay *mathplus.Ray) bool {
	reflectDirection := mathplus.Reflect(inRay.Direction().Normalize(), hitInfo.Normal)
	*scatteredRay = *mathplus.NewRay(hitInfo.Point, reflectDirection.Add((mathplus.RandomUnitVector3()).Normalize().Mul(m.fuzz)))
	*attenuation = m.albedo

	return scatteredRay.Direction().Dot(hitInfo.Normal) > 0
}
