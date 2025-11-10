package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/smpp"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSmppuser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSmppuserFunc,
		ReadContext:   readSmppuserFunc,
		UpdateContext: updateSmppuserFunc,
		DeleteContext: deleteSmppuserFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSmppuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Get("username").(string)
	smppuser := smpp.Smppuser{
		Password: d.Get("password").(string),
		Username: d.Get("username").(string),
	}

	_, err := client.AddResource("smppuser", smppuserName, &smppuser)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(smppuserName)

	return readSmppuserFunc(ctx, d, meta)
}

func readSmppuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading smppuser state %s", smppuserName)
	data, err := client.FindResource("smppuser", smppuserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing smppuser state %s", smppuserName)
		d.SetId("")
		return nil
	}
	d.Set("username", data["username"])
	d.Set("username", data["username"])

	return nil

}

func updateSmppuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Get("username").(string)

	smppuser := smpp.Smppuser{
		Username: d.Get("username").(string),
	}
	hasChange := false
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for smppuser %s, starting update", smppuserName)
		smppuser.Password = d.Get("password").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("smppuser", &smppuser)
		if err != nil {
			return diag.Errorf("Error updating smppuser %s", smppuserName)
		}
	}
	return readSmppuserFunc(ctx, d, meta)
}

func deleteSmppuserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSmppuserFunc")
	client := meta.(*NetScalerNitroClient).client
	smppuserName := d.Id()
	err := client.DeleteResource("smppuser", smppuserName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
