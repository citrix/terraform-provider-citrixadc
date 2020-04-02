package citrixadc

import (
	"log"

	"github.com/chiradeep/go-nitro/netscaler"
)

func isTargetAdcCluster(nsClient *netscaler.NitroClient) bool {
	log.Printf("[DEBUG]  citrixadc-provider-test: In isTargetAdcCluster")

	datalist, err := nsClient.FindAllResources(netscaler.Clusterinstance.Type())
	if err != nil {
		panic(err)
	}

	if len(datalist) == 0 {
		return false
	} else {
		return true
	}
}
