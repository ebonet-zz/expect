package expect

import (
	"testing"
)

func TestExpectCondition(t *testing.T) {

	testMoc := &testing.T{}

	Condition(testMoc, 1 == 1, "")

	if testMoc.Failed() {
		t.Errorf("Condition: failing on 1==1")
	}

	Condition(testMoc, 1 == 0, "")

	if !testMoc.Failed() {
		t.Errorf("Condition: not failing on 1==0")
	}

}

func TestEqual(t *testing.T) {

	testMoc := &testing.T{}

	Equal(testMoc, 1, 1, "")

	if testMoc.Failed() {
		t.Errorf("Equal: failing on Equal(1,1)")
	}

	testMoc = &testing.T{}

	Equal(testMoc, 1, 0, "")

	if !testMoc.Failed() {
		t.Errorf("Condition: not failing on Equal(1,0)")
	}

	testMoc = &testing.T{}

	Equal(testMoc, 1, "1", "")

	if !testMoc.Failed() {
		t.Errorf("Condition: not failing on Equal(1,\"1\")")
	}

	testMoc = &testing.T{}

	Equal(testMoc, "1", "1", "")

	if testMoc.Failed() {
		t.Errorf("Equal: failing on Equal(\"1\",\"1\")")
	}

}
