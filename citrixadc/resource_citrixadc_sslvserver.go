package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslvserverFunc,
		ReadContext:   readSslvserverFunc,
		UpdateContext: updateSslvserverFunc,
		DeleteContext: deleteSslvserverFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
			"cleartextport": {
				Type:     schema.TypeInt,
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
			"dhekeyexchangewithpsk": {
				Type:     schema.TypeString,
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
			"hsts": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includesubdomains": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocspstapling": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preload": {
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
			"tls13sessionticketsperauthcontext": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vservername": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zerorttearlydata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	var sslvserverName string
	if v, ok := d.GetOk("vservername"); ok {
		sslvserverName = v.(string)
	} else {
		log.Printf("[ERROR] No pre-existing ssl vserver  was found")
		return nil
	}
	sslvserver := ssl.Sslvserver{
		Cipherredirect:                    d.Get("cipherredirect").(string),
		Cipherurl:                         d.Get("cipherurl").(string),
		Cleartextport:                     d.Get("cleartextport").(int),
		Clientauth:                        d.Get("clientauth").(string),
		Clientcert:                        d.Get("clientcert").(string),
		Dh:                                d.Get("dh").(string),
		Dhcount:                           d.Get("dhcount").(int),
		Dhekeyexchangewithpsk:             d.Get("dhekeyexchangewithpsk").(string),
		Dhfile:                            d.Get("dhfile").(string),
		Dhkeyexpsizelimit:                 d.Get("dhkeyexpsizelimit").(string),
		Dtls1:                             d.Get("dtls1").(string),
		Dtls12:                            d.Get("dtls12").(string),
		Dtlsprofilename:                   d.Get("dtlsprofilename").(string),
		Ersa:                              d.Get("ersa").(string),
		Ersacount:                         d.Get("ersacount").(int),
		Hsts:                              d.Get("hsts").(string),
		Includesubdomains:                 d.Get("includesubdomains").(string),
		Maxage:                            d.Get("maxage").(int),
		Ocspstapling:                      d.Get("ocspstapling").(string),
		Preload:                           d.Get("preload").(string),
		Pushenctrigger:                    d.Get("pushenctrigger").(string),
		Redirectportrewrite:               d.Get("redirectportrewrite").(string),
		Sendclosenotify:                   d.Get("sendclosenotify").(string),
		Sessreuse:                         d.Get("sessreuse").(string),
		Sesstimeout:                       d.Get("sesstimeout").(int),
		Snienable:                         d.Get("snienable").(string),
		Ssl2:                              d.Get("ssl2").(string),
		Ssl3:                              d.Get("ssl3").(string),
		Sslprofile:                        d.Get("sslprofile").(string),
		Sslredirect:                       d.Get("sslredirect").(string),
		Sslv2redirect:                     d.Get("sslv2redirect").(string),
		Sslv2url:                          d.Get("sslv2url").(string),
		Strictsigdigestcheck:              d.Get("strictsigdigestcheck").(string),
		Tls1:                              d.Get("tls1").(string),
		Tls11:                             d.Get("tls11").(string),
		Tls12:                             d.Get("tls12").(string),
		Tls13:                             d.Get("tls13").(string),
		Tls13sessionticketsperauthcontext: d.Get("tls13sessionticketsperauthcontext").(int),
		Vservername:                       d.Get("vservername").(string),
		Zerorttearlydata:                  d.Get("zerorttearlydata").(string),
	}

	_, err := client.UpdateResource(service.Sslvserver.Type(), sslvserverName, &sslvserver)
	if err != nil {
		return diag.Errorf("Error updating sslvserver %v: %v", sslvserverName, err)
	}

	d.SetId(sslvserverName)

	return readSslvserverFunc(ctx, d, meta)
}

func readSslvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	sslvserverName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslvserver state %s", sslvserverName)
	data, err := client.FindResource(service.Sslvserver.Type(), sslvserverName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslvserver state %s", sslvserverName)
		d.SetId("")
		return nil
	}

	sslvserverAttributes := [40]string{
		"cipherredirect",
		"cipherurl",
		// "cleartextport",
		"clientauth",
		"clientcert",
		"dh",
		"dhcount",
		"dhekeyexchangewithpsk",
		"dhfile",
		"dhkeyexpsizelimit",
		"dtls1",
		"dtls12",
		"dtlsprofilename",
		"ersa",
		"ersacount",
		"hsts",
		"includesubdomains",
		"maxage",
		"ocspstapling",
		"preload",
		"pushenctrigger",
		"redirectportrewrite",
		"sendclosenotify",
		"sessreuse",
		"sesstimeout",
		"snienable",
		"ssl2",
		"ssl3",
		"sslprofile",
		"sslredirect",
		"sslv2redirect",
		"sslv2url",
		"strictsigdigestcheck",
		"tls1",
		"tls11",
		"tls12",
		"tls13",
		"tls13sessionticketsperauthcontext",
		"vservername",
		"zerorttearlydata",
	}

	for _, val := range sslvserverAttributes {
		if _, exists := data[val]; exists {
			if data[val] != "" || data[val] != nil {
				d.Set(val, data[val])
			}
		}
	}

	return nil

}

func updateSslvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslvserverFunc")
	client := meta.(*NetScalerNitroClient).client
	sslvserverName := d.Get("vservername").(string)

	sslvserver := ssl.Sslvserver{
		Vservername: d.Get("vservername").(string),
	}
	hasChange := false
	if d.HasChange("cipherredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherredirect has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Cipherredirect = d.Get("cipherredirect").(string)
		hasChange = true
	}
	if d.HasChange("cipherurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherurl has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Cipherurl = d.Get("cipherurl").(string)
		hasChange = true
	}
	if d.HasChange("cleartextport") {
		log.Printf("[DEBUG]  citrixadc-provider: Cleartextport has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Cleartextport = d.Get("cleartextport").(int)
		hasChange = true
	}
	if d.HasChange("clientauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientauth has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Clientauth = d.Get("clientauth").(string)
		hasChange = true
	}
	if d.HasChange("clientcert") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcert has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Clientcert = d.Get("clientcert").(string)
		hasChange = true
	}
	if d.HasChange("dh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dh has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dh = d.Get("dh").(string)
		hasChange = true
	}
	if d.HasChange("dhcount") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhcount has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dhcount = d.Get("dhcount").(int)
		hasChange = true
	}
	if d.HasChange("dhekeyexchangewithpsk") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhekeyexchangewithpsk has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dhekeyexchangewithpsk = d.Get("dhekeyexchangewithpsk").(string)
		hasChange = true
	}
	if d.HasChange("dhfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhfile has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dhfile = d.Get("dhfile").(string)
		hasChange = true
	}
	if d.HasChange("dhkeyexpsizelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhkeyexpsizelimit has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dhkeyexpsizelimit = d.Get("dhkeyexpsizelimit").(string)
		hasChange = true
	}
	if d.HasChange("dtls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtls1 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dtls1 = d.Get("dtls1").(string)
		hasChange = true
	}
	if d.HasChange("dtls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtls12 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dtls12 = d.Get("dtls12").(string)
		hasChange = true
	}
	if d.HasChange("dtlsprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Dtlsprofilename has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Dtlsprofilename = d.Get("dtlsprofilename").(string)
		hasChange = true
	}
	if d.HasChange("ersa") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersa has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Ersa = d.Get("ersa").(string)
		hasChange = true
	}
	if d.HasChange("ersacount") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersacount has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Ersacount = d.Get("ersacount").(int)
		hasChange = true
	}
	if d.HasChange("hsts") {
		log.Printf("[DEBUG]  citrixadc-provider: Hsts has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Hsts = d.Get("hsts").(string)
		hasChange = true
	}
	if d.HasChange("includesubdomains") {
		log.Printf("[DEBUG]  citrixadc-provider: Includesubdomains has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Includesubdomains = d.Get("includesubdomains").(string)
		hasChange = true
	}
	if d.HasChange("maxage") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxage has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Maxage = d.Get("maxage").(int)
		hasChange = true
	}
	if d.HasChange("ocspstapling") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspstapling has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Ocspstapling = d.Get("ocspstapling").(string)
		hasChange = true
	}
	if d.HasChange("preload") {
		log.Printf("[DEBUG]  citrixadc-provider: Preload has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Preload = d.Get("preload").(string)
		hasChange = true
	}
	if d.HasChange("pushenctrigger") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctrigger has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Pushenctrigger = d.Get("pushenctrigger").(string)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectportrewrite has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("sendclosenotify") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendclosenotify has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sendclosenotify = d.Get("sendclosenotify").(string)
		hasChange = true
	}
	if d.HasChange("sessreuse") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessreuse has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sessreuse = d.Get("sessreuse").(string)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("snienable") {
		log.Printf("[DEBUG]  citrixadc-provider: Snienable has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Snienable = d.Get("snienable").(string)
		hasChange = true
	}
	if d.HasChange("ssl2") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl2 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Ssl2 = d.Get("ssl2").(string)
		hasChange = true
	}
	if d.HasChange("ssl3") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl3 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Ssl3 = d.Get("ssl3").(string)
		hasChange = true
	}
	if d.HasChange("sslprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslprofile has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sslprofile = d.Get("sslprofile").(string)
		hasChange = true
	}
	if d.HasChange("sslredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslredirect has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sslredirect = d.Get("sslredirect").(string)
		hasChange = true
	}
	if d.HasChange("sslv2redirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslv2redirect has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sslv2redirect = d.Get("sslv2redirect").(string)
		hasChange = true
	}
	if d.HasChange("sslv2url") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslv2url has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Sslv2url = d.Get("sslv2url").(string)
		hasChange = true
	}
	if d.HasChange("strictsigdigestcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictsigdigestcheck has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Strictsigdigestcheck = d.Get("strictsigdigestcheck").(string)
		hasChange = true
	}
	if d.HasChange("tls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls1 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Tls1 = d.Get("tls1").(string)
		hasChange = true
	}
	if d.HasChange("tls11") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls11 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Tls11 = d.Get("tls11").(string)
		hasChange = true
	}
	if d.HasChange("tls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls12 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Tls12 = d.Get("tls12").(string)
		hasChange = true
	}
	if d.HasChange("tls13") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13 has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Tls13 = d.Get("tls13").(string)
		hasChange = true
	}
	if d.HasChange("tls13sessionticketsperauthcontext") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13sessionticketsperauthcontext has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Tls13sessionticketsperauthcontext = d.Get("tls13sessionticketsperauthcontext").(int)
		hasChange = true
	}
	if d.HasChange("vservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Vservername has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Vservername = d.Get("vservername").(string)
		hasChange = true
	}
	if d.HasChange("zerorttearlydata") {
		log.Printf("[DEBUG]  citrixadc-provider: Zerorttearlydata has changed for sslvserver %s, starting update", sslvserverName)
		sslvserver.Zerorttearlydata = d.Get("zerorttearlydata").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslvserver.Type(), sslvserverName, &sslvserver)
		if err != nil {
			return diag.Errorf("Error updating sslvserver %s", sslvserverName)
		}
	}
	return readSslvserverFunc(ctx, d, meta)
}

func deleteSslvserverFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslvserverFunc")
	// sslvserver does not have DELETE operation, but this function is required to set the ID to ""

	// this unsets the sslprofile for the sslvserver, making it no issues to delete the sslprofile
	if _, ok := d.GetOk("sslprofile"); !ok {

		client := meta.(*NetScalerNitroClient).client
		sslvserverName := d.Get("vservername").(string)

		sslvserverunset := make(map[string]interface{})
		sslvserverunset["sslprofile"] = true
		sslvserverunset["vservername"] = sslvserverName

		err := client.ActOnResource("sslvserver", sslvserverunset, "unset")
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")

	return nil
}
