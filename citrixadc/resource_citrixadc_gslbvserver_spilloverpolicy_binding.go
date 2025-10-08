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

func resourceCitrixAdcGslbvserver_spilloverpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createGslbvserver_spilloverpolicy_bindingFunc,
		ReadContext:   readGslbvserver_spilloverpolicy_bindingFunc,
		DeleteContext: deleteGslbvserver_spilloverpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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
			"gotopriorityexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGslbvserver_spilloverpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createGslbvserver_spilloverpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	policyname := d.Get("policyname")

	bindingId := fmt.Sprintf("%s,%s", name, policyname)
	gslbvserver_spilloverpolicy_binding := gslb.Gslbvserverspilloverpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	err := client.UpdateUnnamedResource(service.Gslbvserver_spilloverpolicy_binding.Type(), &gslbvserver_spilloverpolicy_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readGslbvserver_spilloverpolicy_bindingFunc(ctx, d, meta)
}

func readGslbvserver_spilloverpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readGslbvserver_spilloverpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policyname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading gslbvserver_spilloverpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "gslbvserver_spilloverpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_spilloverpolicy_binding state %s", bindingId)
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
		log.Printf("[WARN] citrixadc-provider: Clearing gslbvserver_spilloverpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("name", data["name"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteGslbvserver_spilloverpolicy_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteGslbvserver_spilloverpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	policyname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))

	err := client.DeleteResourceWithArgs(service.Gslbvserver_spilloverpolicy_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
