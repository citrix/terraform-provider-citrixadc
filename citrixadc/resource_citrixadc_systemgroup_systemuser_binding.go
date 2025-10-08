package citrixadc

import (
	"context"
	"net/url"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcSystemgroup_systemuser_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemgroup_systemuser_bindingFunc,
		ReadContext:   readSystemgroup_systemuser_bindingFunc,
		DeleteContext: deleteSystemgroup_systemuser_bindingFunc,
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
		},
	}
}

func createSystemgroup_systemuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemgroup_systemuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	username := d.Get("username")
	bindingId := fmt.Sprintf("%s,%s", groupname, username)
	systemgroup_systemuser_binding := system.Systemgroupsystemuserbinding{
		Groupname: d.Get("groupname").(string),
		Username:  d.Get("username").(string),
	}

	_, err := client.AddResource(service.Systemgroup_systemuser_binding.Type(), bindingId, &systemgroup_systemuser_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSystemgroup_systemuser_bindingFunc(ctx, d, meta)
}

func readSystemgroup_systemuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemgroup_systemuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	username := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading systemgroup_systemuser_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "systemgroup_systemuser_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemgroup_systemuser_binding state %s", bindingId)
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
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing systemgroup_systemuser_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("username", data["username"])

	return nil

}

func deleteSystemgroup_systemuser_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemgroup_systemuser_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	username := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("username:%s", url.QueryEscape(username)))

	err := client.DeleteResourceWithArgs(service.Systemgroup_systemuser_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
