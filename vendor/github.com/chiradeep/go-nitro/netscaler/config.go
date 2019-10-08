/*
Copyright 2016 Citrix Systems, Inc. All rights reserved.

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
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225#28596225
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

//AddResource adds a resource of supplied type and name
func (c *NitroClient) AddResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	if !strings.EqualFold(resourceType, "systemfile") {
		log.Printf("[TRACE] go-nitro: Resourcejson is " + string(resourceJSON))
	}
	body, err := c.createResource(resourceType, resourceJSON)
	if err != nil {
		return "", fmt.Errorf("[ERROR] go-nitro: Failed to create resource of type %s, name=%s, err=%s", resourceType, name, err)
	}
	_ = body

	return name, nil
}

// ApplyResource applies the configured settings to for the supplied type
func (c *NitroClient) ApplyResource(resourceType string, resourceStruct interface{}) error {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	log.Printf("[TRACE] go-nitro: Resourcejson is " + string(resourceJSON))

	body, err := c.applyResource(resourceType, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to apply resource of type %s,  err=%s", resourceType, err)
	}
	_ = body

	return nil
}

// ActOnResource applies the configured settings using the action (enable, disable, unset, apply, rename)
func (c *NitroClient) ActOnResource(resourceType string, resourceStruct interface{}, action string) error {

	nsResource := make(map[string]interface{})
	nsResource[resourceType] = resourceStruct

	resourceJSON, err := JSONMarshal(nsResource)

	log.Printf("[TRACE] go-nitro: Resourcejson is " + string(resourceJSON))

	_, err = c.actOnResource(resourceType, resourceJSON, action)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to apply action on resource of type %s,  action=%s err=%s", resourceType, action, err)
	}

	return nil
}

//UpdateResource updates a resource of supplied type and name
func (c *NitroClient) UpdateResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	if c.ResourceExists(resourceType, name) == true {
		nsResource := make(map[string]interface{})
		nsResource[resourceType] = resourceStruct
		resourceJSON, err := JSONMarshal(nsResource)

		log.Printf("[DEBUG] go-nitro: UpdateResource: Resourcejson is " + string(resourceJSON))

		body, err := c.updateResource(resourceType, name, resourceJSON)
		if err != nil {
			return "", fmt.Errorf("[ERROR] go-nitro: Failed to update resource of type %s, name=%s err=%s", resourceType, name, err)
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

	log.Printf("[DEBUG] go-nitro: UpdateResource: Resourcejson is " + string(resourceJSON))

	body, err := c.updateUnnamedResource(resourceType, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to update resource of type %s,  err=%s", resourceType, err)
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

		log.Printf("[DEBUG] go-nitro: ChangeResource: Resourcejson is " + string(resourceJSON))

		body, err := c.changeResource(resourceType, name, resourceJSON)
		if err != nil {
			return "", fmt.Errorf("[ERROR] go-nitro: Failed to change resource of type %s, name=%s err=%s", resourceType, name, err)
		}
		_ = body
	}

	return name, nil
}

//DeleteResource deletes a resource of supplied type and name
func (c *NitroClient) DeleteResource(resourceType string, resourceName string) error {

	_, err := c.listResource(resourceType, resourceName)
	if err == nil { // resource exists
		log.Printf("[DEBUG] go-nitro: DeleteResource Found resource of type %s: %s", resourceType, resourceName)
		_, err = c.deleteResource(resourceType, resourceName)
		if err != nil {
			log.Printf("[ERROR] go-nitro: Failed to delete resourceType %s: %s, err=%s", resourceType, resourceName, err)
			return err
		}
	} else {
		log.Printf("[INFO] go-nitro: DeleteResource: Resource %s already deleted ", resourceName)
	}
	return nil
}

//DeleteResourceWithArgs deletes a resource of supplied type and name. Args are supplied as an array of strings
//Each array entry is formatted as "key:value"
func (c *NitroClient) DeleteResourceWithArgs(resourceType string, resourceName string, args []string) error {

	_, err := c.listResourceWithArgs(resourceType, resourceName, args)
	if err == nil { // resource exists
		log.Printf("[INFO] go-nitro: DeleteResource found resource of type %s: %s", resourceType, resourceName)
		_, err = c.deleteResourceWithArgs(resourceType, resourceName, args)
		if err != nil {
			log.Printf("[ERROR] go-nitro: Failed to delete resourceType %s: %s, err=%s", resourceType, resourceName, err)
			return err
		}
	} else {
		log.Printf("[INFO] go-nitro: Resource %s already deleted ", resourceName)
	}
	return nil
}

//DeleteResourceWithArgsMap deletes a resource of supplied type and name. Args are supplied as map of key value
func (c *NitroClient) DeleteResourceWithArgsMap(resourceType string, resourceName string, args map[string]string) error {

	_, err := c.listResourceWithArgsMap(resourceType, resourceName, args)
	if err == nil { // resource exists
		log.Printf("[INFO] go-nitro: DeleteResource found resource of type %s: %s", resourceType, resourceName)
		_, err = c.deleteResourceWithArgsMap(resourceType, resourceName, args)
		if err != nil {
			log.Printf("[ERROR] go-nitro: Failed to delete resourceType %s: %s, err=%s", resourceType, resourceName, err)
			return err
		}
	} else {
		log.Printf("[INFO] go-nitro: Resource %s already deleted ", resourceName)
	}
	return nil
}

//BindResource binds the 'bindingResourceName' to the 'bindToResourceName'.
func (c *NitroClient) BindResource(bindToResourceType string, bindToResourceName string, bindingResourceType string, bindingResourceName string, bindingStruct interface{}) error {
	if c.ResourceExists(bindToResourceType, bindToResourceName) == false {
		return fmt.Errorf("[ERROR] go-nitro: BindTo Resource %s of type %s does not exist", bindToResourceType, bindToResourceName)
	}

	if c.ResourceExists(bindingResourceType, bindingResourceName) == false {
		return fmt.Errorf("[ERROR] go-nitro: Binding Resource %s of type %s does not exist", bindingResourceType, bindingResourceName)
	}
	bindingName := bindToResourceType + "_" + bindingResourceType + "_binding"
	nsBinding := make(map[string]interface{})
	nsBinding[bindingName] = bindingStruct

	resourceJSON, err := JSONMarshal(nsBinding)

	body, err := c.createResource(bindingName, resourceJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to bind resource %s to resource %s, err=%s", bindToResourceName, bindingResourceName, err)
	}
	_ = body
	return nil
}

//UnbindResource unbinds 'boundResourceName' from 'boundToResource'
func (c *NitroClient) UnbindResource(boundToResourceType string, boundToResourceName string, boundResourceType string, boundResourceName string, bindingFilterName string) error {

	if c.ResourceExists(boundToResourceType, boundToResourceName) == false {
		log.Printf("[INFO] go-nitro: Unbind: BoundTo Resource %s of type %s does not exist", boundToResourceType, boundToResourceName)
		return nil
	}

	if c.ResourceExists(boundResourceType, boundResourceName) == false {
		log.Printf("[INFO] go-nitro: Unbind: Bound Resource %s of type %s does not exist", boundResourceType, boundResourceName)
		return nil
	}

	_, err := c.unbindResource(boundToResourceType, boundToResourceName, boundResourceType, boundResourceName, bindingFilterName)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to unbind  %s:%s from %s:%s, err=%s", boundResourceType, boundResourceName, boundToResourceType, boundToResourceName, err)
	}

	return nil
}

//ResourceExists returns true if supplied resource name and type exists
func (c *NitroClient) ResourceExists(resourceType string, resourceName string) bool {
	_, err := c.listResource(resourceType, resourceName)
	if err != nil {
		log.Printf("[INFO] go-nitro: No %s %s found", resourceType, resourceName)
		return false
	}
	log.Printf("[INFO] go-nitro: %s %s is already present", resourceType, resourceName)
	return true
}

//FindResourceArray returns the config of the supplied resource name and type if it exists. Use when the resource to be returned is an array
func (c *NitroClient) FindResourceArray(resourceType string, resourceName string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, resourceName)
	if err != nil {
		log.Printf("[WARN] go-nitro: FindResourceArray: No %s %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindResourceArray: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindResourceArray: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindResourceArray: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		log.Printf("[WARN] go-nitro: FindResourceArray No %s type with name %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindResourceArray: No resource %s of type %s found", resourceName, resourceType)
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
		log.Printf("[WARN] go-nitro: FindFilteredResourceArray: No %s found matching filter %s", resourceType, filter)
		return nil, fmt.Errorf("[INFO] go-nitro: FindFilteredResourceArray: No resource of type %s found matching filter %s", resourceType, filter)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Println("[ERROR] go-nitro: FindFilteredResourceArray: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindFilteredResourceArray: Failed to unmarshal Netscaler Response:resource of type %s matching filter %s", resourceType, filter)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		log.Printf("[WARN] go-nitro: FindFilteredResourceArray No %s type matching filter %s found", resourceType, filter)
		return nil, fmt.Errorf("[INFO] go-nitro: FindFilteredResourceArray: No resource type %s matching filter %s found", filter, resourceType)
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
		log.Printf("[WARN] go-nitro: FindResource: No %s %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindResource: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindResource: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindResource: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrc, ok := data[resourceType]
	if !ok || rsrc == nil {
		log.Printf("[WARN] go-nitro: FindResource No %s type with name %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindResource: No resource %s of type %s found", resourceName, resourceType)
	}

	switch result := data[resourceType].(type) {
	case map[string]interface{}:
		return result, nil
	case []interface{}:
		return result[0].(map[string]interface{}), nil
	default:
		log.Printf("[WARN] go-nitro: FindResource Unable to determine type of response")
		return nil, fmt.Errorf("[INFO] go-nitro: FindResource: Unable to determine type of response")
	}
}

//FindAllResources finds all config objects of the supplied resource type and returns them in an array
func (c *NitroClient) FindAllResources(resourceType string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, "")
	if err != nil {
		log.Printf("[INFO] go-nitro: FindAllResources: No %s objects found", resourceType)
		return make([]map[string]interface{}, 0, 0), nil
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindAllResources: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindAllResources: Failed to unmarshal Netscaler Response: of type %s", resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		log.Printf("[INFO] go-nitro: FindAllResources: No %s found", resourceType)
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
		log.Printf("[INFO] go-nitro: ResourceBindingExists: No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
		return false
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: ResourceBindingExists: Failed to unmarshal Netscaler Response!")
		return false
	}
	binding := fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)
	rsrc, ok := data[binding]
	if !ok || rsrc == nil {
		return false
	}

	log.Printf("[INFO] go-nitro: ResourceBindingExists: %s of type  %s is bound to %s type and name %s", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
	return true
}

//FindBoundResource finds a bound resource if it exists
func (c *NitroClient) FindBoundResource(resourceType string, resourceName string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) (map[string]interface{}, error) {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, boundResourceFilterName, boundResourceFilterValue)
	if err != nil {
		log.Printf("[INFO] go-nitro: FindBoundResource: No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
		return nil, fmt.Errorf("[INFO] go-nitro: No %s %s to %s %s binding found, err=%s", resourceType, resourceName, boundResourceType, boundResourceFilterValue, err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindBoundResource: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindBoundResource: Failed to unmarshal Netscaler Response!, err=%s", err)
	}
	bindingType := fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)
	rsrc, ok := data[bindingType]
	if !ok || rsrc == nil {
		return nil, fmt.Errorf("[WARN] go-nitro: FindBoundResource: No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
	}

	resource := data[bindingType].([]interface{})[0] //only one resource obviously
	return resource.(map[string]interface{}), nil

}

//FindAllBoundResources returns an array of bound config objects of the type specified that are bound to the resource specified
func (c *NitroClient) FindAllBoundResources(resourceType string, resourceName string, boundResourceType string) ([]map[string]interface{}, error) {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, "", "")
	if err != nil {
		log.Printf("[INFO] go-nitro: FindAllBoundResources: No %s %s to %s  binding found", resourceType, resourceName, boundResourceType)
		return nil, fmt.Errorf("[ERROR] go-nitro: No %s %s to %s binding found, err=%s", resourceType, resourceName, boundResourceType, err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindAllBoundResources: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindAllBoundResources: Failed to unmarshal Netscaler Response!, err=%s", err)
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
		log.Printf("[ERROR] go-nitro: EnableFeatures: Failed to marshal features to JSON")
		return fmt.Errorf("[ERROR] go-nitro: Failed to marshal features to JSON")
	}

	_, err = c.enableFeatures(featureJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to enable feature %v", err)
	}
	return nil
}

//ListEnabledFeatures returns a string array of the list of features enabled on the NetScaler appliance
func (c *NitroClient) ListEnabledFeatures() ([]string, error) {

	bytes, err := c.listEnabledFeatures()
	if err != nil {
		return []string{}, fmt.Errorf("[ERROR] go-nitro: Failed to list features %v", err)
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindAllBoundResources: Failed to unmarshal Netscaler Response!")
		return []string{}, fmt.Errorf("[ERROR] go-nitro: Failed to unmarshal Netscaler Response to list Features")
	}
	feat, ok := data["nsfeature"]
	if !ok || feat == nil {
		log.Printf("No features found")
		return []string{}, fmt.Errorf("[ERROR] go-nitro: No features found")
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
		log.Printf("[ERROR] go-nitro: SaveConfig: Failed to marshal save config to JSON")
		return fmt.Errorf("[ERROR] go-nitro: Failed to marshal save config to JSON")
	}
	log.Printf("saveJSON is " + string(saveJSON))

	err = c.saveConfig(saveJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to save config %v", err)
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
		log.Printf("[ERROR] go-nitro: ClearConfig: Failed to marshal clear config to JSON")
		return fmt.Errorf("[ERROR] go-nitro: Failed to marshal clear config to JSON")
	}
	log.Printf("clearJSON is " + string(clearJSON))

	err = c.clearConfig(clearJSON)
	if err != nil {
		return fmt.Errorf("[ERROR] go-nitro: Failed to clear config %v", err)
	}
	return nil
}
