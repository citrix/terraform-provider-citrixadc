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
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type NetscalerService struct {
	Name        string `json:"name"`
	Ip          string `json:"ip"`
	ServiceType string `json:"servicetype"`
	Port        int    `json:"port"`
}

type NetscalerLB struct {
	Name            string `json:"name"`
	Ipv46           string `json:"ipv46"`
	ServiceType     string `json:"servicetype"`
	Port            int    `json:"port"`
	PersistenceType string `json:"persistencetype,omitempty"`
	LbMethod        string `json:"lbmethod,omitempty"`
}

type NetscalerLBServiceBinding struct {
	Name        string `json:"name"`
	ServiceName string `json:"serviceName"`
}

type NetscalerCsAction struct {
	Name            string `json:"name"`
	TargetLBVserver string `json:"targetLBVserver"`
}

type NetscalerCsPolicy struct {
	PolicyName string `json:"policyName"`
	Rule       string `json:"rule"`
	Action     string `json:"action"`
}

type NetscalerCsPolicyBinding struct {
	Name       string `json:"name"`
	PolicyName string `json:"policyName"`
	Priority   int    `json:"priority"`
	Bindpoint  string `json:"bindpoint"`
}

type NetscalerCsVserver struct {
	Name        string `json:"name"`
	ServiceType string `json:"servicetype"`
	Ipv46       string `json:"ipv46"`
	Port        int    `json:"port"`
}

func (c *nitroClient) DeleteService(sname string) error {
	resourceType := "service"
	_, err := c.deleteResource(resourceType, sname)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to delete service %s err=%s", sname, err))
		return err
	}
	return nil
}

func (c *nitroClient) CreateService(serviceStruct *NetscalerService) (string, error) {
	resourceType := "service"
	sname := serviceStruct.Name
	if !c.ResourceExists("service", sname) {
		if serviceStruct.ServiceType == "" {
			serviceStruct.ServiceType = "HTTP"
		}
		nsService := &struct {
			Service NetscalerService `json:"service"`
		}{Service: *serviceStruct}
		resourceJson, err := json.Marshal(nsService)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to marshal service %s err=", sname, err))
			return sname, err
		}
		log.Println(string(resourceJson))

		body, err := c.createResource(resourceType, resourceJson)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to create service %s err=%s", sname, err))
			return sname, err
		}
		log.Printf("Created service with name %s", sname)
		_ = body
	} else {
		log.Printf("Found service with name %s, not recreating", sname)
	}
	return sname, nil

}

func (c *nitroClient) AddAndBindService(lbName string, svcStruct *NetscalerService) error {
	_, err := c.CreateService(svcStruct)
	if err != nil {
		return err
	}
	sname := svcStruct.Name
	//bind the lb to the service
	resourceType := "lbvserver"
	boundResourceType := "service"
	if !c.ResourceBindingExists(resourceType, lbName, boundResourceType, "servicename", sname) {
		nsLbSvcBinding := &struct {
			Lbvserver_service_binding NetscalerLBServiceBinding `json:"lbvserver_service_binding"`
		}{Lbvserver_service_binding: NetscalerLBServiceBinding{Name: lbName, ServiceName: sname}}
		resourceJson, err := json.Marshal(nsLbSvcBinding)

		resourceType = "lbvserver_service_binding"

		body, err := c.createResource(resourceType, resourceJson)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to bind lb %s to service %s, err=%s", lbName, sname, err))
			_ = c.DeleteService(sname) //TODO: the problem might be that it is bound to a different lb
			return err
		}
		_ = body
	} else {
		log.Printf("lb %s already bound to service %s", lbName, sname)
	}
	return nil
}

func (c *nitroClient) CreateLBVserver(lbStruct *NetscalerLB) (string, error) {

	resourceType := "lbvserver"
	if c.ResourceExists(resourceType, lbStruct.Name) == false {
		if lbStruct.ServiceType == "" {
			lbStruct.ServiceType = "HTTP"
		}
		if lbStruct.Ipv46 == "" || lbStruct.Ipv46 == "0.0.0.0" {
			errstr := fmt.Sprintf("VIP cannot be empty or 0.0.0.0 for lb %s", lbStruct.Name)
			log.Fatal(errstr)
			return "", errors.New(errstr)
		}
		if lbStruct.Port == 0 && lbStruct.ServiceType == "HTTP" {
			lbStruct.Port = 80
		}
		if lbStruct.Port == 0 {
			errstr := fmt.Sprintf("Port cannot be 0 for lb %s", lbStruct.Name)
			log.Fatal(errstr)
			return "", errors.New(errstr)
		}

		nsLB := &struct {
			Lbvserver NetscalerLB `json:"lbvserver"`
		}{Lbvserver: *lbStruct}
		resourceJson, err := json.Marshal(nsLB)

		log.Println("Resourcejson is " + string(resourceJson))

		body, err := c.createResource(resourceType, resourceJson)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to create lb %s, err=%s", lbStruct.Name, err))
			//TODO roll back
			return "", err
		}
		_ = body
	}

	return lbStruct.Name, nil
}

func (c *nitroClient) DeleteResource(resourceType string, resourceName string) error {

	_, err := c.listResource(resourceType, resourceName)
	if err == nil { // resource exists
		log.Printf("Found resource of type %s: %s", resourceType, resourceName)
		_, err = c.deleteResource(resourceType, resourceName)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to delete resourceType %: %s, err=%s", resourceType, resourceName, err))
			return err
		}
	} else {
		log.Printf("Resource %s already deleted ", resourceName)
	}
	return nil
}

func (c *nitroClient) DeleteLBVserver(lbName string) error {
	resourceType := "lbvserver"

	_, err := c.listResource(resourceType, lbName)
	if err == nil { // resource exists
		log.Printf("Found resource of type %s: %s", resourceType, lbName)
		_, err = c.deleteResource(resourceType, lbName)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to delete lb %s, err=%s", lbName, err))
			return err
		}
	} else {
		log.Printf("Lb %s already deleted ", lbName)
	}
	return nil
}

func (c *nitroClient) ResourceExists(resourceType string, resourceName string) bool {
	_, err := c.listResource(resourceType, resourceName)
	if err != nil {
		log.Printf("No %s %s found", resourceType, resourceName)
		return false
	}
	log.Printf("%s %s is already present", resourceType, resourceName)
	return true
}

func (c *nitroClient) ResourceBindingExists(resourceType string, resourceName string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) bool {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, boundResourceFilterName, boundResourceFilterValue)
	if err != nil {
		log.Printf("No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
		return false
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		log.Println("Failed to unmarshal Netscaler Response!")
		return false
	}
	if data[fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)] == nil {
		return false
	}

	log.Printf("%s %s is already bound to %s %s", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
	return true
}
