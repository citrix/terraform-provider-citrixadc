package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAutoscaleprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAutoscaleprofileFunc,
		Read:          readAutoscaleprofileFunc,
		Update:        updateAutoscaleprofileFunc,
		Delete:        deleteAutoscaleprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apikey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sharedsecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAutoscaleprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Get("name").(string)
	autoscaleprofile := autoscale.Autoscaleprofile{
		Apikey:       d.Get("apikey").(string),
		Name:         d.Get("name").(string),
		Sharedsecret: d.Get("sharedsecret").(string),
		Type:         d.Get("type").(string),
		Url:          d.Get("url").(string),
	}

	_, err := client.AddResource(service.Autoscaleprofile.Type(), autoscaleprofileName, &autoscaleprofile)
	if err != nil {
		return err
	}

	d.SetId(autoscaleprofileName)

	err = readAutoscaleprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this autoscaleprofile but we can't read it ?? %s", autoscaleprofileName)
		return nil
	}
	return nil
}

func readAutoscaleprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscaleprofile state %s", autoscaleprofileName)
	data, err := client.FindResource(service.Autoscaleprofile.Type(), autoscaleprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscaleprofile state %s", autoscaleprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("type", data["type"])
	d.Set("url", data["url"])

	return nil

}

func updateAutoscaleprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Get("name").(string)

	autoscaleprofile := autoscale.Autoscaleprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("apikey") {
		log.Printf("[DEBUG]  citrixadc-provider: Apikey has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Apikey = d.Get("apikey").(string)
		hasChange = true
	}
	if d.HasChange("sharedsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Sharedsecret has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Sharedsecret = d.Get("sharedsecret").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG]  citrixadc-provider: Url has changed for autoscaleprofile %s, starting update", autoscaleprofileName)
		autoscaleprofile.Url = d.Get("url").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Autoscaleprofile.Type(), &autoscaleprofile)
		if err != nil {
			return fmt.Errorf("Error updating autoscaleprofile %s", autoscaleprofileName)
		}
	}
	return readAutoscaleprofileFunc(d, meta)
}

func deleteAutoscaleprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscaleprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscaleprofileName := d.Id()
	err := client.DeleteResource(service.Autoscaleprofile.Type(), autoscaleprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
