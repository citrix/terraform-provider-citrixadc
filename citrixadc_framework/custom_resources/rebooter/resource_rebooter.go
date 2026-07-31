package rebooter

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RebooterResource{}
var _ resource.ResourceWithConfigure = (*RebooterResource)(nil)

func NewRebooterResource() resource.Resource {
	return &RebooterResource{}
}

// RebooterResource models the NITRO `reboot` action.
//
// It is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops and every input attribute is RequiresReplace. Create
// issues the reboot (warm or cold) and, when wait_until_reachable is true, polls the
// appliance's nslicense endpoint until the ADC comes back up. A synthetic ID keeps the
// action-only resource addressable by Terraform. This mirrors the legacy SDKv2
// citrixadc_rebooter resource exactly.
type RebooterResource struct {
	client *service.NitroClient
}

// RebooterResourceModel describes the resource data model. Every schema attribute has
// a matching tfsdk field.
type RebooterResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Warm                  types.Bool   `tfsdk:"warm"`
	Timestamp             types.String `tfsdk:"timestamp"`
	WaitUntilReachable    types.Bool   `tfsdk:"wait_until_reachable"`
	ReachableTimeout      types.String `tfsdk:"reachable_timeout"`
	ReachablePollDelay    types.String `tfsdk:"reachable_poll_delay"`
	ReachablePollInterval types.String `tfsdk:"reachable_poll_interval"`
	ReachablePollTimeout  types.String `tfsdk:"reachable_poll_timeout"`
}

func (r *RebooterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rebooter"
}

func (r *RebooterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RebooterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rebooter resource.",
			},
			"warm": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Restarts the Citrix ADC software without rebooting the underlying operating system.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp value used to trigger the reboot. Change it to force a new reboot.",
			},
			"wait_until_reachable": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "When true, wait until the ADC is reachable again after the reboot.",
			},
			"reachable_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("10m"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Maximum duration to wait for the ADC to become reachable.",
			},
			"reachable_poll_delay": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("60s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Delay before the first reachability poll.",
			},
			"reachable_poll_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("60s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interval between reachability polls.",
			},
			"reachable_poll_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("20s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Per-poll HTTP timeout when checking reachability.",
			},
		},
	}
}

func (r *RebooterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RebooterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rebooter resource")

	// Issue the reboot action.
	if err := rebooterRebootAdcInstance(ctx, r.client, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reboot ADC instance, got error: %s", err))
		return
	}

	// Optionally block until the appliance is reachable again.
	if data.WaitUntilReachable.ValueBool() {
		if err := rebooterWaitReachable(ctx, r.client, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("ADC did not become reachable after reboot, got error: %s", err))
			return
		}
	}

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("tf-rebooter-%d", time.Now().UnixNano()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RebooterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The reboot is a one-shot action. NITRO has no GET endpoint that reports its
	// state, so Read is a pure preserve-state no-op.
	var data RebooterResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for rebooter; NITRO has no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RebooterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the reboot action; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state RebooterResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for rebooter; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RebooterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// The reboot is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete only removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for rebooter; NITRO has no inverse of the reboot action")
}

// rebooterRebootAdcInstance issues the NITRO `reboot` action (warm or cold).
//
// The appliance commonly tears down the TCP connection the instant it accepts the
// reboot command — before the HTTP response is written — which surfaces as a
// transport error ("EOF" or "read: connection reset by peer") rather than a NITRO
// errorcode. In that case the reboot WAS accepted (GitHub issue #980: the box
// reboots but the Terraform apply errors out and breaks dependent resources). Treat
// such connection-teardown errors as success and let Create fall through to the
// wait_until_reachable poll, which is the authoritative signal that the ADC actually
// went down and came back up. Genuine errors (auth failures, NITRO errorcodes,
// connection-refused / timeouts from a box that was never reachable, etc.) are still
// returned so real problems are not masked.
func rebooterRebootAdcInstance(ctx context.Context, client *service.NitroClient, data *RebooterResourceModel) error {
	reboot := ns.Reboot{
		Warm: data.Warm.ValueBool(),
	}
	err := client.ActOnResource("reboot", &reboot, "")
	if err != nil && rebooterIsConnectionTeardownError(err) {
		tflog.Warn(ctx, fmt.Sprintf("citrixadc-provider: reboot request returned a connection-teardown error (%s); the appliance closed the connection as it rebooted — treating the reboot as accepted (GitHub #980)", err))
		return nil
	}
	return err
}

// rebooterIsConnectionTeardownError reports whether err is the transport-level
// connection teardown the appliance produces when it reboots mid-request, as
// opposed to a genuine NITRO/auth error or a box that was never reachable. A reset
// or EOF means the TCP connection was established and then torn down by the peer,
// which is the signature of an accepted reboot; "connection refused" / timeouts are
// deliberately NOT matched here because they mean the appliance was never reached.
// See GitHub issue #980.
func rebooterIsConnectionTeardownError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "EOF") ||
		strings.Contains(msg, "connection reset by peer") ||
		strings.Contains(msg, "connection reset")
}

// rebooterWaitReachable blocks until the ADC becomes reachable again or the configured
// timeout elapses. It mirrors the legacy SDKv2 StateChangeConf-based polling: an initial
// delay, then polls of the nslicense endpoint at the configured interval.
func rebooterWaitReachable(ctx context.Context, client *service.NitroClient, data *RebooterResourceModel) error {
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

	// Initial delay before the first poll.
	time.Sleep(pollDelay)

	deadline := time.Now().Add(timeout)
	for {
		err := rebooterPollLicense(client, data)
		if err == nil {
			tflog.Debug(ctx, "citrixadc-provider: ADC is reachable")
			return nil
		}
		if err.Error() != "Timeout" {
			// Unexpected error; surface it.
			return err
		}
		// Still unreachable; check the overall timeout before polling again.
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for ADC to become reachable after %s", timeout)
		}
		time.Sleep(pollInterval)
	}
}

// rebooterPollLicense performs a single GET against the ADC nslicense endpoint using a
// per-poll HTTP timeout. It returns a "Timeout" error when the appliance is still
// unreachable, nil once it responds, or the underlying error otherwise.
func rebooterPollLicense(client *service.NitroClient, data *RebooterResourceModel) error {
	username := client.GetUsername()
	password := client.GetPassword()
	endpoint := client.GetURL()
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	timeout, err := time.ParseDuration(data.ReachablePollTimeout.ValueString())
	if err != nil {
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
		}
		// Expected timeout error
		return fmt.Errorf("Timeout")
	}
	defer resp.Body.Close()
	// No error
	return nil
}
