package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditnslogaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditnslogactionFunc,
		Read:          readAuditnslogactionFunc,
		Update:        updateAuditnslogactionFunc,
		Delete:        deleteAuditnslogactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"loglevel": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"acl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowexport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contentinspectionlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dateformat": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainresolvenow": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"domainresolveretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"logfacility": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lsn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverdomainname": {
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
			"sslinterception": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urlfiltering": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userdefinedauditlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditnslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditnslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogactionName := d.Get("name").(string)
	auditnslogaction := audit.Auditnslogaction{
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Contentinspectionlog: d.Get("contentinspectionlog").(string),
		Dateformat:           d.Get("dateformat").(string),
		Domainresolvenow:     d.Get("domainresolvenow").(bool),
		Domainresolveretry:   d.Get("domainresolveretry").(int),
		Logfacility:          d.Get("logfacility").(string),
		Loglevel:             toStringList(d.Get("loglevel").([]interface{})),
		Lsn:                  d.Get("lsn").(string),
		Name:                 d.Get("name").(string),
		Serverdomainname:     d.Get("serverdomainname").(string),
		Serverip:             d.Get("serverip").(string),
		Serverport:           d.Get("serverport").(int),
		Sslinterception:      d.Get("sslinterception").(string),
		Subscriberlog:        d.Get("subscriberlog").(string),
		Tcp:                  d.Get("tcp").(string),
		Timezone:             d.Get("timezone").(string),
		Urlfiltering:         d.Get("urlfiltering").(string),
		Userdefinedauditlog:  d.Get("userdefinedauditlog").(string),
	}

	_, err := client.AddResource(service.Auditnslogaction.Type(), auditnslogactionName, &auditnslogaction)
	if err != nil {
		return err
	}

	d.SetId(auditnslogactionName)

	err = readAuditnslogactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditnslogaction but we can't read it ?? %s", auditnslogactionName)
		return nil
	}
	return nil
}

func readAuditnslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditnslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditnslogaction state %s", auditnslogactionName)
	data, err := client.FindResource(service.Auditnslogaction.Type(), auditnslogactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditnslogaction state %s", auditnslogactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("acl", data["acl"])
	d.Set("alg", data["alg"])
	d.Set("appflowexport", data["appflowexport"])
	d.Set("contentinspectionlog", data["contentinspectionlog"])
	d.Set("dateformat", data["dateformat"])
	d.Set("domainresolvenow", data["domainresolvenow"])
	d.Set("domainresolveretry", data["domainresolveretry"])
	d.Set("logfacility", data["logfacility"])
	d.Set("loglevel", data["loglevel"])
	d.Set("lsn", data["lsn"])
	d.Set("serverdomainname", data["serverdomainname"])
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

func updateAuditnslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditnslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogactionName := d.Get("name").(string)

	auditnslogaction := audit.Auditnslogaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acl") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Acl = d.Get("acl").(string)
		hasChange = true
	}
	if d.HasChange("alg") {
		log.Printf("[DEBUG]  citrixadc-provider: Alg has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Alg = d.Get("alg").(string)
		hasChange = true
	}
	if d.HasChange("appflowexport") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowexport has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Appflowexport = d.Get("appflowexport").(string)
		hasChange = true
	}
	if d.HasChange("contentinspectionlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Contentinspectionlog has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Contentinspectionlog = d.Get("contentinspectionlog").(string)
		hasChange = true
	}
	if d.HasChange("dateformat") {
		log.Printf("[DEBUG]  citrixadc-provider: Dateformat has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Dateformat = d.Get("dateformat").(string)
		hasChange = true
	}
	if d.HasChange("domainresolvenow") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolvenow has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Domainresolvenow = d.Get("domainresolvenow").(bool)
		hasChange = true
	}
	if d.HasChange("domainresolveretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolveretry has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Domainresolveretry = d.Get("domainresolveretry").(int)
		hasChange = true
	}
	if d.HasChange("logfacility") {
		log.Printf("[DEBUG]  citrixadc-provider: Logfacility has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Logfacility = d.Get("logfacility").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Loglevel = toStringList(d.Get("loglevel").([]interface{}))
		hasChange = true
	}
	if d.HasChange("lsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Lsn has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Lsn = d.Get("lsn").(string)
		hasChange = true
	}
	if d.HasChange("serverdomainname") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverdomainname has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Serverdomainname = d.Get("serverdomainname").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Serverport = d.Get("serverport").(int)
		hasChange = true
	}
	if d.HasChange("sslinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslinterception has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Sslinterception = d.Get("sslinterception").(string)
		hasChange = true
	}
	if d.HasChange("subscriberlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberlog has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Subscriberlog = d.Get("subscriberlog").(string)
		hasChange = true
	}
	if d.HasChange("tcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcp has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Tcp = d.Get("tcp").(string)
		hasChange = true
	}
	if d.HasChange("timezone") {
		log.Printf("[DEBUG]  citrixadc-provider: Timezone has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Timezone = d.Get("timezone").(string)
		hasChange = true
	}
	if d.HasChange("urlfiltering") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlfiltering has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Urlfiltering = d.Get("urlfiltering").(string)
		hasChange = true
	}
	if d.HasChange("userdefinedauditlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdefinedauditlog has changed for auditnslogaction %s, starting update", auditnslogactionName)
		auditnslogaction.Userdefinedauditlog = d.Get("userdefinedauditlog").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Auditnslogaction.Type(), &auditnslogaction)
		if err != nil {
			return fmt.Errorf("Error updating auditnslogaction %s", auditnslogactionName)
		}
	}
	return readAuditnslogactionFunc(d, meta)
}

func deleteAuditnslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditnslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogactionName := d.Id()
	err := client.DeleteResource(service.Auditnslogaction.Type(), auditnslogactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
