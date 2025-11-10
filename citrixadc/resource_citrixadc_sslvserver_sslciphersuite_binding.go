package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcSslvserver_sslciphersuite_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslvserver_sslciphersuite_bindingFunc,
		ReadContext:   readSslvserver_sslciphersuite_bindingFunc,
		DeleteContext: deleteSslvserver_sslciphersuite_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ciphername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vservername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSslvserver_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserver_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vservername := d.Get("vservername").(string)
	ciphername := d.Get("ciphername").(string)
	bindingId := fmt.Sprintf("%s,%s", vservername, ciphername)
	sslvserver_sslciphersuite_binding := ssl.Sslvserversslciphersuitebinding{
		Ciphername:  d.Get("ciphername").(string),
		Description: d.Get("description").(string),
		Vservername: d.Get("vservername").(string),
	}

	_, err := client.AddResource(service.Sslvserver_sslciphersuite_binding.Type(), vservername, &sslvserver_sslciphersuite_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslvserver_sslciphersuite_bindingFunc(ctx, d, meta)
}

func readSslvserver_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserver_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vservername := idSlice[0]
	ciphername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslvserver_sslciphersuite_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslvserver_sslciphersuite_binding",
		ResourceName:             vservername,
		ResourceMissingErrorCode: 461,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslciphersuite_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_sslciphersuite_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ciphername", data["ciphername"])
	d.Set("description", data["description"])
	d.Set("vservername", data["vservername"])

	return nil

}

func deleteSslvserver_sslciphersuite_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserver_sslciphersuite_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ciphername := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["ciphername"] = url.QueryEscape(ciphername)

	err := client.DeleteResourceWithArgsMap(service.Sslvserver_sslciphersuite_binding.Type(), name, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
