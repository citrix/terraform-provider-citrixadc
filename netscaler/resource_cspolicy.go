package netscaler

import (
	"github.com/chiradeep/go-nitro/config/cs"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerCspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCspolicyFunc,
		Read:          readCspolicyFunc,
		Update:        updateCspolicyFunc,
		Delete:        deleteCspolicyFunc,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"csvserver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func createCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client

	csvserver := d.Get("csvserver").(string)
	targetlbvserver, lbok := d.GetOk("targetlbvserver")
	priority, pok := d.GetOk("priority")
	_, aok := d.GetOk("action")

	if lbok && !pok {
		return fmt.Errorf("Priority needs to be specified if target lb vserver is specified")
	}
	if !lbok && pok {
		return fmt.Errorf("Priority needs to be specified if target lb vserver is specified")
	}
	if !lbok && !aok {
		return fmt.Errorf("Action or targetlbvserver needs to be specified")
	}

	var cspolicyName string
	if v, ok := d.GetOk("policyname"); ok {
		cspolicyName = v.(string)
	} else {
		cspolicyName = resource.PrefixedUniqueId("tf-cspolicy-")
		d.Set("name", cspolicyName)
	}
	cspolicy := cs.Cspolicy{
		Action:     d.Get("action").(string),
		Domain:     d.Get("domain").(string),
		Logaction:  d.Get("logaction").(string),
		Newname:    d.Get("newname").(string),
		Policyname: d.Get("policyname").(string),
		Rule:       d.Get("rule").(string),
		Url:        d.Get("url").(string),
	}

	_, err := client.AddResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
	if err != nil {
		return err
	}

	d.SetId(cspolicyName)

	binding := cs.Csvservercspolicybinding{
		Name:            csvserver,
		Policyname:      cspolicyName,
		Targetlbvserver: targetlbvserver.(string),
		Priority:        priority.(int),
	}

	if !lbok {
		binding = cs.Csvservercspolicybinding{
			Name:       csvserver,
			Policyname: cspolicyName,
		}
	}

	err = client.BindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, &binding)
	if err != nil {
		log.Printf("Failed to bind cspolicy to Csvserver")
		return err
	}
	err = readCspolicyFunc(d, meta)
	if err != nil {
		log.Printf("?? we just created this cspolicy but we can't read it ?? %s", cspolicyName)
		return nil
	}
	return nil
}

func readCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	log.Printf("Reading cspolicy state %s", cspolicyName)
	data, err := client.FindResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		log.Printf("Clearing cspolicy state %s", cspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("domain", data["domain"])
	d.Set("logaction", data["logaction"])
	d.Set("rule", data["rule"])
	d.Set("url", data["url"])

	return nil

}

func updateCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] In update func")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Get("policyname").(string)
	csvserver := d.Get("csvserver").(string)

	cspolicy := cs.Cspolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	lbvserverChanged := false

	if d.HasChange("action") {
		log.Printf("[DEBUG] Action has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG] Domain has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG] Logaction has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG] Newname has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG] Rule has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG] Url has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Url = d.Get("url").(string)
		hasChange = true
	}

	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG] targetlbvserver has changed for cspolicy %s, starting update", cspolicyName)
		lbvserverChanged = true
	}

	if lbvserverChanged {
		//Binding has to be updated
		//First we unbind from cs vserver
		err := client.UnbindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, "policyname")
		if err != nil {
			return fmt.Errorf("Error unbinding cspolicy from csvserver %s", cspolicyName)
		}
		log.Printf("[DEBUG] cspolicy has been unbound from csvserver for cspolicy %s ", cspolicyName)
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
		if err != nil {
			return fmt.Errorf("Error updating cspolicy %s", cspolicyName)
		}
		log.Printf("[DEBUG] cspolicy has been updated  cspolicy %s ", cspolicyName)
	}

	if lbvserverChanged {
		//Binding has to be updated
		//rebind
		targetlbvserver, lbok := d.GetOk("targetlbvserver")
		priority, pok := d.GetOk("priority")

		if !pok && lbok {
			return fmt.Errorf("Need to specify priority if lbvserver is specified")
		}

		binding := cs.Csvservercspolicybinding{
			Name:            csvserver,
			Policyname:      cspolicyName,
			Targetlbvserver: targetlbvserver.(string),
			Priority:        priority.(int),
		}
		err := client.BindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, &binding)
		if err != nil {
			log.Printf("Failed to bind new cspolicy to Csvserver")
			return err
		}
		log.Printf("[DEBUG] cspolicy has been bound to csvserver  cspolicy %s csvserver %s", cspolicyName, csvserver)
	}

	return readCspolicyFunc(d, meta)
}

func deleteCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	csvserver := d.Get("csvserver").(string)

	//First we unbind from cs vserver if necessary
	err := client.UnbindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, "policyname")
	if err != nil {
		return fmt.Errorf("Error unbinding cspolicy from csvserver %s", cspolicyName)
	}
	err = client.DeleteResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		return fmt.Errorf("Error  deleting cspolicy %s", cspolicyName)
	}

	d.SetId("")

	return nil
}
