package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/filter"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFilteraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFilteractionFunc,
		Read:          readFilteractionFunc,
		Update:        updateFilteractionFunc,
		Delete:        deleteFilteractionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"qual": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"respcode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFilteractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Get("name").(string)
	filteraction := filter.Filteraction{
		Name:        d.Get("name").(string),
		Page:        d.Get("page").(string),
		Qual:        d.Get("qual").(string),
		Respcode:    d.Get("respcode").(int),
		Servicename: d.Get("servicename").(string),
		Value:       d.Get("value").(string),
	}

	_, err := client.AddResource(service.Filteraction.Type(), filteractionName, &filteraction)
	if err != nil {
		return err
	}

	d.SetId(filteractionName)

	err = readFilteractionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this filteraction but we can't read it ?? %s", filteractionName)
		return nil
	}
	return nil
}

func readFilteractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading filteraction state %s", filteractionName)
	data, err := client.FindResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing filteraction state %s", filteractionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("page", data["page"])
	d.Set("qual", data["qual"])
	d.Set("respcode", data["respcode"])
	d.Set("servicename", data["servicename"])
	d.Set("value", data["value"])

	return nil

}

func updateFilteractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Get("name").(string)

	filteraction := filter.Filteraction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("page") {
		log.Printf("[DEBUG]  citrixadc-provider: Page has changed for filteraction %s, starting update", filteractionName)
		filteraction.Page = d.Get("page").(string)
		hasChange = true
	}
	if d.HasChange("qual") {
		log.Printf("[DEBUG]  citrixadc-provider: Qual has changed for filteraction %s, starting update", filteractionName)
		filteraction.Qual = d.Get("qual").(string)
		hasChange = true
	}
	if d.HasChange("respcode") {
		log.Printf("[DEBUG]  citrixadc-provider: Respcode has changed for filteraction %s, starting update", filteractionName)
		filteraction.Respcode = d.Get("respcode").(int)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicename has changed for filteraction %s, starting update", filteractionName)
		filteraction.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for filteraction %s, starting update", filteractionName)
		filteraction.Value = d.Get("value").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Filteraction.Type(), filteractionName, &filteraction)
		if err != nil {
			return fmt.Errorf("Error updating filteraction %s", filteractionName)
		}
	}
	return readFilteractionFunc(d, meta)
}

func deleteFilteractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFilteractionFunc")
	client := meta.(*NetScalerNitroClient).client
	filteractionName := d.Id()
	err := client.DeleteResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
