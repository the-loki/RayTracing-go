package hittable

import (
	"cmd/ray-tracing/main.go/pkg/material"
	mathplus "cmd/ray-tracing/main.go/pkg/math-plus"
	"cmd/ray-tracing/main.go/pkg/ray"
)

type HitRecord struct {
	Distance  float64
	FrontFace bool
	Point     mathplus.Vector3
	Normal    mathplus.Vector3
	Material  *material.Material
}

func (r HitRecord) SetFaceNormal(ray *ray.Ray, outWardNormal mathplus.Vector3) {
	r.FrontFace = ray.Direction().Dot(outWardNormal) < 0

	if r.FrontFace {
		r.Normal = outWardNormal
	} else {
		r.Normal = outWardNormal.Mul(-1.0)
	}
}
