package model

import (
	"testing"
)

func TestValidateDNI(t *testing.T) {
	 dni := NewDNI("18047606", "7")
	if !dni.Validate() {
		t.Error("Invalid token")
	}

	 dni2 := NewDNI("12678579", "8")
	 if !dni2.Validate() {
	 	t.Error("Invalid token")
	 }
}
