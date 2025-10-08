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

func resourceCitrixAdcSslprofile_sslcertkey_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslprofile_sslcertkey_bindingFunc,
		ReadContext:   readSslprofile_sslcertkey_bindingFunc,
		DeleteContext: deleteSslprofile_sslcertkey_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cipherpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sslicacertkey": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslprofile_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofile_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	sslicacertkey := d.Get("sslicacertkey")
	bindingId := fmt.Sprintf("%s,%s", name, sslicacertkey)
	sslprofile_sslcertkey_binding := ssl.Sslprofilesslcertkeybinding{
		Cipherpriority: d.Get("cipherpriority").(int),
		Name:           d.Get("name").(string),
		Sslicacertkey:  d.Get("sslicacertkey").(string),
	}

	err := client.UpdateUnnamedResource("sslprofile_sslcertkey_binding", &sslprofile_sslcertkey_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslprofile_sslcertkey_bindingFunc(ctx, d, meta)
}

func readSslprofile_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslprofile_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sslicacertkey := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslprofile_sslcertkey_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslprofile_sslcertkey_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 3248,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["sslicacertkey"].(string) == sslicacertkey {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams sslicacertkey not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile_sslcertkey_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("cipherpriority", d, data["cipherpriority"])
	d.Set("name", data["name"])
	d.Set("sslicacertkey", data["sslicacertkey"])

	return nil

}

func deleteSslprofile_sslcertkey_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslprofile_sslcertkey_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	sslicacertkey := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("sslicacertkey:%s", sslicacertkey))

	err := client.DeleteResourceWithArgs("sslprofile_sslcertkey_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
