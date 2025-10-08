package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAppqoeparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAppqoeparameterFunc,
		ReadContext:   readAppqoeparameterFunc,
		UpdateContext: updateAppqoeparameterFunc,
		DeleteContext: deleteAppqoeparameterFunc,
		Schema: map[string]*schema.Schema{
			"avgwaitingclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dosattackthresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxaltrespbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionlife": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppqoeparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoeparameterName := resource.PrefixedUniqueId("tf-appqoeparameter-")
	appqoeparameter := make(map[string]interface{})

	if v, ok := d.GetOkExists("avgwaitingclient"); ok {
		appqoeparameter["avgwaitingclient"] = v.(int)
	}
	if v, ok := d.GetOkExists("dosattackthresh"); ok {
		appqoeparameter["dosattackthresh"] = v.(int)
	}
	if v, ok := d.GetOk("maxaltrespbandwidth"); ok {
		appqoeparameter["maxaltrespbandwidth"] = v.(int)
	}
	if v, ok := d.GetOk("sessionlife"); ok {
		appqoeparameter["sessionlife"] = v.(int)
	}

	err := client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(appqoeparameterName)

	return readAppqoeparameterFunc(ctx, d, meta)
}

func readAppqoeparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading appqoeparameter state")
	data, err := client.FindResource(service.Appqoeparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appqoeparameter state")
		d.SetId("")
		return nil
	}
	setToInt("avgwaitingclient", d, data["avgwaitingclient"])
	setToInt("dosattackthresh", d, data["dosattackthresh"])
	setToInt("maxaltrespbandwidth", d, data["maxaltrespbandwidth"])
	setToInt("sessionlife", d, data["sessionlife"])

	return nil

}

func updateAppqoeparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppqoeparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	appqoeparameter := appqoe.Appqoeparameter{}
	hasChange := false
	if d.HasChange("avgwaitingclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Avgwaitingclient has changed for appqoeparameter, starting update")
		appqoeparameter.Avgwaitingclient = d.Get("avgwaitingclient").(int)
		hasChange = true
	}
	if d.HasChange("dosattackthresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dosattackthresh has changed for appqoeparameter, starting update")
		appqoeparameter.Dosattackthresh = d.Get("dosattackthresh").(int)
		hasChange = true
	}
	if d.HasChange("maxaltrespbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxaltrespbandwidth has changed for appqoeparameter, starting update")
		appqoeparameter.Maxaltrespbandwidth = d.Get("maxaltrespbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("sessionlife") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlife has changed for appqoeparameter, starting update")
		appqoeparameter.Sessionlife = d.Get("sessionlife").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
		if err != nil {
			return diag.Errorf("Error updating appqoeparameter")
		}
	}
	return readAppqoeparameterFunc(ctx, d, meta)
}

func deleteAppqoeparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppqoeparameterFunc")
	//appqoeparameter does not suppor DELETE operation
	d.SetId("")

	return nil
}
