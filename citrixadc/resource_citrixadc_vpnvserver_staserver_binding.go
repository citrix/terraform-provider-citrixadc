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

func resourceCitrixAdcVpnvserver_staserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnvserver_staserver_bindingFunc,
		ReadContext:   readVpnvserver_staserver_bindingFunc,
		DeleteContext: deleteVpnvserver_staserver_bindingFunc,
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
			"staserver": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"staaddresstype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnvserver_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserver_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	staserver := d.Get("staserver")
	bindingId := fmt.Sprintf("%s,%s", name, staserver)
	vpnvserver_staserver_binding := vpn.Vpnvserverstaserverbinding{
		Name:           d.Get("name").(string),
		Staaddresstype: d.Get("staaddresstype").(string),
		Staserver:      d.Get("staserver").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnvserver_staserver_binding.Type(), &vpnvserver_staserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVpnvserver_staserver_bindingFunc(ctx, d, meta)
}

func readVpnvserver_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserver_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	staserver := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver_staserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_staserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_staserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["staserver"].(string) == staserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver_staserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("staaddresstype", data["staaddresstype"])
	d.Set("staserver", data["staserver"])

	return nil

}

func deleteVpnvserver_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserver_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	staserver := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("staserver:%s", url.QueryEscape(staserver)))

	err := client.DeleteResourceWithArgs(service.Vpnvserver_staserver_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
