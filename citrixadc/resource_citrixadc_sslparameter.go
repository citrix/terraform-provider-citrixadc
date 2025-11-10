package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslparameterFunc,
		ReadContext:   readSslparameterFunc,
		UpdateContext: updateSslparameterFunc,
		DeleteContext: deleteSslparameterFunc, // Thought sslparameter resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"sigdigesttype": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"operationqueuelimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"crlmemorysizemb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cryptodevdisablelimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"denysslreneg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropreqwithnohostheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypttriggerpktcount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heterogeneoussslhw": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hybridfipsmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertcertspace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertionencoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndcppcompliancecertcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocspcachesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushenctriggertimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushflag": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quantumsize": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendclosenotify": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snihttphostmatch": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"softwarecryptothreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslierrorcache": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslimaxerrorcachemem": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssltriggertimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"strictcachecks": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefactioncontrol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefactiondata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslparameterName string

	// there is no primary key in SSLPARAMETER resource. Hence generate one for terraform state maintenance
	sslparameterName = resource.PrefixedUniqueId("tf-sslparameter-")

	sslparameter := ssl.Sslparameter{
		Defaultprofile:           d.Get("defaultprofile").(string),
		Denysslreneg:             d.Get("denysslreneg").(string),
		Dropreqwithnohostheader:  d.Get("dropreqwithnohostheader").(string),
		Heterogeneoussslhw:       d.Get("heterogeneoussslhw").(string),
		Hybridfipsmode:           d.Get("hybridfipsmode").(string),
		Insertcertspace:          d.Get("insertcertspace").(string),
		Insertionencoding:        d.Get("insertionencoding").(string),
		Ndcppcompliancecertcheck: d.Get("ndcppcompliancecertcheck").(string),
		Quantumsize:              d.Get("quantumsize").(string),
		Sendclosenotify:          d.Get("sendclosenotify").(string),
		Snihttphostmatch:         d.Get("snihttphostmatch").(string),
		Sslierrorcache:           d.Get("sslierrorcache").(string),
		Strictcachecks:           d.Get("strictcachecks").(string),
		Undefactioncontrol:       d.Get("undefactioncontrol").(string),
		Undefactiondata:          d.Get("undefactiondata").(string),
		Sigdigesttype:            toStringList(d.Get("sigdigesttype").([]interface{})),
	}
	if raw := d.GetRawConfig().GetAttr("operationqueuelimit"); !raw.IsNull() {
		sslparameter.Operationqueuelimit = intPtr(d.Get("operationqueuelimit").(int))
	}

	if raw := d.GetRawConfig().GetAttr("crlmemorysizemb"); !raw.IsNull() {
		sslparameter.Crlmemorysizemb = intPtr(d.Get("crlmemorysizemb").(int))
	}
	if raw := d.GetRawConfig().GetAttr("cryptodevdisablelimit"); !raw.IsNull() {
		sslparameter.Cryptodevdisablelimit = intPtr(d.Get("cryptodevdisablelimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("encrypttriggerpktcount"); !raw.IsNull() {
		sslparameter.Encrypttriggerpktcount = intPtr(d.Get("encrypttriggerpktcount").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ocspcachesize"); !raw.IsNull() {
		sslparameter.Ocspcachesize = intPtr(d.Get("ocspcachesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pushenctriggertimeout"); !raw.IsNull() {
		sslparameter.Pushenctriggertimeout = intPtr(d.Get("pushenctriggertimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pushflag"); !raw.IsNull() {
		sslparameter.Pushflag = intPtr(d.Get("pushflag").(int))
	}
	if raw := d.GetRawConfig().GetAttr("softwarecryptothreshold"); !raw.IsNull() {
		sslparameter.Softwarecryptothreshold = intPtr(d.Get("softwarecryptothreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("sslimaxerrorcachemem"); !raw.IsNull() {
		sslparameter.Sslimaxerrorcachemem = intPtr(d.Get("sslimaxerrorcachemem").(int))
	}
	if raw := d.GetRawConfig().GetAttr("ssltriggertimeout"); !raw.IsNull() {
		sslparameter.Ssltriggertimeout = intPtr(d.Get("ssltriggertimeout").(int))
	}

	err := client.UpdateUnnamedResource(service.Sslparameter.Type(), &sslparameter)
	if err != nil {
		return diag.Errorf("Error updating sslparameter")
	}

	d.SetId(sslparameterName)

	return readSslparameterFunc(ctx, d, meta)
}

func readSslparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading sslparameter state")
	data, err := client.FindResource(service.Sslparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslparameter state")
		d.SetId("")
		return nil
	}
	setToInt("crlmemorysizemb", d, data["crlmemorysizemb"])
	setToInt("cryptodevdisablelimit", d, data["cryptodevdisablelimit"])
	d.Set("defaultprofile", data["defaultprofile"])
	d.Set("sigdigesttype", data["sigdigesttype"])
	setToInt("operationqueuelimit", d, data["operationqueuelimit"])
	d.Set("denysslreneg", data["denysslreneg"])
	d.Set("dropreqwithnohostheader", data["dropreqwithnohostheader"])
	setToInt("encrypttriggerpktcount", d, data["encrypttriggerpktcount"])
	d.Set("heterogeneoussslhw", data["heterogeneoussslhw"])
	d.Set("hybridfipsmode", data["hybridfipsmode"])
	d.Set("insertcertspace", data["insertcertspace"])
	d.Set("insertionencoding", data["insertionencoding"])
	d.Set("ndcppcompliancecertcheck", data["ndcppcompliancecertcheck"])
	setToInt("ocspcachesize", d, data["ocspcachesize"])
	setToInt("pushenctriggertimeout", d, data["pushenctriggertimeout"])
	setToInt("pushflag", d, data["pushflag"])
	d.Set("quantumsize", data["quantumsize"])
	d.Set("sendclosenotify", data["sendclosenotify"])
	d.Set("snihttphostmatch", data["snihttphostmatch"])
	setToInt("softwarecryptothreshold", d, data["softwarecryptothreshold"])
	d.Set("sslierrorcache", data["sslierrorcache"])
	setToInt("sslimaxerrorcachemem", d, data["sslimaxerrorcachemem"])
	setToInt("ssltriggertimeout", d, data["ssltriggertimeout"])
	d.Set("strictcachecks", data["strictcachecks"])
	d.Set("undefactioncontrol", data["undefactioncontrol"])
	d.Set("undefactiondata", data["undefactiondata"])

	return nil
}

func updateSslparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	sslparameter := ssl.Sslparameter{}

	hasChange := false
	if d.HasChange("sigdigesttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sigdigesttype has changed for sslparameter, starting update")
		sslparameter.Sigdigesttype = toStringList(d.Get("sigdigesttype").([]interface{}))
		hasChange = true
	}
	if d.HasChange("operationqueuelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Operationqueuelimit has changed for sslparameter, starting update")
		sslparameter.Operationqueuelimit = intPtr(d.Get("operationqueuelimit").(int))
		hasChange = true
	}
	if d.HasChange("crlmemorysizemb") {
		log.Printf("[DEBUG]  citrixadc-provider: Crlmemorysizemb has changed for sslparameter, starting update")
		sslparameter.Crlmemorysizemb = intPtr(d.Get("crlmemorysizemb").(int))
		hasChange = true
	}
	if d.HasChange("cryptodevdisablelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Cryptodevdisablelimit has changed for sslparameter, starting update")
		sslparameter.Cryptodevdisablelimit = intPtr(d.Get("cryptodevdisablelimit").(int))
		hasChange = true
	}
	if d.HasChange("defaultprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultprofile has changed for sslparameter, starting update")
		sslparameter.Defaultprofile = d.Get("defaultprofile").(string)
		hasChange = true
	}
	if d.HasChange("denysslreneg") {
		log.Printf("[DEBUG]  citrixadc-provider: Denysslreneg has changed for sslparameter, starting update")
		sslparameter.Denysslreneg = d.Get("denysslreneg").(string)
		hasChange = true
	}
	if d.HasChange("dropreqwithnohostheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropreqwithnohostheader has changed for sslparameter, starting update")
		sslparameter.Dropreqwithnohostheader = d.Get("dropreqwithnohostheader").(string)
		hasChange = true
	}
	if d.HasChange("encrypttriggerpktcount") {
		log.Printf("[DEBUG]  citrixadc-provider: Encrypttriggerpktcount has changed for sslparameter, starting update")
		sslparameter.Encrypttriggerpktcount = intPtr(d.Get("encrypttriggerpktcount").(int))
		hasChange = true
	}
	if d.HasChange("heterogeneoussslhw") {
		log.Printf("[DEBUG]  citrixadc-provider: Heterogeneoussslhw has changed for sslparameter, starting update")
		sslparameter.Heterogeneoussslhw = d.Get("heterogeneoussslhw").(string)
		hasChange = true
	}
	if d.HasChange("hybridfipsmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Hybridfipsmode has changed for sslparameter, starting update")
		sslparameter.Hybridfipsmode = d.Get("hybridfipsmode").(string)
		hasChange = true
	}
	if d.HasChange("insertcertspace") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertcertspace has changed for sslparameter, starting update")
		sslparameter.Insertcertspace = d.Get("insertcertspace").(string)
		hasChange = true
	}
	if d.HasChange("insertionencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertionencoding has changed for sslparameter, starting update")
		sslparameter.Insertionencoding = d.Get("insertionencoding").(string)
		hasChange = true
	}
	if d.HasChange("ndcppcompliancecertcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Ndcppcompliancecertcheck has changed for sslparameter, starting update")
		sslparameter.Ndcppcompliancecertcheck = d.Get("ndcppcompliancecertcheck").(string)
		hasChange = true
	}
	if d.HasChange("ocspcachesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspcachesize has changed for sslparameter, starting update")
		sslparameter.Ocspcachesize = intPtr(d.Get("ocspcachesize").(int))
		hasChange = true
	}
	if d.HasChange("pushenctriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctriggertimeout has changed for sslparameter, starting update")
		sslparameter.Pushenctriggertimeout = intPtr(d.Get("pushenctriggertimeout").(int))
		hasChange = true
	}
	if d.HasChange("pushflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushflag has changed for sslparameter, starting update")
		sslparameter.Pushflag = intPtr(d.Get("pushflag").(int))
		hasChange = true
	}
	if d.HasChange("quantumsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Quantumsize has changed for sslparameter, starting update")
		sslparameter.Quantumsize = d.Get("quantumsize").(string)
		hasChange = true
	}
	if d.HasChange("sendclosenotify") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendclosenotify has changed for sslparameter, starting update")
		sslparameter.Sendclosenotify = d.Get("sendclosenotify").(string)
		hasChange = true
	}
	if d.HasChange("snihttphostmatch") {
		log.Printf("[DEBUG]  citrixadc-provider: Snihttphostmatch has changed for sslparameter, starting update")
		sslparameter.Snihttphostmatch = d.Get("snihttphostmatch").(string)
		hasChange = true
	}
	if d.HasChange("softwarecryptothreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Softwarecryptothreshold has changed for sslparameter, starting update")
		sslparameter.Softwarecryptothreshold = intPtr(d.Get("softwarecryptothreshold").(int))
		hasChange = true
	}
	if d.HasChange("sslierrorcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslierrorcache has changed for sslparameter, starting update")
		sslparameter.Sslierrorcache = d.Get("sslierrorcache").(string)
		hasChange = true
	}
	if d.HasChange("sslimaxerrorcachemem") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslimaxerrorcachemem has changed for sslparameter, starting update")
		sslparameter.Sslimaxerrorcachemem = intPtr(d.Get("sslimaxerrorcachemem").(int))
		hasChange = true
	}
	if d.HasChange("ssltriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssltriggertimeout has changed for sslparameter, starting update")
		sslparameter.Ssltriggertimeout = intPtr(d.Get("ssltriggertimeout").(int))
		hasChange = true
	}
	if d.HasChange("strictcachecks") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictcachecks has changed for sslparameter, starting update")
		sslparameter.Strictcachecks = d.Get("strictcachecks").(string)
		hasChange = true
	}
	if d.HasChange("undefactioncontrol") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefactioncontrol has changed for sslparameter, starting update")
		sslparameter.Undefactioncontrol = d.Get("undefactioncontrol").(string)
		hasChange = true
	}
	if d.HasChange("undefactiondata") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefactiondata has changed for sslparameter, starting update")
		sslparameter.Undefactiondata = d.Get("undefactiondata").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Sslparameter.Type(), &sslparameter)
		if err != nil {
			return diag.Errorf("Error updating sslparameter: %s", err.Error())
		}
	}
	return readSslparameterFunc(ctx, d, meta)
}

func deleteSslparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslparameterFunc")
	// sslparameter do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
