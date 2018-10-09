package netscaler

import (
	"log"
	"testing"

	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
)

func TestNitroClient_FindAllStats(t *testing.T) {
	lbName1 := "test_lb_" + randomString(5)
	lbName2 := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName1,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	lb2 := lb.Lbvserver{
		Name:        lbName2,
		Ipv46:       randomIP(),
		Lbmethod:    "LEASTCONNECTION",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName1, &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName1)
		log.Println("Cannot continue")
		return
	}
	_, err = client.AddResource(Lbvserver.Type(), lbName2, &lb2)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName2)
		log.Println("Cannot continue")
		return
	}
	_, err = client.FindAllStats(Lbvserver.Type())
	if err != nil {
		t.Error("Did not find statistics of type ", err, Lbvserver.Type())
	}
}

func TestNitroClient_FindStats(t *testing.T) {
	lbName1 := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName1,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}

	_, err := client.AddResource(Lbvserver.Type(), lbName1, &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName1)
		log.Println("Cannot continue")
		return
	}

	svcName1 := "test_svc_" + randomString(5)
	svcName2 := "test_svc_" + randomString(5)
	service1 := basic.Service{
		Name:        svcName1,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}
	service2 := basic.Service{
		Name:        svcName2,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName1, &service1)
	if err != nil {
		t.Error("Could not create service service1", err)
		log.Println("Cannot continue")
		return
	}
	_, err = client.AddResource(Service.Type(), svcName2, &service2)
	if err != nil {
		t.Error("Could not create service service2", err)
		log.Println("Cannot continue")
		return
	}
	for _, resourceType := range []string{Lbvserver.Type(), Service.Type(), Gslbvserver.Type(), Gslbservice.Type()} {
		rsrc, err := client.FindAllResources(resourceType)
		if err != nil {
			// Ignore the erratic resource type
			continue
		}
		for _, availableItem := range rsrc {
			_, err = client.FindStat(resourceType, availableItem["name"].(string))
			if err != nil {
				t.Fatal(err)
			}
			// only check one
			break
		}
	}
}
