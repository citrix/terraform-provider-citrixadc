/*
Copyright 2021 Citrix Systems, Inc.

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
	"log"
	"fmt"
	"encoding/json"
)

//FindStats returns the statistics of the supplied resource type if it exists. Use when the resource to be returned is an array
func (c *NitroClient) FindAllStats(resourceType string) ([]map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listStat(resourceType, "")
	if err != nil {
		log.Printf("[WARN] nitro-go: FindStats: No %s found", resourceType)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStats: No type %s found", resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] nitro-go: FindStats: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindStats: Failed to unmarshal Netscaler Response: type %s", resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		log.Printf("[WARN] nitro-go: FindStats No %s type found", resourceType)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStats: No type of %s found", resourceType)
	}
	resources := data[resourceType].([]interface{})
	ret := make([]map[string]interface{}, len(resources), len(resources))
	for i, v := range resources {
		ret[i] = v.(map[string]interface{})
	}
	return ret, nil
}

//FindStat returns the config of the supplied resource name and type if it exists
func (c *NitroClient) FindStat(resourceType string, resourceName string) (map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listStat(resourceType, resourceName)
	if err != nil {
		log.Printf("[WARN] nitro-go: FindStat: No %s %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStat: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] nitro-go: FindStat: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindStat: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrc, ok := data[resourceType]
	if !ok || rsrc == nil {
		log.Printf("[WARN] nitro-go: FindStat No %s type with name %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStat: No resource %s of type %s found", resourceName, resourceType)
	}
	resource := data[resourceType].([]interface{})[0] //only one resource obviously

	return resource.(map[string]interface{}), nil
}

func (c *NitroClient) FindStatWithArgs(resourceType string, resourceName string, args []string) (map[string]interface{}, error) {

	var data map[string]interface{}
	result, err := c.listStatWithArgs(resourceType, resourceName, args)
	if err != nil {
		log.Printf("[WARN] nitro-go: FindStatWithArgs: No %s %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStatWithArgs: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] nitro-go: FindStatWithArgs: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] nitro-go: FindStatWithArgs: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrc, ok := data[resourceType]
	if !ok || rsrc == nil {
		log.Printf("[WARN] nitro-go: FindStatWithArgs No %s type with name %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] nitro-go: FindStatWithArgs: No resource %s of type %s found", resourceName, resourceType)
	}
	switch result := data[resourceType].(type) {
	case map[string]interface{}:
		return result, nil
	case []interface{}:
		return result[0].(map[string]interface{}), nil
	default:
		log.Printf("[WARN] nitro-go: FindStatWithArgs Unable to determine type of response")
		return nil, fmt.Errorf("[INFO] nitro-go: FindStatWithArgs: Unable to determine type of response")
	}
}
