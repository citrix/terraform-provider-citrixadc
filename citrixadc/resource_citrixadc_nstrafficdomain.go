package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceCitrixAdcNstrafficdomain() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstrafficdomainFunc,
		ReadContext:   readNstrafficdomainFunc,
		DeleteContext: deleteNstrafficdomainFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"td": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"aliasname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vmac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNstrafficdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	td_Id := d.Get("td").(int)
	nstrafficdomain := ns.Nstrafficdomain{
		Aliasname: d.Get("aliasname").(string),
		Vmac:      d.Get("vmac").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nstrafficdomain.Td = intPtr(d.Get("td").(int))
	}
	td_IdStr := strconv.Itoa(td_Id)

	_, err := client.AddResource(service.Nstrafficdomain.Type(), td_IdStr, &nstrafficdomain)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(td_IdStr)

	return readNstrafficdomainFunc(ctx, d, meta)
}

func readNstrafficdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	td_IdStr := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain state %s", td_IdStr)
	data, err := client.FindResource(service.Nstrafficdomain.Type(), td_IdStr)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain state %s", td_IdStr)
		d.SetId("")
		return nil
	}
	d.Set("aliasname", data["aliasname"])
	setToInt("td", d, data["td"])
	d.Set("vmac", data["vmac"])

	return nil

}

func deleteNstrafficdomainFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomainFunc")
	client := meta.(*NetScalerNitroClient).client
	nstrafficdomainName := d.Id()
	err := client.DeleteResource(service.Nstrafficdomain.Type(), nstrafficdomainName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
