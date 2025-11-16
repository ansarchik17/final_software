package observer

import "fmt"

type EmailNotifier struct {
	Email string
}

var _ IObserver = (*EmailNotifier)(nil)

func (emailNotifier *EmailNotifier) Update(event string) {
	fmt.Printf("[EMAIL to %s] %s\n", emailNotifier.Email, event)
}
