package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwxmlschema() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwxmlschemaFunc,
		ReadContext:   readAppfwxmlschemaFunc,
		DeleteContext: deleteAppfwxmlschemaFunc,
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
			"src": {
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
		},
	}
}

func createAppfwxmlschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(appfwxmlschemaName)

	return readAppfwxmlschemaFunc(ctx, d, meta)
}

func readAppfwxmlschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteAppfwxmlschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlschemaName := d.Id()
	err := client.DeleteResource(service.Appfwxmlschema.Type(), appfwxmlschemaName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
