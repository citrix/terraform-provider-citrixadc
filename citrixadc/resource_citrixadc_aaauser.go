package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAaauser() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaauserFunc,
		ReadContext:   readAaauserFunc,
		UpdateContext: updateAaauserFunc,
		DeleteContext: deleteAaauserFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loggedin": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaauserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaauserFunc")
	client := meta.(*NetScalerNitroClient).client
	aaauserName := d.Get("username").(string)

	aaauser := aaa.Aaauser{
		Loggedin: d.Get("loggedin").(bool),
		Password: d.Get("password").(string),
		Username: d.Get("username").(string),
	}

	_, err := client.AddResource(service.Aaauser.Type(), aaauserName, &aaauser)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaauserName)

	return readAaauserFunc(ctx, d, meta)
}

func readAaauserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaauserFunc")
	client := meta.(*NetScalerNitroClient).client
	aaauserName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaauser state %s", aaauserName)
	data, err := client.FindResource(service.Aaauser.Type(), aaauserName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaauser state %s", aaauserName)
		d.SetId("")
		return nil
	}
	d.Set("username", data["username"])
	d.Set("loggedin", data["loggedin"])
	//d.Set("password", data["password"])

	return nil

}

func updateAaauserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaauserFunc")
	client := meta.(*NetScalerNitroClient).client
	aaauserName := d.Get("username").(string)

	aaauser := aaa.Aaauser{
		Username: d.Get("username").(string),
	}
	hasChange := false
	if d.HasChange("loggedin") {
		log.Printf("[DEBUG]  citrixadc-provider: Loggedin has changed for aaauser %s, starting update", aaauserName)
		aaauser.Loggedin = d.Get("loggedin").(bool)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for aaauser %s, starting update", aaauserName)
		aaauser.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for aaauser %s, starting update", aaauserName)
		aaauser.Username = d.Get("username").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaauser.Type(), &aaauser)
		if err != nil {
			return diag.Errorf("Error updating aaauser %s", aaauserName)
		}
	}
	return readAaauserFunc(ctx, d, meta)
}

func deleteAaauserFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaauserFunc")
	client := meta.(*NetScalerNitroClient).client
	aaauserName := d.Id()
	err := client.DeleteResource(service.Aaauser.Type(), aaauserName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
