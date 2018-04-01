package main

import "github.com/k0kubun/pp"

func main() {
	animalShelter := newAnimalShelter()
	animalShelter.enqueue(dog{name: "taro"})
	animalShelter.enqueue(dog{name: "poti"})
	animalShelter.enqueue(cat{name: "tama"})
	animalShelter.enqueue(dog{name: "siro"})
	animalShelter.enqueue(cat{name: "pika"})
	animalShelter.enqueue(dog{name: "hoge"})
	animalShelter.enqueue(cat{name: "chou"})
	pp.Println(animalShelter)
	pp.Println(animalShelter.dequeueAny())
	pp.Println(animalShelter.dequeueDog())
	pp.Println(animalShelter.dequeueCat())
	pp.Println(animalShelter.dequeueCat())
}

func newAnimalShelter() *animalShelter {
	dogQueue := &dogQueue{queue: make([]*dog, 0, 0)}
	catQueue := &catQueue{queue: make([]*cat, 0, 0)}
	return &animalShelter{dogs: dogQueue, cats: catQueue, order: 0}
}

type animalShelter struct {
	dogs  *dogQueue
	cats  *catQueue
	order int
}

type dogQueue struct {
	queue []*dog
}

type catQueue struct {
	queue []*cat
}

type animal interface {
	setOrder(v int)
	getOrder() int
	isOlderThan(an animal) bool
}

type dog struct {
	order int
	name  string
}

type cat struct {
	order int
	name  string
}

func (d *dog) setOrder(order int) {
	d.order = order
}

func (c *cat) setOrder(order int) {
	c.order = order
}

func (d *dog) getOrder() int {
	return d.order
}

func (c *cat) getOrder() int {
	return c.order
}

func (d *dog) isOlderThan(a cat) bool {
	return d.order < a.getOrder()
}

func (c *cat) isOlderThan(a dog) bool {
	return c.order < a.getOrder()
}

func (a *animalShelter) enqueue(an interface{}) {
	switch v := an.(type) {
	case dog:
		a.order++
		v.setOrder(a.order)
		a.dogs.queue = append(a.dogs.queue, &v)
	case cat:
		a.order++
		v.setOrder(a.order)
		a.cats.queue = append(a.cats.queue, &v)
	}
}

func (a *animalShelter) dequeueAny() interface{} {
	cat := a.cats.queue[0]
	dog := a.dogs.queue[0]

	if dog.isOlderThan(*cat) {
		if len(a.dogs.queue) == 0 {
			return nil
		}
		a.dogs.queue = a.dogs.queue[1:]
		return dog
	}
	if len(a.cats.queue) == 0 {
		return nil
	}
	a.cats.queue = a.cats.queue[1:]
	return cat
}

func (a *animalShelter) dequeueDog() *dog {
	if len(a.dogs.queue) == 0 {
		return nil
	}
	dog := a.dogs.queue[0]
	a.dogs.queue = a.dogs.queue[1:]
	return dog
}

func (a *animalShelter) dequeueCat() *cat {
	if len(a.cats.queue) == 0 {
		return nil
	}
	cat := a.cats.queue[0]
	a.cats.queue = a.cats.queue[1:]
	return cat
}
