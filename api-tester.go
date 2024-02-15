package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Define API URL and API Key as constants
const apiURL = "https://www.example.com/api"
const apiKey = "your_api_key_here"

func main() {
	number_of_iterations := 10
	number_of_api_calls := 5
	delay_in_ms := 1000

	// Define query parameters as a variable
	queryParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	var wg sync.WaitGroup

	for i := 0; i < number_of_iterations; i++ {
		for j := 0; j < number_of_api_calls; j++ {
			wg.Add(1)
			go makeAPICall(&wg, queryParams)
		}

		time.Sleep(time.Duration(delay_in_ms) * time.Millisecond)
	}

	wg.Wait()
}

func prepareAPIRequest(apiURL, apiKey string, queryParams map[string]string) (*http.Request, error) {
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	qs := parsedURL.Query()
	for key, value := range queryParams {
		qs.Add(key, value)
	}
	parsedURL.RawQuery = qs.Encode()

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-API-Key", apiKey)

	return req, nil
}

func executeAPICall(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func saveResponseToFile(body []byte) (string, error) {
	dirName := "load-testing-results"
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.Mkdir(dirName, os.ModePerm); err != nil {
			return "", err
		}
	}

	fileName := fmt.Sprintf("%s/%s.txt", dirName, uuid.New().String())
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = file.Write(body); err != nil {
		return "", err
	}

	return fileName, nil
}

// Updated function signatures to include queryParams
func makeAPICall(wg *sync.WaitGroup, queryParams map[string]string) {
	defer wg.Done()

	req, err := prepareAPIRequest(apiURL, apiKey, queryParams)
	if err != nil {
		fmt.Println("Error preparing request:", err)
		return
	}

	resp, err := executeAPICall(req)
	if err != nil {
		fmt.Println("Error making API call:", err)
		return
	}
	defer resp.Body.Close()

	body, err := readResponseBody(resp)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fileName, err := saveResponseToFile(body)
	if err != nil {
		fmt.Println("Error saving response to file:", err)
		return
	}

	fmt.Println("API call successful. Response written to", fileName)
}
