package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"bytes"
	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcSystemgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemgroupFunc,
		Read:          readSystemgroupFunc,
		Update:        updateSystemgroupFunc,
		Delete:        deleteSystemgroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"promptstring": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowedmanagementinterface": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"systemusers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cmdpolicybinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      systemgroupCmdpolicybindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policyname": {
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

func createSystemgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	systemgroupName := d.Get("groupname").(string)

	systemgroup := system.Systemgroup{
		Groupname:                  d.Get("groupname").(string),
		Promptstring:               d.Get("promptstring").(string),
		Timeout:                    d.Get("timeout").(int),
		Allowedmanagementinterface: toStringList(d.Get("allowedmanagementinterface").([]interface{})),
	}

	_, err := client.AddResource(service.Systemgroup.Type(), systemgroupName, &systemgroup)
	if err != nil {
		return err
	}

	d.SetId(systemgroupName)

	// Ignore bindings unless there is an explicit configuration for it
	if _, ok := d.GetOk("cmdpolicybinding"); ok {
		err = updateSystemgroupCmdpolicyBindings(d, meta)
		if err != nil {
			return err
		}
	}

	// Ignore bindings unless there is an explicit configuration for it
	if _, ok := d.GetOk("systemusers"); ok {
		err = updateSystemgroupSystemuserBindings(d, meta)
		if err != nil {
			return err
		}
	}

	err = readSystemgroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemgroup but we can't read it ?? %s", systemgroupName)
		return nil
	}
	return nil
}

func readSystemgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	systemgroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemgroup state %s", systemgroupName)
	data, err := client.FindResource(service.Systemgroup.Type(), systemgroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemgroup state %s", systemgroupName)
		d.SetId("")
		return nil
	}
	if _, ok := d.GetOk("cmdpolicybinding"); ok {
		err = readSystemgroupCmdpolicybindings(d, meta)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("cmdpolicybinding"); ok {
		err = readSystemgroupSystemuserbindings(d, meta)
		if err != nil {
			return err
		}
	}

	d.Set("name", data["name"])
	d.Set("groupname", data["groupname"])
	d.Set("promptstring", data["promptstring"])
	d.Set("timeout", data["timeout"])
	d.Set("allowedmanagementinterface", data["allowedmanagementinterface"])

	return nil

}

func updateSystemgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	systemgroupName := d.Get("groupname").(string)

	systemgroup := system.Systemgroup{
		Groupname: d.Get("groupname").(string),
	}
	hasChange := false
	if d.HasChange("promptstring") {
		log.Printf("[DEBUG]  citrixadc-provider: Promptstring has changed for systemgroup %s, starting update", systemgroupName)
		systemgroup.Promptstring = d.Get("promptstring").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for systemgroup %s, starting update", systemgroupName)
		systemgroup.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("allowedmanagementinterface") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for systemgroup %s, starting update", systemgroupName)
		systemgroup.Allowedmanagementinterface = toStringList(d.Get("allowedmanagementinterface").([]interface{}))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Systemgroup.Type(), systemgroupName, &systemgroup)
		if err != nil {
			return fmt.Errorf("Error updating systemgroup %s", systemgroupName)
		}
	}
	if d.HasChange("cmdpolicybinding") {
		err := updateSystemgroupCmdpolicyBindings(d, meta)
		if err != nil {
			return err
		}
	}

	if d.HasChange("systemusers") {
		err := updateSystemgroupSystemuserBindings(d, meta)
		if err != nil {
			return err
		}
	}
	return readSystemgroupFunc(d, meta)
}

func deleteSystemgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	systemgroupName := d.Id()
	err := client.DeleteResource(service.Systemgroup.Type(), systemgroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteSingleSystemgroupCmdpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSystemgroupCmdpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	groupname := d.Get("groupname").(string)
	// Construct args from binding data
	args := make([]string, 0, 1)

	if d, ok := binding["policyname"]; ok {
		s := fmt.Sprintf("policyname:%v", d.(string))
		args = append(args, s)
	}

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs("systemgroup_systemcmdpolicy_binding", groupname, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting cmdpolicy binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleSystemgroupCmdpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSystemgroupCmdpolicyBinding")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("Adding binding %v", binding)

	bindingStruct := system.Systemgroupcmdpolicybinding{}
	bindingStruct.Groupname = d.Get("groupname").(string)

	if d, ok := binding["policyname"]; ok {
		bindingStruct.Policyname = d.(string)
	}

	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = uint32(d.(int))
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("systemgroup_systemcmdpolicy_binding", bindingStruct.Groupname, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateSystemgroupCmdpolicyBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemgroupCmdpolicyBindings")
	oldSet, newSet := d.GetChange("cmdpolicybinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleSystemgroupCmdpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleSystemgroupCmdpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func systemgroupCmdpolicybindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In systemgroupCmdpolicybindingMappingHash")
	var buf bytes.Buffer

	m := v.(map[string]interface{})
	if d, ok := m["policyname"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["priority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}
	return hashcode.String(buf.String())
}

func readSystemgroupCmdpolicybindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSystemgroupCmdpolicybindings")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	bindings, _ := client.FindResourceArray("systemgroup_systemcmdpolicy_binding", groupname)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i := range bindings {
		processedBindings[i] = make(map[string]interface{})
		processedBindings[i].(map[string]interface{})["policyname"] = bindings[i]["policyname"].(string)
		processedBindings[i].(map[string]interface{})["priority"], _ = strconv.Atoi(bindings[i]["priority"].(string))
	}

	updatedSet := schema.NewSet(systemgroupCmdpolicybindingMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("cmdpolicybinding", updatedSet); err != nil {
		return err
	}
	return nil
}

func deleteSingleSystemgroupSystemuserBinding(d *schema.ResourceData, meta interface{}, username string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSystemgroupSystemuserBinding")
	client := meta.(*NetScalerNitroClient).client

	groupname := d.Get("groupname").(string)
	args := make([]string, 0, 1)

	s := fmt.Sprintf("username:%s", username)
	args = append(args, s)

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs("systemgroup_systemuser_binding", groupname, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting systemuser binding %v\n", username)
		return err
	}

	return nil
}

func addSingleSystemgroupSystemuserBinding(d *schema.ResourceData, meta interface{}, username string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSystemgroupSystemuserBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := system.Systemgroupuserbinding{}
	bindingStruct.Groupname = d.Get("groupname").(string)
	bindingStruct.Username = username

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("systemgroup_systemuser_binding", bindingStruct.Groupname, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateSystemgroupSystemuserBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemgroupCmdpolicyBindings")
	oldSet, newSet := d.GetChange("systemusers")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, username := range remove.List() {
		if err := deleteSingleSystemgroupSystemuserBinding(d, meta, username.(string)); err != nil {
			return err
		}
	}

	for _, username := range add.List() {
		if err := addSingleSystemgroupSystemuserBinding(d, meta, username.(string)); err != nil {
			return err
		}
	}
	return nil
}

func readSystemgroupSystemuserbindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSystemgroupSystemuserbindings")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	bindings, _ := client.FindResourceArray("systemgroup_systemuser_binding", groupname)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, val := range bindings {
		processedBindings[i] = val["username"].(string)
	}

	updatedSet := processedBindings
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("systemusers", updatedSet); err != nil {
		return err
	}
	return nil
}
