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

func resourceCitrixAdcLbvserver_servicegroup_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbvserver_servicegroup_bindingFunc,
		ReadContext:   readLbvserver_servicegroup_bindingFunc,
		DeleteContext: deleteLbvserver_servicegroup_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_servicegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	name := d.Get("name").(string)
	servicegroupname := d.Get("servicegroupname").(string)

	bindingId := fmt.Sprintf("%s,%s", name, servicegroupname)

	lbvserver_servicegroup_binding := lb.Lbvserverservicegroupbinding{
		Name:             d.Get("name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("order"); !raw.IsNull() {
		lbvserver_servicegroup_binding.Order = intPtr(d.Get("order").(int))
	}

	_, err := client.AddResource(service.Lbvserver_servicegroup_binding.Type(), name, &lbvserver_servicegroup_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLbvserver_servicegroup_bindingFunc(ctx, d, meta)
}

func readLbvserver_servicegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	servicegroupname := idSlice[1]

	findParams := service.FindParams{
		ResourceType:             "lbvserver_servicegroup_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_servicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right monitor name
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicegroupname"].(string) == servicegroupname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams monitor name not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_servicegroup_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_servicegroup_binding state %s", bindingId)
	d.Set("name", data["name"])
	d.Set("servicegroupname", data["servicegroupname"])
	setToInt("order", d, data["order"])

	return nil

}

func deleteLbvserver_servicegroup_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_servicegroup_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	servicegroupname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicegroupname:%s", servicegroupname))

	err := client.DeleteResourceWithArgs("lbvserver_servicegroup_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
