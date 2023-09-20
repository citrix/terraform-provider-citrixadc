package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcResponderhtmlpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderhtmlpageFunc,
		Read:          readResponderhtmlpageFunc,
		Delete:        deleteResponderhtmlpageFunc,
		Schema: map[string]*schema.Schema{
			"cacertfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"overwrite": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"src": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createResponderhtmlpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderhtmlpageFunc")
	client := meta.(*NetScalerNitroClient).client
	responderhtmlpageName := d.Get("name").(string)
	responderhtmlpage := responder.Responderhtmlpage{
		Cacertfile: d.Get("cacertfile").(string),
		Comment:    d.Get("comment").(string),
		Name:       d.Get("name").(string),
		Overwrite:  d.Get("overwrite").(bool),
		Src:        d.Get("src").(string),
	}

	err := client.ActOnResource(service.Responderhtmlpage.Type(), &responderhtmlpage, "Import")
	if err != nil {
		return err
	}

	d.SetId(responderhtmlpageName)

	err = readResponderhtmlpageFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderhtmlpage but we can't read it ?? %s", responderhtmlpageName)
		return nil
	}
	return nil
}

func readResponderhtmlpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderhtmlpageFunc")
	client := meta.(*NetScalerNitroClient).client
	responderhtmlpageName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderhtmlpage state %s", responderhtmlpageName)
	data, err := client.FindResource(service.Responderhtmlpage.Type(), responderhtmlpageName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderhtmlpage state %s", responderhtmlpageName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteResponderhtmlpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderhtmlpageFunc")
	client := meta.(*NetScalerNitroClient).client
	responderhtmlpageName := d.Id()
	err := client.DeleteResource(service.Responderhtmlpage.Type(), responderhtmlpageName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
