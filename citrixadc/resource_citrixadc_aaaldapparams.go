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

func resourceCitrixAdcAaaldapparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAaaldapparamsFunc,
		ReadContext:   readAaaldapparamsFunc,
		UpdateContext: updateAaaldapparamsFunc,
		DeleteContext: deleteAaaldapparamsFunc,
		Schema: map[string]*schema.Schema{
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnameidentifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchfilter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchsubattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddnpassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldaploginname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxnestinglevel": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nestedgroupextraction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"searchfilter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sectype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssonameattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subattributename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaldapparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaldapparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaldapparamsName := resource.PrefixedUniqueId("tf-aaaldapparams-")

	aaaldapparams := aaa.Aaaldapparams{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Groupnameidentifier:        d.Get("groupnameidentifier").(string),
		Groupsearchattribute:       d.Get("groupsearchattribute").(string),
		Groupsearchfilter:          d.Get("groupsearchfilter").(string),
		Groupsearchsubattribute:    d.Get("groupsearchsubattribute").(string),
		Ldapbase:                   d.Get("ldapbase").(string),
		Ldapbinddn:                 d.Get("ldapbinddn").(string),
		Ldapbinddnpassword:         d.Get("ldapbinddnpassword").(string),
		Ldaploginname:              d.Get("ldaploginname").(string),
		Nestedgroupextraction:      d.Get("nestedgroupextraction").(string),
		Passwdchange:               d.Get("passwdchange").(string),
		Searchfilter:               d.Get("searchfilter").(string),
		Sectype:                    d.Get("sectype").(string),
		Serverip:                   d.Get("serverip").(string),
		Ssonameattribute:           d.Get("ssonameattribute").(string),
		Subattributename:           d.Get("subattributename").(string),
		Svrtype:                    d.Get("svrtype").(string),
	}

	if raw := d.GetRawConfig().GetAttr("authtimeout"); !raw.IsNull() {
		aaaldapparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxnestinglevel"); !raw.IsNull() {
		aaaldapparams.Maxnestinglevel = intPtr(d.Get("maxnestinglevel").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		aaaldapparams.Serverport = intPtr(d.Get("serverport").(int))
	}

	err := client.UpdateUnnamedResource(service.Aaaldapparams.Type(), &aaaldapparams)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(aaaldapparamsName)

	return readAaaldapparamsFunc(ctx, d, meta)
}

func readAaaldapparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaldapparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaldapparams state")
	data, err := client.FindResource(service.Aaaldapparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaldapparams state")
		d.SetId("")
		return nil
	}
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("groupnameidentifier", data["groupnameidentifier"])
	d.Set("groupsearchattribute", data["groupsearchattribute"])
	d.Set("groupsearchfilter", data["groupsearchfilter"])
	d.Set("groupsearchsubattribute", data["groupsearchsubattribute"])
	d.Set("ldapbase", data["ldapbase"])
	d.Set("ldapbinddn", data["ldapbinddn"])
	d.Set("ldapbinddnpassword", data["ldapbinddnpassword"])
	d.Set("ldaploginname", data["ldaploginname"])
	setToInt("maxnestinglevel", d, data["maxnestinglevel"])
	d.Set("nestedgroupextraction", data["nestedgroupextraction"])
	d.Set("passwdchange", data["passwdchange"])
	d.Set("searchfilter", data["searchfilter"])
	d.Set("sectype", data["sectype"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("ssonameattribute", data["ssonameattribute"])
	d.Set("subattributename", data["subattributename"])
	d.Set("svrtype", data["svrtype"])

	return nil

}

func updateAaaldapparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaldapparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaldapparams := aaa.Aaaldapparams{}
	hasChange := false
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for aaaldapparams, starting update")
		aaaldapparams.Authtimeout = intPtr(d.Get("authtimeout").(int))
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaaldapparams, starting update")
		aaaldapparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for aaaldapparams, starting update")
		aaaldapparams.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("groupnameidentifier") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnameidentifier has changed for aaaldapparams, starting update")
		aaaldapparams.Groupnameidentifier = d.Get("groupnameidentifier").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchattribute has changed for aaaldapparams, starting update")
		aaaldapparams.Groupsearchattribute = d.Get("groupsearchattribute").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchfilter has changed for aaaldapparams, starting update")
		aaaldapparams.Groupsearchfilter = d.Get("groupsearchfilter").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchsubattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchsubattribute has changed for aaaldapparams, starting update")
		aaaldapparams.Groupsearchsubattribute = d.Get("groupsearchsubattribute").(string)
		hasChange = true
	}
	if d.HasChange("ldapbase") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbase has changed for aaaldapparams, starting update")
		aaaldapparams.Ldapbase = d.Get("ldapbase").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddn") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddn has changed for aaaldapparams, starting update")
		aaaldapparams.Ldapbinddn = d.Get("ldapbinddn").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddnpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddnpassword has changed for aaaldapparams, starting update")
		aaaldapparams.Ldapbinddnpassword = d.Get("ldapbinddnpassword").(string)
		hasChange = true
	}
	if d.HasChange("ldaploginname") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldaploginname has changed for aaaldapparams, starting update")
		aaaldapparams.Ldaploginname = d.Get("ldaploginname").(string)
		hasChange = true
	}
	if d.HasChange("maxnestinglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnestinglevel has changed for aaaldapparams, starting update")
		aaaldapparams.Maxnestinglevel = intPtr(d.Get("maxnestinglevel").(int))
		hasChange = true
	}
	if d.HasChange("nestedgroupextraction") {
		log.Printf("[DEBUG]  citrixadc-provider: Nestedgroupextraction has changed for aaaldapparams, starting update")
		aaaldapparams.Nestedgroupextraction = d.Get("nestedgroupextraction").(string)
		hasChange = true
	}
	if d.HasChange("passwdchange") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdchange has changed for aaaldapparams, starting update")
		aaaldapparams.Passwdchange = d.Get("passwdchange").(string)
		hasChange = true
	}
	if d.HasChange("searchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Searchfilter has changed for aaaldapparams, starting update")
		aaaldapparams.Searchfilter = d.Get("searchfilter").(string)
		hasChange = true
	}
	if d.HasChange("sectype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sectype has changed for aaaldapparams, starting update")
		aaaldapparams.Sectype = d.Get("sectype").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for aaaldapparams, starting update")
		aaaldapparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for aaaldapparams, starting update")
		aaaldapparams.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("ssonameattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssonameattribute has changed for aaaldapparams, starting update")
		aaaldapparams.Ssonameattribute = d.Get("ssonameattribute").(string)
		hasChange = true
	}
	if d.HasChange("subattributename") {
		log.Printf("[DEBUG]  citrixadc-provider: Subattributename has changed for aaaldapparams, starting update")
		aaaldapparams.Subattributename = d.Get("subattributename").(string)
		hasChange = true
	}
	if d.HasChange("svrtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtype has changed for aaaldapparams, starting update")
		aaaldapparams.Svrtype = d.Get("svrtype").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaaldapparams.Type(), &aaaldapparams)
		if err != nil {
			return diag.Errorf("Error updating aaaldapparams")
		}
	}
	return readAaaldapparamsFunc(ctx, d, meta)
}

func deleteAaaldapparamsFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaldapparamsFunc")
	// aaaldapparams does not support DELETE operation
	d.SetId("")

	return nil
}
