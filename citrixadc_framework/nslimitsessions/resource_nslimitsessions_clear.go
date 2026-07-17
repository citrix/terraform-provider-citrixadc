package nslimitsessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NslimitsessionsClearResource{}
var _ resource.ResourceWithConfigure = (*NslimitsessionsClearResource)(nil)
var _ resource.ResourceWithImportState = (*NslimitsessionsClearResource)(nil)

func NewNslimitsessionsClearResource() resource.Resource {
	return &NslimitsessionsClearResource{}
}

// NslimitsessionsClearResource defines the resource implementation.
type NslimitsessionsClearResource struct {
	client *service.NitroClient
}

// NslimitsessionsClearResourceModel describes the resource data model.
//
// This resource models the NITRO nslimitsessions `?action=clear` action. clear
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The clear payload carries only the mandatory
// limitidentifier (detail is a GET-only filter and lives on the datasource).
type NslimitsessionsClearResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Limitidentifier types.String `tfsdk:"limitidentifier"`
}

func (r *NslimitsessionsClearResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslimitsessionsClearResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitsessions_clear"
}

func (r *NslimitsessionsClearResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslimitsessionsClearResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslimitsessions_clear resource.",
			},
			"limitidentifier": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the rate limit identifier for which to display the sessions.",
			},
		},
	}
}

func (r *NslimitsessionsClearResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslimitsessionsClearResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing nslimitsessions (action-only resource, action=clear)")
	payload := nslimitsessions_clearGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes clear as POST ?action=clear. Use ActOnResource with the
	// case-sensitive "clear" verb (lower-case per the NITRO URL). The clear
	// payload carries only limitidentifier.
	err := r.client.ActOnResource(service.Nslimitsessions.Type(), &payload, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear nslimitsessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared nslimitsessions sessions")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nslimitsessions_clear")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsClearResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// clear is a one-shot action. NITRO has no GET endpoint that reports
	// clear-state, so Read is a pure preserve-state no-op.
	var data NslimitsessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nslimitsessions_clear; NITRO has no query endpoint for clear state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsClearResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for clear; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state NslimitsessionsClearResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nslimitsessions_clear; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsClearResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// clear is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-clear"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nslimitsessions_clear; NITRO has no inverse of the clear action")
}

func nslimitsessions_clearGetThePayloadFromthePlan(ctx context.Context, data *NslimitsessionsClearResourceModel) ns.Nslimitsessions {
	tflog.Debug(ctx, "In nslimitsessions_clearGetThePayloadFromthePlan Function")

	// Create API request body from the model. The clear action accepts only
	// limitidentifier (detail is a GET-only filter, excluded here).
	nslimitsessions := ns.Nslimitsessions{}
	if !data.Limitidentifier.IsNull() && !data.Limitidentifier.IsUnknown() {
		nslimitsessions.Limitidentifier = data.Limitidentifier.ValueString()
	}

	return nslimitsessions
}
