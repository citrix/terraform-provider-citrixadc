package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

type Lbparameter struct {
	Adccookieattributewarningmsg  string      `json:"adccookieattributewarningmsg,omitempty"`
	Allowboundsvcremoval          string      `json:"allowboundsvcremoval,omitempty"`
	Builtin                       interface{} `json:"builtin,omitempty"`
	Computedadccookieattribute    string      `json:"computedadccookieattribute,omitempty"`
	Consolidatedlconn             string      `json:"consolidatedlconn,omitempty"`
	Cookiepassphrase              string      `json:"cookiepassphrase,omitempty"`
	Dbsttl                        int         `json:"dbsttl,omitempty"`
	Dropmqttjumbomessage          string      `json:"dropmqttjumbomessage,omitempty"`
	Feature                       string      `json:"feature,omitempty"`
	Httponlycookieflag            string      `json:"httponlycookieflag,omitempty"`
	Literaladccookieattribute     string      `json:"literaladccookieattribute,omitempty"`
	Maxpipelinenat                int         `json:"maxpipelinenat,omitempty"`
	Monitorconnectionclose        string      `json:"monitorconnectionclose,omitempty"`
	Monitorskipmaxclient          string      `json:"monitorskipmaxclient,omitempty"`
	Preferdirectroute             string      `json:"preferdirectroute,omitempty"`
	Retainservicestate            string      `json:"retainservicestate,omitempty"`
	Sessionsthreshold             int         `json:"sessionsthreshold,omitempty"`
	Startuprrfactor               int         `json:"startuprrfactor,omitempty"`
	Storemqttclientidandusername  string      `json:"storemqttclientidandusername,omitempty"`
	Useencryptedpersistencecookie string      `json:"useencryptedpersistencecookie,omitempty"`
	Useportforhashlb              string      `json:"useportforhashlb,omitempty"`
	Usesecuredpersistencecookie   string      `json:"usesecuredpersistencecookie,omitempty"`
	Vserverspecificmac            string      `json:"vserverspecificmac,omitempty"`
	Lbhashalgorithm               string      `json:"lbhashalgorithm,omitempty"`
	Lbhashfingers                 int         `json:"lbhashfingers,omitempty"`
}

func resourceCitrixAdcLbparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbparameterFunc,
		Read:          readLbparameterFunc,
		Update:        updateLbparameterFunc,
		Delete:        deleteLbparameterFunc,
		Schema: map[string]*schema.Schema{
			"allowboundsvcremoval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"computedadccookieattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"consolidatedlconn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiepassphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbsttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dropmqttjumbomessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httponlycookieflag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"literaladccookieattribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxpipelinenat": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitorconnectionclose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitorskipmaxclient": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferdirectroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retainservicestate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"startuprrfactor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storemqttclientidandusername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionsthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"useencryptedpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useportforhashlb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usesecuredpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vserverspecificmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbhashalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lbhashfingers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}
func createLbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	lbparameterName := resource.PrefixedUniqueId("tf-lbparameter-")

	lbparameter := Lbparameter{
		Allowboundsvcremoval:          d.Get("allowboundsvcremoval").(string),
		Computedadccookieattribute:    d.Get("computedadccookieattribute").(string),
		Consolidatedlconn:             d.Get("consolidatedlconn").(string),
		Cookiepassphrase:              d.Get("cookiepassphrase").(string),
		Dbsttl:                        d.Get("dbsttl").(int),
		Dropmqttjumbomessage:          d.Get("dropmqttjumbomessage").(string),
		Httponlycookieflag:            d.Get("httponlycookieflag").(string),
		Literaladccookieattribute:     d.Get("literaladccookieattribute").(string),
		Maxpipelinenat:                d.Get("maxpipelinenat").(int),
		Monitorconnectionclose:        d.Get("monitorconnectionclose").(string),
		Monitorskipmaxclient:          d.Get("monitorskipmaxclient").(string),
		Preferdirectroute:             d.Get("preferdirectroute").(string),
		Retainservicestate:            d.Get("retainservicestate").(string),
		Startuprrfactor:               d.Get("startuprrfactor").(int),
		Storemqttclientidandusername:  d.Get("storemqttclientidandusername").(string),
		Sessionsthreshold:  		   d.Get("sessionsthreshold").(int),
		Useencryptedpersistencecookie: d.Get("useencryptedpersistencecookie").(string),
		Useportforhashlb:              d.Get("useportforhashlb").(string),
		Usesecuredpersistencecookie:   d.Get("usesecuredpersistencecookie").(string),
		Vserverspecificmac:            d.Get("vserverspecificmac").(string),
		Lbhashalgorithm:               d.Get("lbhashalgorithm").(string),
		Lbhashfingers:                 d.Get("lbhashfingers").(int),
	}

	err := client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		return err
	}

	d.SetId(lbparameterName)

	err = readLbparameterFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbparameter but we can't read it ?? %s", lbparameterName)
		return nil
	}
	return nil
}

func readLbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	lbparameterName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lbparameter state %s", lbparameterName)
	data, err := client.FindResource(service.Lbparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lbparameter state %s", lbparameterName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("allowboundsvcremoval", data["allowboundsvcremoval"])
	d.Set("computedadccookieattribute", data["computedadccookieattribute"])
	d.Set("consolidatedlconn", data["consolidatedlconn"])
	d.Set("cookiepassphrase", data["cookiepassphrase"])
	d.Set("dbsttl", data["dbsttl"])
	d.Set("dropmqttjumbomessage", data["dropmqttjumbomessage"])
	d.Set("httponlycookieflag", data["httponlycookieflag"])
	d.Set("literaladccookieattribute", data["literaladccookieattribute"])
	d.Set("maxpipelinenat", data["maxpipelinenat"])
	d.Set("monitorconnectionclose", data["monitorconnectionclose"])
	d.Set("monitorskipmaxclient", data["monitorskipmaxclient"])
	d.Set("preferdirectroute", data["preferdirectroute"])
	d.Set("retainservicestate", data["retainservicestate"])
	d.Set("startuprrfactor", data["startuprrfactor"])
	d.Set("storemqttclientidandusername", data["storemqttclientidandusername"])
	// d.Set("sessionsthreshold", data["sessionsthreshold"])
	d.Set("useencryptedpersistencecookie", data["useencryptedpersistencecookie"])
	d.Set("useportforhashlb", data["useportforhashlb"])
	d.Set("usesecuredpersistencecookie", data["usesecuredpersistencecookie"])
	d.Set("vserverspecificmac", data["vserverspecificmac"])
	d.Set("lbhashalgorithm", data["lbhashalgorithm"])
	d.Set("lbhashfingers", data["lbhashfingers"])

	return nil

}

func updateLbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	lbparameterName := d.Id()

	lbparameter := Lbparameter{}
	hasChange := false

	if d.HasChange("allowboundsvcremoval") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowboundsvcremoval has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Allowboundsvcremoval = d.Get("allowboundsvcremoval").(string)
		hasChange = true
	}
	if d.HasChange("computedadccookieattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Computedadccookieattribute has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Computedadccookieattribute = d.Get("computedadccookieattribute").(string)
		hasChange = true
	}
	if d.HasChange("consolidatedlconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Consolidatedlconn has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Consolidatedlconn = d.Get("consolidatedlconn").(string)
		hasChange = true
	}
	if d.HasChange("cookiepassphrase") {
		log.Printf("[DEBUG]  citrixadc-provider: Cookiepassphrase has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Cookiepassphrase = d.Get("cookiepassphrase").(string)
		hasChange = true
	}
	if d.HasChange("dbsttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Dbsttl has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Dbsttl = d.Get("dbsttl").(int)
		hasChange = true
	}
	if d.HasChange("dropmqttjumbomessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropmqttjumbomessage has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Dropmqttjumbomessage = d.Get("dropmqttjumbomessage").(string)
		hasChange = true
	}
	if d.HasChange("httponlycookieflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Httponlycookieflag has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Httponlycookieflag = d.Get("httponlycookieflag").(string)
		hasChange = true
	}
	if d.HasChange("literaladccookieattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Literaladccookieattribute has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Literaladccookieattribute = d.Get("literaladccookieattribute").(string)
		hasChange = true
	}
	if d.HasChange("maxpipelinenat") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpipelinenat has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Maxpipelinenat = d.Get("maxpipelinenat").(int)
		hasChange = true
	}
	if d.HasChange("monitorconnectionclose") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitorconnectionclose has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Monitorconnectionclose = d.Get("monitorconnectionclose").(string)
		hasChange = true
	}
	if d.HasChange("monitorskipmaxclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Monitorskipmaxclient has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Monitorskipmaxclient = d.Get("monitorskipmaxclient").(string)
		hasChange = true
	}
	if d.HasChange("preferdirectroute") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferdirectroute has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Preferdirectroute = d.Get("preferdirectroute").(string)
		hasChange = true
	}
	if d.HasChange("retainservicestate") {
		log.Printf("[DEBUG]  citrixadc-provider: Retainservicestate has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Retainservicestate = d.Get("retainservicestate").(string)
		hasChange = true
	}
	if d.HasChange("startuprrfactor") {
		log.Printf("[DEBUG]  citrixadc-provider: Startuprrfactor has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Startuprrfactor = d.Get("startuprrfactor").(int)
		hasChange = true
	}
	if d.HasChange("storemqttclientidandusername") {
		log.Printf("[DEBUG]  citrixadc-provider: Storemqttclientidandusername has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Storemqttclientidandusername = d.Get("storemqttclientidandusername").(string)
		hasChange = true
	}
	if d.HasChange("sessionsthreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Storemqttclientidandusername has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Sessionsthreshold = d.Get("sessionsthreshold").(int)
		hasChange = true
	}
	if d.HasChange("useencryptedpersistencecookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Useencryptedpersistencecookie has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Useencryptedpersistencecookie = d.Get("useencryptedpersistencecookie").(string)
		hasChange = true
	}
	if d.HasChange("useportforhashlb") {
		log.Printf("[DEBUG]  citrixadc-provider: Useportforhashlb has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Useportforhashlb = d.Get("useportforhashlb").(string)
		hasChange = true
	}
	if d.HasChange("usesecuredpersistencecookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Usesecuredpersistencecookie has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Usesecuredpersistencecookie = d.Get("usesecuredpersistencecookie").(string)
		hasChange = true
	}
	if d.HasChange("vserverspecificmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Vserverspecificmac has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Vserverspecificmac = d.Get("vserverspecificmac").(string)
		hasChange = true
	}
	if d.HasChange("lbhashalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbhashalgorithm has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Lbhashalgorithm = d.Get("lbhashalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("lbhashfingers") {
		log.Printf("[DEBUG]  citrixadc-provider: Lbhashfingers has changed for lbparameter %s, starting update", lbparameterName)
		lbparameter.Lbhashfingers = d.Get("lbhashfingers").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
		if err != nil {
			return fmt.Errorf("Error updating lbparameter %s. %s", lbparameterName, err.Error())
		}
	}
	return readLbparameterFunc(d, meta)
}

func deleteLbparameterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbparameterFunc")

	d.SetId("")

	return nil
}
