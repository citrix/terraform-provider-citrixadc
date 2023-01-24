package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSpilloveractionFunc,
		Read:          readSpilloveractionFunc,
		Delete:        deleteSpilloveractionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createSpilloveractionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSpilloveractionFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloveractionName := d.Get("name").(string)
	spilloveraction := spillover.Spilloveraction{
		Action:  d.Get("action").(string),
		Name:    d.Get("name").(string),
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
	spilloveractionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading spilloveraction state %s", spilloveractionName)
	data, err := client.FindResource(service.Spilloveraction.Type(), spilloveractionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing spilloveraction state %s", spilloveractionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])

	return nil

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
