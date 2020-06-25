package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	m1 := map[int]int{}
	// m1[3] = 0
	if v, ok := m1[3]; ok == false {
		t.Log("Key 3 is not existing.")
	} else {
		t.Logf("Key 3's value is %d", v)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
