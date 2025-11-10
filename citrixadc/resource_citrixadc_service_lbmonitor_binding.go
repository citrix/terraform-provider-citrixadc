package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcService_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createService_lbmonitor_bindingFunc,
		ReadContext:   readService_lbmonitor_bindingFunc,
		DeleteContext: deleteService_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"monitor_name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"monstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passive": {
				Type:     schema.TypeBool,
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
		},
	}
}

func createService_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	monitor_name := d.Get("monitor_name")
	bindingId := fmt.Sprintf("%s,%s", name, monitor_name)
	service_lbmonitor_binding := basic.Servicelbmonitorbinding{
		Monitorname: d.Get("monitor_name").(string),
		Monstate:    d.Get("monstate").(string),
		Name:        d.Get("name").(string),
		Passive:     d.Get("passive").(bool),
	}

	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		service_lbmonitor_binding.Weight = intPtr(d.Get("weight").(int))
	}

	err := client.UpdateUnnamedResource(service.Service_lbmonitor_binding.Type(), &service_lbmonitor_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readService_lbmonitor_bindingFunc(ctx, d, meta)
}

func readService_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading service_lbmonitor_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "service_lbmonitor_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing service_lbmonitor_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing service_lbmonitor_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("monitor_name", data["monitor_name"])
	// d.Set("monstate", data["monstate"])
	d.Set("name", data["name"])
	d.Set("passive", data["passive"])
	setToInt("weight", d, data["weight"])

	return nil

}

func deleteService_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteService_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	monitor_name := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitor_name))

	err := client.DeleteResourceWithArgs(service.Service_lbmonitor_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
