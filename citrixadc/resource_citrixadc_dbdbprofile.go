package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcDbdbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createDbdbprofileFunc,
		ReadContext:   readDbdbprofileFunc,
		UpdateContext: updateDbdbprofileFunc,
		DeleteContext: deleteDbdbprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"conmultiplex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablecachingconmuxoff": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interpretquery": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stickiness": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDbdbprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Get("name").(string)
	dbdbprofile := db.Dbdbprofile{
		Conmultiplex:           d.Get("conmultiplex").(string),
		Enablecachingconmuxoff: d.Get("enablecachingconmuxoff").(string),
		Interpretquery:         d.Get("interpretquery").(string),
		Kcdaccount:             d.Get("kcdaccount").(string),
		Name:                   d.Get("name").(string),
		Stickiness:             d.Get("stickiness").(string),
	}

	_, err := client.AddResource(service.Dbdbprofile.Type(), dbdbprofileName, &dbdbprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dbdbprofileName)

	return readDbdbprofileFunc(ctx, d, meta)
}

func readDbdbprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dbdbprofile state %s", dbdbprofileName)
	data, err := client.FindResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dbdbprofile state %s", dbdbprofileName)
		d.SetId("")
		return nil
	}
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("enablecachingconmuxoff", data["enablecachingconmuxoff"])
	d.Set("interpretquery", data["interpretquery"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("name", data["name"])
	d.Set("stickiness", data["stickiness"])

	return nil

}

func updateDbdbprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Get("name").(string)

	dbdbprofile := db.Dbdbprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("conmultiplex") {
		log.Printf("[DEBUG]  citrixadc-provider: Conmultiplex has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Conmultiplex = d.Get("conmultiplex").(string)
		hasChange = true
	}
	if d.HasChange("enablecachingconmuxoff") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablecachingconmuxoff has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Enablecachingconmuxoff = d.Get("enablecachingconmuxoff").(string)
		hasChange = true
	}
	if d.HasChange("interpretquery") {
		log.Printf("[DEBUG]  citrixadc-provider: Interpretquery has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Interpretquery = d.Get("interpretquery").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("stickiness") {
		log.Printf("[DEBUG]  citrixadc-provider: Stickiness has changed for dbdbprofile %s, starting update", dbdbprofileName)
		dbdbprofile.Stickiness = d.Get("stickiness").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Dbdbprofile.Type(), &dbdbprofile)
		if err != nil {
			return diag.Errorf("Error updating dbdbprofile %s", dbdbprofileName)
		}
	}
	return readDbdbprofileFunc(ctx, d, meta)
}

func deleteDbdbprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDbdbprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dbdbprofileName := d.Id()
	err := client.DeleteResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
