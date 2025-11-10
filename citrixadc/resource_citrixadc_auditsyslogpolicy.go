package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"bytes"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuditsyslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuditsyslogpolicyFunc,
		ReadContext:   readAuditsyslogpolicyFunc,
		UpdateContext: updateAuditsyslogpolicyFunc,
		DeleteContext: deleteAuditsyslogpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"globalbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				MaxItems: 1,
				Set:      auditsyslogpolicyGlobalbindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"globalbindtype": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gotopriorityexpression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nextfactor": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createAuditsyslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditsyslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var auditsyslogpolicyName string
	if v, ok := d.GetOk("name"); ok {
		auditsyslogpolicyName = v.(string)
	} else {
		auditsyslogpolicyName = resource.PrefixedUniqueId("tf-auditsyslogpolicy-")
		d.Set("name", auditsyslogpolicyName)
	}
	auditsyslogpolicy := audit.Auditsyslogpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Auditsyslogpolicy.Type(), auditsyslogpolicyName, &auditsyslogpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := syncSystemglobalAuditsyslogpolicyBinding(d, meta); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(auditsyslogpolicyName)

	return readAuditsyslogpolicyFunc(ctx, d, meta)
}

func readAuditsyslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditsyslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditsyslogpolicy state %s", auditsyslogpolicyName)
	data, err := client.FindResource(service.Auditsyslogpolicy.Type(), auditsyslogpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditsyslogpolicy state %s", auditsyslogpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	if err := readSystemglobalAuditsyslogpolicyBinding(d, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil

}

func updateAuditsyslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditsyslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogpolicyName := d.Get("name").(string)

	auditsyslogpolicy := audit.Auditsyslogpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for auditsyslogpolicy %s, starting update", auditsyslogpolicyName)
		auditsyslogpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for auditsyslogpolicy %s, starting update", auditsyslogpolicyName)
		auditsyslogpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for auditsyslogpolicy %s, starting update", auditsyslogpolicyName)
		auditsyslogpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Auditsyslogpolicy.Type(), auditsyslogpolicyName, &auditsyslogpolicy)
		if err != nil {
			return diag.Errorf("Error updating auditsyslogpolicy %s: %s", auditsyslogpolicyName, err.Error())
		}
	}

	if err := syncSystemglobalAuditsyslogpolicyBinding(d, meta); err != nil {
		return diag.FromErr(err)
	}

	return readAuditsyslogpolicyFunc(ctx, d, meta)
}

func deleteAuditsyslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogpolicyName := d.Id()

	// Unbind from global if appropriate
	if v, ok := d.GetOk("globalbinding"); ok {
		// There is only one element
		h := v.(*schema.Set).List()[0]
		err := deleteSystemglobalAuditsyslogpolicyBinding(d, meta, h.(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err := client.DeleteResource(service.Auditsyslogpolicy.Type(), auditsyslogpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSystemglobalAuditsyslogpolicyBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSystemglobalAuditsyslogpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	policyName := d.Get("name").(string)

	findParams := service.FindParams{
		ResourceType: "systemglobal_auditsyslogpolicy_binding",
		FilterMap:    map[string]string{"policyname": url.QueryEscape(policyName)},
	}

	globalBindings, err := client.FindResourceArrayWithParams(findParams)
	log.Printf("[DEBUG]  citrixadc-provider: globalBindings %v", globalBindings)

	if err != nil {
		return err
	}
	processedBindings := make([]interface{}, 0, 1)

	log.Printf("[DEBUG]  citrixadc-provider: processedBindings %v", processedBindings)
	data := make(map[string]interface{})
	if len(globalBindings) > 0 {
		globalBinding := globalBindings[0]
		if v, ok := globalBinding["feature"]; ok {
			data["feature"] = v
		}
		if v, ok := globalBinding["globalbindtype"]; ok {
			data["globalbindtype"] = v
		}
		if v, ok := globalBinding["gotopriorityexpression"]; ok {
			data["gotopriorityexpression"] = v
		}
		if v, ok := globalBinding["nextfactor"]; ok {
			data["nextfactor"] = v
		}
		if v, ok := globalBinding["priority"]; ok {
			v_int, err := strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
			data["priority"] = v_int
		}
		processedBindings = append(processedBindings, data)
	}

	updatedSet := schema.NewSet(auditsyslogpolicyGlobalbindingMappingHash, processedBindings)
	log.Printf("[DEBUG]  citrixadc-provider: updatedSet %v", updatedSet)
	if err := d.Set("globalbinding", updatedSet); err != nil {
		return err
	}

	return nil
}

func syncSystemglobalAuditsyslogpolicyBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In syncSystemglobalAuditsyslogpolicyBinding")
	oldSet, newSet := d.GetChange("globalbinding")

	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v", newSet)

	o := oldSet.(*schema.Set)
	n := newSet.(*schema.Set)

	for _, binding := range o.List() {
		if err := deleteSystemglobalAuditsyslogpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range n.List() {
		if err := addSystemglobalAuditsyslogpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	return nil
}

func addSystemglobalAuditsyslogpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSystemglobalAuditsyslogpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := system.Systemglobalauditsyslogpolicybinding{}
	bindingStruct.Policyname = d.Get("name").(string)

	if d, ok := binding["feature"]; ok {
		bindingStruct.Feature = d.(string)
	}
	if d, ok := binding["globalbindtype"]; ok {
		bindingStruct.Globalbindtype = d.(string)
	}
	if d, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
	}
	if d, ok := binding["nextfactor"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
	}
	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = intPtr(d.(int))
	}

	if err := client.UpdateUnnamedResource("systemglobal_auditsyslogpolicy_binding", bindingStruct); err != nil {
		return err
	}

	return nil
}

func deleteSystemglobalAuditsyslogpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemglobalAuditsyslogpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	policyName := d.Get("name").(string)
	args := make([]string, 0, 1)
	args = append(args, fmt.Sprintf("policyname:%s", policyName))

	if err := client.DeleteResourceWithArgs("systemglobal_auditsyslogpolicy_binding", "", args); err != nil {
		return err
	}

	return nil
}

func auditsyslogpolicyGlobalbindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In auditsyslogpolicyGlobalbindingMappingHash")
	var buf bytes.Buffer

	// All keys added in alphabetical order.
	m := v.(map[string]interface{})
	if d, ok := m["feature"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["globalbindtype"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["gotopriorityexpression"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["nextfactor"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	return hashString(buf.String())
}
