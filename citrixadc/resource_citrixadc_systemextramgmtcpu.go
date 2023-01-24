package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

func resourceCitrixAdcSystemextramgmtcpu() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemextramgmtcpuFunc,
		Read:          readSystemextramgmtcpuFunc,
		Delete:        schema.Noop,
		Update:        schema.Noop,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"reboot": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"reachable_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
				ForceNew: true,
			},
			"reachable_poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
				ForceNew: true,
			},
			"reachable_poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
				ForceNew: true,
			},
			"reachable_poll_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "20s",
				ForceNew: true,
			},
		},
	}
}

func createSystemextramgmtcpuFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemextramgmtcpuFunc")
	client := meta.(*NetScalerNitroClient).client
	systemextramgmtcpuName := resource.PrefixedUniqueId("tf-systemextramgmtcpu-")

	systemextramgmtcpu := system.Systemextramgmtcpu{}

	var action string
	if d.Get("enabled").(bool) {
		action = "enable"
	} else {
		action = "disable"
	}

	err := client.ActOnResource("systemextramgmtcpu", &systemextramgmtcpu, action)
	if err != nil {
		return err
	}

	if d.Get("reboot").(bool) {
		var err error
		err = systemextramgmtcpuRebootAdcInstance(d, meta)
		if err != nil {
			return err
		}

		// Reusing wait function from rebooter resource
		err = rebooterWaitReachable(d, meta)
		if err != nil {
			return err
		}
	}

	d.SetId(systemextramgmtcpuName)

	err = readSystemextramgmtcpuFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemextramgmtcpu but we can't read it ?? %s", systemextramgmtcpuName)
		return nil
	}
	return nil
}

func readSystemextramgmtcpuFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemextramgmtcpuFunc")
	client := meta.(*NetScalerNitroClient).client
	systemextramgmtcpuName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemextramgmtcpu state %s", systemextramgmtcpuName)
	findParams := service.FindParams{
		ResourceType: "systemextramgmtcpu",
	}
	dataarray, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		log.Printf("[ERROR] citrixadc-provider: Error reading state: %s", err.Error())
		log.Printf("[WARN] citrixadc-provider: Clearing systemextramgmtcpu state %s", systemextramgmtcpuName)
		d.SetId("")
		return nil
	}
	data := dataarray[0]
	if data["effectivestate"].(string) == "ENABLED" {
		d.Set("enabled", true)
	} else {
		d.Set("enabled", false)
	}

	return nil

}

func systemextramgmtcpuRebootAdcInstance(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In systemextramgmtcpuRebootAdcInstance")

	client := meta.(*NetScalerNitroClient).client
	reboot := ns.Reboot{
		Warm: true,
	}
	if err := client.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}
	return nil
}
