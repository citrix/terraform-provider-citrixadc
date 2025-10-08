package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSsldhparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSsldhparamFunc,
		Read:          schema.Noop,
		DeleteContext: deleteSsldhparamFunc,
		Schema: map[string]*schema.Schema{
			"bits": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"dhfile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"gen": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSsldhparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSsldhparamFunc")
	client := meta.(*NetScalerNitroClient).client
	ssldhparamName := d.Get("dhfile").(string)

	ssldhparam := ssl.Ssldhparam{
		Bits:   d.Get("bits").(int),
		Dhfile: ssldhparamName,
		Gen:    d.Get("gen").(string),
	}

	err := client.ActOnResource(service.Ssldhparam.Type(), &ssldhparam, "create")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ssldhparamName)

	return nil
}

func deleteSsldhparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSsldhparamFunc")

	d.SetId("")

	return nil
}
