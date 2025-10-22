package citrixadc

import (
	"context"

	"log"

	"github.com/citrix/adc-nitro-go/resource/config/quicbridge"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcQuicbridgeprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createQuicbridgeprofileFunc,
		ReadContext:   readQuicbridgeprofileFunc,
		UpdateContext: updateQuicbridgeprofileFunc,
		DeleteContext: deleteQuicbridgeprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"routingalgorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serveridlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createQuicbridgeprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Get("name").(string)

	quicbridgeprofile := quicbridge.Quicbridgeprofile{
		Name:             d.Get("name").(string),
		Routingalgorithm: d.Get("routingalgorithm").(string),
	}
	if raw := d.GetRawConfig().GetAttr("serveridlength"); !raw.IsNull() {
		quicbridgeprofile.Serveridlength = intPtr(d.Get("serveridlength").(int))
	}

	_, err := client.AddResource("quicbridgeprofile", quicbridgeprofileName, &quicbridgeprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(quicbridgeprofileName)

	return readQuicbridgeprofileFunc(ctx, d, meta)
}

func readQuicbridgeprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading quicbridgeprofile state %s", quicbridgeprofileName)
	data, err := client.FindResource("quicbridgeprofile", quicbridgeprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing quicbridgeprofile state %s", quicbridgeprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("routingalgorithm", data["routingalgorithm"])
	setToInt("serveridlength", d, data["serveridlength"])

	return nil

}

func updateQuicbridgeprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Get("name").(string)

	quicbridgeprofile := quicbridge.Quicbridgeprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("routingalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Routingalgorithm has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Routingalgorithm = d.Get("routingalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("serveridlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Serveridlength has changed for quicbridgeprofile %s, starting update", quicbridgeprofileName)
		quicbridgeprofile.Serveridlength = intPtr(d.Get("serveridlength").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("quicbridgeprofile", quicbridgeprofileName, &quicbridgeprofile)
		if err != nil {
			return diag.Errorf("Error updating quicbridgeprofile %s", quicbridgeprofileName)
		}
	}
	return readQuicbridgeprofileFunc(ctx, d, meta)
}

func deleteQuicbridgeprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteQuicbridgeprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	quicbridgeprofileName := d.Id()
	err := client.DeleteResource("quicbridgeprofile", quicbridgeprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
