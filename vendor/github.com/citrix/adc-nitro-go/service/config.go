/*
Copyright 2021 Citrix Systems, Inc. All rights reserved.

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
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/go-hclog"
)

// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225#28596225
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

type FindParams struct {
	ArgsMap                  map[string]string
	FilterMap                map[string]string
	AttrsMap                 map[string]string
	ResourceType             string
	ResourceName             string
	ResourceMissingErrorCode int
}

type login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
}

type logout struct {
}

func (c *NitroClient) SetLogLevel(level string) {
	c.logger.SetLevel(hclog.LevelFromString(level))
}

func (c *NitroClient) updateSessionid(sessionid string) {
	c.sessionidMux.Lock()
	c.sessionid = sessionid
	c.sessionidMux.Unlock()
}

func (c *NitroClient) clearSessionid() {
	c.updateSessionid("")
}

func (c *NitroClient) getSessionid() string {
	c.sessionidMux.RLock()
	defer c.sessionidMux.RUnlock()
	return c.sessionid
}

func constructQueryString(findParams *FindParams) string {
	// Query string parameters
	var queryBuilder strings.Builder
	concatQueryString(&queryBuilder, constructQueryMapString("args=", findParams.ArgsMap))
	concatQueryString(&queryBuilder, constructQueryMapString("filter=", findParams.FilterMap))
	concatQueryString(&queryBuilder, constructQueryMapString("attrs=", findParams.AttrsMap))

	return queryBuilder.String()
}

func constructQueryMapString(prefix string, queryMap map[string]string) string {
	// Early retrun for empty map
	if len(queryMap) == 0 {
		return ""
	}
	var queryBuilder strings.Builder
	queryBuilder.WriteString(prefix)
	keys := make([]string, 0, len(queryMap))
	for k, _ := range queryMap {
		keys = append(keys, k)
	}
	// Make it deterministic for testing's sake
	sort.Strings(keys)
	lastIndex := len(keys) - 1
	//c.logger.Printf("lastindex %v", lastIndex)

	for i, k := range keys {
		v := queryMap[k]
		//c.logger.Printf("i %v", i)
		if i < lastIndex {
			queryBuilder.WriteString(fmt.Sprintf("%s:%s,", k, v))
		} else {
			queryBuilder.WriteString(fmt.Sprintf("%s:%s", k, v))
		}
	}

	return queryBuilder.String()
}

func concatQueryString(b *strings.Builder, queryString string) {
	// Early return for empty query string
	if len(queryString) == 0 {
		return
	}

	// Fallthrough to processing

	if b.Len() == 0 {
		b.WriteString("?")
	} else {
		b.WriteString("&")
	}
	b.WriteString(queryString)
}

func constructUrlPathString(findParams *FindParams) string {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(findParams.ResourceType)
	if findParams.ResourceName != "" {
		if urlBuilder.Len() > 0 {
			urlBuilder.WriteString("/")
		}
		urlBuilder.WriteString(findParams.ResourceName)
	}
	return urlBuilder.String()
}

// IsLoggedIn tells if user is already logged in
func (c *NitroClient) IsLoggedIn() bool {
	if len(c.getSessionid()) > 0 {
		return true
	}
	return false
}

// Login to netscaler and store the session
func (c *NitroClient) Login() error {
	// Check if login is already done
	if c.IsLoggedIn() {
		return nil
	}
	loginObj := login{
		Username: c.username,
		Password: c.password,
		Timeout:  c.timeout,
	}
	body, err := c.AddResourceReturnBody(Login.Type(), "login", loginObj)
	if err != nil {
		return err
	}
	// Read sessionid from response body
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err == nil {
		c.updateSessionid(data["sessionid"].(string))
	}
	return err
}

// Logout from netscaler and clear the session
func (c *NitroClient) Logout() error {
	logoutObj := logout{}
	_, err := c.AddResource(Logout.Type(), "logout", logoutObj)
	c.clearSessionid()
	return err
}

//AddResourceReturnBody adds a resource of supplied type and name and returns http response body
func (c *NitroClient) AddResourceReturnBody(resourceType string, name string, resourceStruct interface{}) ([]byte, error) {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	var doNotPrintResources = []string{"systemfile", "login", "logout"}
	if !contains(doNotPrintResources, resourceType) {
		c.logger.Trace("AddResourceReturnBody", "resourceJSON", string(resourceJSON))
	}
	body, err := c.createResource(resourceType, resourceJSON)
	if err != nil {
		return body, fmt.Errorf("[ERROR] nitro-go: Failed to create resource of type %s, name=%s, err=%s", resourceType, name, err)
	}
	return body, nil
}

//AddResource adds a resource of supplied type and name
func (c *NitroClient) AddResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	var doNotPrintResources = []string{"systemfile", "login", "logout"}
	if !contains(doNotPrintResources, resourceType) {
		c.logger.Trace("AddResource", "resourceJSON", string(resourceJSON))
	}
	body, err := c.createResource(resourceType, resourceJSON)
	if err != nil {
		return "", fmt.Errorf("[ERROR] nitro-go: Failed to create resource of type %s, name=%s, err=%s", resourceType, name, err)
	}
	_ = body

	return name, nil
}

// ApplyResource applies the configured settings to for the supplied type
func (c *NitroClient) ApplyResource(resourceType string, resourceStruct interface{}) error {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	c.logger.Trace("ApplyResource ", "resourceJSON", string(resourceJSON))

	body, err := c.applyResource(resourceType, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to apply resource of type %s,  err=%s", resourceType, err)
	}
	_ = body

	return nil
}

// ActOnResource applies the configured settings using the action (enable, disable, unset, apply, rename)
func (c *NitroClient) ActOnResource(resourceType string, resourceStruct interface{}, action string) error {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	c.logger.Trace("Resourcejson is ", "resourceJSON", string(resourceJSON))

	_, err = c.actOnResource(resourceType, resourceJSON, action)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to apply action on resource of type %s,  action=%s err=%s", resourceType, action, err)
	}

	return nil
}

//UpdateResource updates a resource of supplied type and name
func (c *NitroClient) UpdateResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	if c.ResourceExists(resourceType, name) == true {
		nsResource := make(map[string]interface{})
		nsResource[resourceType] = resourceStruct
		resourceJSON, err := JSONMarshal(nsResource)

		c.logger.Debug("UpdateResource", "resourceJSON", string(resourceJSON))

		body, err := c.updateResource(resourceType, name, resourceJSON)
		if err != nil {
			return "", fmt.Errorf("[ERROR] nitro-go: Failed to update resource of type %s, name=%s err=%s", resourceType, name, err)
		}
		_ = body
	}

	return name, nil
}

//UpdateUnnamedResource updates a resource of supplied type , which doesn't have a name. E.g., rnat rule
func (c *NitroClient) UpdateUnnamedResource(resourceType string, resourceStruct interface{}) error {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct
	resourceJSON, err := JSONMarshal(nsResource)

	c.logger.Trace("UpdateResource", "resourcejson", string(resourceJSON))

	body, err := c.updateUnnamedResource(resourceType, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to update resource of type %s,  err=%s", resourceType, err)
	}
	_ = body

	return nil
}

//ChangeResource updates a resource of supplied type and name (used for SSL objects)
func (c *NitroClient) ChangeResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	if c.ResourceExists(resourceType, name) == true {
		nsResource := make(map[string]interface{})
		nsResource[resourceType] = resourceStruct
		resourceJSON, err := json.Marshal(nsResource)

		c.logger.Trace("ChangeResource:", "resourcejson", string(resourceJSON))

		body, err := c.changeResource(resourceType, name, resourceJSON)
		if err != nil {
			return "", fmt.Errorf("[ERROR] nitro-go: Failed to change resource of type %s, name=%s err=%s", resourceType, name, err)
		}
		_ = body
	}

	return name, nil
}

//DeleteResource deletes a resource of supplied type and name
func (c *NitroClient) DeleteResource(resourceType string, resourceName string) error {

	_, err := c.listResource(resourceType, resourceName)
	if err == nil { // resource exists
		c.logger.Trace("DeleteResource Found resource ", "resourceType", resourceType, "resourceName", resourceName)
		_, err = c.deleteResource(resourceType, resourceName)
		if err != nil {
			c.logger.Warn("Failed to delete resource", "resourceType", resourceType, "resourceName", resourceName, "error", err)
			return err
		}
	} else {
		c.logger.Info("DeleteResource: Resource already deleted ", "resourceType", resourceType, "resourceName", resourceName)
	}
	return nil
}

//DeleteResourceWithArgs deletes a resource of supplied type and name. Args are supplied as an array of strings
//Each array entry is formatted as "key:value"
func (c *NitroClient) DeleteResourceWithArgs(resourceType string, resourceName string, args []string) error {
	var err error
	if resourceType == "snmptrap_snmpuser_binding" {
		//Remove unwanted argument (username) for listing but keep it for delete operation
		argsWithoutUsername := make([]string, 0)
		for i, v := range args {
			if !strings.Contains(v, "username:") {
				argsWithoutUsername = append(argsWithoutUsername, args[i])
			}
		}
		_, err = c.listResourceWithArgs(resourceType, resourceName, argsWithoutUsername)
	} else if resourceType == "cacheforwardproxy" {
		//Cacheforwardproxy supports only GET(all) Request
		_, err = c.listResource(resourceType, "")
	} else if resourceType == "analyticsglobal_analyticsprofile_binding" {
		// analyticsglobal_analyticsprofile_binding has different GET implementation
		_, err = c.listResource("analyticsglobal", "")
	} else {
		_, err = c.listResourceWithArgs(resourceType, resourceName, args)
	}
	if err == nil { // resource exists
		c.logger.Trace("DeleteResource Found resource ", "resourceType", resourceType, "resourceName", resourceName)
		_, err = c.deleteResourceWithArgs(resourceType, resourceName, args)
		if err != nil {
			c.logger.Warn("Failed to delete resource", "resourceType", resourceType, "resourceName", resourceName, "error", err)
			return err
		}
	} else {
		c.logger.Info("DeleteResource: Resource already deleted ", "resourceType", resourceType, "resourceName", resourceName)
	}
	return nil
}

//DeleteResourceWithArgsMap deletes a resource of supplied type and name. Args are supplied as map of key value
func (c *NitroClient) DeleteResourceWithArgsMap(resourceType string, resourceName string, args map[string]string) error {

	_, err := c.listResourceWithArgsMap(resourceType, resourceName, args)
	if err == nil { // resource exists
		c.logger.Trace("DeleteResource Found resource ", "resourceType", resourceType, "resourceName", resourceName)

		_, err = c.deleteResourceWithArgsMap(resourceType, resourceName, args)
		if err != nil {
			c.logger.Warn("Failed to delete resource", "resourceType", resourceType, "resourceName", resourceName, "error", err)

			return err
		}
	} else {
		c.logger.Info(" Resource already deleted ", "resourceType", resourceType, "resourceName", resourceName)
	}
	return nil
}

//BindResource binds the 'bindingResourceName' to the 'bindToResourceName'.
func (c *NitroClient) BindResource(bindToResourceType string, bindToResourceName string, bindingResourceType string, bindingResourceName string, bindingStruct interface{}) error {
	if !c.ResourceExists(bindToResourceType, bindToResourceName) {
		return fmt.Errorf("[ERROR] nitro-go: BindTo Resource %s of type %s does not exist", bindToResourceType, bindToResourceName)
	}

	if !c.ResourceExists(bindingResourceType, bindingResourceName) {
		return fmt.Errorf("[ERROR] nitro-go: Binding Resource %s of type %s does not exist", bindingResourceType, bindingResourceName)
	}
	bindingName := bindToResourceType + "_" + bindingResourceType + "_binding"
	nsBinding := make(map[string]interface{})
	nsBinding[bindingName] = bindingStruct

	resourceJSON, _ := JSONMarshal(nsBinding)

	body, err := c.createResource(bindingName, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to bind resource %s to resource %s, err=%s", bindToResourceName, bindingResourceName, err)
	}
	_ = body
	return nil
}

//UnbindResource unbinds 'boundResourceName' from 'boundToResource'
func (c *NitroClient) UnbindResource(boundToResourceType string, boundToResourceName string, boundResourceType string, boundResourceName string, bindingFilterName string) error {

	if !c.ResourceExists(boundToResourceType, boundToResourceName) {
		c.logger.Info(" Unbind: BoundTo Resource  does not exist", "boundToResourceType", boundToResourceType, "boundToResourceName", boundToResourceName)
		return nil
	}

	if !c.ResourceExists(boundResourceType, boundResourceName) {
		c.logger.Info(" Unbind: Bound Resource  does not exist", "boundResourceType", boundResourceType, "boundResourceName", boundResourceName)
		return nil
	}

	_, err := c.unbindResource(boundToResourceType, boundToResourceName, boundResourceType, boundResourceName, bindingFilterName)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to unbind  %s:%s from %s:%s, err=%s", boundResourceType, boundResourceName, boundToResourceType, boundToResourceName, err)
	}

	return nil
}

//ResourceExists returns true if supplied resource name and type exists
func (c *NitroClient) ResourceExists(resourceType string, resourceName string) bool {
	_, err := c.listResource(resourceType, resourceName)
	if err != nil {
		c.logger.Debug("ResourceExists: No resource found", "resourceType", resourceType, "resourceName", resourceName)
		return false
	}
	c.logger.Debug("ResourceExists: resource is already present", "resourceType", resourceType, "resourceName", resourceName)
	return true
}

//FindResourceArray returns the config of the supplied resource name and type if it exists. Use when the resource to be returned is an array
func (c *NitroClient) FindResourceArray(resourceType string, resourceName string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, resourceName)
	if err != nil {
		c.logger.Warn("FindResourceArray: No resources found", "resourceType", resourceType, "resourceName", resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindResourceArray: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindResourceArray: Failed to unmarshal Netscaler Response!", "resourceType", resourceType, "resourceName", resourceName)
		return nil, fmt.Errorf("[ERROR] nitro-go: FindResourceArray: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		c.logger.Warn("FindResourceArray No resource found", "resourceType", resourceType, "resourceName", resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindResourceArray: No resource %s of type %s found", resourceName, resourceType)
	}
	resources := data[resourceType].([]interface{})
	ret := make([]map[string]interface{}, len(resources), len(resources))
	for i, v := range resources {
		ret[i] = v.(map[string]interface{})
	}
	return ret, nil
}

//FindFilteredResourceArray returns the config of the supplied resource type, filtered with given filter
func (c *NitroClient) FindFilteredResourceArray(resourceType string, filter map[string]string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listFilteredResource(resourceType, filter)
	if err != nil {
		c.logger.Warn("FindFilteredResourceArray: No resource found matching filter", "resourceType", resourceType, "filter", filter)
		return nil, fmt.Errorf("[INFO] nitro-go: FindFilteredResourceArray: No resource of type %s found matching filter %s", resourceType, filter)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindFilteredResourceArray: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindFilteredResourceArray: Failed to unmarshal Netscaler Response:resource of type %s matching filter %s", resourceType, filter)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		c.logger.Warn("FindFilteredResourceArray No resource matching filter found", "resourceType", resourceType, "filter", filter)
		return nil, fmt.Errorf("[INFO] nitro-go: FindFilteredResourceArray: No resource type %s matching filter %s found", filter, resourceType)
	}
	resources := data[resourceType].([]interface{})
	ret := make([]map[string]interface{}, len(resources), len(resources))
	for i, v := range resources {
		ret[i] = v.(map[string]interface{})
	}
	return ret, nil
}

//FindResource returns the config of the supplied resource name and type if it exists
func (c *NitroClient) FindResource(resourceType string, resourceName string) (map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, resourceName)
	if err != nil {
		c.logger.Warn("FindResource: No resource found", "resourceType", resourceType, "resourceName", resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindResource: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindResource: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindResource: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrc, ok := data[resourceType]
	if !ok || rsrc == nil {
		c.logger.Warn("FindResource No resource found", "resourceType", resourceType, "resourceName", resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindResource: No resource %s of type %s found", resourceName, resourceType)
	}

	switch result := data[resourceType].(type) {
	case map[string]interface{}:
		return result, nil
	case []interface{}:
		return result[0].(map[string]interface{}), nil
	default:
		c.logger.Warn("FindResource Unable to determine type of response")
		return nil, fmt.Errorf("[INFO] nitro-go: FindResource: Unable to determine type of response")
	}
}

// This is meant to be a generic and extendable function to implement all the possible GET methods NITRO will allow.
// Extensibility comes from the use of the FindParams
// Adding fields and handling them inside this function should suffice for any future need.
// When no error is returned the result will always be an array
// An empty array means the resource does not exist
// A single element array means there was only one result
// Multiple elements array means NITRO returned multiple results
// It is left to the user to determine the sanity of the return value
func (c *NitroClient) FindResourceArrayWithParams(findParams FindParams) ([]map[string]interface{}, error) {

	// Construct the url
	var urlBuilder strings.Builder

	// Prefix from nitro client
	urlBuilder.WriteString(c.url)

	// Path from params
	pathStr := constructUrlPathString(&findParams)
	urlBuilder.WriteString(pathStr)

	// Query string from params
	queryStr := constructQueryString(&findParams)
	urlBuilder.WriteString(queryStr)

	url := urlBuilder.String()

	c.logger.Trace("FindResourceArrayWithParams: url is", "url", url)
	result, httpErr := c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}), readResponseHandler)
	c.logger.Trace("FindResourceArrayWithParams: HTTP GET result", "result", string(result), "error", httpErr)

	// Ignore 404.
	// We need to parse the NITRO errorcode value to determine if this is an actual error
	if httpErr != nil {
		if !strings.Contains(httpErr.Error(), "404") {
			return nil, httpErr
		} else {
			c.logger.Debug("FindResourceArrayWithParams: Ignoring 404 http status", "error", httpErr.Error())
		}
	}

	var jsonData interface{}

	err := json.Unmarshal(result, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] nitro-go: FindResourceArrayWithParams: JSON unmarshal error %v", err.Error())
	}

	nitroData, ok := jsonData.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("[ERROR] nitro-go: FindResourceArrayWithParams: Type assertion map[string]interface{} does not hold. Actual type %T", jsonData)
	}

	// Get error code from nitro resource
	errorcode, errok := nitroData["errorcode"]
	if !errok {
		return nil, fmt.Errorf("[ERROR] nitro-go: FindResourceArrayWithParams: there is no error code in nitro response")
	}

	errorcode = int(errorcode.(float64))

	emptyRetval := make([]map[string]interface{}, 0)

	// Resource missing errorcode returned
	missingErrcode := findParams.ResourceMissingErrorCode
	if missingErrcode != 0 && missingErrcode == errorcode {
		c.logger.Debug("FindResourceArrayWithParams: resource missing error code returned", "errorcode", errorcode)
		return emptyRetval, nil
	}
	// Fallthrough

	if errorcode != 0 {
		return nil, fmt.Errorf("[ERROR] FindResourceArrayWithParams: non zero errorcode %d", errorcode)
	}
	// Fallthrough

	// Check if resource type key exists
	resourceData, ok := nitroData[findParams.ResourceType]
	if !ok {
		// Since errorcode is 0 we persume this is the expected behavior for a missing resource
		c.logger.Debug("FindResourceArrayWithParams: resource key missing", "resourceType", findParams.ResourceType)
		return emptyRetval, nil
	}
	// Falthrough

	c.logger.Trace("FindResourceArrayWithParams: retrieved NITRO object", "resourceData", resourceData)

	// resource data assertion
	resourceArray, arrayOk := resourceData.([]interface{})
	resourceMap, mapOk := resourceData.(map[string]interface{})

	if arrayOk {
		retVal := make([]map[string]interface{}, 0, len(resourceArray))
		for _, v := range resourceArray {
			val := v.(map[string]interface{})
			retVal = append(retVal, val)
		}
		return retVal, nil
	}

	if mapOk {
		retVal := make([]map[string]interface{}, 0, 1)
		retVal = append(retVal, resourceMap)
		return retVal, nil

	}
	// Fallthrough to error condition
	return nil, fmt.Errorf("[ERROR] nitro-go: FindResourceArrayWithParams: Cannot handle returned NITRO resource data type %T", resourceData)

}

//FindAllResources finds all config objects of the supplied resource type and returns them in an array
func (c *NitroClient) FindAllResources(resourceType string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, "")
	if err != nil {
		c.logger.Trace(" FindAllResources: No objects found", "resourceType", resourceType)
		return make([]map[string]interface{}, 0, 0), nil
	}
	if err = json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindAllResources: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindAllResources: Failed to unmarshal Netscaler Response: of type %s", resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		c.logger.Trace(" FindAllResources: resource not found", "resourceType", resourceType)
		return make([]map[string]interface{}, 0, 0), nil
	}
	resources := data[resourceType].([]interface{})

	ret := make([]map[string]interface{}, len(resources), len(resources))
	for i, v := range resources {
		ret[i] = v.(map[string]interface{})
	}

	return ret, nil
}

//ResourceBindingExists returns true if the supplied binding exists
func (c *NitroClient) ResourceBindingExists(resourceType string, resourceName string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) bool {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, boundResourceFilterName, boundResourceFilterValue)
	if err != nil {
		c.logger.Trace("ResourceBindingExists: No bound resource to found", "resourceType", resourceType, "name", resourceName, "boundResourceType", boundResourceType, "boundResourceFilterValue", boundResourceFilterValue)
		return false
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		c.logger.Error("ResourceBindingExists: Failed to unmarshal Netscaler Response!")
		return false
	}
	binding := fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)
	rsrc, ok := data[binding]
	if !ok || rsrc == nil {
		return false
	}

	c.logger.Trace("ResourceBindingExists: of type   is bound to  type and name ", "resourceType", resourceType, "resourceName", resourceName, "boundResourceType", boundResourceType, "boundResourceFilterValue", boundResourceFilterValue)
	return true
}

//FindBoundResource finds a bound resource if it exists
func (c *NitroClient) FindBoundResource(resourceType string, resourceName string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) (map[string]interface{}, error) {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, boundResourceFilterName, boundResourceFilterValue)
	if err != nil {
		c.logger.Info(" FindBoundResource: No binding found", "resourceType", resourceType, "resourceName", "resourceName", resourceName, "boundResourceType", boundResourceType, "boundResourceFilterValue", boundResourceFilterValue)
		return nil, fmt.Errorf("[INFO] nitro-go: No %s %s to %s %s binding found, err=%s", resourceType, resourceName, boundResourceType, boundResourceFilterValue, err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindBoundResource: Failed to unmarshal Netscaler Response!", "error", err)
		return nil, fmt.Errorf("[ERROR] nitro-go: FindBoundResource: Failed to unmarshal Netscaler Response!, err=%s", err)
	}
	bindingType := fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)
	rsrc, ok := data[bindingType]
	if !ok || rsrc == nil {
		return nil, fmt.Errorf("[WARN] nitro-go: FindBoundResource: No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
	}

	resource := data[bindingType].([]interface{})[0] //only one resource obviously
	return resource.(map[string]interface{}), nil

}

//FindAllBoundResources returns an array of bound config objects of the type specified that are bound to the resource specified
func (c *NitroClient) FindAllBoundResources(resourceType string, resourceName string, boundResourceType string) ([]map[string]interface{}, error) {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, "", "")
	if err != nil {
		c.logger.Info("FindAllBoundResources: No binding found", "resourceType", resourceType, "resourceName", "resourceName", resourceName, "boundResourceType", boundResourceType)
		return nil, fmt.Errorf("[ERROR] nitro-go: No %s %s to %s binding found, err=%s", resourceType, resourceName, boundResourceType, err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		c.logger.Error("FindAllBoundResources: Failed to unmarshal Netscaler Response!", "error", err)
		return nil, fmt.Errorf("[ERROR] nitro-go: FindAllBoundResources: Failed to unmarshal Netscaler Response!, err=%s", err)
	}
	bindingType := fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)
	binding, ok := data[bindingType]
	if !ok || binding == nil {
		return make([]map[string]interface{}, 0, 0), nil
	}

	resources := binding.([]interface{})
	ret := make([]map[string]interface{}, len(resources), len(resources))
	for i, v := range resources {
		ret[i] = v.(map[string]interface{})
	}
	return ret, nil
}

//EnableFeatures enables the provided list of features. Depending on the licensing of the NetScaler, not all supplied features may actually
//enabled
func (c *NitroClient) EnableFeatures(featureNames []string) error {
	/* construct this:
	{
	        "nsfeature":
		{
		    "feature": [ "LB", ]
		}
	}
	*/
	featureStruct := make(map[string]map[string][]string)
	featureStruct["nsfeature"] = make(map[string][]string)
	featureStruct["nsfeature"]["feature"] = featureNames

	featureJSON, err := JSONMarshal(featureStruct)
	if err != nil {
		c.logger.Error("EnableFeatures: Failed to marshal features to JSON", "error", err)
		return fmt.Errorf("[ERROR] nitro-go: Failed to marshal features to JSON")
	}

	_, err = c.enableFeatures(featureJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to enable feature %v", err)
	}
	return nil
}

