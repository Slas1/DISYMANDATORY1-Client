package main

/*
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "localhost:8080"
)

func main() {
	var client = newClient()

	courses, _ := client.GetCourses(context.Background())

	fmt.Println("Dor jeg her")

	a := len(courses.courses)

	a += 1
}

type ClientAPI struct {
	BaseUrl string
	//apiKey     string
	HTTPClient *http.Client
}

func newClient() *ClientAPI {
	return &ClientAPI{
		BaseUrl: BaseURLV1,
		//apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *ClientAPI) sendRequest(req *http.Request, v interface{}) (interface{}, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}

	return fullResponse, nil
}
*/
