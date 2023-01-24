package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaaradiusparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaaradiusparamsFunc,
		Read:          readAaaradiusparamsFunc,
		Update:        updateAaaradiusparamsFunc,
		Delete:        deleteAaaradiusparamsFunc,
		Schema: map[string]*schema.Schema{
			"radkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authservretry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"callingstationid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipattributetype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ipvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passencoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwdattributetype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pwdvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radattributetype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"radgroupseparator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radgroupsprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radnasip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tunnelendpointclientip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaradiusparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaradiusparamsName := resource.PrefixedUniqueId("tf-aaaradiusparams-")
	
	aaaradiusparams := aaa.Aaaradiusparams{
		Accounting:                 d.Get("accounting").(string),
		Authentication:             d.Get("authentication").(string),
		Authservretry:              d.Get("authservretry").(int),
		Authtimeout:                d.Get("authtimeout").(int),
		Callingstationid:           d.Get("callingstationid").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Ipattributetype:            d.Get("ipattributetype").(int),
		Ipvendorid:                 d.Get("ipvendorid").(int),
		Passencoding:               d.Get("passencoding").(string),
		Pwdattributetype:           d.Get("pwdattributetype").(int),
		Pwdvendorid:                d.Get("pwdvendorid").(int),
		Radattributetype:           d.Get("radattributetype").(int),
		Radgroupseparator:          d.Get("radgroupseparator").(string),
		Radgroupsprefix:            d.Get("radgroupsprefix").(string),
		Radkey:                     d.Get("radkey").(string),
		Radnasid:                   d.Get("radnasid").(string),
		Radnasip:                   d.Get("radnasip").(string),
		Radvendorid:                d.Get("radvendorid").(int),
		Serverip:                   d.Get("serverip").(string),
		Serverport:                 d.Get("serverport").(int),
		Tunnelendpointclientip:     d.Get("tunnelendpointclientip").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaaradiusparams.Type(), &aaaradiusparams)
	if err != nil {
		return err
	}

	d.SetId(aaaradiusparamsName)

	err = readAaaradiusparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaaradiusparams but we can't read it ??")
		return nil
	}
	return nil
}

func readAaaradiusparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaradiusparams state")
	data, err := client.FindResource(service.Aaaradiusparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaradiusparams state")
		d.SetId("")
		return nil
	}
	d.Set("accounting", data["accounting"])
	d.Set("authentication", data["authentication"])
	d.Set("authservretry", data["authservretry"])
	d.Set("authtimeout", data["authtimeout"])
	d.Set("callingstationid", data["callingstationid"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("ipattributetype", data["ipattributetype"])
	d.Set("ipvendorid", data["ipvendorid"])
	d.Set("passencoding", data["passencoding"])
	d.Set("pwdattributetype", data["pwdattributetype"])
	d.Set("pwdvendorid", data["pwdvendorid"])
	d.Set("radattributetype", data["radattributetype"])
	d.Set("radgroupseparator", data["radgroupseparator"])
	d.Set("radgroupsprefix", data["radgroupsprefix"])
	//d.Set("radkey", data["radkey"])
	d.Set("radnasid", data["radnasid"])
	d.Set("radnasip", data["radnasip"])
	d.Set("radvendorid", data["radvendorid"])
	d.Set("serverip", data["serverip"])
	d.Set("serverport", data["serverport"])
	d.Set("tunnelendpointclientip", data["tunnelendpointclientip"])

	return nil

}

func updateAaaradiusparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaradiusparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaradiusparams := aaa.Aaaradiusparams{
		Radkey: d.Get("radkey").(string),
	}
	hasChange := false
	if d.HasChange("accounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Accounting has changed for aaaradiusparams, starting update")
		aaaradiusparams.Accounting = d.Get("accounting").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authservretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Authservretry has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authservretry = d.Get("authservretry").(int)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for aaaradiusparams, starting update")
		aaaradiusparams.Authtimeout = d.Get("authtimeout").(int)
		hasChange = true
	}
	if d.HasChange("callingstationid") {
		log.Printf("[DEBUG]  citrixadc-provider: Callingstationid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Callingstationid = d.Get("callingstationid").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for aaaradiusparams, starting update")
		aaaradiusparams.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("ipattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Ipattributetype = d.Get("ipattributetype").(int)
		hasChange = true
	}
	if d.HasChange("ipvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Ipvendorid = d.Get("ipvendorid").(int)
		hasChange = true
	}
	if d.HasChange("passencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Passencoding has changed for aaaradiusparams, starting update")
		aaaradiusparams.Passencoding = d.Get("passencoding").(string)
		hasChange = true
	}
	if d.HasChange("pwdattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Pwdattributetype = d.Get("pwdattributetype").(int)
		hasChange = true
	}
	if d.HasChange("pwdvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Pwdvendorid = d.Get("pwdvendorid").(int)
		hasChange = true
	}
	if d.HasChange("radattributetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Radattributetype has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radattributetype = d.Get("radattributetype").(int)
		hasChange = true
	}
	if d.HasChange("radgroupseparator") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupseparator has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radgroupseparator = d.Get("radgroupseparator").(string)
		hasChange = true
	}
	if d.HasChange("radgroupsprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Radgroupsprefix has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radgroupsprefix = d.Get("radgroupsprefix").(string)
		hasChange = true
	}
	if d.HasChange("radnasid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radnasid = d.Get("radnasid").(string)
		hasChange = true
	}
	if d.HasChange("radnasip") {
		log.Printf("[DEBUG]  citrixadc-provider: Radnasip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radnasip = d.Get("radnasip").(string)
		hasChange = true
	}
	if d.HasChange("radvendorid") {
		log.Printf("[DEBUG]  citrixadc-provider: Radvendorid has changed for aaaradiusparams, starting update")
		aaaradiusparams.Radvendorid = d.Get("radvendorid").(int)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for aaaradiusparams, starting update")
		aaaradiusparams.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("tunnelendpointclientip") {
		log.Printf("[DEBUG]  citrixadc-provider: Tunnelendpointclientip has changed for aaaradiusparams, starting update")
		aaaradiusparams.Tunnelendpointclientip = d.Get("tunnelendpointclientip").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaaradiusparams.Type(), &aaaradiusparams)
		if err != nil {
			return fmt.Errorf("Error updating aaaradiusparams")
		}
	}
	return readAaaradiusparamsFunc(d, meta)
}

func deleteAaaradiusparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaradiusparamsFunc")
	// aaaradiusparams does not support delete operation
	d.SetId("")

	return nil
}
