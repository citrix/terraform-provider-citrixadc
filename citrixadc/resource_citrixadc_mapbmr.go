package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcMapbmr() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createMapbmrFunc,
		ReadContext:   readMapbmrFunc,
		DeleteContext: deleteMapbmrFunc,
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
			"ruleipv6prefix": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"eabitlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psidlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"psidoffset": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createMapbmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Get("name").(string)
	mapbmr := network.Mapbmr{
		Eabitlength:    d.Get("eabitlength").(int),
		Name:           d.Get("name").(string),
		Psidlength:     d.Get("psidlength").(int),
		Psidoffset:     d.Get("psidoffset").(int),
		Ruleipv6prefix: d.Get("ruleipv6prefix").(string),
	}

	_, err := client.AddResource("mapbmr", mapbmrName, &mapbmr)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(mapbmrName)

	return readMapbmrFunc(ctx, d, meta)
}

func readMapbmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading mapbmr state %s", mapbmrName)
	data, err := client.FindResource("mapbmr", mapbmrName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing mapbmr state %s", mapbmrName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("eabitlength", d, data["eabitlength"])
	d.Set("name", data["name"])
	setToInt("psidlength", d, data["psidlength"])
	setToInt("psidoffset", d, data["psidoffset"])
	d.Set("ruleipv6prefix", data["ruleipv6prefix"])

	return nil

}

func deleteMapbmrFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapbmrFunc")
	client := meta.(*NetScalerNitroClient).client
	mapbmrName := d.Id()
	err := client.DeleteResource("mapbmr", mapbmrName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
