package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

/**
* Configuration for DNS profile resource.
 */
type Dnsprofile struct {
	/**
	* Name of the DNS profile
	 */
	Dnsprofilename string `json:"dnsprofilename,omitempty"`
	/**
	* DNS recursive resolution; if enabled, will do recursive resolution for DNS query when the profile is associated with ADNS service, CS Vserver and DNS action
	 */
	Recursiveresolution string `json:"recursiveresolution,omitempty"`
	/**
	* DNS query logging; if enabled, DNS query information such as DNS query id, DNS query flags , DNS domain name and DNS query type will be logged
	 */
	Dnsquerylogging string `json:"dnsquerylogging,omitempty"`
	/**
	* DNS answer section; if enabled, answer section in the response will be logged.
	 */
	Dnsanswerseclogging string `json:"dnsanswerseclogging,omitempty"`
	/**
	* DNS extended logging; if enabled, authority and additional section in the response will be logged.
	 */
	Dnsextendedlogging string `json:"dnsextendedlogging,omitempty"`
	/**
	* DNS error logging; if enabled, whenever error is encountered in DNS module reason for the error will be logged.
	 */
	Dnserrorlogging string `json:"dnserrorlogging,omitempty"`
	/**
	* Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.
	 */
	Cacherecords string `json:"cacherecords,omitempty"`
	/**
	* Cache negative responses in the DNS cache. When disabled, the appliance stops caching negative responses except referral records. This applies to all configurations - proxy, end resolver, and forwarder. However, cached responses are not flushed. The appliance does not serve negative responses from the cache until this parameter is enabled again.
	 */
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	/**
	* Drop the DNS requests containing multiple queries. When enabled, DNS requests containing multiple queries will be dropped. In case of proxy configuration by default the DNS request containing multiple queries is forwarded to the backend and in case of ADNS and Resolver configuration NOCODE error response will be sent to the client.
	 */
	Dropmultiqueryrequest string `json:"dropmultiqueryrequest,omitempty"`
	/**
	* Cache DNS responses with EDNS Client Subnet(ECS) option in the DNS cache. When disabled, the appliance stops caching responses with ECS option. This is relevant to proxy configuration. Enabling/disabling support of ECS option when Citrix ADC is authoritative for a GSLB domain is supported using a knob in GSLB vserver. In all other modes, ECS option is ignored.
	 */
	Cacheecsresponses string `json:"cacheecsresponses,omitempty"`
	/**
	* Insert ECS Option on DNS query
	 */
	Insertecs string `json:"insertecs,omitempty"`
	/**
	* Replace ECS Option on DNS query
	 */
	Replaceecs string `json:"replaceecs,omitempty"`
	/**
	* The maximum ecs prefix length that will be cached
	 */
	Maxcacheableecsprefixlength int `json:"maxcacheableecsprefixlength,omitempty"`
	/**
	* The maximum ecs prefix length that will be cached for IPv6 subnets
	 */
	Maxcacheableecsprefixlength6 int `json:"maxcacheableecsprefixlength6,omitempty"`

	//------- Read only Parameter ---------;

	Referencecount string `json:"referencecount,omitempty"`
}

func resourceCitrixAdcDnsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsprofileFunc,
		Read:          readDnsprofileFunc,
		Update:        updateDnsprofileFunc,
		Delete:        deleteDnsprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cacheecsresponses": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachenegativeresponses": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacherecords": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsanswerseclogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnserrorlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsextendedlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dnsquerylogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropmultiqueryrequest": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recursiveresolution": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertecs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replaceecs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxcacheableecsprefixlength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxcacheableecsprefixlength6": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Get("dnsprofilename").(string)
	dnsprofile := Dnsprofile{
		Cacheecsresponses:            d.Get("cacheecsresponses").(string),
		Cachenegativeresponses:       d.Get("cachenegativeresponses").(string),
		Cacherecords:                 d.Get("cacherecords").(string),
		Dnsanswerseclogging:          d.Get("dnsanswerseclogging").(string),
		Dnserrorlogging:              d.Get("dnserrorlogging").(string),
		Dnsextendedlogging:           d.Get("dnsextendedlogging").(string),
		Dnsprofilename:               d.Get("dnsprofilename").(string),
		Dnsquerylogging:              d.Get("dnsquerylogging").(string),
		Dropmultiqueryrequest:        d.Get("dropmultiqueryrequest").(string),
		Recursiveresolution:          d.Get("recursiveresolution").(string),
		Insertecs:                    d.Get("insertecs").(string),
		Replaceecs:                   d.Get("replaceecs").(string),
		Maxcacheableecsprefixlength:  d.Get("maxcacheableecsprefixlength").(int),
		Maxcacheableecsprefixlength6: d.Get("maxcacheableecsprefixlength6").(int),
	}

	_, err := client.AddResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
	if err != nil {
		return err
	}

	d.SetId(dnsprofileName)

	err = readDnsprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsprofile but we can't read it ?? %s", dnsprofileName)
		return nil
	}
	return nil
}

func readDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnsprofile state %s", dnsprofileName)
	data, err := client.FindResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnsprofile state %s", dnsprofileName)
		d.SetId("")
		return nil
	}
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("cacheecsresponses", data["cacheecsresponses"])
	d.Set("cachenegativeresponses", data["cachenegativeresponses"])
	d.Set("cacherecords", data["cacherecords"])
	d.Set("dnsanswerseclogging", data["dnsanswerseclogging"])
	d.Set("dnserrorlogging", data["dnserrorlogging"])
	d.Set("dnsextendedlogging", data["dnsextendedlogging"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("dnsquerylogging", data["dnsquerylogging"])
	d.Set("dropmultiqueryrequest", data["dropmultiqueryrequest"])
	d.Set("recursiveresolution", data["recursiveresolution"])
	d.Set("insertecs", data["insertecs"])
	d.Set("replaceecs", data["replaceecs"])
	d.Set("maxcacheableecsprefixlength", data["maxcacheableecsprefixlength"])
	d.Set("maxcacheableecsprefixlength6", data["maxcacheableecsprefixlength6"])

	return nil

}

func updateDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Get("dnsprofilename").(string)

	dnsprofile := Dnsprofile{
		Dnsprofilename: d.Get("dnsprofilename").(string),
	}
	hasChange := false
	if d.HasChange("cacheecsresponses") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheecsresponses has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cacheecsresponses = d.Get("cacheecsresponses").(string)
		hasChange = true
	}
	if d.HasChange("cachenegativeresponses") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachenegativeresponses has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cachenegativeresponses = d.Get("cachenegativeresponses").(string)
		hasChange = true
	}
	if d.HasChange("cacherecords") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacherecords has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Cacherecords = d.Get("cacherecords").(string)
		hasChange = true
	}
	if d.HasChange("dnsanswerseclogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsanswerseclogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsanswerseclogging = d.Get("dnsanswerseclogging").(string)
		hasChange = true
	}
	if d.HasChange("dnserrorlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnserrorlogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnserrorlogging = d.Get("dnserrorlogging").(string)
		hasChange = true
	}
	if d.HasChange("dnsextendedlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsextendedlogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsextendedlogging = d.Get("dnsextendedlogging").(string)
		hasChange = true
	}
	if d.HasChange("dnsquerylogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsquerylogging has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dnsquerylogging = d.Get("dnsquerylogging").(string)
		hasChange = true
	}
	if d.HasChange("dropmultiqueryrequest") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropmultiqueryrequest has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Dropmultiqueryrequest = d.Get("dropmultiqueryrequest").(string)
		hasChange = true
	}
	if d.HasChange("recursiveresolution") {
		log.Printf("[DEBUG]  citrixadc-provider: Recursiveresolution has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Recursiveresolution = d.Get("recursiveresolution").(string)
		hasChange = true
	}
	if d.HasChange("insertecs") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertecs has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Insertecs = d.Get("insertecs").(string)
		hasChange = true
	}
	if d.HasChange("replaceecs") {
		log.Printf("[DEBUG]  citrixadc-provider: Replaceecs has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Replaceecs = d.Get("replaceecs").(string)
		hasChange = true
	}
	if d.HasChange("maxcacheableecsprefixlength") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcacheableecsprefixlength has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Maxcacheableecsprefixlength = d.Get("maxcacheableecsprefixlength").(int)
		hasChange = true
	}
	if d.HasChange("maxcacheableecsprefixlength6") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcacheableecsprefixlength6 has changed for dnsprofile %s, starting update", dnsprofileName)
		dnsprofile.Maxcacheableecsprefixlength6 = d.Get("maxcacheableecsprefixlength6").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnsprofile.Type(), dnsprofileName, &dnsprofile)
		if err != nil {
			return fmt.Errorf("Error updating dnsprofile %s", dnsprofileName)
		}
	}
	return readDnsprofileFunc(d, meta)
}

func deleteDnsprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsprofileName := d.Id()
	err := client.DeleteResource(service.Dnsprofile.Type(), dnsprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
