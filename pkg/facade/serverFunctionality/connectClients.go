package serverFunctionality

// ConnectClients implements a subsystem "ConnectClients"
type ConnectClients struct {
}

// Connect implementation.
func (c *ConnectClients) Connect() string {
	return "Connect clients"
}
