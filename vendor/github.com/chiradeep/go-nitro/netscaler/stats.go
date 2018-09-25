package netscaler

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
		log.Printf("[WARN] go-nitro: FindStats: No %s found", resourceType)
		return nil, fmt.Errorf("[INFO] go-nitro: FindStats: No type %s found", resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindStats: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindStats: Failed to unmarshal Netscaler Response: type %s", resourceType)
	}
	rsrcs, ok := data[resourceType]
	if !ok || rsrcs == nil {
		log.Printf("[WARN] go-nitro: FindStats No %s type found", resourceType)
		return nil, fmt.Errorf("[INFO] go-nitro: FindStats: No type of %s found", resourceType)
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
		log.Printf("[WARN] go-nitro: FindStat: No %s %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindStat: No resource %s of type %s found", resourceName, resourceType)
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Printf("[ERROR] go-nitro: FindStat: Failed to unmarshal Netscaler Response!")
		return nil, fmt.Errorf("[ERROR] go-nitro: FindStat: Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType)
	}
	rsrc, ok := data[resourceType]
	if !ok || rsrc == nil {
		log.Printf("[WARN] go-nitro: FindStat No %s type with name %s found", resourceType, resourceName)
		return nil, fmt.Errorf("[INFO] go-nitro: FindStat: No resource %s of type %s found", resourceName, resourceType)
	}
	resource := data[resourceType].([]interface{})[0] //only one resource obviously

	return resource.(map[string]interface{}), nil
}
