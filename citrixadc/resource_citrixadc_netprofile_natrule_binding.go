package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcNetprofile_natrule_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNetprofile_natrule_bindingFunc,
		ReadContext:   readNetprofile_natrule_bindingFunc,
		DeleteContext: deleteNetprofile_natrule_bindingFunc,
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
			"natrule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rewriteip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNetprofile_natrule_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	name := d.Get("name")
	natrule := d.Get("natrule")
	bindingId := fmt.Sprintf("%s,%s", name, natrule)
	netprofile_natrule_binding := network.Netprofilenatrulebinding{
		Name:      d.Get("name").(string),
		Natrule:   d.Get("natrule").(string),
		Netmask:   d.Get("netmask").(string),
		Rewriteip: d.Get("rewriteip").(string),
	}

	err := client.UpdateUnnamedResource(service.Netprofile_natrule_binding.Type(), &netprofile_natrule_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bindingId)

	return readNetprofile_natrule_bindingFunc(ctx, d, meta)
}

func readNetprofile_natrule_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natrule := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading netprofile_natrule_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "netprofile_natrule_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_natrule_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["natrule"].(string) == natrule {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing netprofile_natrule_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("name", data["name"])
	d.Set("natrule", data["natrule"])
	d.Set("netmask", data["netmask"])
	d.Set("rewriteip", data["rewriteip"])

	return nil

}

func deleteNetprofile_natrule_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNetprofile_natrule_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	name := idSlice[0]
	natrule := idSlice[1]

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("natrule:%s", url.QueryEscape(natrule)))
	if val, ok := d.GetOk("netmask"); ok {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(val.(string))))
	}

	err := client.DeleteResourceWithArgs(service.Netprofile_natrule_binding.Type(), name, args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
