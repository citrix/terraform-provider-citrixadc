package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/basic"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcServicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServicegroupFunc,
		Read:          readServicegroupFunc,
		Update:        updateServicegroupFunc,
		Delete:        deleteServicegroupFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autodisabledelay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"autodisablegraceful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customserverid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbsttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dupweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includemembers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memberport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monconnectionclose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nameserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitorindv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"riseapbrstatsmsgcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtspsessionidremap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"useproxyport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	var servicegroupName string
	if v, ok := d.GetOk("name"); ok {
		servicegroupName = v.(string)
	} else {
		servicegroupName = resource.PrefixedUniqueId("tf-servicegroup-")
		d.Set("name", servicegroupName)
	}
	servicegroup := basic.Servicegroup{
		Appflowlog:           d.Get("appflowlog").(string),
		Autodisabledelay:     d.Get("autodisabledelay").(int),
		Autodisablegraceful:  d.Get("autodisablegraceful").(string),
		Autoscale:            d.Get("autoscale").(string),
		Cacheable:            d.Get("cacheable").(string),
		Cachetype:            d.Get("cachetype").(string),
		Cip:                  d.Get("cip").(string),
		Cipheader:            d.Get("cipheader").(string),
		Cka:                  d.Get("cka").(string),
		Clttimeout:           d.Get("clttimeout").(int),
		Cmp:                  d.Get("cmp").(string),
		Comment:              d.Get("comment").(string),
		Customserverid:       d.Get("customserverid").(string),
		Dbsttl:               d.Get("dbsttl").(int),
		Delay:                d.Get("delay").(int),
		Downstateflush:       d.Get("downstateflush").(string),
		Dupweight:            d.Get("dupweight").(int),
		Graceful:             d.Get("graceful").(string),
		Hashid:               d.Get("hashid").(int),
		Healthmonitor:        d.Get("healthmonitor").(string),
		Httpprofilename:      d.Get("httpprofilename").(string),
		Includemembers:       d.Get("includemembers").(bool),
		Maxbandwidth:         d.Get("maxbandwidth").(int),
		Maxclient:            d.Get("maxclient").(int),
		Maxreq:               d.Get("maxreq").(int),
		Memberport:           d.Get("memberport").(int),
		Monconnectionclose:   d.Get("monconnectionclose").(string),
		Monitornamesvc:       d.Get("monitornamesvc").(string),
		Monthreshold:         d.Get("monthreshold").(int),
		Nameserver:           d.Get("nameserver").(string),
		Netprofile:           d.Get("netprofile").(string),
		Newname:              d.Get("newname").(string),
		Pathmonitor:          d.Get("pathmonitor").(string),
		Pathmonitorindv:      d.Get("pathmonitorindv").(string),
		Port:                 d.Get("port").(int),
		Riseapbrstatsmsgcode: d.Get("riseapbrstatsmsgcode").(int),
		Rtspsessionidremap:   d.Get("rtspsessionidremap").(string),
		Sc:                   d.Get("sc").(string),
		Serverid:             d.Get("serverid").(int),
		Servername:           d.Get("servername").(string),
		Servicegroupname:     d.Get("servicegroupname").(string),
		Servicetype:          d.Get("servicetype").(string),
		Sp:                   d.Get("sp").(string),
		State:                d.Get("state").(string),
		Svrtimeout:           d.Get("svrtimeout").(int),
		Tcpb:                 d.Get("tcpb").(string),
		Tcpprofilename:       d.Get("tcpprofilename").(string),
		Td:                   d.Get("td").(int),
		Useproxyport:         d.Get("useproxyport").(string),
		Usip:                 d.Get("usip").(string),
		Weight:               d.Get("weight").(int),
	}

	_, err := client.AddResource(netscaler.Servicegroup.Type(), servicegroupName, &servicegroup)
	if err != nil {
		return err
	}

	d.SetId(servicegroupName)

	err = readServicegroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this servicegroup but we can't read it ?? %s", servicegroupName)
		return nil
	}
	return nil
}

func readServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading servicegroup state %s", servicegroupName)
	data, err := client.FindResource(netscaler.Servicegroup.Type(), servicegroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("autodisabledelay", data["autodisabledelay"])
	d.Set("autodisablegraceful", data["autodisablegraceful"])
	d.Set("autoscale", data["autoscale"])
	d.Set("cacheable", data["cacheable"])
	d.Set("cachetype", data["cachetype"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("cka", data["cka"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("cmp", data["cmp"])
	d.Set("comment", data["comment"])
	d.Set("customserverid", data["customserverid"])
	d.Set("dbsttl", data["dbsttl"])
	d.Set("delay", data["delay"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("dupweight", data["dupweight"])
	d.Set("graceful", data["graceful"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("includemembers", data["includemembers"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxclient", data["maxclient"])
	d.Set("maxreq", data["maxreq"])
	d.Set("memberport", data["memberport"])
	d.Set("monconnectionclose", data["monconnectionclose"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("nameserver", data["nameserver"])
	d.Set("netprofile", data["netprofile"])
	d.Set("newname", data["newname"])
	d.Set("pathmonitor", data["pathmonitor"])
	d.Set("pathmonitorindv", data["pathmonitorindv"])
	d.Set("port", data["port"])
	d.Set("riseapbrstatsmsgcode", data["riseapbrstatsmsgcode"])
	d.Set("rtspsessionidremap", data["rtspsessionidremap"])
	d.Set("sc", data["sc"])
	d.Set("serverid", data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sp", data["sp"])
	d.Set("state", data["state"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("tcpb", data["tcpb"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("useproxyport", data["useproxyport"])
	d.Set("usip", data["usip"])
	d.Set("weight", data["weight"])

	return nil

}

func updateServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Get("name").(string)

	servicegroup := basic.Servicegroup{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("autodisabledelay") {
		log.Printf("[DEBUG]  citrixadc-provider: Autodisabledelay has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autodisabledelay = d.Get("autodisabledelay").(int)
		hasChange = true
	}
	if d.HasChange("autodisablegraceful") {
		log.Printf("[DEBUG]  citrixadc-provider: Autodisablegraceful has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autodisablegraceful = d.Get("autodisablegraceful").(string)
		hasChange = true
	}
	if d.HasChange("autoscale") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoscale has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autoscale = d.Get("autoscale").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG]  citrixadc-provider: Cacheable has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("cachetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Cachetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cachetype = d.Get("cachetype").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  citrixadc-provider: Cip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipheader has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("cka") {
		log.Printf("[DEBUG]  citrixadc-provider: Cka has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cka = d.Get("cka").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clttimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cmp") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cmp = d.Get("cmp").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("customserverid") {
		log.Printf("[DEBUG]  citrixadc-provider: Customserverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Customserverid = d.Get("customserverid").(string)
		hasChange = true
	}
	if d.HasChange("dbsttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Dbsttl has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Dbsttl = d.Get("dbsttl").(int)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  citrixadc-provider: Delay has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Delay = d.Get("delay").(int)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("dupweight") {
		log.Printf("[DEBUG]  citrixadc-provider: Dupweight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Dupweight = d.Get("dupweight").(int)
		hasChange = true
	}
	if d.HasChange("graceful") {
		log.Printf("[DEBUG]  citrixadc-provider: Graceful has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Graceful = d.Get("graceful").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  citrixadc-provider: Hashid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Healthmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("includemembers") {
		log.Printf("[DEBUG]  citrixadc-provider: Includemembers has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Includemembers = d.Get("includemembers").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxbandwidth has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxclient has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxreq has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxreq = d.Get("maxreq").(int)
		hasChange = true
	}
	if d.HasChange("memberport") {
		log.Printf("[DEBUG]  citrixadc-provider: Memberport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Memberport = d.Get("memberport").(int)
		hasChange = true
	}
	if d.HasChange("monconnectionclose") {
		log.Printf("[DEBUG]  citrixadc-provider: Monconnectionclose has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monconnectionclose = d.Get("monconnectionclose").(string)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitornamesvc has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Monthreshold has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("nameserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Nameserver has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Nameserver = d.Get("nameserver").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Pathmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitor = d.Get("pathmonitor").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitorindv") {
		log.Printf("[DEBUG]  citrixadc-provider: Pathmonitorindv has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitorindv = d.Get("pathmonitorindv").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("riseapbrstatsmsgcode") {
		log.Printf("[DEBUG]  citrixadc-provider: Riseapbrstatsmsgcode has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Riseapbrstatsmsgcode = d.Get("riseapbrstatsmsgcode").(int)
		hasChange = true
	}
	if d.HasChange("rtspsessionidremap") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtspsessionidremap has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Rtspsessionidremap = d.Get("rtspsessionidremap").(string)
		hasChange = true
	}
	if d.HasChange("sc") {
		log.Printf("[DEBUG]  citrixadc-provider: Sc has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Sc = d.Get("sc").(string)
		hasChange = true
	}
	if d.HasChange("serverid") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Serverid = d.Get("serverid").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicegroupname has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sp") {
		log.Printf("[DEBUG]  citrixadc-provider: Sp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Sp = d.Get("sp").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("tcpb") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpb has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpb = d.Get("tcpb").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  citrixadc-provider: Useproxyport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}
	if d.HasChange("usip") {
		log.Printf("[DEBUG]  citrixadc-provider: Usip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Usip = d.Get("usip").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  citrixadc-provider: Weight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Servicegroup.Type(), servicegroupName, &servicegroup)
		if err != nil {
			return fmt.Errorf("Error updating servicegroup %s", servicegroupName)
		}
	}
	return readServicegroupFunc(d, meta)
}

func deleteServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	err := client.DeleteResource(netscaler.Servicegroup.Type(), servicegroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
