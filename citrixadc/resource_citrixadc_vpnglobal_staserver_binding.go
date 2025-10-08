package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_staserver_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_staserver_bindingFunc,
		ReadContext:   readVpnglobal_staserver_bindingFunc,
		DeleteContext: deleteVpnglobal_staserver_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"staserver": {
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
			"staaddresstype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createVpnglobal_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	staserver := d.Get("staserver").(string)
	vpnglobal_staserver_binding := vpn.Vpnglobalstaserverbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Staaddresstype:         d.Get("staaddresstype").(string),
		Staserver:              d.Get("staserver").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_staserver_binding.Type(), &vpnglobal_staserver_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(staserver)

	return readVpnglobal_staserver_bindingFunc(ctx, d, meta)
}

func readVpnglobal_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	staserver := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_staserver_binding state %s", staserver)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_staserver_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_staserver_binding state %s", staserver)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["staserver"].(string) == staserver {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_staserver_binding state %s", staserver)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("staaddresstype", data["staaddresstype"])
	d.Set("staserver", data["staserver"])

	return nil

}

func deleteVpnglobal_staserver_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_staserver_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	staserver := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("staserver:%s", url.QueryEscape(staserver)))

	err := client.DeleteResourceWithArgs(service.Vpnglobal_staserver_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
