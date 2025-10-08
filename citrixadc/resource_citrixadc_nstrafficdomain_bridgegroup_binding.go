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

func resourceCitrixAdcNstrafficdomain_bridgegroup_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstrafficdomain_bridgegroup_bindingFunc,
		ReadContext:   readNstrafficdomain_bridgegroup_bindingFunc,
		DeleteContext: deleteNstrafficdomain_bridgegroup_bindingFunc,
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
			"bridgegroup": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNstrafficdomain_bridgegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstrafficdomain_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	td := strconv.Itoa(d.Get("td").(int))
	bridgegroup := strconv.Itoa(d.Get("bridgegroup").(int))
	bindingId := fmt.Sprintf("%s,%s", td, bridgegroup)
	nstrafficdomain_bridgegroup_binding := ns.Nstrafficdomainbridgegroupbinding{
		Bridgegroup: d.Get("bridgegroup").(int),
		Td:          d.Get("td").(int),
	}

	err := client.UpdateUnnamedResource(service.Nstrafficdomain_bridgegroup_binding.Type(), &nstrafficdomain_bridgegroup_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readNstrafficdomain_bridgegroup_bindingFunc(ctx, d, meta)
}

func readNstrafficdomain_bridgegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstrafficdomain_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	td := idSlice[0]
	bridgegroup := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nstrafficdomain_bridgegroup_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nstrafficdomain_bridgegroup_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_bridgegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["bridgegroup"].(string) == bridgegroup {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nstrafficdomain_bridgegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("bridgegroup", d, data["bridgegroup"])
	setToInt("td", d, data["td"])

	return nil

}

func deleteNstrafficdomain_bridgegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstrafficdomain_bridgegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	bridgegroup := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("bridgegroup:%s", bridgegroup))

	err := client.DeleteResourceWithArgs(service.Nstrafficdomain_bridgegroup_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
