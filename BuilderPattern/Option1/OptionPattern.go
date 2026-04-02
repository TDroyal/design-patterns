package option1

// big class
type User struct {
	Options
}

type Options struct {
	name   string // 必填字段
	sex    string
	age    int
	height float32
	weight float32
	fieldA string
	fieldB string
	fieldC string
}

type Option func(opts *Options)

func WithName(name string) Option {
	return func(opts *Options) {
		opts.name = name
	}
}

func WithSex(sex string) Option {
	return func(opts *Options) {
		opts.sex = sex
	}
}

func WithAge(age int) Option {
	return func(opts *Options) {
		opts.age = age
	}
}

func WithHeight(height float32) Option {
	return func(opts *Options) {
		opts.height = height
	}
}

func WithWeight(weight float32) Option {
	return func(opts *Options) {
		opts.weight = weight
	}
}

func WithFieldA(fieldA string) Option {
	return func(opts *Options) {
		opts.fieldA = fieldA
	}
}

func WithFieldB(fieldB string) Option {
	return func(opts *Options) {
		opts.fieldB = fieldB
	}
}

func WithFieldC(fieldC string) Option {
	return func(opts *Options) {
		opts.fieldC = fieldC
	}
}

// 兜底修复方法 repair 完成构造 User 实例过程中一些缺省值的设置.
func repair(o *Options) {
	if o.name == "" {
		o.name = "default_name"
	}
}

func NewUser(opts ...Option) *User {
	u := new(User)

	for _, opt := range opts {
		opt(&u.Options)
	}

	repair(&u.Options)
	return u
}
