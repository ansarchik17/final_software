package subject

import "final_software/rental/observer/observer"

type Subject struct {
	obs []observer.IObserver
}

var _ ISubject = (*Subject)(nil)

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Register(o observer.IObserver) {
	s.obs = append(s.obs, o)
}

func (s *Subject) Notify(event string) {
	for _, o := range s.obs {
		o.Update(event)
	}
}