func (c *NitroClient) DisableFeatures(featureNames []string) error {
	/* construct this:
	{
	        "nsfeature":
		{
		    "feature": [ "LB", ]
		}
	}
	*/
	featureStruct := make(map[string]map[string][]string)
	featureStruct["nsfeature"] = make(map[string][]string)
	featureStruct["nsfeature"]["feature"] = featureNames

	featureJSON, err := JSONMarshal(featureStruct)
	if err != nil {
		c.logger.Error("DisableFeatures: Failed to marshal features to JSON", "error", err)
		return fmt.Errorf("[ERROR] nitro-go: Failed to marshal features to JSON")
	}

	_, err = c.disableFeatures(featureJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to disable feature %v", err)
	}
	return nil
}

//ListEnabledFeatures returns a string array of the list of features enabled on the NetScaler appliance
func (c *NitroClient) ListEnabledFeatures() ([]string, error) {

	bytes, err := c.listEnabledFeatures()
	if err != nil {
		return []string{}, fmt.Errorf("[ERROR] nitro-go: Failed to list features %v", err)
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes, &data); err != nil {
		c.logger.Error("ListEnabledFeatures: Failed to unmarshal Netscaler Response!", "error", err)
		return []string{}, fmt.Errorf("[ERROR] nitro-go: Failed to unmarshal Netscaler Response to list Features")
	}
	feat, ok := data["nsfeature"]
	if !ok || feat == nil {
		c.logger.Error("ListEnabledFeatures: No features found")
		return []string{}, fmt.Errorf("[ERROR] nitro-go: No features found")
	}

	features := data["nsfeature"].(map[string]interface{})
	// since the returned JSON map mixes boolean and array values, the unmarshal fails to figure out there
	// is an array. So we have to convert it to a string and then parse it.
	// this doesn't work: return features["feature"].([]string), nil
	// convert to string: [LB CS SSL] (note: no commas)

	result := fmt.Sprintf("%v", features["feature"])
	result = strings.TrimPrefix(result, "[")
	result = strings.TrimSuffix(result, "]")
	flist := strings.Split(result, " ")
	log.Println("result: ", result, "flist: ", flist)
	return flist, nil
}

