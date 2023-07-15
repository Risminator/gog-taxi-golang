package app

type App interface {
	SayHello(name string) (string, error)
}

type app struct {
}

func (a *app) SayHello(name string) (string, error) {
	message := "Hello, " + name
	return message, nil
}

func NewApp() App {
	return &app{}
}
