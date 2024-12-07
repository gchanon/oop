package main

import "fmt"

func getGun(gunType string) (WeaponCrate, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func getArmor(armorType string) (WeaponCrate, error) {
	if armorType == "wood" {
		return newWoodArmor(), nil
	}
	if armorType == "iron" {
		return newIronArmor(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
