package material

import (
	"cmd/ray-tracing/main.go/pkg/hittable"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
	"math"
	"math/rand"
)

type Dielectric struct {
	refIdx float64
}

func NewDielectric(refIdx float64) *Dielectric {
	return &Dielectric{refIdx: refIdx}
}

func (d Dielectric) scatter(inRay *ray.Ray, hitRecord *hittable.HitRecord, attenuation *mathplus.Vector3, scatteredRay *ray.Ray) bool {
	attenuation = mathplus.NewVector3(1, 1, 1)
	directionNormalized := inRay.Direction().Normalize()
	cosTheta := math.Min(directionNormalized.Mul(-1).Dot(hitRecord.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	e := 1.0 / d.refIdx
	if !hitRecord.FrontFace {
		e = d.refIdx
	}

	if e*sinTheta > 1.0 {
		reflected := mathplus.Reflect(directionNormalized, hitRecord.Normal)
		scatteredRay = ray.NewRay(hitRecord.Point, reflected)
		return true
	}

	reflectProb := mathplus.Schlick(cosTheta, e)
	if rand.Float64() < reflectProb {
		reflected := mathplus.Reflect(directionNormalized, hitRecord.Normal)
		scatteredRay = ray.NewRay(hitRecord.Point, reflected)
		return true
	}

	directionRefracted := mathplus.Refract(directionNormalized, hitRecord.Normal, e)
	scatteredRay = ray.NewRay(hitRecord.Point, directionRefracted)
	return true
}
