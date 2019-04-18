package netscaler

import (
	"github.com/chiradeep/go-nitro/config/basic"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerServer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServerFunc,
		Read:          readServerFunc,
		Update:        updateServerFunc,
		Delete:        deleteServerFunc,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domainresolvenow": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"domainresolveretry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"translationip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translationmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createServerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createServerFunc")
	client := meta.(*NetScalerNitroClient).client
	var serverName string
	if v, ok := d.GetOk("name"); ok {
		serverName = v.(string)
	} else {
		serverName = resource.PrefixedUniqueId("tf-server-")
		d.Set("name", serverName)
	}
	server := basic.Server{
		Comment:            d.Get("comment").(string),
		Delay:              d.Get("delay").(int),
		Domain:             d.Get("domain").(string),
		Domainresolvenow:   d.Get("domainresolvenow").(bool),
		Domainresolveretry: d.Get("domainresolveretry").(int),
		Graceful:           d.Get("graceful").(string),
		Internal:           d.Get("internal").(bool),
		Ipaddress:          d.Get("ipaddress").(string),
		Ipv6address:        d.Get("ipv6address").(string),
		Name:               d.Get("name").(string),
		Newname:            d.Get("newname").(string),
		State:              d.Get("state").(string),
		Td:                 d.Get("td").(int),
		Translationip:      d.Get("translationip").(string),
		Translationmask:    d.Get("translationmask").(string),
	}

	_, err := client.AddResource(netscaler.Server.Type(), serverName, &server)
	if err != nil {
		return err
	}

	d.SetId(serverName)

	err = readServerFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this server but we can't read it ?? %s", serverName)
		return nil
	}
	return nil
}

func readServerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading server state %s", serverName)
	data, err := client.FindResource(netscaler.Server.Type(), serverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing server state %s", serverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("delay", data["delay"])
	d.Set("domain", data["domain"])
	d.Set("domainresolvenow", data["domainresolvenow"])
	d.Set("domainresolveretry", data["domainresolveretry"])
	d.Set("graceful", data["graceful"])
	d.Set("internal", data["internal"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("ipv6address", data["ipv6address"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("state", data["state"])
	d.Set("td", data["td"])
	d.Set("translationip", data["translationip"])
	d.Set("translationmask", data["translationmask"])

	return nil

}

func updateServerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Get("name").(string)

	server := basic.Server{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for server %s, starting update", serverName)
		server.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  netscaler-provider: Delay has changed for server %s, starting update", serverName)
		server.Delay = d.Get("delay").(int)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  netscaler-provider: Domain has changed for server %s, starting update", serverName)
		server.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("domainresolvenow") {
		log.Printf("[DEBUG]  netscaler-provider: Domainresolvenow has changed for server %s, starting update", serverName)
		server.Domainresolvenow = d.Get("domainresolvenow").(bool)
		hasChange = true
	}
	if d.HasChange("domainresolveretry") {
		log.Printf("[DEBUG]  netscaler-provider: Domainresolveretry has changed for server %s, starting update", serverName)
		server.Domainresolveretry = d.Get("domainresolveretry").(int)
		hasChange = true
	}
	if d.HasChange("graceful") {
		log.Printf("[DEBUG]  netscaler-provider: Graceful has changed for server %s, starting update", serverName)
		server.Graceful = d.Get("graceful").(string)
		hasChange = true
	}
	if d.HasChange("internal") {
		log.Printf("[DEBUG]  netscaler-provider: Internal has changed for server %s, starting update", serverName)
		server.Internal = d.Get("internal").(bool)
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  netscaler-provider: Ipaddress has changed for server %s, starting update", serverName)
		server.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("ipv6address") {
		log.Printf("[DEBUG]  netscaler-provider: Ipv6address has changed for server %s, starting update", serverName)
		server.Ipv6address = d.Get("ipv6address").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  netscaler-provider: Name has changed for server %s, starting update", serverName)
		server.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  netscaler-provider: Newname has changed for server %s, starting update", serverName)
		server.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for server %s, starting update", serverName)
		server.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for server %s, starting update", serverName)
		server.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("translationip") {
		log.Printf("[DEBUG]  netscaler-provider: Translationip has changed for server %s, starting update", serverName)
		server.Translationip = d.Get("translationip").(string)
		hasChange = true
	}
	if d.HasChange("translationmask") {
		log.Printf("[DEBUG]  netscaler-provider: Translationmask has changed for server %s, starting update", serverName)
		server.Translationmask = d.Get("translationmask").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Server.Type(), serverName, &server)
		if err != nil {
			return fmt.Errorf("Error updating server %s", serverName)
		}
	}
	return readServerFunc(d, meta)
}

func deleteServerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Id()
	err := client.DeleteResource(netscaler.Server.Type(), serverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
