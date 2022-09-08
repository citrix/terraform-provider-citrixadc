package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/network"
	
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"net/url"
	"log"
)

func resourceCitrixAdcChannel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createChannelFunc,
		Read:          readChannelFunc,
		Update:        updateChannelFunc,
		Delete:        deleteChannelFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"channel_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bandwidthhigh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bandwidthnormal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"conndistr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flowctl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"haheartbeat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hamonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifalias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ifnum": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"lamac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkredundancy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lrminthroughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"macdistr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tagall": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trunk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createChannelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createChannelFunc")
	client := meta.(*NetScalerNitroClient).client
	channelName := d.Get("channel_id").(string)

	channel := network.Channel{
		Bandwidthhigh:   d.Get("bandwidthhigh").(int),
		Bandwidthnormal: d.Get("bandwidthnormal").(int),
		Conndistr:       d.Get("conndistr").(string),
		Flowctl:         d.Get("flowctl").(string),
		Haheartbeat:     d.Get("haheartbeat").(string),
		Hamonitor:       d.Get("hamonitor").(string),
		Id:              d.Get("channel_id").(string),
		Ifalias:         d.Get("ifalias").(string),
		Ifnum:           toStringList(d.Get("ifnum").([]interface{})),
		Lamac:           d.Get("lamac").(string),
		Linkredundancy:  d.Get("linkredundancy").(string),
		Lrminthroughput: d.Get("lrminthroughput").(int),
		Macdistr:        d.Get("macdistr").(string),
		Mode:            d.Get("mode").(string),
		Mtu:             d.Get("mtu").(int),
		Speed:           d.Get("speed").(string),
		State:           d.Get("state").(string),
		Tagall:          d.Get("tagall").(string),
		Throughput:      d.Get("throughput").(int),
		Trunk:           d.Get("trunk").(string),
	}

	_, err := client.AddResource(service.Channel.Type(), channelName, &channel)
	if err != nil {
		return err
	}

	d.SetId(channelName)

	err = readChannelFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this channel but we can't read it ?? %s", channelName)
		return nil
	}
	return nil
}

func readChannelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readChannelFunc")
	client := meta.(*NetScalerNitroClient).client
	channelName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading channel state %s", channelName)
	data, err := client.FindResource(service.Channel.Type(), url.QueryEscape(url.QueryEscape(channelName)))
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing channel state %s", channelName)
		d.SetId("")
		return nil
	}

	d.Set("channel_id", data["id"])
	d.Set("bandwidthhigh", data["bandwidthhigh"])
	d.Set("bandwidthnormal", data["bandwidthnormal"])
	d.Set("conndistr", data["conndistr"])
	d.Set("flowctl", data["flowctl"])
	d.Set("haheartbeat", data["haheartbeat"])
	d.Set("hamonitor", data["hamonitor"])
	d.Set("channel_id", data["id"])
	d.Set("ifalias", data["ifalias"])
	d.Set("ifnum", data["ifnum"])
	d.Set("lamac", data["lamac"])
	d.Set("linkredundancy", data["linkredundancy"])
	d.Set("lrminthroughput", data["lrminthroughput"])
	d.Set("macdistr", data["macdistr"])
	//d.Set("mode", data["mode"])
	d.Set("mtu", data["mtu"])
	//d.Set("speed", data["speed"])
	d.Set("state", data["state"])
	d.Set("tagall", data["tagall"])
	d.Set("throughput", data["throughput"])
	d.Set("trunk", data["trunk"])

	return nil

}

func updateChannelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateChannelFunc")
	client := meta.(*NetScalerNitroClient).client
	channelName := d.Get("channel_id").(string)

	channel := network.Channel{
		Id: d.Get("channel_id").(string),
	}
	hasChange := false
	if d.HasChange("bandwidthhigh") {
		log.Printf("[DEBUG]  citrixadc-provider: Bandwidthhigh has changed for channel %s, starting update", channelName)
		channel.Bandwidthhigh = d.Get("bandwidthhigh").(int)
		hasChange = true
	}
	if d.HasChange("bandwidthnormal") {
		log.Printf("[DEBUG]  citrixadc-provider: Bandwidthnormal has changed for channel %s, starting update", channelName)
		channel.Bandwidthnormal = d.Get("bandwidthnormal").(int)
		hasChange = true
	}
	if d.HasChange("conndistr") {
		log.Printf("[DEBUG]  citrixadc-provider: Conndistr has changed for channel %s, starting update", channelName)
		channel.Conndistr = d.Get("conndistr").(string)
		hasChange = true
	}
	if d.HasChange("flowctl") {
		log.Printf("[DEBUG]  citrixadc-provider: Flowctl has changed for channel %s, starting update", channelName)
		channel.Flowctl = d.Get("flowctl").(string)
		hasChange = true
	}
	if d.HasChange("haheartbeat") {
		log.Printf("[DEBUG]  citrixadc-provider: Haheartbeat has changed for channel %s, starting update", channelName)
		channel.Haheartbeat = d.Get("haheartbeat").(string)
		hasChange = true
	}
	if d.HasChange("hamonitor") {
		log.Printf("[DEBUG]  citrixadc-provider: Hamonitor has changed for channel %s, starting update", channelName)
		channel.Hamonitor = d.Get("hamonitor").(string)
		hasChange = true
	}
	if d.HasChange("ifalias") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifalias has changed for channel %s, starting update", channelName)
		channel.Ifalias = d.Get("ifalias").(string)
		hasChange = true
	}
	if d.HasChange("ifnum") {
		log.Printf("[DEBUG]  citrixadc-provider: Ifnum has changed for channel %s, starting update", channelName)
		channel.Ifnum = toStringList(d.Get("ifnum").([]interface{}))
		hasChange = true
	}
	if d.HasChange("lamac") {
		log.Printf("[DEBUG]  citrixadc-provider: Lamac has changed for channel %s, starting update", channelName)
		channel.Lamac = d.Get("lamac").(string)
		hasChange = true
	}
	if d.HasChange("linkredundancy") {
		log.Printf("[DEBUG]  citrixadc-provider: Linkredundancy has changed for channel %s, starting update", channelName)
		channel.Linkredundancy = d.Get("linkredundancy").(string)
		hasChange = true
	}
	if d.HasChange("lrminthroughput") {
		log.Printf("[DEBUG]  citrixadc-provider: Lrminthroughput has changed for channel %s, starting update", channelName)
		channel.Lrminthroughput = d.Get("lrminthroughput").(int)
		hasChange = true
	}
	if d.HasChange("macdistr") {
		log.Printf("[DEBUG]  citrixadc-provider: Macdistr has changed for channel %s, starting update", channelName)
		channel.Macdistr = d.Get("macdistr").(string)
		hasChange = true
	}
	if d.HasChange("mode") {
		log.Printf("[DEBUG]  citrixadc-provider: Mode has changed for channel %s, starting update", channelName)
		channel.Mode = d.Get("mode").(string)
		hasChange = true
	}
	if d.HasChange("mtu") {
		log.Printf("[DEBUG]  citrixadc-provider: Mtu has changed for channel %s, starting update", channelName)
		channel.Mtu = d.Get("mtu").(int)
		hasChange = true
	}
	if d.HasChange("speed") {
		log.Printf("[DEBUG]  citrixadc-provider: Speed has changed for channel %s, starting update", channelName)
		channel.Speed = d.Get("speed").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for channel %s, starting update", channelName)
		channel.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("tagall") {
		log.Printf("[DEBUG]  citrixadc-provider: Tagall has changed for channel %s, starting update", channelName)
		channel.Tagall = d.Get("tagall").(string)
		hasChange = true
	}
	if d.HasChange("throughput") {
		log.Printf("[DEBUG]  citrixadc-provider: Throughput has changed for channel %s, starting update", channelName)
		channel.Throughput = d.Get("throughput").(int)
		hasChange = true
	}
	if d.HasChange("trunk") {
		log.Printf("[DEBUG]  citrixadc-provider: Trunk has changed for channel %s, starting update", channelName)
		channel.Trunk = d.Get("trunk").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Channel.Type(), &channel)
		if err != nil {
			return fmt.Errorf("Error updating channel %s", channelName)
		}
	}
	return readChannelFunc(d, meta)
}

func deleteChannelFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteChannelFunc")
	client := meta.(*NetScalerNitroClient).client
	channelName := d.Id()
	err := client.DeleteResource(service.Channel.Type(), url.QueryEscape(url.QueryEscape(channelName)))
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
