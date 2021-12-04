package math_plus

import (
	"math"
	"math/rand"
)

const Pi = 3.1415926535897932385
const Float64Infinity = math.MaxFloat64

func RandomFloat64(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func RandomVector3() Vector3 {
	return Vector3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandomVector3InUnitDisk() Vector3 {
	for true {
		p := Vector3{RandomFloat64(-1, 1), RandomFloat64(-1, 1), 0}

		if p.LengthSqr() >= 1 {
			continue
		}

		return p
	}

	return Vector3{v1: 0, v2: 0, v3: 0}
}

func RandomUnitVector3() Vector3 {
	a := RandomFloat64(0, 2*Pi)
	z := RandomFloat64(-1, 1)
	r := math.Sqrt(1 - z*z)

	return Vector3{r * math.Cos(a), r * math.Sin(a), z}
}
