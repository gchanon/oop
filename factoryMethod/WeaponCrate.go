package main

type WeaponCrate interface {
	setName(name string)
	setAbility(status int)
	getName() string
	getAbility() int
}
