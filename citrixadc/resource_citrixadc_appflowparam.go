package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appflow"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppflowparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppflowparamFunc,
		Read:          readAppflowparamFunc,
		Update:        updateAppflowparamFunc,
		Delete:        deleteAppflowparamFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"aaausername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"analyticsauthtoken": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appnamerefresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auditlogs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clienttrafficonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connectionchaining": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cqareporting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distributedtracing": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disttracingsamplingrate": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"emailaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"events": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flowrecordinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gxsessionreporting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpauthorization": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpcontenttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpdomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httphost": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httplocation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpquerywithurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpreferer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpsetcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpsetcookie2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpuseragent": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpvia": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpxforwardedfor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identifiername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identifiersessionname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logstreamovernsip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lsnlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"observationdomainid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"observationdomainname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"observationpointid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"securityinsightrecordinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"securityinsighttraffic": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skipcacheredirectionhttptransaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberawareness": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberidobfuscation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscriberidobfuscationalgo": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpattackcounterinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"templaterefresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeseriesovernsip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udppmtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"urlcategory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usagerecordinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"videoinsight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websaasappusagereporting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAppflowparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppflowparamFunc")
	client := meta.(*NetScalerNitroClient).client

	appflowparamName := resource.PrefixedUniqueId("tf-appflowparam-")

	appflowparam := appflow.Appflowparam{
		Aaausername:                         d.Get("aaausername").(string),
		Analyticsauthtoken:                  d.Get("analyticsauthtoken").(string),
		Appnamerefresh:                      d.Get("appnamerefresh").(int),
		Auditlogs:                           d.Get("auditlogs").(string),
		Cacheinsight:                        d.Get("cacheinsight").(string),
		Clienttrafficonly:                   d.Get("clienttrafficonly").(string),
		Connectionchaining:                  d.Get("connectionchaining").(string),
		Cqareporting:                        d.Get("cqareporting").(string),
		Distributedtracing:                  d.Get("distributedtracing").(string),
		Disttracingsamplingrate:             d.Get("disttracingsamplingrate").(int),
		Emailaddress:                        d.Get("emailaddress").(string),
		Events:                              d.Get("events").(string),
		Flowrecordinterval:                  d.Get("flowrecordinterval").(int),
		Gxsessionreporting:                  d.Get("gxsessionreporting").(string),
		Httpauthorization:                   d.Get("httpauthorization").(string),
		Httpcontenttype:                     d.Get("httpcontenttype").(string),
		Httpcookie:                          d.Get("httpcookie").(string),
		Httpdomain:                          d.Get("httpdomain").(string),
		Httphost:                            d.Get("httphost").(string),
		Httplocation:                        d.Get("httplocation").(string),
		Httpmethod:                          d.Get("httpmethod").(string),
		Httpquerywithurl:                    d.Get("httpquerywithurl").(string),
		Httpreferer:                         d.Get("httpreferer").(string),
		Httpsetcookie:                       d.Get("httpsetcookie").(string),
		Httpsetcookie2:                      d.Get("httpsetcookie2").(string),
		Httpurl:                             d.Get("httpurl").(string),
		Httpuseragent:                       d.Get("httpuseragent").(string),
		Httpvia:                             d.Get("httpvia").(string),
		Httpxforwardedfor:                   d.Get("httpxforwardedfor").(string),
		Identifiername:                      d.Get("identifiername").(string),
		Identifiersessionname:               d.Get("identifiersessionname").(string),
		Logstreamovernsip:                   d.Get("logstreamovernsip").(string),
		Lsnlogging:                          d.Get("lsnlogging").(string),
		Metrics:                             d.Get("metrics").(string),
		Observationdomainid:                 d.Get("observationdomainid").(int),
		Observationdomainname:               d.Get("observationdomainname").(string),
		Observationpointid:                  d.Get("observationpointid").(int),
		Securityinsightrecordinterval:       d.Get("securityinsightrecordinterval").(int),
		Securityinsighttraffic:              d.Get("securityinsighttraffic").(string),
		Skipcacheredirectionhttptransaction: d.Get("skipcacheredirectionhttptransaction").(string),
		Subscriberawareness:                 d.Get("subscriberawareness").(string),
		Subscriberidobfuscation:             d.Get("subscriberidobfuscation").(string),
		Subscriberidobfuscationalgo:         d.Get("subscriberidobfuscationalgo").(string),
		Tcpattackcounterinterval:            d.Get("tcpattackcounterinterval").(int),
		Templaterefresh:                     d.Get("templaterefresh").(int),
		Timeseriesovernsip:                  d.Get("timeseriesovernsip").(string),
		Udppmtu:                             d.Get("udppmtu").(int),
		Urlcategory:                         d.Get("urlcategory").(string),
		Usagerecordinterval:                 d.Get("usagerecordinterval").(int),
		Videoinsight:                        d.Get("videoinsight").(string),
		Websaasappusagereporting:            d.Get("websaasappusagereporting").(string),
	}

	err := client.UpdateUnnamedResource(service.Appflowparam.Type(), &appflowparam)
	if err != nil {
		return err
	}

	d.SetId(appflowparamName)

	err = readAppflowparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appflowparam but we can't read it ??")
		return nil
	}
	return nil
}

func readAppflowparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppflowparamFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading appflowparam state")
	data, err := client.FindResource(service.Appflowparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appflowparam state")
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("aaausername", data["aaausername"])
	d.Set("analyticsauthtoken", data["analyticsauthtoken"])
	d.Set("appnamerefresh", data["appnamerefresh"])
	d.Set("auditlogs", data["auditlogs"])
	d.Set("cacheinsight", data["cacheinsight"])
	d.Set("clienttrafficonly", data["clienttrafficonly"])
	d.Set("connectionchaining", data["connectionchaining"])
	d.Set("cqareporting", data["cqareporting"])
	d.Set("distributedtracing", data["distributedtracing"])
	d.Set("disttracingsamplingrate", data["disttracingsamplingrate"])
	d.Set("emailaddress", data["emailaddress"])
	d.Set("events", data["events"])
	d.Set("flowrecordinterval", data["flowrecordinterval"])
	d.Set("gxsessionreporting", data["gxsessionreporting"])
	d.Set("httpauthorization", data["httpauthorization"])
	d.Set("httpcontenttype", data["httpcontenttype"])
	d.Set("httpcookie", data["httpcookie"])
	d.Set("httpdomain", data["httpdomain"])
	d.Set("httphost", data["httphost"])
	d.Set("httplocation", data["httplocation"])
	d.Set("httpmethod", data["httpmethod"])
	d.Set("httpquerywithurl", data["httpquerywithurl"])
	d.Set("httpreferer", data["httpreferer"])
	d.Set("httpsetcookie", data["httpsetcookie"])
	d.Set("httpsetcookie2", data["httpsetcookie2"])
	d.Set("httpurl", data["httpurl"])
	d.Set("httpuseragent", data["httpuseragent"])
	d.Set("httpvia", data["httpvia"])
	d.Set("httpxforwardedfor", data["httpxforwardedfor"])
	d.Set("identifiername", data["identifiername"])
	d.Set("identifiersessionname", data["identifiersessionname"])
	d.Set("logstreamovernsip", data["logstreamovernsip"])
	d.Set("lsnlogging", data["lsnlogging"])
	d.Set("metrics", data["metrics"])
	d.Set("observationdomainid", data["observationdomainid"])
	d.Set("observationdomainname", data["observationdomainname"])
	d.Set("observationpointid", data["observationpointid"])
	d.Set("securityinsightrecordinterval", data["securityinsightrecordinterval"])
	d.Set("securityinsighttraffic", data["securityinsighttraffic"])
	d.Set("skipcacheredirectionhttptransaction", data["skipcacheredirectionhttptransaction"])
	d.Set("subscriberawareness", data["subscriberawareness"])
	d.Set("subscriberidobfuscation", data["subscriberidobfuscation"])
	d.Set("subscriberidobfuscationalgo", data["subscriberidobfuscationalgo"])
	d.Set("tcpattackcounterinterval", data["tcpattackcounterinterval"])
	d.Set("templaterefresh", data["templaterefresh"])
	d.Set("timeseriesovernsip", data["timeseriesovernsip"])
	d.Set("udppmtu", data["udppmtu"])
	d.Set("urlcategory", data["urlcategory"])
	d.Set("usagerecordinterval", data["usagerecordinterval"])
	d.Set("videoinsight", data["videoinsight"])
	d.Set("websaasappusagereporting", data["websaasappusagereporting"])

	return nil

}

func updateAppflowparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppflowparamFunc")
	client := meta.(*NetScalerNitroClient).client
	appflowparam := appflow.Appflowparam{}

	hasChange := false
	if d.HasChange("aaausername") {
		log.Printf("[DEBUG]  citrixadc-provider: Aaausername has changed for appflowparam, starting update")
		appflowparam.Aaausername = d.Get("aaausername").(string)
		hasChange = true
	}
	if d.HasChange("analyticsauthtoken") {
		log.Printf("[DEBUG]  citrixadc-provider: Analyticsauthtoken has changed for appflowparam, starting update")
		appflowparam.Analyticsauthtoken = d.Get("analyticsauthtoken").(string)
		hasChange = true
	}
	if d.HasChange("appnamerefresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Appnamerefresh has changed for appflowparam, starting update")
		appflowparam.Appnamerefresh = d.Get("appnamerefresh").(int)
		hasChange = true
	}
	if d.HasChange("auditlogs") {
		log.Printf("[DEBUG]  citrixadc-provider: Auditlogs has changed for appflowparam, starting update")
		appflowparam.Auditlogs = d.Get("auditlogs").(string)
		hasChange = true
	}
	if d.HasChange("cacheinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheinsight has changed for appflowparam, starting update")
		appflowparam.Cacheinsight = d.Get("cacheinsight").(string)
		hasChange = true
	}
	if d.HasChange("clienttrafficonly") {
		log.Printf("[DEBUG]  citrixadc-provider: Clienttrafficonly has changed for appflowparam, starting update")
		appflowparam.Clienttrafficonly = d.Get("clienttrafficonly").(string)
		hasChange = true
	}
	if d.HasChange("connectionchaining") {
		log.Printf("[DEBUG]  citrixadc-provider: Connectionchaining has changed for appflowparam, starting update")
		appflowparam.Connectionchaining = d.Get("connectionchaining").(string)
		hasChange = true
	}
	if d.HasChange("cqareporting") {
		log.Printf("[DEBUG]  citrixadc-provider: Cqareporting has changed for appflowparam, starting update")
		appflowparam.Cqareporting = d.Get("cqareporting").(string)
		hasChange = true
	}
	if d.HasChange("distributedtracing") {
		log.Printf("[DEBUG]  citrixadc-provider: Distributedtracing has changed for appflowparam, starting update")
		appflowparam.Distributedtracing = d.Get("distributedtracing").(string)
		hasChange = true
	}
	if d.HasChange("disttracingsamplingrate") {
		log.Printf("[DEBUG]  citrixadc-provider: Disttracingsamplingrate has changed for appflowparam, starting update")
		appflowparam.Disttracingsamplingrate = d.Get("disttracingsamplingrate").(int)
		hasChange = true
	}
	if d.HasChange("emailaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Emailaddress has changed for appflowparam, starting update")
		appflowparam.Emailaddress = d.Get("emailaddress").(string)
		hasChange = true
	}
	if d.HasChange("events") {
		log.Printf("[DEBUG]  citrixadc-provider: Events has changed for appflowparam, starting update")
		appflowparam.Events = d.Get("events").(string)
		hasChange = true
	}
	if d.HasChange("flowrecordinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Flowrecordinterval has changed for appflowparam, starting update")
		appflowparam.Flowrecordinterval = d.Get("flowrecordinterval").(int)
		hasChange = true
	}
	if d.HasChange("gxsessionreporting") {
		log.Printf("[DEBUG]  citrixadc-provider: Gxsessionreporting has changed for appflowparam, starting update")
		appflowparam.Gxsessionreporting = d.Get("gxsessionreporting").(string)
		hasChange = true
	}
	if d.HasChange("httpauthorization") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpauthorization has changed for appflowparam, starting update")
		appflowparam.Httpauthorization = d.Get("httpauthorization").(string)
		hasChange = true
	}
	if d.HasChange("httpcontenttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpcontenttype has changed for appflowparam, starting update")
		appflowparam.Httpcontenttype = d.Get("httpcontenttype").(string)
		hasChange = true
	}
	if d.HasChange("httpcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpcookie has changed for appflowparam, starting update")
		appflowparam.Httpcookie = d.Get("httpcookie").(string)
		hasChange = true
	}
	if d.HasChange("httpdomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpdomain has changed for appflowparam, starting update")
		appflowparam.Httpdomain = d.Get("httpdomain").(string)
		hasChange = true
	}
	if d.HasChange("httphost") {
		log.Printf("[DEBUG]  citrixadc-provider: Httphost has changed for appflowparam, starting update")
		appflowparam.Httphost = d.Get("httphost").(string)
		hasChange = true
	}
	if d.HasChange("httplocation") {
		log.Printf("[DEBUG]  citrixadc-provider: Httplocation has changed for appflowparam, starting update")
		appflowparam.Httplocation = d.Get("httplocation").(string)
		hasChange = true
	}
	if d.HasChange("httpmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpmethod has changed for appflowparam, starting update")
		appflowparam.Httpmethod = d.Get("httpmethod").(string)
		hasChange = true
	}
	if d.HasChange("httpquerywithurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpquerywithurl has changed for appflowparam, starting update")
		appflowparam.Httpquerywithurl = d.Get("httpquerywithurl").(string)
		hasChange = true
	}
	if d.HasChange("httpreferer") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpreferer has changed for appflowparam, starting update")
		appflowparam.Httpreferer = d.Get("httpreferer").(string)
		hasChange = true
	}
	if d.HasChange("httpsetcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpsetcookie has changed for appflowparam, starting update")
		appflowparam.Httpsetcookie = d.Get("httpsetcookie").(string)
		hasChange = true
	}
	if d.HasChange("httpsetcookie2") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpsetcookie2 has changed for appflowparam, starting update")
		appflowparam.Httpsetcookie2 = d.Get("httpsetcookie2").(string)
		hasChange = true
	}
	if d.HasChange("httpurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpurl has changed for appflowparam, starting update")
		appflowparam.Httpurl = d.Get("httpurl").(string)
		hasChange = true
	}
	if d.HasChange("httpuseragent") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpuseragent has changed for appflowparam, starting update")
		appflowparam.Httpuseragent = d.Get("httpuseragent").(string)
		hasChange = true
	}
	if d.HasChange("httpvia") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpvia has changed for appflowparam, starting update")
		appflowparam.Httpvia = d.Get("httpvia").(string)
		hasChange = true
	}
	if d.HasChange("httpxforwardedfor") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpxforwardedfor has changed for appflowparam, starting update")
		appflowparam.Httpxforwardedfor = d.Get("httpxforwardedfor").(string)
		hasChange = true
	}
	if d.HasChange("identifiername") {
		log.Printf("[DEBUG]  citrixadc-provider: Identifiername has changed for appflowparam, starting update")
		appflowparam.Identifiername = d.Get("identifiername").(string)
		hasChange = true
	}
	if d.HasChange("identifiersessionname") {
		log.Printf("[DEBUG]  citrixadc-provider: Identifiersessionname has changed for appflowparam, starting update")
		appflowparam.Identifiersessionname = d.Get("identifiersessionname").(string)
		hasChange = true
	}
	if d.HasChange("logstreamovernsip") {
		log.Printf("[DEBUG]  citrixadc-provider: Logstreamovernsip has changed for appflowparam, starting update")
		appflowparam.Logstreamovernsip = d.Get("logstreamovernsip").(string)
		hasChange = true
	}
	if d.HasChange("lsnlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Lsnlogging has changed for appflowparam, starting update")
		appflowparam.Lsnlogging = d.Get("lsnlogging").(string)
		hasChange = true
	}
	if d.HasChange("metrics") {
		log.Printf("[DEBUG]  citrixadc-provider: Metrics has changed for appflowparam, starting update")
		appflowparam.Metrics = d.Get("metrics").(string)
		hasChange = true
	}
	if d.HasChange("observationdomainid") {
		log.Printf("[DEBUG]  citrixadc-provider: Observationdomainid has changed for appflowparam, starting update")
		appflowparam.Observationdomainid = d.Get("observationdomainid").(int)
		hasChange = true
	}
	if d.HasChange("observationdomainname") {
		log.Printf("[DEBUG]  citrixadc-provider: Observationdomainname has changed for appflowparam, starting update")
		appflowparam.Observationdomainname = d.Get("observationdomainname").(string)
		hasChange = true
	}
	if d.HasChange("observationpointid") {
		log.Printf("[DEBUG]  citrixadc-provider: Observationpointid has changed for appflowparam, starting update")
		appflowparam.Observationpointid = d.Get("observationpointid").(int)
		hasChange = true
	}
	if d.HasChange("securityinsightrecordinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Securityinsightrecordinterval has changed for appflowparam, starting update")
		appflowparam.Securityinsightrecordinterval = d.Get("securityinsightrecordinterval").(int)
		hasChange = true
	}
	if d.HasChange("securityinsighttraffic") {
		log.Printf("[DEBUG]  citrixadc-provider: Securityinsighttraffic has changed for appflowparam, starting update")
		appflowparam.Securityinsighttraffic = d.Get("securityinsighttraffic").(string)
		hasChange = true
	}
	if d.HasChange("skipcacheredirectionhttptransaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Skipcacheredirectionhttptransaction has changed for appflowparam, starting update")
		appflowparam.Skipcacheredirectionhttptransaction = d.Get("skipcacheredirectionhttptransaction").(string)
		hasChange = true
	}
	if d.HasChange("subscriberawareness") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberawareness has changed for appflowparam, starting update")
		appflowparam.Subscriberawareness = d.Get("subscriberawareness").(string)
		hasChange = true
	}
	if d.HasChange("subscriberidobfuscation") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberidobfuscation has changed for appflowparam, starting update")
		appflowparam.Subscriberidobfuscation = d.Get("subscriberidobfuscation").(string)
		hasChange = true
	}
	if d.HasChange("subscriberidobfuscationalgo") {
		log.Printf("[DEBUG]  citrixadc-provider: Subscriberidobfuscationalgo has changed for appflowparam, starting update")
		appflowparam.Subscriberidobfuscationalgo = d.Get("subscriberidobfuscationalgo").(string)
		hasChange = true
	}
	if d.HasChange("tcpattackcounterinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpattackcounterinterval has changed for appflowparam, starting update")
		appflowparam.Tcpattackcounterinterval = d.Get("tcpattackcounterinterval").(int)
		hasChange = true
	}
	if d.HasChange("templaterefresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Templaterefresh has changed for appflowparam, starting update")
		appflowparam.Templaterefresh = d.Get("templaterefresh").(int)
		hasChange = true
	}
	if d.HasChange("timeseriesovernsip") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeseriesovernsip has changed for appflowparam, starting update")
		appflowparam.Timeseriesovernsip = d.Get("timeseriesovernsip").(string)
		hasChange = true
	}
	if d.HasChange("udppmtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Udppmtu has changed for appflowparam, starting update")
		appflowparam.Udppmtu = d.Get("udppmtu").(int)
		hasChange = true
	}
	if d.HasChange("urlcategory") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlcategory has changed for appflowparam, starting update")
		appflowparam.Urlcategory = d.Get("urlcategory").(string)
		hasChange = true
	}
	if d.HasChange("usagerecordinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Usagerecordinterval has changed for appflowparam, starting update")
		appflowparam.Usagerecordinterval = d.Get("usagerecordinterval").(int)
		hasChange = true
	}
	if d.HasChange("videoinsight") {
		log.Printf("[DEBUG]  citrixadc-provider: Videoinsight has changed for appflowparam, starting update")
		appflowparam.Videoinsight = d.Get("videoinsight").(string)
		hasChange = true
	}
	if d.HasChange("websaasappusagereporting") {
		log.Printf("[DEBUG]  citrixadc-provider: Websaasappusagereporting has changed for appflowparam, starting update")
		appflowparam.Websaasappusagereporting = d.Get("websaasappusagereporting").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appflowparam.Type(), &appflowparam)
		if err != nil {
			return fmt.Errorf("Error updating appflowparam")
		}
	}
	return readAppflowparamFunc(d, meta)
}

func deleteAppflowparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppflowparamFunc")
	// appflow parameter does not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
