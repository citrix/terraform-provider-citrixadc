package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceCitrixAdcDnspolicylabel_dnspolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDnspolicylabel_dnspolicy_bindingFunc,
		ReadContext:   readDnspolicylabel_dnspolicy_bindingFunc,
		DeleteContext: deleteDnspolicylabel_dnspolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"labelname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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
			"invokelabelname": {
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
		},
	}
}

func createDnspolicylabel_dnspolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicylabel_dnspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	labelname := d.Get("labelname")
	policyname := d.Get("policyname")
	bindingId := fmt.Sprintf("%s,%s", labelname, policyname)
	dnspolicylabel_dnspolicy_binding := dns.Dnspolicylabeldnspolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Invokelabelname:        d.Get("invokelabelname").(string),
		Labelname:              labelname.(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             policyname.(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Dnspolicylabel_dnspolicy_binding.Type(), &dnspolicylabel_dnspolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readDnspolicylabel_dnspolicy_bindingFunc(ctx, d, meta)
}

func readDnspolicylabel_dnspolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicylabel_dnspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicylabel_dnspolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "dnspolicylabel_dnspolicy_binding",
		ResourceName:             labelname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicylabel_dnspolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
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
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicylabel_dnspolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("invokelabelname", data["invokelabelname"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteDnspolicylabel_dnspolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicylabel_dnspolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	labelname := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))
	args = append(args, fmt.Sprintf("priority:%s", strconv.Itoa(d.Get("priority").(int))))

	err := client.DeleteResourceWithArgs(service.Dnspolicylabel_dnspolicy_binding.Type(), labelname, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
