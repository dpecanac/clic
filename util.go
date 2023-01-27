package main

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func read(fileName string) (*Data, error) {
	f, err := os.ReadFile(DEFAULT_DATA_FILE)
	if err != nil {
		panic(err)
	}

	data := Data{}
	err = yaml.Unmarshal(f, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func filter(requestData Data, config Config) []SR {
	requests := []SR{}

	names := stringToArray(&config.Name)
	if len(names) > 0 {
		for _, r := range requestData.Requests {
			if r.Name == nil {
				continue
			}

			if inArray(*r.Name, names) {
				requests = append(requests, r)
			}
		}
	}

	tags := stringToArray(&config.Tags)
	if len(tags) > 0 {
		for _, r := range requestData.Requests {
			if r.Tags == nil {
				continue
			}

			if len(*r.Tags) > 0 {
				requests = append(requests, r)
			}
		}
	}

	if len(names) == 0 && len(tags) == 0 {
		requests = requestData.Requests
	}

	println("Requests to execute:", len(requests))
	return requests
}

func convert(data Data, requests []SR) []RQ {
	rq := []RQ{}
	for _, r := range requests {
		urls := []string{}
		if data.BaseURL == nil && r.Endpoint == nil {
			continue
		}

		if data.BaseURL != nil {
			urls = append(urls, *data.BaseURL)
		}

		if r.Endpoint != nil {
			urls = append(urls, *r.Endpoint)
		}

		url, err := createRequestURL(urls)
		if err != nil {
			continue
		}

		headers, err := createRequestHeaders(data.Headers, r.Headers)
		if err != nil {
			continue
		}

		rq = append(rq, RQ{
			Name:    r.Name,
			URL:     *url,
			Headers: *headers,
			Method:  r.Method,
		})
	}

	return rq
}

func createRequestURL(args []string) (*string, error) {
	if len(args) == 0 {
		return nil, errors.New("no arguments")
	}

	URL := ""
	for _, url := range args {
		URL += url
	}

	return &URL, nil
}

func createRequestHeaders(gHeaders *[]string, rHeaders *[]string) (*map[string]string, error) {
	headers := map[string]string{}

	if gHeaders != nil {
		for _, header := range *gHeaders {
			h := strings.Split(header, ":")
			headers[h[0]] = h[1]
		}
	}

	if rHeaders != nil {
		for _, header := range *rHeaders {
			h := strings.Split(header, ":")
			headers[h[0]] = h[1]
		}
	}

	return &headers, nil
}

// InArray checks if a given string exists in a list of strings.
func inArray(element string, list []string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}

func stringToArray(names *string) []string {
	if names == nil || len(*names) == 0 {
		return []string{}
	}

	return strings.Split(*names, ",")
}
