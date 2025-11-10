package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwjsonerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwjsonerrorpageFunc,
		ReadContext:   readAppfwjsonerrorpageFunc,
		DeleteContext: deleteAppfwjsonerrorpageFunc,
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
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createAppfwjsonerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(appfwjsonerrorpageName)

	return readAppfwjsonerrorpageFunc(ctx, d, meta)
}

func readAppfwjsonerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteAppfwjsonerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwjsonerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsonerrorpageName := d.Id()
	err := client.DeleteResource("appfwjsonerrorpage", appfwjsonerrorpageName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
