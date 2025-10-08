package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcAaauser_intranetip_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaauser_intranetip_bindingFunc,
		ReadContext:   readAaauser_intranetip_bindingFunc,
		DeleteContext: deleteAaauser_intranetip_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"intranetip": {
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
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createAaauser_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauser_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	intranetip := d.Get("intranetip").(string)
	bindingId := fmt.Sprintf("%s,%s", username, intranetip)
	aaauser_intranetip_binding := aaa.Aaauserintranetipbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetip:             d.Get("intranetip").(string),
		Netmask:                d.Get("netmask").(string),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaauser_intranetip_binding.Type(), &aaauser_intranetip_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAaauser_intranetip_bindingFunc(ctx, d, meta)
}

func readAaauser_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauser_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	intranetip := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser_intranetip_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaauser_intranetip_binding",
		ResourceName:             username,
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetip"].(string) == intranetip {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams intranetip not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser_intranetip_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetip", data["intranetip"])
	d.Set("netmask", data["netmask"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaauser_intranetip_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauser_intranetip_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetip := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetip:%s", intranetip))
	if v, ok := d.GetOk("netmask"); ok {
		netmask := v.(string)
		args = append(args, fmt.Sprintf("netmask:%s", netmask))
	}

	err := client.DeleteResourceWithArgs(service.Aaauser_intranetip_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
