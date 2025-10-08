package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcServer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createServerFunc,
		ReadContext:   readServerFunc,
		UpdateContext: updateServerFunc,
		DeleteContext: deleteServerFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"domainresolvenow": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"domainresolveretry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"querytype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"translationip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translationmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createServerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		Domain:             d.Get("domain").(string),
		Domainresolvenow:   d.Get("domainresolvenow").(bool),
		Domainresolveretry: d.Get("domainresolveretry").(int),
		Internal:           d.Get("internal").(bool),
		Ipaddress:          d.Get("ipaddress").(string),
		Ipv6address:        d.Get("ipv6address").(string),
		Name:               d.Get("name").(string),
		Querytype:          d.Get("querytype").(string),
		State:              d.Get("state").(string),
		Td:                 d.Get("td").(int),
		Translationip:      d.Get("translationip").(string),
		Translationmask:    d.Get("translationmask").(string),
	}

	_, err := client.AddResource(service.Server.Type(), serverName, &server)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(serverName)

	return readServerFunc(ctx, d, meta)
}

func readServerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] netscaler-provider:  In readServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading server state %s", serverName)
	data, err := client.FindResource(service.Server.Type(), serverName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing server state %s", serverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("comment", data["comment"])
	d.Set("domain", data["domain"])
	d.Set("domainresolvenow", data["domainresolvenow"])
	setToInt("domainresolveretry", d, data["domainresolveretry"])
	d.Set("internal", data["internal"])
	d.Set("ipaddress", data["ipaddress"])
	d.Set("ipv6address", data["ipv6address"])
	d.Set("name", data["name"])
	d.Set("querytype", data["querytype"])
	d.Set("state", data["state"])
	setToInt("td", d, data["td"])
	d.Set("translationip", data["translationip"])
	d.Set("translationmask", data["translationmask"])

	return nil

}

func updateServerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In updateServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Get("name").(string)

	server := basic.Server{
		Name: d.Get("name").(string),
	}

	stateChange := false

	hasChange := false

	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for server %s, starting update", serverName)
		server.Comment = d.Get("comment").(string)
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
	if d.HasChange("querytype") {
		log.Printf("[DEBUG]  netscaler-provider: Querytype has changed for server %s, starting update", serverName)
		server.Querytype = d.Get("querytype").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for server %s, starting update", serverName)
		stateChange = true
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
		_, err := client.UpdateResource(service.Server.Type(), serverName, &server)
		if err != nil {
			return diag.Errorf("Error updating server %s", serverName)
		}
	}
	if stateChange {
		err := doServerStateChange(d, client)
		if err != nil {
			return diag.Errorf("Error enabling/disabling server %s", serverName)
		}
	}
	return readServerFunc(ctx, d, meta)
}

func deleteServerFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  netscaler-provider: In deleteServerFunc")
	client := meta.(*NetScalerNitroClient).client
	serverName := d.Id()
	err := client.DeleteResource(service.Server.Type(), serverName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func doServerStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doServerStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	server := basic.Server{
		Name: d.Get("name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Server.Type(), server, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		server.Delay = d.Get("delay").(int)
		server.Graceful = d.Get("graceful").(string)
		err := client.ActOnResource(service.Server.Type(), server, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}
