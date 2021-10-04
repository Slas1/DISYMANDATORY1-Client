package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *ClientAPI) GetCourses(ctx context.Context) (*courseList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/courses", c.BaseUrl), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := courseList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ClientAPI) PostCourse(ctx context.Context, subject course) (bool, error) {
	params := url.Values{}
	params.Add("courseId", strconv.Itoa(int(subject.ID)))
	params.Add("name", subject.Name)
	params.Add("ects", strconv.Itoa(int(subject.Ects)))
	params.Add("courseResponsible", strconv.Itoa(int(subject.CourseResponsible)))
	params.Add("nRatings", strconv.Itoa(int(subject.NRatings)))
	params.Add("avgRatings", fmt.Sprintf("%f", subject.AvgRatings))
	var IDs []string
	IDs = append(IDs, "[")
	for _, i := range subject.ActiveStudents {
		IDs = append(IDs, strconv.Itoa(int(i)))
	}
	IDs = append(IDs, "]")
	params.Add("activeStudents", strings.Join(IDs, ", "))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/courses", c.BaseUrl), strings.NewReader(params.Encode()))
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)

	res := course{}
	if err := c.sendRequest(req, &res); err != nil {
		return false, err
	}

	return true, nil
}

func (c *ClientAPI) GetCourse(ctx context.Context, id int) (*course, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/courses/"+strconv.Itoa(int(id)), c.BaseUrl), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := course{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ClientAPI) PutCourse(ctx context.Context, subject course) (bool, error) {
	params := url.Values{}
	if subject.ID != 0 {
		params.Add("courseId", strconv.Itoa(int(subject.ID)))
	}
	if subject.Name != "" {
		params.Add("name", subject.Name)
	}
	if subject.Ects != 0 {
		params.Add("ects", strconv.Itoa(int(subject.Ects)))
	}
	if subject.CourseResponsible != 0 {
		params.Add("courseResponsible", strconv.Itoa(int(subject.CourseResponsible)))
	}
	if subject.NRatings != 0 {
		params.Add("nRatings", strconv.Itoa(int(subject.NRatings)))
	}
	if subject.AvgRatings != 0 {
		params.Add("avgRatings", fmt.Sprintf("%f", subject.AvgRatings))
	}
	if subject.ActiveStudents != nil {
		var IDs []string
		IDs = append(IDs, "[")
		for _, i := range subject.ActiveStudents {
			IDs = append(IDs, strconv.Itoa(int(i)))
		}
		IDs = append(IDs, "]")
		params.Add("activeStudents", strings.Join(IDs, ", "))
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/courses"+strconv.Itoa(int(subject.ID)), c.BaseUrl), strings.NewReader(params.Encode()))
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)

	res := course{}
	if err := c.sendRequest(req, &res); err != nil {
		return false, err
	}

	return true, nil
}

func (c *ClientAPI) DeleteCourse(ctx context.Context, id int) (bool, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/courses/"+strconv.Itoa(int(id)), c.BaseUrl), nil)
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)

	res := course{}
	if err := c.sendRequest(req, &res); err != nil {
		return false, err
	}

	return true, nil
}

type courseList struct {
	courses []course `json:"courses"`
}

type course struct {
	ID                int32   `json:"courseId"`
	Name              string  `json:"name"`
	Ects              int32   `json:"ects"`
	CourseResponsible int32   `json:"courseResponsibleId"`
	NRatings          int32   `json:"nRatings"`
	AvgRatings        float32 `json:"avgRatings"`
	ActiveStudents    []int32 `json:"activeStudentsId"`
}
