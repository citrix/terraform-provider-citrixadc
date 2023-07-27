package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditnslogglobal_auditnslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditnslogglobal_auditnslogpolicy_bindingFunc,
		Read:          readAuditnslogglobal_auditnslogpolicy_bindingFunc,
		Delete:        deleteAuditnslogglobal_auditnslogpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"globalbindtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAuditnslogglobal_auditnslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditnslogglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	auditnslogglobal_auditnslogpolicy_binding := audit.Auditnslogglobalauditnslogpolicybinding{
		Globalbindtype: d.Get("globalbindtype").(string),
		Policyname:     d.Get("policyname").(string),
		Priority:       d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource("auditnslogglobal_auditnslogpolicy_binding", &auditnslogglobal_auditnslogpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readAuditnslogglobal_auditnslogpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditnslogglobal_auditnslogpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readAuditnslogglobal_auditnslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditnslogglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading auditnslogglobal_auditnslogpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "auditnslogglobal_auditnslogpolicy_binding",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing auditnslogglobal_auditnslogpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing auditnslogglobal_auditnslogpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("policyname", data["policyname"])
	d.Set("priority", data["priority"])

	return nil

}

func deleteAuditnslogglobal_auditnslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditnslogglobal_auditnslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	if v, ok := d.GetOk("globalbindtype"); ok {
		bind_type := v.(string)
		args = append(args, fmt.Sprintf("globalbindtype:%s", bind_type))
	} else {
		args = append(args, fmt.Sprintf("globalbindtype:SYSTEM_GLOBAL"))
	}

	err := client.DeleteResourceWithArgs("auditnslogglobal_auditnslogpolicy_binding", "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
