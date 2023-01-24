package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnvserverFunc,
		Read:          readVpnvserverFunc,
		Update:        updateVpnvserverFunc,
		Delete:        deleteVpnvserverFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicetype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advancedepa": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"appflowlog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certkeynames": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cginfrahomepageredirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deploymenttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devicecert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"doublehop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failedlogintimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icaonly": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icaproxysessionmigration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpvsrresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipset": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv46": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2conn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linuxepapluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpolicy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"loginonce": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logoutonsmartcardremoval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macepapluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxaaausers": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxloginattempts": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "newname": {
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"pcoipvserverprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"range": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rdpserverprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rhistate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samesite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userdomains": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverfqdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windowsepapluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnvserverName := d.Get("name").(string)
	vpnvserver := vpn.Vpnvserver{
		Advancedepa:              d.Get("advancedepa").(string),
		Appflowlog:               d.Get("appflowlog").(string),
		Authentication:           d.Get("authentication").(string),
		Authnprofile:             d.Get("authnprofile").(string),
		Certkeynames:             d.Get("certkeynames").(string),
		Cginfrahomepageredirect:  d.Get("cginfrahomepageredirect").(string),
		Comment:                  d.Get("comment").(string),
		Deploymenttype:           d.Get("deploymenttype").(string),
		Devicecert:               d.Get("devicecert").(string),
		Doublehop:                d.Get("doublehop").(string),
		Downstateflush:           d.Get("downstateflush").(string),
		Dtls:                     d.Get("dtls").(string),
		Failedlogintimeout:       d.Get("failedlogintimeout").(int),
		Httpprofilename:          d.Get("httpprofilename").(string),
		Icaonly:                  d.Get("icaonly").(string),
		Icaproxysessionmigration: d.Get("icaproxysessionmigration").(string),
		Icmpvsrresponse:          d.Get("icmpvsrresponse").(string),
		Ipset:                    d.Get("ipset").(string),
		Ipv46:                    d.Get("ipv46").(string),
		L2conn:                   d.Get("l2conn").(string),
		Linuxepapluginupgrade:    d.Get("linuxepapluginupgrade").(string),
		Listenpolicy:             d.Get("listenpolicy").(string),
		Listenpriority:           d.Get("listenpriority").(int),
		Loginonce:                d.Get("loginonce").(string),
		Logoutonsmartcardremoval: d.Get("logoutonsmartcardremoval").(string),
		Macepapluginupgrade:      d.Get("macepapluginupgrade").(string),
		Maxaaausers:              d.Get("maxaaausers").(int),
		Maxloginattempts:         d.Get("maxloginattempts").(int),
		Name:                     vpnvserverName,
		Netprofile:               d.Get("netprofile").(string),
		// Newname:                  d.Get("newname").(string),
		Pcoipvserverprofilename: d.Get("pcoipvserverprofilename").(string),
		Port:                    d.Get("port").(int),
		Range:                   d.Get("range").(int),
		Rdpserverprofilename:    d.Get("rdpserverprofilename").(string),
		Rhistate:                d.Get("rhistate").(string),
		Samesite:                d.Get("samesite").(string),
		Servicetype:             d.Get("servicetype").(string),
		State:                   d.Get("state").(string),
		Tcpprofilename:          d.Get("tcpprofilename").(string),
		Userdomains:             d.Get("userdomains").(string),
		Vserverfqdn:             d.Get("vserverfqdn").(string),
		Windowsepapluginupgrade: d.Get("windowsepapluginupgrade").(string),
	}

	_, err := client.AddResource(service.Vpnvserver.Type(), vpnvserverName, &vpnvserver)
	if err != nil {
		return err
	}

	d.SetId(vpnvserverName)

	err = readVpnvserverFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnvserver but we can't read it ?? %s", vpnvserverName)
		return nil
	}
	return nil
}

func readVpnvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnvserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnvserver state %s", vpnvserverName)
	data, err := client.FindResource(service.Vpnvserver.Type(), vpnvserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnvserver state %s", vpnvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("advancedepa", data["advancedepa"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authnprofile", data["authnprofile"])
	d.Set("certkeynames", data["certkeynames"])
	d.Set("cginfrahomepageredirect", data["cginfrahomepageredirect"])
	d.Set("comment", data["comment"])
	d.Set("deploymenttype", data["deploymenttype"])
	d.Set("devicecert", data["devicecert"])
	d.Set("doublehop", data["doublehop"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("dtls", data["dtls"])
	d.Set("failedlogintimeout", data["failedlogintimeout"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icaonly", data["icaonly"])
	d.Set("icaproxysessionmigration", data["icaproxysessionmigration"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("ipset", data["ipset"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("linuxepapluginupgrade", data["linuxepapluginupgrade"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("loginonce", data["loginonce"])
	d.Set("logoutonsmartcardremoval", data["logoutonsmartcardremoval"])
	d.Set("macepapluginupgrade", data["macepapluginupgrade"])
	d.Set("maxaaausers", data["maxaaausers"])
	d.Set("maxloginattempts", data["maxloginattempts"])
	// d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	// d.Set("newname", data["newname"])
	d.Set("pcoipvserverprofilename", data["pcoipvserverprofilename"])
	d.Set("port", data["port"])
	d.Set("range", data["range"])
	d.Set("rdpserverprofilename", data["rdpserverprofilename"])
	d.Set("rhistate", data["rhistate"])
	d.Set("samesite", data["samesite"])
	d.Set("servicetype", data["servicetype"])
	d.Set("state", data["state"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("userdomains", data["userdomains"])
	d.Set("vserverfqdn", data["vserverfqdn"])
	d.Set("windowsepapluginupgrade", data["windowsepapluginupgrade"])

	return nil

}

func updateVpnvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnvserverName := d.Get("name").(string)

	vpnvserver := vpn.Vpnvserver{
		Name: vpnvserverName,
	}
	hasChange := false
	if d.HasChange("advancedepa") {
		log.Printf("[DEBUG]  citrixadc-provider: Advancedepa has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Advancedepa = d.Get("advancedepa").(string)
		hasChange = true
	}
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  citrixadc-provider: Appflowlog has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Authnprofile has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Authnprofile = d.Get("authnprofile").(string)
		hasChange = true
	}
	if d.HasChange("certkeynames") {
		log.Printf("[DEBUG]  citrixadc-provider: Certkeynames has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Certkeynames = d.Get("certkeynames").(string)
		hasChange = true
	}
	if d.HasChange("cginfrahomepageredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Cginfrahomepageredirect has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Cginfrahomepageredirect = d.Get("cginfrahomepageredirect").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("deploymenttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Deploymenttype has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Deploymenttype = d.Get("deploymenttype").(string)
		hasChange = true
	}
	if d.HasChange("devicecert") {
		log.Printf("[DEBUG]  citrixadc-provider: Devicecert has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Devicecert = d.Get("devicecert").(string)
		hasChange = true
	}
	if d.HasChange("doublehop") {
		log.Printf("[DEBUG]  citrixadc-provider: Doublehop has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Doublehop = d.Get("doublehop").(string)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  citrixadc-provider: Downstateflush has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("dtls") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtls has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Dtls = d.Get("dtls").(string)
		hasChange = true
	}
	if d.HasChange("failedlogintimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Failedlogintimeout has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Failedlogintimeout = d.Get("failedlogintimeout").(int)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpprofilename has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("icaonly") {
		log.Printf("[DEBUG]  citrixadc-provider: Icaonly has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Icaonly = d.Get("icaonly").(string)
		hasChange = true
	}
	if d.HasChange("icaproxysessionmigration") {
		log.Printf("[DEBUG]  citrixadc-provider: Icaproxysessionmigration has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Icaproxysessionmigration = d.Get("icaproxysessionmigration").(string)
		hasChange = true
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpvsrresponse has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
		hasChange = true
	}
	if d.HasChange("ipset") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipset has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Ipset = d.Get("ipset").(string)
		hasChange = true
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG]  citrixadc-provider: Ipv46 has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Ipv46 = d.Get("ipv46").(string)
		hasChange = true
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG]  citrixadc-provider: L2conn has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.L2conn = d.Get("l2conn").(string)
		hasChange = true
	}
	if d.HasChange("linuxepapluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Linuxepapluginupgrade has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Linuxepapluginupgrade = d.Get("linuxepapluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpolicy has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Listenpolicy = d.Get("listenpolicy").(string)
		hasChange = true
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Listenpriority has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Listenpriority = d.Get("listenpriority").(int)
		hasChange = true
	}
	if d.HasChange("loginonce") {
		log.Printf("[DEBUG]  citrixadc-provider: Loginonce has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Loginonce = d.Get("loginonce").(string)
		hasChange = true
	}
	if d.HasChange("logoutonsmartcardremoval") {
		log.Printf("[DEBUG]  citrixadc-provider: Logoutonsmartcardremoval has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Logoutonsmartcardremoval = d.Get("logoutonsmartcardremoval").(string)
		hasChange = true
	}
	if d.HasChange("macepapluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Macepapluginupgrade has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Macepapluginupgrade = d.Get("macepapluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("maxaaausers") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxaaausers has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Maxaaausers = d.Get("maxaaausers").(int)
		hasChange = true
	}
	if d.HasChange("maxloginattempts") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxloginattempts has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Maxloginattempts = d.Get("maxloginattempts").(int)
		hasChange = true
	}
	// if d.HasChange("name") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Name has changed for vpnvserver %s, starting update", vpnvserverName)
	// 	vpnvserver.Name = d.Get("name").(string)
	// 	hasChange = true
	// }
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Netprofile has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	// if d.HasChange("newname") {
	// 	log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for vpnvserver %s, starting update", vpnvserverName)
	// 	vpnvserver.Newname = d.Get("newname").(string)
	// 	hasChange = true
	// }
	if d.HasChange("pcoipvserverprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcoipvserverprofilename has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Pcoipvserverprofilename = d.Get("pcoipvserverprofilename").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  citrixadc-provider: Port has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG]  citrixadc-provider: Range has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Range = d.Get("range").(int)
		hasChange = true
	}
	if d.HasChange("rdpserverprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpserverprofilename has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Rdpserverprofilename = d.Get("rdpserverprofilename").(string)
		hasChange = true
	}
	if d.HasChange("rhistate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rhistate has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Rhistate = d.Get("rhistate").(string)
		hasChange = true
	}
	if d.HasChange("samesite") {
		log.Printf("[DEBUG]  citrixadc-provider: Samesite has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Samesite = d.Get("samesite").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicetype has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpprofilename has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("userdomains") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdomains has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Userdomains = d.Get("userdomains").(string)
		hasChange = true
	}
	if d.HasChange("vserverfqdn") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserverfqdn has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Vserverfqdn = d.Get("vserverfqdn").(string)
		hasChange = true
	}
	if d.HasChange("windowsepapluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowsepapluginupgrade has changed for vpnvserver %s, starting update", vpnvserverName)
		vpnvserver.Windowsepapluginupgrade = d.Get("windowsepapluginupgrade").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnvserver.Type(), vpnvserverName, &vpnvserver)
		if err != nil {
			return fmt.Errorf("Error updating vpnvserver %s", vpnvserverName)
		}
	}
	return readVpnvserverFunc(d, meta)
}

func deleteVpnvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnvserverName := d.Id()
	err := client.DeleteResource(service.Vpnvserver.Type(), vpnvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
