package mathplus

import "math"

func Clamp(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}

func Reflect(in Vector3, n Vector3) Vector3 {
	return in.Sub(n.Mul(in.Dot(n) * 2))
}

func Refract(in Vector3, n Vector3, e float64) Vector3 {
	cosTheta := in.Mul(-1).Dot(n)
	l1 := in.Add(n.Mul(cosTheta)).Mul(e)
	l2 := n.Mul(-1.0 * math.Sqrt(math.Abs(1.0-l1.LengthSqr())))
	return l1.Add(l2)
}

func Schlick(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
