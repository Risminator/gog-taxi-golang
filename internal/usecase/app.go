package usecase

type Hello interface {
	SayHello(name string) (string, error)
}

type hello struct {
}

func (a *hello) SayHello(name string) (string, error) {
	message := "Hello, " + name
	return message, nil
}

func NewHelloUsecase() Hello {
	return &hello{}
}
