package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcClusternodegroup_service_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusternodegroup_service_bindingFunc,
		ReadContext:   readClusternodegroup_service_bindingFunc,
		DeleteContext: deleteClusternodegroup_service_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternodegroup_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	serviceName := d.Get("service")
	bindingId := fmt.Sprintf("%s,%s", name, serviceName)
	clusternodegroup_service_binding := cluster.Clusternodegroupservicebinding{
		Name:    d.Get("name").(string),
		Service: d.Get("service").(string),
	}

	err := client.UpdateUnnamedResource("clusternodegroup_service_binding", &clusternodegroup_service_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readClusternodegroup_service_bindingFunc(ctx, d, meta)
}

func readClusternodegroup_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	serviceName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_service_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_service_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["service"].(string) == serviceName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams service not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("service", data["service"])

	return nil

}

func deleteClusternodegroup_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	serviceName := idSlice[1]

	args := make([]string, 0)
	//args = append(args, fmt.Sprintf("name:%s", url.QueryEscape(name)))
	args = append(args, fmt.Sprintf("service:%s", url.QueryEscape(serviceName)))

	err := client.DeleteResourceWithArgs("clusternodegroup_service_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
