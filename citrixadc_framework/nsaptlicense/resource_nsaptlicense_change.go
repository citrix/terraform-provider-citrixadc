package nsaptlicense

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
var _ resource.Resource = &NsaptlicenseChangeResource{}
var _ resource.ResourceWithConfigure = (*NsaptlicenseChangeResource)(nil)

func NewNsaptlicenseChangeResource() resource.Resource {
	return &NsaptlicenseChangeResource{}
}

// NsaptlicenseChangeResource defines the resource implementation.
type NsaptlicenseChangeResource struct {
	client *service.NitroClient
}

// NsaptlicenseChangeResourceModel describes the resource data model.
//
// This resource models the NITRO nsaptlicense `change` action, which allocates
// APT license counts. Although the NITRO doc anchor is labelled "change", the
// action is invoked at POST ?action=update; the verb string passed to
// ActOnResource is therefore "update" (there is no `change ns aptlicense` CLI
// command). change is a one-shot side-effect action with no inverse API, so
// Read/Update/Delete are no-ops.
//
// NOTE: allocating APT license counts via this resource is DISRUPTIVE and
// non-idempotent on the appliance.
//
// The Terraform identifier is the NITRO License ID ("id"), which is the
// CLI-mandatory key for the change/update action (it is a real NITRO field, not
// a synthetic marker). "serialno" is a GET-only filter key (Pattern 15) and is
// therefore NOT sent in the change payload.
type NsaptlicenseChangeResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Bindtype       types.String `tfsdk:"bindtype"`
	Countavailable types.String `tfsdk:"countavailable"`
	Licensedir     types.String `tfsdk:"licensedir"`
	Serialno       types.String `tfsdk:"serialno"`
	Sessionid      types.String `tfsdk:"sessionid"`
	Useproxy       types.String `tfsdk:"useproxy"`
}

func (r *NsaptlicenseChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsaptlicense_change"
}

func (r *NsaptlicenseChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsaptlicenseChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				// NITRO License ID: CLI-mandatory key for the change
				// (?action=update) action, and the Terraform resource identifier.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "License ID",
			},
			"bindtype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind type",
			},
			"countavailable": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses",
			},
			"licensedir": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "License Directory",
			},
			"serialno": schema.StringAttribute{
				// GET-only filter key (Pattern 15) - not part of the change
				// action payload.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Hardware Serial Number/License Activation Code(LAC)",
			},
			"sessionid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Session ID",
			},
			"useproxy": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.",
			},
		},
	}
}

func (r *NsaptlicenseChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsaptlicenseChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsaptlicense_change resource")
	payload := nsaptlicense_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes the change op at POST ?action=update (the doc anchor is
	// "change" but the action verb is "update" - there is no `change ns
	// aptlicense` CLI command). Preserve the "update" verb exactly.
	// NOTE: allocating APT license counts is DISRUPTIVE / non-idempotent.
	err := r.client.ActOnResource(service.Nsaptlicense.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change nsaptlicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Changed nsaptlicense resource")

	// The Terraform identifier is the NITRO License ID ("id"), which is the
	// CLI-mandatory key for the change/update action and is already populated
	// from the plan (set once here, Pattern 6). It is deliberately NOT
	// overwritten with a synthetic marker: doing so would strand the mandatory
	// License ID the update payload requires and conflict with the Required id
	// config attribute.

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// change is a one-shot action. Read is a pure preserve-state no-op so the
	// disruptive license allocation is never re-run implicitly and no attribute
	// drifts against the live appliance.
	var data NsaptlicenseChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsaptlicense_change")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the change action; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state NsaptlicenseChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nsaptlicense_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaptlicenseChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// change is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-allocate"). Delete simply removes the resource from Terraform
	// state; the allocated licenses remain on the appliance.
	tflog.Debug(ctx, "Delete is a no-op for nsaptlicense_change; NITRO has no inverse of the change action")
}

func nsaptlicense_changeGetThePayloadFromthePlan(ctx context.Context, data *NsaptlicenseChangeResourceModel) ns.Nsaptlicense {
	tflog.Debug(ctx, "In nsaptlicense_changeGetThePayloadFromthePlan Function")

	// NITRO change (?action=update) accepts id, sessionid, bindtype,
	// countavailable, licensedir and useproxy. serialno is a GET-only filter key
	// (Pattern 15) and is intentionally excluded from the change payload.
	nsaptlicense := ns.Nsaptlicense{}
	if !data.Bindtype.IsNull() && !data.Bindtype.IsUnknown() {
		nsaptlicense.Bindtype = data.Bindtype.ValueString()
	}
	if !data.Countavailable.IsNull() && !data.Countavailable.IsUnknown() {
		nsaptlicense.Countavailable = data.Countavailable.ValueString()
	}
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		nsaptlicense.Id = data.Id.ValueString()
	}
	if !data.Licensedir.IsNull() && !data.Licensedir.IsUnknown() {
		nsaptlicense.Licensedir = data.Licensedir.ValueString()
	}
	if !data.Sessionid.IsNull() && !data.Sessionid.IsUnknown() {
		nsaptlicense.Sessionid = data.Sessionid.ValueString()
	}
	if !data.Useproxy.IsNull() && !data.Useproxy.IsUnknown() {
		nsaptlicense.Useproxy = data.Useproxy.ValueString()
	}

	return nsaptlicense
}
