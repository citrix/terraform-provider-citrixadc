package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcLbvserver_rewritepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLbvserver_rewritepolicy_bindingFunc,
		ReadContext:   readLbvserver_rewritepolicy_bindingFunc,
		DeleteContext: deleteLbvserver_rewritepolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"bindpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_rewritepolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name").(string)
	policyname := d.Get("policyname").(string)
	bindingId := fmt.Sprintf("%s,%s", name, policyname)
	lbvserver_rewritepolicy_binding := lb.Lbvserverpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               uint32(d.Get("priority").(int)),
	}

	err := client.UpdateUnnamedResource(service.Lbvserver_rewritepolicy_binding.Type(), &lbvserver_rewritepolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLbvserver_rewritepolicy_bindingFunc(ctx, d, meta)
}

func readLbvserver_rewritepolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)
	name := idSlice[0]
	policyname := idSlice[1]
	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_rewritepolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbvserver_rewritepolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_rewritepolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right policy name
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_rewritepolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("bindpoint", data["bindpoint"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("name", data["name"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteLbvserver_rewritepolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()

	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policyname := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["policyname"] = url.QueryEscape(policyname)

	if v, ok := d.GetOk("bindpoint"); ok {
		argsMap["bindpoint"] = url.QueryEscape(v.(string))
	}

	if v, ok := d.GetOk("priority"); ok {
		argsMap["priority"] = url.QueryEscape(fmt.Sprintf("%v", v))
	}
	err := client.DeleteResourceWithArgsMap(service.Lbvserver_rewritepolicy_binding.Type(), name, argsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
