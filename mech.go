// mech project mech.go
package mech

import (
	"fmt"
)

// Mech is a basic mech type
type Mech struct {
	structure int
	Weapons   Weapon
	Name      string
}

// NewMech is used to create a new instance of a mech with default structure.
func NewMech(weapon Weapon, name string) *Mech {
	return &Mech{2, weapon, name}
}

// internal call when a mech is hit
func (mech *Mech) hit(damage int) {
	mech.structure -= damage
	fmt.Println(mech.Name, "takes", damage, "damage")
	if mech.structure <= 0 {
		fmt.Println(mech.Name, "destroyed")
	}
}

// A specific weapon.
type Weapon struct {
	Range, Damage int
}

// Used by a mech to fire at another mech.  Requires the range to the opposing mech
// and the mech that is being targeted.
func (weapon *Weapon) Fire(rangeToTarget int, target *Mech) {
	if rangeToTarget <= weapon.Range {
		target.hit(weapon.Damage)
	}
}
