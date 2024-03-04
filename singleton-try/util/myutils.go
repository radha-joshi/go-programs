package util

// definition of behaviour
type StringUtils interface {
	Reverse(text string) string
	SetName(name string)
	GetName() string
}

type myUtils struct {
	name string
}

func (u *myUtils) Reverse(text string) string {
	return u.name
}

func (u *myUtils) SetName(text string) {
	u.name = text
}

func (u *myUtils) GetName() string {
	return u.name
}

var myUtilsInstance = &myUtils{
	name: "default",
}

func NewMyUtils() StringUtils {
	return myUtilsInstance
}
