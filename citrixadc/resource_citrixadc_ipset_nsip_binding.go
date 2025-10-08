package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcIpset_nsip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIpset_nsip_bindingFunc,
		ReadContext:   readIpset_nsip_bindingFunc,
		DeleteContext: deleteIpset_nsip_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createIpset_nsip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIpset_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	ipaddress := d.Get("ipaddress")
	bindingId := fmt.Sprintf("%s,%s", name, ipaddress)
	ipset_nsip_binding := network.Ipsetnsipbinding{
		Ipaddress: d.Get("ipaddress").(string),
		Name:      d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Ipset_nsip_binding.Type(), &ipset_nsip_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readIpset_nsip_bindingFunc(ctx, d, meta)
}

func readIpset_nsip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIpset_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ipaddress := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading ipset_nsip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "ipset_nsip_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing ipset_nsip_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing ipset_nsip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ipaddress", data["ipaddress"])
	d.Set("name", data["name"])

	return nil

}

func deleteIpset_nsip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIpset_nsip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ipaddress := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ipaddress:%s", ipaddress))

	err := client.DeleteResourceWithArgs(service.Ipset_nsip_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
