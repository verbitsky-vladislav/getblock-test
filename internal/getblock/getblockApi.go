package getblock

import (
	"bytes"
	"encoding/json"
	"errors"
	"getblock-test/internal/utils/types"
	"getblock-test/internal/utils/types/eth"
	"net/http"
)

type GetBlockClient struct {
	baseURL string
	client  *http.Client
}

func NewGetBlockClient(apiKey string) *GetBlockClient {
	return &GetBlockClient{
		baseURL: "https://go.getblock.io/" + apiKey + "/",
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
			},
		},
	}
}

func (c *GetBlockClient) doPostRequest(reqBody interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return c.client.Do(req)
}

func (c *GetBlockClient) decodeResponse(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected HTTP status: " + resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *GetBlockClient) GetLastBlockNumber() (*types.GetBlockBaseResponse[string], error) {
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

func (c *GetBlockClient) GetBlock(blockNumber string) (*types.GetBlockBaseResponse[eth.EthBlock], error) {
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
