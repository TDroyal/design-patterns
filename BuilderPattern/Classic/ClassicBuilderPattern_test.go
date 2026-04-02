package classic

import (
	"testing"
)

func TestClassicBuilder(t *testing.T) {
	ub := NewUserBUilder()

	user, err := ub.WithHeight(173).Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("user: %+v", user)
	}

	user, err = ub.WithName("royal_111").WithSex("male").WithAge(24).WithWeight(60).Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("user: %+v", user)
	}

	user, err = ub.WithHeight(175).Build()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("user: %+v", user)
	}
}

/*
=== RUN   TestClassicBuilder
    e:\project\DesignPatterns\BuilderPattern\Classic\ClassicBuilderPattern_test.go:12: miss name info
    e:\project\DesignPatterns\BuilderPattern\Classic\ClassicBuilderPattern_test.go:21: user: &{name:royal_111 sex:male age:24 height:173 weight:60}
    e:\project\DesignPatterns\BuilderPattern\Classic\ClassicBuilderPattern_test.go:28: user: &{name:royal_111 sex:male age:24 height:175 weight:60}
--- FAIL: TestClassicBuilder (0.00s)
FAIL
*/
