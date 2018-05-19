package netscaler

import (
	"github.com/chiradeep/go-nitro/config/gslb"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerGslbservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbserviceFunc,
		Read:          readGslbserviceFunc,
		Update:        updateGslbserviceFunc,
		Delete:        deleteGslbserviceFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
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
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cnameentry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
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
			"maxaaausers": &schema.Schema{
				Type:     schema.TypeInt,
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
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"publicip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sitepersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteprefix": &schema.Schema{
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
			"viewip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"viewname": &schema.Schema{
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

func createGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbserviceName string
	if v, ok := d.GetOk("servicename"); ok {
		gslbserviceName = v.(string)
	} else {
		gslbserviceName = resource.PrefixedUniqueId("tf-gslbservice-")
		d.Set("servicename", gslbserviceName)
	}
	gslbservice := gslb.Gslbservice{
		Appflowlog:      d.Get("appflowlog").(string),
		Cip:             d.Get("cip").(string),
		Cipheader:       d.Get("cipheader").(string),
		Clttimeout:      d.Get("clttimeout").(int),
		Cnameentry:      d.Get("cnameentry").(string),
		Comment:         d.Get("comment").(string),
		Cookietimeout:   d.Get("cookietimeout").(int),
		Downstateflush:  d.Get("downstateflush").(string),
		Hashid:          d.Get("hashid").(int),
		Healthmonitor:   d.Get("healthmonitor").(string),
		Ip:              d.Get("ip").(string),
		Ipaddress:       d.Get("ipaddress").(string),
		Maxaaausers:     d.Get("maxaaausers").(int),
		Maxbandwidth:    d.Get("maxbandwidth").(int),
		Maxclient:       d.Get("maxclient").(int),
		Monitornamesvc:  d.Get("monitornamesvc").(string),
		Monthreshold:    d.Get("monthreshold").(int),
		Newname:         d.Get("newname").(string),
		Port:            d.Get("port").(int),
		Publicip:        d.Get("publicip").(string),
		Publicport:      d.Get("publicport").(int),
		Servername:      d.Get("servername").(string),
		Servicename:     d.Get("servicename").(string),
		Servicetype:     d.Get("servicetype").(string),
		Sitename:        d.Get("sitename").(string),
		Sitepersistence: d.Get("sitepersistence").(string),
		Siteprefix:      d.Get("siteprefix").(string),
		State:           d.Get("state").(string),
		Svrtimeout:      d.Get("svrtimeout").(int),
		Viewip:          d.Get("viewip").(string),
		Viewname:        d.Get("viewname").(string),
		Weight:          d.Get("weight").(int),
	}

	_, err := client.AddResource(netscaler.Gslbservice.Type(), gslbserviceName, &gslbservice)
	if err != nil {
		return err
	}

	d.SetId(gslbserviceName)

	err = readGslbserviceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservice but we can't read it ?? %s", gslbserviceName)
		return nil
	}
	return nil
}

func readGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading gslbservice state %s", gslbserviceName)
	data, err := client.FindResource(netscaler.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing gslbservice state %s", gslbserviceName)
		d.SetId("")
		return nil
	}
	d.Set("servicename", data["servicename"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("cnameentry", data["cnameentry"])
	d.Set("comment", data["comment"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("ip", data["ipaddress"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("maxaaausers", data["maxaaausers"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxclient", data["maxclient"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("newname", data["newname"])
	d.Set("port", data["port"])
	d.Set("publicip", data["publicip"])
	d.Set("publicport", data["publicport"])
	d.Set("servername", data["servername"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitename", data["sitename"])
	d.Set("sitepersistence", data["sitepersistence"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("viewip", data["viewip"])
	d.Set("viewname", data["viewname"])
	d.Set("weight", data["weight"])

	return nil

}

func updateGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Get("servicename").(string)

	gslbservice := gslb.Gslbservice{
		Servicename: d.Get("servicename").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  netscaler-provider: Appflowlog has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  netscaler-provider: Cip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  netscaler-provider: Cipheader has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Clttimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cnameentry") {
		log.Printf("[DEBUG]  netscaler-provider: Cnameentry has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cnameentry = d.Get("cnameentry").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Cookietimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  netscaler-provider: Downstateflush has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  netscaler-provider: Hashid has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Healthmonitor has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("ip") {
		log.Printf("[DEBUG]  netscaler-provider: Ip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ip = d.Get("ip").(string)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  netscaler-provider: Ipaddress has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("maxaaausers") {
		log.Printf("[DEBUG]  netscaler-provider: Maxaaausers has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxaaausers = d.Get("maxaaausers").(int)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  netscaler-provider: Maxbandwidth has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  netscaler-provider: Maxclient has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  netscaler-provider: Monitornamesvc has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  netscaler-provider: Monthreshold has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  netscaler-provider: Newname has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  netscaler-provider: Port has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("publicport") {
		log.Printf("[DEBUG]  netscaler-provider: Publicport has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicport = d.Get("publicport").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  netscaler-provider: Servername has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  netscaler-provider: Servicename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  netscaler-provider: Servicetype has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  netscaler-provider: Sitename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitename = d.Get("sitename").(string)
		hasChange = true
	}
	if d.HasChange("sitepersistence") {
		log.Printf("[DEBUG]  netscaler-provider: Sitepersistence has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitepersistence = d.Get("sitepersistence").(string)
		hasChange = true
	}
	if d.HasChange("siteprefix") {
		log.Printf("[DEBUG]  netscaler-provider: Siteprefix has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Siteprefix = d.Get("siteprefix").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Svrtimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("viewip") {
		log.Printf("[DEBUG]  netscaler-provider: Viewip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewip = d.Get("viewip").(string)
		hasChange = true
	}
	if d.HasChange("viewname") {
		log.Printf("[DEBUG]  netscaler-provider: Viewname has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewname = d.Get("viewname").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  netscaler-provider: Weight has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Gslbservice.Type(), gslbserviceName, &gslbservice)
		if err != nil {
			return fmt.Errorf("Error updating gslbservice %s", gslbserviceName)
		}
	}
	return readGslbserviceFunc(d, meta)
}

func deleteGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	err := client.DeleteResource(netscaler.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
