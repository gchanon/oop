package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	woodArmor, _ := getArmor("wood")
	ironArmor, _ := getArmor("iron")


	printDetails(ak47)
	printDetails(musket)
	printDetails(woodArmor)
	printDetails(ironArmor)
}

func printDetails(g WeaponCrate) {
	fmt.Printf("Weapon: %s", g.getName())
	fmt.Println()
	fmt.Printf("Ability Status: %d", g.getAbility())
	fmt.Println()
}
