package mech

import "testing"

func TestWeapon(t *testing.T) {
	weapon1 := Weapon{Damage:2, Range:2}
	if weapon1.Damage != 2 {
		t.Errorf("weapon1 damage is %i instead of 2", weapon1.Damage)
	}
	if weapon1.Range != 2 {
		t.Errorf("weapon1 range is %i instead of 2", weapon1.Range)
	}
}

func TestWeaponFire(t *testing.T) {
	weapon1 := Weapon{Damage:2, Range:2}
	weapon2 := Weapon{Damage:1, Range:4}

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
