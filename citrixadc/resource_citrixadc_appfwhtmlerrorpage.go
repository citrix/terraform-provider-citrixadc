package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func appfwhtmlerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwhtmlerrorpageFunc,
		Read:          readAppfwhtmlerrorpageFunc,
		Delete:        deleteAppfwhtmlerrorpageFunc,
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

func createAppfwhtmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwhtmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwhtmlerrorpageName := d.Get("name").(string)
	
	appfwhtmlerrorpage := appfw.Appfwhtmlerrorpage{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource(service.Appfwhtmlerrorpage.Type(), &appfwhtmlerrorpage, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwhtmlerrorpageName)

	err = readAppfwhtmlerrorpageFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwhtmlerrorpage but we can't read it ?? %s", appfwhtmlerrorpageName)
		return nil
	}
	return nil
}

func readAppfwhtmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwhtmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwhtmlerrorpageName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwhtmlerrorpage state %s", appfwhtmlerrorpageName)
	data, err := client.FindResource(service.Appfwhtmlerrorpage.Type(), appfwhtmlerrorpageName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwhtmlerrorpage state %s", appfwhtmlerrorpageName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwhtmlerrorpageFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwhtmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwhtmlerrorpageName := d.Id()
	err := client.DeleteResource(service.Appfwhtmlerrorpage.Type(), appfwhtmlerrorpageName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
