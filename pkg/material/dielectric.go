package material

import (
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
	"math"
	"math/rand"
)

type Dielectric struct {
	refIdx float64
}

func NewDielectric(refIdx float64) *Dielectric {
	return &Dielectric{refIdx: refIdx}
}

func (d *Dielectric) Scatter(inRay *mathplus.Ray, hitRecord geometry.HitInfo, attenuation *mathplus.Vector3, scatteredRay *mathplus.Ray) bool {
	*attenuation = *mathplus.NewVector3(1, 1, 1)
	directionNormalized := inRay.Direction().Normalize()
	cosTheta := math.Min(directionNormalized.Mul(-1).Dot(hitRecord.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	e := 1.0 / d.refIdx
	if !hitRecord.FrontFace {
		e = d.refIdx
	}

	if e*sinTheta > 1.0 {
		reflected := mathplus.Reflect(directionNormalized, hitRecord.Normal)
		*scatteredRay = *mathplus.NewRay(hitRecord.Point, reflected)
		return true
	}

	reflectProb := mathplus.Schlick(cosTheta, e)
	if rand.Float64() < reflectProb {
		reflected := mathplus.Reflect(directionNormalized, hitRecord.Normal)
		*scatteredRay = *mathplus.NewRay(hitRecord.Point, reflected)
		return true
	}

	directionRefracted := mathplus.Refract(directionNormalized, hitRecord.Normal, e)
	*scatteredRay = *mathplus.NewRay(hitRecord.Point, directionRefracted)
	return true
}
