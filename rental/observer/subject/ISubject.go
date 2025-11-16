package subject

import "final_software/rental/observer/observer"

type ISubject interface {
	Register(o observer.IObserver)
	Notify(event string)
}
