/*
Copyright 2021 Citrix Systems, Inc, All rights reserved.

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

package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"strings"

	"github.com/hashicorp/go-hclog"
)

// Idempotent flag can't be added for these resources
var idempotentInvalidResources = []string{
	"login",
	"logout",
	"reboot",
	"shutdown",
	"ping",
	"ping6",
	"traceroute",
	"traceroute6",
	"install",
	"appfwjsoncontenttype",
	"appfwxmlcontenttype",
	"dnsnsrec",
	"transformaction",
	"route",
	"lbroute6",
	"vpnnexthopserver",
	"sslcertificatechain",
	"appfwurlencodedformcontenttype",
	"bridgetable",
	"appfwmultipartformcontenttype",
	"snmptrap",
	"dnsmxrec",
	"dnstxtrec",
	"locationfile",
	"locationfile6",
	"cacheforwardproxy",
	"systemuser",
	"dnsaddrec",
	"dnsnameserver",
}

// HTTP Headers to be masked and not shown in logs
var headersToBeMasked = []string{"X-NITRO-USER", "X-NITRO-PASS", "Set-Cookie"}

const (
	nsErrSessionExpired = 444
	nsErrAuthTimeout    = 1027
)

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

type responseHandlerFunc func(resp *http.Response, logger hclog.Logger) ([]byte, error)

func createResponseHandler(resp *http.Response, logger hclog.Logger) ([]byte, error) {
	switch resp.Status {
	case "201 Created", "200 OK":
		body, _ := io.ReadAll(resp.Body)
		return body, nil
	case "409 Conflict":
		body, _ := io.ReadAll(resp.Body)
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")

	case "207 Multi Status":
		//This happens in case of Bulk operations, which we do not support yet
		body, _ := io.ReadAll(resp.Body)
		return body, nil
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"404 Not Found", "405 Method Not Allowed", "406 Not Acceptable",
		"503 Service Unavailable", "599 Netscaler specific error":
		body, _ := io.ReadAll(resp.Body)
		logger.Info("error = ", "body", string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := io.ReadAll(resp.Body)
		return body, err

	}
}

func deleteResponseHandler(resp *http.Response, logger hclog.Logger) ([]byte, error) {
	switch resp.Status {
	case "200 OK", "404 Not Found":
		body, _ := io.ReadAll(resp.Body)
		return body, nil

	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := io.ReadAll(resp.Body)
		logger.Info("delete: error = ", "body", string(body))
		return body, errors.New("[INFO] delete failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := io.ReadAll(resp.Body)
		return body, err

	}
}

func readResponseHandler(resp *http.Response, logger hclog.Logger) ([]byte, error) {
	switch resp.Status {
	case "200 OK":
		body, _ := io.ReadAll(resp.Body)
		return body, nil
	case "404 Not Found":
		body, _ := io.ReadAll(resp.Body)
		logger.Debug("readResponseHandler: 404 not found")
		return body, errors.New("read: 404 not found: ")
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"405 Method Not Allowed", "406 Not Acceptable",
		"409 Conflict", "503 Service Unavailable", "599 Netscaler specific error":
		body, _ := io.ReadAll(resp.Body)
		logger.Info("read: error = ", "body", string(body))
		return body, errors.New("[INFO] failed read: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := io.ReadAll(resp.Body)
		logger.Info("read error = ", "body", string(body))
		return body, err

	}
}

func (c *NitroClient) createHTTPRequest(method string, urlstr string, buff *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, urlstr, buff)
	if err != nil {
		return nil, err
	}
	// Get resourceType from url
	u, err := neturl.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	splitStrings := strings.Split(u.Path, "/")
	resourceType := splitStrings[len(splitStrings)-1]

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.proxiedNs == "" {
		if c.IsLoggedIn() {
			req.Header.Set("Set-Cookie", "NITRO_AUTH_TOKEN="+c.getSessionid())
		} else {
			if resourceType != "login" {
				req.Header.Set("X-NITRO-USER", c.username)
				req.Header.Set("X-NITRO-PASS", c.password)
			}
		}
	} else {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("_MPS_API_PROXY_MANAGED_INSTANCE_IP", c.proxiedNs)
	}

	// User defined headers may overwrite previous headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func maskHeaders(headers http.Header) http.Header {
	maskedHeaders := make(http.Header, len(headers))
	for k, v := range headers {
		upperKey := strings.ToUpper(k)
		if contains(headersToBeMasked, upperKey) {
			maskedHeaders[k] = []string{"*********"}
		} else {
			maskedHeaders[k] = v
		}
	}
	return maskedHeaders
}

func (c *NitroClient) doHTTPRequest(method string, urlstr string, bytes *bytes.Buffer, respHandler responseHandlerFunc) ([]byte, error) {
	req, err := c.createHTTPRequest(method, urlstr, bytes)

	maskedHeaders := maskHeaders(req.Header)
	c.logger.Trace("doHTTPRequest HTTP method", "method", method, "url", urlstr, "headers", maskedHeaders)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}
	c.logger.Trace("response Status:", "status", resp.Status)
	body, err := respHandler(resp, c.logger)
	// Clear sessionid in case of session-expiry
	if resp.Status == "401 Unauthorized" {
		var data map[string]interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 == nil {
			errorcode, ok := data["errorcode"]
			if ok {
				errorcode = int(errorcode.(float64))
				if errorcode == nsErrSessionExpired || errorcode == nsErrAuthTimeout {
					c.clearSessionid()
				}
			}
		}
	}
	return body, err
}

func (c *NitroClient) createResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	c.logger.Trace("Creating ", "resourceType", resourceType)

	url := c.url + resourceType

	if !strings.HasSuffix(resourceType, "_binding") && !contains(idempotentInvalidResources, resourceType) {
		url = url + "?idempotent=yes"
	}
	c.logger.Trace("createResource", "url", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) applyResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	c.logger.Trace("Applying", "resourceType", resourceType)

	url := c.url + resourceType + "?action=apply"
	c.logger.Trace("url is ", "url", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) actOnResource(resourceType string, resourceJSON []byte, action string) ([]byte, error) {
	c.logger.Trace("acting on resource", "resourceType", resourceType)

	var url string
	if action == "" {
		url = c.url + fmt.Sprintf("%s", resourceType)
	} else {
		url = c.url + fmt.Sprintf("%s?action=%s", resourceType, action)
	}
	c.logger.Trace("actOnResource ", "url", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) changeResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	c.logger.Trace("changing resource", "resourceType", resourceType)

	resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
	url := c.url + resourceType + "/" + resourceNameEscaped + "?action=update"
	c.logger.Trace("changeResource", "url", url)

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateResource(resourceType string, resourceName string, resourceJSON []byte) ([]byte, error) {
	c.logger.Trace("Updating resource ", "resourceType", resourceType)

	resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
	url := c.url + resourceType + "/" + resourceNameEscaped
	c.logger.Trace("updateResource ", "url", url)

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) updateUnnamedResource(resourceType string, resourceJSON []byte) ([]byte, error) {
	c.logger.Trace("Updating unnamed resource", "resourceType", resourceType)

	url := c.url + resourceType
	c.logger.Trace("updateUnnamedResource", "url", url)

	return c.doHTTPRequest("PUT", url, bytes.NewBuffer(resourceJSON), createResponseHandler)

}

func (c *NitroClient) deleteResource(resourceType string, resourceName string) ([]byte, error) {
	c.logger.Trace("Deleting resource", "resourceType", resourceType)
	var url string
	if resourceName != "" {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceNameEscaped)
	} else {
		url = c.url + fmt.Sprintf("%s", resourceType)
	}
	c.logger.Trace("deleteResource", "url", url)

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) deleteResourceWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	c.logger.Trace("Deleting resource with args", "resourceType", resourceType, "args ", args)
	var url string
	if resourceName != "" {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.url + fmt.Sprintf("%s/%s?args=", resourceType, resourceNameEscaped)
	} else {
		url = c.url + fmt.Sprintf("%s?args=", resourceType)
	}
	url = url + strings.Join(args, ",")
	c.logger.Trace("deleteResourceWithArgs ", "url", url)

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) deleteResourceWithArgsMap(resourceType string, resourceName string, argsMap map[string]string) ([]byte, error) {
	args := make([]string, len(argsMap))
	i := 0
	for key, value := range argsMap {
		args[i] = fmt.Sprintf("%s:%s", key, value)
		i++
	}
	return c.deleteResourceWithArgs(resourceType, resourceName, args)

}

func (c *NitroClient) unbindResource(resourceType string, resourceName string, boundResourceType string, boundResource string, bindingFilterName string) ([]byte, error) {
	c.logger.Trace("Unbinding resource", "resourceType", resourceType, "resourceName", resourceName)
	bindingName := resourceType + "_" + boundResourceType + "_binding"
	resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))

	url := c.url + "/" + bindingName + "/" + resourceNameEscaped + "?args=" + bindingFilterName + ":" + boundResource

	return c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}), deleteResponseHandler)

}

func (c *NitroClient) listBoundResources(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) ([]byte, error) {
	c.logger.Trace("listing bound resources of type ", "resourceType", resourceType, "resourceName", resourceName)
	var url string
	resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
	if boundResourceFilterName == "" {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s", resourceType, boundResourceType, resourceNameEscaped)
	} else {
		url = c.url + fmt.Sprintf("%s_%s_binding/%s?filter=%s:%s", resourceType, boundResourceType, resourceNameEscaped, boundResourceFilterName, boundResourceFilterValue)
	}

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listFilteredResource(resourceType string, filter map[string]string) ([]byte, error) {
	c.logger.Trace("listing filtered resource of type ", "resourceType", resourceType, "filter: ", filter)

	var filter_strings []string
	for key, value := range filter {
		filter_strings = append(filter_strings, fmt.Sprintf("%s:%s", key, value))
	}

	filter_string := strings.Join(filter_strings, ",")

	url := c.url + fmt.Sprintf("%s?filter=%s", resourceType, filter_string)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResource(resourceType string, resourceName string) ([]byte, error) {
	c.logger.Trace("listing resource of type ", "resourceType", resourceType, "name", resourceName)
	url := c.url + resourceType

	if resourceName != "" {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceNameEscaped)
	}
	c.logger.Trace("listResource", "url", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listResourceWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	c.logger.Trace("listing resource with args ", "resourceType", resourceType, "name", resourceName, "args", args)
	var url string

	if resourceName != "" {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.url + fmt.Sprintf("%s/%s", resourceType, resourceNameEscaped)
	} else {
		url = c.url + fmt.Sprintf("%s", resourceType)
	}
	strArgs := strings.Join(args, ",")
	url2 := url + "?args=" + strArgs
	c.logger.Trace("listResourceWithArgs", "url", url)

	data, err := c.doHTTPRequest("GET", url2, bytes.NewBuffer([]byte{}), readResponseHandler)
	if err != nil {
		c.logger.Trace("listResourceWithArgs: error listing with args, trying filter", "error", err)
		url2 = url + "?filter=" + strArgs
		c.logger.Trace("listResourceWithArgs", "url2", url2)
		return c.doHTTPRequest("GET", url2, bytes.NewBuffer([]byte{}), readResponseHandler)
	}
	return data, err

}

func (c *NitroClient) listResourceWithArgsMap(resourceType string, resourceName string, argsMap map[string]string) ([]byte, error) {
	args := make([]string, len(argsMap))
	i := 0
	for key, value := range argsMap {
		args[i] = fmt.Sprintf("%s:%s", key, value)
		i++
	}
	return c.listResourceWithArgs(resourceType, resourceName, args)

}

func (c *NitroClient) enableFeatures(featureJSON []byte) ([]byte, error) {
	c.logger.Trace("Enabling features")
	url := c.url + "nsfeature?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) disableFeatures(featureJSON []byte) ([]byte, error) {
	c.logger.Trace("Disabling features")
	url := c.url + "nsfeature?action=disable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(featureJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledFeatures() ([]byte, error) {
	c.logger.Trace("listing features")
	url := c.url + "nsfeature"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) enableModes(modeJSON []byte) ([]byte, error) {
	c.logger.Trace("Enabling modes")
	url := c.url + "nsmode?action=enable"

	return c.doHTTPRequest("POST", url, bytes.NewBuffer(modeJSON), createResponseHandler)

}

func (c *NitroClient) listEnabledModes() ([]byte, error) {
	c.logger.Trace("listing modes")
	url := c.url + "nsmode"

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) saveConfig(saveJSON []byte) error {
	c.logger.Trace("Saving config")
	url := c.url + "nsconfig?action=save"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(saveJSON), createResponseHandler)
	return err

}

func (c *NitroClient) clearConfig(clearJSON []byte) error {
	c.logger.Trace("Clearing config")
	url := c.url + "nsconfig?action=clear"

	_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(clearJSON), createResponseHandler)
	return err
}

func (c *NitroClient) listStat(resourceType, resourceName string) ([]byte, error) {
	c.logger.Trace("listing stat of type ", "resourceType", resourceType, "name", resourceName)
	url := c.statsURL + resourceType

	if resourceName != "" {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.statsURL + fmt.Sprintf("%s/%s", resourceType, resourceNameEscaped)
	}
	c.logger.Trace("listStat", "url", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)

}

func (c *NitroClient) listStatWithArgs(resourceType string, resourceName string, args []string) ([]byte, error) {
	c.logger.Trace("listing stat ", "resourceType", resourceType, "name", resourceName, "args", args)
	var url string

	if len(resourceName) > 0 {
		resourceNameEscaped := neturl.PathEscape(neturl.PathEscape(resourceName))
		url = c.statsURL + fmt.Sprintf("%s/%s", resourceType, resourceNameEscaped)
	} else {
		url = c.statsURL + fmt.Sprintf("%s", resourceType)
	}
	strArgs := strings.Join(args, ",")
	url = url + "?args=" + strArgs
	c.logger.Trace("listStatWithArgs", "url", url)

	return c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)
}
