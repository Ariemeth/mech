package mech

import "testing"

var testWeapon1 Weapon = Weapon{Damage:2, Range:2}
var testWeapon2 Weapon = Weapon{Damage:1, Range:4}

func TestNewMech(t *testing.T) {
	mech1 := NewMech(testWeapon1, "testMech1")
	if mech1 == nil {
		t.Errorf("mech1 was unable to be created")
	}

	mech2 := NewMech(testWeapon2, "testMech2")
	if mech2 == nil {
		t.Errorf("mech2 was unable to be created")
	}	
}

func TestHit(t *testing.T) {
	mech1 := NewMech(testWeapon1, "testMech1")
	if mech1 == nil{
		t.Errorf("mech1 was unable to be created")
	}
	
	mech1.hit(0)
	if mech1.structure != 2 {
		t.Errorf("mech1 took damage when it was hit with 0")
	}
	
	mech1.hit(2)
	if mech1.structure != 0 {
		t.Errorf("mech1 was not destroyed by taking 2 damage")
	}	
}

func TestAddWeapon(t *testing.T) {
	
}


func TestMechFire(t *testing.T) {
	
}


