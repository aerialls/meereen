package check

// Check represents a way to verify a component
type Check interface {
	Run()
	GetTitle() string
}
