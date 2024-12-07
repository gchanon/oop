package main

type Armor struct {
	name string
	def  int
}

func (a *Armor) setName(name string) {
	a.name = name
}

func (a *Armor) getName() string {
	return a.name
}

func (a *Armor) setAbility(status int) {
	a.def = status
}

func (a *Armor) getAbility() int {
	return a.def
}
