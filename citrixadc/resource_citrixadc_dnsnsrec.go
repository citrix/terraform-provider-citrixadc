package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/dns"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcDnsnsrec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnsnsrecFunc,
		Read:          readDnsnsrecFunc,
		Delete:        deleteDnsnsrecFunc,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nameserver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createDnsnsrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	domain := d.Get("domain").(string)
	nameserver := d.Get("nameserver").(string)

	dnsnsrecId := fmt.Sprintf("%s,%s", domain, nameserver)
	dnsnsrec := dns.Dnsnsrec{
		Domain:     d.Get("domain").(string),
		Nameserver: d.Get("nameserver").(string),
		Ttl:        d.Get("ttl").(int),
	}

	_, err := client.AddResource(netscaler.Dnsnsrec.Type(), "", &dnsnsrec)
	if err != nil {
		return err
	}

	d.SetId(dnsnsrecId)

	err = readDnsnsrecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnsnsrec but we can't read it ?? %s", dnsnsrecId)
		return nil
	}
	return nil
}

func readDnsnsrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnsrecId := d.Id()

	idSlice := strings.SplitN(dnsnsrecId, ",", 2)
	domain := idSlice[0]
	nameserver := idSlice[1]
	findParams := netscaler.FindParams{
		ResourceType: "dnsnsrec",
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	foundIndex := -1
	for i, v := range dataArr {
		if v["domain"] == domain && v["nameserver"] == nameserver {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams dnsnsrec not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing dnsnsrec state %s", dnsnsrecId)
		d.SetId("")
	}

	data := dataArr[foundIndex]

	d.Set("domain", data["domain"])
	d.Set("nameserver", data["nameserver"])
	d.Set("ttl", data["ttl"])

	return nil

}

func deleteDnsnsrecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnsnsrecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnsnsrecId := d.Id()

	idSlice := strings.SplitN(dnsnsrecId, ",", 2)

	domain := idSlice[0]
	nameserver := idSlice[1]
	argsMap := make(map[string]string)
	argsMap["nameserver"] = url.QueryEscape(nameserver)
	err := client.DeleteResourceWithArgsMap(netscaler.Dnsnsrec.Type(), domain, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
