package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// ServerVersion returns information of the docker client and server host.
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error) {
	resp, err := cli.get(ctx, "/version", nil, nil)
	if err != nil {
		return types.Version{}, err
	}

	var server types.Version
	err = json.NewDecoder(resp.body).Decode(&server)
	ensureReaderClosed(resp)
	return server, err
}

// ServerVersionWithRaw returns information of the docker client and server host
// as well as the raw response from the server.
func (cli *Client) ServerVersionWithRaw(ctx context.Context) (types.Version, []byte, error) {
	resp, err := cli.get(ctx, "/version", nil, nil)
	if err != nil {
		return types.Version{}, nil, err
	}
	defer ensureReaderClosed(resp)

	body, err := ioutil.ReadAll(resp.body)
	if err != nil {
		return types.Version{}, nil, err
	}

	var server types.Version
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&server)
	return server, body, err
}
