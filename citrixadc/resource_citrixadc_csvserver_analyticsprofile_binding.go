package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcCsvserver_analyticsprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserver_analyticsprofile_bindingFunc,
		Read:          readCsvserver_analyticsprofile_bindingFunc,
		Delete:        deleteCsvserver_analyticsprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"analyticsprofile": &schema.Schema{
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

func createCsvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	analyticsprofile := d.Get("analyticsprofile")
	bindingId := fmt.Sprintf("%s,%s", name, analyticsprofile)
	csvserver_analyticsprofile_binding := cs.Csvserveranalyticsprofilebinding{
		Analyticsprofile: d.Get("analyticsprofile").(string),
		Name:             d.Get("name").(string),
	}

	_, err := client.AddResource("csvserver_analyticsprofile_binding", bindingId, &csvserver_analyticsprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readCsvserver_analyticsprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this csvserver_analyticsprofile_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readCsvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	analyticsprofile := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_analyticsprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "csvserver_analyticsprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_analyticsprofile_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_analyticsprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("analyticsprofile", data["analyticsprofile"])
	d.Set("name", data["name"])

	return nil

}

func deleteCsvserver_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserver_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	analyticsprofile := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("analyticsprofile:%s", analyticsprofile))

	err := client.DeleteResourceWithArgs("csvserver_analyticsprofile_binding", name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
