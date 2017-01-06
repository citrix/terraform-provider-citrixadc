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
		log.Println("[INFO] go-nitro: error = " + string(body))
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
		log.Println("[INFO] go-nitro: delete: error = " + string(body))
		return body, errors.New("[INFO] delete failed: " + resp.Status + " (" + string(body) + ")")
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
	case "404 Not Found":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[DEBUG] go-nitro: read: 404 not found")
		return body, errors.New("go-nitro: read: 404 not found: ")
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: read: error = " + string(body))
		return body, errors.New("[INFO] go-nitro: failed read: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: read error = " + string(body))
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
	log.Println("[DEBUG] go-nitro: response Status:", resp.Status)
	return respHandler(resp)
}

func (c *NitroClient) createResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Creating resource of type ", resourceType)

	url := c.url + resourceType

	if !strings.HasSuffix(resourceType, "_binding") {
		url = url + "?idempotent=yes"
	}
	log.Println("[TRACE] go-nitro: url is ", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) changeResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: changing resource of type ", resourceType)

	url := c.url + resourceType + "/" + resourceName + "?action=update"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Updating resource of type ", resourceType)

	url := c.url + resourceType + "/" + resourceName

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) deleteResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Deleting resource of type ", resourceType)
	url := c.url + resourceType + "/" + resourceName

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) deleteResourceWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Deleting resource of type ", resourceType)
	url := c.url + resourceType + "/" + resourceName + "?args="
	for _, arg := range args {
		url = url + arg
	}

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) unbindResource(resourceType string, resourceName string, boundResourceType string, boundResource string, bindingFilterName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: Unbinding resource of type ", resourceType, " name", resourceName)
	bindingName := resourceType + "_" + boundResourceType + "_binding"

	url := c.url + "/" + bindingName + "/" + resourceName + "?args=" + bindingFilterName + ":" + boundResource

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) listBoundResources(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing bound resources of type ", resourceType, ": ", resourceName)
	var url string
	if boundResourceFilterName == "" {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s", resourceType, boundResourceType, resourceName)
	} else {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s?filter=%s:%s", resourceType, boundResourceType, resourceName, boundResourceFilterName, boundResourceFilterValue)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResource(resourceType string, resourceName string) ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing resource of type ", resourceType, ", name: ", resourceName)
	url := c.url + resourceType

	if resourceName != "" {
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceName)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) enableFeatures(featureJSON []byte) ([]byte, error) {
	log.Println("[DEBUG] go-nitro Enabling features")
	url := c.url + "nsfeature?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledFeatures() ([]byte, error) {
	log.Println("[DEBUG] go-nitro: listing features")
	url := c.url + "nsfeature"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) saveConfig(saveJSON []byte) error {
	log.Println("[DEBUG] go-nitro: Saving config")
	url := c.url + "nsconfig?action=save"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(saveJSON), createResponseHandler)
	return err

}

func (c *NitroClient) clearConfig(clearJSON []byte) error {
	log.Println("[DEBUG] go-nitro: Clearing config")
	url := c.url + "nsconfig?action=clear"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(clearJSON), createResponseHandler)
	return err
}
