package domain

import "math/rand"

type Citizen struct {
	Live    bool
	Hunger  int
	Thirst  int
	Health  int
	Fun     int
	Sadness int
}

func New() *Citizen {
	return &Citizen{
		Live:    true,
		Hunger:  0,
		Thirst:  0,
		Health:  100,
		Fun:     100,
		Sadness: 0,
	}
}

func (c *Citizen) IsDead() bool {
	return !c.Live
}

func (c *Citizen) IsAlive() bool {
	return c.Live
}

func (c *Citizen) Eat(value int) {
	c.Hunger = incrementor(c.Hunger, value, 100)
}

func (c *Citizen) Drink(value int) {
	c.Thirst = incrementor(c.Thirst, value, 100)
}

func (c *Citizen) Play(value int) {
	c.Fun = incrementor(c.Fun, value, 100)
}

func (c *Citizen) Relax(value int) {
	c.Sadness = incrementor(c.Sadness, value, 100)
}

func (c *Citizen) Sleep(value int) {
	c.Thirst = decresor(c.Thirst, value, 0)
	c.Hunger = decresor(c.Hunger, value, 0)
	c.Sadness = decresor(c.Sadness, value, 0)
}

func incrementor(oldvalue, value, max int) int {
	newValue := oldvalue + value
	if newValue > max {
		newValue = max
	}
	return newValue
}

func decresor(oldvalue, value, min int) int {
	newValue := oldvalue - value
	if newValue < min {
		newValue = min
	}
	return newValue
}

func (c *Citizen) Update() {
	c.Health = randomDecrector(c.Health, 0)
	c.Fun = randomDecrector(c.Fun, 0)
	c.Sadness = randomDecrector(c.Sadness, 0)
	c.Thirst = randomIncrementor(c.Thirst, 100)
	c.Hunger = randomIncrementor(c.Hunger, 100)
	c.Sadness = randomIncrementor(c.Sadness, 100)

	if c.Health <= 0 && c.Hunger >= 100 {
		c.Live = false
	}
}

func randomDecrector(value, min int) int {
	newValue := value - rand.Intn(2)
	if newValue < min {
		newValue = min
	}
	return newValue
}

func randomIncrementor(value, max int) int {
	newValue := value - rand.Intn(2)
	if newValue < max {
		newValue = max
	}
	return newValue
}
