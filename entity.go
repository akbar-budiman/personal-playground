package main

import (
	"math"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type GeometryObject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Rect struct {
	GeometryObject `json:"geometryObject"`
	Width          int `json:"width"`
	Length         int `json:"length"`
}

func (rect Rect) Area() float64 {
	return float64(rect.Width * rect.Length)
}

func (rect Rect) Perimeter() float64 {
	return float64(2 * (rect.Width + rect.Length))
}

type Circle struct {
	GeometryObject `json:"geometryObject"`
	Radius         float64 `json:"radius"`
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type GeometryProps struct {
	Area      float64
	Perimeter float64
}
