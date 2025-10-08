package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcAuditsyslogglobal_auditsyslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
		ReadContext:   readAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
		DeleteContext: deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
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
		},
	}
}

func createAuditsyslogglobal_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditsyslogglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	auditsyslogglobal_auditsyslogpolicy_binding := audit.Auditsyslogglobalauditsyslogpolicybinding{
		Globalbindtype: d.Get("globalbindtype").(string),
		Policyname:     d.Get("policyname").(string),
		Priority:       d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource("auditsyslogglobal_auditsyslogpolicy_binding", &auditsyslogglobal_auditsyslogpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(policyname)

	return readAuditsyslogglobal_auditsyslogpolicy_bindingFunc(ctx, d, meta)
}

func readAuditsyslogglobal_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditsyslogglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading auditsyslogglobal_auditsyslogpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "auditsyslogglobal_auditsyslogpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing auditsyslogglobal_auditsyslogpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing auditsyslogglobal_auditsyslogpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("feature", data["feature"])
	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(policyname)))
	if v, ok := d.GetOk("globalbindtype"); ok {
		bind_type := v.(string)
		args = append(args, fmt.Sprintf("globalbindtype:%s", bind_type))
	} else {
		args = append(args, fmt.Sprintf("globalbindtype:SYSTEM_GLOBAL"))
	}
	err := client.DeleteResourceWithArgs("auditsyslogglobal_auditsyslogpolicy_binding", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
