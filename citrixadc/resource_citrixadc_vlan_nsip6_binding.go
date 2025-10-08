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

func resourceCitrixAdcVlan_nsip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVlan_nsip6_bindingFunc,
		ReadContext:   readVlan_nsip6_bindingFunc,
		DeleteContext: deleteVlan_nsip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vlanid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createVlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vlanid := strconv.Itoa(d.Get("vlanid").(int))
	ipaddress := d.Get("ipaddress").(string)
	bindingId := fmt.Sprintf("%s,%s", vlanid, ipaddress)
	vlan_nsip6_binding := network.Vlannsip6binding{
		Id:         d.Get("vlanid").(int),
		Ipaddress:  d.Get("ipaddress").(string),
		Netmask:    d.Get("netmask").(string),
		Ownergroup: d.Get("ownergroup").(string),
		Td:         d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource(service.Vlan_nsip6_binding.Type(), &vlan_nsip6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readVlan_nsip6_bindingFunc(ctx, d, meta)
}

func readVlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vrid := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading vlan_nsip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "vlan_nsip6_binding",
		ResourceName:             vrid,
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
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_nsip6_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ipaddress not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vlan_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("vlanid", d, data["id"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("netmask", data["netmask"])
	d.Set("ownergroup", data["ownergroup"])
	setToInt("td", d, data["td"])

	return nil

}

func deleteVlan_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVlan_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vrid := idSlice[0]
	ipaddress := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", url.PathEscape(ipaddress)))
	if v, ok := d.GetOk("netmask"); ok {
		netmask := v.(string)
		args = append(args, fmt.Sprintf("netmask:%s", netmask))
	}
	if v, ok := d.GetOk("td"); ok {
		td := v.(int)
		args = append(args, fmt.Sprintf("td:%v", td))
	}
	if v, ok := d.GetOk("ownergroup"); ok {
		ownergroup := v.(string)
		args = append(args, fmt.Sprintf("ownergroup:%s", ownergroup))
	}

	err := client.DeleteResourceWithArgs(service.Vlan_nsip6_binding.Type(), vrid, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
