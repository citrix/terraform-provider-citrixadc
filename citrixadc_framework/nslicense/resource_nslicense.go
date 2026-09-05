package nslicense

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Default per-operation timeout, matching the SDK v2 resource's 20m timeouts.
const nslicenseOperationTimeout = 20 * time.Minute

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NslicenseResource{}
var _ resource.ResourceWithConfigure = (*NslicenseResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseResource)(nil)

func NewNslicenseResource() resource.Resource {
	return &NslicenseResource{}
}

// NslicenseResource defines the resource implementation.
type NslicenseResource struct {
	client *service.NitroClient
}

func (r *NslicenseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicense"
}

func (r *NslicenseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicense resource")

	// ID scheme matches SDK v2: the license file name.
	data.Id = types.StringValue(data.LicenseFile.ValueString())

	// Resolve the optional+computed ssh_port so state never carries an unknown value.
	if data.SshPort.IsUnknown() {
		data.SshPort = types.Int64Null()
	}

	sshConn, err := r.getSshConnection(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating ssh client. %s", err.Error()))
		return
	}
	defer sshConn.Close()

	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating sftp client. %s", err.Error()))
		return
	}
	defer sftpClient.Close()

	if err := r.uploadLicenseFile(ctx, &data, sftpClient); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error uploading license file. %s", err.Error()))
		return
	}

	if data.Reboot.ValueBool() {
		if err := r.powerCycleAndWait(ctx, &data, nslicenseOperationTimeout); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error power cycling ADC. %s", err.Error()))
			return
		}
	}

	tflog.Trace(ctx, "Created nslicense resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicense resource")

	// Guard against an unknown ssh_port ever leaking into state.
	if data.SshPort.IsUnknown() {
		data.SshPort = types.Int64Null()
	}

	sshConn, err := r.getSshConnection(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating ssh client. %s", err.Error()))
		return
	}
	defer sshConn.Close()

	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating sftp client. %s", err.Error()))
		return
	}
	defer sftpClient.Close()

	fileName := data.LicenseFile.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("nslicense: checking remote license file %s", fileName))
	if err := r.ensureLicenseFileExists(ctx, sftpClient, fileName); err != nil {
		// The license file is gone on the appliance. Mirror the SDK v2 intent
		// (which cleared the ForceNew license_file, forcing recreation) by
		// removing the resource from state so Terraform plans a recreate.
		tflog.Debug(ctx, fmt.Sprintf("nslicense: license file %s not found on ADC, removing from state", fileName))
		resp.State.RemoveResource(ctx)
		return
	}

	// File present: preserve prior state unchanged (all attributes are config-driven).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslicenseResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nslicense resource")

	// Update is a no-op push (matching SDK v2): the only config-changing
	// attribute that touches the appliance is license_file, which is
	// RequiresReplace. All other attributes (ssh_*, reboot, poll_*) are plain
	// state values that Terraform adopts in place. Preserve the ID from state.
	data.Id = state.Id

	if data.SshPort.IsUnknown() {
		data.SshPort = types.Int64Null()
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicense resource")

	sshConn, err := r.getSshConnection(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating ssh client. %s", err.Error()))
		return
	}
	defer sshConn.Close()

	sftpClient, err := sftp.NewClient(sshConn)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error creating sftp client. %s", err.Error()))
		return
	}
	defer sftpClient.Close()

	r.deleteLicenseFile(ctx, sftpClient, data.LicenseFile.ValueString())

	if data.Reboot.ValueBool() {
		if err := r.powerCycleAndWait(ctx, &data, nslicenseOperationTimeout); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error power cycling ADC. %s", err.Error()))
			return
		}
	}

	tflog.Trace(ctx, "Deleted nslicense resource")
}

// ---------------------------------------------------------------------------
// SSH / SFTP / reboot helpers (ported from the SDK v2 resource, adapted to the
// Plugin Framework model + NitroClient accessors).
// ---------------------------------------------------------------------------

