package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strings"
)

func resourceCitrixAdcLbmetrictable_metric_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbmetrictable_metric_bindingFunc,
		Read:          readLbmetrictable_metric_bindingFunc,
		Delete:        deleteLbmetrictable_metric_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snmpoid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLbmetrictable_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbmetrictable_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	metrictable := d.Get("metrictable")
	secondIdComponent := d.Get("metric")
	bindingId := fmt.Sprintf("%s,%s", metrictable, secondIdComponent)
	lbmetrictable_metric_binding := lb.Lbmetrictablemetricbinding{
		Metric:      d.Get("metric").(string),
		Metrictable: d.Get("metrictable").(string),
		Snmpoid:     d.Get("snmpoid").(string),
	}

	_, err := client.AddResource(service.Lbmetrictable_metric_binding.Type(), bindingId, &lbmetrictable_metric_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbmetrictable_metric_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbmetrictable_metric_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbmetrictable_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbmetrictable_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	metrictable := idSlice[0]
	metric := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbmetrictable_metric_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbmetrictable_metric_binding",
		ResourceName:             metrictable,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbmetrictable_metric_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["metric"].(string) == metric {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbmetrictable_metric_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("metric", data["metric"])
	d.Set("metrictable", data["metrictable"])
	d.Set("snmpoid", data["Snmpoid"])

	return nil

}

func deleteLbmetrictable_metric_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbmetrictable_metric_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	metric := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("metric:%s", metric))

	err := client.DeleteResourceWithArgs(service.Lbmetrictable_metric_binding.Type(), name, args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