//EnableModes enables the provided list of Citrix ADC modes.
func (c *NitroClient) EnableModes(modeNames []string) error {
	/* construct this:
	{
	        "nsmode":
		{
		    "mode": [ "USNIP", "L3", "LSN", ]
		}
	}
	*/
	modeStruct := make(map[string]map[string][]string)
	modeStruct["nsmode"] = make(map[string][]string)
	modeStruct["nsmode"]["mode"] = modeNames

	modeJSON, err := JSONMarshal(modeStruct)
	if err != nil {
		c.logger.Error("EnableModes: Failed to marshal modes to JSON", "error", err)
		return fmt.Errorf("[ERROR] nitro-go: Failed to marshal modes to JSON")
	}

	_, err = c.enableModes(modeJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to enable mode %v", err)
	}
	return nil
}

//ListEnabledModes returns a string array of the list of modes enabled on the Citrix ADC appliance
func (c *NitroClient) ListEnabledModes() ([]string, error) {

	bytes, err := c.listEnabledModes()
	if err != nil {
		return []string{}, fmt.Errorf("[ERROR] nitro-go: Failed to list modes %v", err)
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes, &data); err != nil {
		c.logger.Error("ListEnabledModes: Failed to unmarshal Netscaler Response!", "error", err)
		return []string{}, fmt.Errorf("[ERROR] nitro-go: Failed to unmarshal Netscaler Response to list Modes")
	}
	mode, ok := data["nsmode"]
	if !ok || mode == nil {
		c.logger.Error("ListEnabledModes: No modes found")
		return []string{}, fmt.Errorf("[ERROR] No modes found")
	}

	modes := data["nsmode"].(map[string]interface{})
	// since the returned JSON map mixes boolean and array values, the unmarshal fails to figure out there
	// is an array. So we have to convert it to a string and then parse it.
	// this doesn't work: return modes["mode"].([]string), nil
	// convert to string: [USNIP L3 LSN] (note: no commas)

	result := fmt.Sprintf("%v", modes["mode"])
	result = strings.TrimPrefix(result, "[")
	result = strings.TrimSuffix(result, "]")
	mlist := strings.Split(result, " ")
	log.Println("result: ", result, "mlist: ", mlist)
	return mlist, nil
}

