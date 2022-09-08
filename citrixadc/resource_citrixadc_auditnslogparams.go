package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditnslogparams() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditnslogparamsFunc,
		Read:          readAuditnslogparamsFunc,
		Update:        updateAuditnslogparamsFunc,
		Delete:        deleteAuditnslogparamsFunc,
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

func createAuditnslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditnslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogparamsName := resource.PrefixedUniqueId("tf-auditnslogparams-")

	auditnslogparams := audit.Auditnslogparams{
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Contentinspectionlog: d.Get("contentinspectionlog").(string),
		Dateformat:           d.Get("dateformat").(string),
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

	err := client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
	if err != nil {
		return err
	}

	d.SetId(auditnslogparamsName)

	err = readAuditnslogparamsFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditnslogparams but we can't read it ??")
		return nil
	}
	return nil
}

func readAuditnslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditnslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading auditnslogparams state")
	data, err := client.FindResource(service.Auditnslogparams.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditnslogparams state")
		d.SetId("")
		return nil
	}
	d.Set("acl", data["acl"])
	d.Set("alg", data["alg"])
	d.Set("appflowexport", data["appflowexport"])
	d.Set("contentinspectionlog", data["contentinspectionlog"])
	d.Set("dateformat", data["dateformat"])
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

func updateAuditnslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditnslogparamsFunc")
	client := meta.(*NetScalerNitroClient).client

	auditnslogparams := audit.Auditnslogparams{}
	hasChange := false
	if d.HasChange("acl") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl has changed for auditnslogparams, starting update")
		auditnslogparams.Acl = d.Get("acl").(string)
		hasChange = true
	}
	if d.HasChange("alg") {
		log.Printf("[DEBUG]  citrixadc-provider: Alg has changed for auditnslogparams, starting update")
		auditnslogparams.Alg = d.Get("alg").(string)
		hasChange = true
	}
	if d.HasChange("appflowexport") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowexport has changed for auditnslogparams, starting update")
		auditnslogparams.Appflowexport = d.Get("appflowexport").(string)
		hasChange = true
	}
	if d.HasChange("contentinspectionlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Contentinspectionlog has changed for auditnslogparams, starting update")
		auditnslogparams.Contentinspectionlog = d.Get("contentinspectionlog").(string)
		hasChange = true
	}
	if d.HasChange("dateformat") {
		log.Printf("[DEBUG]  citrixadc-provider: Dateformat has changed for auditnslogparams, starting update")
		auditnslogparams.Dateformat = d.Get("dateformat").(string)
		hasChange = true
	}
	if d.HasChange("logfacility") {
		log.Printf("[DEBUG]  citrixadc-provider: Logfacility has changed for auditnslogparams, starting update")
		auditnslogparams.Logfacility = d.Get("logfacility").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditnslogparams, starting update")
		auditnslogparams.Loglevel = toStringList(d.Get("loglevel").([]interface{}))
		hasChange = true
	}
	if d.HasChange("lsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Lsn has changed for auditnslogparams, starting update")
		auditnslogparams.Lsn = d.Get("lsn").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for auditnslogparams, starting update")
		auditnslogparams.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for auditnslogparams, starting update")
		auditnslogparams.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("sslinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslinterception has changed for auditnslogparams, starting update")
		auditnslogparams.Sslinterception = d.Get("sslinterception").(string)
		hasChange = true
	}
	if d.HasChange("subscriberlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberlog has changed for auditnslogparams, starting update")
		auditnslogparams.Subscriberlog = d.Get("subscriberlog").(string)
		hasChange = true
	}
	if d.HasChange("tcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcp has changed for auditnslogparams, starting update")
		auditnslogparams.Tcp = d.Get("tcp").(string)
		hasChange = true
	}
	if d.HasChange("timezone") {
		log.Printf("[DEBUG]  citrixadc-provider: Timezone has changed for auditnslogparams, starting update")
		auditnslogparams.Timezone = d.Get("timezone").(string)
		hasChange = true
	}
	if d.HasChange("urlfiltering") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlfiltering has changed for auditnslogparams, starting update")
		auditnslogparams.Urlfiltering = d.Get("urlfiltering").(string)
		hasChange = true
	}
	if d.HasChange("userdefinedauditlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdefinedauditlog has changed for auditnslogparams, starting update")
		auditnslogparams.Userdefinedauditlog = d.Get("userdefinedauditlog").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
		if err != nil {
			return fmt.Errorf("Error updating auditnslogparams")
		}
	}
	return readAuditnslogparamsFunc(d, meta)
}

func deleteAuditnslogparamsFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditnslogparamsFunc")
	// auditnslogparams does not support DELETE operation
	d.SetId("")

	return nil
}
