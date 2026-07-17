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
var _ resource.Resource = &NsconfigConvertResource{}
var _ resource.ResourceWithConfigure = (*NsconfigConvertResource)(nil)
var _ resource.ResourceWithImportState = (*NsconfigConvertResource)(nil)

func NewNsconfigConvertResource() resource.Resource {
	return &NsconfigConvertResource{}
}

// NsconfigConvertResource defines the resource implementation.
type NsconfigConvertResource struct {
	client *service.NitroClient
}

// NsconfigConvertResourceModel describes the resource data model.
// Action-only resource (NITRO nsconfig `?action=convert`). `timestamp` is a
// synthetic TF-only field used as the resource ID; re-running the action requires
// bumping it (RequiresReplace).
type NsconfigConvertResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Configfile   types.String `tfsdk:"configfile"`
	Responsefile types.String `tfsdk:"responsefile"`
	Async        types.Bool   `tfsdk:"async"`
	Outtype      types.String `tfsdk:"outtype"`
	Timestamp    types.String `tfsdk:"timestamp"`
}

func (r *NsconfigConvertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig_convert"
}

func (r *NsconfigConvertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconfig_convert resource (equals the configured timestamp).",
			},
			// CLI + NITRO doc both mark configfile mandatory (tfdata mismatch; Pattern 8).
			"configfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Full path of config file to be converted to nitro.",
			},
			"responsefile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Full path of file to store the nitro graph. If not specified, nitro graph is returned as part of the API response.",
			},
			"async": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Using this option will run the operation in async mode and return the job id. The job ID can be used later to track the conversion progress via show ns job <id> Command. This option is mostly useful for API to avoid timeouts for large input configuration.",
			},
			"outtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Format to display the difference in configurations.",
			},
			"timestamp": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timestamp marker used as the resource ID. Change it to re-run convert ns config.",
			},
		},
	}
}

func (r *NsconfigConvertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsconfigConvertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsconfigConvertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsconfigConvertResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsconfig_convert resource (convert ns config)")
	payload := nsconfig_convertGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes convert as POST ?action=convert. Verb casing is lower-case per
	// the NITRO URL.
	if err := r.client.ActOnResource(service.Nsconfig.Type(), &payload, "convert"); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert ns config, got error: %s", err))
		return
	}

	// Synthetic ID equals the configured timestamp.
	data.Id = data.Timestamp

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigConvertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No GET endpoint for the convert action (Pattern 13): preserve state as-is.
	var data NsconfigConvertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Read is a no-op for nsconfig_convert (action-only, no GET endpoint)")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigConvertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All attributes are RequiresReplace; Update is never expected to run.
	var data, state NsconfigConvertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsconfig_convert; all attributes are RequiresReplace")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsconfigConvertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Action-only resource: there is no inverse API. Delete just removes from state.
	tflog.Debug(ctx, "Deleting nsconfig_convert: action-only, removing from state only")
}

func nsconfig_convertGetThePayloadFromthePlan(ctx context.Context, data *NsconfigConvertResourceModel) ns.Nsconfig {
	tflog.Debug(ctx, "In nsconfig_convertGetThePayloadFromthePlan Function")

	// NITRO nsconfig `?action=convert` accepts: configfile, responsefile, async
	// (Go json tag `Async`, capital A), outtype.
	nsconfig := ns.Nsconfig{}
	if !data.Configfile.IsNull() && !data.Configfile.IsUnknown() {
		nsconfig.Configfile = data.Configfile.ValueString()
	}
	if !data.Responsefile.IsNull() && !data.Responsefile.IsUnknown() {
		nsconfig.Responsefile = data.Responsefile.ValueString()
	}
	if !data.Async.IsNull() && !data.Async.IsUnknown() {
		nsconfig.Async = data.Async.ValueBool()
	}
	if !data.Outtype.IsNull() && !data.Outtype.IsUnknown() {
		nsconfig.Outtype = data.Outtype.ValueString()
	}

	return nsconfig
}
