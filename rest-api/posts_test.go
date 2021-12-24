package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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

func TestRetrieveSingleTask(t *testing.T) {
	testMethod := http.MethodGet
	testPath := "/tasks/"

	tests := []struct {
		name       string
		id         string
		expected   Post
		statusCode int
	}{
		{
			name: "status ok",
			id:   "1",
			expected: Post{
				Id:      "1",
				Title:   "Created Post",
				Content: "What's up guys!",
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "status not found",
			id:         "2",
			expected:   Post{},
			statusCode: http.StatusNotFound,
		},
	}

	pm["1"] = &Post{
		Id:      "1",
		Title:   "Created Post",
		Content: "What's up guys!",
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath+tt.id, nil)
			testResp := httptest.NewRecorder()

			vars := map[string]string{
				"id": tt.id,
			}

			testReq = mux.SetURLVars(testReq, vars)

			RetrieveSinglePost(testResp, testReq)
			var got Post
			body, err := ioutil.ReadAll(testResp.Body)

			if err != nil {
				panic(err)
			}

			json.Unmarshal(body, &got)

			assert.Equal(t, tt.statusCode, testResp.Result().StatusCode)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	testMethod := http.MethodPut
	testPath := "/tasks"

	tests := []struct {
		name            string
		post            *Post
		id              string
		expectedId      string
		expectedTitle   string
		expectedContent string
		statusCode      int
	}{
		{
			name: "status ok",
			post: &Post{
				Title:   "Updated Post",
				Content: "Yo yo",
			},
			id: "1",
			expectedId:      "1",
			expectedTitle:   "Updated Post",
			expectedContent: "Yo yo",
			statusCode:      http.StatusOK,
		},
		{
			name: "status not found",
			post: &Post{
				Title:   "Updated Post",
				Content: "What's up guys!",
			},
			id: "2",
			expectedId:      "",
			expectedTitle:   "",
			expectedContent: "",
			statusCode:      http.StatusNotFound,
		},
		{
			name: "status bad request",
			post: &Post{
				Title:   "",
				Content: "this might fail...",
			},
			id: "1",
			expectedId:      "",
			expectedTitle:   "",
			expectedContent: "",
			statusCode:      http.StatusBadRequest,
		},
	}

	pm["1"] = &Post{
		Id:      "1",
		Title:   "Created Post",
		Content: "What's up guys!",
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.post)
			if err != nil {
				panic(err)
			}
			b := bytes.NewBuffer(body)

			testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath+tt.id, b)
			testResp := httptest.NewRecorder()

			vars := map[string]string{
				"id": tt.id,
			}

			testReq = mux.SetURLVars(testReq, vars)

			UpdatePost(testResp, testReq)
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

func TestDeletePost(t *testing.T) {
	testMethod := http.MethodDelete
	testPath := "/tasks"

	tests := []struct {
		name            string
		id              string
		statusCode      int
	}{
		{
			name: "status ok",
			id: "1",
			statusCode:      http.StatusOK,
		},
		{
			name: "status not found",
			id: "2",
			statusCode:      http.StatusNotFound,
		},
	}

	pm["1"] = &Post{
		Id:      "1",
		Title:   "Created Post",
		Content: "What's up guys!",
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath+tt.id, nil)
			testResp := httptest.NewRecorder()

			vars := map[string]string{
				"id": tt.id,
			}
			testReq = mux.SetURLVars(testReq, vars)

			DeletePost(testResp, testReq)

			assert.Equal(t, tt.statusCode, testResp.Result().StatusCode)
		})
	}
}

func TestRetrieveAllPosts(t *testing.T) {
	testMethod := http.MethodGet
	testPath := "/tasks"

	pm["3"] = &Post{
		Id:      "3",
		Title:   "Created Post 3",
		Content: "Hello",
	}
	pm["1"] = &Post{
		Id:      "1",
		Title:   "Created Post",
		Content: "What's up guys!",
	}
	pm["4"] = &Post{
		Id:      "4",
		Title:   "No wAAY",
		Content: "Sup",
	}

	expected := []Post{
		{
			Id:      "1",
			Title:   "Created Post",
			Content: "What's up guys!",
		},
		{
			Id:      "3",
			Title:   "Created Post 3",
			Content: "Hello",
		},
		{
			Id:      "4",
			Title:   "No wAAY",
			Content: "Sup",
		},
	}

	t.Run("status ok", func(t *testing.T) {
		testReq := httptest.NewRequest(testMethod, "http://localhost:8080"+testPath, nil)
		testResp := httptest.NewRecorder()

		RetrieveAllPosts(testResp, testReq)
		var got = []Post{}

		body, err := ioutil.ReadAll(testResp.Body)

		if err != nil {
			panic(err)
		}

		json.Unmarshal(body, &got)

		assert.Equal(t, http.StatusOK, testResp.Result().StatusCode)
		assert.Equal(t, expected, got)
	})
}