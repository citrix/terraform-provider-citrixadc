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
var _ resource.Resource = &NsconfigDiffResource{}
var _ resource.ResourceWithConfigure = (*NsconfigDiffResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigDiffResource)(nil)

func NewNsconfigDiffResource() resource.Resource {
	return &NsconfigDiffResource{}
}

// NsconfigDiffResource defines the resource implementation.
type NsconfigDiffResource struct {
	client *service.NitroClient
}

// NsconfigDiffResourceModel describes the resource data model.
// Action-only resource (NITRO nsconfig `?action=diff`). `timestamp` is a
// synthetic TF-only field used as the resource ID; re-running the action requires
// bumping it (RequiresReplace).
type NsconfigDiffResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Config1              types.String `tfsdk:"config1"`
	Config2              types.String `tfsdk:"config2"`
	Outtype              types.String `tfsdk:"outtype"`
	Template             types.Bool   `tfsdk:"template"`
	Ignoredevicespecific types.Bool   `tfsdk:"ignoredevicespecific"`
	Timestamp            types.String `tfsdk:"timestamp"`
}

func (r *NsconfigDiffResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_diff"
}

func (r *NsconfigDiffResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconfig_diff resource (equals the configured timestamp).",
			},
			"config1": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Location of the configurations.",
			},
			"config2": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Location of the configurations.",
			},
			"outtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Format to display the difference in configurations. Possible values: [ cli, xml ]",
			},
			"template": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "File that contains the commands to be compared.",
			},
			"ignoredevicespecific": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Suppress device specific differences.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp marker used as the resource ID. Change it to re-run diff ns config.",
			},
		},
	}
}

func (r *NsconfigDiffResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigDiffResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsconfigDiffResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigDiffResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_diff resource (diff ns config)")
	payload := nsconfig_diffGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes diff as POST ?action=diff. Verb casing is lower-case per the
	// NITRO URL.
	if err := r.client.ActOnResource(service.Nsconfig.Type(), &payload, "diff"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to diff ns config, got error: %s", err))
		return
	}

	// Synthetic ID equals the configured timestamp.
	data.Id = data.Timestamp

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigDiffResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No GET endpoint for the diff action (Pattern 13): preserve state as-is.
	var data NsconfigDiffResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read is a no-op for nsconfig_diff (action-only, no GET endpoint)")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigDiffResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; Update is never expected to run.
	var data, state NsconfigDiffResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsconfig_diff; all attributes are RequiresReplace")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigDiffResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Action-only resource: there is no inverse API. Delete just removes from state.
	tflog.Debug(ctx, "Deleting nsconfig_diff: action-only, removing from state only")
}

func nsconfig_diffGetThePayloadFromthePlan(ctx context.Context, data *NsconfigDiffResourceModel) ns.Nsconfig {
	tflog.Debug(ctx, "In nsconfig_diffGetThePayloadFromthePlan Function")

	// NITRO nsconfig `?action=diff` accepts: config1, config2, outtype, template,
	// ignoredevicespecific.
	nsconfig := ns.Nsconfig{}
	if !data.Config1.IsNull() && !data.Config1.IsUnknown() {
		nsconfig.Config1 = data.Config1.ValueString()
	}
	if !data.Config2.IsNull() && !data.Config2.IsUnknown() {
		nsconfig.Config2 = data.Config2.ValueString()
	}
	if !data.Outtype.IsNull() && !data.Outtype.IsUnknown() {
		nsconfig.Outtype = data.Outtype.ValueString()
	}
	if !data.Template.IsNull() && !data.Template.IsUnknown() {
		nsconfig.Template = data.Template.ValueBool()
	}
	if !data.Ignoredevicespecific.IsNull() && !data.Ignoredevicespecific.IsUnknown() {
		nsconfig.Ignoredevicespecific = data.Ignoredevicespecific.ValueBool()
	}

	return nsconfig
}
