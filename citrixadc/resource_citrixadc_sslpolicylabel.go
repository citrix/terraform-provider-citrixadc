package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSslpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslpolicylabelFunc,
		ReadContext:   readSslpolicylabelFunc,
		DeleteContext: deleteSslpolicylabelFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslpolicylabelName = d.Get("labelname").(string)

	sslpolicylabel := ssl.Sslpolicylabel{
		Labelname: sslpolicylabelName,
		Type:      d.Get("type").(string),
	}

	_, err := client.AddResource(service.Sslpolicylabel.Type(), sslpolicylabelName, &sslpolicylabel)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslpolicylabelName)

	return readSslpolicylabelFunc(ctx, d, meta)
}

func readSslpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicylabelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslpolicylabel state %s", sslpolicylabelName)
	data, err := client.FindResource(service.Sslpolicylabel.Type(), sslpolicylabelName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslpolicylabel state %s", sslpolicylabelName)
		d.SetId("")
		return nil
	}
	d.Set("labelname", data["labelname"])
	d.Set("labelname", data["labelname"])
	d.Set("type", data["type"])

	return nil

}

func deleteSslpolicylabelFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslpolicylabelFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicylabelName := d.Id()
	err := client.DeleteResource(service.Sslpolicylabel.Type(), sslpolicylabelName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
