package subject

import "final_software/rental/observer/observer"

type ISubject interface {
	Register(observer observer.IObserver)
	Notify(event string)
}
