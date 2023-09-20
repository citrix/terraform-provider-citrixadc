package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcLbgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbgroupFunc,
		Read:          readLbgroupFunc,
		Update:        updateLbgroupFunc,
		Delete:        deleteLbgroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"persistencetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistencebackup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backuppersistencetimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"persistmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookiename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"v6persistmasklen": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cookiedomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usevserverpersistency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLbgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createLbgroupFunc")
	client := meta.(*NetScalerNitroClient).client

	LbgroupName := d.Get("name").(string)
	Lbgroup := lb.Lbgroup{
		Name:                     d.Get("name").(string),
		Persistencetype:          d.Get("persistencetype").(string),
		Persistencebackup:        d.Get("persistencebackup").(string),
		Backuppersistencetimeout: d.Get("backuppersistencetimeout").(int),
		Persistmask:              d.Get("persistmask").(string),
		Cookiename:               d.Get("cookiename").(string),
		V6persistmasklen:         d.Get("v6persistmasklen").(int),
		Cookiedomain:             d.Get("cookiedomain").(string),
		Timeout:                  d.Get("timeout").(int),
		Rule:                     d.Get("rule").(string),
		Usevserverpersistency:    d.Get("usevserverpersistency").(string),
	}

	_, err := client.AddResource("lbgroup", LbgroupName, &Lbgroup)
	if err != nil {
		return err
	}

	d.SetId(LbgroupName)

	err = readLbgroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this Lbgroup but we can't read it ?? %s", LbgroupName)
		return nil
	}
	return nil
}

func readLbgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readLbgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	LbgroupName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading Lbgroup state %s", LbgroupName)
	data, err := client.FindResource("lbgroup", LbgroupName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing Lbgroup state %s", LbgroupName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("persistencetype", data["persistencetype"])
	d.Set("persistencebackup", data["persistencebackup"])
	d.Set("backuppersistencetimeout", data["backuppersistencetimeout"])
	d.Set("persistmask", data["persistmask"])
	d.Set("cookiename", data["cookiename"])
	d.Set("v6persistmasklen", data["v6persistmasklen"])
	d.Set("cookiedomain", data["cookiedomain"])
	d.Set("timeout", data["timeout"])
	d.Set("rule", data["rule"])
	d.Set("usevserverpersistency", data["usevserverpersistency"])

	return nil
}

func updateLbgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateLbgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	LbgroupName := d.Get("name").(string)

	Lbgroup := lb.Lbgroup{
		Name: d.Get("name").(string),
	}

	hasChange := false

	if d.HasChange("persistencetype") {
		log.Printf("[DEBUG]  netscaler-provider: Persistencetype has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Persistencetype = d.Get("persistencetype").(string)
		hasChange = true
	}
	if d.HasChange("persistencebackup") {
		log.Printf("[DEBUG]  netscaler-provider: Persistencebackup has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Persistencebackup = d.Get("persistencebackup").(string)
		hasChange = true
	}
	if d.HasChange("backuppersistencetimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Backuppersistencetimeout has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Backuppersistencetimeout = d.Get("backuppersistencetimeout").(int)
		hasChange = true
	}
	if d.HasChange("persistmask") {
		log.Printf("[DEBUG]  netscaler-provider: Persistmask has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Persistmask = d.Get("persistmask").(string)
		hasChange = true
	}
	if d.HasChange("cookiename") {
		log.Printf("[DEBUG]  netscaler-provider: Cookiename has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Cookiename = d.Get("cookiename").(string)
		hasChange = true
	}
	if d.HasChange("v6persistmasklen") {
		log.Printf("[DEBUG]  netscaler-provider: V6persistmasklen has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.V6persistmasklen = d.Get("v6persistmasklen").(int)
		hasChange = true
	}
	if d.HasChange("cookiedomain") {
		log.Printf("[DEBUG]  netscaler-provider: Cookiedomain has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Cookiedomain = d.Get("cookiedomain").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  netscaler-provider: Timeout has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  netscaler-provider: Rule has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("usevserverpersistency") {
		log.Printf("[DEBUG]  netscaler-provider: Usevserverpersistency has changed for Lbgroup %s, starting update", LbgroupName)
		Lbgroup.Usevserverpersistency = d.Get("usevserverpersistency").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("lbgroup", LbgroupName, &Lbgroup)
		if err != nil {
			return fmt.Errorf("Error updating Lbgroup %s", LbgroupName)
		}
	}

	return readLbgroupFunc(d, meta)
}

func deleteLbgroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteLbgroupFunc")
	client := meta.(*NetScalerNitroClient).client
	LbgroupName := d.Id()
	err := client.DeleteResource("lbgroup", LbgroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
