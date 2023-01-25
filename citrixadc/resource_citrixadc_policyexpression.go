package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcPolicyexpression() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicyexpressionFunc,
		Read:          readPolicyexpressionFunc,
		Update:        updatePolicyexpressionFunc,
		Delete:        deletePolicyexpressionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clientsecuritymessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createPolicyexpressionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Get("name").(string)
	policyexpression := policy.Policyexpression{
		Clientsecuritymessage: d.Get("clientsecuritymessage").(string),
		Comment:               d.Get("comment").(string),
		Description:           d.Get("description").(string),
		Name:                  d.Get("name").(string),
		Value:                 d.Get("value").(string),
	}

	_, err := client.AddResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
	if err != nil {
		return err
	}

	d.SetId(policyexpressionName)

	err = readPolicyexpressionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policyexpression but we can't read it ?? %s", policyexpressionName)
		return nil
	}
	return nil
}

func readPolicyexpressionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policyexpression state %s", policyexpressionName)
	data, err := client.FindResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policyexpression state %s", policyexpressionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("clientsecuritymessage", data["clientsecuritymessage"])
	d.Set("comment", data["comment"])
	d.Set("description", data["description"])
	d.Set("name", data["name"])
	d.Set("value", data["value"])

	return nil

}

func updatePolicyexpressionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Get("name").(string)

	policyexpression := policy.Policyexpression{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("clientsecuritymessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritymessage has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Clientsecuritymessage = d.Get("clientsecuritymessage").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("description") {
		log.Printf("[DEBUG]  citrixadc-provider: Description has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Description = d.Get("description").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("value") {
		log.Printf("[DEBUG]  citrixadc-provider: Value has changed for policyexpression %s, starting update", policyexpressionName)
		policyexpression.Value = d.Get("value").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Policyexpression.Type(), policyexpressionName, &policyexpression)
		if err != nil {
			return fmt.Errorf("Error updating policyexpression %s: %s", policyexpressionName, err.Error())
		}
	}
	return readPolicyexpressionFunc(d, meta)
}

func deletePolicyexpressionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicyexpressionFunc")
	client := meta.(*NetScalerNitroClient).client
	policyexpressionName := d.Id()
	err := client.DeleteResource(service.Policyexpression.Type(), policyexpressionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
