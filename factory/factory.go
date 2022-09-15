package factory

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type ak47 struct {
	Gun
}

func newAk47() IGun {
	return &ak47{Gun{
		name:  "AK47 gun",
		power: 4,
	}}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{Gun{
		name:  "Musket gun",
		power: 1,
	}}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}
