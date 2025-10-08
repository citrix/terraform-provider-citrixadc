package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcVpnvserver_intranetip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnvserver_intranetip_bindingFunc,
		ReadContext:   readVpnvserver_intranetip_bindingFunc,
		DeleteContext: deleteVpnvserver_intranetip_bindingFunc,
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
			"intranetip": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	intranetip := d.Get("intranetip")
	bindingId := fmt.Sprintf("%s,%s", name, intranetip)
	vpnvserver_intranetip_binding := vpn.Vpnvserverintranetipbinding{
		Intranetip: d.Get("intranetip").(string),
		Name:       d.Get("name").(string),
		Netmask:    d.Get("netmask").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_intranetip_binding.Type(), &vpnvserver_intranetip_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVpnvserver_intranetip_bindingFunc(ctx, d, meta)
}

func readVpnvserver_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_intranetip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_intranetip_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip"].(string) == intranetip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("intranetip", data["intranetip"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])

	return nil

}

func deleteVpnvserver_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip:%s", intranetip))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Vpnvserver_intranetip_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
