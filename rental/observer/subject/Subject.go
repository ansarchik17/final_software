package subject

import "final_software/rental/observer/observer"

type Subject struct {
	obs []observer.IObserver
}

var _ ISubject = (*Subject)(nil)

func NewSubject() *Subject {
	return &Subject{}
}

func (subject *Subject) Register(o observer.IObserver) {
	subject.obs = append(subject.obs, o)
}

func (subject *Subject) Notify(event string) {
	for _, o := range subject.obs {
		o.Update(event)
	}
}
