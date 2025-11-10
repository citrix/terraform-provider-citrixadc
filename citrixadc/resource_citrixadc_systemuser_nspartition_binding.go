package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcSystemuser_nspartition_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSystemuser_nspartition_bindingFunc,
		ReadContext:   readSystemuser_nspartition_bindingFunc,
		DeleteContext: deleteSystemuser_nspartition_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"partitionname": {
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

func createSystemuser_nspartition_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemuser_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	username := d.Get("username").(string)
	partitionname := d.Get("partitionname").(string)
	bindingId := fmt.Sprintf("%s,%s", username, partitionname)
	systemuser_nspartition_binding := system.Systemusernspartitionbinding{
		Partitionname: d.Get("partitionname").(string),
		Username:      d.Get("username").(string),
	}

	err := client.UpdateUnnamedResource(service.Systemuser_nspartition_binding.Type(), &systemuser_nspartition_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readSystemuser_nspartition_bindingFunc(ctx, d, meta)
}

func readSystemuser_nspartition_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemuser_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	username := idSlice[0]
	partitionname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading systemuser_nspartition_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "systemuser_nspartition_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemuser_nspartition_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["partitionname"].(string) == partitionname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams partitionname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing systemuser_nspartition_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("partitionname", data["partitionname"])
	d.Set("username", data["username"])

	return nil

}

func deleteSystemuser_nspartition_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemuser_nspartition_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	partitionname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("partitionname:%s", partitionname))

	err := client.DeleteResourceWithArgs(service.Systemuser_nspartition_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
