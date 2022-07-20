package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppflowcollectorFunc,
		Read:          readAppflowcollectorFunc,
		Update:        updateAppflowcollectorFunc,
		Delete:        deleteAppflowcollectorFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Get("name").(string)

	appflowcollector := appflow.Appflowcollector{
		Ipaddress:  d.Get("ipaddress").(string),
		Name:       d.Get("name").(string),
		Netprofile: d.Get("netprofile").(string),
		Port:       d.Get("port").(int),
		Transport:  d.Get("transport").(string),
	}

	_, err := client.AddResource(service.Appflowcollector.Type(), appflowcollectorName, &appflowcollector)
	if err != nil {
		return err
	}

	d.SetId(appflowcollectorName)

	err = readAppflowcollectorFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appflowcollector but we can't read it ?? %s", appflowcollectorName)
		return nil
	}
	return nil
}

func readAppflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowcollector state %s", appflowcollectorName)
	data, err := client.FindResource(service.Appflowcollector.Type(), appflowcollectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowcollector state %s", appflowcollectorName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("port", data["port"])
	d.Set("transport", data["transport"])

	return nil

}

func updateAppflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Get("name").(string)

	appflowcollector := appflow.Appflowcollector{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipaddress has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for appflowcollector %s, starting update", appflowcollectorName)
		appflowcollector.Port = d.Get("port").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Appflowcollector.Type(), appflowcollectorName, &appflowcollector)
		if err != nil {
			return fmt.Errorf("Error updating appflowcollector %s", appflowcollectorName)
		}
	}
	return readAppflowcollectorFunc(d, meta)
}

func deleteAppflowcollectorFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowcollectorName := d.Id()
	err := client.DeleteResource(service.Appflowcollector.Type(), appflowcollectorName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
