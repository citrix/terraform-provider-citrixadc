package citrixadc

import (
	"log"

	"github.com/citrix/adc-nitro-go/service"
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
