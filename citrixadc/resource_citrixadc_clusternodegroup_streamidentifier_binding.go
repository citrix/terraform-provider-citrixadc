package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcClusternodegroup_streamidentifier_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroup_streamidentifier_bindingFunc,
		Read:          readClusternodegroup_streamidentifier_bindingFunc,
		Delete:        deleteClusternodegroup_streamidentifier_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"identifiername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternodegroup_streamidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	identifiername := d.Get("identifiername")
	bindingId := fmt.Sprintf("%s,%s", name, identifiername)
	clusternodegroup_streamidentifier_binding := cluster.Clusternodegroupstreamidentifierbinding{
		Identifiername: d.Get("identifiername").(string),
		Name:           d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternodegroup_streamidentifier_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup_streamidentifier_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternodegroup_streamidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_streamidentifier_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_streamidentifier_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_streamidentifier_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_streamidentifier_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("identifiername", data["identifiername"])
	d.Set("name", data["name"])

	return nil

}

func deleteClusternodegroup_streamidentifier_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("identifiername:%s", identifiername))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_streamidentifier_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
