package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/tm"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcTmglobal_auditnslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmglobal_auditnslogpolicy_bindingFunc,
		ReadContext:   readTmglobal_auditnslogpolicy_bindingFunc,
		DeleteContext: deleteTmglobal_auditnslogpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
		},
	}
}

func createTmglobal_auditnslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	tmglobal_auditnslogpolicy_binding := tm.Tmglobalauditnslogpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Tmglobal_auditnslogpolicy_binding.Type(), &tmglobal_auditnslogpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readTmglobal_auditnslogpolicy_bindingFunc(ctx, d, meta)
}

func readTmglobal_auditnslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading tmglobal_auditnslogpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "tmglobal_auditnslogpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing tmglobal_auditnslogpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing tmglobal_auditnslogpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteTmglobal_auditnslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))

	err := client.DeleteResourceWithArgs(service.Tmglobal_auditnslogpolicy_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
