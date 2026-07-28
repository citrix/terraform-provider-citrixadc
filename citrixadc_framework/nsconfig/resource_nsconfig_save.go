package nsconfig

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsconfigSaveResource{}
var _ resource.ResourceWithConfigure = (*NsconfigSaveResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigSaveResource)(nil)

func NewNsconfigSaveResource() resource.Resource {
	return &NsconfigSaveResource{}
}

// NsconfigSaveResource defines the resource implementation.
type NsconfigSaveResource struct {
	client *service.NitroClient
}

// NsconfigSaveResourceModel describes the resource data model.
// This is an action-only resource (NITRO `save ns config`). The `timestamp`
// attribute is a synthetic TF-only field used as the resource ID; re-running the
// action requires bumping it (RequiresReplace).
type NsconfigSaveResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	All                    types.Bool   `tfsdk:"all"`
	Timestamp              types.String `tfsdk:"timestamp"`
	ConcurrentSaveOk       types.Bool   `tfsdk:"concurrent_save_ok"`
	ConcurrentSaveRetries  types.Int64  `tfsdk:"concurrent_save_retries"`
	ConcurrentSaveTimeout  types.String `tfsdk:"concurrent_save_timeout"`
	ConcurrentSaveInterval types.String `tfsdk:"concurrent_save_interval"`
	SaveOnDestroy          types.Bool   `tfsdk:"save_on_destroy"`
}

func (r *NsconfigSaveResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_save"
}

func (r *NsconfigSaveResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconfig_save resource (equals the configured timestamp).",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to do saveconfig for all partitions.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp marker used as the resource ID. Change it to re-run save ns config.",
			},
			"concurrent_save_ok": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Tolerate a concurrent save (NITRO errorcode 293) by retrying the save operation.",
			},
			"concurrent_save_retries": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(0),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of times to retry the save when a concurrent save is in progress.",
			},
			"concurrent_save_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("5m"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Total timeout (Go duration) for the concurrent-save retry loop.",
			},
			"concurrent_save_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("10s"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interval (Go duration) between concurrent-save retries.",
			},
			"save_on_destroy": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "If true, perform save ns config again when the resource is destroyed.",
			},
		},
	}
}

func (r *NsconfigSaveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigSaveResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// doSaveConfig performs the NITRO `save ns config` action.
func (r *NsconfigSaveResource) doSaveConfig(ctx context.Context, data *NsconfigSaveResourceModel) error {
	tflog.Debug(ctx, "In NsconfigSaveResource doSaveConfig")
	nsconfig := ns.Nsconfig{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		nsconfig.All = data.All.ValueBool()
	}
	return r.client.ActOnResource(service.Nsconfig.Type(), &nsconfig, "save")
}

// saveWithConcurrencyHandling mirrors the SDK v2 retry logic: a NITRO errorcode
// 293 indicates another save is in progress; retry up to concurrent_save_retries
// times honoring the configured interval and timeout.
func (r *NsconfigSaveResource) saveWithConcurrencyHandling(ctx context.Context, data *NsconfigSaveResourceModel) error {
	err := r.doSaveConfig(ctx, data)
	if err == nil {
		return nil
	}
	if !strings.Contains(err.Error(), "\"errorcode\": 293") {
		return err
	}
	// Concurrent save in progress.
	if !data.ConcurrentSaveOk.ValueBool() {
		return err
	}
	retries := int(data.ConcurrentSaveRetries.ValueInt64())
	if retries <= 0 {
		// Tolerated but no retries requested: treat 293 as success (matches SDK v2
		// fallthrough where SetId is still called).
		return nil
	}

	interval, perr := time.ParseDuration(data.ConcurrentSaveInterval.ValueString())
	if perr != nil {
		return perr
	}
	timeout, perr := time.ParseDuration(data.ConcurrentSaveTimeout.ValueString())
	if perr != nil {
		return perr
	}

	deadline := time.Now().Add(timeout)
	var lastErr error = err
	for i := 0; i < retries; i++ {
		time.Sleep(interval)
		if time.Now().After(deadline) {
			break
		}
		lastErr = r.doSaveConfig(ctx, data)
		if lastErr == nil {
			return nil
		}
		if !strings.Contains(lastErr.Error(), "\"errorcode\": 293") {
			return lastErr
		}
	}
	return lastErr
}

func (r *NsconfigSaveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigSaveResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_save resource (save ns config)")
	if err := r.saveWithConcurrencyHandling(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to save ns config, got error: %s", err))
		return
	}

	// Synthetic ID equals the configured timestamp (matches SDK v2 SetId(timestamp)).
	data.Id = data.Timestamp

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigSaveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No GET endpoint for the save action (Pattern 13): preserve state as-is.
	var data NsconfigSaveResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read is a no-op for nsconfig_save (action-only, no GET endpoint)")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigSaveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; Update is never expected to run.
	var data, state NsconfigSaveResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsconfig_save; all attributes are RequiresReplace")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigSaveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsconfigSaveResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !data.SaveOnDestroy.ValueBool() {
		tflog.Debug(ctx, "Deleting nsconfig_save: save_on_destroy is false, removing from state only")
		return
	}

	tflog.Debug(ctx, "Deleting nsconfig_save: save_on_destroy is true, performing save ns config")
	if err := r.saveWithConcurrencyHandling(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to save ns config on destroy, got error: %s", err))
		return
	}
}
