package citrixadc

import (
	"context"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcLbmonitor_metric_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbmonitor_metric_bindingFunc,
		ReadContext:   readLbmonitor_metric_bindingFunc,
		DeleteContext: deleteLbmonitor_metric_bindingFunc,
		Schema: map[string]*schema.Schema{
			"monitorname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metric": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metricthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"metricweight": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func createLbmonitor_metric_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmonitor_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmonitorName := d.Get("monitorname").(string)
	metricName := d.Get("metric").(string)
	bindingId := fmt.Sprintf("%s,%s", lbmonitorName, metricName)

	lbmonitor_metric_binding := lb.Lbmonitormetricbinding{
		Metric:      metricName,
		Monitorname: lbmonitorName,
	}

	if raw := d.GetRawConfig().GetAttr("metricthreshold"); !raw.IsNull() {
		lbmonitor_metric_binding.Metricthreshold = intPtr(d.Get("metricthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("metricweight"); !raw.IsNull() {
		lbmonitor_metric_binding.Metricweight = intPtr(d.Get("metricweight").(int))
	}

	err := client.UpdateUnnamedResource("lbmonitor_metric_binding", &lbmonitor_metric_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLbmonitor_metric_bindingFunc(ctx, d, meta)
}

func readLbmonitor_metric_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbmonitor_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbmonitor_metric_binding state %s", bindingId)
	idSlice := strings.SplitN(bindingId, ",", 2)
	lbmonitorName := idSlice[0]
	metricName := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading lbmonitor_metric_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbmonitor_metric_binding",
		ResourceName:             lbmonitorName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbmonitor_metric_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right metric name
	foundIndex := -1
	for i, v := range dataArr {
		if v["metric"].(string) == metricName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbmonitor_metric_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Fallthrough
	data := dataArr[foundIndex]

	d.Set("monitorname", data["monitorname"])
	d.Set("metric", data["metric"])
	setToInt("metricthreshold", d, data["metricthreshold"])
	setToInt("metricweight", d, data["metricweight"])

	return nil

}

func deleteLbmonitor_metric_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbmonitor_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	lbmonitorName := idSlice[0]
	metricName := idSlice[1]

	args := make(map[string]string)
	args["metric"] = metricName
	err := client.DeleteResourceWithArgsMap("lbmonitor_metric_binding", lbmonitorName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
