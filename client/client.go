package client

// Client is Zai client.
// It shouldn't be used after calling Close().
type Client interface {

	// Close closes the client.
	Close()
}
