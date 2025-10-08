package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcSystemglobal_authenticationpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemglobal_authenticationpolicy_bindingFunc,
		ReadContext:   readSystemglobal_authenticationpolicy_bindingFunc,
		DeleteContext: deleteSystemglobal_authenticationpolicy_bindingFunc,
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
			"builtin": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"feature": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"nextfactor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSystemglobal_authenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	systemglobal_authenticationpolicy_binding := system.Systemglobalauthenticationpolicybinding{
		//Builtin:                d.Get("builtin").([]interface{}),
		//Feature:                d.Get("feature").(string),
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Nextfactor:             d.Get("nextfactor").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Systemglobal_authenticationpolicy_binding.Type(), &systemglobal_authenticationpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readSystemglobal_authenticationpolicy_bindingFunc(ctx, d, meta)
}

func readSystemglobal_authenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading systemglobal_authenticationpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "systemglobal_authenticationpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemglobal_authenticationpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemglobal_authenticationpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("feature", data["feature"])
	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("nextfactor", data["nextfactor"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteSystemglobal_authenticationpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))

	err := client.DeleteResourceWithArgs(service.Systemglobal_authenticationpolicy_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
