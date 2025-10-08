package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNsparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNsparamFunc,
		ReadContext:   readNsparamFunc,
		DeleteContext: deleteNsparamFunc,
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
			"icaports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"secureicaports": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ipttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNsparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	if listVal, ok := d.Get("secureicaports").([]interface{}); ok {
		nsparam["secureicaports"] = toStringList(listVal)
	}
	if listVal, ok := d.Get("icaports").([]interface{}); ok {
		nsparam["icaports"] = toStringList(listVal)
	}
	if data, ok := d.GetOk("ipttl"); ok {
		nsparam["ipttl"] = data.(int)
	}

	err := client.UpdateUnnamedResource(service.Nsparam.Type(), &nsparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nsparamId)

	return readNsparamFunc(ctx, d, meta)
}

func readNsparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	setToInt("exclusivequotamaxclient", d, data["exclusivequotamaxclient"])
	setToInt("exclusivequotaspillover", d, data["exclusivequotaspillover"])
	d.Set("ftpportrange", data["ftpportrange"])
	setToInt("grantquotamaxclient", d, data["grantquotamaxclient"])
	setToInt("grantquotaspillover", d, data["grantquotaspillover"])
	d.Set("internaluserlogin", data["internaluserlogin"])
	setToInt("maxconn", d, data["maxconn"])
	setToInt("maxreq", d, data["maxreq"])
	setToInt("mgmthttpport", d, data["mgmthttpport"])
	setToInt("mgmthttpsport", d, data["mgmthttpsport"])
	setToInt("pmtumin", d, data["pmtumin"])
	setToInt("pmtutimeout", d, data["pmtutimeout"])
	d.Set("proxyprotocol", data["proxyprotocol"])
	d.Set("securecookie", data["securecookie"])
	setToInt("servicepathingressvlan", d, data["servicepathingressvlan"])
	d.Set("tcpcip", data["tcpcip"])
	// d.Set("timezone", data["timezone"]) // This is received as different value from the NetScaler
	d.Set("useproxyport", data["useproxyport"])
	if val, ok := data["secureicaports"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("secureicaports", toStringList(list))
		}
	} else {
		d.Set("secureicaports", nil)
	}
	if val, ok := data["icaports"]; ok {
		if list, ok := val.([]interface{}); ok {
			d.Set("icaports", toStringList(list))
		}
	} else {
		d.Set("icaports", nil)
	}
	setToInt("ipttl", d, data["ipttl"])

	return nil

}

func deleteNsparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsparamFunc")

	d.SetId("")

	return nil
}
