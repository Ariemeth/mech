package mech

import "testing"

func TestMech(t *testing.T) {
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
}

func TestWeapon(t *testing.T) {
	weapon1 := Weapon{2, 2}
	if weapon1.Damage != 2 {
		t.Errorf("weapon1 damage is %i instead of 2", weapon1.Damage)
	}
	if weapon1.Range != 2 {
		t.Errorf("weapon1 range is %i instead of 2", weapon1.Range)
	}
}
