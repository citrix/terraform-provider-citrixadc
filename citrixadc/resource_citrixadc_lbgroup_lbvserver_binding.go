package citrixadc

import (
	"context"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcLbgroup_lbvserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbgroup_lbvserver_bindingFunc,
		ReadContext:   readLbgroup_lbvserver_bindingFunc,
		DeleteContext: deleteLbgroup_lbvserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vservername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLbgroup_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbgroup_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbgroupName := d.Get("name").(string)
	lbvserverName := d.Get("vservername").(string)
	bindingId := fmt.Sprintf("%s,%s", lbgroupName, lbvserverName)

	lbgroup_lbvserver_binding := lb.Lbgrouplbvserverbinding{
		Name:        lbgroupName,
		Vservername: lbvserverName,
	}

	_, err := client.AddResource(service.Lbgroup_lbvserver_binding.Type(), lbgroupName, &lbgroup_lbvserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLbgroup_lbvserver_bindingFunc(ctx, d, meta)
}

func readLbgroup_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbgroup_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbgroup_lbvserver_binding bindingId: %s", bindingId)

	idSlice := strings.SplitN(bindingId, ",", 2)
	lbgroupName := idSlice[0]
	lbvserverName := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             service.Lbgroup_lbvserver_binding.Type(),
		ResourceName:             lbgroupName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbgroup_lbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right lbvserver name
	foundIndex := -1
	for i, v := range dataArr {
		if v["vservername"].(string) == lbvserverName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbgroup_lbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("vservername", data["vservername"])

	return nil
}

func deleteLbgroup_lbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbgroup_lbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	lbgroupName := idSlice[0]
	lbvserverName := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["vservername"] = url.QueryEscape(lbvserverName)

	err := client.DeleteResourceWithArgsMap(service.Lbgroup_lbvserver_binding.Type(), lbgroupName, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
