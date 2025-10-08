package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcMapdomain_mapbmr_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createMapdomain_mapbmr_bindingFunc,
		ReadContext:   readMapdomain_mapbmr_bindingFunc,
		DeleteContext: deleteMapdomain_mapbmr_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"mapbmrname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createMapdomain_mapbmr_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	mapbmrname := d.Get("mapbmrname")
	bindingId := fmt.Sprintf("%s,%s", name, mapbmrname)
	mapdomain_mapbmr_binding := network.Mapdomainmapbmrbinding{
		Mapbmrname: d.Get("mapbmrname").(string),
		Name:       d.Get("name").(string),
	}

	err := client.UpdateUnnamedResource("mapdomain_mapbmr_binding", &mapdomain_mapbmr_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readMapdomain_mapbmr_bindingFunc(ctx, d, meta)
}

func readMapdomain_mapbmr_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	mapbmrname := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading mapdomain_mapbmr_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "mapdomain_mapbmr_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain_mapbmr_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["mapbmrname"].(string) == mapbmrname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing mapdomain_mapbmr_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("mapbmrname", data["mapbmrname"])
	d.Set("name", data["name"])

	return nil

}

func deleteMapdomain_mapbmr_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteMapdomain_mapbmr_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	mapbmrname := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("mapbmrname:%s", mapbmrname))

	err := client.DeleteResourceWithArgs("mapdomain_mapbmr_binding", name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
