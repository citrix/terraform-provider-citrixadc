package nslaslicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NslaslicenseApplyResource{}
var _ resource.ResourceWithConfigure = (*NslaslicenseApplyResource)(nil)

func NewNslaslicenseApplyResource() resource.Resource {
	return &NslaslicenseApplyResource{}
}

// NslaslicenseApplyResource defines the resource implementation.
type NslaslicenseApplyResource struct {
	client *service.NitroClient
}

// NslaslicenseApplyResourceModel describes the resource data model.
//
// This resource models the NITRO nslaslicense `?action=apply` action. apply is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. NOTE: applying a LAS license is DISRUPTIVE /
// non-idempotent on the appliance.
type NslaslicenseApplyResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Filelocation   types.String `tfsdk:"filelocation"`
	Filename       types.String `tfsdk:"filename"`
	Fixedbandwidth types.Bool   `tfsdk:"fixedbandwidth"`
}

func (r *NslaslicenseApplyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslaslicense_apply"
}

func (r *NslaslicenseApplyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslaslicenseApplyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslaslicense_apply resource.",
			},
			"filelocation": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "location of the file on Citrix ADC.",
			},
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the file. It should not include filepath.",
			},
			"fixedbandwidth": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "apply fixed bandwidth license on ADC",
			},
		},
	}
}

func (r *NslaslicenseApplyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslaslicenseApplyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslaslicense_apply resource")
	payload := nslaslicense_applyGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: NITRO exposes only the `apply` verb
	// (POST ?action=apply). NOTE: applying a LAS license is DISRUPTIVE /
	// non-idempotent on the appliance.
	err := r.client.ActOnResource(service.Nslaslicense.Type(), &payload, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslaslicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nslaslicense_apply resource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nslaslicense_apply")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseApplyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// apply is a one-shot action. NITRO has no GET endpoint that reports
	// apply-state, so Read is a pure preserve-state no-op.
	var data NslaslicenseApplyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nslaslicense_apply; NITRO has no query endpoint for apply state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseApplyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for apply; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state NslaslicenseApplyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nslaslicense_apply; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseApplyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// apply is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-apply"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nslaslicense_apply; NITRO has no inverse of the apply action")
}

func nslaslicense_applyGetThePayloadFromthePlan(ctx context.Context, data *NslaslicenseApplyResourceModel) ns.Nslaslicense {
	tflog.Debug(ctx, "In nslaslicense_applyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nslaslicense := ns.Nslaslicense{}
	if !data.Filelocation.IsNull() && !data.Filelocation.IsUnknown() {
		nslaslicense.Filelocation = data.Filelocation.ValueString()
	}
	if !data.Filename.IsNull() && !data.Filename.IsUnknown() {
		nslaslicense.Filename = data.Filename.ValueString()
	}
	if !data.Fixedbandwidth.IsNull() && !data.Fixedbandwidth.IsUnknown() {
		nslaslicense.Fixedbandwidth = data.Fixedbandwidth.ValueBool()
	}

	return nslaslicense
}
