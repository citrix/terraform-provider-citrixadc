/*
Copyright 2016 Citrix Systems, Inc

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
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/config/ns"
)

var client *NitroClient

//Used to generate random config object names
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(125)+1, rand.Intn(252)+1, rand.Intn(252)+1, rand.Intn(252)+1)
}

//init random and client
func init() {
	rand.Seed(time.Now().UnixNano())
	var err error
	client, err = NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}

}

// Functional tests
func TestClearConfig(t *testing.T) {
	err := client.ClearConfig()
	if err != nil {
		t.Error("Could not clear config: ", err)
	}
}

func TestAdd(t *testing.T) {

	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	client.AddResource(Lbvserver.Type(), lbName, &lb1)

	rsrc, err := client.FindResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Did not find resource of type ", Lbvserver.Type(), ":", lbName)
	}
	val, ok := rsrc["ipv46"]
	if ok {
		if val != rndIP {
			t.Error("Wrong ipv46 for lb ", lbName, ": ", val)
		}
		val, ok = rsrc["lbmethod"]
		if val != "ROUNDROBIN" {
			t.Error("Wrong lbmethod for lb ", lbName, ": ", val)
		}
		val, ok = rsrc["servicetype"]
		if val != "HTTP" {
			t.Error("Wrong servicetype for lb ", lbName, ": ", val)
		}
	}
	if !ok {
		t.Error("Non existent property in retrieved lb ", lbName)
	}

	svcName := randomString(5)
	rndIP2 := randomIP()

	service1 := basic.Service{
		Name:        svcName,
		Ip:          rndIP2,
		Port:        80,
		Servicetype: "HTTP",
	}

	client.AddResource(Service.Type(), svcName, &service1)

	_, err = client.FindResource(Service.Type(), svcName)
	if err != nil {
		t.Error("Did not find resource of type ", Service.Type(), ":", svcName)
	}
}

func TestApply(t *testing.T) {
	aclName := "test_acl_" + randomString(5)
	acl1 := ns.Nsacl{
		Aclname:   aclName,
		Aclaction: "ALLOW",
		Srcip:     true,
		Srcipval:  "192.168.11.10",
		Destip:    true,
		Destipval: "192.183.83.11",
		Priority:  1100,
	}

	client.AddResource(Nsacl.Type(), aclName, &acl1)

	acls := ns.Nsacls{}
	client.ApplyResource(Nsacls.Type(), &acls)

	readAcls, err := client.FindResourceArray(Nsacl.Type(), aclName)
	if err != nil {
		t.Error("Did not find resource of type ", Nsacl.Type(), ":", aclName)
	}
	acl2 := readAcls[0]
	log.Println("Found acl, kernelstate= ", acl2["kernelstate"])
	if acl2["kernelstate"].(string) != "APPLIED" {
		t.Error("ACL created but not APPLIED ", Nsacl.Type(), ":", aclName)
	}
}

func TestUpdate(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB")
	}

	lb1 = lb.Lbvserver{
		Name:     lbName,
		Lbmethod: "LEASTCONNECTION",
	}
	client.UpdateResource(Lbvserver.Type(), lbName, &lb1)
	rsrc, err := client.FindResource(Lbvserver.Type(), lbName)
	if err != nil {
		t.Error("Did not find resource of type ", Lbvserver.Type(), ":", lbName)
	}
	val, ok := rsrc["lbmethod"]
	if ok {
		if val != "LEASTCONNECTION" {
			t.Error("Did not update lb method to LEASTCONNECTION")
		}
	}
	if !ok {
		t.Error("Failed to retrieve lb vserver object")
	}
}

func TestBindUnBind(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)
	rndIP2 := randomIP()
	svcName := "test_svc_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB")
	}
	service1 := basic.Service{
		Name:        svcName,
		Ip:          rndIP2,
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName, &service1)
	if err != nil {
		t.Error("Could not create service")
	}

	binding := lb.Lbvserverservicebinding{
		Name:        lbName,
		Servicename: svcName,
	}

	client.BindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, &binding)
	exists := client.ResourceBindingExists(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if !exists {
		t.Error("Failed to bind service to lb vserver")
	}
	err = client.UnbindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, "servicename")
	exists = client.ResourceBindingExists(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if exists {
		t.Error("Failed to unbind service to lb vserver")
	}

}

func TestFindBoundResource(t *testing.T) {
	lbName := "test_lb_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", "sample_lb_1")
	}
	svcName := "test_svc_" + randomString(5)
	service1 := basic.Service{
		Name:        svcName,
		Ip:          randomIP(),
		Port:        80,
		Servicetype: "HTTP",
	}

	_, err = client.AddResource(Service.Type(), svcName, &service1)
	if err != nil {
		t.Error("Failed to add resource of type ", Service.Type(), ":", svcName)

	}
	binding := lb.Lbvserverservicebinding{
		Name:        lbName,
		Servicename: svcName,
	}
	err = client.BindResource(Lbvserver.Type(), lbName, Service.Type(), svcName, &binding)
	if err != nil {
		t.Error("Failed to bind resource of type ", Service.Type(), ":", svcName)

	}
	result, err := client.FindBoundResource(Lbvserver.Type(), lbName, Service.Type(), "servicename", svcName)
	if err != nil {
		t.Error("Failed to find bound resource of type ", Service.Type(), ":", svcName)

	}
	//log.Println("Found bound resource ", result)
	if result["servicename"] != svcName {
		t.Error("Failed to find bound resource of type ", Service.Type(), ":", svcName)
	}

}

func TestDelete(t *testing.T) {
	rndIP := randomIP()
	lbName := "test_lb_" + randomString(5)

	lb1 := lb.Lbvserver{
		Name:        lbName,
		Ipv46:       rndIP,
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName, &lb1)
	if err != nil {
		t.Error("Could not create LB")
	}

	client.DeleteResource(Lbvserver.Type(), lbName)
	if client.ResourceExists(Lbvserver.Type(), lbName) {
		t.Error("Failed to delete ", lbName)
	}
}

func TestEnableFeatures(t *testing.T) {
	features := []string{"SSL", "CS"}
	err := client.EnableFeatures(features)
	if err != nil {
		t.Error("Failed to enable features")
	}
	result, err := client.ListEnabledFeatures()
	if err != nil {
		t.Error("Failed to retrieve features")
	}
	found := 0
	for _, f := range features {
		for _, r := range result {
			if f == r {
				found = found + 1
			}
		}
	}
	if found != len(features) {
		t.Error("Requested features do not match enabled features=", features, "result=", result)
	}
}

func TestSaveConfig(t *testing.T) {
	err := client.SaveConfig()
	if err != nil {
		t.Error("Failed to save config")
	}
}

func TestFindAllResources(t *testing.T) {
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
	}
	_, err = client.AddResource(Lbvserver.Type(), lbName2, &lb2)
	if err != nil {
		t.Error("Failed to add resource of type ", Lbvserver.Type(), ":", lbName2)
	}
	rsrcs, err := client.FindAllResources(Lbvserver.Type())
	if err != nil {
		t.Error("Did not find resources of type ", Lbvserver.Type())
	}
	if len(rsrcs) < 2 {
		t.Error("Found only ", len(rsrcs), " resources of type ", Lbvserver.Type(), " expected at least 2")
	}

	found := 0
	for _, v := range rsrcs {
		name := v["name"].(string)
		if name == lbName1 || name == lbName2 {
			found = found + 1
		}
	}
	if found != 2 {
		t.Error("Did not find all configured lbvservers")
	}

}

func TestFindAllBoundResources(t *testing.T) {
	lbName1 := "test_lb_" + randomString(5)
	svcName1 := "test_svc_" + randomString(5)
	svcName2 := "test_svc_" + randomString(5)
	lb1 := lb.Lbvserver{
		Name:        lbName1,
		Ipv46:       randomIP(),
		Lbmethod:    "ROUNDROBIN",
		Servicetype: "HTTP",
		Port:        8000,
	}
	_, err := client.AddResource(Lbvserver.Type(), lbName1, &lb1)
	if err != nil {
		t.Error("Could not create LB")
	}
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
		t.Error("Could not create service service1")
	}
	_, err = client.AddResource(Service.Type(), svcName2, &service2)
	if err != nil {
		t.Error("Could not create service service2")
	}

	binding1 := lb.Lbvserverservicebinding{
		Name:        lbName1,
		Servicename: svcName1,
	}
	binding2 := lb.Lbvserverservicebinding{
		Name:        lbName1,
		Servicename: svcName2,
	}

	err = client.BindResource(Lbvserver.Type(), lbName1, Service.Type(), svcName1, &binding1)
	if err != nil {
		t.Error("Could not bind service service1")
	}

	err = client.BindResource(Lbvserver.Type(), lbName1, Service.Type(), svcName2, &binding2)
	if err != nil {
		t.Error("Could not bind service service2")
	}
	rsrcs, err := client.FindAllBoundResources(Lbvserver.Type(), lbName1, Service.Type())
	if err != nil {
		t.Error("Did not find bound resources of type ", Service.Type())
	}
	if len(rsrcs) < 2 {
		t.Error("Found only ", len(rsrcs), " resources of type ", Service.Type(), " expected at least 2")
	}

	found := 0
	for _, v := range rsrcs {
		name := v["servicename"].(string)
		if name == svcName1 || name == svcName2 {
			found = found + 1
		}
	}
	if found != 2 {
		t.Error("Did not find all bound services")
	}

}
