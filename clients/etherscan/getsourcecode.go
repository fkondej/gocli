package etherscan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// API: contract -> getsourcecode
type EtherscanGetSourcecodeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		SourceCode           string `json:"SourceCode"`
		ABI                  string `json:"ABI"`
		ContractName         string `json:"ContractName"`
		CompilerVersion      string `json:"CompilerVersion"`
		ConstructorArguments string `json:"ConstructorArguments"`
	} `json:"result"`
}

func (c *EtherscanClient) GetSourcecode(ctx context.Context, hexAddress string) (*EtherscanGetSourcecodeResponse, string, error) {
	downloadURL := fmt.Sprintf(
		"%s?module=contract&action=getsourcecode&address=%s",
		c.apiURL, hexAddress,
	)
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return nil, downloadURL, fmt.Errorf("Failed rate limiter for Get Sourcecode for smart contract: %s. %w", hexAddress, err)
	}
	downloadURLWithApikey := fmt.Sprintf("%s&apikey=%s", downloadURL, c.apiKey)
	resp, err := http.Get(downloadURLWithApikey)
	if err != nil {
		return nil, downloadURL, fmt.Errorf("Failed to Get Sourcecode for smart contract: %s. %w", hexAddress, err)
	}
	defer resp.Body.Close()
	var payload EtherscanGetSourcecodeResponse
	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, downloadURL, fmt.Errorf("Failed to parse Get Sourcecode response to json: %s. %w", hexAddress, err)
	}
	if payload.Status != "1" {
		return &payload, downloadURL, fmt.Errorf("Get Sourcecode response Status is not 1: %s. Error: %s", payload.Status, payload.Message)
	}
	return &payload, downloadURL, nil
}
