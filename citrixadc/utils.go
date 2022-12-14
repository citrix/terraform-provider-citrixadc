package citrixadc

import (
	"log"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"
)

func isTargetAdcCluster(nsClient *service.NitroClient) bool {
	log.Printf("[DEBUG]  citrixadc-provider-test: In isTargetAdcCluster")

	datalist, err := nsClient.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		//lintignore:R009
		panic(err)
	}

	if len(datalist) == 0 {
		return false
	} else {
		return true
	}
}

func toStringList(in []interface{}) []string {
	out := make([]string, 0, len(in))
	for _, val := range in {
		out = append(out, val.(string))
	}
	return out
}

func toIntegerList(in []interface{}) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = in[i].(int)
	}
	return out
}

// Check if the attribute is int, if not convert to int and set the value
func setToInt(attributeName string, d *schema.ResourceData, value interface{}) {
	var v int
	var err error

	switch valueTyped := value.(type) {
	case int:
		v = valueTyped
	case string:
		v, _ = strconv.Atoi(valueTyped)
	case nil:
		v = 0
	default:
		log.Printf("[DEBUG] got unexpected type %T for int", value)
		return
	}

	if err != nil {
		return
	}

	d.Set(attributeName, v)
}
