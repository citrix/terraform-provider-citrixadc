package citrixadc

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/analytics"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAnalyticsglobal_analyticsprofile_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAnalyticsglobal_analyticsprofile_bindingFunc,
		ReadContext:   readAnalyticsglobal_analyticsprofile_bindingFunc,
		DeleteContext: deleteAnalyticsglobal_analyticsprofile_bindingFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"analyticsprofile": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createAnalyticsglobal_analyticsprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	analyticsprofile := d.Get("analyticsprofile").(string)

	analyticsglobal_analyticsprofile_binding := analytics.Analyticsglobalanalyticsprofilebinding{
		Analyticsprofile: d.Get("analyticsprofile").(string),
	}

	err := client.UpdateUnnamedResource("analyticsglobal_analyticsprofile_binding", &analyticsglobal_analyticsprofile_binding)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(analyticsprofile)

	return readAnalyticsglobal_analyticsprofile_bindingFunc(ctx, d, meta)
}

func readAnalyticsglobal_analyticsprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofile := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)

	findParams := service.FindParams{
		ResourceType:             "analyticsglobal",
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
		log.Printf("[WARN] citrixadc-provider: Clearing analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["analyticsprofile"].(string) == analyticsprofile {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing analyticsglobal_analyticsprofile_binding state %s", analyticsprofile)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("analyticsprofile", data["analyticsprofile"])

	return nil

}

func deleteAnalyticsglobal_analyticsprofile_bindingFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAnalyticsglobal_analyticsprofile_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	analyticsprofile := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("analyticsprofile:%s", analyticsprofile))

	err := client.DeleteResourceWithArgs("analyticsglobal_analyticsprofile_binding", "", args)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
