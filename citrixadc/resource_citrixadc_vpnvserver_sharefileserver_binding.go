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

func resourceCitrixAdcVpnvserver_sharefileserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnvserver_sharefileserver_bindingFunc,
		ReadContext:   readVpnvserver_sharefileserver_bindingFunc,
		DeleteContext: deleteVpnvserver_sharefileserver_bindingFunc,
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
			"sharefile": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	sharefile := d.Get("sharefile")
	bindingId := fmt.Sprintf("%s,%s", name, sharefile)
	vpnvserver_sharefileserver_binding := vpn.Vpnvserversharefileserverbinding{
		Name:      d.Get("name").(string),
		Sharefile: d.Get("sharefile").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_sharefileserver_binding.Type(), &vpnvserver_sharefileserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVpnvserver_sharefileserver_bindingFunc(ctx, d, meta)
}

func readVpnvserver_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sharefile := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_sharefileserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_sharefileserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_sharefileserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["sharefile"].(string) == sharefile {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_sharefileserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("sharefile", data["sharefile"])

	return nil

}

func deleteVpnvserver_sharefileserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_sharefileserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sharefile := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("sharefile:%s", sharefile))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_sharefileserver_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
