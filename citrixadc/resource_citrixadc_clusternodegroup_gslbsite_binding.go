package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcClusternodegroup_gslbsite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createClusternodegroup_gslbsite_bindingFunc,
		Read:          readClusternodegroup_gslbsite_bindingFunc,
		Delete:        deleteClusternodegroup_gslbsite_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"gslbsite": &schema.Schema{
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

func createClusternodegroup_gslbsite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	gslbsite := d.Get("gslbsite")
	bindingId := fmt.Sprintf("%s,%s", name, gslbsite)
	clusternodegroup_gslbsite_binding := cluster.Clusternodegroupgslbsitebinding{
		Gslbsite: d.Get("gslbsite").(string),
		Name:     d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_gslbsite_binding.Type(), &clusternodegroup_gslbsite_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readClusternodegroup_gslbsite_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this clusternodegroup_gslbsite_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readClusternodegroup_gslbsite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	gslbsite := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_gslbsite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_gslbsite_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_gslbsite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["gslbsite"].(string) == gslbsite {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams gslbsite not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_gslbsite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gslbsite", data["gslbsite"])
	d.Set("name", data["name"])

	return nil

}

func deleteClusternodegroup_gslbsite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_gslbsite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	gslbsite := idSlice[1]

	args := make([]string, 0)
	
	args = append(args, fmt.Sprintf("gslbsite:%s", url.QueryEscape(gslbsite)))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_gslbsite_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
