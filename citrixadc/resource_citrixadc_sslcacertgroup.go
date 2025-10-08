package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSslcacertgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcacertgroupFunc,
		ReadContext:   readSslcacertgroupFunc,
		DeleteContext: deleteSslcacertgroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cacertgroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcacertgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Get("cacertgroupname").(string)

	sslcacertgroup := ssl.Sslcacertgroup{
		Cacertgroupname: sslcacertgroupName,
	}

	_, err := client.AddResource("sslcacertgroup", sslcacertgroupName, &sslcacertgroup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslcacertgroupName)

	return readSslcacertgroupFunc(ctx, d, meta)
}

func readSslcacertgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslcacertgroup state %s", sslcacertgroupName)
	data, err := client.FindResource("sslcacertgroup", sslcacertgroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslcacertgroup state %s", sslcacertgroupName)
		d.SetId("")
		return nil
	}
	d.Set("cacertgroupname", data["cacertgroupname"])

	return nil

}

func deleteSslcacertgroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcacertgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	sslcacertgroupName := d.Id()
	err := client.DeleteResource("sslcacertgroup", sslcacertgroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
