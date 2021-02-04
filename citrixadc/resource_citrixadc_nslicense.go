package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ns"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func resourceCitrixAdcNslicense() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNslicenseFunc,
		Read:          readNslicenseFunc,
		Update:        updateNslicenseFunc,
		Delete:        deleteNslicenseFunc,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"license_file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_username": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"ssh_host_pubkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reboot": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			"poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			"poll_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
		},
	}
}

func rebootAdcInstance(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In rebootAdcInstance")

	client := meta.(*NetScalerNitroClient).client
	reboot := ns.Reboot{
		Warm: true,
	}
	if err := client.ActOnResource("reboot", &reboot, ""); err != nil {
		return err
	}
	return nil
}

func pollLicense(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In pollLicense")

	username := meta.(*NetScalerNitroClient).Username
	password := meta.(*NetScalerNitroClient).Password
	endpoint := meta.(*NetScalerNitroClient).Endpoint
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	var timeout time.Duration
	var err error
	if timeout, err = time.ParseDuration(d.Get("poll_timeout").(string)); err != nil {
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

func resourceAdcinstanceLicensePoll(d *schema.ResourceData, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] netscaler-provider: In resourceAdcinstanceLicensePoll")
		err := pollLicense(d, meta)
		if err != nil {
			if err.Error() == "Timeout" {
				return nil, "unreachable", nil
			} else {
				return nil, "unreachable", err
			}
		}
		log.Printf("[DEBUG] netscaler-provider: Returning \"reachable\"")
		return "reachable", "reachable", nil
	}
}

func powerCycleAndWait(d *schema.ResourceData, meta interface{}, t time.Duration) error {
	log.Printf("[DEBUG] netscaler-provider: In powerCycleAndWait")
	var err error

	if err = rebootAdcInstance(d, meta); err != nil {
		return fmt.Errorf("Error rebooting ADC. %s", err.Error())
	}

	var poll_interval time.Duration
	if poll_interval, err = time.ParseDuration(d.Get("poll_interval").(string)); err != nil {
		return err
	}

	var poll_delay time.Duration
	if poll_delay, err = time.ParseDuration(d.Get("poll_delay").(string)); err != nil {
		return err
	}
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"unreachable"},
		Target:       []string{"reachable"},
		Refresh:      resourceAdcinstanceLicensePoll(d, meta),
		Timeout:      t,
		PollInterval: poll_interval,
		Delay:        poll_delay,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}
	return nil
}

func getSshConnection(d *schema.ResourceData, meta interface{}) (*ssh.Client, error) {
	log.Printf("[DEBUG] netscaler-provider: In getSshConnection")

	nsClient := meta.(*NetScalerNitroClient)

	var username, password, host, port string

	// Configure ssh username
	if val, ok := d.GetOk("ssh_username"); ok {
		username = val.(string)
	} else {
		username = nsClient.Username
	}

	// Configure ssh password
	if val, ok := d.GetOk("ssh_password"); ok {
		password = val.(string)
	} else {
		password = nsClient.Password
	}

	// Configure ssh host
	if val, ok := d.GetOk("ssh_host"); ok {
		host = val.(string)
	} else {
		// Parse the NITRO API endpoint for ssh host
		u, err := url.Parse(nsClient.Endpoint)
		if err != nil {
			return nil, err
		}
		host = strings.Split(u.Host, ":")[0]
	}

	if val, ok := d.GetOk("ssh_port"); ok {
		port = strconv.Itoa(val.(int))
	} else {
		port = "22"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	// Confgiure host key verification
	var hostKeyCallBack ssh.HostKeyCallback

	val := d.Get("ssh_host_pubkey")
	publickey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(val.(string)))
	if err != nil {
		return nil, err
	}

	hostKeyCallBack = ssh.FixedHostKey(publickey)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: hostKeyCallBack,
	}

	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getSftpClient(d *schema.ResourceData, meta interface{}, sshConn *ssh.Client) (*sftp.Client, error) {
	log.Printf("[DEBUG] netscaler-provider: In getSftpClient")
	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		return nil, err
	}
	return sftpClient, nil
}

