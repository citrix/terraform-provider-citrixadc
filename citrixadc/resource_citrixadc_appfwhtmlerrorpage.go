package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func appfwhtmlerrorpage() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwhtmlerrorpageFunc,
		ReadContext:   readAppfwhtmlerrorpageFunc,
		DeleteContext: deleteAppfwhtmlerrorpageFunc,
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

func createAppfwhtmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(appfwhtmlerrorpageName)

	return readAppfwhtmlerrorpageFunc(ctx, d, meta)
}

func readAppfwhtmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func deleteAppfwhtmlerrorpageFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwhtmlerrorpageFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwhtmlerrorpageName := d.Id()
	err := client.DeleteResource(service.Appfwhtmlerrorpage.Type(), appfwhtmlerrorpageName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
