// mech project mech.go
package mech

import (
	"fmt"
)

// Mech is a basic mech type
type mech struct {
	structure int
	weapons   Weapon
	name      string
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(weapon Weapon, name string) *mech {
	newMech := new(mech)
	newMech.structure = 2
	newMech.weapons = weapon
	newMech.name = name

	return newMech
}

// internal call when a mech is hit
func (mech *mech) hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.name, "destroyed")
	}
}

// Add a Weapon to the mech
func (mech *mech) AddWeapon(weapon Weapon) {

}

// Tell the Mech to fire at a Target
func (mech *mech) Fire(target Target) {

}
