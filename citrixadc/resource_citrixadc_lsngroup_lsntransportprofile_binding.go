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

func resourceCitrixAdcLsngroup_lsntransportprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createLsngroup_lsntransportprofile_bindingFunc,
		ReadContext:   readLsngroup_lsntransportprofile_bindingFunc,
		DeleteContext: deleteLsngroup_lsntransportprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transportprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createLsngroup_lsntransportprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	groupname := d.Get("groupname")
	transportprofilename := d.Get("transportprofilename")
	bindingId := fmt.Sprintf("%s,%s", groupname, transportprofilename)
	lsngroup_lsntransportprofile_binding := lsn.Lsngrouplsntransportprofilebinding{
		Groupname:            d.Get("groupname").(string),
		Transportprofilename: d.Get("transportprofilename").(string),
	}

	err := client.UpdateUnnamedResource("lsngroup_lsntransportprofile_binding", &lsngroup_lsntransportprofile_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readLsngroup_lsntransportprofile_bindingFunc(ctx, d, meta)
}

func readLsngroup_lsntransportprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	groupname := idSlice[0]
	transportprofilename := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup_lsntransportprofile_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lsngroup_lsntransportprofile_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsntransportprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["transportprofilename"].(string) == transportprofilename {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams transportprofilename not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup_lsntransportprofile_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("groupname", data["groupname"])
	d.Set("transportprofilename", data["transportprofilename"])

	return nil

}

func deleteLsngroup_lsntransportprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroup_lsntransportprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	transportprofilename := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("transportprofilename:%s", transportprofilename))

	err := client.DeleteResourceWithArgs("lsngroup_lsntransportprofile_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
