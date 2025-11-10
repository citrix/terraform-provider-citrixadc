package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVxlanvlanmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVxlanvlanmapFunc,
		ReadContext:   readVxlanvlanmapFunc,
		DeleteContext: deleteVxlanvlanmapFunc,
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
		},
	}
}

func createVxlanvlanmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Get("name").(string)
	vxlanvlanmap := network.Vxlanvlanmap{
		Name: d.Get("name").(string),
	}

	_, err := client.AddResource("vxlanvlanmap", vxlanvlanmapName, &vxlanvlanmap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vxlanvlanmapName)

	return readVxlanvlanmapFunc(ctx, d, meta)
}

func readVxlanvlanmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vxlanvlanmap state %s", vxlanvlanmapName)
	data, err := client.FindResource("vxlanvlanmap", vxlanvlanmapName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vxlanvlanmap state %s", vxlanvlanmapName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])

	return nil

}

func deleteVxlanvlanmapFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlanvlanmapFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanvlanmapName := d.Id()
	err := client.DeleteResource("vxlanvlanmap", vxlanvlanmapName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
