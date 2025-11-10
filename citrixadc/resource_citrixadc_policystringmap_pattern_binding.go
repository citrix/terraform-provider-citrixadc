package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcPolicystringmap_pattern_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicystringmap_pattern_bindingFunc,
		ReadContext:   readPolicystringmap_pattern_bindingFunc,
		UpdateContext: updatePolicystringmap_pattern_bindingFunc,
		DeleteContext: deletePolicystringmap_pattern_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createPolicystringmap_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	key := d.Get("key").(string)
	bindingId := fmt.Sprintf("%s,%s", name, key)
	policystringmap_pattern_binding := policy.Policystringmappatternbinding{
		Key:     d.Get("key").(string),
		Name:    d.Get("name").(string),
		Value:   d.Get("value").(string),
		Comment: d.Get("comment").(string),
	}

	_, err := client.AddResource(service.Policystringmap_pattern_binding.Type(), name, &policystringmap_pattern_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readPolicystringmap_pattern_bindingFunc(ctx, d, meta)
}

func readPolicystringmap_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	key := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading policystringmap_pattern_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "policystringmap_pattern_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing policystringmap_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["key"].(string) == key {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policystringmap_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("key", data["key"])
	d.Set("value", data["value"])
	d.Set("comment", data["comment"])

	return nil

}

func updatePolicystringmap_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	bindingId := d.Id()
	policystringmap_pattern_binding := policy.Policystringmappatternbinding{
		Key:   d.Get("key").(string),
		Name:  name,
		Value: d.Get("value").(string),
	}
	hasUpdate := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG] netscaler-provider:  Comment has changed for policystringmap_pattern_binding %s, starting update", bindingId)
		policystringmap_pattern_binding.Comment = d.Get("comment").(string)
		hasUpdate = true
	}

	if hasUpdate {
		_, err := client.AddResource(service.Policystringmap_pattern_binding.Type(), name, &policystringmap_pattern_binding)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return readPolicystringmap_pattern_bindingFunc(ctx, d, meta)
}

func deletePolicystringmap_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicystringmap_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	key := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("key:%s", url.QueryEscape(key)))

	err := client.DeleteResourceWithArgs("policystringmap_pattern_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
