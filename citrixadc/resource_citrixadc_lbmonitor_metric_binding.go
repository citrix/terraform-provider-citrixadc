package citrixadc

import (
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbmonitor_metric_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmonitor_metric_bindingFunc,
		Read:          readLbmonitor_metric_bindingFunc,
		Delete:        deleteLbmonitor_metric_bindingFunc,
		Schema: map[string]*schema.Schema{
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metricthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"metricweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLbmonitor_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmonitor_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbmonitorName := d.Get("monitorname").(string)
	metricName := d.Get("metric").(string)
	bindingId := fmt.Sprintf("%s,%s", lbmonitorName, metricName)

	lbmonitor_metric_binding := lb.Lbmonitormetricbinding{
		Metric:          metricName,
		Metricthreshold: d.Get("metricthreshold").(int),
		Monitorname:     lbmonitorName,
		Metricweight:    d.Get("metricweight").(int),
	}

	err := client.UpdateUnnamedResource("lbmonitor_metric_binding", &lbmonitor_metric_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbmonitor_metric_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbmonitor_metric_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbmonitor_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
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
	d.Set("metricthreshold", data["metricthreshold"])
	d.Set("metricweight", data["metricweight"])

	return nil

}

func deleteLbmonitor_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
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
		return err
	}

	d.SetId("")

	return nil
}
