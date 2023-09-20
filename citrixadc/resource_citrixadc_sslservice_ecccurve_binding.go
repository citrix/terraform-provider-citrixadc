package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslservice_ecccurve_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslservice_ecccurve_bindingFunc,
		Read:          readSslservice_ecccurve_bindingFunc,
		Delete:        deleteSslservice_ecccurve_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ecccurvename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservice_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservice_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename")
	ecccurvename := d.Get("ecccurvename")
	bindingId := fmt.Sprintf("%s,%s", servicename, ecccurvename)
	sslservice_ecccurve_binding := ssl.Sslserviceecccurvebinding{
		Ecccurvename: d.Get("ecccurvename").(string),
		Servicename:  d.Get("servicename").(string),
	}

	_, err := client.AddResource(service.Sslservice_ecccurve_binding.Type(), bindingId, &sslservice_ecccurve_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslservice_ecccurve_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservice_ecccurve_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslservice_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservice_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	ecccurvename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservice_ecccurve_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservice_ecccurve_binding",
		ResourceName:             servicename,
		ResourceMissingErrorCode: 463,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_ecccurve_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ecccurvename", data["ecccurvename"])
	d.Set("servicename", data["servicename"])

	return nil

}

func deleteSslservice_ecccurve_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservice_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ecccurvename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ecccurvename:%s", ecccurvename))

	err := client.DeleteResourceWithArgs(service.Sslservice_ecccurve_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
