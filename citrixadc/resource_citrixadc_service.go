package citrixadc

import (
	"fmt"
	"log"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCitrixAdcService() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServiceFunc,
		Read:          readServiceFunc,
		Update:        updateServiceFunc,
		Delete:        deleteServiceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"accessdown": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"all": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"appflowlog": {
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
				ForceNew: true,
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
			"cleartextport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"contentinspectionprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customserverid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
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
			"internal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"netprofile": &schema.Schema{
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
				ForceNew: true,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"riseapbrstatsmsgcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
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

			"lbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"lbmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// SSL service parameters
			"snienable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Wait for disabled state parameters
			"wait_until_disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"disabled_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default:  "2m",
			},
			"disabled_poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default:  "2s",
			},
			"disabled_poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default:  "5s",
			},
		},
	}
}

func createServiceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In createServiceFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	var serviceName string
	if v, ok := d.GetOk("name"); ok {
		serviceName = v.(string)
	} else {
		serviceName = resource.PrefixedUniqueId("tf-service-")
		d.Set("name", serviceName)
	}
	lbmonitor, mok := d.GetOk("lbmonitor")
	if mok {
		exists := client.ResourceExists(service.Lbmonitor.Type(), lbmonitor.(string))
		if !exists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified lb monitor does not exist on netscaler!")
		}
	}

	lbvserver, lok := d.GetOk("lbvserver")
	if lok {
		exists := client.ResourceExists(service.Lbvserver.Type(), lbvserver.(string))
		if !exists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified lb vserver does not exist on netscaler!")
		}
	}
	svc := basic.Service{
		Name:                         serviceName,
		Accessdown:                   d.Get("accessdown").(string),
		All:                          d.Get("all").(bool),
		Appflowlog:                   d.Get("appflowlog").(string),
		Cacheable:                    d.Get("cacheable").(string),
		Cachetype:                    d.Get("cachetype").(string),
		Cip:                          d.Get("cip").(string),
		Cipheader:                    d.Get("cipheader").(string),
		Cka:                          d.Get("cka").(string),
		Cleartextport:                d.Get("cleartextport").(int),
		Clttimeout:                   d.Get("clttimeout").(int),
		Cmp:                          d.Get("cmp").(string),
		Comment:                      d.Get("comment").(string),
		Contentinspectionprofilename: d.Get("contentinspectionprofilename").(string),
		Customserverid:               d.Get("customserverid").(string),
		Dnsprofilename:               d.Get("dnsprofilename").(string),
		Downstateflush:               d.Get("downstateflush").(string),
		Hashid:                       d.Get("hashid").(int),
		Healthmonitor:                d.Get("healthmonitor").(string),
		Httpprofilename:              d.Get("httpprofilename").(string),
		Internal:                     d.Get("internal").(bool),
		Ip:                           d.Get("ip").(string),
		Ipaddress:                    d.Get("ipaddress").(string),
		Maxbandwidth:                 d.Get("maxbandwidth").(int),
		Maxclient:                    d.Get("maxclient").(int),
		Maxreq:                       d.Get("maxreq").(int),
		Monconnectionclose:           d.Get("monconnectionclose").(string),
		Monitornamesvc:               d.Get("monitornamesvc").(string),
		Monthreshold:                 d.Get("monthreshold").(int),
		Netprofile:                   d.Get("netprofile").(string),
		Pathmonitor:                  d.Get("pathmonitor").(string),
		Pathmonitorindv:              d.Get("pathmonitorindv").(string),
		Port:                         d.Get("port").(int),
		Processlocal:                 d.Get("processlocal").(string),
		Rtspsessionidremap:           d.Get("rtspsessionidremap").(string),
		Sc:                           d.Get("sc").(string),
		Serverid:                     d.Get("serverid").(int),
		Servername:                   d.Get("servername").(string),
		Servicetype:                  d.Get("servicetype").(string),
		Sp:                           d.Get("sp").(string),
		State:                        d.Get("state").(string),
		Svrtimeout:                   d.Get("svrtimeout").(int),
		Tcpb:                         d.Get("tcpb").(string),
		Tcpprofilename:               d.Get("tcpprofilename").(string),
		Td:                           d.Get("td").(int),
		Useproxyport:                 d.Get("useproxyport").(string),
		Usip:                         d.Get("usip").(string),
		Weight:                       d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Service.Type(), serviceName, &svc)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: could not add resource %s of type %s", service.Service.Type(), serviceName)
		return err
	}
	if lok { //lbvserver is specified
		lbvserverName := d.Get("lbvserver").(string)
		binding := lb.Lbvserverservicebinding{
			Name:        lbvserverName,
			Servicename: serviceName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding service %s to lbvserver %s", serviceName, lbvserverName)
		err = client.BindResource(service.Lbvserver.Type(), lbvserverName, service.Service.Type(), serviceName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind service %s to lbvserver %s", serviceName, lbvserverName)
			err2 := client.DeleteResource(service.Service.Type(), serviceName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete service %s after bind to lb vserver failed", serviceName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete service %s after bind to lbvserver failed", serviceName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind  service %s to lbvserver %s", serviceName, lbvserverName)
		}
	}
	if mok { //lbmonitor is specified
		lbmonitorName := d.Get("lbmonitor").(string)
		binding := lb.Lbmonitorservicebinding{
			Monitorname: lbmonitorName,
			Servicename: serviceName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding service %s to lbmonitor %s", serviceName, lbmonitorName)
		err = client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Service.Type(), serviceName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind service %s to lbmonitor %s", serviceName, lbmonitorName)
			err2 := client.DeleteResource(service.Service.Type(), serviceName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete service %s after bind to lbmonitor failed", serviceName)
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to delete service %s after bind to lbmonitor failed", serviceName)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind  service %s to lbmonitor %s", serviceName, lbmonitorName)
		}
	}

	if hasSslserviceProperties(d) {
		err := syncSslservice(d, client)
		if err != nil {
			return err
		}
	}

	d.SetId(serviceName)
	err = readServiceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this service but we can't read it ?? %s", serviceName)
		return nil
	}
	return nil
}

func readServiceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readServiceFunc")

	client := meta.(*NetScalerNitroClient).client
	serviceName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading service state %s", serviceName)
	data, err := client.FindResource(service.Service.Type(), serviceName)
	log.Printf("Reading service state %v", data)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing service state %s", serviceName)
		d.SetId("")
		return nil
	}
	//read bound vserver.
	var vserverBindings []map[string]interface{}
	if _, ok := d.GetOk("lbvserver"); ok {
		vserverBindings, err = client.FindResourceArray(service.Svcbindings.Type(), serviceName)
		if err != nil {
			log.Printf("[WARN] netscaler-provider: Clearing service state %s", serviceName)
			d.SetId("")
			return nil
		}
	}
	//read bound lb monitor.
	boundMonitors, err := client.FindAllBoundResources(service.Service.Type(), serviceName, service.Lbmonitor.Type())
	if err != nil {
		//This is actually OK in most cases
		log.Printf("[WARN] netscaler-provider: Clearing servicestate %s", serviceName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("accessdown", data["accessdown"])
	d.Set("all", data["all"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("cacheable", data["cacheable"])
	d.Set("cachetype", data["cachetype"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("cka", data["cka"])
	d.Set("cleartextport", data["cleartextport"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("cmp", data["cmp"])
	d.Set("comment", data["comment"])
	d.Set("contentinspectionprofilename", data["contentinspectionprofilename"])
	d.Set("customserverid", data["customserverid"])
	d.Set("dnsprofilename", data["dnsprofilename"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("internal", data["internal"])
	/*
		if data["ip"] != "" {
			d.Set("ip", data["ip"])
		}*/
	d.Set("ipaddress", data["ipaddress"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxclient", data["maxclient"])
	d.Set("maxreq", data["maxreq"])
	d.Set("monconnectionclose", data["monconnectionclose"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("pathmonitor", data["pathmonitor"])
	d.Set("pathmonitorindv", data["pathmonitorindv"])
	d.Set("port", data["port"])
	d.Set("processlocal", data["processlocal"])
	d.Set("riseapbrstatsmsgcode", data["riseapbrstatsmsgcode"])
	d.Set("rtspsessionidremap", data["rtspsessionidremap"])
	d.Set("sc", data["sc"])
	d.Set("serverid", data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sp", data["sp"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("tcpb", data["tcpb"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("useproxyport", data["useproxyport"])
	d.Set("usip", data["usip"])
	d.Set("weight", data["weight"])

	// Set state according to svrstate
	if data["svrstate"] == "OUT OF SERVICE" {
		d.Set("state", "DISABLED")
	} else {
		d.Set("state", "ENABLED")
	}

	var boundVserver string
	if _, ok := d.GetOk("lbvserver"); ok {
		for _, vserver := range vserverBindings {
			vs, ok := vserver["vservername"]
			if ok {
				boundVserver = vs.(string)
				break
			}
		}
	}
	d.Set("lbvserver", boundVserver)

	var boundMonitor string
	for _, monitor := range boundMonitors {
		mon, ok := monitor["monitor_name"]
		if ok {
			boundMonitor = mon.(string)
			log.Printf("[INFO] netscaler-provider:  Found %s  lbmonitor bound to %s", boundMonitor, serviceName)
			break
		}
	}
	d.Set("lbmonitor", boundMonitor)

	if hasSslserviceProperties(d) {
		err := readSslservice(d, client)
		if err != nil {
			return err
		}
	}

	return nil

}

func updateServiceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In updateServiceFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	serviceName := d.Get("name").(string)

	stateChange := false
	hasChange := false
	lbvserverChanged := false
	lbmonitorChanged := false
	svc := basic.Service{
		Name: d.Get("name").(string),
	}
	if d.HasChange("accessdown") {
		log.Printf("[DEBUG] netscaler-provider:  Accessdown has changed for service %s, starting update", serviceName)
		svc.Accessdown = d.Get("accessdown").(string)
		hasChange = true
	}
	if d.HasChange("all") {
		log.Printf("[DEBUG] netscaler-provider:  All has changed for service %s, starting update", serviceName)
		svc.All = d.Get("all").(bool)
		hasChange = true
	}
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG] netscaler-provider:  Appflowlog has changed for service %s, starting update", serviceName)
		svc.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG] netscaler-provider:  Cacheable has changed for service %s, starting update", serviceName)
		svc.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("cachetype") {
		log.Printf("[DEBUG] netscaler-provider:  Cachetype has changed for service %s, starting update", serviceName)
		svc.Cachetype = d.Get("cachetype").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG] netscaler-provider:  Cip has changed for service %s, starting update", serviceName)
		svc.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG] netscaler-provider:  Cipheader has changed for service %s, starting update", serviceName)
		svc.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("cka") {
		log.Printf("[DEBUG] netscaler-provider:  Cka has changed for service %s, starting update", serviceName)
		svc.Cka = d.Get("cka").(string)
		hasChange = true
	}
	if d.HasChange("cleartextport") {
		log.Printf("[DEBUG] netscaler-provider:  Cleartextport has changed for service %s, starting update", serviceName)
		svc.Cleartextport = d.Get("cleartextport").(int)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Clttimeout has changed for service %s, starting update", serviceName)
		svc.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cmp") {
		log.Printf("[DEBUG] netscaler-provider:  Cmp has changed for service %s, starting update", serviceName)
		svc.Cmp = d.Get("cmp").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG] netscaler-provider:  Comment has changed for service %s, starting update", serviceName)
		svc.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("contentinspectionprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Contentinspectionprofilename has changed for service %s, starting update", serviceName)
		svc.Contentinspectionprofilename = d.Get("contentinspectionprofilename").(string)
		hasChange = true
	}
	if d.HasChange("customserverid") {
		log.Printf("[DEBUG] netscaler-provider:  Customserverid has changed for service %s, starting update", serviceName)
		svc.Customserverid = d.Get("customserverid").(string)
		hasChange = true
	}
	if d.HasChange("dnsprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Dnsprofilename has changed for service %s, starting update", serviceName)
		svc.Dnsprofilename = d.Get("dnsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG] netscaler-provider:  Downstateflush has changed for service %s, starting update", serviceName)
		svc.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG] netscaler-provider:  Hashid has changed for service %s, starting update", serviceName)
		svc.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG] netscaler-provider:  Healthmonitor has changed for service %s, starting update", serviceName)
		svc.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Httpprofilename has changed for service %s, starting update", serviceName)
		svc.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("internal") {
		log.Printf("[DEBUG] netscaler-provider:  Internal has changed for service %s, starting update", serviceName)
		svc.Internal = d.Get("internal").(bool)
		hasChange = true
	}
	if d.HasChange("ip") {
		log.Printf("[DEBUG] netscaler-provider:  Ip has changed for service %s, starting update", serviceName)
		svc.Ipaddress = d.Get("ip").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG] netscaler-provider:  Ipaddress has changed for service %s, starting update", serviceName)
		svc.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG] netscaler-provider:  Maxbandwidth has changed for service %s, starting update", serviceName)
		svc.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG] netscaler-provider:  Maxclient has changed for service %s, starting update", serviceName)
		svc.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG] netscaler-provider:  Maxreq has changed for service %s, starting update", serviceName)
		svc.Maxreq = d.Get("maxreq").(int)
		hasChange = true
	}
	if d.HasChange("monconnectionclose") {
		log.Printf("[DEBUG]  netscaler-provider: Monconnectionclose has changed for service %s, starting update", serviceName)
		svc.Monconnectionclose = d.Get("monconnectionclose").(string)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG] netscaler-provider:  Monitornamesvc has changed for service %s, starting update", serviceName)
		svc.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG] netscaler-provider:  Monthreshold has changed for service %s, starting update", serviceName)
		svc.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG] netscaler-provider:  Name has changed for service %s, starting update", serviceName)
		svc.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] netscaler-provider:  Netprofile has changed for service %s, starting update", serviceName)
		svc.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitor") {
		log.Printf("[DEBUG] netscaler-provider:  Pathmonitor has changed for service %s, starting update", serviceName)
		svc.Pathmonitor = d.Get("pathmonitor").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitorindv") {
		log.Printf("[DEBUG] netscaler-provider:  Pathmonitorindv has changed for service %s, starting update", serviceName)
		svc.Pathmonitorindv = d.Get("pathmonitorindv").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG] netscaler-provider:  Port has changed for service %s, starting update", serviceName)
		svc.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("processlocal") {
		log.Printf("[DEBUG]  netscaler-provider: Processlocal has changed for service %s, starting update", serviceName)
		svc.Processlocal = d.Get("processlocal").(string)
		hasChange = true
	}
	if d.HasChange("rtspsessionidremap") {
		log.Printf("[DEBUG] netscaler-provider:  Rtspsessionidremap has changed for service %s, starting update", serviceName)
		svc.Rtspsessionidremap = d.Get("rtspsessionidremap").(string)
		hasChange = true
	}
	if d.HasChange("sc") {
		log.Printf("[DEBUG] netscaler-provider:  Sc has changed for service %s, starting update", serviceName)
		svc.Sc = d.Get("sc").(string)
		hasChange = true
	}
	if d.HasChange("serverid") {
		log.Printf("[DEBUG] netscaler-provider:  Serverid has changed for service %s, starting update", serviceName)
		svc.Serverid = d.Get("serverid").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG] netscaler-provider:  Servername has changed for service %s, starting update", serviceName)
		svc.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG] netscaler-provider:  Servicetype has changed for service %s, starting update", serviceName)
		svc.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sp") {
		log.Printf("[DEBUG] netscaler-provider:  Sp has changed for service %s, starting update", serviceName)
		svc.Sp = d.Get("sp").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] netscaler-provider:  State has changed for service %s, starting update", serviceName)
		stateChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG] netscaler-provider:  Svrtimeout has changed for service %s, starting update", serviceName)
		svc.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("tcpb") {
		log.Printf("[DEBUG] netscaler-provider:  Tcpb has changed for service %s, starting update", serviceName)
		svc.Tcpb = d.Get("tcpb").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG] netscaler-provider:  Tcpprofilename has changed for service %s, starting update", serviceName)
		svc.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG] netscaler-provider:  Td has changed for service %s, starting update", serviceName)
		svc.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG] netscaler-provider:  Useproxyport has changed for service %s, starting update", serviceName)
		svc.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}
	if d.HasChange("usip") {
		log.Printf("[DEBUG] netscaler-provider:  Usip has changed for service %s, starting update", serviceName)
		svc.Usip = d.Get("usip").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG] netscaler-provider:  Weight has changed for service %s, starting update", serviceName)
		svc.Weight = d.Get("weight").(int)
		hasChange = true
	}
	if d.HasChange("lbmonitor") {
		log.Printf("[DEBUG] netscaler-provider:  lb monitor has changed for service %s, starting update", serviceName)
		lbmonitorChanged = true
	}
	if d.HasChange("lbvserver") {
		log.Printf("[DEBUG] netscaler-provider:  lb vserver has changed for service %s, starting update", serviceName)
		lbvserverChanged = true
	}

	lbmonitor := d.Get("lbmonitor")
	lbmonitorName := lbmonitor.(string)

	lbvserverName := d.Get("lbvserver").(string)
	if lbmonitorChanged {
		//Binding has to be updated
		//First we unbind from lb monitor
		oldLbmonitor, _ := d.GetChange("lbmonitor")
		oldLbmonitorName := oldLbmonitor.(string)

		// Default monitors cannot be unbound
		// Instead they are replaced when the new monitor is bound
		oldMonitorIsDefault := oldLbmonitorName == "ping-default" || oldLbmonitorName == "tcp-default"

		if oldLbmonitorName != "" && !oldMonitorIsDefault {
			err := client.UnbindResource(service.Lbmonitor.Type(), oldLbmonitorName, service.Service.Type(), serviceName, "servicename")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding lbmonitor from service %s", oldLbmonitorName)
			}
			log.Printf("[DEBUG] netscaler-provider: lbmonitor has been unbound from service for lb monitor %s ", oldLbmonitorName)
		}
	}
	if lbvserverChanged {
		//Binding has to be updated
		//First we unbind from lb vserver
		oldLbvserver, _ := d.GetChange("lbvserver")
		oldLbvserverName := oldLbvserver.(string)
		if oldLbvserverName != "" {
			err := client.UnbindResource(service.Lbvserver.Type(), oldLbvserverName, service.Service.Type(), serviceName, "servicename")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding lbvserver from service %s", oldLbvserverName)
			}
			log.Printf("[DEBUG] netscaler-provider: lbvserver has been unbound from service for lb vserver %s ", oldLbvserverName)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Service.Type(), serviceName, &svc)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error updating service %s", serviceName)
		}
		log.Printf("[DEBUG] netscaler-provider: service has been updated  service %s ", serviceName)
	}

	// Default monitors cannot be explicitely bound
	// Instead they are bound upon the unbind of the last non default monitor from the service
	newMonitorIsDefault := lbmonitorName == "ping-default" || lbmonitorName == "tcp-default"

	if lbmonitorChanged && lbmonitorName != "" && !newMonitorIsDefault {
		//Binding has to be updated
		//rebind
		binding := lb.Lbmonitorservicebinding{
			Monitorname: lbmonitorName,
			Servicename: serviceName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding monitor %s to service %s", lbmonitorName, serviceName)
		err := client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Service.Type(), serviceName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind  lbmonitor %s to service %s", lbmonitorName, serviceName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind lb monitor %s to service %s", lbmonitorName, serviceName)
		}
		log.Printf("[DEBUG] netscaler-provider: new lbmonitor has been bound to service  lbmonitor %s service %s", lbmonitorName, serviceName)
	}
	if lbvserverChanged && lbvserverName != "" {
		//Binding has to be updated
		//rebind
		binding := lb.Lbvserverservicebinding{
			Name:        lbvserverName,
			Servicename: serviceName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding vserver %s to service %s", lbvserverName, serviceName)
		err := client.BindResource(service.Lbvserver.Type(), lbvserverName, service.Service.Type(), serviceName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind  lbvserver %s to service %s", lbvserverName, serviceName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind lb vserver %s to service %s", lbvserverName, serviceName)
		}
		log.Printf("[DEBUG] netscaler-provider: new lbvserver has been bound to service  lbvserver %s service %s", lbvserverName, serviceName)
	}

	if hasSslserviceProperties(d) {
		err := syncSslservice(d, client)
		if err != nil {
			return err
		}
	}

	if stateChange {
		err := doServiceStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling service %s", serviceName)
		}
	}

	return readServiceFunc(d, meta)
}

func deleteServiceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In deleteServiceFunc")
	client := meta.(*NetScalerNitroClient).client

	meta.(*NetScalerNitroClient).lock.Lock()
	defer meta.(*NetScalerNitroClient).lock.Unlock()

	serviceName := d.Id()
	err := client.DeleteResource(service.Service.Type(), serviceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doServiceStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doServiceStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	svc := basic.Service{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Service.Type(), svc, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		svc.Delay = d.Get("delay").(int)
		svc.Graceful = d.Get("graceful").(string)
		err := client.ActOnResource(service.Service.Type(), svc, "disable")
		if err != nil {
			return err
		}
		// Wait for state change
		if d.Get("wait_until_disabled").(bool) {
			serviceWaitDisableState(d, client)
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}

func hasSslserviceProperties(d *schema.ResourceData) bool {
	hasProperties := false
	if _, ok := d.GetOk("snienable"); ok {
		hasProperties = true
	}
	if _, ok := d.GetOk("commonname"); ok {
		hasProperties = true
	}
	return hasProperties
}

func readSslservice(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In readSslservice")

	// Only go ahead and read if sslservice parameters are defined
	if !hasSslserviceProperties(d) {
		return nil
	}

	// Fallthrough
	sslserviceName := d.Get("name").(string)
	findParams := service.FindParams{
		ResourceType: "sslservice",
		ResourceName: sslserviceName,
	}
	arr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	if len(arr) > 1 {
		return fmt.Errorf("Too many sslservice results \"%v\"", arr)
	}
	d.Set("snienable", arr[0]["snienable"])
	d.Set("commonname", arr[0]["commonname"])
	return nil
}

func syncSslservice(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In syncSslservice")
	if !hasSslserviceProperties(d) {
		return nil
	}

	// Faltrhough
	sslserviceName := d.Get("name").(string)
	sslservice := ssl.Sslservice{
		Servicename: sslserviceName,
		Snienable:   d.Get("snienable").(string),
		Commonname:  d.Get("commonname").(string),
	}
	err := client.UpdateUnnamedResource("sslservice", &sslservice)
	if err != nil {
		return err
	}

	return nil
}

func serviceWaitDisableState(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG] citrixadc-provider: In serviceWaitDisableState")

	var err error
	var timeout time.Duration
	if timeout, err = time.ParseDuration(d.Get("disabled_timeout").(string)); err != nil {
		return err
	}

	var poll_interval time.Duration
	if poll_interval, err = time.ParseDuration(d.Get("disabled_poll_interval").(string)); err != nil {
		return err
	}

	var poll_delay time.Duration
	if poll_delay, err = time.ParseDuration(d.Get("disabled_poll_delay").(string)); err != nil {
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"enabled"},
		Target:       []string{"disabled"},
		Refresh:      serviceStatePoll(d, client),
		Timeout:      timeout,
		PollInterval: poll_interval,
		Delay:        poll_delay,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	return nil
}

func serviceStatePoll(d *schema.ResourceData, client *service.NitroClient) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] citrixadc-provider: In serviceStatePoll")
		serviceName := d.Id()
		data, err := client.FindResource(service.Service.Type(), serviceName)
		if err != nil {
			return nil, "", err
		}
		if data["svrstate"] == "OUT OF SERVICE" {
			return "disabled", "disabled", nil
		} else {
			return "enabled", "enabled", nil
		}
	}
}
