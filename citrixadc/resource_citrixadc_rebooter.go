package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func resourceCitrixAdcRebooter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRebooterFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"warm": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"wait_until_reachable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
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

func createRebooterFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRebooterFunc")
	rebooterId := resource.PrefixedUniqueId("tf-rebooter-")

	err := rebooterRebootAdcInstance(d, meta)
	if err != nil {
		return err
	}

	if d.Get("wait_until_reachable").(bool) {
		err := rebooterWaitReachable(d, meta)
		if err != nil {
			return err
		}
	}

	d.SetId(rebooterId)

	return nil
}

func rebooterRebootAdcInstance(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In rebooterRebootAdcInstance")

	client := meta.(*NetScalerNitroClient).client
	reboot := ns.Reboot{
		Warm: d.Get("warm").(bool),
	}
	if err := client.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}
	return nil
}

func rebooterWaitReachable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In rebooterWaitReachable")

	var err error
	var timeout time.Duration
	if timeout, err = time.ParseDuration(d.Get("reachable_timeout").(string)); err != nil {
		return err
	}

	var poll_interval time.Duration
	if poll_interval, err = time.ParseDuration(d.Get("reachable_poll_interval").(string)); err != nil {
		return err
	}

	var poll_delay time.Duration
	if poll_delay, err = time.ParseDuration(d.Get("reachable_poll_delay").(string)); err != nil {
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"unreachable"},
		Target:       []string{"reachable"},
		Refresh:      rebooterInstancePoll(d, meta),
		Timeout:      timeout,
		PollInterval: poll_interval,
		Delay:        poll_delay,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	return nil
}

func rebooterPollLicense(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In rebooterPollLicense")

	username := meta.(*NetScalerNitroClient).Username
	password := meta.(*NetScalerNitroClient).Password
	endpoint := meta.(*NetScalerNitroClient).Endpoint
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	var timeout time.Duration
	var err error
	if timeout, err = time.ParseDuration(d.Get("reachable_poll_timeout").(string)); err != nil {
		return err
	}
	c := http.Client{
		Timeout: timeout,
	}
	buff := &bytes.Buffer{}
	req, _ := http.NewRequest("GET", url, buff)
	req.Header.Set("X-NITRO-USER", username)
	req.Header.Set("X-NITRO-PASS", password)
	resp, err := c.Do(req)
	if err != nil {
		if !strings.Contains(err.Error(), "Client.Timeout exceeded") {
			// Unexpected error
			return err
		} else {
			// Expected timeout error
			return fmt.Errorf("Timeout")
		}
	} else {
		log.Printf("Status code is %v\n", resp.Status)
	}
	// No error
	return nil
}

func rebooterInstancePoll(d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] citrixadc-provider: In reboooterInstancePoll")
		err := rebooterPollLicense(d, meta)
		if err != nil {
			if err.Error() == "Timeout" {
				return nil, "unreachable", nil
			} else {
				return nil, "unreachable", err
			}
		}
		log.Printf("[DEBUG] citrixadc-provider: Returning \"reachable\"")
		return "reachable", "reachable", nil
	}
}
