package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCmpaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCmpactionFunc,
		Read:          readCmpactionFunc,
		Update:        updateCmpactionFunc,
		Delete:        deleteCmpactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cmptype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"addvaryheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deltatype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"varyheadervalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCmpactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Get("name").(string)
	cmpaction := cmp.Cmpaction{
		Addvaryheader:   d.Get("addvaryheader").(string),
		Cmptype:         d.Get("cmptype").(string),
		Deltatype:       d.Get("deltatype").(string),
		Name:            d.Get("name").(string),
		Varyheadervalue: d.Get("varyheadervalue").(string),
	}

	_, err := client.AddResource(service.Cmpaction.Type(), cmpactionName, &cmpaction)
	if err != nil {
		return err
	}

	d.SetId(cmpactionName)

	err = readCmpactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this cmpaction but we can't read it ?? %s", cmpactionName)
		return nil
	}
	return nil
}

func readCmpactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading cmpaction state %s", cmpactionName)
	data, err := client.FindResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing cmpaction state %s", cmpactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	//d.Set("addvaryheader", data["addvaryheader"])
	d.Set("cmptype", data["cmptype"])
	//d.Set("deltatype", data["deltatype"])
	d.Set("varyheadervalue", data["varyheadervalue"])

	return nil

}

func updateCmpactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Get("name").(string)

	cmpaction := cmp.Cmpaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("addvaryheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Addvaryheader has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Addvaryheader = d.Get("addvaryheader").(string)
		hasChange = true
	}
	if d.HasChange("cmptype") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmptype has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Cmptype = d.Get("cmptype").(string)
		hasChange = true
	}
	if d.HasChange("varyheadervalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Varyheadervalue has changed for cmpaction %s, starting update", cmpactionName)
		cmpaction.Varyheadervalue = d.Get("varyheadervalue").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Cmpaction.Type(), &cmpaction)
		if err != nil {
			return fmt.Errorf("Error updating cmpaction %s", cmpactionName)
		}
	}
	return readCmpactionFunc(d, meta)
}

func deleteCmpactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCmpactionFunc")
	client := meta.(*NetScalerNitroClient).client
	cmpactionName := d.Id()
	err := client.DeleteResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
