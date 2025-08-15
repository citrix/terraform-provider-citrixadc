package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNshttpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNshttpprofileFunc,
		Read:          readNshttpprofileFunc,
		Update:        updateNshttpprofileFunc,
		Delete:        deleteNshttpprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"adpttimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altsvc": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altsvcvalue": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apdexcltresptimethreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientiphdrexpr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmponpush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conmultiplex": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropextracrlf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropextradata": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropinvalreqs": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowonlywordcharactersandhyphen": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"grpcholdlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"grpcholdtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"grpclengthdelimitation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2altsvcframe": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2direct": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2headertablesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2initialconnwindowsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2initialwindowsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxconcurrentstreams": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxemptyframespermin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxframesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxheaderlistsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxpingframespermin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxresetframespermin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxsettingsframespermin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2minseverconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2strictcipher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http3maxheaderblockedstreams": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http3maxheaderfieldsectionsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http3maxheadertablesize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"incomphdrdelay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"markconnreqinval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markhttp09inval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markhttpheaderextrawserror": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markrfc7230noncompliantinval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marktracereqinval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxheaderlen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreq": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreusepool": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minreusepool": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"persistentetag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reqtimeoutaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reusepooltimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtsptunnel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spdy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weblog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websocket": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshttpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Get("name").(string)

	nshttpprofile := ns.Nshttpprofile{
		Adpttimeout:                      d.Get("adpttimeout").(string),
		Allowonlywordcharactersandhyphen: d.Get("allowonlywordcharactersandhyphen").(string),
		Altsvc:                           d.Get("altsvc").(string),
		Altsvcvalue:                      d.Get("altsvcvalue").(string),
		Apdexcltresptimethreshold:        d.Get("apdexcltresptimethreshold").(int),
		Clientiphdrexpr:                  d.Get("clientiphdrexpr").(string),
		Cmponpush:                        d.Get("cmponpush").(string),
		Conmultiplex:                     d.Get("conmultiplex").(string),
		Dropextracrlf:                    d.Get("dropextracrlf").(string),
		Dropextradata:                    d.Get("dropextradata").(string),
		Dropinvalreqs:                    d.Get("dropinvalreqs").(string),
		Grpcholdlimit:                    d.Get("grpcholdlimit").(int),
		Grpcholdtimeout:                  d.Get("grpcholdtimeout").(int),
		Grpclengthdelimitation:           d.Get("grpclengthdelimitation").(string),
		Http2:                            d.Get("http2").(string),
		Http2altsvcframe:                 d.Get("http2altsvcframe").(string),
		Http2direct:                      d.Get("http2direct").(string),
		Http2headertablesize:             d.Get("http2headertablesize").(int),
		Http2initialconnwindowsize:       d.Get("http2initialconnwindowsize").(int),
		Http2initialwindowsize:           d.Get("http2initialwindowsize").(int),
		Http2maxconcurrentstreams:        d.Get("http2maxconcurrentstreams").(int),
		Http2maxemptyframespermin:        d.Get("http2maxemptyframespermin").(int),
		Http2maxframesize:                d.Get("http2maxframesize").(int),
		Http2maxheaderlistsize:           d.Get("http2maxheaderlistsize").(int),
		Http2maxpingframespermin:         d.Get("http2maxpingframespermin").(int),
		Http2maxresetframespermin:        d.Get("http2maxresetframespermin").(int),
		Http2maxsettingsframespermin:     d.Get("http2maxsettingsframespermin").(int),
		Http2minseverconn:                d.Get("http2minseverconn").(int),
		Http2strictcipher:                d.Get("http2strictcipher").(string),
		Http3:                            d.Get("http3").(string),
		Http3maxheaderblockedstreams:     d.Get("http3maxheaderblockedstreams").(int),
		Http3maxheaderfieldsectionsize:   d.Get("http3maxheaderfieldsectionsize").(int),
		Http3maxheadertablesize:          d.Get("http3maxheadertablesize").(int),
		Incomphdrdelay:                   d.Get("incomphdrdelay").(int),
		Markconnreqinval:                 d.Get("markconnreqinval").(string),
		Markhttp09inval:                  d.Get("markhttp09inval").(string),
		Markhttpheaderextrawserror:       d.Get("markhttpheaderextrawserror").(string),
		Markrfc7230noncompliantinval:     d.Get("markrfc7230noncompliantinval").(string),
		Marktracereqinval:                d.Get("marktracereqinval").(string),
		Maxheaderlen:                     d.Get("maxheaderlen").(int),
		Maxreq:                           d.Get("maxreq").(int),
		Maxreusepool:                     d.Get("maxreusepool").(int),
		Minreusepool:                     d.Get("minreusepool").(int),
		Name:                             d.Get("name").(string),
		Persistentetag:                   d.Get("persistentetag").(string),
		Reqtimeout:                       d.Get("reqtimeout").(int),
		Reqtimeoutaction:                 d.Get("reqtimeoutaction").(string),
		Reusepooltimeout:                 d.Get("reusepooltimeout").(int),
		Rtsptunnel:                       d.Get("rtsptunnel").(string),
		Spdy:                             d.Get("spdy").(string),
		Weblog:                           d.Get("weblog").(string),
		Websocket:                        d.Get("websocket").(string),
	}

	_, err := client.AddResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		return err
	}

	d.SetId(nshttpprofileName)

	err = readNshttpprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nshttpprofile but we can't read it ?? %s", nshttpprofileName)
		return nil
	}
	return nil
}

func readNshttpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nshttpprofile state %s", nshttpprofileName)
	data, err := client.FindResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshttpprofile state %s", nshttpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("adpttimeout", data["adpttimeout"])
	d.Set("allowonlywordcharactersandhyphen", data["allowonlywordcharactersandhyphen"])
	d.Set("altsvc", data["altsvc"])
	d.Set("altsvcvalue", data["altsvcvalue"])
	d.Set("apdexcltresptimethreshold", data["apdexcltresptimethreshold"])
	d.Set("clientiphdrexpr", data["clientiphdrexpr"])
	d.Set("cmponpush", data["cmponpush"])
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("dropextracrlf", data["dropextracrlf"])
	d.Set("dropextradata", data["dropextradata"])
	d.Set("dropinvalreqs", data["dropinvalreqs"])
	d.Set("grpcholdlimit", data["grpcholdlimit"])
	d.Set("grpcholdtimeout", data["grpcholdtimeout"])
	d.Set("grpclengthdelimitation", data["grpclengthdelimitation"])
	d.Set("http2", data["http2"])
	d.Set("http2altsvcframe", data["http2altsvcframe"])
	d.Set("http2direct", data["http2direct"])
	d.Set("http2headertablesize", data["http2headertablesize"])
	d.Set("http2initialconnwindowsize", data["http2initialconnwindowsize"])
	d.Set("http2initialwindowsize", data["http2initialwindowsize"])
	d.Set("http2maxconcurrentstreams", data["http2maxconcurrentstreams"])
	d.Set("http2maxemptyframespermin", data["http2maxemptyframespermin"])
	d.Set("http2maxframesize", data["http2maxframesize"])
	d.Set("http2maxheaderlistsize", data["http2maxheaderlistsize"])
	d.Set("http2maxpingframespermin", data["http2maxpingframespermin"])
	d.Set("http2maxresetframespermin", data["http2maxresetframespermin"])
	d.Set("http2maxsettingsframespermin", data["http2maxsettingsframespermin"])
	d.Set("http2minseverconn", data["http2minseverconn"])
	d.Set("http2strictcipher", data["http2strictcipher"])
	d.Set("http3", data["http3"])
	d.Set("http3maxheaderblockedstreams", data["http3maxheaderblockedstreams"])
	d.Set("http3maxheaderfieldsectionsize", data["http3maxheaderfieldsectionsize"])
	d.Set("http3maxheadertablesize", data["http3maxheadertablesize"])
	d.Set("incomphdrdelay", data["incomphdrdelay"])
	d.Set("markconnreqinval", data["markconnreqinval"])
	d.Set("markhttp09inval", data["markhttp09inval"])
	d.Set("markhttpheaderextrawserror", data["markhttpheaderextrawserror"])
	d.Set("markrfc7230noncompliantinval", data["markrfc7230noncompliantinval"])
	d.Set("marktracereqinval", data["marktracereqinval"])
	d.Set("maxheaderlen", data["maxheaderlen"])
	d.Set("maxreq", data["maxreq"])
	d.Set("maxreusepool", data["maxreusepool"])
	d.Set("minreusepool", data["minreusepool"])
	d.Set("name", data["name"])
	d.Set("persistentetag", data["persistentetag"])
	d.Set("reqtimeout", data["reqtimeout"])
	d.Set("reqtimeoutaction", data["reqtimeoutaction"])
	d.Set("reusepooltimeout", data["reusepooltimeout"])
	d.Set("rtsptunnel", data["rtsptunnel"])
	d.Set("spdy", data["spdy"])
	d.Set("weblog", data["weblog"])
	d.Set("websocket", data["websocket"])

	return nil

}

func updateNshttpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Get("name").(string)

	nshttpprofile := ns.Nshttpprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("adpttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Adpttimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Adpttimeout = d.Get("adpttimeout").(string)
		hasChange = true
	}
	if d.HasChange("allowonlywordcharactersandhyphen") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowonlywordcharactersandhyphen has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Allowonlywordcharactersandhyphen = d.Get("allowonlywordcharactersandhyphen").(string)
		hasChange = true
	}
	if d.HasChange("altsvc") {
		log.Printf("[DEBUG]  citrixadc-provider: Altsvc has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Altsvc = d.Get("altsvc").(string)
		hasChange = true
	}
	if d.HasChange("altsvcvalue") {
		log.Printf("[DEBUG]  citrixadc-provider: Altsvcvalue has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Altsvcvalue = d.Get("altsvcvalue").(string)
		hasChange = true
	}
	if d.HasChange("apdexcltresptimethreshold") {
		log.Printf("[DEBUG]  citrixadc-provider: Apdexcltresptimethreshold has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Apdexcltresptimethreshold = d.Get("apdexcltresptimethreshold").(int)
		hasChange = true
	}
	if d.HasChange("clientiphdrexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientiphdrexpr has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Clientiphdrexpr = d.Get("clientiphdrexpr").(string)
		hasChange = true
	}
	if d.HasChange("cmponpush") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmponpush has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Cmponpush = d.Get("cmponpush").(string)
		hasChange = true
	}
	if d.HasChange("conmultiplex") {
		log.Printf("[DEBUG]  citrixadc-provider: Conmultiplex has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Conmultiplex = d.Get("conmultiplex").(string)
		hasChange = true
	}
	if d.HasChange("dropextracrlf") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropextracrlf has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Dropextracrlf = d.Get("dropextracrlf").(string)
		hasChange = true
	}
	if d.HasChange("dropextradata") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropextradata has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Dropextradata = d.Get("dropextradata").(string)
		hasChange = true
	}
	if d.HasChange("dropinvalreqs") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropinvalreqs has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Dropinvalreqs = d.Get("dropinvalreqs").(string)
		hasChange = true
	}
	if d.HasChange("grpcholdlimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Grpcholdlimit has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Grpcholdlimit = d.Get("grpcholdlimit").(int)
		hasChange = true
	}
	if d.HasChange("grpcholdtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Grpcholdtimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Grpcholdtimeout = d.Get("grpcholdtimeout").(int)
		hasChange = true
	}
	if d.HasChange("grpclengthdelimitation") {
		log.Printf("[DEBUG]  citrixadc-provider: Grpclengthdelimitation has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Grpclengthdelimitation = d.Get("grpclengthdelimitation").(string)
		hasChange = true
	}
	if d.HasChange("http2") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2 has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2 = d.Get("http2").(string)
		hasChange = true
	}
	if d.HasChange("http2altsvcframe") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2altsvcframe has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2altsvcframe = d.Get("http2altsvcframe").(string)
		hasChange = true
	}
	if d.HasChange("http2direct") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2direct has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2direct = d.Get("http2direct").(string)
		hasChange = true
	}
	if d.HasChange("http2headertablesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2headertablesize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2headertablesize = d.Get("http2headertablesize").(int)
		hasChange = true
	}
	if d.HasChange("http2initialconnwindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2initialconnwindowsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2initialconnwindowsize = d.Get("http2initialconnwindowsize").(int)
		hasChange = true
	}
	if d.HasChange("http2initialwindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2initialwindowsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2initialwindowsize = d.Get("http2initialwindowsize").(int)
		hasChange = true
	}
	if d.HasChange("http2maxconcurrentstreams") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxconcurrentstreams has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxconcurrentstreams = d.Get("http2maxconcurrentstreams").(int)
		hasChange = true
	}
	if d.HasChange("http2maxemptyframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxemptyframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxemptyframespermin = d.Get("http2maxemptyframespermin").(int)
		hasChange = true
	}
	if d.HasChange("http2maxframesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxframesize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxframesize = d.Get("http2maxframesize").(int)
		hasChange = true
	}
	if d.HasChange("http2maxheaderlistsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxheaderlistsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxheaderlistsize = d.Get("http2maxheaderlistsize").(int)
		hasChange = true
	}
	if d.HasChange("http2maxpingframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxpingframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxpingframespermin = d.Get("http2maxpingframespermin").(int)
		hasChange = true
	}
	if d.HasChange("http2maxresetframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxresetframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxresetframespermin = d.Get("http2maxresetframespermin").(int)
		hasChange = true
	}
	if d.HasChange("http2maxsettingsframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxsettingsframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxsettingsframespermin = d.Get("http2maxsettingsframespermin").(int)
		hasChange = true
	}
	if d.HasChange("http2minseverconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2minseverconn has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2minseverconn = d.Get("http2minseverconn").(int)
		hasChange = true
	}
	if d.HasChange("http2strictcipher") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2strictcipher has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2strictcipher = d.Get("http2strictcipher").(string)
		hasChange = true
	}
	if d.HasChange("http3") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3 has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3 = d.Get("http3").(string)
		hasChange = true
	}
	if d.HasChange("http3maxheaderblockedstreams") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3maxheaderblockedstreams has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3maxheaderblockedstreams = d.Get("http3maxheaderblockedstreams").(int)
		hasChange = true
	}
	if d.HasChange("http3maxheaderfieldsectionsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3maxheaderfieldsectionsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3maxheaderfieldsectionsize = d.Get("http3maxheaderfieldsectionsize").(int)
		hasChange = true
	}
	if d.HasChange("http3maxheadertablesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3maxheadertablesize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3maxheadertablesize = d.Get("http3maxheadertablesize").(int)
		hasChange = true
	}
	if d.HasChange("incomphdrdelay") {
		log.Printf("[DEBUG]  citrixadc-provider: Incomphdrdelay has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Incomphdrdelay = d.Get("incomphdrdelay").(int)
		hasChange = true
	}
	if d.HasChange("markconnreqinval") {
		log.Printf("[DEBUG]  citrixadc-provider: Markconnreqinval has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Markconnreqinval = d.Get("markconnreqinval").(string)
		hasChange = true
	}
	if d.HasChange("markhttp09inval") {
		log.Printf("[DEBUG]  citrixadc-provider: Markhttp09inval has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Markhttp09inval = d.Get("markhttp09inval").(string)
		hasChange = true
	}
	if d.HasChange("markhttpheaderextrawserror") {
		log.Printf("[DEBUG]  citrixadc-provider: Markhttpheaderextrawserror has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Markhttpheaderextrawserror = d.Get("markhttpheaderextrawserror").(string)
		hasChange = true
	}
	if d.HasChange("markrfc7230noncompliantinval") {
		log.Printf("[DEBUG]  citrixadc-provider: Markrfc7230noncompliantinval has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Markrfc7230noncompliantinval = d.Get("markrfc7230noncompliantinval").(string)
		hasChange = true
	}
	if d.HasChange("marktracereqinval") {
		log.Printf("[DEBUG]  citrixadc-provider: Marktracereqinval has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Marktracereqinval = d.Get("marktracereqinval").(string)
		hasChange = true
	}
	if d.HasChange("maxheaderlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxheaderlen has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxheaderlen = d.Get("maxheaderlen").(int)
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxreq has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxreq = d.Get("maxreq").(int)
		hasChange = true
	}
	if d.HasChange("maxreusepool") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxreusepool has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxreusepool = d.Get("maxreusepool").(int)
		hasChange = true
	}
	if d.HasChange("minreusepool") {
		log.Printf("[DEBUG]  citrixadc-provider: Minreusepool has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Minreusepool = d.Get("minreusepool").(int)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("persistentetag") {
		log.Printf("[DEBUG]  citrixadc-provider: Persistentetag has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Persistentetag = d.Get("persistentetag").(string)
		hasChange = true
	}
	if d.HasChange("reqtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqtimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Reqtimeout = d.Get("reqtimeout").(int)
		hasChange = true
	}
	if d.HasChange("reqtimeoutaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqtimeoutaction has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Reqtimeoutaction = d.Get("reqtimeoutaction").(string)
		hasChange = true
	}
	if d.HasChange("reusepooltimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reusepooltimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Reusepooltimeout = d.Get("reusepooltimeout").(int)
		hasChange = true
	}
	if d.HasChange("rtsptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtsptunnel has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Rtsptunnel = d.Get("rtsptunnel").(string)
		hasChange = true
	}
	if d.HasChange("spdy") {
		log.Printf("[DEBUG]  citrixadc-provider: Spdy has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Spdy = d.Get("spdy").(string)
		hasChange = true
	}
	if d.HasChange("weblog") {
		log.Printf("[DEBUG]  citrixadc-provider: Weblog has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Weblog = d.Get("weblog").(string)
		hasChange = true
	}
	if d.HasChange("websocket") {
		log.Printf("[DEBUG]  citrixadc-provider: Websocket has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Websocket = d.Get("websocket").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
		if err != nil {
			return fmt.Errorf("Error updating nshttpprofile %s", nshttpprofileName)
		}
	}
	return readNshttpprofileFunc(d, meta)
}

func deleteNshttpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Id()
	err := client.DeleteResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
