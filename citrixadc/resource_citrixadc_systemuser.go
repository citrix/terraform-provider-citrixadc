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

func resourceCitrixAdcSystemuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemuserFunc,
		Read:          readSystemuserFunc,
		Update:        updateSystemuserFunc,
		Delete:        deleteSystemuserFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"externalauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxsession": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  false,
				Sensitive: true,
			},
			"hashedpassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"promptstring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"cmdpolicybinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      cmdpolicybindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
		},
	}
}

func createSystemuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemuserFunc")
	client := meta.(*NetScalerNitroClient).client
	login_username := (*meta.(*NetScalerNitroClient)).Username
	username := d.Get("username").(string)

	if (username == login_username) {
		return fmt.Errorf("It seems you are trying to change the password of the Admin user. If so, please use the resource \"citrixadc_change_password\"")
	}
	systemuser := system.Systemuser{
		Externalauth: d.Get("externalauth").(string),
		Logging:      d.Get("logging").(string),
		Maxsession:   d.Get("maxsession").(int),
		Password:     d.Get("password").(string),
		Promptstring: d.Get("promptstring").(string),
		Timeout:      d.Get("timeout").(int),
		Username:     username,
	}

	_, err := client.AddResource(service.Systemuser.Type(), username, &systemuser)
	if err != nil {
		return err
	}

	err = updateCmdpolicyBindings(d, meta)
	if err != nil {
		return err
	}

	d.SetId(username)

	err = readSystemuserFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemuser but we can't read it ?? %s", username)
		return nil
	}
	return nil
}

func readSystemuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemuserFunc")
	client := meta.(*NetScalerNitroClient).client

	systemuserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemuser state %s", systemuserName)
	data, err := client.FindResource(service.Systemuser.Type(), systemuserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemuser state %s", systemuserName)
		d.SetId("")
		return nil
	}

	// Need to assess if the hashed password has changed
	// which would mean some other agent changed it besides
	// the current terraform configuration
	oldHashedPassword := ""
	newHashedPassword := ""
	if d, ok := d.GetOk("hashedpassword"); ok {
		oldHashedPassword = d.(string)
	}
	newHashedPassword = data["password"].(string)
	passwordChanged := d.HasChange("password")

	if oldHashedPassword != "" && oldHashedPassword != newHashedPassword && !passwordChanged {
		d.Set("password", "")
	}
	d.Set("name", data["name"])
	d.Set("externalauth", data["externalauth"])
	d.Set("logging", data["logging"])
	d.Set("maxsession", data["maxsession"])
	d.Set("hashedpassword", data["password"])
	d.Set("promptstring", data["promptstring"])
	d.Set("timeout", data["timeout"])

	err = readCmdpolicybindings(d, meta)
	if err != nil {
		return err
	}

	return nil

}

func updateSystemuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemuserFunc")
	client := meta.(*NetScalerNitroClient).client
	systemuserName := d.Get("username").(string)

	systemuser := system.Systemuser{
		Username: d.Get("username").(string),
	}
	hasChange := false
	if d.HasChange("externalauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Externalauth has changed for systemuser %s, starting update", systemuserName)
		systemuser.Externalauth = d.Get("externalauth").(string)
		hasChange = true
	}
	if d.HasChange("logging") {
		log.Printf("[DEBUG]  citrixadc-provider: Logging has changed for systemuser %s, starting update", systemuserName)
		systemuser.Logging = d.Get("logging").(string)
		hasChange = true
	}
	if d.HasChange("maxsession") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxsession has changed for systemuser %s, starting update", systemuserName)
		systemuser.Maxsession = d.Get("maxsession").(int)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for systemuser %s, starting update", systemuserName)
		systemuser.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("promptstring") {
		log.Printf("[DEBUG]  citrixadc-provider: Promptstring has changed for systemuser %s, starting update", systemuserName)
		systemuser.Promptstring = d.Get("promptstring").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for systemuser %s, starting update", systemuserName)
		systemuser.Timeout = d.Get("timeout").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Systemuser.Type(), systemuserName, &systemuser)
		if err != nil {
			return fmt.Errorf("Error updating systemuser %s", systemuserName)
		}
	}
	if d.HasChange("cmdpolicybinding") {
		if err := updateCmdpolicyBindings(d, meta); err != nil {
			return err
		}
	}
	return readSystemuserFunc(d, meta)
}

func deleteSystemuserFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemuserFunc")
	client := meta.(*NetScalerNitroClient).client
	systemuserName := d.Id()
	err := client.DeleteResource(service.Systemuser.Type(), systemuserName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteSingleCmdpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleCmdpolicyBinding")
	client := meta.(*NetScalerNitroClient).client

	username := d.Get("username").(string)
	// Construct args from binding data
	args := make([]string, 0, 1)

	if d, ok := binding["policyname"]; ok {
		s := fmt.Sprintf("policyname:%v", d.(string))
		args = append(args, s)
	}

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs("systemuser_systemcmdpolicy_binding", username, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting cmdpolicy binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleCmdpolicyBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleCmdpolicyBinding")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("Adding binding %v", binding)

	bindingStruct := system.Systemusercmdpolicybinding{}
	bindingStruct.Username = d.Get("username").(string)

	if d, ok := binding["policyname"]; ok {
		bindingStruct.Policyname = d.(string)
	}

	if d, ok := binding["priority"]; ok {
		bindingStruct.Priority = uint32(d.(int))
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource("systemuser_systemcmdpolicy_binding", bindingStruct.Username, bindingStruct); err != nil {
		return err
	}
	return nil
}

func updateCmdpolicyBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCmdpolicyBindings")
	oldSet, newSet := d.GetChange("cmdpolicybinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleCmdpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleCmdpolicyBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func cmdpolicybindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In cmdpolicybindingMappingHash")
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

func readCmdpolicybindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readCmdpolicybindings")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	// Read the lb vserver bindings registered under this policy name
	bindings, _ := client.FindResourceArray("systemuser_systemcmdpolicy_binding", username)
	log.Printf("bindings %v\n", bindings)

	// Process values into new list of maps
	processedBindings := make([]interface{}, len(bindings))
	// Initialize maps
	for i, _ := range bindings {
		processedBindings[i] = make(map[string]interface{})
		processedBindings[i].(map[string]interface{})["policyname"] = bindings[i]["policyname"].(string)
		processedBindings[i].(map[string]interface{})["priority"], _ = strconv.Atoi(bindings[i]["priority"].(string))
	}

	updatedSet := schema.NewSet(cmdpolicybindingMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("cmdpolicybinding", updatedSet); err != nil {
		return err
	}
	return nil
}
