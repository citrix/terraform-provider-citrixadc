package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/cs"
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/config/rewrite"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcRewritepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewritepolicyFunc,
		Read:          readRewritepolicyFunc,
		Update:        updateRewritepolicyFunc,
		Delete:        deleteRewritepolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"globalbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      rewritepolicyGlobalbindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"globalbindtype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gotopriorityexpression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invoke": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labelname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labeltype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"policyname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lbvserverbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      rewritepolicyLbVserverMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bindpoint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gotopriorityexpression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invoke": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labelname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labeltype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"csvserverbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      rewritepolicyLbVserverMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bindpoint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gotopriorityexpression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invoke": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"labelname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labeltype": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"targetlbvserver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var rewritepolicyName string
	if v, ok := d.GetOk("name"); ok {
		rewritepolicyName = v.(string)
	} else {
		rewritepolicyName = resource.PrefixedUniqueId("tf-rewritepolicy-")
		d.Set("name", rewritepolicyName)
	}
	rewritepolicy := rewrite.Rewritepolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Newname:     d.Get("newname").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(netscaler.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
	if err != nil {
		return err
	}

	d.SetId(rewritepolicyName)

	err = readRewritepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewritepolicy but we can't read it ?? %s", rewritepolicyName)
		return nil
	}

	if err := updateRewriteGlobalBinding(d, meta); err != nil {
		return err
	}

	if err := updateRewriteLbvserverBindings(d, meta); err != nil {
		return err
	}

	if err := updateRewriteCsvserverBindings(d, meta); err != nil {
		return err
	}

	return nil
}

func readRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading rewritepolicy state %s", rewritepolicyName)
	data, err := client.FindResource(netscaler.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing rewritepolicy state %s", rewritepolicyName)
		d.SetId("")
		return nil
	}

	if err := readRewriteGlobalBinding(d, meta); err != nil {
		return err
	}

	if err := readRewriteLbvserverBindings(d, meta); err != nil {
		return err
	}

	if err := readRewriteCsvserverBindings(d, meta); err != nil {
		return err
	}

	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Get("name").(string)

	rewritepolicy := rewrite.Rewritepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for rewritepolicy %s, starting update", rewritepolicyName)
		rewritepolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Rewritepolicy.Type(), rewritepolicyName, &rewritepolicy)
		if err != nil {
			return fmt.Errorf("Error updating rewritepolicy %s", rewritepolicyName)
		}
	}

	if err := updateRewriteGlobalBinding(d, meta); err != nil {
		return err
	}

	if err := updateRewriteLbvserverBindings(d, meta); err != nil {
		return err
	}

	if err := updateRewriteCsvserverBindings(d, meta); err != nil {
		return err
	}

	return readRewritepolicyFunc(d, meta)
}

func deleteRewritepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewritepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	rewritepolicyName := d.Id()

	// Delete all bindings prior to deleting rewrite policy

	if err := deleteRewriteGlobalBinding(d, meta); err != nil {
		return err
	}

	if err := deleteRewriteLbvserverBindings(d, meta); err != nil {
		return err
	}

	if err := deleteRewriteCsvserverBindings(d, meta); err != nil {
		return err
	}

	err := client.DeleteResource(netscaler.Rewritepolicy.Type(), rewritepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func rewritepolicyGlobalbindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In rewritepolicyGlobalbindingMappingHash")
	var buf bytes.Buffer

	// All keys added in alphabetical order.
	m := v.(map[string]interface{})
	if d, ok := m["gotopriorityexpression"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["labelname"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["labeltype"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}
	if d, ok := m["type"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	return hashcode.String(buf.String())
}

func addSingleRewriteGlobalBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleRewriteGlobalBinding")

	client := meta.(*NetScalerNitroClient).client

	bindingStruct := rewrite.Rewriteglobalrewritepolicybinding{}
	bindingStruct.Policyname = d.Get("name").(string)
	if d, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
	}
	if d, ok := binding["invoke"]; ok {
		bindingStruct.Invoke = d.(bool)
	}
	if d, ok := binding["labelname"]; ok {
		log.Printf("Labelname %v\n", d)
		bindingStruct.Labelname = d.(string)
	}
	if d, ok := binding["labeltype"]; ok {
		log.Printf("Labeltype %v\n", d)
		bindingStruct.Labeltype = d.(string)
	}
	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = d.(int)
	}
	if d, ok := binding["type"]; ok {
		log.Printf("Type %v\n", d)
		bindingStruct.Type = d.(string)
	}

	if err := client.UpdateUnnamedResource("rewriteglobal_rewritepolicy_binding", bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleRewriteGlobalBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteSingleGlobalBinding")

	client := meta.(*NetScalerNitroClient).client

	// Construct args from binding data
	args := make([]string, 0, 3)

	if d, ok := d.GetOk("name"); ok {
		s := fmt.Sprintf("policyname:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["type"]; ok {
		s := fmt.Sprintf("type:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["priority"]; ok {
		s := fmt.Sprintf("priority:%d", d.(int))
		args = append(args, s)
	}

	if err := client.DeleteResourceWithArgs("rewriteglobal_rewritepolicy_binding", binding["policyname"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting rewrite global binding %v\n", binding)
		return err
	}
	return nil
}

func readRewriteGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readRewriteGlobalBinding")
	client := meta.(*NetScalerNitroClient).client

	log.Printf("Existing global binding %v \n", d.Get("globalbinding"))
	o, n := d.GetChange("globalbinding")
	log.Printf("o %v n %v\n", o, n)

	name := d.Get("name").(string)
	globalBindings, _ := client.FindResourceArray("rewritepolicy_rewriteglobal_binding", name)
	log.Printf("Global bindings read %v\n", globalBindings)
	log.Printf("Global bindings len %v\n", len(globalBindings))
	// processedBindings will be used to update the Set
	processedBindings := make([]interface{}, len(globalBindings))
	for i, val := range globalBindings {
		processedBindings[i] = make(map[string]interface{})
		boundtoSlice := strings.Split(val["boundto"].(string), " ")
		processedBindings[i].(map[string]interface{})["type"] = boundtoSlice[1]
		if v, ok := val["gotopriorityexpression"]; ok {
			processedBindings[i].(map[string]interface{})["gotopriorityexpression"] = v
		}
		if v, ok := val["labelname"]; ok {
			processedBindings[i].(map[string]interface{})["labelname"] = v
		}
		if v, ok := val["labeltype"]; ok {
			processedBindings[i].(map[string]interface{})["labeltype"] = v
		}
		if v, ok := val["priority"]; ok {
			var err error
			if processedBindings[i].(map[string]interface{})["priority"], err = strconv.Atoi(v.(string)); err != nil {
				return err
			}
		}
		// deduce invoke from the existence of labeltype and labelname
		// since the NITRO object does not contain it explicitely
		processedBindings[i].(map[string]interface{})["invoke"] = false
		if v, ok := processedBindings[i].(map[string]interface{})["labelname"]; ok {
			if v != "" {
				processedBindings[i].(map[string]interface{})["invoke"] = true
			}
		}
		if v, ok := processedBindings[i].(map[string]interface{})["labeltype"]; ok {
			if v != "" {
				processedBindings[i].(map[string]interface{})["invoke"] = true
			}
		}
	}
	log.Printf("processedBindings %v\n", processedBindings)

	updatedSet := schema.NewSet(rewritepolicyGlobalbindingMappingHash, processedBindings)
	log.Printf("global updatedSet %v\n", updatedSet)
	if err := d.Set("globalbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated global binding Set %v\n", d.Get("globalbinding"))
	return nil
}

func updateRewriteGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteGlobalBinding")

	oldSet, newSet := d.GetChange("globalbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleRewriteGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleRewriteGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func deleteRewriteGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteGlobalBinding")
	if bindings, ok := d.GetOk("globalbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleRewriteGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}

func rewritepolicyLbVserverMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In rewritepolicyLbVserverMappingHash")
	var buf bytes.Buffer

	// All keys added in alphabetical order.
	m := v.(map[string]interface{})
	if d, ok := m["bindpoint"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["gotopriorityexpression"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["invoke"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", d.(bool)))
	}
	if d, ok := m["labelname"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["labeltype"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}
	return hashcode.String(buf.String())
}

func readRewriteLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readRewriteLbvserverBindings")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	// Read the lb vserver bindings registered under this policy name
	lbVserverBindings, _ := client.FindResourceArray("rewritepolicy_lbvserver_binding", name)
	log.Printf("lbVserverBindings %v\n", lbVserverBindings)

	// Process values into new list of maps
	processedBindings := make([]interface{}, len(lbVserverBindings))
	// Initialize maps
	for i, _ := range processedBindings {
		processedBindings[i] = make(map[string]interface{})
	}

	for i, a := range lbVserverBindings {
		//acls[i] = a.(map[string]interface{})
		log.Printf("rewrite lbvserver binding key %v value %v\n", i, a)

		// Process boundto key to deduce lbvserver and bindpoint
		boundtoSlice := strings.Split(a["boundto"].(string), " ")
		log.Printf("boundtoSlice %v\n", boundtoSlice)
		if boundtoSlice[0] == "REQ" {
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "RES" {
			processedBindings[i].(map[string]interface{})["bindpoint"] = "RESPONSE"
		} else {
			return fmt.Errorf("Unexpected bindpoint string \"%v\"", boundtoSlice[0])
		}

		processedBindings[i].(map[string]interface{})["name"] = boundtoSlice[2]
		//processedBindings[i].(map[string]interface{})["policyname"] = d.Get("name").(string)

	}

	log.Printf("processedBindings %v\n", processedBindings)

	// Parse bindings from the lb vserver side to complete processedBindings
	for _, binding := range processedBindings {
		lbVserverName := binding.(map[string]interface{})["name"].(string)
		lbVserverPolicyBindings, _ := client.FindResourceArray("lbvserver_rewritepolicy_binding", lbVserverName)
		// Find the one binding that refers to this particular policy
		var relevantBinding map[string]interface{}
		for _, b := range lbVserverPolicyBindings {
			if b["policyname"] == d.Get("name").(string) {
				relevantBinding = b
				break
			}
		}
		log.Printf("[DEBUG] citrixadc-provider: Relevant binding %v\n", relevantBinding)
		// Fill in the rest of the fields
		keys := []string{"gotopriorityexpression", "invoke", "labelname", "labeltype", "priority"}
		for _, key := range keys {
			if v, ok := relevantBinding[key]; ok {
				if key == "priority" {
					var err error
					if binding.(map[string]interface{})[key], err = strconv.Atoi(v.(string)); err != nil {
						return err
					}
				} else {
					binding.(map[string]interface{})[key] = v
				}
			}
		}
	}

	log.Printf("processedBindings %v\n", processedBindings)

	updatedSet := schema.NewSet(rewritepolicyLbVserverMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("lbvserverbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated binding Set %v\n", d.Get("lbvserverbinding"))
	return nil
}

func updateRewriteLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteLbvserverBindings")
	oldSet, newSet := d.GetChange("lbvserverbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleRewriteLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleRewriteLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func deleteSingleRewriteLbvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleRewriteLbvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	// Construct args from binding data
	args := make([]string, 0, 3)

	if d, ok := d.GetOk("name"); ok {
		s := fmt.Sprintf("policyname:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["bindpoint"]; ok {
		s := fmt.Sprintf("bindpoint:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["priority"]; ok {
		s := fmt.Sprintf("priority:%d", d.(int))
		args = append(args, s)
	}

	if err := client.DeleteResourceWithArgs("lbvserver_rewritepolicy_binding", binding["name"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting lb vserver binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleRewriteLbvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleRewriteLbvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := lb.Lbvserverrewritepolicybinding{}
	bindingStruct.Policyname = d.Get("name").(string)

	if d, ok := binding["bindpoint"]; ok {
		bindingStruct.Bindpoint = d.(string)
	}
	if d, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
	}
	if d, ok := binding["invoke"]; ok {
		bindingStruct.Invoke = d.(bool)
	}
	if d, ok := binding["labelname"]; ok {
		bindingStruct.Labelname = d.(string)
	}

	if d, ok := binding["labeltype"]; ok {
		bindingStruct.Labeltype = d.(string)
	}
	if d, ok := binding["name"]; ok {
		bindingStruct.Name = d.(string)
	}
	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = d.(int)
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("lbvserver_rewritepolicy_binding", binding["name"].(string), bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteRewriteLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteLbvserverBindings")
	if bindings, ok := d.GetOk("lbvserverbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleRewriteLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}

func readRewriteCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readRewriteCsvserverBindings")

	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	// Read the lb vserver bindings registered under this policy name
	csVserverBindings, _ := client.FindResourceArray("rewritepolicy_csvserver_binding", name)
	log.Printf("csVserverBindings %v\n", csVserverBindings)

	// Process values into new list of maps
	processedBindings := make([]interface{}, len(csVserverBindings))
	// Initialize maps
	for i, _ := range processedBindings {
		processedBindings[i] = make(map[string]interface{})
	}

	for i, a := range csVserverBindings {
		//acls[i] = a.(map[string]interface{})
		log.Printf("rewrite csvserver binding key %v value %v\n", i, a)

		// Process boundto key to deduce lbvserver and bindpoint
		boundtoSlice := strings.Split(a["boundto"].(string), " ")
		log.Printf("boundtoSlice %v\n", boundtoSlice)
		if boundtoSlice[0] == "REQ" {
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "RES" {
			processedBindings[i].(map[string]interface{})["bindpoint"] = "RESPONSE"
		} else {
			return fmt.Errorf("Unexpected bindpoint string \"%v\"", boundtoSlice[0])
		}

		processedBindings[i].(map[string]interface{})["name"] = boundtoSlice[2]
		//processedBindings[i].(map[string]interface{})["policyname"] = d.Get("name").(string)

	}

	log.Printf("processedBindings %v\n", processedBindings)

	// Parse bindings from the lb vserver side to complete processedBindings
	for _, binding := range processedBindings {
		lbVserverName := binding.(map[string]interface{})["name"].(string)
		lbVserverPolicyBindings, _ := client.FindResourceArray("csvserver_rewritepolicy_binding", lbVserverName)
		// Find the one binding that refers to this particular policy
		var relevantBinding map[string]interface{}
		for _, b := range lbVserverPolicyBindings {
			if b["policyname"] == d.Get("name").(string) {
				relevantBinding = b
				break
			}
		}
		log.Printf("[DEBUG] citrixadc-provider: Relevant binding %v\n", relevantBinding)
		// Fill in the rest of the fields
		keys := []string{"gotopriorityexpression", "invoke", "labelname", "labeltype", "priority"}
		for _, key := range keys {
			if v, ok := relevantBinding[key]; ok {
				if key == "priority" {
					var err error
					if binding.(map[string]interface{})[key], err = strconv.Atoi(v.(string)); err != nil {
						return err
					}
				} else {
					binding.(map[string]interface{})[key] = v
				}
			}
		}
	}

	log.Printf("processedBindings %v\n", processedBindings)

	updatedSet := schema.NewSet(rewritepolicyLbVserverMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("csvserverbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated binding Set %v\n", d.Get("csvserverbinding"))
	return nil
}

func updateRewriteCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateRewriteCsvserverBindings")
	oldSet, newSet := d.GetChange("csvserverbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleCsvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleCsvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func addSingleRewriteCsvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleRewriteCsvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := cs.Csvserverrewritepolicybinding{}
	bindingStruct.Policyname = d.Get("name").(string)

	if d, ok := binding["bindpoint"]; ok {
		bindingStruct.Bindpoint = d.(string)
	}
	if d, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
	}
	if d, ok := binding["invoke"]; ok {
		bindingStruct.Invoke = d.(bool)
	}
	if d, ok := binding["labelname"]; ok {
		bindingStruct.Labelname = d.(string)
	}

	if d, ok := binding["labeltype"]; ok {
		bindingStruct.Labeltype = d.(string)
	}
	if d, ok := binding["name"]; ok {
		bindingStruct.Name = d.(string)
	}
	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = d.(int)
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("csvserver_rewritepolicy_binding", binding["name"].(string), bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleRewriteCsvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleRewriteCsvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	// Construct args from binding data
	args := make([]string, 0, 3)

	if d, ok := d.GetOk("name"); ok {
		s := fmt.Sprintf("policyname:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["bindpoint"]; ok {
		s := fmt.Sprintf("bindpoint:%v", d.(string))
		args = append(args, s)
	}

	if d, ok := binding["priority"]; ok {
		s := fmt.Sprintf("priority:%d", d.(int))
		args = append(args, s)
	}

	if err := client.DeleteResourceWithArgs("csvserver_rewritepolicy_binding", binding["name"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting cs vserver binding %v\n", binding)
		return err
	}

	return nil
}

func deleteRewriteCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteCsvserverBindings")
	if bindings, ok := d.GetOk("csvserverbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleCsvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}
