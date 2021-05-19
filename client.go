package bitcoinRpc

import (
	"github.com/evgeny-klyopov/bitcoin-rpc/response"
	jsonRpcClient "github.com/evgeny-klyopov/golang-json-rpc"
	"net/http"
)

type client struct {
	jsonRpcClient jsonRpcClient.JsonRpcConnector
}

type BitcoinRpc interface {
	GetBalance(params []string) (*response.Balance, error)
	ListSinceBlock(params []string) (*response.ListSinceBlock, error)
	GetTransaction(params []string) (*response.GetTransaction, error)
}

func NewClient(jsonRpcClient jsonRpcClient.JsonRpcConnector) BitcoinRpc {
	return &client{
		jsonRpcClient: jsonRpcClient,
	}
}

func (c *client) ListSinceBlock(params []string) (*response.ListSinceBlock, error) {
	var result response.ListSinceBlock

	err := c.request("listsinceblock", params, &result)

	return &result, err
}
func (c *client) GetTransaction(params []string) (*response.GetTransaction, error) {
	var result response.GetTransaction

	err := c.request("gettransaction", params, &result)

	return &result, err
}

func (c *client) GetBalance(params []string) (*response.Balance, error) {
	var balance response.Balance

	err := c.request("getbalance", params, &balance)
	return &balance, err
}

func (c *client) request(method string, params interface{}, data interface{}) error {
	statusCode, err := c.jsonRpcClient.Request(method, params, &data)

	if *(statusCode) != http.StatusOK {
		return err
	}

	return nil
}
