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
	BaseURLV1 = "http://localhost:8080"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func main() {
	c := NewClient()

	printCurrentCourses(c)

	fmt.Println("Adding course")
	addcourse := new(course)
	addcourse.ID = 5
	addcourse.Name = "Test"
	addcourse.Ects = 200
	addcourse.CourseResponsible = 6
	addcourse.NRatings = 10000
	addcourse.AvgRatings = 0.1
	addcourse.ActiveStudents = make([]int32, 3)

	c.PostCourse(context.Background(), *addcourse)

	printCurrentCourses(c)

	fmt.Println("Edditing course")
	editcourse := new(course)
	editcourse.ID = 5
	editcourse.Name = "Test2"
	editcourse.Ects = 400
	editcourse.CourseResponsible = 6
	editcourse.AvgRatings = 0.2
	editcourse.ActiveStudents = make([]int32, 6)

	c.UpdateCourse(context.Background(), *editcourse)

	printCurrentCourses(c)

	fmt.Println("Deleteing course")
	c.DeleteCourse(context.Background(), 5)

	printCurrentCourses(c)

}

func NewClient() *Client {
	return &Client{
		BaseURL: BaseURLV1,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) (interface{}, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req.Close = true

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {

			fmt.Println("sendRequest error 2")
			return nil, errors.New(errRes.Message)
		}

		fmt.Println("sendRequest error 3")
		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	switch vType := v.(type) {
	case []course:
		fullResponse := []course{}
		if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {

			fmt.Print(req.Method + ": ")
			fmt.Println(err)
			return nil, err
		}

		//fmt.Println(fullResponse)
		return fullResponse, nil
	case *course:
		fullResponse := course{}
		if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {

			fmt.Print(req.Method + ": ")
			fmt.Println(err)
			return nil, err
		}

		//fmt.Println(fullResponse)
		return fullResponse, nil
	default:
		return nil, fmt.Errorf("unknown type: %t", vType)
	}
}