//SaveConfig persists the config on the NetScaler to the NetScaler's persistent storage. This could take a few seconds
func (c *NitroClient) SaveConfig() error {
	/* construct this:
	{
	        "nsconfig": {}
	}
	*/
	saveStruct := make(map[string]interface{})
	saveStruct["nsconfig"] = make(map[string]interface{})

	saveJSON, err := JSONMarshal(saveStruct)
	if err != nil {
		c.logger.Error("SaveConfig: Failed to marshal save config to JSON", "error", err)
		return fmt.Errorf("[ERROR] nitro-go: Failed to marshal save config to JSON")
	}
	c.logger.Debug("saveJSON", "json", string(saveJSON))

	err = c.saveConfig(saveJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to save config %v", err)
	}
	return nil
}

//ClearConfig deletes the config on the NetScaler
func (c *NitroClient) ClearConfig() error {
	/* construct this:
	{
	    "nsconfig": {"level": "basic"}
	}
	*/
	clearStruct := make(map[string]map[string]string)
	clearStruct["nsconfig"] = make(map[string]string)

	clearStruct["nsconfig"]["level"] = "basic"

	clearJSON, err := JSONMarshal(clearStruct)
	if err != nil {
		c.logger.Error("ClearConfig: Failed to marshal clear config to JSON", "error", err)
		return fmt.Errorf("[ERROR] nitro-go: Failed to marshal clear config to JSON")
	}
	c.logger.Trace("clearJSON is ", "text", string(clearJSON))

	err = c.clearConfig(clearJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] nitro-go: Failed to clear config %v", err)
	}
	return nil
}
