package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cr"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcCrvserver_lbvserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCrvserver_lbvserver_bindingFunc,
		ReadContext:   readCrvserver_lbvserver_bindingFunc,
		DeleteContext: deleteCrvserver_lbvserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"lbvserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createCrvserver_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCrvserver_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	lbvserver := d.Get("lbvserver")
	bindingId := fmt.Sprintf("%s,%s", name, lbvserver)
	crvserver_lbvserver_binding := cr.Crvserverlbvserverbinding{
		Lbvserver: lbvserver.(string),
		Name:      name.(string),
	}

	err := client.UpdateUnnamedResource(service.Crvserver_lbvserver_binding.Type(), &crvserver_lbvserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readCrvserver_lbvserver_bindingFunc(ctx, d, meta)
}

func readCrvserver_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCrvserver_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	lbvserver := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading crvserver_lbvserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "crvserver_lbvserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing crvserver_lbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["lbvserver"].(string) == lbvserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams lbvserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing crvserver_lbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("lbvserver", data["lbvserver"])
	d.Set("name", data["name"])

	return nil

}

func deleteCrvserver_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrvserver_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	lbvserver := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("lbvserver:%s", lbvserver))

	err := client.DeleteResourceWithArgs(service.Crvserver_lbvserver_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
