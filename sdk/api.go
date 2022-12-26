package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoreResponse struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type Response struct {
	CoreResponse
	Data string `json:"data"`
}

func NewResponse(id string, status int, bodyBytes []byte, err error) (resp *Response) {
	// if there is an error, return the error
	if err != nil {
		return &Response{
			CoreResponse: CoreResponse{
				ID:     id,
				Status: status,
			},
			Data: err.Error(),
		}
	} else {
		// if there is no error, return the body and data
		if status == http.StatusOK {
			return &Response{
				CoreResponse: CoreResponse{
					ID:     id,
					Status: status,
				},
				Data: string(bodyBytes),
			}
		} else {
			// if not OK, return the body as an http status error
			return &Response{
				CoreResponse: CoreResponse{
					ID:     id,
					Status: status,
				},
				Data: err.Error(),
			}
		}
	}
}

func (r *Response) UnmarshalData(v interface{}) error {
	return json.Unmarshal([]byte(r.Data), v)
}

func (r *Response) UnmarshalError(v interface{}) error {
	return json.Unmarshal([]byte(r.Data), v)
}

func (r *Response) UnmarshalCore(v interface{}) error {
	return json.Unmarshal([]byte(r.Data), v)
}

func (r *Response) Unmarshal(v interface{}) error {
	return json.Unmarshal([]byte(r.Data), v)
}

func (r *Response) GetCoreResponse() *CoreResponse {
	return &r.CoreResponse
}

func (r *Response) GetResponse() *Response {
	return r
}

func (r *Response) GetStatus() int {
	return r.Status
}

func (r *Response) GetData() string {
	return r.Data
}

func (r *Response) GetID() string {
	return r.ID
}

func (r *Response) GetError() error {
	return fmt.Errorf(r.Data)
}

func (r *Response) GetCoreError() error {
	return fmt.Errorf(r.Data)
}

func requestAPI() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "basic;bearer/jwt") // TODO: add auth (encoded with wallet address, password & device id)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	//var responseObject Response
	responseObject := NewResponse("1", resp.StatusCode, bodyBytes, err)
	//json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
