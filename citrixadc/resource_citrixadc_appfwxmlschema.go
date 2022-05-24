package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwxmlschema() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwxmlschemaFunc,
		Read:          readAppfwxmlschemaFunc,
		Delete:        deleteAppfwxmlschemaFunc,
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
			"src": &schema.Schema{
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
		},
	}
}

func createAppfwxmlschemaFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwxmlschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlschemaName := d.Get("name").(string)
	appfwxmlschema := appfw.Appfwxmlschema{
		Comment:   d.Get("comment").(string),
		Name:      d.Get("name").(string),
		Overwrite: d.Get("overwrite").(bool),
		Src:       d.Get("src").(string),
	}

	err := client.ActOnResource(service.Appfwxmlschema.Type(), &appfwxmlschema, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwxmlschemaName)

	err = readAppfwxmlschemaFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwxmlschema but we can't read it ?? %s", appfwxmlschemaName)
		return nil
	}
	return nil
}

func readAppfwxmlschemaFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwxmlschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlschemaName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwxmlschema state %s", appfwxmlschemaName)
	data, err := client.FindResource(service.Appfwxmlschema.Type(), appfwxmlschemaName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwxmlschema state %s", appfwxmlschemaName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteAppfwxmlschemaFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlschemaName := d.Id()
	err := client.DeleteResource(service.Appfwxmlschema.Type(), appfwxmlschemaName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
