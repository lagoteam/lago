package providers

type AppServiceProvider struct {
}

func (r AppServiceProvider) Boot() {
	RouteServiceProvider{}.Boot()
	ConsoleServiceProvider{}.Boot()
	EventServiceProvider{}.Boot()
	QueueServiceProvider{}.Boot()
}

func (r AppServiceProvider) Register() {
	RouteServiceProvider{}.Register()
	ConsoleServiceProvider{}.Register()
	EventServiceProvider{}.Register()
	QueueServiceProvider{}.Register()
}
