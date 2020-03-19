package serverFunctionality

// ShowStatistics implements a subsystem "ShowStatistics"
type ShowStatistics struct {
}

// Show implementation.
func (s *ShowStatistics) Show() string {
	return "Show statistics"
}
