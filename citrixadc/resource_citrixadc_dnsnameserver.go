package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
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
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"local": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	} else if dnsvserver,ok := d.GetOk("dnsvservername");ok{
		PrimaryId = dnsvserver.(string)
		dnsnameserver.Dnsvservername = PrimaryId
	}

	_, err := client.AddResource(service.Dnsnameserver.Type(), PrimaryId, &dnsnameserver)
	if err != nil {
		return err
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
	foundIndex := -1
	for i, dnsnameserver := range dataArray {
		match := false
		if dnsnameserver["ip"] == d.Get("ip").(string) {
			match = true
		}else if dnsnameserver["dnsvservername"] == d.Get("dnsvservername").(string) {
			match = true
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
	d.Set("local", data["local"])
	d.Set("state", data["state"])
	d.Set("type", data["type"])

	return nil

}

func updateDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var PrimaryId string
	if Name, ok := d.GetOk("ip"); ok {
		PrimaryId = Name.(string)

	} else if Name,ok :=d.GetOk("dnsvservername");ok {
		PrimaryId = Name.(string)
	}

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
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnsnameserver %s, starting update", PrimaryId)
		dnsnameserver.Type = d.Get("type").(string)
		hasChange = true
	}
	if stateChange {
		err := doDnsvserverStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling cs vserver %s", PrimaryId)
		}
	}
	if hasChange {
		_, err := client.UpdateResource(service.Dnsnameserver.Type(), PrimaryId, &dnsnameserver)
		if err != nil {
			return fmt.Errorf("Error updating dnsnameserver %s", PrimaryId)
		}
	}
	return readDnsnameserverFunc(d, meta)
}

func deleteDnsnameserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnameserverFunc")
	client := meta.(*NetScalerNitroClient).client
	Name := d.Id()
	err := client.DeleteResource(service.Dnsnameserver.Type(), Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
func doDnsvserverStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doLbvserverStateChange")

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
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
