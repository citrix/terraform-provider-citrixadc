package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLsnip6profile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsnip6profileFunc,
		ReadContext:   readLsnip6profileFunc,
		DeleteContext: deleteLsnip6profileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"natprefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"network6": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLsnip6profileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Get("name").(string)
	lsnip6profile := lsn.Lsnip6profile{
		Name:      d.Get("name").(string),
		Natprefix: d.Get("natprefix").(string),
		Network6:  d.Get("network6").(string),
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource("lsnip6profile", lsnip6profileName, &lsnip6profile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lsnip6profileName)

	return readLsnip6profileFunc(ctx, d, meta)
}

func readLsnip6profileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsnip6profile state %s", lsnip6profileName)
	data, err := client.FindResource("lsnip6profile", lsnip6profileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsnip6profile state %s", lsnip6profileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("network6", data["network6"])
	d.Set("type", data["type"])

	return nil

}

func deleteLsnip6profileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsnip6profileFunc")
	client := meta.(*NetScalerNitroClient).client
	lsnip6profileName := d.Id()
	err := client.DeleteResource("lsnip6profile", lsnip6profileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
