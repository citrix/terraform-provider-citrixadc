package nsconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsconfigClearResource{}
var _ resource.ResourceWithConfigure = (*NsconfigClearResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigClearResource)(nil)

func NewNsconfigClearResource() resource.Resource {
	return &NsconfigClearResource{}
}

// NsconfigClearResource defines the resource implementation.
type NsconfigClearResource struct {
	client *service.NitroClient
}

// NsconfigClearResourceModel describes the resource data model.
// Action-only resource (NITRO `clear ns config`). `timestamp` is a synthetic
// TF-only field used as the resource ID.
type NsconfigClearResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Level     types.String `tfsdk:"level"`
	Force     types.Bool   `tfsdk:"force"`
	Rbaconfig types.String `tfsdk:"rbaconfig"`
	Timestamp types.String `tfsdk:"timestamp"`
}

func (r *NsconfigClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_clear"
}

func (r *NsconfigClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconfig_clear resource (equals the configured timestamp).",
			},
			"level": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Types of configurations to be cleared: basic, extended, or full.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Configurations will be cleared without prompting for confirmation.",
			},
			"rbaconfig": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "RBA configurations and TACACS policies bound to system global will not be cleared if RBA is set to NO. Applicable only for the basic clear level. Default is YES.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp marker used as the resource ID. Change it to re-run clear ns config.",
			},
		},
	}
}

func (r *NsconfigClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsconfigClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigClearResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_clear resource (clear ns config)")
	nsconfig := ns.Nsconfig{}
	if !data.Level.IsNull() && !data.Level.IsUnknown() {
		nsconfig.Level = data.Level.ValueString()
	}
	if !data.Force.IsNull() && !data.Force.IsUnknown() {
		nsconfig.Force = data.Force.ValueBool()
	}
	if !data.Rbaconfig.IsNull() && !data.Rbaconfig.IsUnknown() {
		nsconfig.Rbaconfig = data.Rbaconfig.ValueString()
	}

	if err := r.client.ActOnResource(service.Nsconfig.Type(), &nsconfig, "clear"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear ns config, got error: %s", err))
		return
	}

	// Synthetic ID equals the configured timestamp (matches SDK v2 SetId(timestamp)).
	data.Id = data.Timestamp

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No GET endpoint for the clear action (Pattern 13): preserve state as-is.
	var data NsconfigClearResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read is a no-op for nsconfig_clear (action-only, no GET endpoint)")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; Update is never expected to run.
	var data, state NsconfigClearResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsconfig_clear; all attributes are RequiresReplace")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Action-only resource: there is no inverse API. Delete just removes from state
	// (matches SDK v2 schema.Noop delete).
	tflog.Debug(ctx, "Deleting nsconfig_clear: action-only, removing from state only")
}
