package systemextramgmtcpu

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemextramgmtcpuResource{}
var _ resource.ResourceWithConfigure = (*SystemextramgmtcpuResource)(nil)
var _ resource.ResourceWithImportState = (*SystemextramgmtcpuResource)(nil)

func NewSystemextramgmtcpuResource() resource.Resource {
	return &SystemextramgmtcpuResource{}
}

// SystemextramgmtcpuResource defines the resource implementation.
type SystemextramgmtcpuResource struct {
	client *service.NitroClient
}

func (r *SystemextramgmtcpuResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemextramgmtcpuResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemextramgmtcpu"
}

func (r *SystemextramgmtcpuResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemextramgmtcpuResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemextramgmtcpu resource")

	systemextramgmtcpu := systemextramgmtcpuGetThePayloadFromtheConfig(ctx, &data)

	// systemextramgmtcpu is an action-only resource: enable/disable the extra
	// management CPU (mirrors SDK v2 createFunc which called ActOnResource with
	// "enable"/"disable" based on the `enabled` attribute).
	action := "disable"
	if data.Enabled.ValueBool() {
		action = "enable"
	}

	if err := r.client.ActOnResource(service.Systemextramgmtcpu.Type(), &systemextramgmtcpu, action); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to %s systemextramgmtcpu, got error: %s", action, err))
		return
	}

	// Optionally reboot the ADC and wait for it to become reachable again
	// (mirrors SDK v2 reboot handling using the rebooter wait logic).
	if data.Reboot.ValueBool() {
		if err := r.rebootAdcInstance(ctx); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reboot ADC after configuring systemextramgmtcpu, got error: %s", err))
			return
		}
		if err := r.waitReachable(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("ADC did not become reachable after reboot, got error: %s", err))
			return
		}
	}

	// Singleton resource -> static ID.
	data.Id = types.StringValue("systemextramgmtcpu-config")

	tflog.Trace(ctx, "Created systemextramgmtcpu resource")

	// NOTE: `enabled` is intentionally NOT read back from the ADC here. The
	// configured (plan) value is authoritative for the Create result; reading the
	// effectivestate back (which stays DISABLED on an unlicensed appliance) would
	// trigger an inconsistent-result-after-apply error. Drift is detected on the
	// next Read/refresh, matching SDK v2 ForceNew semantics.

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemextramgmtcpu resource")

	r.readSystemextramgmtcpuFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemextramgmtcpu resource")

	// systemextramgmtcpu exposes no NITRO 'set'/update: `enabled` and every
	// reachable_* knob are RequiresReplace, so the only attribute that can reach
	// Update is the `reboot` behavior flag. There is no API call to make here;
	// just persist the plan to state (mirrors SDK v2 Update: schema.Noop).

	tflog.Trace(ctx, "Updated systemextramgmtcpu resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemextramgmtcpuResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemextramgmtcpuResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemextramgmtcpu resource")

	// systemextramgmtcpu is a global/action-only configuration: there is nothing
	// to delete on the ADC (mirrors SDK v2 Delete: schema.Noop). Terraform removes
	// the resource from state automatically.
	tflog.Trace(ctx, "Deleted systemextramgmtcpu resource from state")
}

// Helper function to read systemextramgmtcpu data from API
func (r *SystemextramgmtcpuResource) readSystemextramgmtcpuFromApi(ctx context.Context, data *SystemextramgmtcpuResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemextramgmtcpu.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemextramgmtcpu, got error: %s", err))
		return
	}

	systemextramgmtcpuSetAttrFromGet(ctx, data, getResponseData)
}

// rebootAdcInstance issues a warm reboot of the ADC (mirrors SDK v2
// systemextramgmtcpuRebootAdcInstance).
func (r *SystemextramgmtcpuResource) rebootAdcInstance(ctx context.Context) error {
	tflog.Debug(ctx, "In systemextramgmtcpu rebootAdcInstance")
	reboot := ns.Reboot{
		Warm: true,
	}
	return r.client.ActOnResource("reboot", &reboot, "")
}

// waitReachable blocks until the ADC becomes reachable again after a reboot, or
// until reachable_timeout elapses (mirrors SDK v2 rebooterWaitReachable).
func (r *SystemextramgmtcpuResource) waitReachable(ctx context.Context, data *SystemextramgmtcpuResourceModel) error {
	tflog.Debug(ctx, "In systemextramgmtcpu waitReachable")

	timeout, err := time.ParseDuration(data.ReachableTimeout.ValueString())
	if err != nil {
		return err
	}
	pollInterval, err := time.ParseDuration(data.ReachablePollInterval.ValueString())
	if err != nil {
		return err
	}
	pollDelay, err := time.ParseDuration(data.ReachablePollDelay.ValueString())
	if err != nil {
		return err
	}
	pollTimeout, err := time.ParseDuration(data.ReachablePollTimeout.ValueString())
	if err != nil {
		return err
	}

	// Initial delay before we start polling (the box is going down).
	time.Sleep(pollDelay)

	deadline := time.Now().Add(timeout)
	for {
		if pollErr := r.pollLicense(pollTimeout); pollErr == nil {
			tflog.Debug(ctx, "ADC is reachable again after reboot")
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timed out after %s waiting for the ADC to become reachable following reboot", timeout)
		}
		time.Sleep(pollInterval)
	}
}

// pollLicense performs a single reachability probe against the nslicense endpoint
// (mirrors SDK v2 rebooterPollLicense). A nil return means the ADC responded.
func (r *SystemextramgmtcpuResource) pollLicense(timeout time.Duration) error {
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", r.client.GetURL())

	c := http.Client{
		Timeout: timeout,
	}
	buff := &bytes.Buffer{}
	req, err := http.NewRequest("GET", url, buff)
	if err != nil {
		return err
	}
	req.Header.Set("X-NITRO-USER", r.client.GetUsername())
	req.Header.Set("X-NITRO-PASS", r.client.GetPassword())

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
