package vehicle

// For the Visitor pattern.
type IVisitable interface {
	Accept(v IVisitor)
}
