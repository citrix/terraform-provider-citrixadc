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

func resourceCitrixAdcRnat6_nsip6_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createRnat6_nsip6_bindingFunc,
		ReadContext:   readRnat6_nsip6_bindingFunc,
		DeleteContext: deleteRnat6_nsip6_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"natip6": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ownergroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createRnat6_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createRnat6_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	natip6 := d.Get("natip6").(string)
	bindingId := fmt.Sprintf("%s,%s", name, natip6)
	rnat6_nsip6_binding := network.Rnat6nsip6binding{
		Name:       d.Get("name").(string),
		Natip6:     d.Get("natip6").(string),
		Ownergroup: d.Get("ownergroup").(string),
	}

	err := client.UpdateUnnamedResource(service.Rnat6_nsip6_binding.Type(), &rnat6_nsip6_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readRnat6_nsip6_bindingFunc(ctx, d, meta)
}

func readRnat6_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readRnat6_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natip6 := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading rnat6_nsip6_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "rnat6_nsip6_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing rnat6_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["natip6"].(string) == natip6 {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams natip6 not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing rnat6_nsip6_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("natip6", data["natip6"])
	d.Set("ownergroup", data["ownergroup"])

	return nil

}

func deleteRnat6_nsip6_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRnat6_nsip6_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natip6 := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("natip6:%s", natip6))
	if v, ok := d.GetOk("ownergroup"); ok {
		ownergroup := v.(string)
		args = append(args, fmt.Sprintf("ownergroup:%s", ownergroup))
	}

	err := client.DeleteResourceWithArgs(service.Rnat6_nsip6_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
