package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslservicegroup_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslservicegroup_sslciphersuite_bindingFunc,
		Read:          readSslservicegroup_sslciphersuite_bindingFunc,
		Delete:        deleteSslservicegroup_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservicegroup_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslservicegroup_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupname := d.Get("servicegroupname")
	ciphername := d.Get("ciphername")
	bindingId := fmt.Sprintf("%s,%s", servicegroupname, ciphername)
	sslservicegroup_sslciphersuite_binding := ssl.Sslservicegroupsslciphersuitebinding{
		Ciphername:       d.Get("ciphername").(string),
		Description:      d.Get("description").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslservicegroup_sslciphersuite_binding.Type(), &sslservicegroup_sslciphersuite_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslservicegroup_sslciphersuite_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservicegroup_sslciphersuite_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslservicegroup_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslservicegroup_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicegroupname := idSlice[0]
	ciphername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslservicegroup_sslciphersuite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslservicegroup_sslciphersuite_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_sslciphersuite_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslservicegroup_sslciphersuite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ciphername", data["ciphername"])
	d.Set("description", data["description"])
	d.Set("servicegroupname", data["servicegroupname"])

	return nil

}

func deleteSslservicegroup_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslservicegroup_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ciphername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ciphername:%s", ciphername))

	err := client.DeleteResourceWithArgs(service.Sslservicegroup_sslciphersuite_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
