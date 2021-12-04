package mathplus

import (
	"math"
	"math/rand"
)

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomFloat64InRange(min float64, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func RandomVector3() Vector3 {
	return Vector3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandomVector3InUnitDisk() Vector3 {
	for true {
		p := Vector3{RandomFloat64InRange(-1, 1), RandomFloat64InRange(-1, 1), 0}

		if p.LengthSqr() >= 1 {
			continue
		}

		return p
	}

	return Vector3{v1: 0, v2: 0, v3: 0}
}

func RandomUnitVector3() Vector3 {
	a := RandomFloat64InRange(0, 2*math.Pi)
	z := RandomFloat64InRange(-1, 1)
	r := math.Sqrt(1 - z*z)

	return Vector3{r * math.Cos(a), r * math.Sin(a), z}
}
