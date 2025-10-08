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

func resourceCitrixAdcSslvserver_ecccurve_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslvserver_ecccurve_bindingFunc,
		ReadContext:   readSslvserver_ecccurve_bindingFunc,
		DeleteContext: deleteSslvserver_ecccurve_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"ecccurvename": {
				Type:     schema.TypeString,
				Required: true,
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

func createSslvserver_ecccurve_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserver_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	vservername := d.Get("vservername").(string)
	ecccurvename := d.Get("ecccurvename").(string)
	bindingId := fmt.Sprintf("%s,%s", vservername, ecccurvename)
	sslvserver_ecccurve_binding := ssl.Sslvserverecccurvebinding{
		Ecccurvename: d.Get("ecccurvename").(string),
		Vservername:  d.Get("vservername").(string),
	}

	_, err := client.AddResource(service.Sslvserver_ecccurve_binding.Type(), vservername, &sslvserver_ecccurve_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSslvserver_ecccurve_bindingFunc(ctx, d, meta)
}

func readSslvserver_ecccurve_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserver_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	vservername := idSlice[0]
	ecccurvename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading sslvserver_ecccurve_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "sslvserver_ecccurve_binding",
		ResourceName:             vservername,
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_ecccurve_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver_ecccurve_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("ecccurvename", data["ecccurvename"])
	d.Set("vservername", data["vservername"])

	return nil

}

func deleteSslvserver_ecccurve_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserver_ecccurve_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	ecccurvename := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["ecccurvename"] = url.QueryEscape(ecccurvename)

	err := client.DeleteResourceWithArgsMap(service.Sslvserver_ecccurve_binding.Type(), name, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
