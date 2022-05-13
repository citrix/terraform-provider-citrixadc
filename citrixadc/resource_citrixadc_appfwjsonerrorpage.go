package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwjsonerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwjsonerrorpageFunc,
		Read:          readAppfwjsonerrorpageFunc,
		Delete:        deleteAppfwjsonerrorpageFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwjsonerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwjsonerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsonerrorpageName := d.Get("name").(string)
	
	appfwjsonerrorpage := appfw.Appfwjsonerrorpage{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource("appfwjsonerrorpage", &appfwjsonerrorpage, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwjsonerrorpageName)

	err = readAppfwjsonerrorpageFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwjsonerrorpage but we can't read it ?? %s", appfwjsonerrorpageName)
		return nil
	}
	return nil
}

func readAppfwjsonerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwjsonerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsonerrorpageName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwjsonerrorpage state %s", appfwjsonerrorpageName)
	data, err := client.FindResource("appfwjsonerrorpage", appfwjsonerrorpageName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwjsonerrorpage state %s", appfwjsonerrorpageName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwjsonerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwjsonerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsonerrorpageName := d.Id()
	err := client.DeleteResource("appfwjsonerrorpage", appfwjsonerrorpageName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
