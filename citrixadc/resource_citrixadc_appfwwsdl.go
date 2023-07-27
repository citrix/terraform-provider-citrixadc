package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwwsdl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwwsdlFunc,
		Read:          readAppfwwsdlFunc,
		Delete:        deleteAppfwwsdlFunc,
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
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createAppfwwsdlFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwwsdlFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwwsdlName := d.Get("name").(string)

	appfwwsdl := appfw.Appfwwsdl{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource(service.Appfwwsdl.Type(), &appfwwsdl, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwwsdlName)

	err = readAppfwwsdlFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwwsdl but we can't read it ?? %s", appfwwsdlName)
		return nil
	}
	return nil
}

func readAppfwwsdlFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwwsdlFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwwsdlName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwwsdl state %s", appfwwsdlName)
	data, err := client.FindResource(service.Appfwwsdl.Type(), appfwwsdlName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwwsdl state %s", appfwwsdlName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwwsdlFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwwsdlFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwwsdlName := d.Id()
	err := client.DeleteResource(service.Appfwwsdl.Type(), appfwwsdlName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
