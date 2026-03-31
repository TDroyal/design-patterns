package interfaceadapter

import "fmt"

// 课程实现方
type CourseService struct {
}

func NewCourseService() *CourseService {
	return &CourseService{}
}

func (c *CourseService) LearnGolang() {
	fmt.Println("learn golang...")
}

func (c *CourseService) LearnJava() {
	fmt.Println("learn java...")
}

func (c *CourseService) LearnCPlusPlus() {
	fmt.Println("learn c++...")
}

// interface由使用方定义
/*
对于使用方来说，需要根据自身对于 CourseService 的适用范围，完成 interface 的定义.
这个定义过程即明确了自身对于 CourseService 职责定位，也实现了对无关方法的消音屏蔽.

在这个过程中，interface 就扮演了适配器的角色，对 class 的范围起到适配和收敛的作为，
使其在使用方手中，有一个更加恰到好处的定位和空间.
*/
type GolangCourseProxy interface {
	LearnGolang()
}

type JavaCourseProxy interface {
	LearnJava()
}
