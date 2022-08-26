package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/spillover"
	
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)


func resourceCitrixAdcSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSpilloveractionFunc,
		Read:          readSpilloveractionFunc,
		Update:        updateSpilloveractionFunc,
		Delete:        deleteSpilloveractionFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			
			
		},
	}
}

func createSpilloveractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	var spilloveractionName string
	if v, ok := d.GetOk("name"); ok {
		spilloveractionName = v.(string)
	} else {
		spilloveractionName= resource.PrefixedUniqueId("tf-spilloveraction-")
		d.Set("name", spilloveractionName)
	}
	spilloveraction := spillover.Spilloveraction{
		Action:           d.Get("action").(string),
		Name:           d.Get("name").(string),
		Newname:           d.Get("newname").(string),
		
	}

	_, err := client.AddResource(service.Spilloveraction.Type(), spilloveractionName, &spilloveraction)
	if err != nil {
		return err
	}

	d.SetId(spilloveractionName)
	
	err = readSpilloveractionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this spilloveraction but we can't read it ?? %s", spilloveractionName)
		return nil
	}
	return nil
}

func readSpilloveractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName:= d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading spilloveraction state %s", spilloveractionName)
	data, err := client.FindResource(service.Spilloveraction.Type(), spilloveractionName)
	if err != nil {
	log.Printf("[WARN] citrixadc-provider: Clearing spilloveraction state %s", spilloveractionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	

	return nil

}

func updateSpilloveractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Get("name").(string)

	spilloveraction := spillover.Spilloveraction{
		Name : d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for spilloveraction %s, starting update", spilloveractionName)
		spilloveraction.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for spilloveraction %s, starting update", spilloveractionName)
		spilloveraction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for spilloveraction %s, starting update", spilloveractionName)
		spilloveraction.Newname = d.Get("newname").(string)
		hasChange = true
	}
	

	if hasChange {
		_, err := client.UpdateResource(service.Spilloveraction.Type(), spilloveractionName, &spilloveraction)
		if err != nil {
			return fmt.Errorf("Error updating spilloveraction %s", spilloveractionName)
		}
	}
	return readSpilloveractionFunc(d, meta)
}

func deleteSpilloveractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Id()
	err := client.DeleteResource(service.Spilloveraction.Type(), spilloveractionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
