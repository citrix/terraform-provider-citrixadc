package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/analytics"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAnalyticsglobal_analyticsprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAnalyticsglobal_analyticsprofile_bindingFunc,
		Read:          readAnalyticsglobal_analyticsprofile_bindingFunc,
		Delete:        deleteAnalyticsglobal_analyticsprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"analyticsprofile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAnalyticsglobal_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	analyticsprofile := d.Get("analyticsprofile").(string)

	analyticsglobal_analyticsprofile_binding := analytics.Analyticsglobalanalyticsprofilebinding{
		Analyticsprofile: d.Get("analyticsprofile").(string),
	}

	err := client.UpdateUnnamedResource("analyticsglobal_analyticsprofile_binding", &analyticsglobal_analyticsprofile_binding)
	if err != nil {
		return err
	}

	d.SetId(analyticsprofile)

	err = readAnalyticsglobal_analyticsprofile_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this analyticsglobal_analyticsprofile_binding but we can't read it ?? %s", analyticsprofile)
		return nil
	}
	return nil
}

func readAnalyticsglobal_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofile := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)

	findParams := service.FindParams{
		ResourceType:             "analyticsglobal",
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
		log.Printf("[WARN] citrixadc-provider: Clearing analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("analyticsprofile", data["analyticsprofile"])

	return nil

}

func deleteAnalyticsglobal_analyticsprofile_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	analyticsprofile := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("analyticsprofile:%s", analyticsprofile))

	err := client.DeleteResourceWithArgs("analyticsglobal_analyticsprofile_binding", "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
