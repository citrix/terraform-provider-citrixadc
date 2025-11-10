package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLbmetrictable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbmetrictableFunc,
		ReadContext:   readLbmetrictableFunc,
		DeleteContext: deleteLbmetrictableFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"metrictable": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLbmetrictableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client

	lbmetrictableName := d.Get("metrictable").(string)

	lbmetrictable := lb.Lbmetrictable{
		Metrictable: lbmetrictableName,
	}

	_, err := client.AddResource("lbmetrictable", lbmetrictableName, &lbmetrictable)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(lbmetrictableName)

	return readLbmetrictableFunc(ctx, d, meta)
}

func readLbmetrictableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmetrictableName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbmetrictable state %s", lbmetrictableName)
	data, err := client.FindResource("lbmetrictable", lbmetrictableName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbmetrictable state %s", lbmetrictableName)
		d.SetId("")
		return nil
	}
	d.Set("metrictable", data["metrictable"])

	return nil

}

func deleteLbmetrictableFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbmetrictableFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmetrictableName := d.Id()
	err := client.DeleteResource("lbmetrictable", lbmetrictableName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
