package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAaaparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAaaparameterFunc,
		Read:          readAaaparameterFunc,
		Update:        updateAaaparameterFunc,
		Delete:        deleteAaaparameterFunc,
		Schema: map[string]*schema.Schema{
			"aaadloglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aaadnatip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aaasessionloglevel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apitokencache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultcspheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynaddr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enableenhancedauthfeedback": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablesessionstickiness": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enablestaticpagecaching": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failedlogintimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ftmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loginencryption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxaaausers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxkbquestions": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxloginattempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxsamldeflatesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistentloginattempts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwdexpirynotificationdays": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samesite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tokenintrospectioninterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httponlycookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enhancedepa": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wafprotection": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"securityinsights": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAaaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAaaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	aaaparameterName := resource.PrefixedUniqueId("tf-aaaparameter-")

	aaaparameter := aaa.Aaaparameter{
		Aaadloglevel:               d.Get("aaadloglevel").(string),
		Aaadnatip:                  d.Get("aaadnatip").(string),
		Aaasessionloglevel:         d.Get("aaasessionloglevel").(string),
		Apitokencache:              d.Get("apitokencache").(string),
		Defaultauthtype:            d.Get("defaultauthtype").(string),
		Defaultcspheader:           d.Get("defaultcspheader").(string),
		Dynaddr:                    d.Get("dynaddr").(string),
		Enableenhancedauthfeedback: d.Get("enableenhancedauthfeedback").(string),
		Enablesessionstickiness:    d.Get("enablesessionstickiness").(string),
		Enablestaticpagecaching:    d.Get("enablestaticpagecaching").(string),
		Failedlogintimeout:         d.Get("failedlogintimeout").(int),
		Ftmode:                     d.Get("ftmode").(string),
		Loginencryption:            d.Get("loginencryption").(string),
		Maxaaausers:                d.Get("maxaaausers").(int),
		Maxkbquestions:             d.Get("maxkbquestions").(int),
		Maxloginattempts:           d.Get("maxloginattempts").(int),
		Maxsamldeflatesize:         d.Get("maxsamldeflatesize").(int),
		Persistentloginattempts:    d.Get("persistentloginattempts").(string),
		Pwdexpirynotificationdays:  d.Get("pwdexpirynotificationdays").(int),
		Samesite:                   d.Get("samesite").(string),
		Tokenintrospectioninterval: d.Get("tokenintrospectioninterval").(int),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Enhancedepa:                d.Get("enhancedepa").(string),
		Wafprotection:              toStringList(d.Get("wafprotection").([]interface{})),
		Securityinsights:           d.Get("securityinsights").(string),
	}

	err := client.UpdateUnnamedResource(service.Aaaparameter.Type(), &aaaparameter)
	if err != nil {
		return err
	}

	d.SetId(aaaparameterName)

	err = readAaaparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this aaaparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readAaaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAaaparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading aaaparameter state")
	data, err := client.FindResource(service.Aaaparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing aaaparameter state")
		d.SetId("")
		return nil
	}
	d.Set("aaadloglevel", data["aaadloglevel"])
	d.Set("aaadnatip", data["aaadnatip"])
	d.Set("aaasessionloglevel", data["aaasessionloglevel"])
	d.Set("apitokencache", data["apitokencache"])
	d.Set("defaultauthtype", data["defaultauthtype"])
	d.Set("defaultcspheader", data["defaultcspheader"])
	d.Set("dynaddr", data["dynaddr"])
	d.Set("enableenhancedauthfeedback", data["enableenhancedauthfeedback"])
	d.Set("enablesessionstickiness", data["enablesessionstickiness"])
	d.Set("enablestaticpagecaching", data["enablestaticpagecaching"])
	d.Set("failedlogintimeout", data["failedlogintimeout"])
	d.Set("ftmode", data["ftmode"])
	d.Set("loginencryption", data["loginencryption"])
	d.Set("maxaaausers", data["maxaaausers"])
	d.Set("maxkbquestions", data["maxkbquestions"])
	d.Set("maxloginattempts", data["maxloginattempts"])
	d.Set("maxsamldeflatesize", data["maxsamldeflatesize"])
	d.Set("persistentloginattempts", data["persistentloginattempts"])
	d.Set("pwdexpirynotificationdays", data["pwdexpirynotificationdays"])
	d.Set("samesite", data["samesite"])
	d.Set("tokenintrospectioninterval", data["tokenintrospectioninterval"])
	d.Set("httponlycookie", data["httponlycookie"])
	d.Set("enhancedepa", data["enhancedepa"])
	d.Set("wafprotection", data["wafprotection"])
	d.Set("securityinsights", data["securityinsights"])

	return nil

}

func updateAaaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAaaparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	aaaparameter := aaa.Aaaparameter{}
	hasChange := false
	if d.HasChange("aaadloglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Aaadloglevel has changed for aaaparameter, starting update")
		aaaparameter.Aaadloglevel = d.Get("aaadloglevel").(string)
		hasChange = true
	}
	if d.HasChange("aaadnatip") {
		log.Printf("[DEBUG]  citrixadc-provider: Aaadnatip has changed for aaaparameter, starting update")
		aaaparameter.Aaadnatip = d.Get("aaadnatip").(string)
		hasChange = true
	}
	if d.HasChange("aaasessionloglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Aaasessionloglevel has changed for aaaparameter, starting update")
		aaaparameter.Aaasessionloglevel = d.Get("aaasessionloglevel").(string)
		hasChange = true
	}
	if d.HasChange("apitokencache") {
		log.Printf("[DEBUG]  citrixadc-provider: Apitokencache has changed for aaaparameter, starting update")
		aaaparameter.Apitokencache = d.Get("apitokencache").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthtype has changed for aaaparameter, starting update")
		aaaparameter.Defaultauthtype = d.Get("defaultauthtype").(string)
		hasChange = true
	}
	if d.HasChange("defaultcspheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultcspheader has changed for aaaparameter, starting update")
		aaaparameter.Defaultcspheader = d.Get("defaultcspheader").(string)
		hasChange = true
	}
	if d.HasChange("dynaddr") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynaddr has changed for aaaparameter, starting update")
		aaaparameter.Dynaddr = d.Get("dynaddr").(string)
		hasChange = true
	}
	if d.HasChange("enableenhancedauthfeedback") {
		log.Printf("[DEBUG]  citrixadc-provider: Enableenhancedauthfeedback has changed for aaaparameter, starting update")
		aaaparameter.Enableenhancedauthfeedback = d.Get("enableenhancedauthfeedback").(string)
		hasChange = true
	}
	if d.HasChange("enablesessionstickiness") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablesessionstickiness has changed for aaaparameter, starting update")
		aaaparameter.Enablesessionstickiness = d.Get("enablesessionstickiness").(string)
		hasChange = true
	}
	if d.HasChange("enablestaticpagecaching") {
		log.Printf("[DEBUG]  citrixadc-provider: Enablestaticpagecaching has changed for aaaparameter, starting update")
		aaaparameter.Enablestaticpagecaching = d.Get("enablestaticpagecaching").(string)
		hasChange = true
	}
	if d.HasChange("failedlogintimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Failedlogintimeout has changed for aaaparameter, starting update")
		aaaparameter.Failedlogintimeout = d.Get("failedlogintimeout").(int)
		aaaparameter.Maxloginattempts = d.Get("maxloginattempts").(int)
		hasChange = true
	}
	if d.HasChange("ftmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftmode has changed for aaaparameter, starting update")
		aaaparameter.Ftmode = d.Get("ftmode").(string)
		hasChange = true
	}
	if d.HasChange("loginencryption") {
		log.Printf("[DEBUG]  citrixadc-provider: Loginencryption has changed for aaaparameter, starting update")
		aaaparameter.Loginencryption = d.Get("loginencryption").(string)
		hasChange = true
	}
	if d.HasChange("maxaaausers") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxaaausers has changed for aaaparameter, starting update")
		aaaparameter.Maxaaausers = d.Get("maxaaausers").(int)
		hasChange = true
	}
	if d.HasChange("maxkbquestions") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxkbquestions has changed for aaaparameter, starting update")
		aaaparameter.Maxkbquestions = d.Get("maxkbquestions").(int)
		hasChange = true
	}
	if d.HasChange("maxloginattempts") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxloginattempts has changed for aaaparameter, starting update")
		aaaparameter.Maxloginattempts = d.Get("maxloginattempts").(int)
		hasChange = true
	}
	if d.HasChange("maxsamldeflatesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxsamldeflatesize has changed for aaaparameter, starting update")
		aaaparameter.Maxsamldeflatesize = d.Get("maxsamldeflatesize").(int)
		hasChange = true
	}
	if d.HasChange("persistentloginattempts") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentloginattempts has changed for aaaparameter, starting update")
		aaaparameter.Persistentloginattempts = d.Get("persistentloginattempts").(string)
		hasChange = true
	}
	if d.HasChange("pwdexpirynotificationdays") {
		log.Printf("[DEBUG]  citrixadc-provider: Pwdexpirynotificationdays has changed for aaaparameter, starting update")
		aaaparameter.Pwdexpirynotificationdays = d.Get("pwdexpirynotificationdays").(int)
		hasChange = true
	}
	if d.HasChange("samesite") {
		log.Printf("[DEBUG]  citrixadc-provider: Samesite has changed for aaaparameter, starting update")
		aaaparameter.Samesite = d.Get("samesite").(string)
		hasChange = true
	}
	if d.HasChange("tokenintrospectioninterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Tokenintrospectioninterval has changed for aaaparameter, starting update")
		aaaparameter.Tokenintrospectioninterval = d.Get("tokenintrospectioninterval").(int)
		hasChange = true
	}
	if d.HasChange("httponlycookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookie has changed for aaaparameter, starting update")
		aaaparameter.Httponlycookie = d.Get("httponlycookie").(string)
		hasChange = true
	}
	if d.HasChange("enhancedepa") {
		log.Printf("[DEBUG]  citrixadc-provider: Enhancedepa has changed for aaaparameter, starting update")
		aaaparameter.Enhancedepa = d.Get("enhancedepa").(string)
		hasChange = true
	}
	if d.HasChange("wafprotection") {
		log.Printf("[DEBUG]  citrixadc-provider: wafprotection has changed for aaaparameter, starting update")
		aaaparameter.Wafprotection = toStringList(d.Get("wafprotection").([]interface{}))
		hasChange = true
	}
	if d.HasChange("securityinsights") {
		log.Printf("[DEBUG]  citrixadc-provider: Securityinsights has changed for aaaparameter, starting update")
		aaaparameter.Securityinsights = d.Get("securityinsights").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Aaaparameter.Type(), &aaaparameter)
		if err != nil {
			return fmt.Errorf("Error updating aaaparameter")
		}
	}
	return readAaaparameterFunc(d, meta)
}

func deleteAaaparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAaaparameterFunc")
	// aaaparameter does not support DELETE operation
	d.SetId("")

	return nil
}
