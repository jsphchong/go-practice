package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	testMethod := http.MethodPost
	testPath := "/tasks"

	tests := []struct {
		name            string
		post            *Post
		expectedId      string
		expectedTitle   string
		expectedContent string
		statusCode      int
	}{
		{
			name: "status ok",
			post: &Post{
				Title:   "Created Post",
				Content: "What's up guys!",
			},
			expectedId:      "1",
			expectedTitle:   "Created Post",
			expectedContent: "What's up guys!",
			statusCode:      http.StatusOK,
		},
		{
			name: "status bad request",
			post: &Post{
				Title:   "",
				Content: "What's up guys!",
			},
			expectedId:      "",
			expectedTitle:   "",
			expectedContent: "",
			statusCode:      http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.post)
			if err != nil {
				panic(err)
			}
			b := bytes.NewBuffer(body)
			testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath, b)
			testResp := httptest.NewRecorder()

			CreatePost(testResp, testReq)

			var got Post
			body, err = ioutil.ReadAll(testResp.Body)

			if err != nil {
				panic(err)
			}

			json.Unmarshal(body, &got)

			assert.Equal(t, tt.statusCode, testResp.Result().StatusCode)
			assert.Equal(t, tt.expectedId, got.Id)
			assert.Equal(t, tt.expectedTitle, got.Title)
			assert.Equal(t, tt.expectedContent, got.Content)
		})
	}
}

// func TestRetrieveSingleTask(t *testing.T) {
// 	testMethod := http.MethodGet
// 	testPath := "/tasks/"

// 	tests := []struct {
// 		name       string
// 		id         string
// 		expected   Post
// 		statusCode int
// 	}{
// 		{
// 			name: "status ok",
// 			id: "1",
// 			expected: Post{},
// 			statusCode: http.StatusOK,
// 		},
// 		{
// 			name: "status not found",
// 			id: "1",
// 			expected: Post{},
// 			statusCode: http.StatusNotFound,
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt

// 		t.Run(tt.name, func(t *testing.T) {
// 			body, err := json.Marshal(tt.post)
// 			if err != nil {
// 				panic(err)
// 			}
// 			b := bytes.NewBuffer(body)
// 			testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath, b)
// 			testResp := httptest.NewRecorder()

// 			CreatePost(testResp, testReq)

// 			var got Post
// 			body, err = ioutil.ReadAll(testResp.Body)

// 			if err != nil {
// 				panic(err)
// 			}

// 			json.Unmarshal(body, &got)

// 			assert.Equal(t, tt.statusCode, testResp.Result().StatusCode)
// 			assert.Equal(t, tt.expectedId, got.Id)
// 			assert.Equal(t, tt.expectedTitle, got.Title)
// 			assert.Equal(t, tt.expectedContent, got.Content)
// 		})
// 	}
// }
