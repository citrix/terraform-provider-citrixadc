package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwwsdl() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwwsdlFunc,
		ReadContext:   readAppfwwsdlFunc,
		DeleteContext: deleteAppfwwsdlFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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

func createAppfwwsdlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(appfwwsdlName)

	return readAppfwwsdlFunc(ctx, d, meta)
}

func readAppfwwsdlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteAppfwwsdlFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwwsdlFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwwsdlName := d.Id()
	err := client.DeleteResource(service.Appfwwsdl.Type(), appfwwsdlName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
