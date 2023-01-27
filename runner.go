package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
)

func run(data Data, config Config) {
	requests := filter(data, config)
	rq := convert(data, requests)

	for _, r := range rq {
		err := execute(r, config.Verbose)
		if err == nil {
			continue
		}
	}
}

func execute(requestData RQ, verbose bool) error {
	client := http.Client{}
	defer client.CloseIdleConnections()
	var body io.Reader
	if requestData.Body != nil && (len(*requestData.Body) > 0) {
		body.Read([]byte(*requestData.Body))
	}

	req, _ := http.NewRequest(strings.ToUpper(requestData.Method), requestData.URL, body)
	for k, v := range requestData.Headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.Body != nil {
		body, err := httputil.DumpResponse(resp, true)
		out(requestData.Name, body)
		return err
	}

	return nil
}

func out(name *string, body []byte) {
	if name != nil {
		println(*name)
	}

	println(string(body))
}
