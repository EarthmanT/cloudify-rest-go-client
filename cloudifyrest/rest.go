/*
Copyright (c) 2017 GigaSpaces Technologies Ltd. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudifyrest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const JsonContentType = "application/json"
const DataContentType = "application/octet-stream"

func (r *CloudifyRestClient) GetRequest(url, method string, body io.Reader) *http.Request {
	log.Printf("Use: %v:%v@%v#%s\n", r.User, r.Password, r.RestURL+url, r.Tenant)

	var auth_string string
	auth_string = r.User + ":" + r.Password
	req, err := http.NewRequest(method, r.RestURL+url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth_string)))
	if len(r.Tenant) > 0 {
		req.Header.Add("Tenant", r.Tenant)
	}

	return req
}

func (r *CloudifyRestClient) Get(url, acceptedContentType string) []byte {
	req := r.GetRequest(url, "GET", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if len(contentType) < len(acceptedContentType) || contentType[:len(acceptedContentType)] != acceptedContentType {
		log.Fatal(fmt.Sprintf("Wrong content type: %+v", contentType))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if acceptedContentType == JsonContentType {
		log.Printf("Response %s\n", string(body))
	} else {
		log.Printf("Binary response length: %d\n", len(body))
	}
	return body
}

func (r *CloudifyRestClient) Delete(url string) []byte {
	req := r.GetRequest(url, "DELETE", nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if contentType[:len(JsonContentType)] != JsonContentType {
		log.Fatal(fmt.Sprintf("Wrong content type: %+v", contentType))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response %s\n", string(body))
	return body
}

func (r *CloudifyRestClient) Post(url string, data []byte) []byte {
	req := r.GetRequest(url, "POST", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", JsonContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if contentType[:len(JsonContentType)] != JsonContentType {
		log.Fatal(fmt.Sprintf("Wrong content type: %+v", contentType))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response %s\n", string(body))
	return body
}

func (r *CloudifyRestClient) Put(url, providedContentType string, data []byte) []byte {
	req := r.GetRequest(url, "PUT", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", providedContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if contentType[:len(JsonContentType)] != JsonContentType {
		log.Fatal(fmt.Sprintf("Wrong content type: %+v", contentType))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response %s\n", string(body))
	return body
}