func (r *NslicenseResource) getSshConnection(ctx context.Context, data *NslicenseResourceModel) (*ssh.Client, error) {
	tflog.Debug(ctx, "netscaler-provider: In getSshConnection")

	var username, password, host, port string

	// Configure ssh username (fall back to the provider NITRO username).
	if !data.SshUsername.IsNull() && data.SshUsername.ValueString() != "" {
		username = data.SshUsername.ValueString()
	} else {
		username = r.client.GetUsername()
	}

	// Configure ssh password (fall back to the provider NITRO password).
	if !data.SshPassword.IsNull() && data.SshPassword.ValueString() != "" {
		password = data.SshPassword.ValueString()
	} else {
		password = r.client.GetPassword()
	}

	// Configure ssh host (fall back to the host parsed from the NITRO endpoint).
	if !data.SshHost.IsNull() && data.SshHost.ValueString() != "" {
		host = data.SshHost.ValueString()
	} else {
		u, err := url.Parse(r.client.GetURL())
		if err != nil {
			return nil, err
		}
		host = strings.Split(u.Host, ":")[0]
	}

	if !data.SshPort.IsNull() && !data.SshPort.IsUnknown() && data.SshPort.ValueInt64() != 0 {
		port = strconv.FormatInt(data.SshPort.ValueInt64(), 10)
	} else {
		port = "22"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	// Configure host key verification.
	publickey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(data.SshHostPubkey.ValueString()))
	if err != nil {
		return nil, err
	}
	hostKeyCallBack := ssh.FixedHostKey(publickey)

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

func (r *NslicenseResource) uploadLicenseFile(ctx context.Context, data *NslicenseResourceModel, sftpClient *sftp.Client) error {
	tflog.Debug(ctx, "netscaler-provider: In uploadLicenseFile")

	fileName := data.LicenseFile.ValueString()

	// Stat file to verify it exists locally.
	if _, err := os.Stat(fileName); err != nil {
		return err
	}

	localFile, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		return err
	}
	defer func() {
		if err := localFile.Close(); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("netscaler-provider: error closing license file %v", err))
		}
	}()

	fileBytes, err := ioutil.ReadAll(localFile)
	if err != nil {
		return err
	}

	remotePath := fmt.Sprintf("/nsconfig/license/%s", fileName)
	remoteFile, err := sftpClient.Create(remotePath)
	if err != nil {
		return err
	}
	defer remoteFile.Close()

	if _, err = remoteFile.Write(fileBytes); err != nil {
		return err
	}
	return nil
}

func (r *NslicenseResource) deleteLicenseFile(ctx context.Context, sftpClient *sftp.Client, fileName string) {
	tflog.Debug(ctx, "netscaler-provider: In deleteLicenseFile")
	remotePath := fmt.Sprintf("/nsconfig/license/%s", fileName)
	sftpClient.Remove(remotePath)
}

func (r *NslicenseResource) ensureLicenseFileExists(ctx context.Context, sftpClient *sftp.Client, fileName string) error {
	tflog.Debug(ctx, "netscaler-provider: In ensureLicenseFileExists")
	remotePath := fmt.Sprintf("/nsconfig/license/%s", fileName)
	_, err := sftpClient.Stat(remotePath)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("netscaler-provider: error for remote license file: %s", err.Error()))
	}
	return err
}

func (r *NslicenseResource) rebootAdcInstance(ctx context.Context) error {
	tflog.Debug(ctx, "netscaler-provider: In rebootAdcInstance")
	reboot := ns.Reboot{
		Warm: true,
	}
	return r.client.ActOnResource("reboot", &reboot, "")
}

// pollLicense probes the appliance's NITRO endpoint to test reachability after
// a reboot. It returns nil when the appliance is reachable, an error otherwise.
func (r *NslicenseResource) pollLicense(ctx context.Context, data *NslicenseResourceModel) error {
	tflog.Debug(ctx, "netscaler-provider: In pollLicense")

	username := r.client.GetUsername()
	password := r.client.GetPassword()
	endpoint := r.client.GetURL()
	urlStr := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	timeout, err := time.ParseDuration(data.PollTimeout.ValueString())
	if err != nil {
		return err
	}
	c := http.Client{
		Timeout: timeout,
	}
	buff := &bytes.Buffer{}
	req, _ := http.NewRequest("GET", urlStr, buff)
	req.Header.Set("X-NITRO-USER", username)
	req.Header.Set("X-NITRO-PASS", password)
	resp, err := c.Do(req)
	if err != nil {
		if !strings.Contains(err.Error(), "Client.Timeout exceeded") {
			// Unexpected error
			return err
		}
		// Expected timeout error
		return fmt.Errorf("Timeout")
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	// Reachable
	return nil
}

// powerCycleAndWait reboots the appliance and blocks until it is reachable
// again (or the timeout elapses). It mirrors the SDK v2 StateChangeConf poll:
// an initial poll_delay, then a poll every poll_interval until reachable.
func (r *NslicenseResource) powerCycleAndWait(ctx context.Context, data *NslicenseResourceModel, timeout time.Duration) error {
	tflog.Debug(ctx, "netscaler-provider: In powerCycleAndWait")

	if err := r.rebootAdcInstance(ctx); err != nil {
		return fmt.Errorf("Error rebooting ADC. %s", err.Error())
	}

	pollInterval, err := time.ParseDuration(data.PollInterval.ValueString())
	if err != nil {
		return err
	}
	pollDelay, err := time.ParseDuration(data.PollDelay.ValueString())
	if err != nil {
		return err
	}

	// Initial delay before the first reachability poll (the appliance needs
	// time to actually go down before we start polling for it to come back).
	time.Sleep(pollDelay)

	deadline := time.Now().Add(timeout)
	for {
		if err := r.pollLicense(ctx, data); err == nil {
			tflog.Debug(ctx, "netscaler-provider: ADC reachable")
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for ADC to become reachable")
		}
		time.Sleep(pollInterval)
	}
}
