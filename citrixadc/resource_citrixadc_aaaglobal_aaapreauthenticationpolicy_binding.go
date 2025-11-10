package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaaglobal_aaapreauthenticationpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		ReadContext:   readAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		DeleteContext: deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"builtin": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"policy": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAaaglobal_aaapreauthenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Get("policy").(string)
	aaaglobal_aaapreauthenticationpolicy_binding := aaa.Aaaglobalaaapreauthenticationpolicybinding{
		Builtin: toStringList(d.Get("builtin").([]interface{})),
		Policy:  d.Get("policy").(string),
	}

	if raw := d.GetRawConfig().GetAttr("priority"); !raw.IsNull() {
		aaaglobal_aaapreauthenticationpolicy_binding.Priority = intPtr(d.Get("priority").(int))
	}

	err := client.UpdateUnnamedResource(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), &aaaglobal_aaapreauthenticationpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policy)

	return readAaaglobal_aaapreauthenticationpolicy_bindingFunc(ctx, d, meta)
}

func readAaaglobal_aaapreauthenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policy := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)

	findParams := service.FindParams{
		ResourceType:             "aaaglobal_aaapreauthenticationpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policy {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaaglobal_aaapreauthenticationpolicy_binding state %s", policy)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("policy", data["policy"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaglobal_aaapreauthenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policy := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))

	err := client.DeleteResourceWithArgs(service.Aaaglobal_aaapreauthenticationpolicy_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
