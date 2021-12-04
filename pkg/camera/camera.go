package camera

import (
	. "github.com/404Polaris/RayTracing-go/pkg/mathplus"
	"math"
)

type Camera struct {
	apertureRadius  float64
	u, v, w         Vector3
	origin          Vector3
	vertical        Vector3
	horizontal      Vector3
	lowerLeftConner Vector3
}

func NewCamera(fov float64, vUp Vector3, lookAt Vector3, lookFrom Vector3, aspectRatio float64, apertureRadius float64, distToFocus float64) *Camera {
	camera := &Camera{apertureRadius: apertureRadius}
	theta := fov * math.Pi / 180
	h := math.Tan(theta / 2)
	vh := 2.0 * h
	vw := aspectRatio * vh

	camera.w = lookFrom.Sub(lookAt).Normalize()
	camera.u = vUp.Cross(camera.w).Normalize()
	camera.v = camera.w.Cross(camera.u)

	camera.origin = lookFrom
	camera.vertical = camera.v.Mul(distToFocus * vh)
	camera.horizontal = camera.u.Mul(distToFocus * vw)
	camera.lowerLeftConner = lookFrom.Sub(camera.horizontal.Div(2)).Sub(camera.vertical.Div(2))
	camera.lowerLeftConner = camera.lowerLeftConner.Sub(camera.w.Mul(distToFocus))
	return camera
}

func (c *Camera) GetRay(s float64, t float64) *Ray {
	rd := RandomVector3InUnitDisk().Mul(c.apertureRadius)
	offset := c.u.Mul(rd.X()).Add(c.v.Mul(rd.Y()))
	direction := c.lowerLeftConner.Add(c.horizontal.Mul(s)).Add(c.vertical.Mul(t)).Sub(c.origin).Sub(offset)
	return NewRay(c.origin.Add(offset), direction)
}
