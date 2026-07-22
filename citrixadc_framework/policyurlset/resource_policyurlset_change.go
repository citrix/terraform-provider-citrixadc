package policyurlset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicyurlsetChangeResource{}
var _ resource.ResourceWithConfigure = (*PolicyurlsetChangeResource)(nil)

func NewPolicyurlsetChangeResource() resource.Resource {
	return &PolicyurlsetChangeResource{}
}

// PolicyurlsetChangeResource defines the resource implementation.
type PolicyurlsetChangeResource struct {
	client *service.NitroClient
}

// PolicyurlsetChangeResourceModel describes the resource data model.
//
// This resource models the NITRO policyurlset change action. NOTE: the NITRO
// doc labels the section `change`, but the real HTTP action (and CLI verb) is
// `update`; Create therefore calls ActOnResource with the verb "update". This
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The payload carries only `name`.
type PolicyurlsetChangeResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *PolicyurlsetChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyurlset_change"
}

func (r *PolicyurlsetChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyurlsetChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyurlset_change resource.",
			},
			// NITRO update payload marks `name` as mandatory.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique name of the url set. Maximum length: 127.",
			},
		},
	}
}

func (r *PolicyurlsetChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyurlsetChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyurlset_change resource")
	payload := policyurlset_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes this action as POST ?action=update (the doc's `change`
	// section documents action=update; there is no `change` CLI verb).
	err := r.client.ActOnResource(service.Policyurlset.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policyurlset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated policyurlset resource")

	data.Id = types.StringValue(fmt.Sprintf("policyurlset_change-%v", data.Name.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// update is a one-shot action. NITRO has no GET endpoint that reports
	// this action's state, so Read is a pure preserve-state no-op.
	var data PolicyurlsetChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for policyurlset_change; NITRO has no query endpoint for this action")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update-in-place endpoint for this action; every schema
	// attribute is RequiresReplace, so Terraform never invokes Update for a
	// real change.
	var data, state PolicyurlsetChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policyurlset_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// This is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for policyurlset_change; NITRO has no inverse of the update action")
}

func policyurlset_changeGetThePayloadFromthePlan(ctx context.Context, data *PolicyurlsetChangeResourceModel) policy.Policyurlset {
	tflog.Debug(ctx, "In policyurlset_changeGetThePayloadFromthePlan Function")

	// NITRO `?action=update` accepts only `name`.
	policyurlset := policy.Policyurlset{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policyurlset.Name = data.Name.ValueString()
	}

	return policyurlset
}
