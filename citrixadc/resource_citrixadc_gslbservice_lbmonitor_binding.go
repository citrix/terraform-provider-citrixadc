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

func resourceCitrixAdcGslbservice_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createGslbservice_lbmonitor_bindingFunc,
		ReadContext:   readGslbservice_lbmonitor_bindingFunc,
		DeleteContext: deleteGslbservice_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"monitor_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"monstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		},
	}
}

func createGslbservice_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename").(string)
	monitor_name := d.Get("monitor_name").(string)

	bindingId := fmt.Sprintf("%s,%s", servicename, monitor_name)
	gslbservice_lbmonitor_binding := gslb.Gslbservicelbmonitorbinding{
		Monitorname: d.Get("monitor_name").(string),
		Monstate:    d.Get("monstate").(string),
		Servicename: d.Get("servicename").(string),
	}

	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		gslbservice_lbmonitor_binding.Weight = intPtr(d.Get("weight").(int))
	}

	err := client.UpdateUnnamedResource(service.Gslbservice_lbmonitor_binding.Type(), &gslbservice_lbmonitor_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readGslbservice_lbmonitor_bindingFunc(ctx, d, meta)
}

func readGslbservice_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce monitor_name from id string")
	}

	if len(idSlice) > 2 {
		return diag.Errorf("Too many separators \",\" in id string")
	}

	servicename := idSlice[0]
	monitor_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbservice_lbmonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbservice_lbmonitor_binding",
		ResourceName:             servicename,
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["monitor_name"].(string) == monitor_name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing gslbservice_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("monitor_name", data["monitor_name"])
	d.Set("monstate", data["monstate"])
	d.Set("servicename", data["servicename"])
	setToInt("weight", d, data["weight"])

	return nil

}

func deleteGslbservice_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbservice_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitor_name))

	err := client.DeleteResourceWithArgs(service.Gslbservice_lbmonitor_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
