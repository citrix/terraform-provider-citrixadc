package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

type Auditsyslogaction struct {
	/**
	* Name of the syslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog action is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my syslog action" or 'my syslog action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address of the syslog server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* SYSLOG server name as a FQDN.  Mutually exclusive with serverIP/lbVserverName
	*/
	Serverdomainname string `json:"serverdomainname,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.
	*/
	Domainresolveretry int `json:"domainresolveretry,omitempty"`
	/**
	* Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName
	*/
	Lbvservername string `json:"lbvservername,omitempty"`
	/**
	* Port on which the syslog server accepts connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Audit log level, which specifies the types of events to log.
		Available values function as follows:
		* ALL - All events.
		* EMERGENCY - Events that indicate an immediate crisis on the server.
		* ALERT - Events that might require action.
		* CRITICAL - Events that indicate an imminent server crisis.
		* ERROR - Events that indicate some type of error.
		* WARNING - Events that require action in the near future.
		* NOTICE - Events that the administrator should know about.
		* INFORMATIONAL - All but low-level events.
		* DEBUG - All events, in extreme detail.
		* NONE - No events.
	*/
	Loglevel []string `json:"loglevel,omitempty"`
	/**
	* Management log specifies the categories of log files to be exported.
		It use destination and transport from PE params.
		Available values function as follows:
		* ALL - All categories (SHELL, NSMGMT and ACCESS).
		* SHELL -  bash.log, and sh.log.
		* ACCESS - auth.log, nsvpn.log, httpaccess.log, httperror.log, httpaccess-vpn.log and httperror-vpn.log.
		* NSMGMT - notice.log and ns.log.
		* NONE - No logs.
	*/
	Managementlog []string `json:"managementlog,omitempty"`
	/**
	* Management log level, which specifies the types of events to log.
		Available values function as follows:
		* ALL - All events.
		* EMERGENCY - Events that indicate an immediate crisis on the server.
		* ALERT - Events that might require action.
		* CRITICAL - Events that indicate an imminent server crisis.
		* ERROR - Events that indicate some type of error.
		* WARNING - Events that require action in the near future.
		* NOTICE - Events that the administrator should know about.
		* INFORMATIONAL - All but low-level events.
		* DEBUG - All events, in extreme detail.
		* NONE - No events.
	*/
	Mgmtloglevel []string `json:"mgmtloglevel,omitempty"`
	/**
	* Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC. For example, set it to RFC5424 to ensure RFC 5424 compliance
	*/
	Syslogcompliance string `json:"syslogcompliance,omitempty"`
	/**
	* Format of dates in the logs.
		Supported formats are:
		* MMDDYYYY. -U.S. style month/date/year format.
		* DDMMYYYY - European style date/month/year format.
		* YYYYMMDD - ISO style year/month/date format.
	*/
	Dateformat string `json:"dateformat,omitempty"`
	/**
	* Facility value, as defined in RFC 3164, assigned to the log message.
		Log facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.
	*/
	Logfacility string `json:"logfacility,omitempty"`
	/**
	* Log TCP messages.
	*/
	Tcp string `json:"tcp,omitempty"`
	/**
	* Log access control list (ACL) messages.
	*/
	Acl string `json:"acl,omitempty"`
	/**
	* Time zone used for date and timestamps in the logs.
		Supported settings are:
		* GMT_TIME. Coordinated Universal time.
		* LOCAL_TIME. Use the server's timezone setting.
	*/
	Timezone string `json:"timezone,omitempty"`
	/**
	* Log user-configurable log messages to syslog.
		Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.
	*/
	Userdefinedauditlog string `json:"userdefinedauditlog,omitempty"`
	/**
	* Export log messages to AppFlow collectors.
		Appflow collectors are entities to which log messages can be sent so that some action can be performed on them.
	*/
	Appflowexport string `json:"appflowexport,omitempty"`
	/**
	* Log lsn info
	*/
	Lsn string `json:"lsn,omitempty"`
	/**
	* Log alg info
	*/
	Alg string `json:"alg,omitempty"`
	/**
	* Log subscriber session event information
	*/
	Subscriberlog string `json:"subscriberlog,omitempty"`
	/**
	* Transport type used to send auditlogs to syslog server. Default type is UDP.
	*/
	Transport string `json:"transport,omitempty"`
	/**
	* Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorization header is required to be of the form - Splunk <auth-token>.
	*/
	Httpauthtoken string `json:"httpauthtoken,omitempty"`
	/**
	* The URL at which to upload the logs messages on the endpoint
	*/
	Httpendpointurl string `json:"httpendpointurl,omitempty"`
	/**
	* Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Max size of log data that can be held in NSB chain of server info.
	*/
	Maxlogdatasizetohold int `json:"maxlogdatasizetohold,omitempty"`
	/**
	* Log DNS related syslog messages
	*/
	Dns string `json:"dns,omitempty"`
	/**
	* Log Content Inspection event information
	*/
	Contentinspectionlog string `json:"contentinspectionlog,omitempty"`
	/**
	* Name of the network profile.
		The SNIP configured in the network profile will be used as source IP while sending log messages.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Log SSL Interception event information
	*/
	Sslinterception string `json:"sslinterception,omitempty"`
	/**
	* Log URL filtering event information
	*/
	Urlfiltering string `json:"urlfiltering,omitempty"`
	/**
	* Export log stream analytics statistics to syslog server.
	*/
	Streamanalytics string `json:"streamanalytics,omitempty"`
	/**
	* Immediately send a DNS query to resolve the server's domain name.
	*/
	Domainresolvenow bool `json:"domainresolvenow,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}

func resourceCitrixAdcAuditsyslogaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditsyslogactionFunc,
		Read:          readAuditsyslogactionFunc,
		Update:        updateAuditsyslogactionFunc,
		Delete:        deleteAuditsyslogactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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

func createAuditsyslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client

	auditsyslogactionName := d.Get("name").(string)

	auditsyslogaction := Auditsyslogaction{
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Contentinspectionlog: d.Get("contentinspectionlog").(string),
		Dateformat:           d.Get("dateformat").(string),
		Dns:                  d.Get("dns").(string),
		Domainresolvenow:     d.Get("domainresolvenow").(bool),
		Domainresolveretry:   d.Get("domainresolveretry").(int),
		Lbvservername:        d.Get("lbvservername").(string),
		Logfacility:          d.Get("logfacility").(string),
		Loglevel:             toStringList(loglevelValue(d)),
		Lsn:                  d.Get("lsn").(string),
		Maxlogdatasizetohold: d.Get("maxlogdatasizetohold").(int),
		Name:                 d.Get("name").(string),
		Netprofile:           d.Get("netprofile").(string),
		Serverdomainname:     d.Get("serverdomainname").(string),
		Serverip:             d.Get("serverip").(string),
		Serverport:           d.Get("serverport").(int),
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
	}
	if listVal, ok := d.Get("managementlog").([]interface{}); ok {
		auditsyslogaction.Managementlog = toStringList(listVal)
	}
	if listVal, ok := d.Get("mgmtloglevel").([]interface{}); ok {
		auditsyslogaction.Mgmtloglevel = toStringList(listVal)
	}

	_, err := client.AddResource(service.Auditsyslogaction.Type(), auditsyslogactionName, &auditsyslogaction)
	if err != nil {
		return err
	}

	d.SetId(auditsyslogactionName)

	err = readAuditsyslogactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditsyslogaction but we can't read it ?? %s", auditsyslogactionName)
		return nil
	}
	return nil
}

func readAuditsyslogactionFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("acl", data["acl"])
	d.Set("alg", data["alg"])
	d.Set("appflowexport", data["appflowexport"])
	d.Set("contentinspectionlog", data["contentinspectionlog"])
	d.Set("dateformat", data["dateformat"])
	d.Set("dns", data["dns"])
	d.Set("domainresolvenow", data["domainresolvenow"])
	d.Set("domainresolveretry", data["domainresolveretry"])
	d.Set("lbvservername", data["lbvservername"])
	d.Set("logfacility", data["logfacility"])
	d.Set("loglevel", data["loglevel"])
	d.Set("lsn", data["lsn"])
	d.Set("maxlogdatasizetohold", data["maxlogdatasizetohold"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("serverdomainname", data["serverdomainname"])
	d.Set("serverip", data["serverip"])
	d.Set("serverport", data["serverport"])
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

func updateAuditsyslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogactionName := d.Get("name").(string)

	auditsyslogaction := Auditsyslogaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
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
		auditsyslogaction.Domainresolveretry = d.Get("domainresolveretry").(int)
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
		auditsyslogaction.Maxlogdatasizetohold = d.Get("maxlogdatasizetohold").(int)
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
		auditsyslogaction.Serverport = d.Get("serverport").(int)
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
			return fmt.Errorf("Error updating auditsyslogaction %s", auditsyslogactionName)
		}
	}
	return readAuditsyslogactionFunc(d, meta)
}

func deleteAuditsyslogactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditsyslogactionFunc")
	client := meta.(*NetScalerNitroClient).client
	auditsyslogactionName := d.Id()
	err := client.DeleteResource(service.Auditsyslogaction.Type(), auditsyslogactionName)
	if err != nil {
		return err
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
