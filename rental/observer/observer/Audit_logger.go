package observer

import "fmt"

type AuditLogger struct{}

var _ IObserver = (*AuditLogger)(nil)

func (logger *AuditLogger) Update(event string) {
	fmt.Println("[AUDIT]", event)
}
