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

func resourceCitrixAdcAaagroup_vpnintranetapplication_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaagroup_vpnintranetapplication_bindingFunc,
		ReadContext:   readAaagroup_vpnintranetapplication_bindingFunc,
		DeleteContext: deleteAaagroup_vpnintranetapplication_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"intranetapplication": {
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
		},
	}
}

func createAaagroup_vpnintranetapplication_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroup_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	intranetapplication := d.Get("intranetapplication").(string)
	bindingId := fmt.Sprintf("%s,%s", groupname, intranetapplication)
	aaagroup_vpnintranetapplication_binding := aaa.Aaagroupvpnintranetapplicationbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupname:              d.Get("groupname").(string),
		Intranetapplication:    d.Get("intranetapplication").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaagroup_vpnintranetapplication_binding.Type(), &aaagroup_vpnintranetapplication_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAaagroup_vpnintranetapplication_bindingFunc(ctx, d, meta)
}

func readAaagroup_vpnintranetapplication_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroup_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	intranetapplication := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup_vpnintranetapplication_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaagroup_vpnintranetapplication_binding",
		ResourceName:             groupname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_vpnintranetapplication_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetapplication"].(string) == intranetapplication {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams intranetapplication not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_vpnintranetapplication_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("intranetapplication", data["intranetapplication"])

	return nil

}

func deleteAaagroup_vpnintranetapplication_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroup_vpnintranetapplication_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	intranetapplication := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetapplication:%s", intranetapplication))

	err := client.DeleteResourceWithArgs(service.Aaagroup_vpnintranetapplication_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
