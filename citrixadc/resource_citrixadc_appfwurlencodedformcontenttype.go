package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAppfwurlencodedformcontenttype() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppfwurlencodedformcontenttypeFunc,
		ReadContext:   readAppfwurlencodedformcontenttypeFunc,
		DeleteContext: deleteAppfwurlencodedformcontenttypeFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"urlencodedformcontenttypevalue": {
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

func createAppfwurlencodedformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Get("urlencodedformcontenttypevalue").(string)
	appfwurlencodedformcontenttype := appfw.Appfwurlencodedformcontenttype{
		Isregex:                        d.Get("isregex").(string),
		Urlencodedformcontenttypevalue: d.Get("urlencodedformcontenttypevalue").(string),
	}

	_, err := client.AddResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName, &appfwurlencodedformcontenttype)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appfwurlencodedformcontenttypeName)

	return readAppfwurlencodedformcontenttypeFunc(ctx, d, meta)
}

func readAppfwurlencodedformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwurlencodedformcontenttype state %s", appfwurlencodedformcontenttypeName)
	data, err := client.FindResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwurlencodedformcontenttype state %s", appfwurlencodedformcontenttypeName)
		d.SetId("")
		return nil
	}
	d.Set("isregex", data["isregex"])
	d.Set("urlencodedformcontenttypevalue", data["urlencodedformcontenttypevalue"])

	return nil

}

func deleteAppfwurlencodedformcontenttypeFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwurlencodedformcontenttypeFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwurlencodedformcontenttypeName := d.Id()
	err := client.DeleteResource("appfwurlencodedformcontenttype", appfwurlencodedformcontenttypeName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
