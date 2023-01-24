package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cr"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcCrvserver_analyticsprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCrvserver_analyticsprofile_bindingFunc,
		Read:          readCrvserver_analyticsprofile_bindingFunc,
		Delete:        deleteCrvserver_analyticsprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"analyticsprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createCrvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCrvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	analyticsprofile := d.Get("analyticsprofile")
	bindingId := fmt.Sprintf("%s,%s", name, analyticsprofile)
	crvserver_analyticsprofile_binding := cr.Crvserveranalyticsprofilebinding{
		Analyticsprofile: analyticsprofile.(string),
		Name:             name.(string),
	}

	err := client.UpdateUnnamedResource("crvserver_analyticsprofile_binding", &crvserver_analyticsprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readCrvserver_analyticsprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this crvserver_analyticsprofile_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readCrvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCrvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	analyticsprofile := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading crvserver_analyticsprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "crvserver_analyticsprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing crvserver_analyticsprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["analyticsprofile"].(string) == analyticsprofile {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams analyticsprofile not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing crvserver_analyticsprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("analyticsprofile", data["analyticsprofile"])
	d.Set("name", data["name"])

	return nil

}

func deleteCrvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	analyticsprofile := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("analyticsprofile:%s", analyticsprofile))

	err := client.DeleteResourceWithArgs("crvserver_analyticsprofile_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
