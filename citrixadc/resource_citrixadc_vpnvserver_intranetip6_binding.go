package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcVpnvserver_intranetip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnvserver_intranetip6_bindingFunc,
		ReadContext:   readVpnvserver_intranetip6_bindingFunc,
		DeleteContext: deleteVpnvserver_intranetip6_bindingFunc,
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
			"intranetip6": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"numaddr": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	intranetip6 := d.Get("intranetip6")
	bindingId := fmt.Sprintf("%s,%s", name, intranetip6)
	vpnvserver_intranetip6_binding := vpn.Vpnvserverintranetip6binding{
		Intranetip6: d.Get("intranetip6").(string),
		Name:        d.Get("name").(string),
	}

	if raw := d.GetRawConfig().GetAttr("numaddr"); !raw.IsNull() {
		vpnvserver_intranetip6_binding.Numaddr = intPtr(d.Get("numaddr").(int))
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_intranetip6_binding.Type(), &vpnvserver_intranetip6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVpnvserver_intranetip6_bindingFunc(ctx, d, meta)
}

func readVpnvserver_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip6 := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_intranetip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_intranetip6_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return diag.FromErr(err)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip6"].(string) == intranetip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("intranetip6", data["intranetip6"])
	d.Set("name", data["name"])
	setToInt("numaddr", d, data["numaddr"])

	return nil

}

func deleteVpnvserver_intranetip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_intranetip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip6 := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip6:%s", intranetip6))
	if val, ok := d.GetOk("numaddr"); ok {
		args = append(args, fmt.Sprintf("numaddr:%d", (val.(int))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnvserver_intranetip6_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
