package domain

type Citizen struct {
	Live bool
	Hunger int
	Thirst int 
	Health int 
	Fun int 
	Sadness int
}

func New() Citizen {
	return Citizen{
		Live:    true,
		Hunger:  0,
		Thirst:  0,
		Health:  100,
		Fun:     100,
		Sadness: 0,
	}	
}