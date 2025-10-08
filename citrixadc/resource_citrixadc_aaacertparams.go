package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaacertparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaacertparamsFunc,
		ReadContext:   readAaacertparamsFunc,
		UpdateContext: updateAaacertparamsFunc,
		DeleteContext: deleteAaacertparamsFunc,
		Schema: map[string]*schema.Schema{
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnamefield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usernamefield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaacertparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaacertparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaacertparamsName := resource.PrefixedUniqueId("tf-aaacertparams-")

	aaacertparams := aaa.Aaacertparams{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupnamefield:             d.Get("groupnamefield").(string),
		Usernamefield:              d.Get("usernamefield").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaacertparams.Type(), &aaacertparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaacertparamsName)

	return readAaacertparamsFunc(ctx, d, meta)
}

func readAaacertparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaacertparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaacertparams state")
	data, err := client.FindResource(service.Aaacertparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaacertparams state")
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupnamefield", data["groupnamefield"])
	d.Set("usernamefield", data["usernamefield"])

	return nil

}

func updateAaacertparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaacertparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaacertparams := aaa.Aaacertparams{}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaacertparams, starting update")
		aaacertparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupnamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnamefield has changed for aaacertparams, starting update")
		aaacertparams.Groupnamefield = d.Get("groupnamefield").(string)
		hasChange = true
	}
	if d.HasChange("usernamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Usernamefield has changed for aaacertparams, starting update")
		aaacertparams.Usernamefield = d.Get("usernamefield").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaacertparams.Type(), &aaacertparams)
		if err != nil {
			return diag.Errorf("Error updating aaacertparams")
		}
	}
	return readAaacertparamsFunc(ctx, d, meta)
}

func deleteAaacertparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaacertparamsFunc")
	// aaacertparams does not support delete operations
	d.SetId("")

	return nil
}
