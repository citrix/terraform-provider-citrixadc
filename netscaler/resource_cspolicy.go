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
	log.Printf("[DEBUG] netscaler-provider: In createCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client

	csvserver := d.Get("csvserver").(string)
	targetlbvserver, lbok := d.GetOk("targetlbvserver")
	priority, pok := d.GetOk("priority")
	action, aok := d.GetOk("action")
	_, dok := d.GetOk("domain")
	_, uok := d.GetOk("url")
	_, rok := d.GetOk("rule")

	if lbok && rok && !pok {
		return fmt.Errorf("[ERROR] netscaler-provider: Priority needs to be specified if target lb vserver and rule is specified")
	}
	if !lbok && pok {
		return fmt.Errorf("[ERROR] netscaler-provider: Priority needs to be specified if target lb vserver is not specified")
	}
	if !lbok && !aok {
		return fmt.Errorf("[ERROR] netscaler-provider: Action or targetlbvserver needs to be specified")
	}
	if aok {
		actionExists := client.ResourceExists(netscaler.Csaction.Type(), action.(string))
		if !actionExists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified Action %s does not exist", action.(string))
		}
		if !rok {
			return fmt.Errorf("[ERROR] netscaler-provider: Action  %s specified without rule", action.(string))
		}
		if dok || uok {
			return fmt.Errorf("[ERROR] netscaler-provider: Cannot specify url or domain when action  %s is specified", action.(string))
		}
	}
	if uok && dok {
		return fmt.Errorf("[ERROR] netscaler-provider: Cannot specify both url and domain ")
	}
	if rok && (uok || dok) {
		return fmt.Errorf("[ERROR] netscaler-provider: Cannot specify both rule and domain or url ")
	}
	if (uok || dok) && pok {
		return fmt.Errorf("[ERROR] netscaler-provider: Cannot specify both priority and domain or url ")
	}

	var cspolicyName string
	if v, ok := d.GetOk("policyname"); ok {
		cspolicyName = v.(string)
	} else {
		cspolicyName = resource.PrefixedUniqueId("tf-cspolicy-")
		d.Set("name", cspolicyName)
	}
	cspolicy := cs.Cspolicy{
		Policyname: d.Get("policyname").(string),
		Action:     d.Get("action").(string),
		Domain:     d.Get("domain").(string),
		Logaction:  d.Get("logaction").(string),
		Rule:       d.Get("rule").(string),
		Url:        d.Get("url").(string),
	}

	_, err := client.AddResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
	if err != nil {
		return err
	}

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
		d.SetId("")
		err2 := client.DeleteResource(netscaler.Cspolicy.Type(), cspolicyName)
		if err2 != nil {
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to undo add cspolicy after bind cspolicy %s to Csvserver failed", cspolicyName, err2)
		}
		return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind cspolicy %s to Csvserver", cspolicyName, err)
	}
	d.SetId(cspolicyName)
	err = readCspolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider:  ?? we just created this cspolicy but we can't read it ?? %s", cspolicyName)
		return nil
	}
	return nil
}

func readCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In readCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	log.Printf("[DEBUG] netscaler-provider:  Reading cspolicy state %s", cspolicyName)
	data, err := client.FindResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider:  Clearing cspolicy state %s", cspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("domain", data["domain"])
	d.Set("logaction", data["logaction"])
	d.Set("rule", data["rule"])
	d.Set("url", data["url"])

	//read the csvserver binding and update
	bindings, err := client.FindAllBoundResources(netscaler.Cspolicy.Type(), cspolicyName, netscaler.Csvserver.Type())
	if err != nil {
		log.Printf("[WARN] netscaler-provider: cspolicy binding to csvserver error %s", cspolicyName)
		return nil
	}
	var boundCsvserver string
	for _, binding := range bindings {
		csv, ok := binding["domain"]
		if ok {
			boundCsvserver = csv.(string)
			break
		}
	}
	d.Set("csvserver", boundCsvserver)

	return nil

}

func updateCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In updateCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Get("policyname").(string)
	csvserver := d.Get("csvserver").(string)

	cspolicy := cs.Cspolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	lbvserverChanged := false
	priorityChanged := false

	if d.HasChange("action") {
		log.Printf("[DEBUG] netscaler-provider: Action has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG] netscaler-provider: Domain has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG] netscaler-provider: Logaction has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG] netscaler-provider: Newname has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG] netscaler-provider: Rule has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("url") {
		log.Printf("[DEBUG] netscaler-provider: Url has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Url = d.Get("url").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG] netscaler-provider: Priority has changed for cspolicy %s, starting update", cspolicyName)
		priorityChanged = true
	}

	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG] netscaler-provider: targetlbvserver has changed for cspolicy %s, starting update", cspolicyName)
		lbvserverChanged = true
	}

	if lbvserverChanged || priorityChanged {
		//Binding has to be updated
		//First we unbind from cs vserver
		err := client.UnbindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, "policyname")
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding cspolicy from csvserver %s", cspolicyName)
		}
		log.Printf("[DEBUG] netscaler-provider: cspolicy has been unbound from csvserver for cspolicy %s ", cspolicyName)
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Cspolicy.Type(), cspolicyName, &cspolicy)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error updating cspolicy %s", cspolicyName)
		}
		log.Printf("[DEBUG] netscaler-provider: cspolicy has been updated  cspolicy %s ", cspolicyName)
	}

	if lbvserverChanged || priorityChanged {
		//Binding has to be updated
		//rebind
		targetlbvserver, lbok := d.GetOk("targetlbvserver")
		priority, pok := d.GetOk("priority")

		if !pok && lbok {
			return fmt.Errorf("[ERROR] netscaler-provider: Need to specify priority if lbvserver is specified")
		}

		binding := cs.Csvservercspolicybinding{
			Name:            csvserver,
			Policyname:      cspolicyName,
			Targetlbvserver: targetlbvserver.(string),
			Priority:        priority.(int),
		}
		err := client.BindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, &binding)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Failed to bind new cspolicy to Csvserver")
		}
		log.Printf("[DEBUG] netscaler-provider: cspolicy has been bound to csvserver  cspolicy %s csvserver %s", cspolicyName, csvserver)
	}

	return readCspolicyFunc(d, meta)
}

func deleteCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In deleteCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	csvserver := d.Get("csvserver").(string)

	//First we unbind from cs vserver if necessary
	err := client.UnbindResource(netscaler.Csvserver.Type(), csvserver, netscaler.Cspolicy.Type(), cspolicyName, "policyname")
	if err != nil {
		return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding cspolicy from csvserver %s", cspolicyName)
	}
	err = client.DeleteResource(netscaler.Cspolicy.Type(), cspolicyName)
	if err != nil {
		return fmt.Errorf("[ERROR] netscaler-provider: Error  deleting cspolicy %s", cspolicyName)
	}

	d.SetId("")

	return nil
}
