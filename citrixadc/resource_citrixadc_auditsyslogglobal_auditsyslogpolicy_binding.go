package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
)

func resourceCitrixAdcAuditsyslogglobal_auditsyslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
		Read:          readAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
		Delete:        deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"builtin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"feature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"globalbindtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAuditsyslogglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId(policyname)

	err = readAuditsyslogglobal_auditsyslogpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditsyslogglobal_auditsyslogpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readAuditsyslogglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
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

func deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(policyname)))
	if v, ok := d.GetOk("globalbindtype"); ok {
		bind_type := v.(string)
		args = append(args, fmt.Sprintf("globalbindtype:%s", bind_type))
	} else {
		args = append(args,fmt.Sprintf("globalbindtype:SYSTEM_GLOBAL"))
	}
	err := client.DeleteResourceWithArgs("auditsyslogglobal_auditsyslogpolicy_binding", "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
