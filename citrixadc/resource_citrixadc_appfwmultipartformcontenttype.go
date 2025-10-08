package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwmultipartformcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwmultipartformcontenttypeFunc,
		ReadContext:   readAppfwmultipartformcontenttypeFunc,
		DeleteContext: deleteAppfwmultipartformcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"multipartformcontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwmultipartformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Get("multipartformcontenttypevalue").(string)
	appfwmultipartformcontenttype := appfw.Appfwmultipartformcontenttype{
		Isregex:                       d.Get("isregex").(string),
		Multipartformcontenttypevalue: d.Get("multipartformcontenttypevalue").(string),
	}

	_, err := client.AddResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName, &appfwmultipartformcontenttype)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwmultipartformcontenttypeName)

	return readAppfwmultipartformcontenttypeFunc(ctx, d, meta)
}

func readAppfwmultipartformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwmultipartformcontenttype state %s", appfwmultipartformcontenttypeName)
	data, err := client.FindResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwmultipartformcontenttype state %s", appfwmultipartformcontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("multipartformcontenttypevalue", data["multipartformcontenttypevalue"])

	return nil

}

func deleteAppfwmultipartformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwmultipartformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwmultipartformcontenttypeName := d.Id()
	err := client.DeleteResource("appfwmultipartformcontenttype", appfwmultipartformcontenttypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
