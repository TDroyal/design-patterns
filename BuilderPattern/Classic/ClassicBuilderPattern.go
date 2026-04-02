package classic

import (
	"errors"
)

type User struct {
	name   string // 必填字段
	sex    string
	age    int
	height float32
	weight float32
}

type UserBuilder struct {
	User
}

func NewUserBUilder() *UserBuilder {
	return &UserBuilder{}
}

// 通过链式调用方式完成User成员属性的赋值
func (u *UserBuilder) WithName(name string) *UserBuilder {
	u.name = name
	return u
}

func (u *UserBuilder) WithSex(sex string) *UserBuilder {
	u.sex = sex
	return u
}

func (u *UserBuilder) WithAge(age int) *UserBuilder {
	u.age = age
	return u
}

func (u *UserBuilder) WithHeight(height float32) *UserBuilder {
	u.height = height
	return u
}

func (u *UserBuilder) WithWeight(weight float32) *UserBuilder {
	u.weight = weight
	return u
}

func (u *UserBuilder) Build() (*User, error) {
	// 校验必填字段
	if u.name == "" {
		return nil, errors.New("miss name info")
	}

	return &u.User, nil
}
