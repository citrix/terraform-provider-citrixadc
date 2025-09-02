package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCspolicyFunc,
		Read:          readCspolicyFunc,
		Update:        updateCspolicyFunc,
		Delete:        deleteCspolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"csvserver": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"targetlbvserver": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"forcenew_id_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"priority": {
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
	_, rok := d.GetOk("rule")

	if aok {
		actionExists := client.ResourceExists(service.Csaction.Type(), action.(string))
		if !actionExists {
			return fmt.Errorf("[ERROR] netscaler-provider: Specified Action %s does not exist", action.(string))
		}
		if !rok {
			return fmt.Errorf("[ERROR] netscaler-provider: Action  %s specified without rule", action.(string))
		}
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
		Logaction:  d.Get("logaction").(string),
		Rule:       d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Cspolicy.Type(), cspolicyName, &cspolicy)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("csvserver"); ok {
		binding := cs.Csvserverpolicybinding{
			Name:       csvserver,
			Policyname: cspolicyName,
		}

		if pok {
			binding.Priority = uint32(priority.(int))
		}

		if lbok {
			binding.Targetlbvserver = targetlbvserver.(string)
		}

		err = client.BindResource(service.Csvserver.Type(), csvserver, service.Cspolicy.Type(), cspolicyName, &binding)
		if err != nil {
			d.SetId("")
			err2 := client.DeleteResource(service.Cspolicy.Type(), cspolicyName)
			if err2 != nil {
				return fmt.Errorf("[ERROR] netscaler-provider:  Failed to undo add cspolicy after bind cspolicy %s to Csvserver failed err=%v", cspolicyName, err2)
			}
			return fmt.Errorf("[ERROR] netscaler-provider:  Failed to bind cspolicy %s to Csvserver, err=%v", cspolicyName, err)
		}
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
	data, err := client.FindResource(service.Cspolicy.Type(), cspolicyName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider:  Clearing cspolicy state %s", cspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("policyname", data["policyname"])
	d.Set("action", data["action"])
	d.Set("logaction", data["logaction"])
	d.Set("rule", data["rule"])

	//read the csvserver binding and update
	if _, ok := d.GetOk("csvserver"); ok {
		bindings, err := client.FindAllBoundResources(service.Cspolicy.Type(), cspolicyName, service.Csvserver.Type())
		if err != nil {
			log.Printf("[WARN] netscaler-provider: cspolicy binding to csvserver error %s", cspolicyName)
			return nil
		}
		var boundCsvserver string
		for _, binding := range bindings {
			log.Printf("[TRACE] netscaler-provider: csvserver_cspolicy binding %v", binding)
			var ok bool
			var csv interface{}
			if _, ok = binding["boundto"]; ok {
				csv = binding["boundto"]
			}
			if ok {
				boundCsvserver = csv.(string)
				break
			}
		}
		log.Printf("[TRACE] netscaler-provider: boundCsvserver %v", boundCsvserver)
		d.Set("csvserver", boundCsvserver)
	}

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
	if d.HasChange("logaction") {
		log.Printf("[DEBUG] netscaler-provider: Logaction has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG] netscaler-provider: Rule has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG] netscaler-provider: Priority has changed for cspolicy %s, starting update", cspolicyName)
		priorityChanged = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG] netscaler-provider: Newname has changed for cspolicy %s, starting update", cspolicyName)
		cspolicy.Newname = d.Get("newname").(string)
		d.SetId(cspolicy.Newname)
		hasChange = true
	}
	if d.HasChange("targetlbvserver") {
		log.Printf("[DEBUG] netscaler-provider: targetlbvserver has changed for cspolicy %s, starting update", cspolicyName)
		lbvserverChanged = true
	}
	if _, ok := d.GetOk("csvserver"); ok {
		if lbvserverChanged || priorityChanged {
			//Binding has to be updated
			//First we unbind from cs vserver
			err := client.UnbindResource(service.Csvserver.Type(), csvserver, service.Cspolicy.Type(), cspolicyName, "policyname")
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding cspolicy from csvserver %s", cspolicyName)
			}
			log.Printf("[DEBUG] netscaler-provider: cspolicy has been unbound from csvserver for cspolicy %s ", cspolicyName)
		}
	}

	if hasChange {
		_, err := client.UpdateResource(service.Cspolicy.Type(), cspolicyName, &cspolicy)
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error updating cspolicy %s", cspolicyName)
		}
		log.Printf("[DEBUG] netscaler-provider: cspolicy has been updated  cspolicy %s ", cspolicyName)
	}

	if _, ok := d.GetOk("csvserver"); ok {
		if lbvserverChanged || priorityChanged {
			//Binding has to be updated
			//rebind
			targetlbvserver, lbok := d.GetOk("targetlbvserver")
			priority, pok := d.GetOk("priority")
			cspolicyName = d.Id()

			if !pok && lbok {
				return fmt.Errorf("[ERROR] netscaler-provider: Need to specify priority if lbvserver is specified")
			}

			binding := cs.Csvserverpolicybinding{
				Name:            csvserver,
				Policyname:      cspolicyName,
				Targetlbvserver: targetlbvserver.(string),
				Priority:        uint32(priority.(int)),
			}
			err := client.BindResource(service.Csvserver.Type(), csvserver, service.Cspolicy.Type(), cspolicyName, &binding)
			if err != nil {
				return fmt.Errorf("[ERROR] netscaler-provider: Failed to bind new cspolicy to Csvserver")
			}
			log.Printf("[DEBUG] netscaler-provider: cspolicy has been bound to csvserver  cspolicy %s csvserver %s", cspolicyName, csvserver)
		}
	}

	return readCspolicyFunc(d, meta)
}

func deleteCspolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In deleteCspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	cspolicyName := d.Id()
	csvserver := d.Get("csvserver").(string)

	//First we unbind from cs vserver if necessary
	if _, ok := d.GetOk("csvserver"); ok {
		err := client.UnbindResource(service.Csvserver.Type(), csvserver, service.Cspolicy.Type(), cspolicyName, "policyname")
		if err != nil {
			return fmt.Errorf("[ERROR] netscaler-provider: Error unbinding cspolicy \"%s\" from csvserver \"%v\": %v ", cspolicyName, csvserver, err.Error())
		}
	}
	err := client.DeleteResource(service.Cspolicy.Type(), cspolicyName)
	if err != nil {
		return fmt.Errorf("[ERROR] netscaler-provider: Error  deleting cspolicy %s", cspolicyName)
	}

	d.SetId("")

	return nil
}
