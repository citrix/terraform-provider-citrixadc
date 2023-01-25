package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func resourceCitrixAdcInstaller() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createInstallerFunc,
		Read:          schema.Noop,
		Delete:        schema.Noop,
		Schema: map[string]*schema.Schema{
			"enhancedupgrade": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"l": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"resizeswapvar": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"y": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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

func createInstallerFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createInstallFunc")
	installerId := resource.PrefixedUniqueId("tf-installer-")

	err := installerInstallBuild(d, meta)
	if err != nil {
		return err
	}

	if d.Get("wait_until_reachable").(bool) {
		err := installerWaitReachable(d, meta)
		if err != nil {
			return err
		}
	}

	d.SetId(installerId)

	return nil
}

func installerInstallBuild(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In installerInstallBuild")

	client := meta.(*NetScalerNitroClient).client

	install := utility.Install{
		Enhancedupgrade: d.Get("enhancedupgrade").(bool),
		L:               d.Get("l").(bool),
		Resizeswapvar:   d.Get("resizeswapvar").(bool),
		Url:             d.Get("url").(string),
		Y:               d.Get("y").(bool),
	}

	if err := client.ActOnResource("install", &install, ""); err != nil {
		errorStr := err.Error()
		if strings.HasSuffix(errorStr, "EOF") || strings.HasSuffix(errorStr, "connection reset by peer") {
			// This is expected since the operation results in a TCP conection reset some times
			// especially when y = true
			log.Printf("[DEBUG] citrixadc-provider: Ignoring go-nitro error \"%s\"", errorStr)
			return nil
		} else {
			return err
		}
	}
	return nil
}

func installerWaitReachable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In installerWaitReachable")

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
		Refresh:      installerInstancePoll(d, meta),
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

func installerPollLicense(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider: In installerPollLicense")

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

func installerInstancePoll(d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] citrixadc-provider: In installerInstancePoll")
		err := installerPollLicense(d, meta)
		if err != nil {
			log.Printf("[DEBUG] citrixadc-provider: Unreachable: %v", err.Error())
			return nil, "unreachable", nil
		}
		log.Printf("[DEBUG] citrixadc-provider: Returning \"reachable\"")
		return "reachable", "reachable", nil
	}
}
