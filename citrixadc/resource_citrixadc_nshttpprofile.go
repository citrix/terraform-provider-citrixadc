package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

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
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"adpttimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altsvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apdexcltresptimethreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientiphdrexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmponpush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropextracrlf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropextradata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropinvalreqs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http2headertablesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2initialwindowsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxconcurrentstreams": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxframesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2maxheaderlistsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2minseverconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2strictcipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incomphdrdelay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"markconnreqinval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"markhttp09inval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"marktracereqinval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxheaderlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistentetag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reqtimeoutaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reusepooltimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtsptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spdy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weblog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websocket": &schema.Schema{
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
		Name:                      nshttpprofileName,
		Adpttimeout:               d.Get("adpttimeout").(string),
		Altsvc:                    d.Get("altsvc").(string),
		Apdexcltresptimethreshold: d.Get("apdexcltresptimethreshold").(int),
		Clientiphdrexpr:           d.Get("clientiphdrexpr").(string),
		Cmponpush:                 d.Get("cmponpush").(string),
		Conmultiplex:              d.Get("conmultiplex").(string),
		Dropextracrlf:             d.Get("dropextracrlf").(string),
		Dropextradata:             d.Get("dropextradata").(string),
		Dropinvalreqs:             d.Get("dropinvalreqs").(string),
		Http2:                     d.Get("http2").(string),
		Http2direct:               d.Get("http2direct").(string),
		Http2headertablesize:      d.Get("http2headertablesize").(int),
		Http2initialwindowsize:    d.Get("http2initialwindowsize").(int),
		Http2maxconcurrentstreams: d.Get("http2maxconcurrentstreams").(int),
		Http2maxframesize:         d.Get("http2maxframesize").(int),
		Http2maxheaderlistsize:    d.Get("http2maxheaderlistsize").(int),
		Http2minseverconn:         d.Get("http2minseverconn").(int),
		Http2strictcipher:         d.Get("http2strictcipher").(string),
		Incomphdrdelay:            d.Get("incomphdrdelay").(int),
		Markconnreqinval:          d.Get("markconnreqinval").(string),
		Markhttp09inval:           d.Get("markhttp09inval").(string),
		Marktracereqinval:         d.Get("marktracereqinval").(string),
		Maxheaderlen:              d.Get("maxheaderlen").(int),
		Maxreq:                    d.Get("maxreq").(int),
		Maxreusepool:              d.Get("maxreusepool").(int),
		Minreusepool:              d.Get("minreusepool").(int),
		Persistentetag:            d.Get("persistentetag").(string),
		Reqtimeout:                d.Get("reqtimeout").(int),
		Reqtimeoutaction:          d.Get("reqtimeoutaction").(string),
		Reusepooltimeout:          d.Get("reusepooltimeout").(int),
		Rtsptunnel:                d.Get("rtsptunnel").(string),
		Spdy:                      d.Get("spdy").(string),
		Weblog:                    d.Get("weblog").(string),
		Websocket:                 d.Get("websocket").(string),
	}

	_, err := client.AddResource(netscaler.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
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
	data, err := client.FindResource(netscaler.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nshttpprofile state %s", nshttpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("adpttimeout", data["adpttimeout"])
	d.Set("altsvc", data["altsvc"])
	d.Set("apdexcltresptimethreshold", data["apdexcltresptimethreshold"])
	d.Set("clientiphdrexpr", data["clientiphdrexpr"])
	d.Set("cmponpush", data["cmponpush"])
	d.Set("conmultiplex", data["conmultiplex"])
	d.Set("dropextracrlf", data["dropextracrlf"])
	d.Set("dropextradata", data["dropextradata"])
	d.Set("dropinvalreqs", data["dropinvalreqs"])
	d.Set("http2", data["http2"])
	d.Set("http2direct", data["http2direct"])
	d.Set("http2headertablesize", data["http2headertablesize"])
	d.Set("http2initialwindowsize", data["http2initialwindowsize"])
	d.Set("http2maxconcurrentstreams", data["http2maxconcurrentstreams"])
	d.Set("http2maxframesize", data["http2maxframesize"])
	d.Set("http2maxheaderlistsize", data["http2maxheaderlistsize"])
	d.Set("http2minseverconn", data["http2minseverconn"])
	d.Set("http2strictcipher", data["http2strictcipher"])
	d.Set("incomphdrdelay", data["incomphdrdelay"])
	d.Set("markconnreqinval", data["markconnreqinval"])
	d.Set("markhttp09inval", data["markhttp09inval"])
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
	if d.HasChange("altsvc") {
		log.Printf("[DEBUG]  citrixadc-provider: Altsvc has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Altsvc = d.Get("altsvc").(string)
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
	if d.HasChange("http2") {
		log.Printf("[DEBUG]  citrixadc-provider: Http2 has changed for nshttpprofile %s, starting update", nshttpprofileName)
		nshttpprofile.Http2 = d.Get("http2").(string)
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
		_, err := client.UpdateResource(netscaler.Nshttpprofile.Type(), nshttpprofileName, &nshttpprofile)
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
	err := client.DeleteResource(netscaler.Nshttpprofile.Type(), nshttpprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
