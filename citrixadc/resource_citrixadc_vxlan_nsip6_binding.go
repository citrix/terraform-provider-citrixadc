package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVxlan_nsip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVxlan_nsip6_bindingFunc,
		ReadContext:   readVxlan_nsip6_bindingFunc,
		DeleteContext: deleteVxlan_nsip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vxlanid": {
				Type:     schema.TypeInt,
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
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVxlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVxlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vxlanid := strconv.Itoa(d.Get("vxlanid").(int))
	ipaddress := d.Get("ipaddress").(string)
	bindingId := fmt.Sprintf("%s,%s", vxlanid, ipaddress)
	vxlan_nsip6_binding := network.Vxlannsip6binding{
		Id:        d.Get("vxlanid").(int),
		Ipaddress: d.Get("ipaddress").(string),
		Netmask:   d.Get("netmask").(string),
	}

	err := client.UpdateUnnamedResource(service.Vxlan_nsip6_binding.Type(), &vxlan_nsip6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVxlan_nsip6_bindingFunc(ctx, d, meta)
}

func readVxlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVxlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vxlan_nsip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vxlan_nsip6_binding",
		ResourceName:             vxlanid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ipaddress"].(string) == ipaddress {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vxlan_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("vxlanid", d, data["id"])
	d.Set("ipaddress", data["ipaddress"])
	// d.Set("netmask", data["netmask"])

	return nil

}

func deleteVxlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVxlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vxlanid := idSlice[0]
	ipaddress := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", url.QueryEscape(ipaddress)))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}
	err := client.DeleteResourceWithArgs(service.Vxlan_nsip6_binding.Type(), vxlanid, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
