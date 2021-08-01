package numary

//Connects to numary, handles auth and basic request stuff
type Connection struct {
	Server string
}

func Connect(server string) Connection {
	return Connection{Server: server}
}

type Ledger struct {
}
