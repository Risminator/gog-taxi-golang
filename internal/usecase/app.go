package usecase

type Hello interface {
	SayHello(name string) (string, error)
}

type hello struct {
}

func NewHelloUsecase() Hello {
	return &hello{}
}

func (h *hello) SayHello(name string) (string, error) {
	message := "Hello, " + name
	return message, nil
}
