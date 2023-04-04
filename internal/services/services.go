package services

type Keenetic interface {
	Check() error
}

type Services struct {
	Keenetic Keenetic
}

type Deps struct {
	KeeneticDnsDomain string
}

func NewServices(deps *Deps) *Services {
	return &Services{
		Keenetic: NewKeeneticService(deps.KeeneticDnsDomain),
	}
}