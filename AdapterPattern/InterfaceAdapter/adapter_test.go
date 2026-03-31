package interfaceadapter

import "testing"

func TestAdapter(t *testing.T) {
	var golangProxy GolangCourseProxy = NewCourseService()
	golangProxy.LearnGolang()

	var javaProxy JavaCourseProxy = NewCourseService()
	javaProxy.LearnJava()
}

/*
=== RUN   TestAdapter
learn golang...
learn java...
--- PASS: TestAdapter (0.00s)
PASS
*/
