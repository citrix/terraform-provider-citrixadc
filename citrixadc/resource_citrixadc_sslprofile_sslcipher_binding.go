package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcSslprofile_sslcipher_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslprofile_sslcipher_bindingFunc,
		ReadContext:   readSslprofile_sslcipher_bindingFunc,
		DeleteContext: deleteSslprofile_sslcipher_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cipherpriority": {
				Type:     schema.TypeInt,
				Optional: true, // this is optional attribute
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslprofile_sslcipher_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslprofile_sslcipher_bindingFunc(ctx, d, meta)
}

func readSslprofile_sslcipher_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslprofile_sslcipher_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.Split(bindingId, ",")

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce ciphername from id string")
	}

	if len(idSlice) > 2 {
		return diag.Errorf("Too many separators \",\" in id string")
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
			return diag.FromErr(err)
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
	setToInt("cipherpriority", d, data["cipherpriority"])

	return nil

}

func deleteSslprofile_sslcipher_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
