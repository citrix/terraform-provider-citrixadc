package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcLbvserver_service_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbvserver_service_bindingFunc,
		ReadContext:   readLbvserver_service_bindingFunc,
		DeleteContext: deleteLbvserver_service_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	servicename := d.Get("servicename").(string)

	bindingId := fmt.Sprintf("%s,%s", name, servicename)

	lbvserver_service_binding := lb.Lbvserverservicebinding{
		Name:        d.Get("name").(string),
		Servicename: d.Get("servicename").(string),
		Weight:      d.Get("weight").(int),
		Order:       d.Get("order").(int),
	}

	_, err := client.AddResource(service.Lbvserver_service_binding.Type(), name, &lbvserver_service_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLbvserver_service_bindingFunc(ctx, d, meta)
}

func readLbvserver_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	servicename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_service_binding state %v", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbvserver_service_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicename"].(string) == servicename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_service_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("servicename", data["servicename"])
	setToInt("weight", d, data["weight"])
	setToInt("order", d, data["order"])

	return nil

}

func deleteLbvserver_service_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_service_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicename:%s", servicename))

	err := client.DeleteResourceWithArgs("lbvserver_service_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
