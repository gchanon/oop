package main

type Ak47 struct {
	Gun
}

func newAk47() WeaponCrate {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
