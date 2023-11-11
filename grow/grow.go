package main

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func plotGraph(x, y []float64, text string) error {
	p := plot.New()

	f, err := os.Create(fmt.Sprintf("%s.png", text))
	if err != nil {
		return fmt.Errorf("error creating png: %v", err)
	}
	defer f.Close()

	pxys := make(plotter.XYs, len(x))
	for i := range x {
		pxys[i].X = x[i]
		pxys[i].Y = y[i]
	}

	s, err := plotter.NewScatter(pxys)
	if err != nil {
		return fmt.Errorf("error creating scatter: %v", err)
	}
	s.Color = color.RGBA{R: 255, A: 255}

	p.Add(s)

	l, err := plotter.NewLine(pxys)
	if err != nil {
		return fmt.Errorf("error creating lines: %v", err)
	}
	l.Color = color.RGBA{G: 255, A: 255}

	p.Add(l)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("error init plot writer: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("error writting plot: %v", err)
	}

	return nil
}

func calculateGrow(iterations int, x0, r float64) []float64 {
	population := make([]float64, iterations+1)
	population[0] = x0

	for i := 0; i < iterations; i++ {
		population[i+1] = (1+r)*population[i] - r*population[i]*population[i]
		if population[i+1] < 0 {
			population[i+1] = 0
		}
	}
	return population
}

func calculatePower(iterations int, x0, r float64) []float64 {
	power := make([]float64, iterations+1)
	power[0] = x0

	for i := 0; i < iterations; i++ {
		power[i+1] = r * power[i] * (1 - power[i])
		if power[i+1] < 0 {
			power[i+1] = 0
		}
	}
	return power
}

func iterator(n int) []float64 {
	r := make([]float64, n)
	for i := range r {
		r[i] = float64(i)
	}
	return r
}

func main() {
	n := 70
	var r float64 = 3.5
	var x0 float64 = 0.2

	//p := calculateGrow(n, x0, r)
	p := calculatePower(n, x0, r)

	for i := 0; i < n+1; i++ {
		fmt.Printf("Нагрузка %d:\t%.2f\n", i, p[i])
	}

	plotGraph(iterator(n+1), p[:n+1], "Динамика роста")

	plotGraph(p[:n], p[1:n+1], "Фазовый портрет")
}
