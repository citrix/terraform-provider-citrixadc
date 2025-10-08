package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAaassoprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaassoprofileFunc,
		ReadContext:   readAaassoprofileFunc,
		UpdateContext: updateAaassoprofileFunc,
		DeleteContext: deleteAaassoprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createAaassoprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Get("name").(string)
	aaassoprofile := aaa.Aaassoprofile{
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
		Username: d.Get("username").(string),
	}

	_, err := client.AddResource("aaassoprofile", aaassoprofileName, &aaassoprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaassoprofileName)

	return readAaassoprofileFunc(ctx, d, meta)
}

func readAaassoprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaassoprofile state %s", aaassoprofileName)
	data, err := client.FindResource("aaassoprofile", aaassoprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaassoprofile state %s", aaassoprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	//d.Set("password", data["password"])
	d.Set("username", data["username"])

	return nil

}

func updateAaassoprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Get("name").(string)

	aaassoprofile := aaa.Aaassoprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false

	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for aaassoprofile %s, starting update", aaassoprofileName)
		aaassoprofile.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for aaassoprofile %s, starting update", aaassoprofileName)
		aaassoprofile.Username = d.Get("username").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("aaassoprofile", &aaassoprofile)
		if err != nil {
			return diag.Errorf("Error updating aaassoprofile %s", aaassoprofileName)
		}
	}
	return readAaassoprofileFunc(ctx, d, meta)
}

func deleteAaassoprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaassoprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	aaassoprofileName := d.Id()
	err := client.DeleteResource("aaassoprofile", aaassoprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
