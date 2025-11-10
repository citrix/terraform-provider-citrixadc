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

func resourceCitrixAdcSslcipher_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslcipher_sslciphersuite_bindingFunc,
		ReadContext:   readSslcipher_sslciphersuite_bindingFunc,
		DeleteContext: deleteSslcipher_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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

func createSslcipher_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslcipher_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	ciphergroupname := d.Get("ciphergroupname")
	ciphername := d.Get("ciphername")
	bindingId := fmt.Sprintf("%s,%s", ciphergroupname, ciphername)
	sslcipher_sslciphersuite_binding := ssl.Sslciphersslciphersuitebinding{
		Ciphergroupname: d.Get("ciphergroupname").(string),
		Ciphername:      d.Get("ciphername").(string),
		Cipheroperation: d.Get("cipheroperation").(string),
		Ciphgrpals:      d.Get("ciphgrpals").(string),
		Description:     d.Get("description").(string),
	}

	if raw := d.GetRawConfig().GetAttr("cipherpriority"); !raw.IsNull() {
		sslcipher_sslciphersuite_binding.Cipherpriority = intPtr(d.Get("cipherpriority").(int))
	}

	_, err := client.AddResource(service.Sslcipher_sslciphersuite_binding.Type(), bindingId, &sslcipher_sslciphersuite_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslcipher_sslciphersuite_bindingFunc(ctx, d, meta)
}

func readSslcipher_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
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
	setToInt("cipherpriority", d, data["cipherpriority"])
	d.Set("ciphgrpals", data["ciphgrpals"])
	d.Set("description", data["description"])

	return nil

}

func deleteSslcipher_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
