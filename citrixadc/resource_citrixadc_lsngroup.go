package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lsn"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcLsngroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLsngroupFunc,
		Read:          readLsngroupFunc,
		Update:        updateLsngroupFunc,
		Delete:        deleteLsngroupFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"groupname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"clientname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allocpolicy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ftp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftpcm": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6profile": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nattype": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"portblocksize": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pptp": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspalg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionlogging": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionsync": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sipalg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmptraplimit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createLsngroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLsngroupFunc")
	client := meta.(*NetScalerNitroClient).client
	lsngroupName := d.Get("groupname").(string)

	lsngroup := make(map[string]interface{})
	if v, ok := d.GetOk("allocpolicy"); ok {
		lsngroup["allocpolicy"] = v.(string)
	}
	if v, ok := d.GetOk("clientname"); ok {
		lsngroup["clientname"] = v.(string)
	}
	if v, ok := d.GetOk("ftp"); ok {
		lsngroup["ftp"] = v.(string)
	}
	if v, ok := d.GetOk("ftpcm"); ok {
		lsngroup["ftpcm"] = v.(string)
	}
	if v, ok := d.GetOk("groupname"); ok {
		lsngroup["groupname"] = v.(string)
	}
	if v, ok := d.GetOk("ip6profile"); ok {
		lsngroup["ip6profile"] = v.(string)
	}
	if v, ok := d.GetOk("logging"); ok {
		lsngroup["logging"] = v.(string)
	}
	if v, ok := d.GetOk("nattype"); ok {
		lsngroup["nattype"] = v.(string)
	}
	if v, ok := d.GetOkExists("portblocksize"); ok {
		lsngroup["portblocksize"] = v.(int)
	}
	if v, ok := d.GetOk("pptp"); ok {
		lsngroup["pptp"] = v.(string)
	}
	if v, ok := d.GetOk("rtspalg"); ok {
		lsngroup["rtspalg"] = v.(string)
	}
	if v, ok := d.GetOk("sessionlogging"); ok {
		lsngroup["sessionlogging"] = v.(string)
	}
	if v, ok := d.GetOk("sessionsync"); ok {
		lsngroup["sessionsync"] = v.(string)
	}
	if v, ok := d.GetOk("sipalg"); ok {
		lsngroup["sipalg"] = v.(string)
	}
	if v, ok := d.GetOkExists("snmptraplimit"); ok {
		val, _ := strconv.Atoi(v.(string))
		lsngroup["snmptraplimit"] = val
	}

	_, err := client.AddResource("lsngroup", lsngroupName, &lsngroup)
	if err != nil {
		return err
	}

	d.SetId(lsngroupName)

	err = readLsngroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lsngroup but we can't read it ?? %s", lsngroupName)
		return nil
	}
	return nil
}

func readLsngroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLsngroupFunc")
	client := meta.(*NetScalerNitroClient).client
	lsngroupName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading lsngroup state %s", lsngroupName)
	data, err := client.FindResource("lsngroup", lsngroupName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing lsngroup state %s", lsngroupName)
		d.SetId("")
		return nil
	}
	d.Set("groupname", data["groupname"])
	d.Set("allocpolicy", data["allocpolicy"])
	d.Set("clientname", data["clientname"])
	d.Set("ftp", data["ftp"])
	d.Set("ftpcm", data["ftpcm"])
	d.Set("groupname", data["groupname"])
	d.Set("ip6profile", data["ip6profile"])
	d.Set("logging", data["logging"])
	d.Set("nattype", data["nattype"])
	d.Set("portblocksize", data["portblocksize"])
	d.Set("pptp", data["pptp"])
	d.Set("rtspalg", data["rtspalg"])
	d.Set("sessionlogging", data["sessionlogging"])
	d.Set("sessionsync", data["sessionsync"])
	d.Set("sipalg", data["sipalg"])
	d.Set("snmptraplimit", data["snmptraplimit"])

	return nil

}

func updateLsngroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLsngroupFunc")
	client := meta.(*NetScalerNitroClient).client
	lsngroupName := d.Get("groupname").(string)

	lsngroup := lsn.Lsngroup{
		Groupname: d.Get("groupname").(string),
	}
	hasChange := false
	if d.HasChange("ftp") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftp has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Ftp = d.Get("ftp").(string)
		hasChange = true
	}
	if d.HasChange("ftpcm") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftpcm has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Ftpcm = d.Get("ftpcm").(string)
		hasChange = true
	}

	if d.HasChange("logging") {
		log.Printf("[DEBUG]  citrixadc-provider: Logging has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Logging = d.Get("logging").(string)
		hasChange = true
	}
	if d.HasChange("portblocksize") {
		log.Printf("[DEBUG]  citrixadc-provider: Portblocksize has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Portblocksize = d.Get("portblocksize").(int)
		hasChange = true
	}
	if d.HasChange("pptp") {
		log.Printf("[DEBUG]  citrixadc-provider: Pptp has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Pptp = d.Get("pptp").(string)
		hasChange = true
	}
	if d.HasChange("rtspalg") {
		log.Printf("[DEBUG]  citrixadc-provider: Rtspalg has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Rtspalg = d.Get("rtspalg").(string)
		hasChange = true
	}
	if d.HasChange("sessionlogging") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionlogging has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Sessionlogging = d.Get("sessionlogging").(string)
		hasChange = true
	}
	if d.HasChange("sessionsync") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionsync has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Sessionsync = d.Get("sessionsync").(string)
		hasChange = true
	}
	if d.HasChange("sipalg") {
		log.Printf("[DEBUG]  citrixadc-provider: Sipalg has changed for lsngroup %s, starting update", lsngroupName)
		lsngroup.Sipalg = d.Get("sipalg").(string)
		hasChange = true
	}
	if d.HasChange("snmptraplimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Snmptraplimit has changed for lsngroup %s, starting update", lsngroupName)
		val, _ := strconv.Atoi(d.Get("snmptraplimit").(string))
		lsngroup.Snmptraplimit = val
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("lsngroup", &lsngroup)
		if err != nil {
			return fmt.Errorf("Error updating lsngroup %s", lsngroupName)
		}
	}
	return readLsngroupFunc(d, meta)
}

func deleteLsngroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLsngroupFunc")
	client := meta.(*NetScalerNitroClient).client
	lsngroupName := d.Id()
	err := client.DeleteResource("lsngroup", lsngroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
