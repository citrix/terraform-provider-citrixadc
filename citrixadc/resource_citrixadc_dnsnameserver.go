package citrixadc

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCitrixAdcDnsnameserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsnameserverFunc,
		Read:          readDnsnameserverFunc,
		Update:        updateDnsnameserverFunc,
		Delete:        deleteDnsnameserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"dnsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"local": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true, // Computed is often used to represent values that are not user configurable or can not be known at time of terraform plan or apply
				ForceNew: true,
			},
		},
	}
}

func createDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnameserver := dns.Dnsnameserver{
		Dnsprofilename: d.Get("dnsprofilename").(string),
		Local:          d.Get("local").(bool),
		State:          d.Get("state").(string),
		Type:           d.Get("type").(string),
	}
	var PrimaryId string
	if Ip, ok := d.GetOk("ip"); ok {
		PrimaryId = Ip.(string)
		dnsnameserver.Ip = PrimaryId
	} else if dnsvserver, ok := d.GetOk("dnsvservername"); ok {
		PrimaryId = dnsvserver.(string)
		dnsnameserver.Dnsvservername = PrimaryId
	}

	_, err := client.AddResource(service.Dnsnameserver.Type(), PrimaryId, &dnsnameserver)
	if err != nil {
		return err
	}
	if val, ok := d.GetOk("type"); ok {
		PrimaryId = PrimaryId + "," + val.(string)
	} else {
		// the default value of attribute type is "UDP". So, it is appended implicitly when not specified by the user.
		PrimaryId = PrimaryId + ",UDP"
	}

	d.SetId(PrimaryId)

	err = readDnsnameserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this resource but we can't read it ?? %s", PrimaryId)
		return nil
	}
	return nil
}

func readDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	PrimaryId := d.Id()

	// To make the resource backward compatible, in the prev state file user will have ID with 1 value, but in release v1.27.0 we have updated Id. So here we are changing the code to make it backward compatible
	// here we are checking for id, if it has 1 elements then we are appending the 2rd attribute to the old Id.
	oldIdSlice := strings.Split(PrimaryId, ",")

	if len(oldIdSlice) == 1 {
		if val, ok := d.GetOk("type"); ok {
			PrimaryId = PrimaryId + "," + val.(string)
		} else {
			PrimaryId = PrimaryId + ",UDP"
		}

		d.SetId(PrimaryId)
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading dnsnameserver state %s", PrimaryId)
	findParams := service.FindParams{
		ResourceType: service.Dnsnameserver.Type(),
	}

	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnameserver state %s", PrimaryId)
		d.SetId("")
		return nil
	}
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: dns nameserver does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	idSlice := strings.SplitN(PrimaryId, ",", 2)
	name := idSlice[0]
	dns_type := idSlice[1]

	foundIndex := -1
	for i, dnsnameserver := range dataArray {
		match := false
		if dnsnameserver["ip"] == name || dnsnameserver["dnsvservername"] == name {
			match = true
		}
		if match == true {
			if dnsnameserver["type"] != dns_type {
				match = false
			}
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsnameserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnameserver state %s", PrimaryId)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("ip", data["ip"])
	// attribute local is not part of GET response.
	d.Set("local", d.Get("local").(bool))
	d.Set("state", data["state"])
	d.Set("type", data["type"])

	return nil

}

func updateDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client

	PrimaryId := d.Id()

	idSlice := strings.SplitN(PrimaryId, ",", 2)
	name := idSlice[0]

	dnsnameserver := dns.Dnsnameserver{
		Ip:             d.Get("ip").(string),
		Dnsvservername: d.Get("dnsvservername").(string),
	}
	hasChange := false
	stateChange := false
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsprofilename has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}

	if d.HasChange("local") {
		log.Printf("[DEBUG]  citrixadc-provider: Local has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.Local = d.Get("local").(bool)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.State = d.Get("state").(string)
		stateChange = true
	}
	if stateChange {
		err := doDnsvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling dnsnameserver %s", PrimaryId)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Dnsnameserver.Type(), name, &dnsnameserver)
		if err != nil {
			return fmt.Errorf("Error updating dnsnameserver %s", PrimaryId)
		}
	}
	return readDnsnameserverFunc(d, meta)
}

func deleteDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	PrimaryId := d.Id()
	idSlice := strings.SplitN(PrimaryId, ",", 2)
	Name := idSlice[0]
	dns_type := idSlice[1]

	argsMap := make(map[string]string)
	if val, ok := d.GetOk("dnsvservername"); ok && Name == val { // if the user gives `dnsvservername`, then we need to directly call delete operation.
		err := client.DeleteResource(service.Dnsnameserver.Type(), Name)
		if err != nil {
			return err
		}

		d.SetId("")

		return nil
	}
	if val, ok := d.GetOk("type"); ok {
		argsMap["type"] = url.QueryEscape(val.(string))
	} else {
		argsMap["type"] = dns_type
	}
	err := client.DeleteResourceWithArgsMap(service.Dnsnameserver.Type(), Name, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
func doDnsvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doDnsvserverStateChange")

	dnsvserver := dns.Dnsnameserver{
		Ip:             d.Get("ip").(string),
		Dnsvservername: d.Get("dnsvservername").(string),
	}

	newstate := d.Get("state")

	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Dnsnameserver.Type(), dnsvserver, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		err := client.ActOnResource(service.Dnsnameserver.Type(), dnsvserver, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("doDnsvserverStateChange : \"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
