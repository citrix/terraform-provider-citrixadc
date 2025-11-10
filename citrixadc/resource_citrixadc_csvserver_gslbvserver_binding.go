package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcCsvserver_gslbvserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCsvserver_gslbvserver_bindingFunc,
		ReadContext:   readCsvserver_gslbvserver_bindingFunc,
		DeleteContext: deleteCsvserver_gslbvserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vserver": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCsvserver_gslbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createCsvserver_gslbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	vserver := d.Get("vserver").(string)
	bindingId := fmt.Sprintf("%s,%s", name, vserver)
	csvserver_gslbvserver_binding := cs.Csvservergslbvserverbinding{
		Name:    name,
		Vserver: vserver,
	}

	_, err := client.AddResource(service.Csvserver_gslbvserver_binding.Type(), name, &csvserver_gslbvserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readCsvserver_gslbvserver_bindingFunc(ctx, d, meta)
}

func readCsvserver_gslbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readCsvserver_gslbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vserver := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading csvserver_gslbvserver_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "csvserver_gslbvserver_binding",
		ResourceName:             name,
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
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_gslbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["vserver"].(string) == vserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams vserver not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing csvserver_gslbvserver_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("vserver", data["vserver"])

	return nil

}

func deleteCsvserver_gslbvserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCsvserver_gslbvserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	vserver := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("vserver:%s", vserver))

	err := client.DeleteResourceWithArgs(service.Csvserver_gslbvserver_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
