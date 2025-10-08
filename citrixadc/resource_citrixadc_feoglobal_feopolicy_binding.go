package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/feo"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcFeoglobal_feopolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFeoglobal_feopolicy_bindingFunc,
		ReadContext:   readFeoglobal_feopolicy_bindingFunc,
		DeleteContext: deleteFeoglobal_feopolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"globalbindtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createFeoglobal_feopolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeoglobal_feopolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	feoglobal_feopolicy_binding := feo.Feoglobalfeopolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource("feoglobal_feopolicy_binding", &feoglobal_feopolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readFeoglobal_feopolicy_bindingFunc(ctx, d, meta)
}

func readFeoglobal_feopolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeoglobal_feopolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading feoglobal_feopolicy_binding state %s", policyname)

	argsMap := make(map[string]string)
	if v, ok := d.GetOk("type"); ok {
		argsMap["type"] = v.(string)
	} else {
		argsMap["type"] = "REQ_DEFAULT"
	}

	findParams := service.FindParams{
		ResourceType:             "feoglobal_feopolicy_binding",
		ArgsMap:                  argsMap,
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
		log.Printf("[WARN] citrixadc-provider: Clearing feoglobal_feopolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing feoglobal_feopolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteFeoglobal_feopolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeoglobal_feopolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	args = append(args, fmt.Sprintf("priority:%v", d.Get("priority").(int)))
	if v, ok := d.GetOk("type"); ok {
		type_val := v.(string)
		args = append(args, fmt.Sprintf("type:%s", type_val))
	} else {
		args = append(args, fmt.Sprintf("type:%s", "REQ_DEFAULT"))
	}

	err := client.DeleteResourceWithArgs("feoglobal_feopolicy_binding", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
