package citrixadc

import (
	"context"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcLbroute6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbroute6Func,
		ReadContext:   readLbroute6Func,
		DeleteContext: deleteLbroute6Func,
		Schema: map[string]*schema.Schema{
			"gatewayname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbroute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbroute6Func")
	client := meta.(*NetScalerNitroClient).client
	var network = d.Get("network").(string)
	lbroute6 := lb.Lbroute6{
		Gatewayname: d.Get("gatewayname").(string),
		Network:     d.Get("network").(string),
	}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		lbroute6.Td = intPtr(d.Get("td").(int))
	}

	_, err := client.AddResource(service.Lbroute6.Type(), "", &lbroute6)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(network)

	return readLbroute6Func(ctx, d, meta)
}

func readLbroute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbroute6Func")
	client := meta.(*NetScalerNitroClient).client
	network := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading lbroute6 state %s", network)
	findParams := service.FindParams{
		ResourceType: service.Lbroute6.Type(),
	}

	dataArray, err := client.FindResourceArrayWithParams(findParams)

	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lb route6 state %s", network)
		d.SetId("")
		return nil
	}

	if len(dataArray) == 0 {
		log.Printf("[WARN] citrixadc-provider: lb route6 does not exist. Clearing state.")
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, lbroute6 := range dataArray {
		if lbroute6["network"] == network {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams route6 not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lb route6 state %s", network)
		d.SetId("")
		return nil
	}

	data := dataArray[foundIndex]

	d.Set("gatewayname", data["gatewayname"])
	d.Set("network", data["network"])
	setToInt("td", d, data["td"])

	return nil

}

func deleteLbroute6Func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbroute6Func")
	client := meta.(*NetScalerNitroClient).client

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(d.Get("network").(string))

	err := client.DeleteResourceWithArgsMap(service.Lbroute6.Type(), "", argsMap)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
