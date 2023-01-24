package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/analytics"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAnalyticsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAnalyticsprofileFunc,
		Read:          readAnalyticsprofileFunc,
		Update:        updateAnalyticsprofileFunc,
		Delete:        deleteAnalyticsprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"allhttpheaders": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auditlogs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collectors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cqareporting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grpcstatus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpauthentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpclientsidemeasurements": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpcontenttype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpcookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpdomainname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httphost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httplocation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httppagetracking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpreferer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpsetcookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpsetcookie2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpurlquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpuseragent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpvia": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpxforwardedforheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"integratedcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metrics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outputmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpburstreporting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urlcategory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAnalyticsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAnalyticsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofileName := d.Get("name").(string)
	analyticsprofile := analytics.Analyticsprofile{
		Allhttpheaders:             d.Get("allhttpheaders").(string),
		Auditlogs:                  d.Get("auditlogs").(string),
		Collectors:                 d.Get("collectors").(string),
		Cqareporting:               d.Get("cqareporting").(string),
		Events:                     d.Get("events").(string),
		Grpcstatus:                 d.Get("grpcstatus").(string),
		Httpauthentication:         d.Get("httpauthentication").(string),
		Httpclientsidemeasurements: d.Get("httpclientsidemeasurements").(string),
		Httpcontenttype:            d.Get("httpcontenttype").(string),
		Httpcookie:                 d.Get("httpcookie").(string),
		Httpdomainname:             d.Get("httpdomainname").(string),
		Httphost:                   d.Get("httphost").(string),
		Httplocation:               d.Get("httplocation").(string),
		Httpmethod:                 d.Get("httpmethod").(string),
		Httppagetracking:           d.Get("httppagetracking").(string),
		Httpreferer:                d.Get("httpreferer").(string),
		Httpsetcookie:              d.Get("httpsetcookie").(string),
		Httpsetcookie2:             d.Get("httpsetcookie2").(string),
		Httpurl:                    d.Get("httpurl").(string),
		Httpurlquery:               d.Get("httpurlquery").(string),
		Httpuseragent:              d.Get("httpuseragent").(string),
		Httpvia:                    d.Get("httpvia").(string),
		Httpxforwardedforheader:    d.Get("httpxforwardedforheader").(string),
		Integratedcache:            d.Get("integratedcache").(string),
		Metrics:                    d.Get("metrics").(string),
		Name:                       d.Get("name").(string),
		Outputmode:                 d.Get("outputmode").(string),
		Tcpburstreporting:          d.Get("tcpburstreporting").(string),
		Type:                       d.Get("type").(string),
		Urlcategory:                d.Get("urlcategory").(string),
	}

	_, err := client.AddResource("analyticsprofile", analyticsprofileName, &analyticsprofile)
	if err != nil {
		return err
	}

	d.SetId(analyticsprofileName)

	err = readAnalyticsprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this analyticsprofile but we can't read it ?? %s", analyticsprofileName)
		return nil
	}
	return nil
}

func readAnalyticsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAnalyticsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading analyticsprofile state %s", analyticsprofileName)
	data, err := client.FindResource("analyticsprofile", analyticsprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing analyticsprofile state %s", analyticsprofileName)
		d.SetId("")
		return nil
	}
	d.Set("allhttpheaders", data["allhttpheaders"])
	d.Set("auditlogs", data["auditlogs"])
	d.Set("collectors", data["collectors"])
	d.Set("cqareporting", data["cqareporting"])
	d.Set("events", data["events"])
	d.Set("grpcstatus", data["grpcstatus"])
	d.Set("httpauthentication", data["httpauthentication"])
	d.Set("httpclientsidemeasurements", data["httpclientsidemeasurements"])
	d.Set("httpcontenttype", data["httpcontenttype"])
	d.Set("httpcookie", data["httpcookie"])
	d.Set("httpdomainname", data["httpdomainname"])
	d.Set("httphost", data["httphost"])
	d.Set("httplocation", data["httplocation"])
	d.Set("httpmethod", data["httpmethod"])
	d.Set("httppagetracking", data["httppagetracking"])
	d.Set("httpreferer", data["httpreferer"])
	d.Set("httpsetcookie", data["httpsetcookie"])
	d.Set("httpsetcookie2", data["httpsetcookie2"])
	d.Set("httpurl", data["httpurl"])
	d.Set("httpurlquery", data["httpurlquery"])
	d.Set("httpuseragent", data["httpuseragent"])
	d.Set("httpvia", data["httpvia"])
	d.Set("httpxforwardedforheader", data["httpxforwardedforheader"])
	d.Set("integratedcache", data["integratedcache"])
	d.Set("metrics", data["metrics"])
	d.Set("name", data["name"])
	d.Set("outputmode", data["outputmode"])
	d.Set("tcpburstreporting", data["tcpburstreporting"])
	d.Set("type", data["type"])
	d.Set("urlcategory", data["urlcategory"])

	return nil

}

func updateAnalyticsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAnalyticsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofileName := d.Get("name").(string)

	analyticsprofile := analytics.Analyticsprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("allhttpheaders") {
		log.Printf("[DEBUG]  citrixadc-provider: Allhttpheaders has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Allhttpheaders = d.Get("allhttpheaders").(string)
		hasChange = true
	}
	if d.HasChange("auditlogs") {
		log.Printf("[DEBUG]  citrixadc-provider: Auditlogs has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Auditlogs = d.Get("auditlogs").(string)
		hasChange = true
	}
	if d.HasChange("collectors") {
		log.Printf("[DEBUG]  citrixadc-provider: Collectors has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Collectors = d.Get("collectors").(string)
		hasChange = true
	}
	if d.HasChange("cqareporting") {
		log.Printf("[DEBUG]  citrixadc-provider: Cqareporting has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Cqareporting = d.Get("cqareporting").(string)
		hasChange = true
	}
	if d.HasChange("events") {
		log.Printf("[DEBUG]  citrixadc-provider: Events has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Events = d.Get("events").(string)
		hasChange = true
	}
	if d.HasChange("grpcstatus") {
		log.Printf("[DEBUG]  citrixadc-provider: Grpcstatus has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Grpcstatus = d.Get("grpcstatus").(string)
		hasChange = true
	}
	if d.HasChange("httpauthentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpauthentication has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpauthentication = d.Get("httpauthentication").(string)
		hasChange = true
	}
	if d.HasChange("httpclientsidemeasurements") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpclientsidemeasurements has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpclientsidemeasurements = d.Get("httpclientsidemeasurements").(string)
		hasChange = true
	}
	if d.HasChange("httpcontenttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpcontenttype has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpcontenttype = d.Get("httpcontenttype").(string)
		hasChange = true
	}
	if d.HasChange("httpcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpcookie has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpcookie = d.Get("httpcookie").(string)
		hasChange = true
	}
	if d.HasChange("httpdomainname") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpdomainname has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpdomainname = d.Get("httpdomainname").(string)
		hasChange = true
	}
	if d.HasChange("httphost") {
		log.Printf("[DEBUG]  citrixadc-provider: Httphost has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httphost = d.Get("httphost").(string)
		hasChange = true
	}
	if d.HasChange("httplocation") {
		log.Printf("[DEBUG]  citrixadc-provider: Httplocation has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httplocation = d.Get("httplocation").(string)
		hasChange = true
	}
	if d.HasChange("httpmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpmethod has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpmethod = d.Get("httpmethod").(string)
		hasChange = true
	}
	if d.HasChange("httppagetracking") {
		log.Printf("[DEBUG]  citrixadc-provider: Httppagetracking has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httppagetracking = d.Get("httppagetracking").(string)
		hasChange = true
	}
	if d.HasChange("httpreferer") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpreferer has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpreferer = d.Get("httpreferer").(string)
		hasChange = true
	}
	if d.HasChange("httpsetcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpsetcookie has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpsetcookie = d.Get("httpsetcookie").(string)
		hasChange = true
	}
	if d.HasChange("httpsetcookie2") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpsetcookie2 has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpsetcookie2 = d.Get("httpsetcookie2").(string)
		hasChange = true
	}
	if d.HasChange("httpurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpurl has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpurl = d.Get("httpurl").(string)
		hasChange = true
	}
	if d.HasChange("httpurlquery") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpurlquery has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpurlquery = d.Get("httpurlquery").(string)
		hasChange = true
	}
	if d.HasChange("httpuseragent") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpuseragent has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpuseragent = d.Get("httpuseragent").(string)
		hasChange = true
	}
	if d.HasChange("httpvia") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpvia has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpvia = d.Get("httpvia").(string)
		hasChange = true
	}
	if d.HasChange("httpxforwardedforheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpxforwardedforheader has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Httpxforwardedforheader = d.Get("httpxforwardedforheader").(string)
		hasChange = true
	}
	if d.HasChange("integratedcache") {
		log.Printf("[DEBUG]  citrixadc-provider: Integratedcache has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Integratedcache = d.Get("integratedcache").(string)
		hasChange = true
	}
	if d.HasChange("metrics") {
		log.Printf("[DEBUG]  citrixadc-provider: Metrics has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Metrics = d.Get("metrics").(string)
		hasChange = true
	}
	if d.HasChange("outputmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Outputmode has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Outputmode = d.Get("outputmode").(string)
		hasChange = true
	}
	if d.HasChange("tcpburstreporting") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpburstreporting has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Tcpburstreporting = d.Get("tcpburstreporting").(string)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("urlcategory") {
		log.Printf("[DEBUG]  citrixadc-provider: Urlcategory has changed for analyticsprofile %s, starting update", analyticsprofileName)
		analyticsprofile.Urlcategory = d.Get("urlcategory").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("analyticsprofile", analyticsprofileName, &analyticsprofile)
		if err != nil {
			return fmt.Errorf("Error updating analyticsprofile %s", analyticsprofileName)
		}
	}
	return readAnalyticsprofileFunc(d, meta)
}

func deleteAnalyticsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAnalyticsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	analyticsprofileName := d.Id()
	err := client.DeleteResource("analyticsprofile", analyticsprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
