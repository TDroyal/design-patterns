package option1

import "testing"

func TestOptionPattern(t *testing.T) {
	u1 := NewUser()
	u2 := NewUser(WithName("royal_111"), WithFieldA("aaa"), WithAge(35))
	u3 := NewUser(WithSex("female"))

	t.Logf("user1: %+v\nuser2: %+v\nuser3: %+v", u1, u2, u3)
}

/*
=== RUN   TestOptionPattern
    e:\project\DesignPatterns\BuilderPattern\Option1\OptionPattern_test.go:10: user1: &{Options:{name:default_name sex: age:0 height:0 weight:0 fieldA: fieldB: fieldC:}}
        user2: &{Options:{name:royal_111 sex: age:35 height:0 weight:0 fieldA:aaa fieldB: fieldC:}}
        user3: &{Options:{name:default_name sex:female age:0 height:0 weight:0 fieldA: fieldB: fieldC:}}
--- PASS: TestOptionPattern (0.00s)
PASS
*/
