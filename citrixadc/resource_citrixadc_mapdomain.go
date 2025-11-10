package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcMapdomain() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createMapdomainFunc,
		ReadContext:   readMapdomainFunc,
		DeleteContext: deleteMapdomainFunc,
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
			"mapdmrname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createMapdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Get("name").(string)
	mapdomain := network.Mapdomain{
		Mapdmrname: d.Get("mapdmrname").(string),
		Name:       d.Get("name").(string),
	}

	_, err := client.AddResource("mapdomain", mapdomainName, &mapdomain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(mapdomainName)

	return readMapdomainFunc(ctx, d, meta)
}

func readMapdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading mapdomain state %s", mapdomainName)
	data, err := client.FindResource("mapdomain", mapdomainName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain state %s", mapdomainName)
		d.SetId("")
		return nil
	}
	d.Set("mapdmrname", data["mapdmrname"])
	d.Set("name", data["name"])

	return nil

}

func deleteMapdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	mapdomainName := d.Id()
	err := client.DeleteResource("mapdomain", mapdomainName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
