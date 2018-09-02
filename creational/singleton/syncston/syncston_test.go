package syncston

import "testing"

func TestNew(t *testing.T) {
	singleton := New()

	cases := []struct {
		addKey       string
		addValue     []byte
		wantMap      map[string][]byte
		wantLenOfMap int
	}{
		{"First", []byte{1, 2, 3},
			map[string][]byte{"First": []byte{1, 2, 3}}, 1},
		{"Second", []byte{1, 2},
			map[string][]byte{"First": []byte{1, 2, 3}, "Second": []byte{1, 2}}, 2},
	}
	for _, v := range cases {
		//create new type, should be the same
		singleton = New()
		singleton[v.addKey] = v.addValue

		if !equalOfMaps(v.wantMap, singleton) {
			t.Errorf("Two maps should be the same. Got [%v], but want[%v]", singleton, v.wantMap)
		}

	}
}

//test helper function for checking two maps are equal
func equalOfMaps(first, second map[string][]byte) bool {
	if len(first) != len(second) {
		return false
	}

	for key, value := range first {
		secValue, ok := second[key]
		if !ok {
			return false
		}
		//check slice of bytes ([]bytes)
		for bytKey, bytVal := range value {
			if bytVal != secValue[bytKey] {
				return false
			}
		}
	}
	return true
}
