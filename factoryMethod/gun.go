package main

type Gun struct {
	name   string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setAbility(status int) {
	g.power = status
}

func (g *Gun) getAbility() int {
	return g.power
}
