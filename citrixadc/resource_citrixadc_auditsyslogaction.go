package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuditsyslogaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuditsyslogactionFunc,
		ReadContext:   readAuditsyslogactionFunc,
		UpdateContext: updateAuditsyslogactionFunc,
		DeleteContext: deleteAuditsyslogactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"protocolviolations": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dns": {
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
			"lbvservername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logfacility": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loglevel": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lsn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxlogdatasizetohold": {

				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"netprofile": {
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
			"tcpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"managementlog": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mgmtloglevel": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"syslogcompliance": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpauthtoken": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpendpointurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"streamanalytics": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditsyslogactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client

	auditsyslogactionName := d.Get("name").(string)

	auditsyslogaction := audit.Auditsyslogaction{
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Contentinspectionlog: d.Get("contentinspectionlog").(string),
		Dateformat:           d.Get("dateformat").(string),
		Dns:                  d.Get("dns").(string),
		Domainresolvenow:     d.Get("domainresolvenow").(bool),
		Lbvservername:        d.Get("lbvservername").(string),
		Logfacility:          d.Get("logfacility").(string),
		Loglevel:             toStringList(loglevelValue(d)),
		Lsn:                  d.Get("lsn").(string),
		Name:                 d.Get("name").(string),
		Netprofile:           d.Get("netprofile").(string),
		Serverdomainname:     d.Get("serverdomainname").(string),
		Serverip:             d.Get("serverip").(string),
		Sslinterception:      d.Get("sslinterception").(string),
		Subscriberlog:        d.Get("subscriberlog").(string),
		Tcp:                  d.Get("tcp").(string),
		Tcpprofilename:       d.Get("tcpprofilename").(string),
		Timezone:             d.Get("timezone").(string),
		Transport:            d.Get("transport").(string),
		Urlfiltering:         d.Get("urlfiltering").(string),
		Userdefinedauditlog:  d.Get("userdefinedauditlog").(string),
		Syslogcompliance:     d.Get("syslogcompliance").(string),
		Httpauthtoken:        d.Get("httpauthtoken").(string),
		Httpendpointurl:      d.Get("httpendpointurl").(string),
		Streamanalytics:      d.Get("streamanalytics").(string),
		Protocolviolations:   d.Get("protocolviolations").(string),
	}

	if raw := d.GetRawConfig().GetAttr("domainresolveretry"); !raw.IsNull() {
		auditsyslogaction.Domainresolveretry = intPtr(d.Get("domainresolveretry").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxlogdatasizetohold"); !raw.IsNull() {
		auditsyslogaction.Maxlogdatasizetohold = intPtr(d.Get("maxlogdatasizetohold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		auditsyslogaction.Serverport = intPtr(d.Get("serverport").(int))
	}
	if listVal, ok := d.Get("managementlog").([]interface{}); ok {
		auditsyslogaction.Managementlog = toStringList(listVal)
	}
	if listVal, ok := d.Get("mgmtloglevel").([]interface{}); ok {
		auditsyslogaction.Mgmtloglevel = toStringList(listVal)
	}

	_, err := client.AddResource(service.Auditsyslogaction.Type(), auditsyslogactionName, &auditsyslogaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(auditsyslogactionName)

	return readAuditsyslogactionFunc(ctx, d, meta)
}

func readAuditsyslogactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditsyslogaction state %s", auditsyslogactionName)
	data, err := client.FindResource(service.Auditsyslogaction.Type(), auditsyslogactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditsyslogaction state %s", auditsyslogactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("protocolviolations", data["protocolviolations"])
	d.Set("acl", data["acl"])
	d.Set("alg", data["alg"])
	d.Set("appflowexport", data["appflowexport"])
	d.Set("contentinspectionlog", data["contentinspectionlog"])
	d.Set("dateformat", data["dateformat"])
	d.Set("dns", data["dns"])
	d.Set("domainresolvenow", data["domainresolvenow"])
	setToInt("domainresolveretry", d, data["domainresolveretry"])
	d.Set("lbvservername", data["lbvservername"])
	d.Set("logfacility", data["logfacility"])
	d.Set("loglevel", data["loglevel"])
	d.Set("lsn", data["lsn"])
	setToInt("maxlogdatasizetohold", d, data["maxlogdatasizetohold"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("serverdomainname", data["serverdomainname"])
	d.Set("serverip", data["serverip"])
	setToInt("serverport", d, data["serverport"])
	d.Set("sslinterception", data["sslinterception"])
	d.Set("subscriberlog", data["subscriberlog"])
	d.Set("tcp", data["tcp"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("timezone", data["timezone"])
	d.Set("transport", data["transport"])
	d.Set("urlfiltering", data["urlfiltering"])
	d.Set("userdefinedauditlog", data["userdefinedauditlog"])
	d.Set("syslogcompliance", data["syslogcompliance"])
	d.Set("httpauthtoken", data["httpauthtoken"])
	d.Set("httpendpointurl", data["httpendpointurl"])
	d.Set("streamanalytics", data["streamanalytics"])
	if val, ok := data["managementlog"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("managementlog", toStringList(list))
		}
	} else {
		d.Set("managementlog", nil)
	}
	if val, ok := data["mgmtloglevel"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("mgmtloglevel", toStringList(list))
		}
	} else {
		d.Set("mgmtloglevel", nil)
	}

	return nil

}

func updateAuditsyslogactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogactionName := d.Get("name").(string)

	auditsyslogaction := audit.Auditsyslogaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("protocolviolations") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolviolations has changed for auditsyslogaction, starting update")
		auditsyslogaction.Protocolviolations = d.Get("protocolviolations").(string)
		hasChange = true
	}
	if d.HasChange("acl") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Acl = d.Get("acl").(string)
		hasChange = true
	}
	if d.HasChange("alg") {
		log.Printf("[DEBUG]  citrixadc-provider: Alg has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Alg = d.Get("alg").(string)
		hasChange = true
	}
	if d.HasChange("appflowexport") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowexport has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Appflowexport = d.Get("appflowexport").(string)
		hasChange = true
	}
	if d.HasChange("contentinspectionlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Contentinspectionlog has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Contentinspectionlog = d.Get("contentinspectionlog").(string)
		hasChange = true
	}
	if d.HasChange("dateformat") {
		log.Printf("[DEBUG]  citrixadc-provider: Dateformat has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Dateformat = d.Get("dateformat").(string)
		hasChange = true
	}
	if d.HasChange("dns") {
		log.Printf("[DEBUG]  citrixadc-provider: Dns has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Dns = d.Get("dns").(string)
		hasChange = true
	}
	if d.HasChange("domainresolvenow") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolvenow has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Domainresolvenow = d.Get("domainresolvenow").(bool)
		hasChange = true
	}
	if d.HasChange("domainresolveretry") {
		log.Printf("[DEBUG]  citrixadc-provider: Domainresolveretry has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Domainresolveretry = intPtr(d.Get("domainresolveretry").(int))
		hasChange = true
	}
	if d.HasChange("lbvservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbvservername has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Lbvservername = d.Get("lbvservername").(string)
		hasChange = true
	}
	if d.HasChange("logfacility") {
		log.Printf("[DEBUG]  citrixadc-provider: Logfacility has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Logfacility = d.Get("logfacility").(string)
		hasChange = true
	}
	if d.HasChange("loglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Loglevel has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Loglevel = toStringList(loglevelValue(d))
		hasChange = true
	}
	if d.HasChange("lsn") {
		log.Printf("[DEBUG]  citrixadc-provider: Lsn has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Lsn = d.Get("lsn").(string)
		hasChange = true
	}
	if d.HasChange("maxlogdatasizetohold") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxlogdatasizetohold has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Maxlogdatasizetohold = intPtr(d.Get("maxlogdatasizetohold").(int))
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("serverdomainname") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverdomainname has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Serverdomainname = d.Get("serverdomainname").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("sslinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslinterception has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Sslinterception = d.Get("sslinterception").(string)
		hasChange = true
	}
	if d.HasChange("subscriberlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberlog has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Subscriberlog = d.Get("subscriberlog").(string)
		hasChange = true
	}
	if d.HasChange("tcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcp has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Tcp = d.Get("tcp").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("timezone") {
		log.Printf("[DEBUG]  citrixadc-provider: Timezone has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Timezone = d.Get("timezone").(string)
		hasChange = true
	}
	if d.HasChange("transport") {
		log.Printf("[DEBUG]  citrixadc-provider: Transport has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Transport = d.Get("transport").(string)
		hasChange = true
	}
	if d.HasChange("urlfiltering") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlfiltering has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Urlfiltering = d.Get("urlfiltering").(string)
		hasChange = true
	}
	if d.HasChange("userdefinedauditlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdefinedauditlog has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Userdefinedauditlog = d.Get("userdefinedauditlog").(string)
		hasChange = true
	}
	if d.HasChange("syslogcompliance") {
		log.Printf("[DEBUG]  citrixadc-provider: Syslogcompliance has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Syslogcompliance = d.Get("syslogcompliance").(string)
		hasChange = true
	}
	if d.HasChange("httpauthtoken") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpauthtoken has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Httpauthtoken = d.Get("httpauthtoken").(string)
		hasChange = true
	}
	if d.HasChange("httpendpointurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpendpointurl has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Httpendpointurl = d.Get("httpendpointurl").(string)
		hasChange = true
	}
	if d.HasChange("streamanalytics") {
		log.Printf("[DEBUG]  citrixadc-provider: Streamanalytics has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Streamanalytics = d.Get("streamanalytics").(string)
		hasChange = true
	}
	if d.HasChange("managementlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Managementlog has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Managementlog = toStringList(d.Get("managementlog").([]interface{}))
		hasChange = true
	}
	if d.HasChange("mgmtloglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Mgmtloglevel has changed for auditsyslogaction %s, starting update", auditsyslogactionName)
		auditsyslogaction.Mgmtloglevel = toStringList(d.Get("mgmtloglevel").([]interface{}))
		hasChange = true
	}
	if hasChange {
		_, err := client.UpdateResource(service.Auditsyslogaction.Type(), auditsyslogactionName, &auditsyslogaction)
		if err != nil {
			return diag.Errorf("Error updating auditsyslogaction %s", auditsyslogactionName)
		}
	}
	return readAuditsyslogactionFunc(ctx, d, meta)
}

func deleteAuditsyslogactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogactionName := d.Id()
	err := client.DeleteResource(service.Auditsyslogaction.Type(), auditsyslogactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func loglevelValue(d *schema.ResourceData) []interface{} {
	if val, ok := d.GetOk("loglevel"); ok {
		return val.(*schema.Set).List()
	} else {
		return make([]interface{}, 0, 0)
	}
}
