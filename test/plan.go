package test

// Plan represents an execution plan for tests to be run
type Plan struct {
	Tests []*Test
}

// Add takes a test and adds it to the current plan
func (p *Plan) Add(test *Test) {
	p.Tests = append(p.Tests, test)
}
