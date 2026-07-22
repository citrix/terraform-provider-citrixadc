package nsextension

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsextensionChangeResource{}
var _ resource.ResourceWithConfigure = (*NsextensionChangeResource)(nil)

func NewNsextensionChangeResource() resource.Resource {
	return &NsextensionChangeResource{}
}

// NsextensionChangeResource defines the resource implementation.
type NsextensionChangeResource struct {
	client *service.NitroClient
}

// NsextensionChangeResourceModel describes the resource data model.
//
// This resource models the NITRO nsextension `change` action, which reloads /
// recompiles the extension object from its stored source file. change is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The change payload carries only `name`.
//
// NOTE: although the NITRO doc anchor is named "change", the action is invoked
// at ?action=update (POST). The literal "change" verb is rejected by NITRO with
// errorcode 1240. The verb string passed to ActOnResource is therefore "update".
type NsextensionChangeResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func (r *NsextensionChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsextension_change"
}

func (r *NsextensionChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsextensionChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsextension_change resource.",
			},
			// NITRO change (?action=update) payload marks `name` as mandatory.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the extension object to reload from its stored source file.",
			},
		},
	}
}

func (r *NsextensionChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsextensionChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsextension_change resource")
	payload := nsextension_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes the change (reload) op at POST ?action=update. The literal
	// "change" verb is rejected with errorcode 1240, so the verb string is
	// "update".
	err := r.client.ActOnResource(service.Nsextension.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change nsextension, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Changed nsextension resource")

	// ID = the extension name that was reloaded; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("nsextension_change-%v", data.Name.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// change is a one-shot action. NITRO has no GET endpoint that reports
	// change-state, so Read is a pure preserve-state no-op.
	var data NsextensionChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsextension_change; NITRO has no query endpoint for change state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the change action; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state NsextensionChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsextension_change; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// change is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-change"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nsextension_change; NITRO has no inverse of the change action")
}

func nsextension_changeGetThePayloadFromthePlan(ctx context.Context, data *NsextensionChangeResourceModel) ns.Nsextension {
	tflog.Debug(ctx, "In nsextension_changeGetThePayloadFromthePlan Function")

	// NITRO change (?action=update) accepts only `name`.
	nsextension := ns.Nsextension{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsextension.Name = data.Name.ValueString()
	}

	return nsextension
}
