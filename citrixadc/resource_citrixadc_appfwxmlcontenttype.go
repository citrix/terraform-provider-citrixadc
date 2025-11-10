package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwxmlcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwxmlcontenttypeFunc,
		ReadContext:   readAppfwxmlcontenttypeFunc,
		DeleteContext: deleteAppfwxmlcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"isregex": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"xmlcontenttypevalue": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAppfwxmlcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Get("xmlcontenttypevalue").(string)
	appfwxmlcontenttype := appfw.Appfwxmlcontenttype{
		Isregex:             d.Get("isregex").(string),
		Xmlcontenttypevalue: appfwxmlcontenttypeName,
	}

	_, err := client.AddResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName, &appfwxmlcontenttype)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwxmlcontenttypeName)

	return readAppfwxmlcontenttypeFunc(ctx, d, meta)
}

func readAppfwxmlcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwxmlcontenttype state %s", appfwxmlcontenttypeName)
	data, err := client.FindResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwxmlcontenttype state %s", appfwxmlcontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("xmlcontenttypevalue", data["xmlcontenttypevalue"])

	return nil

}

func deleteAppfwxmlcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwxmlcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwxmlcontenttypeName := d.Id()
	err := client.DeleteResource(service.Appfwxmlcontenttype.Type(), appfwxmlcontenttypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
