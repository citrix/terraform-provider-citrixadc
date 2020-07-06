package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNstcpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstcpprofileFunc,
		Read:          readNstcpprofileFunc,
		Update:        updateNstcpprofileFunc,
		Delete:        deleteNstcpprofileFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ackaggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ackonpush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"applyadaptivetcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"buffersize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"burstratecontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientiptcpoption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientiptcpoptionnumber": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delayedack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dropestconnontimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drophalfclosedconnontimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dsack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dupackthresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamicreceivebuffering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"establishclientconn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"frto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hystart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"initialcwnd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kaconnidletime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"kamaxprobes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"kaprobeinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"kaprobeupdatelastactivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxburst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxcwnd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxpktpermss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minrto": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mptcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mptcpdropdataonpreestsf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mptcpfastopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mptcpsessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nagle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oooqsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pktperretx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rateqmax": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rstmaxack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rstwindowattenuate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendbuffsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"slowstartincr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"spoofsyndrop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"taillossprobe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpfastopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpfastopencookiesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcprate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpsegoffload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ws": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wsval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstcpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	var nstcpprofileName string
	if v, ok := d.GetOk("name"); ok {
		nstcpprofileName = v.(string)
	} else {
		nstcpprofileName = resource.PrefixedUniqueId("tf-nstcpprofile-")
		d.Set("name", nstcpprofileName)
	}
	nstcpprofile := ns.Nstcpprofile{
		Ackaggregation:              d.Get("ackaggregation").(string),
		Ackonpush:                   d.Get("ackonpush").(string),
		Applyadaptivetcp:            d.Get("applyadaptivetcp").(string),
		Buffersize:                  d.Get("buffersize").(int),
		Burstratecontrol:            d.Get("burstratecontrol").(string),
		Clientiptcpoption:           d.Get("clientiptcpoption").(string),
		Clientiptcpoptionnumber:     d.Get("clientiptcpoptionnumber").(int),
		Delayedack:                  d.Get("delayedack").(int),
		Dropestconnontimeout:        d.Get("dropestconnontimeout").(string),
		Drophalfclosedconnontimeout: d.Get("drophalfclosedconnontimeout").(string),
		Dsack:                       d.Get("dsack").(string),
		Dupackthresh:                d.Get("dupackthresh").(int),
		Dynamicreceivebuffering:     d.Get("dynamicreceivebuffering").(string),
		Ecn:                         d.Get("ecn").(string),
		Establishclientconn:         d.Get("establishclientconn").(string),
		Fack:                        d.Get("fack").(string),
		Flavor:                      d.Get("flavor").(string),
		Frto:                        d.Get("frto").(string),
		Hystart:                     d.Get("hystart").(string),
		Initialcwnd:                 d.Get("initialcwnd").(int),
		Ka:                          d.Get("ka").(string),
		Kaconnidletime:              d.Get("kaconnidletime").(int),
		Kamaxprobes:                 d.Get("kamaxprobes").(int),
		Kaprobeinterval:             d.Get("kaprobeinterval").(int),
		Kaprobeupdatelastactivity:   d.Get("kaprobeupdatelastactivity").(string),
		Maxburst:                    d.Get("maxburst").(int),
		Maxcwnd:                     d.Get("maxcwnd").(int),
		Maxpktpermss:                d.Get("maxpktpermss").(int),
		Minrto:                      d.Get("minrto").(int),
		Mptcp:                       d.Get("mptcp").(string),
		Mptcpdropdataonpreestsf:     d.Get("mptcpdropdataonpreestsf").(string),
		Mptcpfastopen:               d.Get("mptcpfastopen").(string),
		Mptcpsessiontimeout:         d.Get("mptcpsessiontimeout").(int),
		Mss:                         d.Get("mss").(int),
		Nagle:                       d.Get("nagle").(string),
		Name:                        d.Get("name").(string),
		Oooqsize:                    d.Get("oooqsize").(int),
		Pktperretx:                  d.Get("pktperretx").(int),
		Rateqmax:                    d.Get("rateqmax").(int),
		Rstmaxack:                   d.Get("rstmaxack").(string),
		Rstwindowattenuate:          d.Get("rstwindowattenuate").(string),
		Sack:                        d.Get("sack").(string),
		Sendbuffsize:                d.Get("sendbuffsize").(int),
		Slowstartincr:               d.Get("slowstartincr").(int),
		Spoofsyndrop:                d.Get("spoofsyndrop").(string),
		Syncookie:                   d.Get("syncookie").(string),
		Taillossprobe:               d.Get("taillossprobe").(string),
		Tcpfastopen:                 d.Get("tcpfastopen").(string),
		Tcpfastopencookiesize:       d.Get("tcpfastopencookiesize").(int),
		Tcpmode:                     d.Get("tcpmode").(string),
		Tcprate:                     d.Get("tcprate").(int),
		Tcpsegoffload:               d.Get("tcpsegoffload").(string),
		Timestamp:                   d.Get("timestamp").(string),
		Ws:                          d.Get("ws").(string),
		Wsval:                       d.Get("wsval").(int),
	}

	_, err := client.AddResource(netscaler.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
	if err != nil {
		return err
	}

	d.SetId(nstcpprofileName)

	err = readNstcpprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstcpprofile but we can't read it ?? %s", nstcpprofileName)
		return nil
	}
	return nil
}

func readNstcpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstcpprofile state %s", nstcpprofileName)
	data, err := client.FindResource(netscaler.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstcpprofile state %s", nstcpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("ackaggregation", data["ackaggregation"])
	d.Set("ackonpush", data["ackonpush"])
	d.Set("applyadaptivetcp", data["applyadaptivetcp"])
	d.Set("buffersize", data["buffersize"])
	d.Set("burstratecontrol", data["burstratecontrol"])
	d.Set("clientiptcpoption", data["clientiptcpoption"])
	d.Set("clientiptcpoptionnumber", data["clientiptcpoptionnumber"])
	d.Set("delayedack", data["delayedack"])
	d.Set("dropestconnontimeout", data["dropestconnontimeout"])
	d.Set("drophalfclosedconnontimeout", data["drophalfclosedconnontimeout"])
	d.Set("dsack", data["dsack"])
	d.Set("dupackthresh", data["dupackthresh"])
	d.Set("dynamicreceivebuffering", data["dynamicreceivebuffering"])
	d.Set("ecn", data["ecn"])
	d.Set("establishclientconn", data["establishclientconn"])
	d.Set("fack", data["fack"])
	d.Set("flavor", data["flavor"])
	d.Set("frto", data["frto"])
	d.Set("hystart", data["hystart"])
	d.Set("initialcwnd", data["initialcwnd"])
	d.Set("ka", data["ka"])
	d.Set("kaconnidletime", data["kaconnidletime"])
	d.Set("kamaxprobes", data["kamaxprobes"])
	d.Set("kaprobeinterval", data["kaprobeinterval"])
	d.Set("kaprobeupdatelastactivity", data["kaprobeupdatelastactivity"])
	d.Set("maxburst", data["maxburst"])
	d.Set("maxcwnd", data["maxcwnd"])
	d.Set("maxpktpermss", data["maxpktpermss"])
	d.Set("minrto", data["minrto"])
	d.Set("mptcp", data["mptcp"])
	d.Set("mptcpdropdataonpreestsf", data["mptcpdropdataonpreestsf"])
	d.Set("mptcpfastopen", data["mptcpfastopen"])
	d.Set("mptcpsessiontimeout", data["mptcpsessiontimeout"])
	d.Set("mss", data["mss"])
	d.Set("nagle", data["nagle"])
	d.Set("name", data["name"])
	d.Set("oooqsize", data["oooqsize"])
	d.Set("pktperretx", data["pktperretx"])
	d.Set("rateqmax", data["rateqmax"])
	d.Set("rstmaxack", data["rstmaxack"])
	d.Set("rstwindowattenuate", data["rstwindowattenuate"])
	d.Set("sack", data["sack"])
	d.Set("sendbuffsize", data["sendbuffsize"])
	d.Set("slowstartincr", data["slowstartincr"])
	d.Set("spoofsyndrop", data["spoofsyndrop"])
	d.Set("syncookie", data["syncookie"])
	d.Set("taillossprobe", data["taillossprobe"])
	d.Set("tcpfastopen", data["tcpfastopen"])
	d.Set("tcpfastopencookiesize", data["tcpfastopencookiesize"])
	d.Set("tcpmode", data["tcpmode"])
	d.Set("tcprate", data["tcprate"])
	d.Set("tcpsegoffload", data["tcpsegoffload"])
	d.Set("timestamp", data["timestamp"])
	d.Set("ws", data["ws"])
	d.Set("wsval", data["wsval"])

	return nil

}

func updateNstcpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpprofileName := d.Get("name").(string)

	nstcpprofile := ns.Nstcpprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("ackaggregation") {
		log.Printf("[DEBUG]  citrixadc-provider: Ackaggregation has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Ackaggregation = d.Get("ackaggregation").(string)
		hasChange = true
	}
	if d.HasChange("ackonpush") {
		log.Printf("[DEBUG]  citrixadc-provider: Ackonpush has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Ackonpush = d.Get("ackonpush").(string)
		hasChange = true
	}
	if d.HasChange("applyadaptivetcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Applyadaptivetcp has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Applyadaptivetcp = d.Get("applyadaptivetcp").(string)
		hasChange = true
	}
	if d.HasChange("buffersize") {
		log.Printf("[DEBUG]  citrixadc-provider: Buffersize has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Buffersize = d.Get("buffersize").(int)
		hasChange = true
	}
	if d.HasChange("burstratecontrol") {
		log.Printf("[DEBUG]  citrixadc-provider: Burstratecontrol has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Burstratecontrol = d.Get("burstratecontrol").(string)
		hasChange = true
	}
	if d.HasChange("clientiptcpoption") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientiptcpoption has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Clientiptcpoption = d.Get("clientiptcpoption").(string)
		hasChange = true
	}
	if d.HasChange("clientiptcpoptionnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientiptcpoptionnumber has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Clientiptcpoptionnumber = d.Get("clientiptcpoptionnumber").(int)
		hasChange = true
	}
	if d.HasChange("delayedack") {
		log.Printf("[DEBUG]  citrixadc-provider: Delayedack has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Delayedack = d.Get("delayedack").(int)
		hasChange = true
	}
	if d.HasChange("dropestconnontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropestconnontimeout has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Dropestconnontimeout = d.Get("dropestconnontimeout").(string)
		hasChange = true
	}
	if d.HasChange("drophalfclosedconnontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Drophalfclosedconnontimeout has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Drophalfclosedconnontimeout = d.Get("drophalfclosedconnontimeout").(string)
		hasChange = true
	}
	if d.HasChange("dsack") {
		log.Printf("[DEBUG]  citrixadc-provider: Dsack has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Dsack = d.Get("dsack").(string)
		hasChange = true
	}
	if d.HasChange("dupackthresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dupackthresh has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Dupackthresh = d.Get("dupackthresh").(int)
		hasChange = true
	}
	if d.HasChange("dynamicreceivebuffering") {
		log.Printf("[DEBUG]  citrixadc-provider: Dynamicreceivebuffering has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Dynamicreceivebuffering = d.Get("dynamicreceivebuffering").(string)
		hasChange = true
	}
	if d.HasChange("ecn") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecn has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Ecn = d.Get("ecn").(string)
		hasChange = true
	}
	if d.HasChange("establishclientconn") {
		log.Printf("[DEBUG]  citrixadc-provider: Establishclientconn has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Establishclientconn = d.Get("establishclientconn").(string)
		hasChange = true
	}
	if d.HasChange("fack") {
		log.Printf("[DEBUG]  citrixadc-provider: Fack has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Fack = d.Get("fack").(string)
		hasChange = true
	}
	if d.HasChange("flavor") {
		log.Printf("[DEBUG]  citrixadc-provider: Flavor has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Flavor = d.Get("flavor").(string)
		hasChange = true
	}
	if d.HasChange("frto") {
		log.Printf("[DEBUG]  citrixadc-provider: Frto has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Frto = d.Get("frto").(string)
		hasChange = true
	}
	if d.HasChange("hystart") {
		log.Printf("[DEBUG]  citrixadc-provider: Hystart has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Hystart = d.Get("hystart").(string)
		hasChange = true
	}
	if d.HasChange("initialcwnd") {
		log.Printf("[DEBUG]  citrixadc-provider: Initialcwnd has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Initialcwnd = d.Get("initialcwnd").(int)
		hasChange = true
	}
	if d.HasChange("ka") {
		log.Printf("[DEBUG]  citrixadc-provider: Ka has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Ka = d.Get("ka").(string)
		hasChange = true
	}
	if d.HasChange("kaconnidletime") {
		log.Printf("[DEBUG]  citrixadc-provider: Kaconnidletime has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Kaconnidletime = d.Get("kaconnidletime").(int)
		hasChange = true
	}
	if d.HasChange("kamaxprobes") {
		log.Printf("[DEBUG]  citrixadc-provider: Kamaxprobes has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Kamaxprobes = d.Get("kamaxprobes").(int)
		hasChange = true
	}
	if d.HasChange("kaprobeinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Kaprobeinterval has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Kaprobeinterval = d.Get("kaprobeinterval").(int)
		hasChange = true
	}
	if d.HasChange("kaprobeupdatelastactivity") {
		log.Printf("[DEBUG]  citrixadc-provider: Kaprobeupdatelastactivity has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Kaprobeupdatelastactivity = d.Get("kaprobeupdatelastactivity").(string)
		hasChange = true
	}
	if d.HasChange("maxburst") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxburst has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Maxburst = d.Get("maxburst").(int)
		hasChange = true
	}
	if d.HasChange("maxcwnd") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxcwnd has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Maxcwnd = d.Get("maxcwnd").(int)
		hasChange = true
	}
	if d.HasChange("maxpktpermss") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxpktpermss has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Maxpktpermss = d.Get("maxpktpermss").(int)
		hasChange = true
	}
	if d.HasChange("minrto") {
		log.Printf("[DEBUG]  citrixadc-provider: Minrto has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Minrto = d.Get("minrto").(int)
		hasChange = true
	}
	if d.HasChange("mptcp") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcp has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Mptcp = d.Get("mptcp").(string)
		hasChange = true
	}
	if d.HasChange("mptcpdropdataonpreestsf") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcpdropdataonpreestsf has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Mptcpdropdataonpreestsf = d.Get("mptcpdropdataonpreestsf").(string)
		hasChange = true
	}
	if d.HasChange("mptcpfastopen") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcpfastopen has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Mptcpfastopen = d.Get("mptcpfastopen").(string)
		hasChange = true
	}
	if d.HasChange("mptcpsessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Mptcpsessiontimeout has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Mptcpsessiontimeout = d.Get("mptcpsessiontimeout").(int)
		hasChange = true
	}
	if d.HasChange("mss") {
		log.Printf("[DEBUG]  citrixadc-provider: Mss has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Mss = d.Get("mss").(int)
		hasChange = true
	}
	if d.HasChange("nagle") {
		log.Printf("[DEBUG]  citrixadc-provider: Nagle has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Nagle = d.Get("nagle").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("oooqsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Oooqsize has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Oooqsize = d.Get("oooqsize").(int)
		hasChange = true
	}
	if d.HasChange("pktperretx") {
		log.Printf("[DEBUG]  citrixadc-provider: Pktperretx has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Pktperretx = d.Get("pktperretx").(int)
		hasChange = true
	}
	if d.HasChange("rateqmax") {
		log.Printf("[DEBUG]  citrixadc-provider: Rateqmax has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Rateqmax = d.Get("rateqmax").(int)
		hasChange = true
	}
	if d.HasChange("rstmaxack") {
		log.Printf("[DEBUG]  citrixadc-provider: Rstmaxack has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Rstmaxack = d.Get("rstmaxack").(string)
		hasChange = true
	}
	if d.HasChange("rstwindowattenuate") {
		log.Printf("[DEBUG]  citrixadc-provider: Rstwindowattenuate has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Rstwindowattenuate = d.Get("rstwindowattenuate").(string)
		hasChange = true
	}
	if d.HasChange("sack") {
		log.Printf("[DEBUG]  citrixadc-provider: Sack has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Sack = d.Get("sack").(string)
		hasChange = true
	}
	if d.HasChange("sendbuffsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendbuffsize has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Sendbuffsize = d.Get("sendbuffsize").(int)
		hasChange = true
	}
	if d.HasChange("slowstartincr") {
		log.Printf("[DEBUG]  citrixadc-provider: Slowstartincr has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Slowstartincr = d.Get("slowstartincr").(int)
		hasChange = true
	}
	if d.HasChange("spoofsyndrop") {
		log.Printf("[DEBUG]  citrixadc-provider: Spoofsyndrop has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Spoofsyndrop = d.Get("spoofsyndrop").(string)
		hasChange = true
	}
	if d.HasChange("syncookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Syncookie has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Syncookie = d.Get("syncookie").(string)
		hasChange = true
	}
	if d.HasChange("taillossprobe") {
		log.Printf("[DEBUG]  citrixadc-provider: Taillossprobe has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Taillossprobe = d.Get("taillossprobe").(string)
		hasChange = true
	}
	if d.HasChange("tcpfastopen") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpfastopen has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Tcpfastopen = d.Get("tcpfastopen").(string)
		hasChange = true
	}
	if d.HasChange("tcpfastopencookiesize") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpfastopencookiesize has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Tcpfastopencookiesize = d.Get("tcpfastopencookiesize").(int)
		hasChange = true
	}
	if d.HasChange("tcpmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpmode has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Tcpmode = d.Get("tcpmode").(string)
		hasChange = true
	}
	if d.HasChange("tcprate") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcprate has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Tcprate = d.Get("tcprate").(int)
		hasChange = true
	}
	if d.HasChange("tcpsegoffload") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpsegoffload has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Tcpsegoffload = d.Get("tcpsegoffload").(string)
		hasChange = true
	}
	if d.HasChange("timestamp") {
		log.Printf("[DEBUG]  citrixadc-provider: Timestamp has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Timestamp = d.Get("timestamp").(string)
		hasChange = true
	}
	if d.HasChange("ws") {
		log.Printf("[DEBUG]  citrixadc-provider: Ws has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Ws = d.Get("ws").(string)
		hasChange = true
	}
	if d.HasChange("wsval") {
		log.Printf("[DEBUG]  citrixadc-provider: Wsval has changed for nstcpprofile %s, starting update", nstcpprofileName)
		nstcpprofile.Wsval = d.Get("wsval").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Nstcpprofile.Type(), nstcpprofileName, &nstcpprofile)
		if err != nil {
			return fmt.Errorf("Error updating nstcpprofile %s", nstcpprofileName)
		}
	}
	return readNstcpprofileFunc(d, meta)
}

func deleteNstcpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstcpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpprofileName := d.Id()
	err := client.DeleteResource(netscaler.Nstcpprofile.Type(), nstcpprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
