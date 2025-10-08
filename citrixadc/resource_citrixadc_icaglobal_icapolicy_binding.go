package citrixadc

import (
	"context"
	"fmt"
	"log"

	"github.com/citrix/adc-nitro-go/resource/config/ica"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcaglobal_icapolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcaglobal_icapolicy_bindingFunc,
		ReadContext:   readIcaglobal_icapolicy_bindingFunc,
		DeleteContext: deleteIcaglobal_icapolicy_bindingFunc,
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
			"type": {
				Type:     schema.TypeString,
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
		},
	}
}

func createIcaglobal_icapolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcaglobal_icapolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	icaglobal_icapolicy_binding := ica.Icaglobalicapolicybinding{
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource("icaglobal_icapolicy_binding", &icaglobal_icapolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readIcaglobal_icapolicy_bindingFunc(ctx, d, meta)
}

func readIcaglobal_icapolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcaglobal_icapolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading icaglobal_icapolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "icaglobal_icapolicy_binding",
		ArgsMap:                  map[string]string{"type": d.Get("type").(string)},
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
		log.Printf("[WARN] citrixadc-provider: Clearing icaglobal_icapolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing icaglobal_icapolicy_binding state %s", policyname)
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

func deleteIcaglobal_icapolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcaglobal_icapolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	args = append(args, fmt.Sprintf("priority:%v", d.Get("priority").(int)))
	args = append(args, fmt.Sprintf("type:%s", d.Get("type").(string)))

	err := client.DeleteResourceWithArgs("icaglobal_icapolicy_binding", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
