package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslcipher_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslcipher_sslciphersuite_bindingFunc,
		Read:          readSslcipher_sslciphersuite_bindingFunc,
		Delete:        deleteSslcipher_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ciphergroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ciphername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cipherpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cipheroperation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ciphgrpals": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslcipher_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcipher_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	ciphergroupname := d.Get("ciphergroupname")
	ciphername := d.Get("ciphername")
	bindingId := fmt.Sprintf("%s,%s", ciphergroupname, ciphername)
	sslcipher_sslciphersuite_binding := ssl.Sslciphersslciphersuitebinding{
		Ciphergroupname: d.Get("ciphergroupname").(string),
		Ciphername:      d.Get("ciphername").(string),
		Cipheroperation: d.Get("cipheroperation").(string),
		Cipherpriority:  d.Get("cipherpriority").(int),
		Ciphgrpals:      d.Get("ciphgrpals").(string),
		Description:     d.Get("description").(string),
	}

	_, err := client.AddResource(service.Sslcipher_sslciphersuite_binding.Type(), bindingId, &sslcipher_sslciphersuite_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslcipher_sslciphersuite_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslcipher_sslciphersuite_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslcipher_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslcipher_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	ciphergroupname := idSlice[0]
	ciphername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslcipher_sslciphersuite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslcipher_sslciphersuite_binding",
		ResourceName:             ciphergroupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslcipher_sslciphersuite_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslcipher_sslciphersuite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ciphergroupname", data["ciphergroupname"])
	d.Set("ciphername", data["ciphername"])
	d.Set("cipheroperation", data["cipheroperation"])
	d.Set("cipherpriority", data["cipherpriority"])
	d.Set("ciphgrpals", data["ciphgrpals"])
	d.Set("description", data["description"])

	return nil

}

func deleteSslcipher_sslciphersuite_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslcipher_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	ciphergroupname := idSlice[0]
	ciphername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ciphername:%s", ciphername))

	err := client.DeleteResourceWithArgs(service.Sslcipher_sslciphersuite_binding.Type(), ciphergroupname, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
