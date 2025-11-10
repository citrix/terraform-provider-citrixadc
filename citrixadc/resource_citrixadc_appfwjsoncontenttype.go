package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwjsoncontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwjsoncontenttypeFunc,
		ReadContext:   readAppfwjsoncontenttypeFunc,
		DeleteContext: deleteAppfwjsoncontenttypeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"jsoncontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwjsoncontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Get("jsoncontenttypevalue").(string)
	appfwjsoncontenttype := appfw.Appfwjsoncontenttype{
		Isregex:              d.Get("isregex").(string),
		Jsoncontenttypevalue: appfwjsoncontenttypeName,
	}

	_, err := client.AddResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeName, &appfwjsoncontenttype)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwjsoncontenttypeName)

	return readAppfwjsoncontenttypeFunc(ctx, d, meta)
}

func readAppfwjsoncontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwjsoncontenttype state %s", appfwjsoncontenttypeName)
	data, err := client.FindResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwjsoncontenttype state %s", appfwjsoncontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("jsoncontenttypevalue", data["jsoncontenttypevalue"])

	return nil

}

func deleteAppfwjsoncontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwjsoncontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwjsoncontenttypeName := d.Id()
	err := client.DeleteResource(service.Appfwjsoncontenttype.Type(), appfwjsoncontenttypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
