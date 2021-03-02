package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/policy"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcPolicypatset_pattern_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicypatset_pattern_bindingFunc,
		Read:          readPolicypatset_pattern_bindingFunc,
		Delete:        deletePolicypatset_pattern_bindingFunc,
		Schema: map[string]*schema.Schema{
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"feature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicypatset_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	stringText := d.Get("string").(string)
	bindingId := fmt.Sprintf("%v,%v", name, stringText)

	policypatset_pattern_binding := policy.Policypatsetpatternbinding{
		Charset: d.Get("charset").(string),
		Comment: d.Get("comment").(string),
		Feature: d.Get("feature").(string),
		Index:   d.Get("index").(int),
		Name:    d.Get("name").(string),
		String:  d.Get("string").(string),
	}

	err := client.UpdateUnnamedResource(netscaler.Policypatset_pattern_binding.Type(), &policypatset_pattern_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readPolicypatset_pattern_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policypatset_pattern_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readPolicypatset_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	stringText := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading policypatset_pattern_binding state %s", bindingId)
	findParams := netscaler.FindParams{
		ResourceType:             "policypatset_pattern_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 2823,
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
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right String
	foundIndex := -1
	for i, v := range dataArr {
		if v["String"].(string) == stringText {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams String not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policypatset_pattern_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("charset", data["charset"])
	d.Set("comment", data["comment"])
	d.Set("feature", data["feature"])
	d.Set("index", data["index"])
	d.Set("name", data["name"])
	d.Set("string", data["String"])

	return nil

}

func deletePolicypatset_pattern_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicypatset_pattern_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	stringText := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["String"] = url.QueryEscape(stringText)

	err := client.DeleteResourceWithArgsMap(netscaler.Policypatset_pattern_binding.Type(), name, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
