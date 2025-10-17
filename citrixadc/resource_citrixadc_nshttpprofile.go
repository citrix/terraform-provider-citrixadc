package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNshttpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNshttpprofileFunc,
		ReadContext:   readNshttpprofileFunc,
		UpdateContext: updateNshttpprofileFunc,
		DeleteContext: deleteNshttpprofileFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
			"maxheaderfieldlen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxrxresetframespermin": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http3webtransport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http3minseverconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httppipelinebuffsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"allowonlywordcharactersandhyphen": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostheadervalidation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxduplicateheaderfields": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passprotocolupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2extendedconnect": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNshttpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Get("name").(string)

	nshttpprofile := ns.Nshttpprofile{
		Adpttimeout:                      d.Get("adpttimeout").(string),
		Altsvc:                           d.Get("altsvc").(string),
		Altsvcvalue:                      d.Get("altsvcvalue").(string),
		Clientiphdrexpr:                  d.Get("clientiphdrexpr").(string),
		Cmponpush:                        d.Get("cmponpush").(string),
		Conmultiplex:                     d.Get("conmultiplex").(string),
		Dropextracrlf:                    d.Get("dropextracrlf").(string),
		Dropextradata:                    d.Get("dropextradata").(string),
		Dropinvalreqs:                    d.Get("dropinvalreqs").(string),
		Grpclengthdelimitation:           d.Get("grpclengthdelimitation").(string),
		Http2:                            d.Get("http2").(string),
		Http2altsvcframe:                 d.Get("http2altsvcframe").(string),
		Http2direct:                      d.Get("http2direct").(string),
		Http2strictcipher:                d.Get("http2strictcipher").(string),
		Http3:                            d.Get("http3").(string),
		Markconnreqinval:                 d.Get("markconnreqinval").(string),
		Markhttp09inval:                  d.Get("markhttp09inval").(string),
		Markhttpheaderextrawserror:       d.Get("markhttpheaderextrawserror").(string),
		Markrfc7230noncompliantinval:     d.Get("markrfc7230noncompliantinval").(string),
		Marktracereqinval:                d.Get("marktracereqinval").(string),
		Name:                             d.Get("name").(string),
		Persistentetag:                   d.Get("persistentetag").(string),
		Reqtimeoutaction:                 d.Get("reqtimeoutaction").(string),
		Rtsptunnel:                       d.Get("rtsptunnel").(string),
		Weblog:                           d.Get("weblog").(string),
		Websocket:                        d.Get("websocket").(string),
		Http3webtransport:                d.Get("http3webtransport").(string),
		Allowonlywordcharactersandhyphen: d.Get("allowonlywordcharactersandhyphen").(string),
		Hostheadervalidation:             d.Get("hostheadervalidation").(string),
		Passprotocolupgrade:              d.Get("passprotocolupgrade").(string),
		Http2extendedconnect:             d.Get("http2extendedconnect").(string),
	}

	if raw := d.GetRawConfig().GetAttr("apdexcltresptimethreshold"); !raw.IsNull() {
		nshttpprofile.Apdexcltresptimethreshold = intPtr(d.Get("apdexcltresptimethreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("grpcholdlimit"); !raw.IsNull() {
		nshttpprofile.Grpcholdlimit = intPtr(d.Get("grpcholdlimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("grpcholdtimeout"); !raw.IsNull() {
		nshttpprofile.Grpcholdtimeout = intPtr(d.Get("grpcholdtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2headertablesize"); !raw.IsNull() {
		nshttpprofile.Http2headertablesize = intPtr(d.Get("http2headertablesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2initialconnwindowsize"); !raw.IsNull() {
		nshttpprofile.Http2initialconnwindowsize = intPtr(d.Get("http2initialconnwindowsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2initialwindowsize"); !raw.IsNull() {
		nshttpprofile.Http2initialwindowsize = intPtr(d.Get("http2initialwindowsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxconcurrentstreams"); !raw.IsNull() {
		nshttpprofile.Http2maxconcurrentstreams = intPtr(d.Get("http2maxconcurrentstreams").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxemptyframespermin"); !raw.IsNull() {
		nshttpprofile.Http2maxemptyframespermin = intPtr(d.Get("http2maxemptyframespermin").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxframesize"); !raw.IsNull() {
		nshttpprofile.Http2maxframesize = intPtr(d.Get("http2maxframesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxheaderlistsize"); !raw.IsNull() {
		nshttpprofile.Http2maxheaderlistsize = intPtr(d.Get("http2maxheaderlistsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxpingframespermin"); !raw.IsNull() {
		nshttpprofile.Http2maxpingframespermin = intPtr(d.Get("http2maxpingframespermin").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxresetframespermin"); !raw.IsNull() {
		nshttpprofile.Http2maxresetframespermin = intPtr(d.Get("http2maxresetframespermin").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxsettingsframespermin"); !raw.IsNull() {
		nshttpprofile.Http2maxsettingsframespermin = intPtr(d.Get("http2maxsettingsframespermin").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2minseverconn"); !raw.IsNull() {
		nshttpprofile.Http2minseverconn = intPtr(d.Get("http2minseverconn").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http3maxheaderblockedstreams"); !raw.IsNull() {
		nshttpprofile.Http3maxheaderblockedstreams = intPtr(d.Get("http3maxheaderblockedstreams").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http3maxheaderfieldsectionsize"); !raw.IsNull() {
		nshttpprofile.Http3maxheaderfieldsectionsize = intPtr(d.Get("http3maxheaderfieldsectionsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http3maxheadertablesize"); !raw.IsNull() {
		nshttpprofile.Http3maxheadertablesize = intPtr(d.Get("http3maxheadertablesize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("incomphdrdelay"); !raw.IsNull() {
		nshttpprofile.Incomphdrdelay = intPtr(d.Get("incomphdrdelay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxheaderlen"); !raw.IsNull() {
		nshttpprofile.Maxheaderlen = intPtr(d.Get("maxheaderlen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxreq"); !raw.IsNull() {
		nshttpprofile.Maxreq = intPtr(d.Get("maxreq").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxreusepool"); !raw.IsNull() {
		nshttpprofile.Maxreusepool = intPtr(d.Get("maxreusepool").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minreusepool"); !raw.IsNull() {
		nshttpprofile.Minreusepool = intPtr(d.Get("minreusepool").(int))
	}
	if raw := d.GetRawConfig().GetAttr("reqtimeout"); !raw.IsNull() {
		nshttpprofile.Reqtimeout = intPtr(d.Get("reqtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("reusepooltimeout"); !raw.IsNull() {
		nshttpprofile.Reusepooltimeout = intPtr(d.Get("reusepooltimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxheaderfieldlen"); !raw.IsNull() {
		nshttpprofile.Maxheaderfieldlen = intPtr(d.Get("maxheaderfieldlen").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http2maxrxresetframespermin"); !raw.IsNull() {
		nshttpprofile.Http2maxrxresetframespermin = intPtr(d.Get("http2maxrxresetframespermin").(int))
	}
	if raw := d.GetRawConfig().GetAttr("http3minseverconn"); !raw.IsNull() {
		nshttpprofile.Http3minseverconn = intPtr(d.Get("http3minseverconn").(int))
	}
	if raw := d.GetRawConfig().GetAttr("httppipelinebuffsize"); !raw.IsNull() {
		nshttpprofile.Httppipelinebuffsize = intPtr(d.Get("httppipelinebuffsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxduplicateheaderfields"); !raw.IsNull() {
		nshttpprofile.Maxduplicateheaderfields = intPtr(d.Get("maxduplicateheaderfields").(int))
	}

	_, err := client.AddResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nshttpprofileName)

	return readNshttpprofileFunc(ctx, d, meta)
}

func readNshttpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	d.Set("altsvc", data["altsvc"])
	d.Set("altsvcvalue", data["altsvcvalue"])
	setToInt("apdexcltresptimethreshold", d, data["apdexcltresptimethreshold"])
	d.Set("clientiphdrexpr", data["clientiphdrexpr"])
	d.Set("cmponpush", data["cmponpush"])
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("dropextracrlf", data["dropextracrlf"])
	d.Set("dropextradata", data["dropextradata"])
	d.Set("dropinvalreqs", data["dropinvalreqs"])
	setToInt("grpcholdlimit", d, data["grpcholdlimit"])
	setToInt("grpcholdtimeout", d, data["grpcholdtimeout"])
	d.Set("grpclengthdelimitation", data["grpclengthdelimitation"])
	d.Set("http2", data["http2"])
	d.Set("http2altsvcframe", data["http2altsvcframe"])
	d.Set("http2direct", data["http2direct"])
	setToInt("http2headertablesize", d, data["http2headertablesize"])
	setToInt("http2initialconnwindowsize", d, data["http2initialconnwindowsize"])
	setToInt("http2initialwindowsize", d, data["http2initialwindowsize"])
	setToInt("http2maxconcurrentstreams", d, data["http2maxconcurrentstreams"])
	setToInt("http2maxemptyframespermin", d, data["http2maxemptyframespermin"])
	setToInt("http2maxframesize", d, data["http2maxframesize"])
	setToInt("http2maxheaderlistsize", d, data["http2maxheaderlistsize"])
	setToInt("http2maxpingframespermin", d, data["http2maxpingframespermin"])
	setToInt("http2maxresetframespermin", d, data["http2maxresetframespermin"])
	setToInt("http2maxsettingsframespermin", d, data["http2maxsettingsframespermin"])
	setToInt("http2minseverconn", d, data["http2minseverconn"])
	d.Set("http2strictcipher", data["http2strictcipher"])
	d.Set("http3", data["http3"])
	setToInt("http3maxheaderblockedstreams", d, data["http3maxheaderblockedstreams"])
	setToInt("http3maxheaderfieldsectionsize", d, data["http3maxheaderfieldsectionsize"])
	setToInt("http3maxheadertablesize", d, data["http3maxheadertablesize"])
	setToInt("incomphdrdelay", d, data["incomphdrdelay"])
	d.Set("markconnreqinval", data["markconnreqinval"])
	d.Set("markhttp09inval", data["markhttp09inval"])
	d.Set("markhttpheaderextrawserror", data["markhttpheaderextrawserror"])
	d.Set("markrfc7230noncompliantinval", data["markrfc7230noncompliantinval"])
	d.Set("marktracereqinval", data["marktracereqinval"])
	setToInt("maxheaderlen", d, data["maxheaderlen"])
	setToInt("maxreq", d, data["maxreq"])
	setToInt("maxreusepool", d, data["maxreusepool"])
	setToInt("minreusepool", d, data["minreusepool"])
	d.Set("name", data["name"])
	d.Set("persistentetag", data["persistentetag"])
	setToInt("reqtimeout", d, data["reqtimeout"])
	d.Set("reqtimeoutaction", data["reqtimeoutaction"])
	setToInt("reusepooltimeout", d, data["reusepooltimeout"])
	d.Set("rtsptunnel", data["rtsptunnel"])
	d.Set("weblog", data["weblog"])
	d.Set("websocket", data["websocket"])
	setToInt("maxheaderfieldlen", d, data["maxheaderfieldlen"])
	setToInt("http2maxrxresetframespermin", d, data["http2maxrxresetframespermin"])
	d.Set("http3webtransport", data["http3webtransport"])
	setToInt("http3minseverconn", d, data["http3minseverconn"])
	setToInt("httppipelinebuffsize", d, data["httppipelinebuffsize"])
	d.Set("allowonlywordcharactersandhyphen", data["allowonlywordcharactersandhyphen"])
	d.Set("hostheadervalidation", data["hostheadervalidation"])
	setToInt("maxduplicateheaderfields", d, data["maxduplicateheaderfields"])
	d.Set("passprotocolupgrade", data["passprotocolupgrade"])
	d.Set("http2extendedconnect", data["http2extendedconnect"])

	return nil

}

func updateNshttpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		nshttpprofile.Apdexcltresptimethreshold = intPtr(d.Get("apdexcltresptimethreshold").(int))
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
		nshttpprofile.Grpcholdlimit = intPtr(d.Get("grpcholdlimit").(int))
		hasChange = true
	}
	if d.HasChange("grpcholdtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Grpcholdtimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Grpcholdtimeout = intPtr(d.Get("grpcholdtimeout").(int))
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
		nshttpprofile.Http2headertablesize = intPtr(d.Get("http2headertablesize").(int))
		hasChange = true
	}
	if d.HasChange("http2initialconnwindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2initialconnwindowsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2initialconnwindowsize = intPtr(d.Get("http2initialconnwindowsize").(int))
		hasChange = true
	}
	if d.HasChange("http2initialwindowsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2initialwindowsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2initialwindowsize = intPtr(d.Get("http2initialwindowsize").(int))
		hasChange = true
	}
	if d.HasChange("http2maxconcurrentstreams") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxconcurrentstreams has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxconcurrentstreams = intPtr(d.Get("http2maxconcurrentstreams").(int))
		hasChange = true
	}
	if d.HasChange("http2maxemptyframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxemptyframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxemptyframespermin = intPtr(d.Get("http2maxemptyframespermin").(int))
		hasChange = true
	}
	if d.HasChange("http2maxframesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxframesize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxframesize = intPtr(d.Get("http2maxframesize").(int))
		hasChange = true
	}
	if d.HasChange("http2maxheaderlistsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxheaderlistsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxheaderlistsize = intPtr(d.Get("http2maxheaderlistsize").(int))
		hasChange = true
	}
	if d.HasChange("http2maxpingframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxpingframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxpingframespermin = intPtr(d.Get("http2maxpingframespermin").(int))
		hasChange = true
	}
	if d.HasChange("http2maxresetframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxresetframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxresetframespermin = intPtr(d.Get("http2maxresetframespermin").(int))
		hasChange = true
	}
	if d.HasChange("http2maxsettingsframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxsettingsframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxsettingsframespermin = intPtr(d.Get("http2maxsettingsframespermin").(int))
		hasChange = true
	}
	if d.HasChange("http2minseverconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2minseverconn has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2minseverconn = intPtr(d.Get("http2minseverconn").(int))
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
		nshttpprofile.Http3maxheaderblockedstreams = intPtr(d.Get("http3maxheaderblockedstreams").(int))
		hasChange = true
	}
	if d.HasChange("http3maxheaderfieldsectionsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3maxheaderfieldsectionsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3maxheaderfieldsectionsize = intPtr(d.Get("http3maxheaderfieldsectionsize").(int))
		hasChange = true
	}
	if d.HasChange("http3maxheadertablesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3maxheadertablesize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3maxheadertablesize = intPtr(d.Get("http3maxheadertablesize").(int))
		hasChange = true
	}
	if d.HasChange("incomphdrdelay") {
		log.Printf("[DEBUG]  citrixadc-provider: Incomphdrdelay has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Incomphdrdelay = intPtr(d.Get("incomphdrdelay").(int))
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
		nshttpprofile.Maxheaderlen = intPtr(d.Get("maxheaderlen").(int))
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxreq has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxreq = intPtr(d.Get("maxreq").(int))
		hasChange = true
	}
	if d.HasChange("maxreusepool") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxreusepool has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxreusepool = intPtr(d.Get("maxreusepool").(int))
		hasChange = true
	}
	if d.HasChange("minreusepool") {
		log.Printf("[DEBUG]  citrixadc-provider: Minreusepool has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Minreusepool = intPtr(d.Get("minreusepool").(int))
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
		nshttpprofile.Reqtimeout = intPtr(d.Get("reqtimeout").(int))
		hasChange = true
	}
	if d.HasChange("reqtimeoutaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqtimeoutaction has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Reqtimeoutaction = d.Get("reqtimeoutaction").(string)
		hasChange = true
	}
	if d.HasChange("reusepooltimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reusepooltimeout has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Reusepooltimeout = intPtr(d.Get("reusepooltimeout").(int))
		hasChange = true
	}
	if d.HasChange("rtsptunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtsptunnel has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Rtsptunnel = d.Get("rtsptunnel").(string)
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
	if d.HasChange("maxheaderfieldlen") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxheaderfieldlen has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxheaderfieldlen = intPtr(d.Get("maxheaderfieldlen").(int))
		hasChange = true
	}
	if d.HasChange("http2maxrxresetframespermin") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2maxrxresetframespermin has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2maxrxresetframespermin = intPtr(d.Get("http2maxrxresetframespermin").(int))
		hasChange = true
	}
	if d.HasChange("http3webtransport") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3webtransport has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3webtransport = d.Get("http3webtransport").(string)
		hasChange = true
	}
	if d.HasChange("http3minseverconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Http3minseverconn has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http3minseverconn = intPtr(d.Get("http3minseverconn").(int))
		hasChange = true
	}
	if d.HasChange("httppipelinebuffsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Httppipelinebuffsize has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Httppipelinebuffsize = intPtr(d.Get("httppipelinebuffsize").(int))
		hasChange = true
	}
	if d.HasChange("allowonlywordcharactersandhyphen") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowonlywordcharactersandhyphen has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Allowonlywordcharactersandhyphen = d.Get("allowonlywordcharactersandhyphen").(string)
		hasChange = true
	}
	if d.HasChange("hostheadervalidation") {
		log.Printf("[DEBUG]  citrixadc-provider: Hostheadervalidation has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Hostheadervalidation = d.Get("hostheadervalidation").(string)
		hasChange = true
	}
	if d.HasChange("maxduplicateheaderfields") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxduplicateheaderfields has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Maxduplicateheaderfields = intPtr(d.Get("maxduplicateheaderfields").(int))
		hasChange = true
	}
	if d.HasChange("passprotocolupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Passprotocolupgrade has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Passprotocolupgrade = d.Get("passprotocolupgrade").(string)
		hasChange = true
	}
	if d.HasChange("http2extendedconnect") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2extendedconnect has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2extendedconnect = d.Get("http2extendedconnect").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
		if err != nil {
			return diag.Errorf("Error updating nshttpprofile %s", nshttpprofileName)
		}
	}
	return readNshttpprofileFunc(ctx, d, meta)
}

func deleteNshttpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNshttpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nshttpprofileName := d.Id()
	err := client.DeleteResource(service.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
