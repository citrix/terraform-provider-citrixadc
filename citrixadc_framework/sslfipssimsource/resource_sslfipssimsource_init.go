package sslfipssimsource

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslfipssimsourceInitResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimsourceInitResource)(nil)

func NewSslfipssimsourceInitResource() resource.Resource {
	return &SslfipssimsourceInitResource{}
}

// SslfipssimsourceInitResource defines the resource implementation.
type SslfipssimsourceInitResource struct {
	client *service.NitroClient
}

// SslfipssimsourceInitResourceModel describes the resource data model.
//
// This resource models the NITRO sslfipssimsource `?action=init` action. init
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The init payload carries only the mandatory
// certfile attribute (sourcesecret/targetsecret are enable-only).
type SslfipssimsourceInitResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Certfile types.String `tfsdk:"certfile"`
}

func (r *SslfipssimsourceInitResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimsource_init"
}

func (r *SslfipssimsourceInitResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimsourceInitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimsource exposes only the
		// `enable` and `init` NITRO actions (no get/add/delete). This resource
		// models the `init` action, whose payload is {certfile}.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimsource_init resource.",
			},
			"certfile": schema.StringAttribute{
				// Required for the init action (Pattern 8: tfdata wrongly marks optional).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the source FIPS appliance's certificate file. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func (r *SslfipssimsourceInitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimsourceInitResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Initializing sslfipssimsource (action-only resource)")
	payload := sslfipssimsource_initGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes init as POST ?action=init. Use ActOnResource with the
	// case-sensitive "init" verb (lower-case per the NITRO URL).
	// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
	err := r.client.ActOnResource(service.Sslfipssimsource.Type(), &payload, "init")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to init sslfipssimsource, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Initialized sslfipssimsource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimsource_init")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceInitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// init is a one-shot action. NITRO has no GET endpoint that reports
	// init-state, so Read is a pure preserve-state no-op.
	var data SslfipssimsourceInitResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslfipssimsource_init; NITRO has no query endpoint for init state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceInitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for init; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SslfipssimsourceInitResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslfipssimsource_init; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceInitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// init is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfipssimsource_init; NITRO has no inverse of the init action")
}

func sslfipssimsource_initGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimsourceInitResourceModel) ssl.Sslfipssimsource {
	tflog.Debug(ctx, "In sslfipssimsource_initGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only the init action's field
	// (certfile) is set; sourcesecret/targetsecret are enable-only and omitted.
	sslfipssimsource := ssl.Sslfipssimsource{}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslfipssimsource.Certfile = data.Certfile.ValueString()
	}

	return sslfipssimsource
}
