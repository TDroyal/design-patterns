package option2

import "testing"

func TestOptionPattern(t *testing.T) {
	u := NewUser(WithName("royal_111"), WithAge(24))
	t.Logf("user: %+v", u)
}

/*
=== RUN   TestOptionPattern
    e:\project\DesignPatterns\BuilderPattern\Option2\OptionPattern_test.go:7: user: &{opts:{name:royal_111 sex: age:24 height:0 weight:0 fieldA: fieldB: fieldC:}}
--- PASS: TestOptionPattern (0.00s)
PASS
*/
