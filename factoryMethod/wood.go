package main

type Wood struct {
	Armor
}

func newWoodArmor() WeaponCrate {
	return &Wood{
		Armor{
			name: "wood_armor",
			def:  100,
		},
	}
}
