package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwxmlerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwxmlerrorpageFunc,
		ReadContext:   readAppfwxmlerrorpageFunc,
		DeleteContext: deleteAppfwxmlerrorpageFunc,
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

func createAppfwxmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(appfwxmlerrorpageName)

	return readAppfwxmlerrorpageFunc(ctx, d, meta)
}

func readAppfwxmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteAppfwxmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlerrorpageName := d.Id()
	err := client.DeleteResource(service.Appfwxmlerrorpage.Type(), appfwxmlerrorpageName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
