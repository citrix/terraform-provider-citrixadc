package nsacls6

import (
	"context"
	"fmt"

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
var _ resource.Resource = &Nsacls6RenumberResource{}
var _ resource.ResourceWithConfigure = (*Nsacls6RenumberResource)(nil)
var _ resource.ResourceWithImportState = (*Nsacls6RenumberResource)(nil)

func NewNsacls6RenumberResource() resource.Resource {
	return &Nsacls6RenumberResource{}
}

// Nsacls6RenumberResource defines the resource implementation.
type Nsacls6RenumberResource struct {
	client *service.NitroClient
}

// Nsacls6RenumberResourceModel describes the resource data model.
//
// This resource models the NITRO nsacls6 `?action=renumber` action. renumber is
// a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The renumber payload carries the optional
// attribute type.
type Nsacls6RenumberResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (r *Nsacls6RenumberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nsacls6RenumberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsacls6_renumber"
}

func (r *Nsacls6RenumberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nsacls6RenumberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacls6_renumber resource.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of the acl ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
		},
	}
}

func (r *Nsacls6RenumberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nsacls6RenumberResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Renumbering nsacls6 (action-only resource)")
	payload := nsacls6_renumberGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes renumber as POST ?action=renumber. Use ActOnResource with
	// the case-sensitive "renumber" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Nsacls6.Type(), &payload, "renumber")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to renumber nsacls6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Renumbered nsacls6")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nsacls6_renumber")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6RenumberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// renumber is a one-shot action. NITRO has no GET endpoint that reports
	// renumber-state, so Read is a pure preserve-state no-op.
	var data Nsacls6RenumberResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsacls6_renumber; NITRO has no query endpoint for renumber state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6RenumberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for renumber; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state Nsacls6RenumberResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsacls6_renumber; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacls6RenumberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// renumber is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-renumber"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nsacls6_renumber; NITRO has no inverse of the renumber action")
}

// nsacls6_renumberGetThePayloadFromthePlan builds the renumber action payload,
// including only this action's fields.
func nsacls6_renumberGetThePayloadFromthePlan(ctx context.Context, data *Nsacls6RenumberResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nsacls6_renumberGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		payload["type"] = data.Type.ValueString()
	} else {
		// type defaults to CLASSIC per NITRO doc
		payload["type"] = "CLASSIC"
	}

	return payload
}
