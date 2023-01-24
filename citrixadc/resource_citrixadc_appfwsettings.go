package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppfwsettings() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppfwsettingsFunc,
		Read:          readAppfwsettingsFunc,
		Update:        updateAppfwsettingsFunc,
		Delete:        deleteAppfwsettingsFunc,
		Schema: map[string]*schema.Schema{
			"ceflogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"centralizedlearning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientiploggingheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiepostencryptprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entitydecoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geolocationlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"importsizelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"learnratelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"logmalformedreq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformedreqaction": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"proxyport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxyserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessioncookiename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"signatureautoupdate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatureurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useconfigurablesecretkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppfwsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppfwsettingsFunc")
	client := meta.(*NetScalerNitroClient).client

	appfwsettingsName := resource.PrefixedUniqueId("tf-appfwsettings-")

	appfwsettings := appfw.Appfwsettings{
		Ceflogging:               d.Get("ceflogging").(string),
		Centralizedlearning:      d.Get("centralizedlearning").(string),
		Clientiploggingheader:    d.Get("clientiploggingheader").(string),
		Cookiepostencryptprefix:  d.Get("cookiepostencryptprefix").(string),
		Defaultprofile:           d.Get("defaultprofile").(string),
		Entitydecoding:           d.Get("entitydecoding").(string),
		Geolocationlogging:       d.Get("geolocationlogging").(string),
		Importsizelimit:          d.Get("importsizelimit").(int),
		Learnratelimit:           d.Get("learnratelimit").(int),
		Logmalformedreq:          d.Get("logmalformedreq").(string),
		Malformedreqaction:       toStringList(d.Get("malformedreqaction").([]interface{})),
		Proxyport:                d.Get("proxyport").(int),
		Proxyserver:              d.Get("proxyserver").(string),
		Sessioncookiename:        d.Get("sessioncookiename").(string),
		Sessionlifetime:          d.Get("sessionlifetime").(int),
		Sessionlimit:             d.Get("sessionlimit").(int),
		Sessiontimeout:           d.Get("sessiontimeout").(int),
		Signatureautoupdate:      d.Get("signatureautoupdate").(string),
		Signatureurl:             d.Get("signatureurl").(string),
		Undefaction:              d.Get("undefaction").(string),
		Useconfigurablesecretkey: d.Get("useconfigurablesecretkey").(string),
	}

	err := client.UpdateUnnamedResource(service.Appfwsettings.Type(), &appfwsettings)
	if err != nil {
		return err
	}

	d.SetId(appfwsettingsName)

	err = readAppfwsettingsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appfwsettings but we can't read it ?? %s", appfwsettingsName)
		return nil
	}
	return nil
}

func readAppfwsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppfwsettingsFunc")
	client := meta.(*NetScalerNitroClient).client
	appfwsettingsName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appfwsettings state %s", appfwsettingsName)
	data, err := client.FindResource(service.Appfwsettings.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appfwsettings state %s", appfwsettingsName)
		d.SetId("")
		return nil
	}
	d.Set("ceflogging", data["ceflogging"])
	d.Set("centralizedlearning", data["centralizedlearning"])
	d.Set("clientiploggingheader", data["clientiploggingheader"])
	d.Set("cookiepostencryptprefix", data["cookiepostencryptprefix"])
	d.Set("defaultprofile", data["defaultprofile"])
	d.Set("entitydecoding", data["entitydecoding"])
	d.Set("geolocationlogging", data["geolocationlogging"])
	d.Set("importsizelimit", data["importsizelimit"])
	d.Set("learnratelimit", data["learnratelimit"])
	d.Set("logmalformedreq", data["logmalformedreq"])
	d.Set("malformedreqaction", data["malformedreqaction"])
	d.Set("proxyport", data["proxyport"])
	d.Set("proxyserver", data["proxyserver"])
	d.Set("sessioncookiename", data["sessioncookiename"])
	d.Set("sessionlifetime", data["sessionlifetime"])
	d.Set("sessionlimit", data["sessionlimit"])
	d.Set("sessiontimeout", data["sessiontimeout"])
	d.Set("signatureautoupdate", data["signatureautoupdate"])
	d.Set("signatureurl", data["signatureurl"])
	d.Set("undefaction", data["undefaction"])
	d.Set("useconfigurablesecretkey", data["useconfigurablesecretkey"])

	return nil

}

func updateAppfwsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppfwsettingsFunc")
	client := meta.(*NetScalerNitroClient).client

	appfwsettings := appfw.Appfwsettings{}
	hasChange := false

	if d.HasChange("ceflogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Ceflogging has changed for appfwsettings, starting update")
		appfwsettings.Ceflogging = d.Get("ceflogging").(string)
		hasChange = true
	}
	if d.HasChange("centralizedlearning") {
		log.Printf("[DEBUG]  citrixadc-provider: Centralizedlearning has changed for appfwsettings, starting update")
		appfwsettings.Centralizedlearning = d.Get("centralizedlearning").(string)
		hasChange = true
	}
	if d.HasChange("clientiploggingheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientiploggingheader has changed for appfwsettings, starting update")
		appfwsettings.Clientiploggingheader = d.Get("clientiploggingheader").(string)
		hasChange = true
	}
	if d.HasChange("cookiepostencryptprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiepostencryptprefix has changed for appfwsettings, starting update")
		appfwsettings.Cookiepostencryptprefix = d.Get("cookiepostencryptprefix").(string)
		hasChange = true
	}
	if d.HasChange("defaultprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultprofile has changed for appfwsettings, starting update")
		appfwsettings.Defaultprofile = d.Get("defaultprofile").(string)
		hasChange = true
	}
	if d.HasChange("entitydecoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Entitydecoding has changed for appfwsettings, starting update")
		appfwsettings.Entitydecoding = d.Get("entitydecoding").(string)
		hasChange = true
	}
	if d.HasChange("geolocationlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Geolocationlogging has changed for appfwsettings, starting update")
		appfwsettings.Geolocationlogging = d.Get("geolocationlogging").(string)
		hasChange = true
	}
	if d.HasChange("importsizelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Importsizelimit has changed for appfwsettings, starting update")
		appfwsettings.Importsizelimit = d.Get("importsizelimit").(int)
		hasChange = true
	}
	if d.HasChange("learnratelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Learnratelimit has changed for appfwsettings, starting update")
		appfwsettings.Learnratelimit = d.Get("learnratelimit").(int)
		hasChange = true
	}
	if d.HasChange("logmalformedreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Logmalformedreq has changed for appfwsettings, starting update")
		appfwsettings.Logmalformedreq = d.Get("logmalformedreq").(string)
		hasChange = true
	}
	if d.HasChange("malformedreqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Malformedreqaction has changed for appfwsettings, starting update")
		appfwsettings.Malformedreqaction = toStringList(d.Get("malformedreqaction").([]interface{}))
		hasChange = true
	}
	if d.HasChange("proxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyport has changed for appfwsettings, starting update")
		appfwsettings.Proxyport = d.Get("proxyport").(int)
		hasChange = true
	}
	if d.HasChange("proxyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyserver has changed for appfwsettings, starting update")
		appfwsettings.Proxyserver = d.Get("proxyserver").(string)
		hasChange = true
	}
	if d.HasChange("sessioncookiename") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessioncookiename has changed for appfwsettings, starting update")
		appfwsettings.Sessioncookiename = d.Get("sessioncookiename").(string)
		hasChange = true
	}
	if d.HasChange("sessionlifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlifetime has changed for appfwsettings, starting update")
		appfwsettings.Sessionlifetime = d.Get("sessionlifetime").(int)
		hasChange = true
	}
	if d.HasChange("sessionlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlimit has changed for appfwsettings, starting update")
		appfwsettings.Sessionlimit = d.Get("sessionlimit").(int)
		hasChange = true
	}
	if d.HasChange("sessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessiontimeout has changed for appfwsettings, starting update")
		appfwsettings.Sessiontimeout = d.Get("sessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("signatureautoupdate") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureautoupdate has changed for appfwsettings, starting update")
		appfwsettings.Signatureautoupdate = d.Get("signatureautoupdate").(string)
		hasChange = true
	}
	if d.HasChange("signatureurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureurl has changed for appfwsettings, starting update")
		appfwsettings.Signatureurl = d.Get("signatureurl").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for appfwsettings, starting update")
		appfwsettings.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}
	if d.HasChange("useconfigurablesecretkey") {
		log.Printf("[DEBUG]  citrixadc-provider: Useconfigurablesecretkey has changed for appfwsettings, starting update")
		appfwsettings.Useconfigurablesecretkey = d.Get("useconfigurablesecretkey").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appfwsettings.Type(), &appfwsettings)
		if err != nil {
			return fmt.Errorf("Error updating appfwsettings %s", err.Error())
		}
	}
	return readAppfwsettingsFunc(d, meta)
}

func deleteAppfwsettingsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppfwsettingsFunc")

	d.SetId("")

	return nil
}
