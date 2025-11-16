package observer

import "fmt"

type AuditLogger struct{}

var _ IObserver = (*AuditLogger)(nil)

func (l *AuditLogger) Update(event string) {
	fmt.Println("[AUDIT]", event)
}
