package main

type AreaCounter struct {
	obj Geometry
}

func (ac AreaCounter) Count() float64 {
	return ac.obj.Area()
}

type PerimeterCounter struct {
	obj Geometry
}

func (pc PerimeterCounter) Count() float64 {
	return pc.obj.Perimeter()
}
