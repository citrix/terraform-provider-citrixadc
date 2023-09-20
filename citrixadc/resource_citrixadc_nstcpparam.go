package citrixadc

import (
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

// We need to convert fields that are int and accept zero values to string for correct operation
type Nstcpparam struct {
	Ackonpush                           string `json:"ackonpush,omitempty"`
	Autosyncookietimeout                int    `json:"autosyncookietimeout,omitempty"`
	Connflushifnomem                    string `json:"connflushifnomem,omitempty"`
	Connflushthres                      int    `json:"connflushthres,omitempty"`
	Delayedack                          int    `json:"delayedack,omitempty"`
	Downstaterst                        string `json:"downstaterst,omitempty"`
	Feature                             string `json:"feature,omitempty"`
	Initialcwnd                         int    `json:"initialcwnd,omitempty"`
	Kaprobeupdatelastactivity           string `json:"kaprobeupdatelastactivity,omitempty"`
	Learnvsvrmss                        string `json:"learnvsvrmss,omitempty"`
	Limitedpersist                      string `json:"limitedpersist,omitempty"`
	Maxburst                            int    `json:"maxburst,omitempty"`
	Maxdynserverprobes                  int    `json:"maxdynserverprobes,omitempty"`
	Maxpktpermss                        string `json:"maxpktpermss,omitempty"` // was int
	Maxsynackretx                       int    `json:"maxsynackretx,omitempty"`
	Maxsynhold                          int    `json:"maxsynhold,omitempty"`
	Maxsynholdperprobe                  int    `json:"maxsynholdperprobe,omitempty"`
	Maxtimewaitconn                     int    `json:"maxtimewaitconn,omitempty"`
	Minrto                              int    `json:"minrto,omitempty"`
	Mptcpchecksum                       string `json:"mptcpchecksum,omitempty"`
	Mptcpclosemptcpsessiononlastsfclose string `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	Mptcpconcloseonpassivesf            string `json:"mptcpconcloseonpassivesf,omitempty"`
	Mptcpimmediatesfcloseonfin          string `json:"mptcpimmediatesfcloseonfin,omitempty"`
	Mptcpmaxpendingsf                   string `json:"mptcpmaxpendingsf,omitempty"` // was int
	Mptcpmaxsf                          int    `json:"mptcpmaxsf,omitempty"`
	Mptcppendingjointhreshold           string `json:"mptcppendingjointhreshold,omitempty"` // was int
	Mptcprtostoswitchsf                 int    `json:"mptcprtostoswitchsf,omitempty"`
	Mptcpsfreplacetimeout               string `json:"mptcpsfreplacetimeout,omitempty"` // was int
	Mptcpsftimeout                      string `json:"mptcpsftimeout,omitempty"`        // was int
	Mptcpusebackupondss                 string `json:"mptcpusebackupondss,omitempty"`
	Msslearndelay                       int    `json:"msslearndelay,omitempty"`
	Msslearninterval                    int    `json:"msslearninterval,omitempty"`
	Nagle                               string `json:"nagle,omitempty"`
	Oooqsize                            string `json:"oooqsize,omitempty"` // was int
	Pktperretx                          int    `json:"pktperretx,omitempty"`
	Recvbuffsize                        int    `json:"recvbuffsize,omitempty"`
	Sack                                string `json:"sack,omitempty"`
	Slowstartincr                       int    `json:"slowstartincr,omitempty"`
	Synattackdetection                  string `json:"synattackdetection,omitempty"`
	Synholdfastgiveup                   int    `json:"synholdfastgiveup,omitempty"`
	Tcpfastopencookietimeout            string `json:"tcpfastopencookietimeout,omitempty"` // was int
	Tcpfintimeout                       int    `json:"tcpfintimeout,omitempty"`
	Tcpmaxretries                       int    `json:"tcpmaxretries,omitempty"`
	Ws                                  string `json:"ws,omitempty"`
	Wsval                               string `json:"wsval,omitempty"` // was int
}

func resourceCitrixAdcNstcpparam() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstcpparamFunc,
		Read:          readNstcpparamFunc,
		Delete:        deleteNstcpparamFunc,
		Schema: map[string]*schema.Schema{
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"mptcpsftimeout": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createNstcpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstcpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpparamName := resource.PrefixedUniqueId("tf-nstcpparam-")

	nstcpparam := Nstcpparam{
		Ackonpush:                           d.Get("ackonpush").(string),
		Autosyncookietimeout:                d.Get("autosyncookietimeout").(int),
		Connflushifnomem:                    d.Get("connflushifnomem").(string),
		Connflushthres:                      d.Get("connflushthres").(int),
		Delayedack:                          d.Get("delayedack").(int),
		Downstaterst:                        d.Get("downstaterst").(string),
		Initialcwnd:                         d.Get("initialcwnd").(int),
		Kaprobeupdatelastactivity:           d.Get("kaprobeupdatelastactivity").(string),
		Learnvsvrmss:                        d.Get("learnvsvrmss").(string),
		Limitedpersist:                      d.Get("limitedpersist").(string),
		Maxburst:                            d.Get("maxburst").(int),
		Maxdynserverprobes:                  d.Get("maxdynserverprobes").(int),
		Maxpktpermss:                        d.Get("maxpktpermss").(string),
		Maxsynackretx:                       d.Get("maxsynackretx").(int),
		Maxsynhold:                          d.Get("maxsynhold").(int),
		Maxsynholdperprobe:                  d.Get("maxsynholdperprobe").(int),
		Maxtimewaitconn:                     d.Get("maxtimewaitconn").(int),
		Minrto:                              d.Get("minrto").(int),
		Mptcpchecksum:                       d.Get("mptcpchecksum").(string),
		Mptcpclosemptcpsessiononlastsfclose: d.Get("mptcpclosemptcpsessiononlastsfclose").(string),
		Mptcpconcloseonpassivesf:            d.Get("mptcpconcloseonpassivesf").(string),
		Mptcpimmediatesfcloseonfin:          d.Get("mptcpimmediatesfcloseonfin").(string),
		Mptcpmaxpendingsf:                   d.Get("mptcpmaxpendingsf").(string),
		Mptcpmaxsf:                          d.Get("mptcpmaxsf").(int),
		Mptcppendingjointhreshold:           d.Get("mptcppendingjointhreshold").(string),
		Mptcprtostoswitchsf:                 d.Get("mptcprtostoswitchsf").(int),
		Mptcpsfreplacetimeout:               d.Get("mptcpsfreplacetimeout").(string),
		Mptcpsftimeout:                      d.Get("mptcpsftimeout").(string),
		Mptcpusebackupondss:                 d.Get("mptcpusebackupondss").(string),
		Msslearndelay:                       d.Get("msslearndelay").(int),
		Msslearninterval:                    d.Get("msslearninterval").(int),
		Nagle:                               d.Get("nagle").(string),
		Oooqsize:                            d.Get("oooqsize").(string),
		Pktperretx:                          d.Get("pktperretx").(int),
		Recvbuffsize:                        d.Get("recvbuffsize").(int),
		Sack:                                d.Get("sack").(string),
		Slowstartincr:                       d.Get("slowstartincr").(int),
		Synattackdetection:                  d.Get("synattackdetection").(string),
		Synholdfastgiveup:                   d.Get("synholdfastgiveup").(int),
		Tcpfastopencookietimeout:            d.Get("tcpfastopencookietimeout").(string),
		Tcpfintimeout:                       d.Get("tcpfintimeout").(int),
		Tcpmaxretries:                       d.Get("tcpmaxretries").(int),
		Ws:                                  d.Get("ws").(string),
		Wsval:                               d.Get("wsval").(string),
	}

	err := client.UpdateUnnamedResource(service.Nstcpparam.Type(), &nstcpparam)
	if err != nil {
		return err
	}

	d.SetId(nstcpparamName)

	err = readNstcpparamFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstcpparam but we can't read it ?? %s", nstcpparamName)
		return nil
	}
	return nil
}

func readNstcpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstcpparamFunc")
	client := meta.(*NetScalerNitroClient).client
	nstcpparamName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nstcpparam state %s", nstcpparamName)
	findParams := service.FindParams{
		ResourceType: "nstcpparam",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	// There is always a single entry
	data := dataArr[0]
	log.Printf("[DEBUG] citrixadc-provider: data read %v", data)

	d.Set("ackonpush", data["ackonpush"])
	d.Set("autosyncookietimeout", data["autosyncookietimeout"])
	d.Set("connflushifnomem", data["connflushifnomem"])
	d.Set("connflushthres", data["connflushthres"])
	d.Set("delayedack", data["delayedack"])
	d.Set("downstaterst", data["downstaterst"])
	d.Set("initialcwnd", data["initialcwnd"])
	d.Set("kaprobeupdatelastactivity", data["kaprobeupdatelastactivity"])
	d.Set("learnvsvrmss", data["learnvsvrmss"])
	d.Set("limitedpersist", data["limitedpersist"])
	d.Set("maxburst", data["maxburst"])
	d.Set("maxdynserverprobes", data["maxdynserverprobes"])
	d.Set("maxpktpermss", data["maxpktpermss"])
	d.Set("maxsynackretx", data["maxsynackretx"])
	d.Set("maxsynhold", data["maxsynhold"])
	d.Set("maxsynholdperprobe", data["maxsynholdperprobe"])
	d.Set("maxtimewaitconn", data["maxtimewaitconn"])
	d.Set("minrto", data["minrto"])
	d.Set("mptcpchecksum", data["mptcpchecksum"])
	d.Set("mptcpclosemptcpsessiononlastsfclose", data["mptcpclosemptcpsessiononlastsfclose"])
	d.Set("mptcpconcloseonpassivesf", data["mptcpconcloseonpassivesf"])
	d.Set("mptcpimmediatesfcloseonfin", data["mptcpimmediatesfcloseonfin"])
	d.Set("mptcpmaxpendingsf", data["mptcpmaxpendingsf"])
	d.Set("mptcpmaxsf", data["mptcpmaxsf"])
	d.Set("mptcppendingjointhreshold", data["mptcppendingjointhreshold"])
	d.Set("mptcprtostoswitchsf", data["mptcprtostoswitchsf"])
	d.Set("mptcpsfreplacetimeout", data["mptcpsfreplacetimeout"])
	d.Set("mptcpsftimeout", data["mptcpsftimeout"])
	d.Set("mptcpusebackupondss", data["mptcpusebackupondss"])
	d.Set("msslearndelay", data["msslearndelay"])
	d.Set("msslearninterval", data["msslearninterval"])
	d.Set("nagle", data["nagle"])
	d.Set("oooqsize", data["oooqsize"])
	d.Set("pktperretx", data["pktperretx"])
	d.Set("recvbuffsize", data["recvbuffsize"])
	d.Set("sack", data["sack"])
	d.Set("slowstartincr", data["slowstartincr"])
	d.Set("synattackdetection", data["synattackdetection"])
	d.Set("synholdfastgiveup", data["synholdfastgiveup"])
	d.Set("tcpfastopencookietimeout", data["tcpfastopencookietimeout"])
	d.Set("tcpfintimeout", data["tcpfintimeout"])
	d.Set("tcpmaxretries", data["tcpmaxretries"])
	d.Set("ws", data["ws"])
	d.Set("wsval", data["wsval"])

	return nil

}

func deleteNstcpparamFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstcpparamFunc")

	d.SetId("")

	return nil
}
