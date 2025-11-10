package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/citrix/adc-nitro-go/service"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"net/url"
)

func resourceCitrixAdcVpnglobal_domain_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnglobal_domain_bindingFunc,
		ReadContext:   readVpnglobal_domain_bindingFunc,
		DeleteContext: deleteVpnglobal_domain_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"intranetdomain": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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

func createVpnglobal_domain_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnglobal_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetdomain := d.Get("intranetdomain").(string)
	vpnglobal_domain_binding := vpn.Vpnglobaldomainbinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Intranetdomain:         d.Get("intranetdomain").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnglobal_domain_binding.Type(), &vpnglobal_domain_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(intranetdomain)

	return readVpnglobal_domain_bindingFunc(ctx, d, meta)
}

func readVpnglobal_domain_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnglobal_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	intranetdomain := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading vpnglobal_domain_binding state %s", intranetdomain)

	findParams := service.FindParams{
		ResourceType:             "vpnglobal_domain_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_domain_binding state %s", intranetdomain)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["intranetdomain"].(string) == intranetdomain {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing vpnglobal_domain_binding state %s", intranetdomain)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("intranetdomain", data["intranetdomain"])

	return nil

}

func deleteVpnglobal_domain_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnglobal_domain_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	intranetdomain := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("intranetdomain:%s", url.QueryEscape(intranetdomain)))
	err := client.DeleteResourceWithArgs(service.Vpnglobal_domain_binding.Type(), "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
