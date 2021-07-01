package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSslparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslparameterFunc,
		Read:          readSslparameterFunc,
		Update:        updateSslparameterFunc,
		Delete:        deleteSslparameterFunc, // Thought sslparameter resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"crlmemorysizemb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cryptodevdisablelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defaultprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"denysslreneg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropreqwithnohostheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypttriggerpktcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"heterogeneoussslhw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hybridfipsmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertcertspace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertionencoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndcppcompliancecertcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocspcachesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushenctriggertimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushflag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quantumsize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendclosenotify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snihttphostmatch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"softwarecryptothreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslierrorcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslimaxerrorcachemem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssltriggertimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"strictcachecks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefactioncontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefactiondata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslparameterName string

	// there is no primary key in SSLPARAMETER resource. Hence generate one for terraform state maintenance
	sslparameterName = resource.PrefixedUniqueId("tf-sslparameter-")

	sslparameter := ssl.Sslparameter{
		Crlmemorysizemb:          uint32(d.Get("crlmemorysizemb").(int)),
		Cryptodevdisablelimit:    uint32(d.Get("cryptodevdisablelimit").(int)),
		Defaultprofile:           d.Get("defaultprofile").(string),
		Denysslreneg:             d.Get("denysslreneg").(string),
		Dropreqwithnohostheader:  d.Get("dropreqwithnohostheader").(string),
		Encrypttriggerpktcount:   uint32(d.Get("encrypttriggerpktcount").(int)),
		Heterogeneoussslhw:       d.Get("heterogeneoussslhw").(string),
		Hybridfipsmode:           d.Get("hybridfipsmode").(string),
		Insertcertspace:          d.Get("insertcertspace").(string),
		Insertionencoding:        d.Get("insertionencoding").(string),
		Ndcppcompliancecertcheck: d.Get("ndcppcompliancecertcheck").(string),
		Ocspcachesize:            uint32(d.Get("ocspcachesize").(int)),
		Pushenctriggertimeout:    uint32(d.Get("pushenctriggertimeout").(int)),
		Pushflag:                 uint32(d.Get("pushflag").(int)),
		Quantumsize:              d.Get("quantumsize").(string),
		Sendclosenotify:          d.Get("sendclosenotify").(string),
		Snihttphostmatch:         d.Get("snihttphostmatch").(string),
		Softwarecryptothreshold:  uint32(d.Get("softwarecryptothreshold").(int)),
		Sslierrorcache:           d.Get("sslierrorcache").(string),
		Sslimaxerrorcachemem:     uint32(d.Get("sslimaxerrorcachemem").(int)),
		Ssltriggertimeout:        uint32(d.Get("ssltriggertimeout").(int)),
		Strictcachecks:           d.Get("strictcachecks").(string),
		Undefactioncontrol:       d.Get("undefactioncontrol").(string),
		Undefactiondata:          d.Get("undefactiondata").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslparameter.Type(), &sslparameter)
	if err != nil {
		return fmt.Errorf("Error updating sslparameter")
	}

	d.SetId(sslparameterName)

	err = readSslparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just updated the sslparameter but we can't read it ??")
		return nil
	}
	return nil
}

func readSslparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading sslparameter state")
	data, err := client.FindResource(service.Sslparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslparameter state")
		d.SetId("")
		return nil
	}
	d.Set("crlmemorysizemb", data["crlmemorysizemb"])
	d.Set("cryptodevdisablelimit", data["cryptodevdisablelimit"])
	d.Set("defaultprofile", data["defaultprofile"])
	d.Set("denysslreneg", data["denysslreneg"])
	d.Set("dropreqwithnohostheader", data["dropreqwithnohostheader"])
	d.Set("encrypttriggerpktcount", data["encrypttriggerpktcount"])
	d.Set("heterogeneoussslhw", data["heterogeneoussslhw"])
	d.Set("hybridfipsmode", data["hybridfipsmode"])
	d.Set("insertcertspace", data["insertcertspace"])
	d.Set("insertionencoding", data["insertionencoding"])
	d.Set("ndcppcompliancecertcheck", data["ndcppcompliancecertcheck"])
	d.Set("ocspcachesize", data["ocspcachesize"])
	d.Set("pushenctriggertimeout", data["pushenctriggertimeout"])
	d.Set("pushflag", data["pushflag"])
	d.Set("quantumsize", data["quantumsize"])
	d.Set("sendclosenotify", data["sendclosenotify"])
	d.Set("snihttphostmatch", data["snihttphostmatch"])
	d.Set("softwarecryptothreshold", data["softwarecryptothreshold"])
	d.Set("sslierrorcache", data["sslierrorcache"])
	d.Set("sslimaxerrorcachemem", data["sslimaxerrorcachemem"])
	d.Set("ssltriggertimeout", data["ssltriggertimeout"])
	d.Set("strictcachecks", data["strictcachecks"])
	d.Set("undefactioncontrol", data["undefactioncontrol"])
	d.Set("undefactiondata", data["undefactiondata"])

	return nil
}

func updateSslparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslparameterFunc")
	client := meta.(*NetScalerNitroClient).client

	sslparameter := ssl.Sslparameter{}

	hasChange := false
	if d.HasChange("crlmemorysizemb") {
		log.Printf("[DEBUG]  citrixadc-provider: Crlmemorysizemb has changed for sslparameter, starting update")
		sslparameter.Crlmemorysizemb = uint32(d.Get("crlmemorysizemb").(int))
		hasChange = true
	}
	if d.HasChange("cryptodevdisablelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Cryptodevdisablelimit has changed for sslparameter, starting update")
		sslparameter.Cryptodevdisablelimit = uint32(d.Get("cryptodevdisablelimit").(int))
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
		sslparameter.Encrypttriggerpktcount = uint32(d.Get("encrypttriggerpktcount").(int))
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
		sslparameter.Ocspcachesize = uint32(d.Get("ocspcachesize").(int))
		hasChange = true
	}
	if d.HasChange("pushenctriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctriggertimeout has changed for sslparameter, starting update")
		sslparameter.Pushenctriggertimeout = uint32(d.Get("pushenctriggertimeout").(int))
		hasChange = true
	}
	if d.HasChange("pushflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushflag has changed for sslparameter, starting update")
		sslparameter.Pushflag = uint32(d.Get("pushflag").(int))
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
		sslparameter.Softwarecryptothreshold = uint32(d.Get("softwarecryptothreshold").(int))
		hasChange = true
	}
	if d.HasChange("sslierrorcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslierrorcache has changed for sslparameter, starting update")
		sslparameter.Sslierrorcache = d.Get("sslierrorcache").(string)
		hasChange = true
	}
	if d.HasChange("sslimaxerrorcachemem") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslimaxerrorcachemem has changed for sslparameter, starting update")
		sslparameter.Sslimaxerrorcachemem = uint32(d.Get("sslimaxerrorcachemem").(int))
		hasChange = true
	}
	if d.HasChange("ssltriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssltriggertimeout has changed for sslparameter, starting update")
		sslparameter.Ssltriggertimeout = uint32(d.Get("ssltriggertimeout").(int))
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
			return fmt.Errorf("Error updating sslparameter: %s", err.Error())
		}
	}
	return readSslparameterFunc(d, meta)
}

func deleteSslparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslparameterFunc")
	// sslparameter do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")
	return nil
}
