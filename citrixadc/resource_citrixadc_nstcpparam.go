package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcNstcpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createNstcpparamFunc,
		ReadContext:   readNstcpparamFunc,
		DeleteContext: deleteNstcpparamFunc,
		Schema: map[string]*schema.Schema{
			"rfc5961chlgacklimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpsendsfresetoption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpreliableaddaddr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpfastcloseoption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enhancedisngeneration": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"delinkclientserveronrst": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compacttcpoptionnoop": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ackonpush": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"autosyncookietimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"connflushifnomem": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"connflushthres": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"delayedack": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"downstaterst": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"initialcwnd": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"kaprobeupdatelastactivity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"learnvsvrmss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"limitedpersist": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxburst": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxdynserverprobes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxpktpermss": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxsynackretx": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxsynhold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxsynholdperprobe": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxtimewaitconn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"minrto": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpchecksum": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpclosemptcpsessiononlastsfclose": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpconcloseonpassivesf": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpimmediatesfcloseonfin": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpmaxpendingsf": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpmaxsf": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcppendingjointhreshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcprtostoswitchsf": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpsfreplacetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpsftimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpusebackupondss": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"msslearndelay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"msslearninterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nagle": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"oooqsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pktperretx": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"recvbuffsize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sack": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"slowstartincr": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"synattackdetection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"synholdfastgiveup": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpfastopencookietimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpfintimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpmaxretries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ws": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wsval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNstcpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstcpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpparamName := resource.PrefixedUniqueId("tf-nstcpparam-")

	nstcpparam := ns.Nstcpparam{
		Ackonpush:                           d.Get("ackonpush").(string),
		Connflushifnomem:                    d.Get("connflushifnomem").(string),
		Downstaterst:                        d.Get("downstaterst").(string),
		Kaprobeupdatelastactivity:           d.Get("kaprobeupdatelastactivity").(string),
		Learnvsvrmss:                        d.Get("learnvsvrmss").(string),
		Limitedpersist:                      d.Get("limitedpersist").(string),
		Mptcpchecksum:                       d.Get("mptcpchecksum").(string),
		Mptcpclosemptcpsessiononlastsfclose: d.Get("mptcpclosemptcpsessiononlastsfclose").(string),
		Mptcpconcloseonpassivesf:            d.Get("mptcpconcloseonpassivesf").(string),
		Mptcpimmediatesfcloseonfin:          d.Get("mptcpimmediatesfcloseonfin").(string),
		Mptcpusebackupondss:                 d.Get("mptcpusebackupondss").(string),
		Nagle:                               d.Get("nagle").(string),
		Sack:                                d.Get("sack").(string),
		Synattackdetection:                  d.Get("synattackdetection").(string),
		Ws:                                  d.Get("ws").(string),
		Compacttcpoptionnoop:                d.Get("compacttcpoptionnoop").(string),
		Delinkclientserveronrst:             d.Get("delinkclientserveronrst").(string),
		Enhancedisngeneration:               d.Get("enhancedisngeneration").(string),
		Mptcpfastcloseoption:                d.Get("mptcpfastcloseoption").(string),
		Mptcpreliableaddaddr:                d.Get("mptcpreliableaddaddr").(string),
		Mptcpsendsfresetoption:              d.Get("mptcpsendsfresetoption").(string),
	}
	if raw := d.GetRawConfig().GetAttr("rfc5961chlgacklimit"); !raw.IsNull() {
		nstcpparam.Rfc5961chlgacklimit = intPtr(d.Get("rfc5961chlgacklimit").(int))
	}
	if raw := d.GetRawConfig().GetAttr("autosyncookietimeout"); !raw.IsNull() {
		nstcpparam.Autosyncookietimeout = intPtr(d.Get("autosyncookietimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("connflushthres"); !raw.IsNull() {
		nstcpparam.Connflushthres = intPtr(d.Get("connflushthres").(int))
	}
	if raw := d.GetRawConfig().GetAttr("delayedack"); !raw.IsNull() {
		nstcpparam.Delayedack = intPtr(d.Get("delayedack").(int))
	}
	if raw := d.GetRawConfig().GetAttr("initialcwnd"); !raw.IsNull() {
		nstcpparam.Initialcwnd = intPtr(d.Get("initialcwnd").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxburst"); !raw.IsNull() {
		nstcpparam.Maxburst = intPtr(d.Get("maxburst").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxdynserverprobes"); !raw.IsNull() {
		nstcpparam.Maxdynserverprobes = intPtr(d.Get("maxdynserverprobes").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxpktpermss"); !raw.IsNull() {
		nstcpparam.Maxpktpermss = intPtr(d.Get("maxpktpermss").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxsynackretx"); !raw.IsNull() {
		nstcpparam.Maxsynackretx = intPtr(d.Get("maxsynackretx").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcpmaxpendingsf"); !raw.IsNull() {
		nstcpparam.Mptcpmaxpendingsf = intPtr(d.Get("mptcpmaxpendingsf").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcppendingjointhreshold"); !raw.IsNull() {
		nstcpparam.Mptcppendingjointhreshold = intPtr(d.Get("mptcppendingjointhreshold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcpsfreplacetimeout"); !raw.IsNull() {
		nstcpparam.Mptcpsfreplacetimeout = intPtr(d.Get("mptcpsfreplacetimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcpsftimeout"); !raw.IsNull() {
		nstcpparam.Mptcpsftimeout = intPtr(d.Get("mptcpsftimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("oooqsize"); !raw.IsNull() {
		nstcpparam.Oooqsize = intPtr(d.Get("oooqsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpfastopencookietimeout"); !raw.IsNull() {
		nstcpparam.Tcpfastopencookietimeout = intPtr(d.Get("tcpfastopencookietimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("wsval"); !raw.IsNull() {
		nstcpparam.Wsval = intPtr(d.Get("wsval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxsynhold"); !raw.IsNull() {
		nstcpparam.Maxsynhold = intPtr(d.Get("maxsynhold").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxsynholdperprobe"); !raw.IsNull() {
		nstcpparam.Maxsynholdperprobe = intPtr(d.Get("maxsynholdperprobe").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxtimewaitconn"); !raw.IsNull() {
		nstcpparam.Maxtimewaitconn = intPtr(d.Get("maxtimewaitconn").(int))
	}
	if raw := d.GetRawConfig().GetAttr("minrto"); !raw.IsNull() {
		nstcpparam.Minrto = intPtr(d.Get("minrto").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcpmaxsf"); !raw.IsNull() {
		nstcpparam.Mptcpmaxsf = intPtr(d.Get("mptcpmaxsf").(int))
	}
	if raw := d.GetRawConfig().GetAttr("mptcprtostoswitchsf"); !raw.IsNull() {
		nstcpparam.Mptcprtostoswitchsf = intPtr(d.Get("mptcprtostoswitchsf").(int))
	}
	if raw := d.GetRawConfig().GetAttr("msslearndelay"); !raw.IsNull() {
		nstcpparam.Msslearndelay = intPtr(d.Get("msslearndelay").(int))
	}
	if raw := d.GetRawConfig().GetAttr("msslearninterval"); !raw.IsNull() {
		nstcpparam.Msslearninterval = intPtr(d.Get("msslearninterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("pktperretx"); !raw.IsNull() {
		nstcpparam.Pktperretx = intPtr(d.Get("pktperretx").(int))
	}
	if raw := d.GetRawConfig().GetAttr("recvbuffsize"); !raw.IsNull() {
		nstcpparam.Recvbuffsize = intPtr(d.Get("recvbuffsize").(int))
	}
	if raw := d.GetRawConfig().GetAttr("slowstartincr"); !raw.IsNull() {
		nstcpparam.Slowstartincr = intPtr(d.Get("slowstartincr").(int))
	}
	if raw := d.GetRawConfig().GetAttr("synholdfastgiveup"); !raw.IsNull() {
		nstcpparam.Synholdfastgiveup = intPtr(d.Get("synholdfastgiveup").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpfintimeout"); !raw.IsNull() {
		nstcpparam.Tcpfintimeout = intPtr(d.Get("tcpfintimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("tcpmaxretries"); !raw.IsNull() {
		nstcpparam.Tcpmaxretries = intPtr(d.Get("tcpmaxretries").(int))
	}
	err := client.UpdateUnnamedResource(service.Nstcpparam.Type(), &nstcpparam)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(nstcpparamName)

	return readNstcpparamFunc(ctx, d, meta)
}

func readNstcpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstcpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstcpparam state %s", nstcpparamName)
	findParams := service.FindParams{
		ResourceType: "nstcpparam",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return diag.FromErr(err)
	}
	// There is always a single entry
	data := dataArr[0]
	log.Printf("[DEBUG] citrixadc-provider: data read %v", data)

	d.Set("ackonpush", data["ackonpush"])
	setToInt("rfc5961chlgacklimit", d, data["rfc5961chlgacklimit"])
	d.Set("mptcpsendsfresetoption", data["mptcpsendsfresetoption"])
	d.Set("mptcpreliableaddaddr", data["mptcpreliableaddaddr"])
	d.Set("mptcpfastcloseoption", data["mptcpfastcloseoption"])
	d.Set("enhancedisngeneration", data["enhancedisngeneration"])
	d.Set("delinkclientserveronrst", data["delinkclientserveronrst"])
	d.Set("compacttcpoptionnoop", data["compacttcpoptionnoop"])
	setToInt("autosyncookietimeout", d, data["autosyncookietimeout"])
	d.Set("connflushifnomem", data["connflushifnomem"])
	setToInt("connflushthres", d, data["connflushthres"])
	setToInt("delayedack", d, data["delayedack"])
	d.Set("downstaterst", data["downstaterst"])
	setToInt("initialcwnd", d, data["initialcwnd"])
	d.Set("kaprobeupdatelastactivity", data["kaprobeupdatelastactivity"])
	d.Set("learnvsvrmss", data["learnvsvrmss"])
	d.Set("limitedpersist", data["limitedpersist"])
	setToInt("maxburst", d, data["maxburst"])
	setToInt("maxdynserverprobes", d, data["maxdynserverprobes"])
	setToInt("maxpktpermss", d, data["maxpktpermss"])
	setToInt("maxsynackretx", d, data["maxsynackretx"])
	setToInt("maxsynhold", d, data["maxsynhold"])
	setToInt("maxsynholdperprobe", d, data["maxsynholdperprobe"])
	setToInt("maxtimewaitconn", d, data["maxtimewaitconn"])
	setToInt("minrto", d, data["minrto"])
	d.Set("mptcpchecksum", data["mptcpchecksum"])
	d.Set("mptcpclosemptcpsessiononlastsfclose", data["mptcpclosemptcpsessiononlastsfclose"])
	d.Set("mptcpconcloseonpassivesf", data["mptcpconcloseonpassivesf"])
	d.Set("mptcpimmediatesfcloseonfin", data["mptcpimmediatesfcloseonfin"])
	setToInt("mptcpmaxpendingsf", d, data["mptcpmaxpendingsf"])
	setToInt("mptcpmaxsf", d, data["mptcpmaxsf"])
	setToInt("mptcppendingjointhreshold", d, data["mptcppendingjointhreshold"])
	setToInt("mptcprtostoswitchsf", d, data["mptcprtostoswitchsf"])
	setToInt("mptcpsfreplacetimeout", d, data["mptcpsfreplacetimeout"])
	setToInt("mptcpsftimeout", d, data["mptcpsftimeout"])
	d.Set("mptcpusebackupondss", data["mptcpusebackupondss"])
	setToInt("msslearndelay", d, data["msslearndelay"])
	setToInt("msslearninterval", d, data["msslearninterval"])
	d.Set("nagle", data["nagle"])
	setToInt("oooqsize", d, data["oooqsize"])
	setToInt("pktperretx", d, data["pktperretx"])
	setToInt("recvbuffsize", d, data["recvbuffsize"])
	d.Set("sack", data["sack"])
	setToInt("slowstartincr", d, data["slowstartincr"])
	d.Set("synattackdetection", data["synattackdetection"])
	setToInt("synholdfastgiveup", d, data["synholdfastgiveup"])
	setToInt("tcpfastopencookietimeout", d, data["tcpfastopencookietimeout"])
	setToInt("tcpfintimeout", d, data["tcpfintimeout"])
	setToInt("tcpmaxretries", d, data["tcpmaxretries"])
	d.Set("ws", data["ws"])
	setToInt("wsval", d, data["wsval"])

	return nil

}

func deleteNstcpparamFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstcpparamFunc")

	d.SetId("")

	return nil
}
