package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcNstrafficdomain_vxlan_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstrafficdomain_vxlan_bindingFunc,
		ReadContext:   readNstrafficdomain_vxlan_bindingFunc,
		DeleteContext: deleteNstrafficdomain_vxlan_bindingFunc,
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
			"vxlan": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNstrafficdomain_vxlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	td := strconv.Itoa(d.Get("td").(int))
	vxlan := strconv.Itoa(d.Get("vxlan").(int))
	bindingId := fmt.Sprintf("%s,%s", td, vxlan)
	nstrafficdomain_vxlan_binding := ns.Nstrafficdomainvxlanbinding{}

	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		nstrafficdomain_vxlan_binding.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("vxlan"); !raw.IsNull() {
		nstrafficdomain_vxlan_binding.Vxlan = intPtr(d.Get("vxlan").(int))
	}

	err := client.UpdateUnnamedResource(service.Nstrafficdomain_vxlan_binding.Type(), &nstrafficdomain_vxlan_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readNstrafficdomain_vxlan_bindingFunc(ctx, d, meta)
}

func readNstrafficdomain_vxlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	td := idSlice[0]
	vxlan := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain_vxlan_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nstrafficdomain_vxlan_binding",
		ResourceName:             td,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vxlan"].(string) == vxlan {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_vxlan_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("td", d, data["td"])
	setToInt("vxlan", d, data["vxlan"])

	return nil

}

func deleteNstrafficdomain_vxlan_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomain_vxlan_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vxlan := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vxlan:%s", vxlan))

	err := client.DeleteResourceWithArgs(service.Nstrafficdomain_vxlan_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
