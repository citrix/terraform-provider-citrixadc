package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNstimeout() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNstimeoutFunc,
		Read:          readNstimeoutFunc,
		Update:        updateNstimeoutFunc,
		Delete:        deleteNstimeoutFunc, // Thought nstimeout resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
			"anyclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anyserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anytcpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"anytcpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"client": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"halfclose": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"httpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"newconnidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nontcpzombie": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reducedfintimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reducedrsttimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"server": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpclient": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpserver": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zombie": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNstimeoutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client
	nstimeoutName := resource.PrefixedUniqueId("tf-nstimeout-")

	nstimeout := ns.Nstimeout{
		Anyclient:          d.Get("anyclient").(int),
		Anyserver:          d.Get("anyserver").(int),
		Anytcpclient:       d.Get("anytcpclient").(int),
		Anytcpserver:       d.Get("anytcpserver").(int),
		Client:             d.Get("client").(int),
		Halfclose:          d.Get("halfclose").(int),
		Httpclient:         d.Get("httpclient").(int),
		Httpserver:         d.Get("httpserver").(int),
		Newconnidletimeout: d.Get("newconnidletimeout").(int),
		Nontcpzombie:       d.Get("nontcpzombie").(int),
		Reducedfintimeout:  d.Get("reducedfintimeout").(int),
		Reducedrsttimeout:  d.Get("reducedrsttimeout").(int),
		Server:             d.Get("server").(int),
		Tcpclient:          d.Get("tcpclient").(int),
		Tcpserver:          d.Get("tcpserver").(int),
		Zombie:             d.Get("zombie").(int),
	}

	err := client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
	if err != nil {
		return err
	}

	d.SetId(nstimeoutName)

	err = readNstimeoutFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nstimeout but we can't read it ??")
		return nil
	}
	return nil
}

func readNstimeoutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading nstimeout state")
	data, err := client.FindResource(service.Nstimeout.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nstimeout state")
		d.SetId("")
		return nil
	}
	d.Set("anyclient", data["anyclient"])
	d.Set("anyserver", data["anyserver"])
	d.Set("anytcpclient", data["anytcpclient"])
	d.Set("anytcpserver", data["anytcpserver"])
	d.Set("client", data["client"])
	d.Set("halfclose", data["halfclose"])
	d.Set("httpclient", data["httpclient"])
	d.Set("httpserver", data["httpserver"])
	d.Set("newconnidletimeout", data["newconnidletimeout"])
	d.Set("nontcpzombie", data["nontcpzombie"])
	d.Set("reducedfintimeout", data["reducedfintimeout"])
	d.Set("reducedrsttimeout", data["reducedrsttimeout"])
	d.Set("server", data["server"])
	d.Set("tcpclient", data["tcpclient"])
	d.Set("tcpserver", data["tcpserver"])
	d.Set("zombie", data["zombie"])

	return nil

}

func updateNstimeoutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNstimeoutFunc")
	client := meta.(*NetScalerNitroClient).client

	nstimeout := ns.Nstimeout{}

	hasChange := false
	if d.HasChange("anyclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Anyclient has changed for nstimeout, starting update")
		nstimeout.Anyclient = d.Get("anyclient").(int)
		hasChange = true
	}
	if d.HasChange("anyserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Anyserver has changed for nstimeout, starting update")
		nstimeout.Anyserver = d.Get("anyserver").(int)
		hasChange = true
	}
	if d.HasChange("anytcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Anytcpclient has changed for nstimeout, starting update")
		nstimeout.Anytcpclient = d.Get("anytcpclient").(int)
		hasChange = true
	}
	if d.HasChange("anytcpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Anytcpserver has changed for nstimeout, starting update")
		nstimeout.Anytcpserver = d.Get("anytcpserver").(int)
		hasChange = true
	}
	if d.HasChange("client") {
		log.Printf("[DEBUG]  citrixadc-provider: Client has changed for nstimeout, starting update")
		nstimeout.Client = d.Get("client").(int)
		hasChange = true
	}
	if d.HasChange("halfclose") {
		log.Printf("[DEBUG]  citrixadc-provider: Halfclose has changed for nstimeout, starting update")
		nstimeout.Halfclose = d.Get("halfclose").(int)
		hasChange = true
	}
	if d.HasChange("httpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpclient has changed for nstimeout, starting update")
		nstimeout.Httpclient = d.Get("httpclient").(int)
		hasChange = true
	}
	if d.HasChange("httpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpserver has changed for nstimeout, starting update")
		nstimeout.Httpserver = d.Get("httpserver").(int)
		hasChange = true
	}
	if d.HasChange("newconnidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Newconnidletimeout has changed for nstimeout, starting update")
		nstimeout.Newconnidletimeout = d.Get("newconnidletimeout").(int)
		hasChange = true
	}
	if d.HasChange("nontcpzombie") {
		log.Printf("[DEBUG]  citrixadc-provider: Nontcpzombie has changed for nstimeout, starting update")
		nstimeout.Nontcpzombie = d.Get("nontcpzombie").(int)
		hasChange = true
	}
	if d.HasChange("reducedfintimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reducedfintimeout has changed for nstimeout, starting update")
		nstimeout.Reducedfintimeout = d.Get("reducedfintimeout").(int)
		hasChange = true
	}
	if d.HasChange("reducedrsttimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Reducedrsttimeout has changed for nstimeout, starting update")
		nstimeout.Reducedrsttimeout = d.Get("reducedrsttimeout").(int)
		hasChange = true
	}
	if d.HasChange("server") {
		log.Printf("[DEBUG]  citrixadc-provider: Server has changed for nstimeout, starting update")
		nstimeout.Server = d.Get("server").(int)
		hasChange = true
	}
	if d.HasChange("tcpclient") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpclient has changed for nstimeout, starting update")
		nstimeout.Tcpclient = d.Get("tcpclient").(int)
		hasChange = true
	}
	if d.HasChange("tcpserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Tcpserver has changed for nstimeout, starting update")
		nstimeout.Tcpserver = d.Get("tcpserver").(int)
		hasChange = true
	}
	if d.HasChange("zombie") {
		log.Printf("[DEBUG]  citrixadc-provider: Zombie has changed for nstimeout, starting update")
		nstimeout.Zombie = d.Get("zombie").(int)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
		if err != nil {
			return fmt.Errorf("Error updating nstimeout")
		}
	}
	return readNstimeoutFunc(d, meta)
}

func deleteNstimeoutFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNstimeoutFunc")
	// nstimeout do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
