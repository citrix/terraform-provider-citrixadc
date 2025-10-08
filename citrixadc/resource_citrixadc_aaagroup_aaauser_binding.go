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

func resourceCitrixAdcAaagroup_aaauser_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaagroup_aaauser_bindingFunc,
		ReadContext:   readAaagroup_aaauser_bindingFunc,
		DeleteContext: deleteAaagroup_aaauser_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
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

func createAaagroup_aaauser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaagroup_aaauser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname").(string)
	username := d.Get("username").(string)
	bindingId := fmt.Sprintf("%s,%s", groupname, username)
	aaagroup_aaauser_binding := aaa.Aaagroupaaauserbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Groupname:              d.Get("groupname").(string),
		Username:               d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaagroup_aaauser_binding.Type(), &aaagroup_aaauser_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readAaagroup_aaauser_bindingFunc(ctx, d, meta)
}

func readAaagroup_aaauser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaagroup_aaauser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	username := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading aaagroup_aaauser_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "aaagroup_aaauser_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_aaauser_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["username"].(string) == username {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams username not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing aaagroup_aaauser_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("groupname", data["groupname"])
	d.Set("username", data["username"])

	return nil

}

func deleteAaagroup_aaauser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaagroup_aaauser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	username := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("username:%s", username))

	err := client.DeleteResourceWithArgs(service.Aaagroup_aaauser_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
