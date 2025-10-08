package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcFis() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFisFunc,
		ReadContext:   readFisFunc,
		DeleteContext: deleteFisFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"ownernode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createFisFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Get("name").(string)
	fis := network.Fis{
		Name:      d.Get("name").(string),
		Ownernode: d.Get("ownernode").(int),
	}

	_, err := client.AddResource(service.Fis.Type(), fisName, &fis)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fisName)

	return readFisFunc(ctx, d, meta)
}

func readFisFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading fis state %s", fisName)
	data, err := client.FindResource(service.Fis.Type(), fisName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing fis state %s", fisName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("ownernode", d, data["ownernode"])

	return nil

}

func deleteFisFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFisFunc")
	client := meta.(*NetScalerNitroClient).client
	fisName := d.Id()
	err := client.DeleteResource(service.Fis.Type(), fisName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
