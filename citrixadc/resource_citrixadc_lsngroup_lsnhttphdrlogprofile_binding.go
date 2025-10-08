package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcLsngroup_lsnhttphdrlogprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsngroup_lsnhttphdrlogprofile_bindingFunc,
		ReadContext:   readLsngroup_lsnhttphdrlogprofile_bindingFunc,
		DeleteContext: deleteLsngroup_lsnhttphdrlogprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"httphdrlogprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsnhttphdrlogprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsnhttphdrlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	httphdrlogprofilename := d.Get("httphdrlogprofilename")
	bindingId := fmt.Sprintf("%s,%s", groupname, httphdrlogprofilename)
	lsngroup_lsnhttphdrlogprofile_binding := lsn.Lsngrouplsnhttphdrlogprofilebinding{
		Groupname:             d.Get("groupname").(string),
		Httphdrlogprofilename: d.Get("httphdrlogprofilename").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsnhttphdrlogprofile_binding", &lsngroup_lsnhttphdrlogprofile_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLsngroup_lsnhttphdrlogprofile_bindingFunc(ctx, d, meta)
}

func readLsngroup_lsnhttphdrlogprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsnhttphdrlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	httphdrlogprofilename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsnhttphdrlogprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsnhttphdrlogprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnhttphdrlogprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["httphdrlogprofilename"].(string) == httphdrlogprofilename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams httphdrlogprofilename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsnhttphdrlogprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("httphdrlogprofilename", data["httphdrlogprofilename"])

	return nil

}

func deleteLsngroup_lsnhttphdrlogprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsnhttphdrlogprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	httphdrlogprofilename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("httphdrlogprofilename:%s", httphdrlogprofilename))

	err := client.DeleteResourceWithArgs("lsngroup_lsnhttphdrlogprofile_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
