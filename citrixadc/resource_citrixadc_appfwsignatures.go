package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcAppfwsignatures() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwsignaturesFunc,
		Read:          readAppfwsignaturesFunc,
		Delete:        deleteAppfwsignaturesFunc,
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
			"merge": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mergedefault": &schema.Schema{
				Type:     schema.TypeBool,
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
			"preservedefactions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sha1": &schema.Schema{
				Type:     schema.TypeString,
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
			"vendortype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"xslt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Get("name").(string)
	
	appfwsignatures := appfw.Appfwsignatures{
		Comment:            d.Get("comment").(string),
		Merge:              d.Get("merge").(bool),
		Mergedefault:       d.Get("mergedefault").(bool),
		Name:               d.Get("name").(string),
		Overwrite:          d.Get("overwrite").(bool),
		Preservedefactions: d.Get("preservedefactions").(bool),
		Sha1:               d.Get("sha1").(string),
		Src:                d.Get("src").(string),
		Vendortype:         d.Get("vendortype").(string),
		Xslt:               d.Get("xslt").(string),
	}

	err := client.ActOnResource(service.Appfwsignatures.Type(), &appfwsignatures, "Import")
	if err != nil {
		return err
	}

	d.SetId(appfwsignaturesName)

	err = readAppfwsignaturesFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwsignatures but we can't read it ?? %s", appfwsignaturesName)
		return nil
	}
	return nil
}

func readAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwsignatures state %s", appfwsignaturesName)
	data, err := client.FindResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwsignatures state %s", appfwsignaturesName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	

	return nil

}

func deleteAppfwsignaturesFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwsignaturesFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsignaturesName := d.Id()
	err := client.DeleteResource(service.Appfwsignatures.Type(), appfwsignaturesName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
