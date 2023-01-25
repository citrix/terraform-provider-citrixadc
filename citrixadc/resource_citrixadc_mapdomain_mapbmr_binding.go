package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcMapdomain_mapbmr_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createMapdomain_mapbmr_bindingFunc,
		Read:          readMapdomain_mapbmr_bindingFunc,
		Delete:        deleteMapdomain_mapbmr_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"mapbmrname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createMapdomain_mapbmr_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	mapbmrname := d.Get("mapbmrname")
	bindingId := fmt.Sprintf("%s,%s", name, mapbmrname)
	mapdomain_mapbmr_binding := network.Mapdomainmapbmrbinding{
		Mapbmrname: d.Get("mapbmrname").(string),
		Name:       d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("mapdomain_mapbmr_binding", &mapdomain_mapbmr_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readMapdomain_mapbmr_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this mapdomain_mapbmr_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readMapdomain_mapbmr_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	mapbmrname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading mapdomain_mapbmr_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "mapdomain_mapbmr_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain_mapbmr_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["mapbmrname"].(string) == mapbmrname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain_mapbmr_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("mapbmrname", data["mapbmrname"])
	d.Set("name", data["name"])

	return nil

}

func deleteMapdomain_mapbmr_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	mapbmrname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("mapbmrname:%s", mapbmrname))

	err := client.DeleteResourceWithArgs("mapdomain_mapbmr_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
