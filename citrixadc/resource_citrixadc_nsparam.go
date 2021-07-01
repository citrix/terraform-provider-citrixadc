package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

func resourceCitrixAdcNsparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsparamFunc,
		Read:          readNsparamFunc,
		Delete:        deleteNsparamFunc,
		Schema: map[string]*schema.Schema{
			"advancedanalyticsstats": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"aftpallowrandomsourceport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cookieversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"crportrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exclusivequotamaxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"exclusivequotaspillover": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ftpportrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"grantquotamaxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"grantquotaspillover": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"internaluserlogin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mgmthttpport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mgmthttpsport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pmtumin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pmtutimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"proxyprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"securecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicepathingressvlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"useproxyport": &schema.Schema{
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
	nsparam := ns.Nsparam{
		Advancedanalyticsstats:    d.Get("advancedanalyticsstats").(string),
		Aftpallowrandomsourceport: d.Get("aftpallowrandomsourceport").(string),
		Cip:                       d.Get("cip").(string),
		Cipheader:                 d.Get("cipheader").(string),
		Cookieversion:             d.Get("cookieversion").(string),
		Crportrange:               d.Get("crportrange").(string),
		Exclusivequotamaxclient:   uint32(d.Get("exclusivequotamaxclient").(int)),
		Exclusivequotaspillover:   uint32(d.Get("exclusivequotaspillover").(int)),
		Ftpportrange:              d.Get("ftpportrange").(string),
		Grantquotamaxclient:       uint32(d.Get("grantquotamaxclient").(int)),
		Grantquotaspillover:       uint32(d.Get("grantquotaspillover").(int)),
		Internaluserlogin:         d.Get("internaluserlogin").(string),
		Maxconn:                   uint32(d.Get("maxconn").(int)),
		Maxreq:                    uint32(d.Get("maxreq").(int)),
		Mgmthttpport:              int32(d.Get("mgmthttpport").(int)),
		Mgmthttpsport:             int32(d.Get("mgmthttpsport").(int)),
		Pmtumin:                   uint32(d.Get("pmtumin").(int)),
		Pmtutimeout:               uint32(d.Get("pmtutimeout").(int)),
		Proxyprotocol:             d.Get("proxyprotocol").(string),
		Securecookie:              d.Get("securecookie").(string),
		Servicepathingressvlan:    uint32(d.Get("servicepathingressvlan").(int)),
		Tcpcip:                    d.Get("tcpcip").(string),
		Timezone:                  d.Get("timezone").(string),
		Useproxyport:              d.Get("useproxyport").(string),
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
	d.Set("name", data["name"])
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
