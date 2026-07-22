package appfwlearningdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwlearningdataResetResource{}
var _ resource.ResourceWithConfigure = (*AppfwlearningdataResetResource)(nil)

func NewAppfwlearningdataResetResource() resource.Resource {
	return &AppfwlearningdataResetResource{}
}

// AppfwlearningdataResetResource defines the resource implementation.
type AppfwlearningdataResetResource struct {
	client *service.NitroClient
}

// AppfwlearningdataResetResourceModel describes the resource data model.
//
// This resource models the NITRO appfwlearningdata `?action=reset` action, which
// clears ALL learned-data databases and zeroes the transaction count. reset is a
// one-shot side-effect action: NITRO accepts an EMPTY payload
// ({"appfwlearningdata":{}}) and takes no arguments (confirmed by the NetScaler
// CLI `reset appfw learningdata`), and there is no GET endpoint reporting
// reset-state and no inverse API. Read/Update/Delete are therefore no-ops.
type AppfwlearningdataResetResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *AppfwlearningdataResetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningdata_reset"
}

func (r *AppfwlearningdataResetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwlearningdataResetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwlearningdata_reset resource.",
			},
			// The reset action accepts no attributes: NITRO's reset payload is
			// empty and the CLI `reset appfw learningdata` takes no arguments.
		},
	}
}

func (r *AppfwlearningdataResetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwlearningdataResetResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Resetting appfwlearningdata (action-only resource)")
	payload := appfwlearningdata_resetGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes reset as POST ?action=reset. Use ActOnResource with the
	// case-sensitive "reset" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Appfwlearningdata.Type(), &payload, "reset")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reset appfwlearningdata, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered appfwlearningdata reset")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("appfwlearningdata_reset")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataResetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// reset is a one-shot action. NITRO has no GET endpoint that reports
	// reset-state, so Read is a pure preserve-state no-op.
	var data AppfwlearningdataResetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for appfwlearningdata_reset; NITRO has no query endpoint for reset state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataResetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for reset; there are no config attributes to
	// change, so Terraform never invokes Update for a real change.
	var data, state AppfwlearningdataResetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwlearningdata_reset; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataResetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// reset is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-reset"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for appfwlearningdata_reset; NITRO has no inverse of the reset action")
}

func appfwlearningdata_resetGetThePayloadFromthePlan(ctx context.Context, data *AppfwlearningdataResetResourceModel) appfw.Appfwlearningdata {
	tflog.Debug(ctx, "In appfwlearningdata_resetGetThePayloadFromthePlan Function")

	// The reset action carries an empty payload ({"appfwlearningdata":{}}); it
	// accepts no attributes.
	appfwlearningdata := appfw.Appfwlearningdata{}

	return appfwlearningdata
}
