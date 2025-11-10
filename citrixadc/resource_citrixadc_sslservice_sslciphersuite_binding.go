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

func resourceCitrixAdcSslservice_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslservice_sslciphersuite_bindingFunc,
		ReadContext:   readSslservice_sslciphersuite_bindingFunc,
		DeleteContext: deleteSslservice_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": {
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
			"servicename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSslservice_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslservice_sslciphersuite_bindingFunc(ctx, d, meta)
}

func readSslservice_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
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

func deleteSslservice_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
