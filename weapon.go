
package mech

// A specific weapon.
type Weapon struct {
	Range, Damage int
}



// Used by a mech to fire at another mech.  Requires the range to the opposing mech
// and the mech that is being targeted.
func (weapon Weapon) Fire(rangeToTarget int, target Target) {
	if rangeToTarget <= weapon.Range {
		target.hit(weapon.Damage)
	}
}