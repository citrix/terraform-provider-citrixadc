/*
Copyright 2016 Citrix Systems, Inc, All rights reserved.

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

package netscaler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type nitroClient struct {
	url      string
	username string
	password string
}

func NewNitroClient(url string, username string, password string) *nitroClient {
	c := new(nitroClient)
	c.url = strings.Trim(url, " /") + "/nitro/v1/config/"
	c.username = username
	c.password = password
	return c
}

func (c *nitroClient) createHttpRequest(method string, url string, buff *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-NITRO-USER", c.username)
	req.Header.Set("X-NITRO-PASS", c.password)
	return req, nil
}

func (c *nitroClient) createResource(resourceType string, resourceJson []byte) ([]byte, error) {
	log.Println("Creating resource of type ", resourceType)

	url := c.url + resourceType

	method := "POST"
	if strings.HasSuffix(resourceType, "_binding") {
		method = "POST" //FIXME: docs say it should be PUT
	}

	req, err := c.createHttpRequest(method, url, bytes.NewBuffer(resourceJson))

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	} else {
		log.Println("response Status:", resp.Status)

		switch resp.Status {
		case "201 Created", "200 OK", "409 Conflict":
			body, _ := ioutil.ReadAll(resp.Body)
			return body, nil

		case "207 Multi Status":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			return body, err
		case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
			"404 Not Found", "405 Method Not Allowed", "406 Not Acceptable",
			"503 Service Unavailable", "599 Netscaler specific error":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
		default:
			body, err := ioutil.ReadAll(resp.Body)
			return body, err

		}
	}
}

func (c *nitroClient) deleteResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("Deleting resource of type ", resourceType)
	url := c.url + resourceType + "/" + resourceName

	req, err := c.createHttpRequest("DELETE", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	} else {

		log.Println("response Status:", resp.Status)

		switch resp.Status {
		case "200 OK", "404 Not Found":
			body, _ := ioutil.ReadAll(resp.Body)
			return body, nil

		case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
			"405 Method Not Allowed", "406 Not Acceptable",
			"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
		default:
			body, err := ioutil.ReadAll(resp.Body)
			return body, err

		}
	}
}

func (c *nitroClient) unbindResource(resourceType string, resourceName string, boundResourceType string, boundResource string) ([]byte, error) {
	log.Println("Unbinding resource of type ", resourceType)

	url := c.url + "/" + resourceName + "?args=" + boundResourceType + ":" + boundResource

	req, err := c.createHttpRequest("DELETE", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	} else {
		log.Println("response Status:", resp.Status)

		switch resp.Status {
		case "200 OK", "404 Not Found":
			body, _ := ioutil.ReadAll(resp.Body)
			return body, nil

		case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
			"405 Method Not Allowed", "406 Not Acceptable",
			"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
		default:
			body, err := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, err

		}
	}
}

func (c *nitroClient) listBoundResources(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) ([]byte, error) {
	log.Println("listing resource of type ", resourceType)
	var url string
	if boundResourceFilterName == "" {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s", resourceType, boundResourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s?filter=%s:%s", resourceType, boundResourceType, resourceName, boundResourceFilterName, boundResourceFilterValue)
	}

	req, err := c.createHttpRequest("GET", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	} else {
		log.Println("response Status:", resp.Status)

		switch resp.Status {
		case "200 OK":
			body, _ := ioutil.ReadAll(resp.Body)
			return body, nil
		case "400 Bad Request", "401 Unauthorized", "403 Forbidden", "404 Not Found",
			"405 Method Not Allowed", "406 Not Acceptable",
			"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
		default:
			body, err := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, err

		}
	}
}

func (c *nitroClient) listResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("listing resource of type ", resourceType)
	url := c.url + resourceType

	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	}

	req, err := c.createHttpRequest("GET", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	} else {
		log.Println("response Status:", resp.Status)

		switch resp.Status {
		case "200 OK":
			body, _ := ioutil.ReadAll(resp.Body)
			return body, nil
		case "400 Bad Request", "401 Unauthorized", "403 Forbidden", "404 Not Found",
			"405 Method Not Allowed", "406 Not Acceptable",
			"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
			//TODO
			body, _ := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
		default:
			body, err := ioutil.ReadAll(resp.Body)
			log.Println("error = " + string(body))
			return body, err

		}
	}
}
