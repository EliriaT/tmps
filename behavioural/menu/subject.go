package menu

import "fmt"

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type Item struct {
	observerList []Observer
	Menu         []string
}

func NewMenu(name []string) *Item {
	return &Item{
		Menu: name,
	}
}
func (i *Item) UpdateMenu(newDish string) {
	fmt.Printf("A new menu is available %s is now in stock with \n", newDish)
	i.Menu = append(i.Menu, newDish)
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) Deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.Menu)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
