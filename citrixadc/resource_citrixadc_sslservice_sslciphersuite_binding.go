package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslservice_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslservice_sslciphersuite_bindingFunc,
		Read:          readSslservice_sslciphersuite_bindingFunc,
		Delete:        deleteSslservice_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservice_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservice_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename")
	ciphername := d.Get("ciphername")
	bindingId := fmt.Sprintf("%s,%s", servicename, ciphername)
	sslservice_sslciphersuite_binding := ssl.Sslservicesslciphersuitebinding{
		Ciphername:  d.Get("ciphername").(string),
		Description: d.Get("description").(string),
		Servicename: d.Get("servicename").(string),
	}

	_, err := client.AddResource(service.Sslservice_sslciphersuite_binding.Type(), bindingId, &sslservice_sslciphersuite_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslservice_sslciphersuite_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservice_sslciphersuite_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslservice_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservice_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	ciphername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservice_sslciphersuite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservice_sslciphersuite_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_sslciphersuite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["ciphername"].(string) == ciphername {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams ciphername not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice_sslciphersuite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ciphername", data["ciphername"])
	d.Set("description", data["description"])
	d.Set("servicename", data["servicename"])

	return nil

}

func deleteSslservice_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservice_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicename := idSlice[0]
	ciphername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ciphername:%s", ciphername))

	err := client.DeleteResourceWithArgs(service.Sslservice_sslciphersuite_binding.Type(), servicename, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
