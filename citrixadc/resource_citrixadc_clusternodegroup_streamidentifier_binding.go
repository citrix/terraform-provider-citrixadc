package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcClusternodegroup_streamidentifier_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createClusternodegroup_streamidentifier_bindingFunc,
		ReadContext:   readClusternodegroup_streamidentifier_bindingFunc,
		DeleteContext: deleteClusternodegroup_streamidentifier_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"identifiername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createClusternodegroup_streamidentifier_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	identifiername := d.Get("identifiername")
	bindingId := fmt.Sprintf("%s,%s", name, identifiername)
	clusternodegroup_streamidentifier_binding := cluster.Clusternodegroupstreamidentifierbinding{
		Identifiername: d.Get("identifiername").(string),
		Name:           d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readClusternodegroup_streamidentifier_bindingFunc(ctx, d, meta)
}

func readClusternodegroup_streamidentifier_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading clusternodegroup_streamidentifier_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "clusternodegroup_streamidentifier_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_streamidentifier_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["identifiername"].(string) == identifiername {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams identifiername not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing clusternodegroup_streamidentifier_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("identifiername", data["identifiername"])
	d.Set("name", data["name"])

	return nil

}

func deleteClusternodegroup_streamidentifier_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteClusternodegroup_streamidentifier_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	identifiername := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("identifiername:%s", identifiername))

	err := client.DeleteResourceWithArgs(service.Clusternodegroup_streamidentifier_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
