package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcClusternodegroup_gslbsite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusternodegroup_gslbsite_bindingFunc,
		ReadContext:   readClusternodegroup_gslbsite_bindingFunc,
		DeleteContext: deleteClusternodegroup_gslbsite_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"gslbsite": {
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

func createClusternodegroup_gslbsite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	gslbsite := d.Get("gslbsite")
	bindingId := fmt.Sprintf("%s,%s", name, gslbsite)
	clusternodegroup_gslbsite_binding := cluster.Clusternodegroupgslbsitebinding{
		Gslbsite: d.Get("gslbsite").(string),
		Name:     d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_gslbsite_binding.Type(), &clusternodegroup_gslbsite_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readClusternodegroup_gslbsite_bindingFunc(ctx, d, meta)
}

func readClusternodegroup_gslbsite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	gslbsite := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_gslbsite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_gslbsite_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_gslbsite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["gslbsite"].(string) == gslbsite {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams gslbsite not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_gslbsite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gslbsite", data["gslbsite"])
	d.Set("name", data["name"])

	return nil

}

func deleteClusternodegroup_gslbsite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	gslbsite := idSlice[1]

	args := make([]string, 0)

	args = append(args, fmt.Sprintf("gslbsite:%s", url.QueryEscape(gslbsite)))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_gslbsite_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
