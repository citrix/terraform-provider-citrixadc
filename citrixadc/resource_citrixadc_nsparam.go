package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcNsparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsparamFunc,
		Read:          readNsparamFunc,
		Delete:        deleteNsparamFunc,
		Schema: map[string]*schema.Schema{
			"advancedanalyticsstats": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"aftpallowrandomsourceport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cipheader": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookieversion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"crportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exclusivequotamaxclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exclusivequotaspillover": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ftpportrange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"grantquotamaxclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"grantquotaspillover": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"internaluserlogin": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxreq": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mgmthttpport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mgmthttpsport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pmtumin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pmtutimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"proxyprotocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"securecookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicepathingressvlan": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpcip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"useproxyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsparamId := resource.PrefixedUniqueId("tf-nsparam-")
	nsparam := make(map[string]interface{})
	if data, ok := d.GetOk("advancedanalyticsstats"); ok {
		nsparam["advancedanalyticsstats"] = data.(string)
	}
	if data, ok := d.GetOk("aftpallowrandomsourceport"); ok {
		nsparam["aftpallowrandomsourceport"] = data.(string)
	}
	if data, ok := d.GetOk("cip"); ok {
		nsparam["cip"] = data.(string)
	}
	if data, ok := d.GetOk("cipheader"); ok {
		nsparam["cipheader"] = data.(string)
	}
	if data, ok := d.GetOk("cookieversion"); ok {
		nsparam["cookieversion"] = data.(string)
	}
	if data, ok := d.GetOk("crportrange"); ok {
		nsparam["crportrange"] = data.(string)
	}
	if data, ok := d.GetOk("exclusivequotamaxclient"); ok {
		nsparam["exclusivequotamaxclient"] = data.(string)
	}
	if data, ok := d.GetOk("exclusivequotaspillover"); ok {
		nsparam["exclusivequotaspillover"] = data.(string)
	}
	if data, ok := d.GetOk("ftpportrange"); ok {
		nsparam["ftpportrange"] = data.(string)
	}
	if data, ok := d.GetOk("grantquotaspillover"); ok {
		nsparam["grantquotaspillover"] = data.(string)
	}
	if data, ok := d.GetOk("grantquotamaxclient"); ok {
		nsparam["grantquotamaxclient"] = data.(string)
	}
	if data, ok := d.GetOk("internaluserlogin"); ok {
		nsparam["internaluserlogin"] = data.(string)
	}
	if data, ok := d.GetOkExists("maxconn"); ok {
		nsparam["maxconn"] = data.(int)
	}
	if data, ok := d.GetOkExists("maxreq"); ok {
		nsparam["maxreq"] = data.(int)
	}
	if data, ok := d.GetOk("mgmthttpport"); ok {
		nsparam["mgmthttpport"] = data.(int)
	}
	if data, ok := d.GetOk("mgmthttpsport"); ok {
		nsparam["mgmthttpsport"] = data.(int)
	}
	if data, ok := d.GetOk("pmtumin"); ok {
		nsparam["pmtumin"] = data.(int)
	}
	if data, ok := d.GetOk("pmtutimeout"); ok {
		nsparam["pmtutimeout"] = data.(int)
	}
	if data, ok := d.GetOk("proxyprotocol"); ok {
		nsparam["proxyprotocol"] = data.(string)
	}
	if data, ok := d.GetOk("securecookie"); ok {
		nsparam["securecookie"] = data.(string)
	}
	if data, ok := d.GetOk("servicepathingressvlan"); ok {
		nsparam["servicepathingressvlan"] = data.(string)
	}
	if data, ok := d.GetOk("tcpcip"); ok {
		nsparam["tcpcip"] = data.(string)
	}
	if data, ok := d.GetOk("timezone"); ok {
		nsparam["timezone"] = data.(string)
	}
	if data, ok := d.GetOk("useproxyport"); ok {
		nsparam["useproxyport"] = data.(string)
	}
	err := client.UpdateUnnamedResource(service.Nsparam.Type(), &nsparam)
	if err != nil {
		return err
	}

	d.SetId(nsparamId)

	err = readNsparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsparam but we can't read it ??")
		return err
	}
	return nil
}

func readNsparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nsparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsparam state %s", nsparamName)
	data, err := client.FindResource(service.Nsparam.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsparam state %s", nsparamName)
		d.SetId("")
		return nil
	}
	d.Set("advancedanalyticsstats", data["advancedanalyticsstats"])
	d.Set("aftpallowrandomsourceport", data["aftpallowrandomsourceport"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("cookieversion", data["cookieversion"])
	d.Set("crportrange", data["crportrange"])
	d.Set("exclusivequotamaxclient", data["exclusivequotamaxclient"])
	d.Set("exclusivequotaspillover", data["exclusivequotaspillover"])
	d.Set("ftpportrange", data["ftpportrange"])
	d.Set("grantquotamaxclient", data["grantquotamaxclient"])
	d.Set("grantquotaspillover", data["grantquotaspillover"])
	d.Set("internaluserlogin", data["internaluserlogin"])
	d.Set("maxconn", data["maxconn"])
	d.Set("maxreq", data["maxreq"])
	d.Set("mgmthttpport", data["mgmthttpport"])
	d.Set("mgmthttpsport", data["mgmthttpsport"])
	d.Set("pmtumin", data["pmtumin"])
	d.Set("pmtutimeout", data["pmtutimeout"])
	d.Set("proxyprotocol", data["proxyprotocol"])
	d.Set("securecookie", data["securecookie"])
	d.Set("servicepathingressvlan", data["servicepathingressvlan"])
	d.Set("tcpcip", data["tcpcip"])
	d.Set("timezone", data["timezone"])
	d.Set("useproxyport", data["useproxyport"])

	return nil

}

func deleteNsparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsparamFunc")

	d.SetId("")

	return nil
}
