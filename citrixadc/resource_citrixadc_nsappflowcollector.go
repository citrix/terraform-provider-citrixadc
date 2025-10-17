package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcNsappflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsappflowcollectorFunc,
		ReadContext:   readNsappflowcollectorFunc,
		DeleteContext: deleteNsappflowcollectorFunc,
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
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsappflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Get("name").(string)
	nsappflowcollector := ns.Nsappflowcollector{
		Ipaddress: d.Get("ipaddress").(string),
		Name:      d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		nsappflowcollector.Port = intPtr(d.Get("port").(int))
	}

	_, err := client.AddResource(service.Nsappflowcollector.Type(), nsappflowcollectorName, &nsappflowcollector)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsappflowcollectorName)

	return readNsappflowcollectorFunc(ctx, d, meta)
}

func readNsappflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsappflowcollector state %s", nsappflowcollectorName)
	data, err := client.FindResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsappflowcollector state %s", nsappflowcollectorName)
		d.SetId("")
		return nil
	}
	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])
	setToInt("port", d, data["port"])

	return nil

}

func deleteNsappflowcollectorFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsappflowcollectorFunc")
	client := meta.(*NetScalerNitroClient).client
	nsappflowcollectorName := d.Id()
	err := client.DeleteResource(service.Nsappflowcollector.Type(), nsappflowcollectorName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
