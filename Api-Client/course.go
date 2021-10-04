package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type course struct {
	ID                int32   `json:"courseId"`
	Name              string  `json:"name"`
	Ects              int32   `json:"ects"`
	CourseResponsible int32   `json:"courseResponsibleId"`
	NRatings          int32   `json:"nRatings"`
	AvgRatings        float32 `json:"avgRatings"`
	ActiveStudents    []int32 `json:"activeStudents"`
}

func (c *Client) GetCourses(ctx context.Context) (interface{}, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/courses", c.BaseURL), nil)
	if err != nil {
		fmt.Println("Here error 1")
		return nil, err
	}

	req = req.WithContext(ctx)

	res := make([]course, 0)
	response, err := c.sendRequest(req, res)
	if err != nil {
		fmt.Println("Here error 2")
		return nil, err
	}

	return response, nil
}

func (c *Client) PostCourse(ctx context.Context, data course) (interface{}, error) {

	courseJSON, JSONerr := json.Marshal(data)

	if JSONerr != nil {
		return false, JSONerr
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/courses", c.BaseURL), strings.NewReader(string(courseJSON)))
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)

	var res course
	response, err := c.sendRequest(req, &res)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return response, nil

}

func (c *Client) GetCourse(ctx context.Context, id int) (interface{}, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/courses/"+strconv.Itoa(id), c.BaseURL), nil)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false, err
	}

	req = req.WithContext(ctx)

	var res course
	response, err := c.sendRequest(req, &res)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return false, err
	}

	return response, nil
}

func (c *Client) DeleteCourse(ctx context.Context, id int) (interface{}, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/courses/"+strconv.Itoa(id), c.BaseURL), nil)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false, err
	}

	req = req.WithContext(ctx)

	res := make([]course, 0)
	response, err := c.sendRequest(req, res)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return false, err
	}

	return response, nil
}

func (c *Client) UpdateCourse(ctx context.Context, data course) (interface{}, error) {
	//

	courseJSON, JSONerr := json.Marshal(data)
	if JSONerr != nil {
		return false, JSONerr
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/courses/"+strconv.Itoa(int(data.ID)), c.BaseURL), strings.NewReader(string(courseJSON)))
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return false, err
	}

	req = req.WithContext(ctx)

	var res course
	response, err := c.sendRequest(req, &res)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return false, err
	}

	return response, nil
}

func printCurrentCourses(c *Client) {
	courses, _ := c.GetCourses(context.Background())

	courses1 := courses.([]course)

	fmt.Println("Current courses:")
	for i := 0; i < len(courses1); i++ {
		fmt.Println(courses1[i])
	}
}
