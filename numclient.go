package numary

//Connects to numary, handles auth and basic request stuff
type Connection struct {
	Server string
}

func Connect(server string) Ledger {
	return Connection{Server: server}
}

type Ledger struct {
}
