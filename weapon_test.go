package mech

import "testing"

func TestMechFire(t *testing.T) {
	weapon1 := Weapon{2, 2}
	weapon2 := Weapon{1, 4}

	mech1 := NewMech(weapon1, "testMech1")
	if mech1 == nil {
		t.Errorf("mech1 was unable to be created")
	}

	mech2 := NewMech(weapon2, "testMech2")
	if mech2 == nil {
		t.Errorf("mech2 was unable to be created")
	}

	mech1.Weapons.Fire(3, mech2)
	if mech2.structure != 2 {
		t.Errorf("mech1 destroyed at range 3 by range 2 weapon")
	}

	mech1.Weapons.Fire(2, mech2)
	if mech2.structure != 0 {
		t.Errorf("mech2 not destroyed by mech1 at range 2 by range 2, damage 2 weapon")
	}
}
