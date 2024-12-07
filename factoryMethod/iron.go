package main

type Iron struct {
	Armor
}

func newIronArmor() WeaponCrate {
	return &Iron{
		Armor{
			name: "iron_armor",
			def:  300,
		},
	}
}