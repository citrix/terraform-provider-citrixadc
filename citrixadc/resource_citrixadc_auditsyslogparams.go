package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditsyslogparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditsyslogparamsFunc,
		Read:          readAuditsyslogparamsFunc,
		Update:        updateAuditsyslogparamsFunc,
		Delete:        deleteAuditsyslogparamsFunc,
		Schema: map[string]*schema.Schema{
			"acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowexport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contentinspectionlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dateformat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logfacility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loglevel": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lsn": &schema.Schema{
				Type:     schema.TypeString,
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
			"sslinterception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urlfiltering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userdefinedauditlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditsyslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditsyslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogparamsName := resource.PrefixedUniqueId("tf-auditsyslogparams-")
	
	auditsyslogparams := audit.Auditsyslogparams{
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Contentinspectionlog: d.Get("contentinspectionlog").(string),
		Dateformat:           d.Get("dateformat").(string),
		Dns:                  d.Get("dns").(string),
		Logfacility:          d.Get("logfacility").(string),
		Loglevel:             toStringList(d.Get("loglevel").([]interface{})),
		Lsn:                  d.Get("lsn").(string),
		Serverip:             d.Get("serverip").(string),
		Serverport:           d.Get("serverport").(int),
		Sslinterception:      d.Get("sslinterception").(string),
		Subscriberlog:        d.Get("subscriberlog").(string),
		Tcp:                  d.Get("tcp").(string),
		Timezone:             d.Get("timezone").(string),
		Urlfiltering:         d.Get("urlfiltering").(string),
		Userdefinedauditlog:  d.Get("userdefinedauditlog").(string),
	}

	err := client.UpdateUnnamedResource(service.Auditsyslogparams.Type(), &auditsyslogparams)
	if err != nil {
		return err
	}

	d.SetId(auditsyslogparamsName)

	err = readAuditsyslogparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditsyslogparams but we can't read it ??")
		return nil
	}
	return nil
}

func readAuditsyslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditsyslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading auditsyslogparams state")
	data, err := client.FindResource(service.Auditsyslogparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditsyslogparams state")
		d.SetId("")
		return nil
	}
	d.Set("acl", data["acl"])
	d.Set("alg", data["alg"])
	d.Set("appflowexport", data["appflowexport"])
	d.Set("contentinspectionlog", data["contentinspectionlog"])
	d.Set("dateformat", data["dateformat"])
	d.Set("dns", data["dns"])
	d.Set("logfacility", data["logfacility"])
	d.Set("loglevel", data["loglevel"])
	d.Set("lsn", data["lsn"])
	d.Set("serverip", data["serverip"])
	d.Set("serverport", data["serverport"])
	d.Set("sslinterception", data["sslinterception"])
	d.Set("subscriberlog", data["subscriberlog"])
	d.Set("tcp", data["tcp"])
	d.Set("timezone", data["timezone"])
	d.Set("urlfiltering", data["urlfiltering"])
	d.Set("userdefinedauditlog", data["userdefinedauditlog"])

	return nil

}

func updateAuditsyslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditsyslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	auditsyslogparams := audit.Auditsyslogparams{}
	hasChange := false
	if d.HasChange("acl") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl has changed for auditsyslogparams, starting update")
		auditsyslogparams.Acl = d.Get("acl").(string)
		hasChange = true
	}
	if d.HasChange("alg") {
		log.Printf("[DEBUG]  citrixadc-provider: Alg has changed for auditsyslogparams, starting update")
		auditsyslogparams.Alg = d.Get("alg").(string)
		hasChange = true
	}
	if d.HasChange("appflowexport") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowexport has changed for auditsyslogparams, starting update")
		auditsyslogparams.Appflowexport = d.Get("appflowexport").(string)
		hasChange = true
	}
	if d.HasChange("contentinspectionlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Contentinspectionlog has changed for auditsyslogparams, starting update")
		auditsyslogparams.Contentinspectionlog = d.Get("contentinspectionlog").(string)
		hasChange = true
	}
	if d.HasChange("dateformat") {
		log.Printf("[DEBUG]  citrixadc-provider: Dateformat has changed for auditsyslogparams, starting update")
		auditsyslogparams.Dateformat = d.Get("dateformat").(string)
		hasChange = true
	}
	if d.HasChange("dns") {
		log.Printf("[DEBUG]  citrixadc-provider: Dns has changed for auditsyslogparams, starting update")
		auditsyslogparams.Dns = d.Get("dns").(string)
		hasChange = true
	}
	if d.HasChange("logfacility") {
		log.Printf("[DEBUG]  citrixadc-provider: Logfacility has changed for auditsyslogparams, starting update")
		auditsyslogparams.Logfacility = d.Get("logfacility").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditsyslogparams, starting update")
		auditsyslogparams.Loglevel = toStringList(d.Get("loglevel").([]interface{}))
		hasChange = true
	}
	if d.HasChange("lsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Lsn has changed for auditsyslogparams, starting update")
		auditsyslogparams.Lsn = d.Get("lsn").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for auditsyslogparams, starting update")
		auditsyslogparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for auditsyslogparams, starting update")
		auditsyslogparams.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("sslinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslinterception has changed for auditsyslogparams, starting update")
		auditsyslogparams.Sslinterception = d.Get("sslinterception").(string)
		hasChange = true
	}
	if d.HasChange("subscriberlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberlog has changed for auditsyslogparams, starting update")
		auditsyslogparams.Subscriberlog = d.Get("subscriberlog").(string)
		hasChange = true
	}
	if d.HasChange("tcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcp has changed for auditsyslogparams, starting update")
		auditsyslogparams.Tcp = d.Get("tcp").(string)
		hasChange = true
	}
	if d.HasChange("timezone") {
		log.Printf("[DEBUG]  citrixadc-provider: Timezone has changed for auditsyslogparams, starting update")
		auditsyslogparams.Timezone = d.Get("timezone").(string)
		hasChange = true
	}
	if d.HasChange("urlfiltering") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlfiltering has changed for auditsyslogparams, starting update")
		auditsyslogparams.Urlfiltering = d.Get("urlfiltering").(string)
		hasChange = true
	}
	if d.HasChange("userdefinedauditlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdefinedauditlog has changed for auditsyslogparams, starting update")
		auditsyslogparams.Userdefinedauditlog = d.Get("userdefinedauditlog").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Auditsyslogparams.Type(), &auditsyslogparams)
		if err != nil {
			return fmt.Errorf("Error updating auditsyslogparams")
		}
	}
	return readAuditsyslogparamsFunc(d, meta)
}

func deleteAuditsyslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogparamsFunc")
	//auditsyslogparams does not support DELETE operation
	d.SetId("")

	return nil
}
