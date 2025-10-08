package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcMapdmr() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createMapdmrFunc,
		ReadContext:   readMapdmrFunc,
		DeleteContext: deleteMapdmrFunc,
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
			"bripv6prefix": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createMapdmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapdmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdmrName := d.Get("name").(string)
	mapdmr := network.Mapdmr{
		Bripv6prefix: d.Get("bripv6prefix").(string),
		Name:         d.Get("name").(string),
	}

	_, err := client.AddResource("mapdmr", mapdmrName, &mapdmr)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(mapdmrName)

	return readMapdmrFunc(ctx, d, meta)
}

func readMapdmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapdmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdmrName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading mapdmr state %s", mapdmrName)
	data, err := client.FindResource("mapdmr", mapdmrName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing mapdmr state %s", mapdmrName)
		d.SetId("")
		return nil
	}
	d.Set("bripv6prefix", data["bripv6prefix"])
	d.Set("name", data["name"])

	return nil

}

func deleteMapdmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapdmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdmrName := d.Id()
	err := client.DeleteResource("mapdmr", mapdmrName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
