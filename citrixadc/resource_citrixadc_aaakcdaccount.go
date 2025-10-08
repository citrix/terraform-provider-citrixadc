package citrixadc

import (
	"context"
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcAaakcdaccount() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaakcdaccountFunc,
		ReadContext:   readAaakcdaccountFunc,
		UpdateContext: updateAaakcdaccountFunc,
		DeleteContext: deleteAaakcdaccountFunc,
		Schema: map[string]*schema.Schema{
			"kcdaccount": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cacert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delegateduser": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enterpriserealm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdpassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keytab": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realmstr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicespn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usercert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userrealm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaakcdaccountFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaakcdaccountFunc")
	client := meta.(*NetScalerNitroClient).client
	aaakcdaccountName := d.Get("kcdaccount").(string)

	aaakcdaccount := aaa.Aaakcdaccount{
		Cacert:          d.Get("cacert").(string),
		Delegateduser:   d.Get("delegateduser").(string),
		Enterpriserealm: d.Get("enterpriserealm").(string),
		Kcdaccount:      d.Get("kcdaccount").(string),
		Kcdpassword:     d.Get("kcdpassword").(string),
		Keytab:          d.Get("keytab").(string),
		Realmstr:        d.Get("realmstr").(string),
		Servicespn:      d.Get("servicespn").(string),
		Usercert:        d.Get("usercert").(string),
		Userrealm:       d.Get("userrealm").(string),
	}

	_, err := client.AddResource(service.Aaakcdaccount.Type(), aaakcdaccountName, &aaakcdaccount)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaakcdaccountName)

	return readAaakcdaccountFunc(ctx, d, meta)
}

func readAaakcdaccountFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaakcdaccountFunc")
	client := meta.(*NetScalerNitroClient).client
	aaakcdaccountName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading aaakcdaccount state %s", aaakcdaccountName)
	data, err := client.FindResource(service.Aaakcdaccount.Type(), aaakcdaccountName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaakcdaccount state %s", aaakcdaccountName)
		d.SetId("")
		return nil
	}
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("cacert", data["cacert"])
	d.Set("delegateduser", data["delegateduser"])
	d.Set("enterpriserealm", data["enterpriserealm"])
	d.Set("kcdaccount", data["kcdaccount"])
	//d.Set("kcdpassword", data["kcdpassword"])
	d.Set("keytab", data["keytab"])
	//d.Set("realmstr", data["realmstr"])
	d.Set("servicespn", data["servicespn"])
	d.Set("usercert", data["usercert"])
	d.Set("userrealm", data["userrealm"])

	return nil

}

func updateAaakcdaccountFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaakcdaccountFunc")
	client := meta.(*NetScalerNitroClient).client
	aaakcdaccountName := d.Get("kcdaccount").(string)

	aaakcdaccount := aaa.Aaakcdaccount{
		Kcdaccount: d.Get("kcdaccount").(string),
	}
	hasChange := false
	if d.HasChange("cacert") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacert has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Cacert = d.Get("cacert").(string)
		hasChange = true
	}
	if d.HasChange("delegateduser") {
		log.Printf("[DEBUG]  citrixadc-provider: Delegateduser has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Delegateduser = d.Get("delegateduser").(string)
		hasChange = true
	}
	if d.HasChange("enterpriserealm") {
		log.Printf("[DEBUG]  citrixadc-provider: Enterpriserealm has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Enterpriserealm = d.Get("enterpriserealm").(string)
		hasChange = true
	}
	if d.HasChange("kcdpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdpassword has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Kcdpassword = d.Get("kcdpassword").(string)
		hasChange = true
	}
	if d.HasChange("keytab") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytab has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Keytab = d.Get("keytab").(string)
		hasChange = true
	}
	if d.HasChange("realmstr") {
		log.Printf("[DEBUG]  citrixadc-provider: Realmstr has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Realmstr = d.Get("realmstr").(string)
		hasChange = true
	}
	if d.HasChange("servicespn") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicespn has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Servicespn = d.Get("servicespn").(string)
		hasChange = true
	}
	if d.HasChange("usercert") {
		log.Printf("[DEBUG]  citrixadc-provider: Usercert has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Usercert = d.Get("usercert").(string)
		hasChange = true
	}
	if d.HasChange("userrealm") {
		log.Printf("[DEBUG]  citrixadc-provider: Userrealm has changed for aaakcdaccount %s, starting update", aaakcdaccountName)
		aaakcdaccount.Userrealm = d.Get("userrealm").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaakcdaccount.Type(), &aaakcdaccount)
		if err != nil {
			return diag.Errorf("Error updating aaakcdaccount %s", aaakcdaccountName)
		}
	}
	return readAaakcdaccountFunc(ctx, d, meta)
}

func deleteAaakcdaccountFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaakcdaccountFunc")
	client := meta.(*NetScalerNitroClient).client
	aaakcdaccountName := d.Id()
	err := client.DeleteResource(service.Aaakcdaccount.Type(), aaakcdaccountName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
