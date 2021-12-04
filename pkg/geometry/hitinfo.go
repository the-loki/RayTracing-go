package geometry

import (
	mathplus "github.com/404Polaris/RayTracing-go/pkg/mathplus"
)

type HitInfo struct {
	Distance  float64
	FrontFace bool
	Point     mathplus.Vector3
	Normal    mathplus.Vector3
	Material  interface{}
}

func (h *HitInfo) SetFaceNormal(ray *mathplus.Ray, outWardNormal mathplus.Vector3) {
	h.FrontFace = ray.Direction().Dot(outWardNormal) < 0

	if h.FrontFace {
		h.Normal = outWardNormal
	} else {
		h.Normal = outWardNormal.Mul(-1.0)
	}
}
