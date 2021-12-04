package camera

import mathplus "cmd/ray-tracing/main.go/pkg/math-plus"

type Camera struct {
	apertureRadius  float64
	u, v, w         mathplus.Vector3
	origin          mathplus.Vector3
	vertical        mathplus.Vector3
	horizontal      mathplus.Vector3
	lowerLeftConner mathplus.Vector3
}

func (c Camera) GetRay(d float64, t float64) {
}
