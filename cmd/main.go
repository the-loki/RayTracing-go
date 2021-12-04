package main

import (
	"fmt"
	"github.com/404Polaris/RayTracing-go/pkg/camera"
	"github.com/404Polaris/RayTracing-go/pkg/geometry"
	"github.com/404Polaris/RayTracing-go/pkg/material"
	"github.com/404Polaris/RayTracing-go/pkg/mathplus"
	"github.com/404Polaris/RayTracing-go/pkg/scene"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sync"
)

func rayColor(ray *mathplus.Ray, scene geometry.Hittable, depth int) mathplus.Vector3 {
	if depth <= 0 {
		return *mathplus.NewVector3(0, 0, 0)
	}

	var hitInfo geometry.HitInfo
	if scene.Hit(ray, 0.001, math.MaxFloat64, &hitInfo) {
		var scatteredRay mathplus.Ray
		var attenuation mathplus.Vector3

		mat := hitInfo.Material.(material.Material)
		if mat.Scatter(ray, hitInfo, &attenuation, &scatteredRay) {
			return attenuation.MulVector(rayColor(&scatteredRay, scene, depth-1))
		}

		return *mathplus.NewVector3(0, 0, 0)
	}

	direction := ray.Direction().Normalize()
	t := 0.5 * (direction.Y() + 1.0)
	return mathplus.NewVector3(1, 1, 1).Mul(1.0 - t).Add(mathplus.NewVector3(0.5, 0.7, 1.0).Mul(t))
}

func randomScene() geometry.Hittable {
	s := scene.NewScene()
	m1 := material.NewDielectric(1.5)
	s.Add(geometry.NewSphere(1, *mathplus.NewVector3(0, 1, 0), m1))
	m2 := material.NewLambertian(*mathplus.NewVector3(0.4, 0.2, 0.1))
	s.Add(geometry.NewSphere(1, *mathplus.NewVector3(-4, 1, 0), m2))
	m3 := material.NewMetal(0.0, *mathplus.NewVector3(0.7, 0.6, 0.5))
	s.Add(geometry.NewSphere(1, *mathplus.NewVector3(4, 1, 0), m3))
	m4 := material.NewLambertian(*mathplus.NewVector3(0.5, 0.5, 0.5))
	s.Add(geometry.NewSphere(1000, *mathplus.NewVector3(0, -1000, 0), m4))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			rFloat := mathplus.RandomFloat64()
			center := mathplus.NewVector3(float64(a)+0.9*rFloat, 0.2, float64(b)+0.9*rFloat)

			if center.Sub(*mathplus.NewVector3(4, 0.2, 0)).Length() <= 0.9 {
				continue
			}

			if rFloat < 0.8 {
				albedo := mathplus.RandomVector3().MulVector(mathplus.RandomVector3())
				m := material.NewLambertian(albedo)
				s.Add(geometry.NewSphere(0.2, *center, m))
			} else if rFloat < 0.95 {
				albedo := mathplus.RandomVector3()
				fuzz := mathplus.RandomFloat64InRange(0, 0.5)
				m := material.NewMetal(fuzz, albedo)
				s.Add(geometry.NewSphere(0.2, *center, m))
			} else {
				m := material.NewDielectric(1.5)
				s.Add(geometry.NewSphere(0.2, *center, m))
			}
		}
	}

	return s
}

func calcRayTracing(x int, y int, cam *camera.Camera, w int, h int, scene geometry.Hittable, reflectTimes int, samplePerPixel int) color.RGBA {
	wg := sync.WaitGroup{}
	wg.Add(samplePerPixel)
	pixelColor := *mathplus.NewVector3(0, 0, 0)
	pixelColors := make([]mathplus.Vector3, samplePerPixel, samplePerPixel)

	calcRayColor := func(i int, u float64, v float64) {
		ray := cam.GetRay(u, v)
		pixelColors[i] = rayColor(ray, scene, reflectTimes)
		wg.Done()
	}

	for i := 0; i < samplePerPixel; i++ {
		u := (float64(x) + mathplus.RandomFloat64()) / (float64(w) - 1)
		v := (float64(y) + mathplus.RandomFloat64()) / (float64(h) - 1)
		go calcRayColor(i, u, v)
	}

	wg.Wait()
	for i := 0; i < samplePerPixel; i++ {
		pixelColor = pixelColor.Add(pixelColors[i])
	}

	pixelColor = pixelColor.Div(float64(samplePerPixel))
	pixelColor = *mathplus.NewVector3(math.Sqrt(pixelColor.X()), math.Sqrt(pixelColor.Y()), math.Sqrt(pixelColor.Z()))

	R := mathplus.Clamp(pixelColor.R(), 0, 0.999)
	G := mathplus.Clamp(pixelColor.G(), 0, 0.999)
	B := mathplus.Clamp(pixelColor.B(), 0, 0.999)
	pixelColor = *mathplus.NewVector3(R, G, B)
	pixelColor = pixelColor.Mul(256)

	return color.RGBA{R: uint8(pixelColor.R()), G: uint8(pixelColor.G()), B: uint8(pixelColor.B()), A: 255}
}

func main() {
	w := 768
	reflectTimes := 500
	samplePerPixel := 500
	aspectRatio := 16.0 / 9.0
	h := int(float64(w) / aspectRatio)

	distToFocus := 10.0
	vUp := mathplus.NewVector3(0, 1, 0)
	lookAt := mathplus.NewVector3(0, 0, 0)
	lookFrom := mathplus.NewVector3(12, 2, 3)
	cam := camera.NewCamera(20, *vUp, *lookAt, *lookFrom, aspectRatio, 0, distToFocus)
	s := randomScene()

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pixelColor := calcRayTracing(x, y, cam, w, h, s, reflectTimes, samplePerPixel)
			img.SetRGBA(x, h-y, pixelColor)
		}

		fmt.Printf("Current Progress : %f\n", float64(y)*100.0/float64(h))
	}

	imgFile, err := os.Create("image.png")
	defer imgFile.Close()
	err = png.Encode(imgFile, img)

	if err != nil {
		log.Fatal(err)
	}
}
