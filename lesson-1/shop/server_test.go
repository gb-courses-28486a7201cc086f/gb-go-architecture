package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gb-go-architecture/lesson-1/shop/models"
	"gb-go-architecture/lesson-1/shop/repository"
)

const httpTimeout = 30 * time.Second

func doTestReq(client *http.Client, method, url string, payload io.Reader) (resp *http.Response, body *bytes.Buffer, err error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return resp, body, fmt.Errorf("Cannot make test request, err:  %q", err.Error())
	}

	resp, err = client.Do(req)
	if err != nil {
		return resp, body, fmt.Errorf("Request error, got: %q, expected <nil>", err.Error())
	}
	defer resp.Body.Close()

	body = bytes.NewBuffer([]byte{})
	_, err = body.ReadFrom(resp.Body)
	return resp, body, err
}

func passHTTPStatus(resp *http.Response, expectedStatus int) (ok bool, msg string) {
	if resp == nil {
		return false, fmt.Sprint("Bad response, got nil, expected valid response")
	}
	if resp.StatusCode != expectedStatus {
		return false, fmt.Sprintf("Bad status code, got %v, expected %v", resp.StatusCode, expectedStatus)
	}
	return true, ""
}

func TestSuccessHTTPHandlers(t *testing.T) {
	repo := repository.NewMapDBMock()
	s := &server{rep: repo}

	repoItems, err := repo.ListItems(&repository.ItemFilter{Limit: 2})
	if err != nil {
		t.Error("cannot setup test: ", err)
	}
	if len(repoItems) == 0 {
		t.Error("cannot setup test: mock repo is empty")
	}
	itemToUpdate := repoItems[0]
	itemToDelete := repoItems[1]

	// start test server
	router := NewRouter(s)
	server := httptest.NewServer(router)
	defer server.Close()

	client := &http.Client{Timeout: httpTimeout}

	t.Run("createItemHandler", func(t *testing.T) {
		expectedItem := *&models.Item{
			Name:  "create handler test",
			Price: 100,
		}
		testPayload, err := json.Marshal(expectedItem)
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}

		resp, body, err := doTestReq(client, http.MethodPost, server.URL+"/item", bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusOK)
		if !ok {
			t.Error(msg)
			return
		}

		result := &models.Item{}
		if err := json.NewDecoder(body).Decode(result); err != nil {
			t.Error("cannot decode response, error: ", err)
			return
		}

		if expectedItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if expectedItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	})

	t.Run("getItemHandler", func(t *testing.T) {
		expectedItem := itemToUpdate
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, body, err := doTestReq(client, http.MethodGet, testURL, nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusOK)
		if !ok {
			t.Error(msg)
			return
		}

		result := &models.Item{}
		if err := json.NewDecoder(body).Decode(result); err != nil {
			t.Error("cannot decode response, error: ", err)
			return
		}

		if expectedItem.ID != result.ID {
			t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, result.ID)
		}
		if expectedItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if expectedItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	})

	t.Run("listItemHandler empty params", func(t *testing.T) {
		resp, body, err := doTestReq(client, http.MethodGet, server.URL+"/item", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusOK)
		if !ok {
			t.Error(msg)
			return
		}

		result := &ListItemResponse{}
		if err := json.NewDecoder(body).Decode(result); err != nil {
			t.Error("cannot decode response, error: ", err)
			return
		}

		if len(result.Payload) == 0 || result.Limit == 0 {
			// decode result should not be default ListItemResponse.
			t.Error("unexpected response, expected non empty payload")
		}
	})

	t.Run("listItemHandler with params", func(t *testing.T) {
		testURL := fmt.Sprintf("%s/item?price_left=0&price_right=100&limit=1&offset=1", server.URL)
		resp, body, err := doTestReq(client, http.MethodGet, testURL, nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusOK)
		if !ok {
			t.Error(msg)
			return
		}

		result := &ListItemResponse{}
		if err := json.NewDecoder(body).Decode(result); err != nil {
			t.Error("cannot decode response, error: ", err)
			return
		}

		if len(result.Payload) == 0 || result.Limit == 0 {
			// decode result should not be default ListItemResponse.
			t.Error("unexpected response, expected non empty payload")
		}
	})

	t.Run("deleteItemHandler", func(t *testing.T) {
		expectedItem := itemToDelete
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, _, err := doTestReq(client, http.MethodDelete, testURL, nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusNoContent)
		if !ok {
			t.Error(msg)
			return
		}
	})

	t.Run("updateItemHandler", func(t *testing.T) {
		expectedItem := *&models.Item{
			ID:    itemToUpdate.ID,
			Name:  itemToUpdate.Name,
			Price: 100,
		}
		testPayload, err := json.Marshal(expectedItem)
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, body, err := doTestReq(client, http.MethodPut, testURL, bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusOK)
		if !ok {
			t.Error(msg)
			return
		}

		result := &models.Item{}
		if err := json.NewDecoder(body).Decode(result); err != nil {
			t.Error("cannot decode response, error: ", err)
			return
		}

		if expectedItem.ID != result.ID {
			t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, result.ID)
		}
		if expectedItem.Name != result.Name {
			t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, result.Name)
		}
		if expectedItem.Price != result.Price {
			t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, result.Price)
		}
	})
}
func TestUserFailHTTPHandlers(t *testing.T) {
	repo := repository.NewMapDBMock()
	s := &server{rep: repo}

	// start test server
	router := NewRouter(s)
	server := httptest.NewServer(router)
	defer server.Close()

	client := &http.Client{Timeout: httpTimeout}

	t.Run("createItemHandler 400", func(t *testing.T) {
		testPayload, err := json.Marshal("{}")
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}

		resp, _, err := doTestReq(client, http.MethodPost, server.URL+"/item", bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusBadRequest)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("getItemHandler 400", func(t *testing.T) {
		resp, _, err := doTestReq(client, http.MethodGet, server.URL+"/item/badID", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusBadRequest)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("deleteItemHandler 400", func(t *testing.T) {
		resp, _, err := doTestReq(client, http.MethodDelete, server.URL+"/item/1", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusNoContent)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("updateItemHandler 400", func(t *testing.T) {
		testPayload, err := json.Marshal("{}")
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}

		resp, _, err := doTestReq(client, http.MethodPut, server.URL+"/item/1", bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusBadRequest)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("getItemHandler 404", func(t *testing.T) {
		resp, _, err := doTestReq(client, http.MethodGet, server.URL+"/item/-1", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusNotFound)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("deleteItemHandler 404", func(t *testing.T) {
		resp, _, err := doTestReq(client, http.MethodDelete, server.URL+"/item/-1", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusNotFound)
		if !ok {
			t.Error(msg)
		}
	})
}

func TestServerFailHTTPHandlers(t *testing.T) {
	repo := repository.NewMapDBErrorMock()
	s := &server{rep: repo}

	expectedItem := *&models.Item{
		ID:    100,
		Name:  "test item",
		Price: 100,
	}

	// start test server
	router := NewRouter(s)
	server := httptest.NewServer(router)
	defer server.Close()

	client := &http.Client{Timeout: httpTimeout}

	t.Run("createItemHandler", func(t *testing.T) {
		testPayload, err := json.Marshal(expectedItem)
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}

		resp, _, err := doTestReq(client, http.MethodPost, server.URL+"/item", bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusInternalServerError)
		if !ok {
			t.Error(msg)
			return
		}
	})

	t.Run("getItemHandler", func(t *testing.T) {
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, _, err := doTestReq(client, http.MethodGet, testURL, nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusInternalServerError)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("listItemHandler", func(t *testing.T) {
		resp, _, err := doTestReq(client, http.MethodGet, server.URL+"/item", nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusInternalServerError)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("deleteItemHandler", func(t *testing.T) {
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, _, err := doTestReq(client, http.MethodDelete, testURL, nil)
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusInternalServerError)
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("updateItemHandler", func(t *testing.T) {
		testPayload, err := json.Marshal(expectedItem)
		if err != nil {
			t.Error("cannot create test payload: ", err)
			return
		}
		testURL := fmt.Sprintf("%s/item/%d", server.URL, expectedItem.ID)

		resp, _, err := doTestReq(client, http.MethodPut, testURL, bytes.NewBuffer(testPayload))
		if err != nil {
			t.Fatal("unexpected test server error: ", err.Error())
		}

		ok, msg := passHTTPStatus(resp, http.StatusInternalServerError)
		if !ok {
			t.Error(msg)
		}
	})
}
