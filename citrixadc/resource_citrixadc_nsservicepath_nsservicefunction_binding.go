package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func resourceCitrixAdcNsservicepath_nsservicefunction_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsservicepath_nsservicefunction_bindingFunc,
		ReadContext:   readNsservicepath_nsservicefunction_bindingFunc,
		DeleteContext: deleteNsservicepath_nsservicefunction_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"servicepathname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"servicefunction": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
		},
	}
}

func createNsservicepath_nsservicefunction_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	servicepathname := d.Get("servicepathname")
	servicefunction := d.Get("servicefunction")
	bindingId := fmt.Sprintf("%s,%s", servicepathname, servicefunction)
	nsservicepath_nsservicefunction_binding := ns.Nsservicepathnsservicefunctionbinding{
		Servicefunction: d.Get("servicefunction").(string),
		Servicepathname: d.Get("servicepathname").(string),
	}

	if raw := d.GetRawConfig().GetAttr("index"); !raw.IsNull() {
		nsservicepath_nsservicefunction_binding.Index = intPtr(d.Get("index").(int))
	}

	err := client.UpdateUnnamedResource(service.Nsservicepath_nsservicefunction_binding.Type(), &nsservicepath_nsservicefunction_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readNsservicepath_nsservicefunction_bindingFunc(ctx, d, meta)
}

func readNsservicepath_nsservicefunction_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	servicepathname := idSlice[0]
	servicefunction := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading nsservicepath_nsservicefunction_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "nsservicepath_nsservicefunction_binding",
		ResourceName:             servicepathname,
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
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath_nsservicefunction_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["servicefunction"].(string) == servicefunction {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing nsservicepath_nsservicefunction_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	setToInt("index", d, data["index"])
	d.Set("servicefunction", data["servicefunction"])
	d.Set("servicepathname", data["servicepathname"])

	return nil

}

func deleteNsservicepath_nsservicefunction_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsservicepath_nsservicefunction_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	servicefunction := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("servicefunction:%s", servicefunction))

	err := client.DeleteResourceWithArgs(service.Nsservicepath_nsservicefunction_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
