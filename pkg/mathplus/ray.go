package mathplus

type Ray struct {
	origin    Vector3
	direction Vector3
}

func NewRay(origin Vector3, direction Vector3) *Ray {
	return &Ray{origin: origin, direction: direction}
}

func (r *Ray) Origin() Vector3 {
	return r.origin
}

func (r *Ray) Direction() Vector3 {
	return r.direction
}

func (r *Ray) At(length float64) Vector3 {
	return r.origin.Add(r.direction.Mul(length))
}
