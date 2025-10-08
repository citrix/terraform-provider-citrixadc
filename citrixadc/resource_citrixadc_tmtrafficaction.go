package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTmtrafficaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmtrafficactionFunc,
		ReadContext:   readTmtrafficactionFunc,
		UpdateContext: updateTmtrafficactionFunc,
		DeleteContext: deleteTmtrafficactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apptimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"forcedtimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forcedtimeoutval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"formssoaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initiatelogout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlssoprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createTmtrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmtrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficactionName := d.Get("name").(string)

	tmtrafficaction := tm.Tmtrafficaction{
		Apptimeout:       d.Get("apptimeout").(int),
		Forcedtimeout:    d.Get("forcedtimeout").(string),
		Forcedtimeoutval: d.Get("forcedtimeoutval").(int),
		Formssoaction:    d.Get("formssoaction").(string),
		Initiatelogout:   d.Get("initiatelogout").(string),
		Kcdaccount:       d.Get("kcdaccount").(string),
		Name:             d.Get("name").(string),
		Passwdexpression: d.Get("passwdexpression").(string),
		Persistentcookie: d.Get("persistentcookie").(string),
		Samlssoprofile:   d.Get("samlssoprofile").(string),
		Sso:              d.Get("sso").(string),
		Userexpression:   d.Get("userexpression").(string),
	}

	_, err := client.AddResource(service.Tmtrafficaction.Type(), tmtrafficactionName, &tmtrafficaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tmtrafficactionName)

	return readTmtrafficactionFunc(ctx, d, meta)
}

func readTmtrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmtrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmtrafficaction state %s", tmtrafficactionName)
	data, err := client.FindResource(service.Tmtrafficaction.Type(), tmtrafficactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmtrafficaction state %s", tmtrafficactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	setToInt("apptimeout", d, data["apptimeout"])
	d.Set("forcedtimeout", data["forcedtimeout"])
	setToInt("forcedtimeoutval", d, data["forcedtimeoutval"])
	d.Set("formssoaction", data["formssoaction"])
	d.Set("initiatelogout", data["initiatelogout"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("passwdexpression", data["passwdexpression"])
	d.Set("persistentcookie", data["persistentcookie"])
	d.Set("samlssoprofile", data["samlssoprofile"])
	d.Set("sso", data["sso"])
	d.Set("userexpression", data["userexpression"])

	return nil

}

func updateTmtrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmtrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficactionName := d.Get("name").(string)

	tmtrafficaction := tm.Tmtrafficaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("apptimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Apptimeout has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Apptimeout = d.Get("apptimeout").(int)
		hasChange = true
	}
	if d.HasChange("forcedtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeout has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Forcedtimeout = d.Get("forcedtimeout").(string)
		hasChange = true
	}
	if d.HasChange("forcedtimeoutval") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeoutval has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Forcedtimeoutval = d.Get("forcedtimeoutval").(int)
		hasChange = true
	}
	if d.HasChange("formssoaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Formssoaction has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Formssoaction = d.Get("formssoaction").(string)
		hasChange = true
	}
	if d.HasChange("initiatelogout") {
		log.Printf("[DEBUG]  citrixadc-provider: Initiatelogout has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Initiatelogout = d.Get("initiatelogout").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("passwdexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdexpression has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Passwdexpression = d.Get("passwdexpression").(string)
		hasChange = true
	}
	if d.HasChange("persistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentcookie has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Persistentcookie = d.Get("persistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("samlssoprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlssoprofile has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Samlssoprofile = d.Get("samlssoprofile").(string)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("userexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Userexpression has changed for tmtrafficaction %s, starting update", tmtrafficactionName)
		tmtrafficaction.Userexpression = d.Get("userexpression").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmtrafficaction.Type(), &tmtrafficaction)
		if err != nil {
			return diag.Errorf("Error updating tmtrafficaction %s", tmtrafficactionName)
		}
	}
	return readTmtrafficactionFunc(ctx, d, meta)
}

func deleteTmtrafficactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmtrafficactionFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficactionName := d.Id()
	err := client.DeleteResource(service.Tmtrafficaction.Type(), tmtrafficactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
