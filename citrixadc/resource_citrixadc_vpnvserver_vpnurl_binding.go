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

func resourceCitrixAdcVpnvserver_vpnurl_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnvserver_vpnurl_bindingFunc,
		ReadContext:   readVpnvserver_vpnurl_bindingFunc,
		DeleteContext: deleteVpnvserver_vpnurl_bindingFunc,
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
			"urlname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_vpnurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	urlname := d.Get("urlname")
	bindingId := fmt.Sprintf("%s,%s", name, urlname)
	vpnvserver_vpnurl_binding := vpn.Vpnvservervpnurlbinding{
		Name:    d.Get("name").(string),
		Urlname: d.Get("urlname").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_vpnurl_binding.Type(), &vpnvserver_vpnurl_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVpnvserver_vpnurl_bindingFunc(ctx, d, meta)
}

func readVpnvserver_vpnurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	urlname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_vpnurl_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_vpnurl_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpnurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["urlname"].(string) == urlname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_vpnurl_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("urlname", data["urlname"])

	return nil

}

func deleteVpnvserver_vpnurl_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_vpnurl_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	urlname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("urlname:%s", urlname))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_vpnurl_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
