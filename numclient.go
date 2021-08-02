package numary

import (
	"encoding/json"
	"io"
	"net/http"
)

type respFmt = map[string]interface{}

//Connects to numary, handles auth and basic request stuff
type Connection struct {
	Server string
	client *http.Client
}

func Connect(server string) Connection {
	return Connection{Server: server, client: http.DefaultClient}
}

func (c Connection) request(method, path string) (respFmt, error) {
	req, err := http.NewRequest(method, c.Server+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	b, _ := io.ReadAll(resp.Body) //TODO err
	content := respFmt{}
	json.Unmarshal(b, &content)
	return content, nil
}

func (c Connection) get(path string) (respFmt, error) {
	return c.request("GET", path)
}

func (c Connection) Status() error {
	_, err := c.get("/_info")
	return err
}

func (c Connection) Ledger(name string) Ledger {
	return Ledger{name: name, conn: c}
}

type Ledger struct {
	name string
	conn Connection
}

func (l Ledger) Stats() (stats Stats, isError bool) {
	content, err := l.conn.get("/" + l.name + "/stats")
	if err != nil {
		isError = true
		return
	}
	stats = content["stats"].(Stats)
	isError = !content["ok"].(bool)
	return
}
