package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcGslbvserver_gslbservice_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createGslbvserver_gslbservice_bindingFunc,
		ReadContext:   readGslbvserver_gslbservice_bindingFunc,
		DeleteContext: deleteGslbvserver_gslbservice_bindingFunc,
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
			"domainname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func createGslbvserver_gslbservice_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	servicename := d.Get("servicename")

	bindingId := fmt.Sprintf("%s,%s", name, servicename)
	gslbvserver_gslbservice_binding := gslb.Gslbvservergslbservicebinding{
		Domainname:  d.Get("domainname").(string),
		Name:        d.Get("name").(string),
		Servicename: d.Get("servicename").(string),
	}

	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		gslbvserver_gslbservice_binding.Weight = intPtr(d.Get("weight").(int))
	}
	if raw := d.GetRawConfig().GetAttr("order"); !raw.IsNull() {
		gslbvserver_gslbservice_binding.Order = intPtr(d.Get("order").(int))
	}

	err := client.UpdateUnnamedResource(service.Gslbvserver_gslbservice_binding.Type(), &gslbvserver_gslbservice_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readGslbvserver_gslbservice_bindingFunc(ctx, d, meta)
}

func readGslbvserver_gslbservice_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbvserver_gslbservice_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbvserver_gslbservice_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservice_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicename"].(string) == servicename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_gslbservice_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("domainname", data["domainname"])
	d.Set("name", data["name"])
	d.Set("servicename", data["servicename"])
	setToInt("weight", d, data["weight"])
	setToInt("order", d, data["order"])

	return nil

}

func deleteGslbvserver_gslbservice_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbvserver_gslbservice_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicename:%s", servicename))

	err := client.DeleteResourceWithArgs(service.Gslbvserver_gslbservice_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
