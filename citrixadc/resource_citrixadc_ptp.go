package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcPtp() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createPtpFunc,
		ReadContext:   readPtpFunc,
		UpdateContext: updatePtpFunc,
		DeleteContext: deletePtpFunc,
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createPtpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createPtpFunc")
	client := meta.(*NetScalerNitroClient).client
	var ptpName string
	// there is no primary key in ptp resource. Hence generate one for terraform state maintenance
	ptpName = resource.PrefixedUniqueId("tf-ptp-")
	ptp := network.Ptp{
		State: d.Get("state").(string),
	}

	err := client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ptpName)

	return readPtpFunc(ctx, d, meta)
}

func readPtpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readPtpFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading ptp state")
	data, err := client.FindResource(service.Ptp.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing ptp state")
		d.SetId("")
		return nil
	}
	d.Set("state", data["state"])

	return nil

}

func updatePtpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updatePtpFunc")
	client := meta.(*NetScalerNitroClient).client

	ptp := network.Ptp{}
	hasChange := false
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for ptp, starting update")
		ptp.State = d.Get("state").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Ptp.Type(), &ptp)
		if err != nil {
			return diag.Errorf("Error updating ptp ")
		}
	}
	return readPtpFunc(ctx, d, meta)
}

func deletePtpFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deletePtpFunc")

	d.SetId("")

	return nil
}
