package observer

type IObserver interface {
	Update(event string)
}
