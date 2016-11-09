package netscaler

import (
        "github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)


func resourceNetScalerService() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
                Create:        createServiceFunc,
                Read:          readServiceFunc,
                Update:        updateServiceFunc,
                Delete:        deleteServiceFunc,
		Schema: map[string]*schema.Schema{
                       "accessdown": &schema.Schema{
                               Type:     schema.TypeString,
			       Optional: true,
                               Computed: true,
			},
                        "all": &schema.Schema{
                               Type:     schema.TypeBool,
			       Optional: true,
                               Computed: true,
			},
                        "appflowlog": &schema.Schema{
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
                        "cleartextport": &schema.Schema{
                               Type:     schema.TypeInt,
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

func createServiceFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        var serviceName string
	if v, ok := d.GetOk("name"); ok {
             serviceName = v.(string)
	} else {
             serviceName= resource.PrefixedUniqueId("tf-service-")
             d.Set("name", serviceName)
	}
        service := basic.Service{
                Name:            serviceName,
                Accessdown:           d.Get("accessdown").(string),
                All:           d.Get("all").(bool),
                Appflowlog:           d.Get("appflowlog").(string),
                Cacheable:           d.Get("cacheable").(string),
                Cachetype:           d.Get("cachetype").(string),
                Cip:           d.Get("cip").(string),
                Cipheader:           d.Get("cipheader").(string),
                Cka:           d.Get("cka").(string),
                Cleartextport:           d.Get("cleartextport").(int),
                Clttimeout:           d.Get("clttimeout").(int),
                Cmp:           d.Get("cmp").(string),
                Comment:           d.Get("comment").(string),
                Customserverid:           d.Get("customserverid").(string),
                Delay:           d.Get("delay").(int),
                Downstateflush:           d.Get("downstateflush").(string),
                Graceful:           d.Get("graceful").(string),
                Hashid:           d.Get("hashid").(int),
                Healthmonitor:           d.Get("healthmonitor").(string),
                Httpprofilename:           d.Get("httpprofilename").(string),
                Internal:           d.Get("internal").(bool),
                Ip:           d.Get("ip").(string),
                Ipaddress:           d.Get("ipaddress").(string),
                Maxbandwidth:           d.Get("maxbandwidth").(int),
                Maxclient:           d.Get("maxclient").(int),
                Maxreq:           d.Get("maxreq").(int),
                Monitornamesvc:           d.Get("monitornamesvc").(string),
                Monthreshold:           d.Get("monthreshold").(int),
                Netprofile:           d.Get("netprofile").(string),
                Newname:           d.Get("newname").(string),
                Pathmonitor:           d.Get("pathmonitor").(string),
                Pathmonitorindv:           d.Get("pathmonitorindv").(string),
                Port:           d.Get("port").(int),
                Rtspsessionidremap:           d.Get("rtspsessionidremap").(string),
                Sc:           d.Get("sc").(string),
                Serverid:           d.Get("serverid").(int),
                Servername:           d.Get("servername").(string),
                Servicetype:           d.Get("servicetype").(string),
                Sp:           d.Get("sp").(string),
                State:           d.Get("state").(string),
                Svrtimeout:           d.Get("svrtimeout").(int),
                Tcpb:           d.Get("tcpb").(string),
                Tcpprofilename:           d.Get("tcpprofilename").(string),
                Td:           d.Get("td").(int),
                Useproxyport:           d.Get("useproxyport").(string),
                Usip:           d.Get("usip").(string),
                Weight:           d.Get("weight").(int),
                
	}

        _, err := client.AddResource(netscaler.Service.Type(), serviceName, &service)
	if err != nil {
		return err
	}

        d.SetId(serviceName)
        err = readServiceFunc(d, meta)
	if err != nil {
                log.Printf("?? we just created this service but we can't read it ?? %s", serviceName)
		return nil
	}
	return nil
}

func readServiceFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        serviceName:= d.Id()
        log.Printf("Reading service state %s", serviceName)
        data, err := client.FindResource(netscaler.Service.Type(), serviceName)
	if err != nil {
        log.Printf("Clearing service state %s", serviceName)
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
        d.Set("customserverid", data["customserverid"])
        d.Set("delay", data["delay"])
        d.Set("downstateflush", data["downstateflush"])
        d.Set("graceful", data["graceful"])
        d.Set("hashid", data["hashid"])
        d.Set("healthmonitor", data["healthmonitor"])
        d.Set("httpprofilename", data["httpprofilename"])
        d.Set("internal", data["internal"])
        d.Set("ip", data["ip"])
        d.Set("ipaddress", data["ipaddress"])
        d.Set("maxbandwidth", data["maxbandwidth"])
        d.Set("maxclient", data["maxclient"])
        d.Set("maxreq", data["maxreq"])
        d.Set("monitornamesvc", data["monitornamesvc"])
        d.Set("monthreshold", data["monthreshold"])
        d.Set("name", data["name"])
        d.Set("netprofile", data["netprofile"])
        d.Set("newname", data["newname"])
        d.Set("pathmonitor", data["pathmonitor"])
        d.Set("pathmonitorindv", data["pathmonitorindv"])
        d.Set("port", data["port"])
        d.Set("rtspsessionidremap", data["rtspsessionidremap"])
        d.Set("sc", data["sc"])
        d.Set("serverid", data["serverid"])
        d.Set("servername", data["servername"])
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

func updateServiceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] In update func")
	client := meta.(*NetScalerNitroClient).client
        serviceName := d.Get("name").(string)

        service := basic.Service{
		Name: d.Get("name").(string),
	}
        if d.HasChange("accessdown") {
                log.Printf("[DEBUG] Accessdown has changed for service %s, starting update", serviceName)
                service.Accessdown = d.Get("accessdown").(string)
	}
        if d.HasChange("all") {
                log.Printf("[DEBUG] All has changed for service %s, starting update", serviceName)
                service.All = d.Get("all").(bool)
	}
        if d.HasChange("appflowlog") {
                log.Printf("[DEBUG] Appflowlog has changed for service %s, starting update", serviceName)
                service.Appflowlog = d.Get("appflowlog").(string)
	}
        if d.HasChange("cacheable") {
                log.Printf("[DEBUG] Cacheable has changed for service %s, starting update", serviceName)
                service.Cacheable = d.Get("cacheable").(string)
	}
        if d.HasChange("cachetype") {
                log.Printf("[DEBUG] Cachetype has changed for service %s, starting update", serviceName)
                service.Cachetype = d.Get("cachetype").(string)
	}
        if d.HasChange("cip") {
                log.Printf("[DEBUG] Cip has changed for service %s, starting update", serviceName)
                service.Cip = d.Get("cip").(string)
	}
        if d.HasChange("cipheader") {
                log.Printf("[DEBUG] Cipheader has changed for service %s, starting update", serviceName)
                service.Cipheader = d.Get("cipheader").(string)
	}
        if d.HasChange("cka") {
                log.Printf("[DEBUG] Cka has changed for service %s, starting update", serviceName)
                service.Cka = d.Get("cka").(string)
	}
        if d.HasChange("cleartextport") {
                log.Printf("[DEBUG] Cleartextport has changed for service %s, starting update", serviceName)
                service.Cleartextport = d.Get("cleartextport").(int)
	}
        if d.HasChange("clttimeout") {
                log.Printf("[DEBUG] Clttimeout has changed for service %s, starting update", serviceName)
                service.Clttimeout = d.Get("clttimeout").(int)
	}
        if d.HasChange("cmp") {
                log.Printf("[DEBUG] Cmp has changed for service %s, starting update", serviceName)
                service.Cmp = d.Get("cmp").(string)
	}
        if d.HasChange("comment") {
                log.Printf("[DEBUG] Comment has changed for service %s, starting update", serviceName)
                service.Comment = d.Get("comment").(string)
	}
        if d.HasChange("customserverid") {
                log.Printf("[DEBUG] Customserverid has changed for service %s, starting update", serviceName)
                service.Customserverid = d.Get("customserverid").(string)
	}
        if d.HasChange("delay") {
                log.Printf("[DEBUG] Delay has changed for service %s, starting update", serviceName)
                service.Delay = d.Get("delay").(int)
	}
        if d.HasChange("downstateflush") {
                log.Printf("[DEBUG] Downstateflush has changed for service %s, starting update", serviceName)
                service.Downstateflush = d.Get("downstateflush").(string)
	}
        if d.HasChange("graceful") {
                log.Printf("[DEBUG] Graceful has changed for service %s, starting update", serviceName)
                service.Graceful = d.Get("graceful").(string)
	}
        if d.HasChange("hashid") {
                log.Printf("[DEBUG] Hashid has changed for service %s, starting update", serviceName)
                service.Hashid = d.Get("hashid").(int)
	}
        if d.HasChange("healthmonitor") {
                log.Printf("[DEBUG] Healthmonitor has changed for service %s, starting update", serviceName)
                service.Healthmonitor = d.Get("healthmonitor").(string)
	}
        if d.HasChange("httpprofilename") {
                log.Printf("[DEBUG] Httpprofilename has changed for service %s, starting update", serviceName)
                service.Httpprofilename = d.Get("httpprofilename").(string)
	}
        if d.HasChange("internal") {
                log.Printf("[DEBUG] Internal has changed for service %s, starting update", serviceName)
                service.Internal = d.Get("internal").(bool)
	}
        if d.HasChange("ip") {
                log.Printf("[DEBUG] Ip has changed for service %s, starting update", serviceName)
                service.Ip = d.Get("ip").(string)
	}
        if d.HasChange("ipaddress") {
                log.Printf("[DEBUG] Ipaddress has changed for service %s, starting update", serviceName)
                service.Ipaddress = d.Get("ipaddress").(string)
	}
        if d.HasChange("maxbandwidth") {
                log.Printf("[DEBUG] Maxbandwidth has changed for service %s, starting update", serviceName)
                service.Maxbandwidth = d.Get("maxbandwidth").(int)
	}
        if d.HasChange("maxclient") {
                log.Printf("[DEBUG] Maxclient has changed for service %s, starting update", serviceName)
                service.Maxclient = d.Get("maxclient").(int)
	}
        if d.HasChange("maxreq") {
                log.Printf("[DEBUG] Maxreq has changed for service %s, starting update", serviceName)
                service.Maxreq = d.Get("maxreq").(int)
	}
        if d.HasChange("monitornamesvc") {
                log.Printf("[DEBUG] Monitornamesvc has changed for service %s, starting update", serviceName)
                service.Monitornamesvc = d.Get("monitornamesvc").(string)
	}
        if d.HasChange("monthreshold") {
                log.Printf("[DEBUG] Monthreshold has changed for service %s, starting update", serviceName)
                service.Monthreshold = d.Get("monthreshold").(int)
	}
        if d.HasChange("name") {
                log.Printf("[DEBUG] Name has changed for service %s, starting update", serviceName)
                service.Name = d.Get("name").(string)
	}
        if d.HasChange("netprofile") {
                log.Printf("[DEBUG] Netprofile has changed for service %s, starting update", serviceName)
                service.Netprofile = d.Get("netprofile").(string)
	}
        if d.HasChange("newname") {
                log.Printf("[DEBUG] Newname has changed for service %s, starting update", serviceName)
                service.Newname = d.Get("newname").(string)
	}
        if d.HasChange("pathmonitor") {
                log.Printf("[DEBUG] Pathmonitor has changed for service %s, starting update", serviceName)
                service.Pathmonitor = d.Get("pathmonitor").(string)
	}
        if d.HasChange("pathmonitorindv") {
                log.Printf("[DEBUG] Pathmonitorindv has changed for service %s, starting update", serviceName)
                service.Pathmonitorindv = d.Get("pathmonitorindv").(string)
	}
        if d.HasChange("port") {
                log.Printf("[DEBUG] Port has changed for service %s, starting update", serviceName)
                service.Port = d.Get("port").(int)
	}
        if d.HasChange("rtspsessionidremap") {
                log.Printf("[DEBUG] Rtspsessionidremap has changed for service %s, starting update", serviceName)
                service.Rtspsessionidremap = d.Get("rtspsessionidremap").(string)
	}
        if d.HasChange("sc") {
                log.Printf("[DEBUG] Sc has changed for service %s, starting update", serviceName)
                service.Sc = d.Get("sc").(string)
	}
        if d.HasChange("serverid") {
                log.Printf("[DEBUG] Serverid has changed for service %s, starting update", serviceName)
                service.Serverid = d.Get("serverid").(int)
	}
        if d.HasChange("servername") {
                log.Printf("[DEBUG] Servername has changed for service %s, starting update", serviceName)
                service.Servername = d.Get("servername").(string)
	}
        if d.HasChange("servicetype") {
                log.Printf("[DEBUG] Servicetype has changed for service %s, starting update", serviceName)
                service.Servicetype = d.Get("servicetype").(string)
	}
        if d.HasChange("sp") {
                log.Printf("[DEBUG] Sp has changed for service %s, starting update", serviceName)
                service.Sp = d.Get("sp").(string)
	}
        if d.HasChange("state") {
                log.Printf("[DEBUG] State has changed for service %s, starting update", serviceName)
                service.State = d.Get("state").(string)
	}
        if d.HasChange("svrtimeout") {
                log.Printf("[DEBUG] Svrtimeout has changed for service %s, starting update", serviceName)
                service.Svrtimeout = d.Get("svrtimeout").(int)
	}
        if d.HasChange("tcpb") {
                log.Printf("[DEBUG] Tcpb has changed for service %s, starting update", serviceName)
                service.Tcpb = d.Get("tcpb").(string)
	}
        if d.HasChange("tcpprofilename") {
                log.Printf("[DEBUG] Tcpprofilename has changed for service %s, starting update", serviceName)
                service.Tcpprofilename = d.Get("tcpprofilename").(string)
	}
        if d.HasChange("td") {
                log.Printf("[DEBUG] Td has changed for service %s, starting update", serviceName)
                service.Td = d.Get("td").(int)
	}
        if d.HasChange("useproxyport") {
                log.Printf("[DEBUG] Useproxyport has changed for service %s, starting update", serviceName)
                service.Useproxyport = d.Get("useproxyport").(string)
	}
        if d.HasChange("usip") {
                log.Printf("[DEBUG] Usip has changed for service %s, starting update", serviceName)
                service.Usip = d.Get("usip").(string)
	}
        if d.HasChange("weight") {
                log.Printf("[DEBUG] Weight has changed for service %s, starting update", serviceName)
                service.Weight = d.Get("weight").(int)
	}
        

        _, err := client.UpdateResource(netscaler.Service.Type(), serviceName, &service)
	if err != nil {
        return fmt.Errorf("Error updating service %s", serviceName)
	}
        return readServiceFunc(d, meta)
}

func deleteServiceFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
        serviceName := d.Id()
        err := client.DeleteResource(netscaler.Service.Type(), serviceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
