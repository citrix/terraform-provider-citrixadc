package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcResponderpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createResponderpolicyFunc,
		Read:          readResponderpolicyFunc,
		Update:        updateResponderpolicyFunc,
		Delete:        deleteResponderpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowaction": &schema.Schema{
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
				Set:      globalbindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gotopriorityexpression": &schema.Schema{
							Type:     schema.TypeString,
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
					},
				},
			},
			"lbvserverbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      lbVserverMappingHash,
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

func createResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	var responderpolicyName string
	if v, ok := d.GetOk("name"); ok {
		responderpolicyName = v.(string)
	} else {
		responderpolicyName = resource.PrefixedUniqueId("tf-responderpolicy-")
		d.Set("name", responderpolicyName)
	}
	responderpolicy := responder.Responderpolicy{
		Action:        d.Get("action").(string),
		Appflowaction: d.Get("appflowaction").(string),
		Comment:       d.Get("comment").(string),
		Logaction:     d.Get("logaction").(string),
		Name:          d.Get("name").(string),
		Rule:          d.Get("rule").(string),
		Undefaction:   d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
	if err != nil {
		return err
	}

	if err := updateLbvserverBindings(d, meta); err != nil {
		return err
	}

	if err := updateGlobalBinding(d, meta); err != nil {
		return err
	}

	if err := updateCsvserverBindings(d, meta); err != nil {
		return err
	}

	d.SetId(responderpolicyName)

	err = readResponderpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this responderpolicy but we can't read it ?? %s", responderpolicyName)
		return nil
	}
	return nil
}

func readResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading responderpolicy state %s", responderpolicyName)
	data, err := client.FindResource(service.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing responderpolicy state %s", responderpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("appflowaction", data["appflowaction"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	if _, ok := d.GetOk("globalbinding"); ok {
		if err := readGlobalBinding(d, meta); err != nil {
			return err
		}
	}
	if _, ok := d.GetOk("lbvserverbinding"); ok {
		if err := readLbvserverBindings(d, meta); err != nil {
			return err
		}
	}
	if _, ok := d.GetOk("csvserverbinding"); ok {
		if err := readCsvserverBindings(d, meta); err != nil {
			return err
		}
	}
	return nil

}

func updateResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	responderpolicyName := d.Get("name").(string)

	responderpolicy := responder.Responderpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("appflowaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Appflowaction = d.Get("appflowaction").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for responderpolicy %s, starting update", responderpolicyName)
		responderpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Responderpolicy.Type(), responderpolicyName, &responderpolicy)
		if err != nil {
			return fmt.Errorf("Error updating responderpolicy %s", responderpolicyName)
		}
	}

	if err := updateGlobalBinding(d, meta); err != nil {
		return err
	}

	if d.HasChange("lbvserverbinding") {
		if err := updateLbvserverBindings(d, meta); err != nil {
			return err
		}
	}

	if d.HasChange("csvserverbinding") {
		if err := updateCsvserverBindings(d, meta); err != nil {
			return err
		}
	}
	return readResponderpolicyFunc(d, meta)
}

func deleteResponderpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteResponderpolicyFunc")
	client := meta.(*NetScalerNitroClient).client

	// Delete bindings prior to deleting policy
	if err := deleteGlobalBinding(d, meta); err != nil {
		return err
	}
	if err := deleteLbvserverBindings(d, meta); err != nil {
		return err
	}
	if err := deleteCsvserverBindings(d, meta); err != nil {
		return err
	}

	// Delete policy
	responderpolicyName := d.Id()
	err := client.DeleteResource(service.Responderpolicy.Type(), responderpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func globalbindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In globalbindingMappingHash")
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

func addSingleGlobalBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleGlobalBinding")

	client := meta.(*NetScalerNitroClient).client

	bindingStruct := responder.Responderpolicyglobalbinding{}
	bindingStruct.Name = d.Get("name").(string)
	if d, ok := binding["gotopriorityexpression"]; ok {
		bindingStruct.Gotopriorityexpression = d.(string)
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
		bindingStruct.Priority = uint32(d.(int))
	}

	if err := client.UpdateUnnamedResource("responderglobal_responderpolicy_binding", bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleGlobalBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleGlobalBinding")

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

	if err := client.DeleteResourceWithArgs("responderglobal_responderpolicy_binding", binding["policyname"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting reposnder global binding %v\n", binding)
		return err
	}
	return nil
}

func readGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readGlobalBinding")
	client := meta.(*NetScalerNitroClient).client

	log.Printf("Existing global binding %v \n", d.Get("globalbinding"))
	o, n := d.GetChange("globalbinding")
	log.Printf("o %v n %v\n", o, n)

	name := d.Get("name").(string)
	globalBindings, _ := client.FindResourceArray("responderpolicy_responderglobal_binding", name)
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
			// Need to convert values to corresponding responderglobal_responderpolicy_binding values
			log.Printf("labeltype read %v\n", v)
			if v == "reqvserver" || v == "resvserver" {
				// Standalone
				processedBindings[i].(map[string]interface{})["labeltype"] = "vserver"
			} else if v == "" {
				// Cluster
				processedBindings[i].(map[string]interface{})["labeltype"] = "vserver"
			} else {
				processedBindings[i].(map[string]interface{})["labeltype"] = v
			}
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

	updatedSet := schema.NewSet(globalbindingMappingHash, processedBindings)
	log.Printf("global updatedSet %v\n", updatedSet)
	if err := d.Set("globalbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated global binding Set %v\n", d.Get("globalbinding"))
	return nil
}

func updateGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateGlobalBinding")

	oldSet, newSet := d.GetChange("globalbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func deleteGlobalBinding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGlobalBinding")
	if bindings, ok := d.GetOk("globalbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleGlobalBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}

func lbVserverMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In lbVserverMappingHash")
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

func readLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readLbvserverBindings")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	// Read the lb vserver bindings registered under this policy name
	lbVserverBindings, _ := client.FindResourceArray("responderpolicy_lbvserver_binding", name)
	log.Printf("lbVserverBindings %v\n", lbVserverBindings)

	// Process values into new list of maps
	processedBindings := make([]interface{}, len(lbVserverBindings))
	// Initialize maps
	for i, _ := range processedBindings {
		processedBindings[i] = make(map[string]interface{})
	}

	for i, a := range lbVserverBindings {
		//acls[i] = a.(map[string]interface{})
		log.Printf("responder lbvserver binding key %v value %v\n", i, a)

		// Process boundto key to deduce lbvserver and bindpoint
		boundtoSlice := strings.Split(a["boundto"].(string), " ")
		log.Printf("boundtoSlice %v\n", boundtoSlice)
		if boundtoSlice[0] == "REQ" {
			// Standalone
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "REQUEST" {
			// Cluster
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "RES" {
			// Standalone
			processedBindings[i].(map[string]interface{})["bindpoint"] = "RESPONSE"
		} else if boundtoSlice[0] == "RESPONSE" {
			// Cluster
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
		lbVserverPolicyBindings, _ := client.FindResourceArray("lbvserver_responderpolicy_binding", lbVserverName)
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

	updatedSet := schema.NewSet(lbVserverMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("lbvserverbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated binding Set %v\n", d.Get("lbvserverbinding"))
	return nil
}

func updateLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbvserverBindings")
	oldSet, newSet := d.GetChange("lbvserverbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func deleteSingleLbvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleLbvserverBinding")
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

	if err := client.DeleteResourceWithArgs("lbvserver_responderpolicy_binding", binding["name"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting lb vserver binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleLbvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleLbvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := lb.Lbvserverpolicybinding{}
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
		bindingStruct.Priority = uint32(d.(int))
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("lbvserver_responderpolicy_binding", binding["name"].(string), bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteLbvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserverBindings")
	if bindings, ok := d.GetOk("lbvserverbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleLbvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}

func readCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readCsvserverBindings")

	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	// Read the lb vserver bindings registered under this policy name
	csVserverBindings, _ := client.FindResourceArray("responderpolicy_csvserver_binding", name)
	log.Printf("csVserverBindings %v\n", csVserverBindings)

	// Process values into new list of maps
	processedBindings := make([]interface{}, len(csVserverBindings))
	// Initialize maps
	for i, _ := range processedBindings {
		processedBindings[i] = make(map[string]interface{})
	}

	for i, a := range csVserverBindings {
		//acls[i] = a.(map[string]interface{})
		log.Printf("responder lbvserver binding key %v value %v\n", i, a)

		// Process boundto key to deduce lbvserver and bindpoint
		boundtoSlice := strings.Split(a["boundto"].(string), " ")
		log.Printf("boundtoSlice %v\n", boundtoSlice)
		if boundtoSlice[0] == "REQ" {
			// Standalone
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "REQUEST" {
			// Cluster
			processedBindings[i].(map[string]interface{})["bindpoint"] = "REQUEST"
		} else if boundtoSlice[0] == "RES" {
			// Standalone
			processedBindings[i].(map[string]interface{})["bindpoint"] = "RESPONSE"
		} else if boundtoSlice[0] == "RESPONSE" {
			// Cluster
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
		lbVserverPolicyBindings, _ := client.FindResourceArray("csvserver_responderpolicy_binding", lbVserverName)
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

	updatedSet := schema.NewSet(lbVserverMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("csvserverbinding", updatedSet); err != nil {
		return err
	}
	log.Printf("Updated binding Set %v\n", d.Get("csvserverbinding"))
	return nil
}

func updateCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCsvserverBindings")
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

func addSingleCsvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleCsvserverBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := cs.Csvserverpolicybinding{}
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
		bindingStruct.Priority = uint32(d.(int))
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("csvserver_responderpolicy_binding", binding["name"].(string), bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleCsvserverBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleCsvserverBinding")
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

	if err := client.DeleteResourceWithArgs("csvserver_responderpolicy_binding", binding["name"].(string), args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting cs vserver binding %v\n", binding)
		return err
	}

	return nil
}

func deleteCsvserverBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserverBindings")
	if bindings, ok := d.GetOk("csvserverbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleCsvserverBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}
