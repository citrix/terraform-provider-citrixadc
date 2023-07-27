package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcClusternodegroup_nslimitidentifier_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroup_nslimitidentifier_bindingFunc,
		Read:          readClusternodegroup_nslimitidentifier_bindingFunc,
		Delete:        deleteClusternodegroup_nslimitidentifier_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"identifiername": {
				Type:     schema.TypeString,
				Required: true,
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

func createClusternodegroup_nslimitidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_nslimitidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	identifiername := d.Get("identifiername")
	bindingId := fmt.Sprintf("%s,%s", name, identifiername)
	clusternodegroup_nslimitidentifier_binding := cluster.Clusternodegroupnslimitidentifierbinding{
		Identifiername: d.Get("identifiername").(string),
		Name:           d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_nslimitidentifier_binding.Type(), &clusternodegroup_nslimitidentifier_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternodegroup_nslimitidentifier_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup_nslimitidentifier_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternodegroup_nslimitidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_nslimitidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_nslimitidentifier_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_nslimitidentifier_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_nslimitidentifier_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["identifiername"].(string) == identifiername {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams identifiername not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_nslimitidentifier_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("identifiername", data["identifiername"])
	d.Set("name", data["name"])

	return nil

}

func deleteClusternodegroup_nslimitidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_nslimitidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("identifiername:%s", identifiername))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_nslimitidentifier_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
