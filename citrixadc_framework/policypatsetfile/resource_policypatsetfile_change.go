package policypatsetfile

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
var _ resource.Resource = &PolicypatsetfileChangeResource{}
var _ resource.ResourceWithConfigure = (*PolicypatsetfileChangeResource)(nil)

func NewPolicypatsetfileChangeResource() resource.Resource {
	return &PolicypatsetfileChangeResource{}
}

// PolicypatsetfileChangeResource defines the resource implementation.
type PolicypatsetfileChangeResource struct {
	client *service.NitroClient
}

// PolicypatsetfileChangeResourceModel describes the resource data model.
//
// This resource models the NITRO policypatsetfile `change` operation. NITRO
// exposes it as POST `?action=update` (the CLI verb is `update`; there is no
// `change`/`?action=change` endpoint). It is a one-shot side-effect action with
// no GET endpoint and no inverse API, so Read/Update/Delete are no-ops. The
// update payload carries only `name`.
type PolicypatsetfileChangeResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *PolicypatsetfileChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policypatsetfile_change"
}

func (r *PolicypatsetfileChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicypatsetfileChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policypatsetfile_change resource.",
			},
			// NITRO update payload marks `name` as mandatory.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.",
			},
		},
	}
}

func (r *PolicypatsetfileChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicypatsetfileChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policypatsetfile_change resource")
	payload := policypatsetfile_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes the change operation as POST ?action=update. Use ActOnResource
	// with the case-sensitive "update" verb (the CLI equivalent is `update`; there
	// is no ?action=change endpoint).
	err := r.client.ActOnResource(service.Policypatsetfile.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policypatsetfile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated policypatsetfile resource")

	// ID = the patset file name that was updated; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("policypatsetfile_change-%v", data.Name.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// change is a one-shot action. NITRO has no GET endpoint that reports
	// update-state, so Read is a pure preserve-state no-op.
	var data PolicypatsetfileChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for policypatsetfile_change; NITRO has no query endpoint for change state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for this action wrapper; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state PolicypatsetfileChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policypatsetfile_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicypatsetfileChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// change is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for policypatsetfile_change; NITRO has no inverse of the change action")
}

func policypatsetfile_changeGetThePayloadFromthePlan(ctx context.Context, data *PolicypatsetfileChangeResourceModel) policy.Policypatsetfile {
	tflog.Debug(ctx, "In policypatsetfile_changeGetThePayloadFromthePlan Function")

	// NITRO `?action=update` accepts only `name`.
	policypatsetfile := policy.Policypatsetfile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policypatsetfile.Name = data.Name.ValueString()
	}

	return policypatsetfile
}
