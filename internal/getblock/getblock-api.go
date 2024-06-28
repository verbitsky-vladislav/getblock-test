package getblock

import (
	"bytes"
	"encoding/json"
	"getblock-test/internal/utils/types"
	"getblock-test/internal/utils/types/eth"
	logger "getblock-test/pkg"
	"net/http"
)

type GetBlockClientService interface {
	GetLastBlockNumber() (*types.GetBlockBaseResponse[string], error)
	GetBlock(blockNumber string) (*types.GetBlockBaseResponse[eth.EthBlock], error)
}

type getBlockClient struct {
	baseURL string
	aPIKey  string
	client  *http.Client
}

func NewGetBlockClient(apiKey string) *getBlockClient {
	return &getBlockClient{
		baseURL: "https://go.getblock.io/" + apiKey + "/",
		client:  &http.Client{},
	}
}

func (c *getBlockClient) doPostRequest(reqBody interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, logger.Error(err, "JSON marshal error")
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, logger.Error(err, "HTTP request creation error")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, logger.Error(err, "HTTP request execution error")
	}

	return resp, nil
}

func (c *getBlockClient) decodeResponse(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return logger.Error(err, "JSON decode error")
	}
	return nil
}

func (c *getBlockClient) GetLastBlockNumber() (*types.GetBlockBaseResponse[string], error) {
	reqBody := types.GetBlockBaseRequest{
		Id:      "getblock.io",
		JsonRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
	}

	resp, err := c.doPostRequest(reqBody)
	if err != nil {
		return nil, err
	}

	var response types.GetBlockBaseResponse[string]
	if err := c.decodeResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *getBlockClient) GetBlock(blockNumber string) (*types.GetBlockBaseResponse[eth.EthBlock], error) {
	reqBody := types.GetBlockBaseRequest{
		Id:      "getblock.io",
		JsonRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params: []interface{}{
			blockNumber, true,
		},
	}

	resp, err := c.doPostRequest(reqBody)
	if err != nil {
		return nil, err
	}

	var response types.GetBlockBaseResponse[eth.EthBlock]
	if err := c.decodeResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
