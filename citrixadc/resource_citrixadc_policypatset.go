package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcPolicypatset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicypatsetFunc,
		Read:          readPolicypatsetFunc,
		Delete:        deletePolicypatsetFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicypatsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Get("name").(string)
	policypatset := policy.Policypatset{
		Comment:   d.Get("comment").(string),
		Indextype: d.Get("indextype").(string),
		Name:      d.Get("name").(string),
	}

	_, err := client.AddResource(service.Policypatset.Type(), policypatsetName, &policypatset)
	if err != nil {
		return err
	}

	d.SetId(policypatsetName)

	err = readPolicypatsetFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policypatset but we can't read it ?? %s", policypatsetName)
		return nil
	}
	return nil
}

func readPolicypatsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading policypatset state %s", policypatsetName)
	data, err := client.FindResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset state %s", policypatsetName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("indextype", data["indextype"])
	d.Set("name", data["name"])

	return nil

}

func deletePolicypatsetFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicypatsetFunc")
	client := meta.(*NetScalerNitroClient).client
	policypatsetName := d.Id()
	err := client.DeleteResource(service.Policypatset.Type(), policypatsetName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
