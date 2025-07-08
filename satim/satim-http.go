package satim

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SatimHttpClient struct {
	endpoint string
}

func newSatimHttpClient(endpoint string) *SatimHttpClient {
	return &SatimHttpClient{
		endpoint: endpoint,
	}
}

func (c *SatimHttpClient) RegisterOrderQuery(query map[string]interface{}) (*SatimRegisterOrderResponse, error) {
	// Build query parameters
	params := url.Values{}
	for key, value := range query {
		// if string set it if int convert it if array json it
		params.Set(key, fmt.Sprintf("%v", value))
	}

	// Final URL
	fullURL := fmt.Sprintf("%s/register.do?%s", c.endpoint, params.Encode())

	// Perform GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response SatimRegisterOrderResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if response.ErrorCode == nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if *response.ErrorCode == "0" {
		return &response, nil
	}

	return nil, errors.New("failed to register order")
}

func (c *SatimHttpClient) GetOrderStatusQuery(query map[string]interface{}) (*SatimOrderStatusResponse, error) {
	// Build query parameters
	params := url.Values{}
	for key, value := range query {
		// if string set it if int convert it if array json it
		params.Set(key, fmt.Sprintf("%v", value))
	}

	// Final URL
	fullURL := fmt.Sprintf("%s/getOrderStatus.do?%s", c.endpoint, params.Encode())

	// Perform GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response SatimOrderStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if response.ErrorCode == nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if *response.ErrorCode == "0" {
		return &response, nil
	}

	if *response.ErrorCode == "2" {
		return &response, nil
	}

	return nil, errors.New("failed to get order status")
}

func (c *SatimHttpClient) ConfirmOrderQuery(query map[string]interface{}) (*SatimOrderConfirmResponse, error) {
	// Build query parameters
	params := url.Values{}
	for key, value := range query {
		// if string set it if int convert it if array json it
		params.Set(key, fmt.Sprintf("%v", value))
	}

	// Final URL
	fullURL := fmt.Sprintf("%s/confirmOrder.do?%s", c.endpoint, params.Encode())

	// Perform GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response SatimOrderConfirmResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if response.ErrorCode == nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if *response.ErrorCode == "0" {
		return &response, nil
	}

	if *response.ErrorCode == "2" {
		return &response, errors.New("order already confirmed")
	}

	return nil, errors.New("failed to confirm order")
}

func (c *SatimHttpClient) RefundOrderQuery(query map[string]interface{}) (*SatimOrderRefundResponse, error) {
	// Build query parameters
	params := url.Values{}
	for key, value := range query {
		// if string set it if int convert it if array json it
		params.Set(key, fmt.Sprintf("%v", value))
	}

	// Final URL
	fullURL := fmt.Sprintf("%s/refund.do?%s", c.endpoint, params.Encode())

	// Perform GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response SatimOrderRefundResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if response.ErrorCode == nil {
		return nil, errors.New("failed to unmarshal response body")
	}

	if *response.ErrorCode == "0" {
		return &response, nil
	}

	return nil, errors.New("failed to refund order")
}
