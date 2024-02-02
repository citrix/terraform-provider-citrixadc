package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSslservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslserviceFunc,
		Read:          readSslserviceFunc,
		Update:        updateSslserviceFunc,
		Delete:        deleteSslserviceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cipherredirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipherurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientcert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhfile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhkeyexpsizelimit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls12": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtlsprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ersa": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ersacount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocspstapling": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushenctrigger": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectportrewrite": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendclosenotify": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverauth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sessreuse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"snienable": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslprofile": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslredirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslv2redirect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslv2url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strictsigdigestcheck": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls11": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls12": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls13": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslserviceName = d.Get("servicename").(string)

	sslservice := ssl.Sslservice{
		Cipherredirect:       d.Get("cipherredirect").(string),
		Cipherurl:            d.Get("cipherurl").(string),
		Clientauth:           d.Get("clientauth").(string),
		Clientcert:           d.Get("clientcert").(string),
		Commonname:           d.Get("commonname").(string),
		Dh:                   d.Get("dh").(string),
		Dhcount:              d.Get("dhcount").(int),
		Dhfile:               d.Get("dhfile").(string),
		Dhkeyexpsizelimit:    d.Get("dhkeyexpsizelimit").(string),
		Dtls1:                d.Get("dtls1").(string),
		Dtls12:               d.Get("dtls12").(string),
		Dtlsprofilename:      d.Get("dtlsprofilename").(string),
		Ersa:                 d.Get("ersa").(string),
		Ersacount:            d.Get("ersacount").(int),
		Ocspstapling:         d.Get("ocspstapling").(string),
		Pushenctrigger:       d.Get("pushenctrigger").(string),
		Redirectportrewrite:  d.Get("redirectportrewrite").(string),
		Sendclosenotify:      d.Get("sendclosenotify").(string),
		Serverauth:           d.Get("serverauth").(string),
		Servicename:          sslserviceName,
		Sessreuse:            d.Get("sessreuse").(string),
		Sesstimeout:          d.Get("sesstimeout").(int),
		Snienable:            d.Get("snienable").(string),
		Ssl2:                 d.Get("ssl2").(string),
		Ssl3:                 d.Get("ssl3").(string),
		Sslprofile:           d.Get("sslprofile").(string),
		Sslredirect:          d.Get("sslredirect").(string),
		Sslv2redirect:        d.Get("sslv2redirect").(string),
		Sslv2url:             d.Get("sslv2url").(string),
		Strictsigdigestcheck: d.Get("strictsigdigestcheck").(string),
		Tls1:                 d.Get("tls1").(string),
		Tls11:                d.Get("tls11").(string),
		Tls12:                d.Get("tls12").(string),
		Tls13:                d.Get("tls13").(string),
	}

	err := client.UpdateUnnamedResource(service.Sslservice.Type(), &sslservice)
	if err != nil {
		return err
	}

	d.SetId(sslserviceName)

	err = readSslserviceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslservice but we can't read it ?? %s", sslserviceName)
		return nil
	}
	return nil
}

func readSslserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	sslserviceName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslservice state %s", sslserviceName)
	data, err := client.FindResource(service.Sslservice.Type(), sslserviceName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslservice state %s", sslserviceName)
		d.SetId("")
		return nil
	}
	d.Set("servicename", data["servicename"])
	// d.Set("cipherredirect", data["cipherredirect"])
	// d.Set("cipherurl", data["cipherurl"])
	// d.Set("clientauth", data["clientauth"])
	// d.Set("clientcert", data["clientcert"])
	// d.Set("commonname", data["commonname"])
	// d.Set("dh", data["dh"])
	// d.Set("dhcount", data["dhcount"])
	// d.Set("dhfile", data["dhfile"])
	// d.Set("dhkeyexpsizelimit", data["dhkeyexpsizelimit"])
	// d.Set("dtls1", data["dtls1"])
	// d.Set("dtls12", data["dtls12"])
	// d.Set("dtlsprofilename", data["dtlsprofilename"])
	// d.Set("ersa", data["ersa"])
	// d.Set("ersacount", data["ersacount"])
	// d.Set("ocspstapling", data["ocspstapling"])
	// d.Set("pushenctrigger", data["pushenctrigger"])
	// d.Set("redirectportrewrite", data["redirectportrewrite"])
	// d.Set("sendclosenotify", data["sendclosenotify"])
	// d.Set("serverauth", data["serverauth"])
	// d.Set("servicename", data["servicename"])
	// d.Set("sessreuse", data["sessreuse"])
	// d.Set("sesstimeout", data["sesstimeout"])
	// d.Set("snienable", data["snienable"])
	// d.Set("ssl2", data["ssl2"])
	// d.Set("ssl3", data["ssl3"])
	d.Set("sslprofile", data["sslprofile"])
	// d.Set("sslredirect", data["sslredirect"])
	// d.Set("sslv2redirect", data["sslv2redirect"])
	// d.Set("sslv2url", data["sslv2url"])
	// d.Set("strictsigdigestcheck", data["strictsigdigestcheck"])
	// d.Set("tls1", data["tls1"])
	// d.Set("tls11", data["tls11"])
	// d.Set("tls12", data["tls12"])
	// d.Set("tls13", data["tls13"])

	return nil

}

func updateSslserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	sslserviceName := d.Get("servicename").(string)

	sslservice := ssl.Sslservice{
		Servicename: d.Get("servicename").(string),
	}
	hasChange := false
	if d.HasChange("cipherredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherredirect has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Cipherredirect = d.Get("cipherredirect").(string)
		hasChange = true
	}
	if d.HasChange("cipherurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherurl has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Cipherurl = d.Get("cipherurl").(string)
		hasChange = true
	}
	if d.HasChange("clientauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientauth has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Clientauth = d.Get("clientauth").(string)
		hasChange = true
	}
	if d.HasChange("clientcert") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcert has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Clientcert = d.Get("clientcert").(string)
		hasChange = true
	}
	if d.HasChange("commonname") {
		log.Printf("[DEBUG]  citrixadc-provider: Commonname has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Commonname = d.Get("commonname").(string)
		hasChange = true
	}
	if d.HasChange("dh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dh has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dh = d.Get("dh").(string)
		hasChange = true
	}
	if d.HasChange("dhcount") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhcount has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dhcount = d.Get("dhcount").(int)
		hasChange = true
	}
	if d.HasChange("dhfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhfile has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dhfile = d.Get("dhfile").(string)
		hasChange = true
	}
	if d.HasChange("dhkeyexpsizelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhkeyexpsizelimit has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dhkeyexpsizelimit = d.Get("dhkeyexpsizelimit").(string)
		hasChange = true
	}
	if d.HasChange("dtls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtls1 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dtls1 = d.Get("dtls1").(string)
		hasChange = true
	}
	if d.HasChange("dtls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtls12 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dtls12 = d.Get("dtls12").(string)
		hasChange = true
	}
	if d.HasChange("dtlsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtlsprofilename has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Dtlsprofilename = d.Get("dtlsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("ersa") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersa has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Ersa = d.Get("ersa").(string)
		hasChange = true
	}
	if d.HasChange("ersacount") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersacount has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Ersacount = d.Get("ersacount").(int)
		hasChange = true
	}
	if d.HasChange("ocspstapling") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspstapling has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Ocspstapling = d.Get("ocspstapling").(string)
		hasChange = true
	}
	if d.HasChange("pushenctrigger") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctrigger has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Pushenctrigger = d.Get("pushenctrigger").(string)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectportrewrite has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("sendclosenotify") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendclosenotify has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sendclosenotify = d.Get("sendclosenotify").(string)
		hasChange = true
	}
	if d.HasChange("serverauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverauth has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Serverauth = d.Get("serverauth").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  citrixadc-provider: Servicename has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("sessreuse") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessreuse has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sessreuse = d.Get("sessreuse").(string)
		hasChange = true
	}
	//sessreuse pre-requisite for sesstimeout
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sesstimeout = d.Get("sesstimeout").(int)
		sslservice.Sessreuse = d.Get("sessreuse").(string)
		hasChange = true
	}
	if d.HasChange("snienable") {
		log.Printf("[DEBUG]  citrixadc-provider: Snienable has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Snienable = d.Get("snienable").(string)
		hasChange = true
	}
	if d.HasChange("ssl2") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl2 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Ssl2 = d.Get("ssl2").(string)
		hasChange = true
	}
	if d.HasChange("ssl3") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl3 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Ssl3 = d.Get("ssl3").(string)
		hasChange = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslprofile has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sslprofile = d.Get("sslprofile").(string)
		hasChange = true
	}
	if d.HasChange("sslredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslredirect has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sslredirect = d.Get("sslredirect").(string)
		hasChange = true
	}
	if d.HasChange("sslv2redirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslv2redirect has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sslv2redirect = d.Get("sslv2redirect").(string)
		hasChange = true
	}
	if d.HasChange("sslv2url") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslv2url has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Sslv2url = d.Get("sslv2url").(string)
		hasChange = true
	}
	if d.HasChange("strictsigdigestcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictsigdigestcheck has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Strictsigdigestcheck = d.Get("strictsigdigestcheck").(string)
		hasChange = true
	}
	if d.HasChange("tls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls1 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Tls1 = d.Get("tls1").(string)
		hasChange = true
	}
	if d.HasChange("tls11") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls11 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Tls11 = d.Get("tls11").(string)
		hasChange = true
	}
	if d.HasChange("tls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls12 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Tls12 = d.Get("tls12").(string)
		hasChange = true
	}
	if d.HasChange("tls13") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13 has changed for sslservice %s, starting update", sslserviceName)
		sslservice.Tls13 = d.Get("tls13").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslservice.Type(), sslserviceName, &sslservice)
		if err != nil {
			return fmt.Errorf("Error updating sslservice %s", sslserviceName)
		}
	}
	return readSslserviceFunc(d, meta)
}

func deleteSslserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslserviceFunc")

	d.SetId("")

	return nil
}
