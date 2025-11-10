package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcServicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createServicegroupFunc,
		ReadContext:   readServicegroupFunc,
		UpdateContext: updateServicegroupFunc,
		DeleteContext: deleteServicegroupFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"topicname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"quicprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bootstrap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autodelayedtrofs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autodisabledelay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"autodisablegraceful": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cacheable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cka": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cmp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customserverid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbsttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dupweight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includemembers": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxbandwidth": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreq": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memberport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monconnectionclose": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nameserver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitorindv": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"riseapbrstatsmsgcode": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rtspsessionidremap": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpb": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"useproxyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lbvservers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"lbmonitor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"servicegroupmembers": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"servicegroupmembers_by_servername": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func createServicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In createServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	var servicegroupName string
	if v, ok := d.GetOk("servicegroupname"); ok {
		servicegroupName = v.(string)
	} else {
		servicegroupName = resource.PrefixedUniqueId("tf-servicegroup-")
		d.Set("servicegroupname", servicegroupName)
	}

	lbmonitor, mok := d.GetOk("lbmonitor")
	if mok {
		exists := client.ResourceExists(service.Lbmonitor.Type(), lbmonitor.(string))
		if !exists {
			return diag.Errorf("[ERROR] netscaler-provider: Specified lb monitor does not exist on netscaler!")
		}
	}

	var lbvservers []string
	l, lok := d.GetOk("lbvservers")
	if lok {
		lbvservers = expandStringList(l.(*schema.Set).List())
		for _, lbvserver := range lbvservers {
			exists := client.ResourceExists(service.Lbvserver.Type(), lbvserver)
			if !exists {
				return diag.Errorf("[ERROR] netscaler-provider: Specified lb vserver %s does not exist on netscaler!", lbvserver)
			}
		}
	}

	var groupmembers []string
	v, gok := d.GetOk("servicegroupmembers")
	if gok {
		groupmembers = expandStringList(v.(*schema.Set).List())
	}

	var groupmembersByServername []string
	vbs, gbsok := d.GetOk("servicegroupmembers_by_servername")
	if gbsok {
		groupmembersByServername = expandStringList(vbs.(*schema.Set).List())
	}

	servicegroup := basic.Servicegroup{
		Appflowlog:          d.Get("appflowlog").(string),
		Autodisablegraceful: d.Get("autodisablegraceful").(string),
		Autoscale:           d.Get("autoscale").(string),
		Cacheable:           d.Get("cacheable").(string),
		Cachetype:           d.Get("cachetype").(string),
		Cip:                 d.Get("cip").(string),
		Cipheader:           d.Get("cipheader").(string),
		Cka:                 d.Get("cka").(string),
		Cmp:                 d.Get("cmp").(string),
		Comment:             d.Get("comment").(string),
		Customserverid:      d.Get("customserverid").(string),
		Downstateflush:      d.Get("downstateflush").(string),
		Healthmonitor:       d.Get("healthmonitor").(string),
		Httpprofilename:     d.Get("httpprofilename").(string),
		Includemembers:      d.Get("includemembers").(bool),
		Monconnectionclose:  d.Get("monconnectionclose").(string),
		Monitornamesvc:      d.Get("monitornamesvc").(string),
		Nameserver:          d.Get("nameserver").(string),
		Netprofile:          d.Get("netprofile").(string),
		Pathmonitor:         d.Get("pathmonitor").(string),
		Pathmonitorindv:     d.Get("pathmonitorindv").(string),
		Rtspsessionidremap:  d.Get("rtspsessionidremap").(string),
		Servername:          d.Get("servername").(string),
		Servicegroupname:    d.Get("servicegroupname").(string),
		Servicetype:         d.Get("servicetype").(string),
		Sp:                  d.Get("sp").(string),
		State:               d.Get("state").(string),
		Tcpb:                d.Get("tcpb").(string),
		Tcpprofilename:      d.Get("tcpprofilename").(string),
		Useproxyport:        d.Get("useproxyport").(string),
		Usip:                d.Get("usip").(string),
		Autodelayedtrofs:    d.Get("autodelayedtrofs").(string),
		Bootstrap:           d.Get("bootstrap").(string),
		Quicprofilename:     d.Get("quicprofilename").(string),
		Topicname:           d.Get("topicname").(string),
	}
	if raw := d.GetRawConfig().GetAttr("autodisabledelay"); !raw.IsNull() {
		servicegroup.Autodisabledelay = intPtr(d.Get("autodisabledelay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("clttimeout"); !raw.IsNull() {
		servicegroup.Clttimeout = intPtr(d.Get("clttimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dbsttl"); !raw.IsNull() {
		servicegroup.Dbsttl = intPtr(d.Get("dbsttl").(int))
	}
	if raw := d.GetRawConfig().GetAttr("dupweight"); !raw.IsNull() {
		servicegroup.Dupweight = intPtr(d.Get("dupweight").(int))
	}
	if raw := d.GetRawConfig().GetAttr("hashid"); !raw.IsNull() {
		servicegroup.Hashid = intPtr(d.Get("hashid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxbandwidth"); !raw.IsNull() {
		servicegroup.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxclient"); !raw.IsNull() {
		servicegroup.Maxclient = intPtr(d.Get("maxclient").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxreq"); !raw.IsNull() {
		servicegroup.Maxreq = intPtr(d.Get("maxreq").(int))
	}
	if raw := d.GetRawConfig().GetAttr("memberport"); !raw.IsNull() {
		servicegroup.Memberport = intPtr(d.Get("memberport").(int))
	}
	if raw := d.GetRawConfig().GetAttr("monthreshold"); !raw.IsNull() {
		servicegroup.Monthreshold = intPtr(d.Get("monthreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		servicegroup.Port = intPtr(d.Get("port").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverid"); !raw.IsNull() {
		servicegroup.Serverid = intPtr(d.Get("serverid").(int))
	}
	if raw := d.GetRawConfig().GetAttr("svrtimeout"); !raw.IsNull() {
		servicegroup.Svrtimeout = intPtr(d.Get("svrtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("td"); !raw.IsNull() {
		servicegroup.Td = intPtr(d.Get("td").(int))
	}
	if raw := d.GetRawConfig().GetAttr("weight"); !raw.IsNull() {
		servicegroup.Weight = intPtr(d.Get("weight").(int))
	}

	_, err := client.AddResource(service.Servicegroup.Type(), servicegroupName, &servicegroup)
	if err != nil {
		return diag.FromErr(err)
	}
	if lok { //lbvservers is specified
		err = addLbvserverBindings(client, servicegroupName, lbvservers)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if mok { //lbmonitor is specified
		lbmonitorName := d.Get("lbmonitor").(string)
		binding := lb.Lbmonitorservicebinding{
			Monitorname:      lbmonitorName,
			Servicegroupname: servicegroupName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding servicegroup %s to lbmonitor %s", servicegroupName, lbmonitorName)
		err = client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Servicegroup.Type(), servicegroupName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind servicegroup %s to lbmonitor %s", servicegroupName, lbmonitorName)
			err2 := client.DeleteResource(service.Servicegroup.Type(), servicegroupName)
			if err2 != nil {
				log.Printf("[ERROR] netscaler-provider:  Failed to delete servicegroup %s after bind to lbmonitor failed", servicegroupName)
				return diag.Errorf("[ERROR] netscaler-provider:  Failed to delete servicegroup %s after bind to lbmonitor failed", servicegroupName)
			}
			return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind  servicegroup %s to lbmonitor %s", servicegroupName, lbmonitorName)
		}
	}

	if gok { //servicegroupmembers is specified
		createServicegroupMemberBindings(client, servicegroupName, groupmembers, false)
	}

	if gbsok { // servicegroupmembers_by_servername is specified
		createServicegroupMemberBindings(client, servicegroupName, groupmembersByServername, true)
	}

	d.SetId(servicegroupName)

	return readServicegroupFunc(ctx, d, meta)
}

func createServicegroupMemberBindings(client *service.NitroClient, servicegroupName string, groupmembers []string, bindByServername bool) error {
	for _, member := range groupmembers {
		//format is ip:port:weight
		parts := strings.Split(member, ":")
		var ip, servername string
		var port int
		var weight int
		if !bindByServername {
			ip = parts[0]
		} else {
			servername = parts[0]
		}
		if len(parts) < 2 {
			log.Printf("[WARN] netscaler-provider:  servicgroupmembers has invalid member: port not specified:%s", member)
			//TODO: take it from memberport
			continue
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("[WARN] netscaler-provider:  servicgroupmembers has invalid port: not an integer: %s", parts[1])
			continue
		}
		weightFound := false
		if len(parts) > 2 {
			weight, err = strconv.Atoi(parts[2])
			weightFound = true
			if err != nil {
				log.Printf("[WARN] netscaler-provider:  servicgroupmembers has invalid weight: not an integer:%s", parts[2])
			}
		}
		var binding basic.Servicegroupservicegroupmemberbinding
		if !bindByServername {
			binding = basic.Servicegroupservicegroupmemberbinding{
				Servicegroupname: servicegroupName,
				Ip:               ip,
				Port:             intPtr(port),
			}
		} else {
			binding = basic.Servicegroupservicegroupmemberbinding{
				Servicegroupname: servicegroupName,
				Servername:       servername,
				Port:             intPtr(port),
			}
		}

		if weightFound {
			binding.Weight = intPtr(weight)
		}
		log.Printf("[INFO] netscaler-provider:  Binding servicegroup %s to ip %s", servicegroupName, ip)
		_, err = client.AddResource(service.Servicegroup_servicegroupmember_binding.Type(), servicegroupName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind servicegroup %s to ip %s", servicegroupName, ip)
			continue //TODO: should be break here?
		}

	}
	return nil
}

func removeServicegroupMemberBindings(client *service.NitroClient, servicegroupName string, groupmembers []string, bindByServername bool) error {
	for _, member := range groupmembers {
		//format is ip:port:weight
		parts := strings.Split(member, ":")
		var ip, servername, port string
		if !bindByServername {
			ip = parts[0]
		} else {
			servername = parts[0]
		}
		if len(parts) < 2 {
			log.Printf("[WARN] netscaler-provider:  servicgroupmembers has invalid member: port not specified:%s", member)
			//TODO: take it from memberport
			continue
		}
		port = parts[1]
		log.Printf("[INFO] netscaler-provider:  UnBinding servicegroup %s from ip %s", servicegroupName, ip)
		args := make([]string, 1, 1)
		if !bindByServername {
			args[0] = fmt.Sprintf("ip:%s,port:%s", ip, port)
		} else {
			args[0] = fmt.Sprintf("servername:%s,port:%s", servername, port)
		}
		err := client.DeleteResourceWithArgs(service.Servicegroup_servicegroupmember_binding.Type(), servicegroupName, args)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to unbind servicegroup %s from ip %s", servicegroupName, ip)
			continue //TODO: should be break here?
		}

	}
	return nil
}

func addLbvserverBindings(client *service.NitroClient, servicegroupName string, lbvservers []string) error {
	for _, lbvserverName := range lbvservers {
		binding := lb.Lbvserverservicegroupbinding{
			Name:             lbvserverName,
			Servicegroupname: servicegroupName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding servicegroup %s to lbvserver %s", servicegroupName, lbvserverName)
		err := client.BindResource(service.Lbvserver.Type(), lbvserverName, service.Servicegroup.Type(), servicegroupName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind servicegroup %s to lbvserver %s", servicegroupName, lbvserverName)
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind  servicegroup %s to lbvserver %s", servicegroupName, lbvserverName)
		}
	}
	return nil
}

func removeLbvserverBindings(client *service.NitroClient, servicegroupName string, lbvservers []string) error {
	for _, lbvserverName := range lbvservers {
		err := client.UnbindResource(service.Lbvserver.Type(), lbvserverName, service.Servicegroup.Type(), servicegroupName, "servicegroupname")
		if err != nil {
			log.Printf("[ERROR] netscaler-provider: Error unbinding lbvserver %s from servicegroup %s", lbvserverName, servicegroupName)
			return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding lbvserver %s from servicegroup %s", lbvserverName, servicegroupName)
		}
	}
	return nil
}

func readServicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading servicegroup state %s", servicegroupName)
	data, err := client.FindResource(service.Servicegroup.Type(), servicegroupName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}
	//read bound service group members. Note that there is no type defined called service.Servicegroupmember.Type()
	boundMembers, err := client.FindAllBoundResources(service.Servicegroup.Type(), servicegroupName, "servicegroupmember")
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}
	//read bound vservers.
	vserverBindings, err := client.FindResourceArray(service.Servicegroupbindings.Type(), servicegroupName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}
	//read bound lb monitor.
	boundMonitors, err := client.FindAllBoundResources(service.Servicegroup.Type(), servicegroupName, service.Lbmonitor.Type())
	if err != nil {
		//This is actually OK in most cases
		log.Printf("[WARN] netscaler-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}

	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("topicname", data["topicname"])
	d.Set("quicprofilename", data["quicprofilename"])
	d.Set("bootstrap", data["bootstrap"])
	d.Set("autodelayedtrofs", data["autodelayedtrofs"])
	d.Set("appflowlog", data["appflowlog"])
	setToInt("autodisabledelay", d, data["autodisabledelay"])
	d.Set("autodisablegraceful", data["autodisablegraceful"])
	d.Set("autoscale", data["autoscale"])
	d.Set("cacheable", data["cacheable"])
	d.Set("cachetype", data["cachetype"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("cka", data["cka"])
	setToInt("clttimeout", d, data["clttimeout"])
	d.Set("cmp", data["cmp"])
	d.Set("comment", data["comment"])
	d.Set("customserverid", data["customserverid"])
	setToInt("dbsttl", d, data["dbsttl"])
	d.Set("downstateflush", data["downstateflush"])
	setToInt("dupweight", d, data["dupweight"])
	setToInt("hashid", d, data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("includemembers", data["includemembers"])
	setToInt("maxbandwidth", d, data["maxbandwidth"])
	setToInt("maxclient", d, data["maxclient"])
	setToInt("maxreq", d, data["maxreq"])
	setToInt("memberport", d, data["memberport"])
	d.Set("monconnectionclose", data["monconnectionclose"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	setToInt("monthreshold", d, data["monthreshold"])
	d.Set("nameserver", data["nameserver"])
	d.Set("netprofile", data["netprofile"])
	d.Set("pathmonitor", data["pathmonitor"])
	d.Set("pathmonitorindv", data["pathmonitorindv"])
	setToInt("port", d, data["port"])
	setToInt("riseapbrstatsmsgcode", d, data["riseapbrstatsmsgcode"])
	d.Set("rtspsessionidremap", data["rtspsessionidremap"])
	setToInt("serverid", d, data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicetype", data["servicetype"])
	if data["sp"] == "ON (but effectively OFF)" {
		d.Set("sp", "ON")
	} else {
		d.Set("sp", data["sp"])
	}
	d.Set("state", data["state"])
	setToInt("svrtimeout", d, data["svrtimeout"])
	d.Set("tcpb", data["tcpb"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	setToInt("td", d, data["td"])
	d.Set("useproxyport", data["useproxyport"])
	d.Set("usip", data["usip"])
	setToInt("weight", d, data["weight"])

	_, membersOk := d.GetOk("servicegroupmembers")
	_, membersByNameOk := d.GetOk("servicegroupmembers_by_servername")
	if membersOk || membersByNameOk {
		//boundMembers is of type []map[string]interface{}
		servicegroupMembers := make([]string, 0, len(boundMembers))
		servicegroupMembersByServername := make([]string, 0, len(boundMembers))
		for _, member := range boundMembers {
			ip := member["ip"].(string)
			servername := member["servername"].(string)
			port := member["port"].(float64) //TODO: why is this not int?
			weight := member["weight"].(string)
			// Heuristic rule
			var strmember string
			if servername == ip {
				strmember = fmt.Sprintf("%s:%.0f:%s", ip, port, weight)
				servicegroupMembers = append(servicegroupMembers, strmember)
			} else {
				strmember = fmt.Sprintf("%s:%.0f:%s", servername, port, weight)
				servicegroupMembersByServername = append(servicegroupMembersByServername, strmember)
			}
		}
		d.Set("servicegroupmembers", servicegroupMembers)
		d.Set("servicegroupmembers_by_servername", servicegroupMembersByServername)
	}

	//vserverBindings is of type []map[string]interface{}
	if _, ok := d.GetOk("lbvservers"); ok {
		var boundVserver string
		lbvservers := make([]string, 0, len(vserverBindings))
		for _, vserver := range vserverBindings {
			vs, ok := vserver["vservername"]
			if ok {
				boundVserver = vs.(string)
				lbvservers = append(lbvservers, boundVserver)
			}
		}
		d.Set("lbvservers", lbvservers)
	}

	var boundMonitor string
	for _, monitor := range boundMonitors {
		mon, ok := monitor["monitor_name"]
		if ok {
			boundMonitor = mon.(string)
			break
		}
	}
	// Need to do this due to explicit binding resource
	// We ignore lbmonitors if not explicitely defined in the resource
	if _, ok := d.GetOk("lbmonitor"); ok {
		d.Set("lbmonitor", boundMonitor)
	}

	return nil

}

func updateServicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In updateServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Get("servicegroupname").(string)

	servicegroup := basic.Servicegroup{
		Servicegroupname: d.Get("servicegroupname").(string),
	}

	stateChange := false
	hasChange := false
	if d.HasChange("quicprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Quicprofilename has changed for servicegroup, starting update")
		servicegroup.Quicprofilename = d.Get("quicprofilename").(string)
		hasChange = true
	}
	if d.HasChange("autodelayedtrofs") {
		log.Printf("[DEBUG]  citrixadc-provider: Autodelayedtrofs has changed for servicegroup, starting update")
		servicegroup.Autodelayedtrofs = d.Get("autodelayedtrofs").(string)
		hasChange = true
	}
	lbvserversChanged := false
	lbmonitorChanged := false
	servicegroupmembersChanged := false
	servicegroupmembersByServernameChanged := false

	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  netscaler-provider: Appflowlog has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("autodisabledelay") {
		log.Printf("[DEBUG]  netscaler-provider: Autodisabledelay has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autodisabledelay = intPtr(d.Get("autodisabledelay").(int))
		hasChange = true
	}
	if d.HasChange("autodisablegraceful") {
		log.Printf("[DEBUG]  netscaler-provider: Autodisablegraceful has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autodisablegraceful = d.Get("autodisablegraceful").(string)
		hasChange = true
	}
	if d.HasChange("autoscale") {
		log.Printf("[DEBUG]  netscaler-provider: Autoscale has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autoscale = d.Get("autoscale").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG]  netscaler-provider: Cacheable has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("cachetype") {
		log.Printf("[DEBUG]  netscaler-provider: Cachetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cachetype = d.Get("cachetype").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  netscaler-provider: Cip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  netscaler-provider: Cipheader has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cipheader = d.Get("cipheader").(string)
		servicegroup.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cka") {
		log.Printf("[DEBUG]  netscaler-provider: Cka has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cka = d.Get("cka").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Clttimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Clttimeout = intPtr(d.Get("clttimeout").(int))
		hasChange = true
	}
	if d.HasChange("cmp") {
		log.Printf("[DEBUG]  netscaler-provider: Cmp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cmp = d.Get("cmp").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("customserverid") {
		log.Printf("[DEBUG]  netscaler-provider: Customserverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Customserverid = d.Get("customserverid").(string)
		hasChange = true
	}
	if d.HasChange("dbsttl") {
		log.Printf("[DEBUG]  netscaler-provider: Dbsttl has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Dbsttl = intPtr(d.Get("dbsttl").(int))
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  netscaler-provider: Downstateflush has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("dupweight") {
		log.Printf("[DEBUG]  netscaler-provider: Dupweight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Dupweight = intPtr(d.Get("dupweight").(int))
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  netscaler-provider: Hashid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Hashid = intPtr(d.Get("hashid").(int))
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Healthmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Httpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("includemembers") {
		log.Printf("[DEBUG]  netscaler-provider: Includemembers has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Includemembers = d.Get("includemembers").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  netscaler-provider: Maxbandwidth has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxbandwidth = intPtr(d.Get("maxbandwidth").(int))
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  netscaler-provider: Maxclient has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxclient = intPtr(d.Get("maxclient").(int))
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG]  netscaler-provider: Maxreq has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxreq = intPtr(d.Get("maxreq").(int))
		hasChange = true
	}
	if d.HasChange("memberport") {
		log.Printf("[DEBUG]  netscaler-provider: Memberport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Memberport = intPtr(d.Get("memberport").(int))
		hasChange = true
	}
	if d.HasChange("monconnectionclose") {
		log.Printf("[DEBUG]  netscaler-provider: Monconnectionclose has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monconnectionclose = d.Get("monconnectionclose").(string)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  netscaler-provider: Monitornamesvc has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  netscaler-provider: Monthreshold has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monthreshold = intPtr(d.Get("monthreshold").(int))
		hasChange = true
	}
	if d.HasChange("nameserver") {
		log.Printf("[DEBUG]  netscaler-provider: Nameserver has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Nameserver = d.Get("nameserver").(string)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  netscaler-provider: Netprofile has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Pathmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitor = d.Get("pathmonitor").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitorindv") {
		log.Printf("[DEBUG]  netscaler-provider: Pathmonitorindv has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitorindv = d.Get("pathmonitorindv").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  netscaler-provider: Port has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Port = intPtr(d.Get("port").(int))
		hasChange = true
	}
	if d.HasChange("rtspsessionidremap") {
		log.Printf("[DEBUG]  netscaler-provider: Rtspsessionidremap has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Rtspsessionidremap = d.Get("rtspsessionidremap").(string)
		hasChange = true
	}
	if d.HasChange("serverid") {
		log.Printf("[DEBUG]  netscaler-provider: Serverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Serverid = intPtr(d.Get("serverid").(int))
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  netscaler-provider: Servername has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG]  netscaler-provider: Servicegroupname has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  netscaler-provider: Servicetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sp") {
		log.Printf("[DEBUG]  netscaler-provider: Sp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Sp = d.Get("sp").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for servicegroup %s, starting update", servicegroupName)
		stateChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Svrtimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Svrtimeout = intPtr(d.Get("svrtimeout").(int))
		hasChange = true
	}
	if d.HasChange("tcpb") {
		log.Printf("[DEBUG]  netscaler-provider: Tcpb has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpb = d.Get("tcpb").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Tcpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Td = intPtr(d.Get("td").(int))
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  netscaler-provider: Useproxyport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}
	if d.HasChange("usip") {
		log.Printf("[DEBUG]  netscaler-provider: Usip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Usip = d.Get("usip").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  netscaler-provider: Weight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Weight = intPtr(d.Get("weight").(int))
		hasChange = true
	}
	if d.HasChange("lbvservers") {
		log.Printf("[DEBUG] netscaler-provider:  lb vservers has changed for servicegroup %s, starting update", servicegroupName)
		lbvserversChanged = true
	}
	if d.HasChange("lbmonitor") {
		log.Printf("[DEBUG] netscaler-provider:  lb monitor has changed for servicegroup %s, starting update", servicegroupName)
		lbmonitorChanged = true
	}
	if d.HasChange("servicegroupmembers") {
		log.Printf("[DEBUG] netscaler-provider:  servicegroup membership has changed for servicegroup %s, starting update", servicegroupName)
		lbmonitorChanged = true
		servicegroupmembersChanged = true
	}
	if d.HasChange("servicegroupmembers_by_servername") {
		log.Printf("[DEBUG] netscaler-provider:  servicegroup membership has changed for servicegroup %s, starting update", servicegroupName)
		lbmonitorChanged = true
		servicegroupmembersByServernameChanged = true
	}

	if lbvserversChanged {
		//Binding has to be updated
		o, n := d.GetChange("lbvservers")
		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		remove := expandStringList(os.Difference(ns).List())
		add := expandStringList(ns.Difference(os).List())

		if len(remove) > 0 {
			err := removeLbvserverBindings(client, servicegroupName, remove)
			if err != nil {
				return diag.FromErr(err)
			}
		}
		if len(add) > 0 {
			err := addLbvserverBindings(client, servicegroupName, add)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	lbmonitor := d.Get("lbmonitor")
	lbmonitorName := lbmonitor.(string)
	if lbmonitorChanged {
		//Binding has to be updated
		//First we unbind from lb monitor
		oldLbmonitor, _ := d.GetChange("lbmonitor")
		oldLbmonitorName := oldLbmonitor.(string)
		if oldLbmonitorName != "" {
			err := client.UnbindResource(service.Lbmonitor.Type(), oldLbmonitorName, service.Servicegroup.Type(), servicegroupName, "servicegroupname")
			if err != nil {
				return diag.Errorf("[ERROR] netscaler-provider: Error unbinding lbmonitor from servicegroup %s", oldLbmonitorName)
			}
			log.Printf("[DEBUG] netscaler-provider: lbmonitor has been unbound from servicegroup for lb monitor %s ", oldLbmonitorName)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Servicegroup.Type(), servicegroupName, &servicegroup)
		if err != nil {
			return diag.Errorf("Error updating servicegroup %s", servicegroupName)
		}
	}

	if lbmonitorChanged && lbmonitorName != "" {
		//Binding has to be updated
		//rebind
		binding := lb.Lbmonitorservicebinding{
			Monitorname:      lbmonitorName,
			Servicegroupname: servicegroupName,
		}
		log.Printf("[INFO] netscaler-provider:  Binding monitor %s to servicegroup %s", lbmonitorName, servicegroupName)
		err := client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Servicegroup.Type(), servicegroupName, &binding)
		if err != nil {
			log.Printf("[ERROR] netscaler-provider:  Failed to bind  lbmonitor %s to servicegroup %s", lbmonitorName, servicegroupName)
			return diag.Errorf("[ERROR] netscaler-provider:  Failed to bind lb monitor %s to servicegroup %s", lbmonitorName, servicegroupName)
		}
		log.Printf("[DEBUG] netscaler-provider: new lbmonitor has been bound to servicegroup  lbmonitor %s servicegroup %s", lbmonitorName, servicegroupName)
	}

	if servicegroupmembersChanged {
		o, n := d.GetChange("servicegroupmembers")
		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		remove := expandStringList(os.Difference(ns).List())
		add := expandStringList(ns.Difference(os).List())

		if len(remove) > 0 {
			removeServicegroupMemberBindings(client, servicegroupName, remove, false)
		}
		if len(add) > 0 {
			createServicegroupMemberBindings(client, servicegroupName, add, false)
		}

	}

	if servicegroupmembersByServernameChanged {
		o, n := d.GetChange("servicegroupmembers_by_servername")
		os := o.(*schema.Set)
		ns := n.(*schema.Set)

		remove := expandStringList(os.Difference(ns).List())
		add := expandStringList(ns.Difference(os).List())

		if len(remove) > 0 {
			removeServicegroupMemberBindings(client, servicegroupName, remove, true)
		}
		if len(add) > 0 {
			createServicegroupMemberBindings(client, servicegroupName, add, true)
		}

	}

	if stateChange {
		err := doServicegroupStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling servicegroup %s", servicegroupName)
		}
	}

	return readServicegroupFunc(ctx, d, meta)
}

func deleteServicegroupFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	err := client.DeleteResource(service.Servicegroup.Type(), servicegroupName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []string
func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		vs = append(vs, v.(string))
	}
	return vs
}

func doServicegroupStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doServicegroupStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	serviceGroup := basic.Servicegroup{
		Servicegroupname: d.Get("servicegroupname").(string),
		Servername:       d.Get("servername").(string),
	}

	if raw := d.GetRawConfig().GetAttr("port"); !raw.IsNull() {
		serviceGroup.Port = intPtr(d.Get("port").(int))
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Servicegroup.Type(), serviceGroup, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		serviceGroup.Delay = intPtr(d.Get("delay").(int))
		serviceGroup.Graceful = d.Get("graceful").(string)
		err := client.ActOnResource(service.Servicegroup.Type(), serviceGroup, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
