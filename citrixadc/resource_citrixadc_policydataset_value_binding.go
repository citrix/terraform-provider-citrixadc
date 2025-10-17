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

func resourceCitrixAdcPolicydataset_value_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPolicydataset_value_bindingFunc,
		ReadContext:   readPolicydataset_value_bindingFunc,
		DeleteContext: deletePolicydataset_value_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
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
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"endrange": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicydataset_value_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	value := d.Get("value").(string)
	// Use `,` as the separator since it is invalid character for servicegroup and monitor name
	valueBindingId := fmt.Sprintf("%s,%s", name, value)

	policydataset_value_binding := policy.Policydatasetvaluebinding{
		Comment:  d.Get("comment").(string),
		Name:     d.Get("name").(string),
		Value:    d.Get("value").(string),
		Endrange: d.Get("endrange").(string),
	}

	if raw := d.GetRawConfig().GetAttr("index"); !raw.IsNull() {
		policydataset_value_binding.Index = intPtr(d.Get("index").(int))
	}

	err := client.UpdateUnnamedResource(service.Policydataset_value_binding.Type(), &policydataset_value_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(valueBindingId)

	return readPolicydataset_value_bindingFunc(ctx, d, meta)
}

func readPolicydataset_value_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	valueBindingId := d.Id()

	idSlice := strings.Split(valueBindingId, ",")

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce value from id string")
	}

	if len(idSlice) > 2 {
		return diag.Errorf("Too many separators \",\" in id string")
	}

	name := idSlice[0]
	value := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_lbmonitor_binding state %s", valueBindingId)
	findParams := service.FindParams{
		ResourceType:             "policydataset_value_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset_value_binding state %s", valueBindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["value"].(string) == value {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset_value_binding state %s", valueBindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	setToInt("index", d, data["index"])
	d.Set("name", data["name"])
	d.Set("value", data["value"])
	d.Set("endrange", data["endrange"])

	return nil

}

func deletePolicydataset_value_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make(map[string]string)
	args["value"] = url.QueryEscape(d.Get("value").(string))
	if val, ok := d.GetOk("endrange"); ok {
		args["endrange"] = url.QueryEscape(val.(string))
	}
	err := client.DeleteResourceWithArgsMap(service.Policydataset_value_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
