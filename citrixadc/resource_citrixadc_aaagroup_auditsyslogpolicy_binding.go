package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcAaagroup_auditsyslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaagroup_auditsyslogpolicy_bindingFunc,
		ReadContext:   readAaagroup_auditsyslogpolicy_bindingFunc,
		DeleteContext: deleteAaagroup_auditsyslogpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func createAaagroup_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroup_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	policy := d.Get("policy").(string)
	bindingId := fmt.Sprintf("%s,%s", groupname, policy)
	aaagroup_auditsyslogpolicy_binding := aaa.Aaagroupauditsyslogpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupname:              d.Get("groupname").(string),
		Policy:                 d.Get("policy").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaagroup_auditsyslogpolicy_binding.Type(), &aaagroup_auditsyslogpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAaagroup_auditsyslogpolicy_bindingFunc(ctx, d, meta)
}

func readAaagroup_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroup_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	policy := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup_auditsyslogpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaagroup_auditsyslogpolicy_binding",
		ResourceName:             groupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_auditsyslogpolicy_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policy not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_auditsyslogpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("policy", data["policy"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteAaagroup_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroup_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policy := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policy:%s", policy))
	if v, ok := d.GetOk("type"); ok {
		type_val := v.(string)
		args = append(args, fmt.Sprintf("type:%s", type_val))
	}

	err := client.DeleteResourceWithArgs(service.Aaagroup_auditsyslogpolicy_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
