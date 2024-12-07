package main

type musket struct {
	Gun
}

func newMusket() WeaponCrate {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
