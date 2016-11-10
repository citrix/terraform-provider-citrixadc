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
)

type responseHandlerFunc func(resp *http.Response) ([]byte, error)

func createResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "201 Created", "200 OK":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "409 Conflict":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")

	case "207 Multi Status":
		//This happens in case of Bulk operations, which we do not support yet
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"404 Not Found", "405 Method Not Allowed", "406 Not Acceptable",
		"503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("error = " + string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return body, err

	}
}

func deleteResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "200 OK", "404 Not Found":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil

	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("error = " + string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return body, err

	}
}

func readResponseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "200 OK":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden", "404 Not Found",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("error = " + string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("error = " + string(body))
		return body, err

	}
}

func (c *NitroClient) createHTTPRequest(method string, url string, buff *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-NITRO-USER", c.username)
	req.Header.Set("X-NITRO-PASS", c.password)
	return req, nil
}

func (c *NitroClient) doHTTPRequest(method string, url string, bytes *bytes.Buffer, respHandler responseHandlerFunc) ([]byte, error) {
	req, err := c.createHTTPRequest(method, url, bytes)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}
	log.Println("response Status:", resp.Status)
	return respHandler(resp)
}

func (c *NitroClient) createResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	log.Println("Creating resource of type ", resourceType)

	url := c.url + resourceType

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	log.Println("Updating resource of type ", resourceType)

	url := c.url + resourceType + "/" + resourceName

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) deleteResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("Deleting resource of type ", resourceType)
	url := c.url + resourceType + "/" + resourceName

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) unbindResource(resourceType string, resourceName string, boundResourceType string, boundResource string, bindingFilterName string) ([]byte, error) {
	log.Println("Unbinding resource of type ", resourceType, " name", resourceName)
	bindingName := resourceType + "_" + boundResourceType + "_binding"

	url := c.url + "/" + bindingName + "/" + resourceName + "?args=" + bindingFilterName + ":" + boundResource

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) listBoundResources(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) ([]byte, error) {
	log.Println("listing resource of type ", resourceType)
	var url string
	if boundResourceFilterName == "" {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s", resourceType, boundResourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s?filter=%s:%s", resourceType, boundResourceType, resourceName, boundResourceFilterName, boundResourceFilterValue)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("listing resource of type ", resourceType)
	url := c.url + resourceType

	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) enableFeatures(featureJSON []byte) ([]byte, error) {
	log.Println("Enabling features")
	url := c.url + "nsfeature?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledFeatures() ([]byte, error) {
	log.Println("listing features")
	url := c.url + "nsfeature"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) saveConfig(saveJSON []byte) error {
	log.Println("Saving config")
	url := c.url + "nsconfig?action=save"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(saveJSON), createResponseHandler)
	return err

}

func (c *NitroClient) clearConfig(clearJSON []byte) error {
	log.Println("Clearing config")
	url := c.url + "nsconfig?action=clear"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(clearJSON), createResponseHandler)
	return err
}
