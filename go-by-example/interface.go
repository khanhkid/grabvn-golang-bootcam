package main

import (
	"fmt"
)

type tinhtoan interface {
	dientich() float64
	chuvi() float64
}

type hinhvuong struct {
	dai, rong float64
}

type chunhat struct {
	dai, rong float64
}

func (h hinhvuong) dientich() float64 {
	return (h.dai * h.rong)
}

func (h hinhvuong) chuvi() float64 {
	return (h.dai + h.rong) * 2
}

func (c chunhat) dientich() float64 {
	return (c.dai * c.rong)
}

func main1() {
	hinhv := hinhvuong{dai: 3, rong: 2}
	fmt.Println(hinhv.dientich())

	fmt.Println("test")
}
