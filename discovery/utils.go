package discovery

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func SetInterval(callback func(), interval time.Duration) chan bool {
	done := make(chan bool, 1)
	go func() {
		callback()
		timer := time.NewTicker(interval)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				callback()
			case <-done:
				return
			}
		}
	}()
	return done
}

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func Post(url string, data []byte) (res []byte, err error) {
	var response *http.Response
	buffer := bytes.NewBuffer(data)
	if response, err = http.Post(url, "application/json;charset=utf-8", buffer); err == nil {
		var body []byte
		if body, err = ioutil.ReadAll(response.Body); err == nil {
			if response.StatusCode != http.StatusOK {
				err = errors.New(string(body))
			} else {
				res = body
			}
		}
		response.Body.Close()
	}
	return res, err
}

func PostBasicAuth(url string, username string, userpwd string, data []byte) (res []byte, err error) {
	var response *http.Response
	var encoded string

	buffer := bytes.NewBuffer(data)
	if request, err := http.NewRequest("POST", url, buffer); err == nil {
		encoder := base64.NewEncoding(base64Table)
		encoded = encoder.EncodeToString([]byte(username + ":" + userpwd))
		request.Header.Add("Authorization", "Basic "+encoded)
		request.Header.Add("Content-Type", "application/json;charset=utf-8")

		client := &http.Client{}
		if response, err = client.Do(request); err == nil {
			var body []byte
			if body, err = ioutil.ReadAll(response.Body); err == nil {
				if response.StatusCode != http.StatusOK {
					err = errors.New(string(body))
				} else {
					res = body
				}
			}
			response.Body.Close()
		}
	}
	return res, err
}

func Get(url string) (data []byte, err error) {
	var response *http.Response
	if response, err = http.Get(url); err == nil {
		if data, err = ioutil.ReadAll(response.Body); err == nil {
			if response.StatusCode != http.StatusOK {
				data = make([]byte, 0)
				err = errors.New(string(data))
			}
		}
	}
	return data, err
}
