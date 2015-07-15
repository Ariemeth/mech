// mech project mech.go
package mech

import (
	"fmt"
)

// Mech is a basic mech type
type mech struct {
	structure int
	Weapons   Weapon
	Name      string
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(weapon Weapon, name string) *mech {
	newMech := new(mech)
	newMech.structure = 2
	newMech.Weapons = weapon
	newMech.Name = name
	
	return newMech
}

// internal call when a mech is hit
func (mech *mech) hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.Name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.Name, "destroyed")
	}
}

// Add a Weapon to the mech
func (mech *mech) AddWeapon(weapon Weapon) {

}

// Tell the Mech to fire at a Target
func (mech *mech) Fire(target Target) {
	
}
