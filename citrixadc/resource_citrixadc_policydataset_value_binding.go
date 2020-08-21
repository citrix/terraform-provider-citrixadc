package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/policy"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcPolicydataset_value_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createPolicydataset_value_bindingFunc,
		Read:          readPolicydataset_value_bindingFunc,
		Delete:        deletePolicydataset_value_bindingFunc,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
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
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createPolicydataset_value_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createPolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	value := d.Get("value").(string)
	// Use `,` as the separator since it is invalid character for servicegroup and monitor name
	valueBindingId := fmt.Sprintf("%s,%s", name, value)

	policydataset_value_binding := policy.Policydatasetvaluebinding{
		Comment: d.Get("comment").(string),
		Index:   d.Get("index").(int),
		Name:    d.Get("name").(string),
		Value:   d.Get("value").(string),
	}

	err := client.UpdateUnnamedResource(netscaler.Policydataset_value_binding.Type(), &policydataset_value_binding)
	if err != nil {
		return err
	}

	d.SetId(valueBindingId)

	err = readPolicydataset_value_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this policydataset_value_binding but we can't read it ?? %s", valueBindingId)
		return nil
	}
	return nil
}

func readPolicydataset_value_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readPolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	valueBindingId := d.Id()

	idSlice := strings.Split(valueBindingId, ",")

	if len(idSlice) < 2 {
		return fmt.Errorf("Cannot deduce value from id string")
	}

	if len(idSlice) > 2 {
		return fmt.Errorf("Too many separators \",\" in id string")
	}

	name := idSlice[0]
	value := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_lbmonitor_binding state %s", valueBindingId)
	findParams := netscaler.FindParams{
		ResourceType:             "policydataset_value_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset_value_binding state %s", valueBindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["value"].(string) == value {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing policydataset_value_binding state %s", valueBindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("index", data["index"])
	d.Set("name", data["name"])
	d.Set("value", data["value"])

	return nil

}

func deletePolicydataset_value_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePolicydataset_value_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	args := make(map[string]string)
	args["value"] = d.Get("value").(string)
	err := client.DeleteResourceWithArgsMap(netscaler.Policydataset_value_binding.Type(), d.Get("name").(string), args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
