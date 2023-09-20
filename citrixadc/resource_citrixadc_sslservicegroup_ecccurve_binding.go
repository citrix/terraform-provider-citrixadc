package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcSslservicegroup_ecccurve_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslservicegroup_ecccurve_bindingFunc,
		Read:          readSslservicegroup_ecccurve_bindingFunc,
		Delete:        deleteSslservicegroup_ecccurve_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ecccurvename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservicegroup_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservicegroup_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupname := d.Get("servicegroupname").(string)
	ecccurvename := d.Get("ecccurvename").(string)
	bindingId := fmt.Sprintf("%s,%s", servicegroupname, ecccurvename)
	sslservicegroup_ecccurve_binding := ssl.Sslservicegroupecccurvebinding{
		Servicegroupname: servicegroupname,
		Ecccurvename:     ecccurvename,
	}

	err := client.UpdateUnnamedResource(service.Sslservicegroup_ecccurve_binding.Type(), &sslservicegroup_ecccurve_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslservicegroup_ecccurve_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservicegroup_ecccurve_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslservicegroup_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservicegroup_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicegroupname := idSlice[0]
	ecccurvename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservicegroup_ecccurve_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservicegroup_ecccurve_binding",
		ResourceName:             servicegroupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ecccurvename"].(string) == ecccurvename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ecccurvename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ecccurvename", data["ecccurvename"])
	d.Set("servicegroupname", data["servicegroupname"])

	return nil

}

func deleteSslservicegroup_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservicegroup_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ecccurvename := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["ecccurvename"] = url.QueryEscape(ecccurvename)

	err := client.DeleteResourceWithArgsMap(service.Sslservicegroup_ecccurve_binding.Type(), name, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
