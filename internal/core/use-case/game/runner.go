package game

import "github.com/andressg79/open-tama/internal/core/domain"

type Runner struct {
	Citizen *domain.Citizen
}

func NewRunner() Runner {
	return Runner{}
}

func (r Runner) Run() {
	r.Citizen.Update()
}

func (r Runner) PrintStats() bool {
	return r.Citizen.IsDead()
}
