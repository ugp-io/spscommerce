package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	AppID       string
	Secret      string
	AccessToken *string
	ExpiresAt   *time.Time
	Transaction TransactionService
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(appID, secret string) *Client {
	c := &Client{
		AppID:  appID,
		Secret: secret,
	}

	c.Transaction = &TransactionServiceOp{client: c}

	return c
}

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {
	if c.AccessToken == nil || c.ExpiresAt == nil || time.Now().After(*c.ExpiresAt) {
		if err := c.GetAccessToken(); err != nil {
			return fmt.Errorf("failed to get access token: %w", err)
		}
	}

	var bodyReader io.Reader
	if body != nil {
		requestJson, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(requestJson)
	}

	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}
	httpReq.Header.Set("Authorization", "Bearer "+*c.AccessToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	fmt.Println("Resp:", string(respBody))
	fmt.Println(url)
	decoder := json.NewDecoder(bytes.NewReader(respBody))
	errDecode := decoder.Decode(&v)
	if errDecode != nil {
		return errDecode
	}

	return nil
}

func (c *Client) GetAccessToken() error {

	body := map[string]string{
		"audience":      "https://spscommerce.com",
		"client_id":     c.AppID,
		"client_secret": c.Secret,
		"grant_type":    "client_credentials",
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://auth.spscommerce.com/oauth/token", bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token request returned %d: %s", resp.StatusCode, string(respBody))
	}

	tokenResp := make(map[string]interface{})
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if accessToken, ok := tokenResp["access_token"].(string); ok {
		expiration := time.Now().Add(time.Duration(tokenResp["expires_in"].(float64)) * time.Second)
		c.ExpiresAt = &expiration
		c.AccessToken = &accessToken
	}

	return nil
}
