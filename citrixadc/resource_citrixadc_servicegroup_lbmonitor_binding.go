package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcServicegroup_lbmonitor_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createServicegroup_lbmonitor_bindingFunc,
		ReadContext:   readServicegroup_lbmonitor_bindingFunc,
		DeleteContext: deleteServicegroup_lbmonitor_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"customserverid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dbsttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": {
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
			"nameserver": {
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
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"serverid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": {
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
		},
	}
}

func createServicegroup_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Get("servicegroupname")
	monitorName := d.Get("monitorname")
	// Use `,` as the separator since it is invalid character for servicegroup and monitor name
	servicegroupLbmonitorBindingId := fmt.Sprintf("%s,%s", servicegroupName, monitorName)

	servicegroup_lbmonitor_binding := basic.Servicegroupmonitorbinding{
		Customserverid:   d.Get("customserverid").(string),
		Dbsttl:           uint64(d.Get("dbsttl").(int)),
		Hashid:           uint32(d.Get("hashid").(int)),
		Monitorname:      d.Get("monitorname").(string),
		Monstate:         d.Get("monstate").(string),
		Nameserver:       d.Get("nameserver").(string),
		Passive:          d.Get("passive").(bool),
		Port:             int32(d.Get("port").(int)),
		Serverid:         uint32(d.Get("serverid").(int)),
		Servicegroupname: d.Get("servicegroupname").(string),
		State:            d.Get("state").(string),
		Weight:           uint32(d.Get("weight").(int)),
	}

	err := client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(servicegroupLbmonitorBindingId)

	return readServicegroup_lbmonitor_bindingFunc(ctx, d, meta)
}

func readServicegroup_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupLbmonitorBindingId := d.Id()
	idSlice := strings.Split(servicegroupLbmonitorBindingId, ",")

	if len(idSlice) < 2 {
		return diag.Errorf("Cannot deduce monitorname from id string")
	}

	if len(idSlice) > 2 {
		return diag.Errorf("Too many separators \",\" in id string")
	}

	servicegroupName := idSlice[0]
	monitorName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
	findParams := service.FindParams{
		ResourceType:             "servicegroup_lbmonitor_binding",
		ResourceName:             servicegroupName,
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
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["monitor_name"].(string) == monitorName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup_lbmonitor_binding state %s", servicegroupLbmonitorBindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("customserverid", data["customserverid"])
	setToInt("dbsttl", d, data["dbsttl"])
	setToInt("hashid", d, data["hashid"])
	d.Set("monitorname", data["monitor_name"])
	d.Set("monstate", data["monstate"])
	d.Set("nameserver", data["nameserver"])
	d.Set("passive", data["passive"])
	setToInt("port", d, data["port"])
	setToInt("serverid", d, data["serverid"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("state", data["state"])
	setToInt("weight", d, data["weight"])

	return nil

}

func deleteServicegroup_lbmonitor_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteServicegroup_lbmonitor_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupLbmonitorBindingId := d.Id()
	idSlice := strings.Split(servicegroupLbmonitorBindingId, ",")

	servicegroupName := idSlice[0]
	monitorName := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("monitor_name:%s", monitorName))
	if v, ok := d.GetOk("port"); ok {
		args = append(args, fmt.Sprintf("port:%v", v))
	}

	err := client.DeleteResourceWithArgs(service.Servicegroup_lbmonitor_binding.Type(), servicegroupName, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
