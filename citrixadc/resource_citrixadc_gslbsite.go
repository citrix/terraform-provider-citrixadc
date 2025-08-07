package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

/**
* Configuration for GSLB site resource.
 */
type Gslbsite struct {
	/**
	* Name for the GSLB site. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my gslbsite" or 'my gslbsite').
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site.
	 */
	Sitetype string `json:"sitetype,omitempty"`
	/**
	* IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).
	 */
	Siteipaddress string `json:"siteipaddress,omitempty"`
	/**
	* Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.
	 */
	Publicip string `json:"publicip,omitempty"`
	/**
	* Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.
		If you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.
	*/
	Metricexchange string `json:"metricexchange,omitempty"`
	/**
	* Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds.
	 */
	Nwmetricexchange string `json:"nwmetricexchange,omitempty"`
	/**
	* Exchange persistent session entries with other GSLB sites every five seconds.
	 */
	Sessionexchange string `json:"sessionexchange,omitempty"`
	/**
	* Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:
		* ALWAYS - Monitor the GSLB service at all times.
		* MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.
		MEPDOWN_SVCDOWN - Monitor the service in either of the following situations:
		* The exchange of metrics through MEP is disabled.
		* The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.
	*/
	Triggermonitor string `json:"triggermonitor,omitempty"`
	/**
	* Parent site of the GSLB site, in a parent-child topology.
	 */
	Parentsite string `json:"parentsite,omitempty"`
	/**
	* Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.
	 */
	Clip string `json:"clip,omitempty"`
	/**
	* IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.
	 */
	Publicclip string `json:"publicclip,omitempty"`
	/**
	* The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.
	 */
	Naptrreplacementsuffix string `json:"naptrreplacementsuffix,omitempty"`
	/**
	* The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.
	 */
	Backupparentlist []string `json:"backupparentlist,omitempty"`
	/**
	* Password to be used for mep communication between gslb site nodes.
	 */
	Sitepassword string `json:"sitepassword,omitempty"`
	/**
	* New name for the GSLB site.
	 */
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Status               string `json:"status,omitempty"`
	Persistencemepstatus string `json:"persistencemepstatus,omitempty"`
	Version              string `json:"version,omitempty"`
	Curbackupparentip    string `json:"curbackupparentip,omitempty"`
	Sitestate            string `json:"sitestate,omitempty"`
	Oldname              string `json:"oldname,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

func resourceCitrixAdcGslbsite() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbsiteFunc,
		Read:          readGslbsiteFunc,
		Update:        updateGslbsiteFunc,
		Delete:        deleteGslbsiteFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"naptrreplacementsuffix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nwmetricexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parentsite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicclip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionexchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"triggermonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupparentlist": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sitepassword": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbsiteName string
	if v, ok := d.GetOk("sitename"); ok {
		gslbsiteName = v.(string)
	} else {
		gslbsiteName = resource.PrefixedUniqueId("tf-gslbsite-")
		d.Set("sitename", gslbsiteName)
	}
	gslbsite := Gslbsite{
		Clip:                   d.Get("clip").(string),
		Metricexchange:         d.Get("metricexchange").(string),
		Naptrreplacementsuffix: d.Get("naptrreplacementsuffix").(string),
		Nwmetricexchange:       d.Get("nwmetricexchange").(string),
		Parentsite:             d.Get("parentsite").(string),
		Publicclip:             d.Get("publicclip").(string),
		Publicip:               d.Get("publicip").(string),
		Sessionexchange:        d.Get("sessionexchange").(string),
		Siteipaddress:          d.Get("siteipaddress").(string),
		Sitename:               d.Get("sitename").(string),
		Sitetype:               d.Get("sitetype").(string),
		Triggermonitor:         d.Get("triggermonitor").(string),
		Sitepassword:           d.Get("sitepassword").(string),
	}
	if listVal, ok := d.Get("backupparentlist").([]interface{}); ok {
		gslbsite.Backupparentlist = toStringList(listVal)
	}
	_, err := client.AddResource(service.Gslbsite.Type(), gslbsiteName, &gslbsite)
	if err != nil {
		return err
	}

	d.SetId(gslbsiteName)

	err = readGslbsiteFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbsite but we can't read it ?? %s", gslbsiteName)
		return nil
	}
	return nil
}

func readGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading gslbsite state %s", gslbsiteName)
	data, err := client.FindResource(service.Gslbsite.Type(), gslbsiteName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing gslbsite state %s", gslbsiteName)
		d.SetId("")
		return nil
	}
	d.Set("sitename", data["sitename"])
	d.Set("clip", data["clip"])
	d.Set("metricexchange", data["metricexchange"])
	d.Set("naptrreplacementsuffix", data["naptrreplacementsuffix"])
	d.Set("nwmetricexchange", data["nwmetricexchange"])
	d.Set("parentsite", data["parentsite"])
	d.Set("publicclip", data["publicclip"])
	d.Set("publicip", data["publicip"])
	d.Set("sessionexchange", data["sessionexchange"])
	d.Set("siteipaddress", data["siteipaddress"])
	d.Set("sitetype", data["sitetype"])
	d.Set("triggermonitor", data["triggermonitor"])
	d.Set("sitepassword", d.Get("sitepassword").(string))
	d.Set("status", data["status"])
	d.Set("persistencemepstatus", data["persistencemepstatus"])
	d.Set("version", data["version"])
	d.Set("curbackupparentip", data["curbackupparentip"])
	d.Set("sitestate", data["sitestate"])
	d.Set("oldname", data["oldname"])
	d.Set("nextgenapiresource", data["_nextgenapiresource"])
	if val, ok := data["backupparentlist"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("backupparentlist", toStringList(list))
		}
	} else {
		d.Set("backupparentlist", nil)
	}

	return nil

}

func updateGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Get("sitename").(string)

	gslbsite := Gslbsite{
		Sitename: gslbsiteName,
	}
	hasRename := false
	hasChange := false
	if d.HasChange("metricexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Metricexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Metricexchange = d.Get("metricexchange").(string)
		hasChange = true
	}
	if d.HasChange("naptrreplacementsuffix") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrreplacementsuffix has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Naptrreplacementsuffix = d.Get("naptrreplacementsuffix").(string)
		hasChange = true
	}
	if d.HasChange("nwmetricexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Nwmetricexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Nwmetricexchange = d.Get("nwmetricexchange").(string)
		hasChange = true
	}
	if d.HasChange("parentsite") {
		log.Printf("[DEBUG]  netscaler-provider: Parentsite has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Parentsite = d.Get("parentsite").(string)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicip has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("sessionexchange") {
		log.Printf("[DEBUG]  netscaler-provider: Sessionexchange has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Sessionexchange = d.Get("sessionexchange").(string)
		hasChange = true
	}
	if d.HasChange("siteipaddress") {
		log.Printf("[DEBUG]  netscaler-provider: Siteipaddress has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Siteipaddress = d.Get("siteipaddress").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  netscaler-provider: Sitename has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Sitename = gslbsiteName
		hasChange = true
	}
	if d.HasChange("triggermonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Triggermonitor has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Triggermonitor = d.Get("triggermonitor").(string)
		hasChange = true
	}
	if d.HasChange("backupparentlist") {
		log.Printf("[DEBUG]  netscaler-provider: Backupparentlist has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Backupparentlist = toStringList(d.Get("backupparentlist").([]interface{}))
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  netscaler-provider: Newname has changed for gslbsite %s, starting update", gslbsiteName)
		gslbsite.Newname = d.Get("newname").(string)
		hasRename = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Gslbsite.Type(), gslbsiteName, &gslbsite)
		if err != nil {
			return err
		}
	}

	if hasRename {
		err := client.ActOnResource(service.Gslbsite.Type(), &gslbsite, "rename")
		if err != nil {
			return err
		}
		d.SetId(gslbsite.Newname)
	}

	return readGslbsiteFunc(d, meta)
}

func deleteGslbsiteFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbsiteFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbsiteName := d.Id()
	err := client.DeleteResource(service.Gslbsite.Type(), gslbsiteName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