func uploadLicenseFile(d *schema.ResourceData, meta interface{}, sftpClient *sftp.Client) error {
	log.Printf("[DEBUG] netscaler-provider: In uploadLicenseFile")

	fileName := d.Get("license_file").(string)

	// Stat file to verify it exists
	_, err := os.Stat(fileName)

	if err != nil {
		return err
	}

	localFile, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return err
	}

	defer func() {
		err := localFile.Close()
		if err != nil {
			log.Printf("[DEBUG] netscaler-provider: error closing license file %v", err)
		}
	}()

	fileBytes, err := ioutil.ReadAll(localFile)
	if err != nil {
		return err
	}

	remotePath := fmt.Sprintf("/nsconfig/license/%s", d.Get("license_file").(string))
	remoteFile, err := sftpClient.Create(remotePath)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	_, err = remoteFile.Write(fileBytes)
	if err != nil {
		return err
	}
	return nil
}

func deleteLicenseFile(sftpClient *sftp.Client, fileName string) error {
	log.Printf("[DEBUG] netscaler-provider: In deleteLicenseFile")
	remotePath := fmt.Sprintf("/nsconfig/license/%s", fileName)
	sftpClient.Remove(remotePath)
	return nil
}

func ensureLicenseFileExists(sftpClient *sftp.Client, fileName string) error {
	log.Printf("[DEBUG] netscaler-provider: In ensureLicenseFileExists")
	remotePath := fmt.Sprintf("/nsconfig/license/%s", fileName)
	_, err := sftpClient.Stat(remotePath)
	if err != nil {
		log.Printf("[DEBUG] netscaler-provider: error for remote license file: %s", err.Error())
	}
	return err
}

func createNslicenseFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In createNslicenseFunc")
	//client := meta.(*NetScalerNitroClient).client
	licenseFile := d.Get("license_file").(string)

	d.SetId(licenseFile)

	sshConn, err := getSshConnection(d, meta)
	if err != nil {
		return fmt.Errorf("Error creating ssh client. %s", err.Error())
	}
	defer sshConn.Close()

	sftpClient, err := getSftpClient(d, meta, sshConn)
	if err != nil {
		return fmt.Errorf("Error creating sftp client. %s", err.Error())
	}
	defer sftpClient.Close()

	err = uploadLicenseFile(d, meta, sftpClient)
	if err != nil {
		return err
	}

	if d.Get("reboot").(bool) {
		err = powerCycleAndWait(d, meta, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("Error power cycling ADC. %s", err.Error())
		}
	}
	return readNslicenseFunc(d, meta)
	// return nil
}

func readNslicenseFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In readNslicenseFunc")
	//client := meta.(*NetScalerNitroClient).client

	sshConn, err := getSshConnection(d, meta)
	if err != nil {
		return fmt.Errorf("Error creating ssh client. %s", err.Error())
	}
	defer sshConn.Close()

	sftpClient, err := getSftpClient(d, meta, sshConn)
	if err != nil {
		return fmt.Errorf("Error creating sftp client. %s", err.Error())
	}
	defer sftpClient.Close()

	fileName := d.Get("license_file").(string)
	log.Printf("[DEBUG] netscaler-provider: %s", fileName)
	err = ensureLicenseFileExists(sftpClient, fileName)
	if err != nil {
		d.Set("license_file", "")
	}

	return nil
}

func updateNslicenseFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In updateNslicenseFunc")
	// Update is a noop
	// The only relevant option change is for license_file
	// which has ForceNew: true

	return nil
}

func deleteNslicenseFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider: In deleteNslicenseFunc")
	//client := meta.(*NetScalerNitroClient).client

	sshConn, err := getSshConnection(d, meta)
	if err != nil {
		return fmt.Errorf("Error creating ssh client. %s", err.Error())
	}
	defer sshConn.Close()

	sftpClient, err := getSftpClient(d, meta, sshConn)
	if err != nil {
		return fmt.Errorf("Error creating sftp client. %s", err.Error())
	}
	defer sftpClient.Close()

	licenseFilename := d.Get("license_file").(string)
	deleteLicenseFile(sftpClient, licenseFilename)

	if d.Get("reboot").(bool) {
		err = powerCycleAndWait(d, meta, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return fmt.Errorf("Error power cycling ADC. %s", err.Error())
		}
	}

	d.SetId("")

	return nil
}
