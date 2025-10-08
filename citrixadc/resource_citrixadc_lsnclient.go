package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLsnclient() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnclientFunc,
		ReadContext:   readLsnclientFunc,
		DeleteContext: deleteLsnclientFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"clientname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnclientFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Get("clientname").(string)
	lsnclient := lsn.Lsnclient{
		Clientname: d.Get("clientname").(string),
	}

	_, err := client.AddResource("lsnclient", lsnclientName, &lsnclient)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnclientName)

	return readLsnclientFunc(ctx, d, meta)
}

func readLsnclientFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnclient state %s", lsnclientName)
	data, err := client.FindResource("lsnclient", lsnclientName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnclient state %s", lsnclientName)
		d.SetId("")
		return nil
	}
	d.Set("clientname", data["clientname"])

	return nil

}

func deleteLsnclientFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnclientFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnclientName := d.Id()
	err := client.DeleteResource("lsnclient", lsnclientName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
