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

func resourceCitrixAdcPolicypatset_pattern_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicypatset_pattern_bindingFunc,
		ReadContext:   readPolicypatset_pattern_bindingFunc,
		DeleteContext: deletePolicypatset_pattern_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"charset": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"feature": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicypatset_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	stringText := d.Get("string").(string)
	bindingId := fmt.Sprintf("%v,%v", name, stringText)

	policypatset_pattern_binding := policy.Policypatsetpatternbinding{
		Charset: d.Get("charset").(string),
		Comment: d.Get("comment").(string),
		Feature: d.Get("feature").(string),
		Index:   d.Get("index").(int),
		Name:    d.Get("name").(string),
		String:  d.Get("string").(string),
	}

	err := client.UpdateUnnamedResource(service.Policypatset_pattern_binding.Type(), &policypatset_pattern_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readPolicypatset_pattern_bindingFunc(ctx, d, meta)
}

func readPolicypatset_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	stringText := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading policypatset_pattern_binding state %s", bindingId)
	findParams := service.FindParams{
		ResourceType:             "policypatset_pattern_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 2823,
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
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right String
	foundIndex := -1
	for i, v := range dataArr {
		if v["String"].(string) == stringText {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams String not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("charset", data["charset"])
	d.Set("comment", data["comment"])
	d.Set("feature", data["feature"])
	setToInt("index", d, data["index"])
	d.Set("name", data["name"])
	d.Set("string", data["String"])

	return nil

}

func deletePolicypatset_pattern_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	stringText := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["String"] = url.QueryEscape(stringText)

	err := client.DeleteResourceWithArgsMap(service.Policypatset_pattern_binding.Type(), name, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
