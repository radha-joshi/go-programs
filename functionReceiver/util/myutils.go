package util

// definition of behaviour
type StringUtils interface {
	Reverse(text string) string
}

type myUtils struct{}

func (u *myUtils) Reverse(text string) string {
	return "Reverse from MyUtils"
}

func NewMyUtils() StringUtils {
	return &myUtils{}
}

type mySuperUtils struct{}

func (u *mySuperUtils) Reverse(text string) string {
	return "Reverse from MySuperUtils"
}

func NewMySuperUtils() StringUtils {
	return &mySuperUtils{}
}
