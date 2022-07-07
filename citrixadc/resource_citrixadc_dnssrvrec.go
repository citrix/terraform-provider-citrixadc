package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcDnssrvrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnssrvrecFunc,
		Read:          readDnssrvrecFunc,
		Update:        updateDnssrvrecFunc,
		Delete:        deleteDnssrvrecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ecssubnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnssrvrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Get("domain").(string) + "," + d.Get("target").(string)
	dnssrvrec := dns.Dnssrvrec{
		Domain:    d.Get("domain").(string),
		Ecssubnet: d.Get("ecssubnet").(string),
		Nodeid:    d.Get("nodeid").(int),
		Port:      d.Get("port").(int),
		Priority:  d.Get("priority").(int),
		Target:    d.Get("target").(string),
		Ttl:       d.Get("ttl").(int),
		Type:      d.Get("type").(string),
		Weight:    d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Dnssrvrec.Type(), dnssrvrecName, &dnssrvrec)
	if err != nil {
		return err
	}

	d.SetId(dnssrvrecName)

	err = readDnssrvrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnssrvrec but we can't read it ?? %s", dnssrvrecName)
		return nil
	}
	return nil
}

func readDnssrvrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssrvrec state %s", dnssrvrecName)
	findParams := service.FindParams{
		ResourceType: service.Dnssrvrec.Type(),
	}
	dataArray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssrvrec state %s", dnssrvrecName)
		d.SetId("")
		return nil
	}
	
	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: dns srvrec does not exist. Clearing state.")
		d.SetId("")
		return nil
	}
	
	foundIndex := -1
	for i, dnssrvrec := range dataArray {
		match := true
		if dnssrvrec["domain"] != d.Get("domain").(string) {
			match = false
		}
		if dnssrvrec["target"] != d.Get("target").(string) {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnssrvrec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnssrvrec state %s", dnssrvrecName)
		d.SetId("")
		return nil
	}
	data := dataArray[foundIndex]
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("nodeid", data["nodeid"])
	d.Set("port", data["port"])
	d.Set("priority", data["priority"])
	d.Set("target", data["target"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])
	d.Set("weight", data["weight"])

	return nil

}

func updateDnssrvrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssrvrecName := d.Get("domain").(string)
	dnssrvrec := dns.Dnssrvrec{}
	log.Println(dnssrvrec)
	hasChange := false

	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Nodeid = d.Get("nodeid").(int)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for dnssrvrec %s, starting update", dnssrvrecName)
		dnssrvrec.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		dnssrvrec.Domain = d.Get("domain").(string)
		dnssrvrec.Target = d.Get("target").(string)
		err := client.UpdateUnnamedResource(service.Dnssrvrec.Type(), &dnssrvrec)
		if err != nil {
			return fmt.Errorf("Error updating dnssrvrec %s", dnssrvrecName)
		}
	}
	return readDnssrvrecFunc(d, meta)
}

func deleteDnssrvrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssrvrecFunc")
	client := meta.(*NetScalerNitroClient).client
	argsMap := make(map[string]string)
	argsMap["target"] = url.QueryEscape(d.Get("target").(string))
	if ecs,ok := d.GetOk("ecssubnet"); ok{
		argsMap["ecssubnet"] = url.QueryEscape(ecs.(string))
	}
	err := client.DeleteResourceWithArgsMap(service.Dnssrvrec.Type(),d.Id(), argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
