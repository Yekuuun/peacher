package builder

type App struct {
	TelegramToken string
	Port          int
}

type AppBuilder struct {
	telegramToken string
	port          int
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

func (b *AppBuilder) WithTelegramToken(token string) *AppBuilder {
	b.telegramToken = token
	return b
}

func (b *AppBuilder) WithPort(port int) *AppBuilder {
	b.port = port
	return b
}

func (b *AppBuilder) Build() *App {
	if b.port == 0 {
		b.port = 8080
	}

	return &App{
		TelegramToken: b.telegramToken,
		Port:          b.port,
	}
}
