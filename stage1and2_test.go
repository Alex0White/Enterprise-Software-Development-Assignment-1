package stage1and2

import (
	"testing"
	"fmt"
	)


func TestUnqueStringItemizer(t *testing.T){
	v := []string{"hello", "whatup", "howdy", "g day"} 
	var s string
	
	s = "g day"
	expected := false
	actual := UnqueStringItemizer(v, s)
	if actual != expected{
		t.Errorf("Test failed", expected, actual)
	}

}