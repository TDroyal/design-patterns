package option2

type User struct {
	opts userOptions
	// ...
}

type userOptions struct {
	name   string
	sex    string
	age    int
	height float32
	weight float32
	fieldA string
	fieldB string
	fieldC string
}

type UserOption interface {
	apply(*userOptions)
}

type funcUserOption struct {
	f func(*userOptions)
}

func newFuncUserOption(f func(*userOptions)) *funcUserOption {
	return &funcUserOption{
		f: f,
	}
}

func (fuo *funcUserOption) apply(opts *userOptions) {
	fuo.f(opts)
}

func WithName(name string) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.name = name
	})
}

func WithAge(age int) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.age = age
	})
}

func WithSex(sex string) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.sex = sex
	})
}

func WithHeight(height float32) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.height = height
	})
}

func WithWeight(weight float32) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.weight = weight
	})
}

func WithFieldA(fieldA string) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.fieldA = fieldA
	})
}

func WithFieldB(fieldB string) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.fieldB = fieldB
	})
}

func WithFieldC(fieldC string) UserOption {
	return newFuncUserOption(func(uo *userOptions) {
		uo.fieldC = fieldC
	})
}

var defaultUserOptions userOptions = userOptions{}

func NewUser(opts ...UserOption) *User {
	uo := defaultUserOptions
	for _, opt := range opts {
		opt.apply(&uo)
	}

	u := &User{
		opts: uo,
		// ...
	}
	return u
}
