package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/subscriber"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSubscriberradiusinterface() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSubscriberradiusinterfaceFunc,
		ReadContext:   readSubscriberradiusinterfaceFunc,
		UpdateContext: updateSubscriberradiusinterfaceFunc,
		DeleteContext: deleteSubscriberradiusinterfaceFunc,
		Schema: map[string]*schema.Schema{
			"listeningservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radiusinterimasstart": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSubscriberradiusinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	subscriberradiusinterfaceName := resource.PrefixedUniqueId("tf-subscriberradiusinterface-")

	subscriberradiusinterface := subscriber.Subscriberradiusinterface{
		Listeningservice:     d.Get("listeningservice").(string),
		Radiusinterimasstart: d.Get("radiusinterimasstart").(string),
	}

	err := client.UpdateUnnamedResource("subscriberradiusinterface", &subscriberradiusinterface)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(subscriberradiusinterfaceName)

	return readSubscriberradiusinterfaceFunc(ctx, d, meta)
}

func readSubscriberradiusinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading subscriberradiusinterface state")
	data, err := client.FindResource("subscriberradiusinterface", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing subscriberradiusinterface state")
		d.SetId("")
		return nil
	}
	d.Set("listeningservice", data["listeningservice"])
	d.Set("radiusinterimasstart", data["radiusinterimasstart"])

	return nil

}

func updateSubscriberradiusinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSubscriberradiusinterfaceFunc")
	client := meta.(*NetScalerNitroClient).client

	subscriberradiusinterface := subscriber.Subscriberradiusinterface{}
	hasChange := false
	if d.HasChange("listeningservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Listeningservice has changed for subscriberradiusinterface, starting update")
		subscriberradiusinterface.Listeningservice = d.Get("listeningservice").(string)
		hasChange = true
	}
	if d.HasChange("radiusinterimasstart") {
		log.Printf("[DEBUG]  citrixadc-provider: Radiusinterimasstart has changed for subscriberradiusinterface, starting update")
		subscriberradiusinterface.Radiusinterimasstart = d.Get("radiusinterimasstart").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("subscriberradiusinterface", &subscriberradiusinterface)
		if err != nil {
			return diag.Errorf("Error updating subscriberradiusinterface")
		}
	}
	return readSubscriberradiusinterfaceFunc(ctx, d, meta)
}

func deleteSubscriberradiusinterfaceFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSubscriberradiusinterfaceFunc")
	//subscriberradiusinterface does not support DELETE operation
	d.SetId("")

	return nil
}
