package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwxmlerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwxmlerrorpageFunc,
		Read:          readAppfwxmlerrorpageFunc,
		Delete:        deleteAppfwxmlerrorpageFunc,
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

func createAppfwxmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwxmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlerrorpageName := d.Get("name").(string)
	appfwxmlerrorpage := appfw.Appfwxmlerrorpage{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource(service.Appfwxmlerrorpage.Type(), &appfwxmlerrorpage, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwxmlerrorpageName)

	err = readAppfwxmlerrorpageFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwxmlerrorpage but we can't read it ?? %s", appfwxmlerrorpageName)
		return nil
	}
	return nil
}

func readAppfwxmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwxmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlerrorpageName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwxmlerrorpage state %s", appfwxmlerrorpageName)
	data, err := client.FindResource(service.Appfwxmlerrorpage.Type(), appfwxmlerrorpageName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwxmlerrorpage state %s", appfwxmlerrorpageName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwxmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlerrorpageName := d.Id()
	err := client.DeleteResource(service.Appfwxmlerrorpage.Type(), appfwxmlerrorpageName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
