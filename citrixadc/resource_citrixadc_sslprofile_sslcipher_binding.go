package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcSslprofile_sslcipher_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslprofile_sslcipher_bindingFunc,
		Read:          readSslprofile_sslcipher_bindingFunc,
		Delete:        deleteSslprofile_sslcipher_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cipherpriority": &schema.Schema{
				Type:     schema.TypeInt,
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

func createSslprofile_sslcipher_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofile_sslcipher_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	profileName := d.Get("name").(string)
	cipherName := d.Get("ciphername").(string)

	// Use `,` as the separator since it is invalid character for adc entity strings
	bindingId := fmt.Sprintf("%s,%s", profileName, cipherName)

	sslprofile_sslcipher_binding := ssl.Sslprofilecipherbinding{
		Ciphername:     d.Get("ciphername").(string),
		Cipherpriority: uint32(d.Get("cipherpriority").(int)),
		Name:           d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslprofile_sslcipher_binding.Type(), &sslprofile_sslcipher_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readSslprofile_sslcipher_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslprofile_sslcipher_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readSslprofile_sslcipher_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslprofile_sslcipher_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	if len(idSlice) < 2 {
		return fmt.Errorf("Cannot deduce ciphername from id string")
	}

	if len(idSlice) > 2 {
		return fmt.Errorf("Too many separators \",\" in id string")
	}

	profileName := idSlice[0]
	cipherName := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             "sslprofile_sslcipher_binding",
		ResourceName:             profileName,
		ResourceMissingErrorCode: 3248,
	}

	dataArr, err := client.FindResourceArrayWithParams(findParams)

	if err != nil {
		if strings.Contains(err.Error(), "\"errorcode\": 3248") {
			return nil
		} else {
			// Unexpected error
			log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
			return err
		}
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_sslcipher_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["cipheraliasname"].(string) == cipherName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams cipher name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_sslcipher_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("ciphername", data["cipheraliasname"])
	d.Set("cipherpriority", data["cipherpriority"])

	return nil

}

func deleteSslprofile_sslcipher_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslprofile_sslcipher_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	profileName := idSlice[0]
	cipherName := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("ciphername:%v", cipherName))

	err := client.DeleteResourceWithArgs(service.Sslprofile_sslcipher_binding.Type(), profileName, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
