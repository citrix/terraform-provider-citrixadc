package policyhttpcallout

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicyhttpcalloutResource{}
var _ resource.ResourceWithConfigure = (*PolicyhttpcalloutResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyhttpcalloutResource)(nil)

func NewPolicyhttpcalloutResource() resource.Resource {
	return &PolicyhttpcalloutResource{}
}

// PolicyhttpcalloutResource defines the resource implementation.
type PolicyhttpcalloutResource struct {
	client *service.NitroClient
}

func (r *PolicyhttpcalloutResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyhttpcalloutResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyhttpcallout"
}

func (r *PolicyhttpcalloutResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyhttpcalloutResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyhttpcallout resource")

	// Create API request body from the model
	policyhttpcallout := policyhttpcalloutGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	policyhttpcalloutName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Policyhttpcallout.Type(), policyhttpcalloutName, &policyhttpcallout)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create policyhttpcallout, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created policyhttpcallout resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(policyhttpcalloutName)

	// Read the updated state back
	if !r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policyhttpcallout not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading policyhttpcallout resource")

	found := r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state, config PolicyhttpcalloutResourceModel

	// Read Terraform prior state to preserve ID and for change detection
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes the user removed from config
	// (plan marks removed Optional+Computed attrs as unknown, config keeps them null).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating policyhttpcallout resource")

	// NITRO enforces mutual exclusivity between fullreqexpr and the request-shaping
	// attributes (bodyexpr, httpmethod, hostexpr, urlstemexpr, headers, parameters):
	// setting fullreqexpr while any of those is still present on the appliance fails
	// with errorcode 703 ("Full request expression and other request attributes
	// cannot be set at the same time"). NITRO also rejects clearing these via an empty
	// value in the update payload. So, for any mutually-exclusive request-shaping
	// attribute the user removed from config (null in config) that still holds a value
	// in prior state, clear it on the appliance with an ?action=unset call before the
	// update PUT. This preserves the SDK v2 user-facing contract (mode switching such
	// as httpmethod-style -> fullReqExpr-style works in a single apply).
	policyhttpcalloutName := data.Name.ValueString()
	unsetPayload := map[string]interface{}{"name": policyhttpcalloutName}
	needUnset := false
	markUnsetString := func(name string, cfg, st types.String) {
		if cfg.IsNull() && !st.IsNull() && st.ValueString() != "" {
			unsetPayload[name] = true
			needUnset = true
		}
	}
	markUnsetList := func(name string, cfg, st types.List) {
		if cfg.IsNull() && !st.IsNull() && len(st.Elements()) > 0 {
			unsetPayload[name] = true
			needUnset = true
		}
	}
	markUnsetString("fullreqexpr", config.Fullreqexpr, state.Fullreqexpr)
	markUnsetString("bodyexpr", config.Bodyexpr, state.Bodyexpr)
	markUnsetString("hostexpr", config.Hostexpr, state.Hostexpr)
	markUnsetString("urlstemexpr", config.Urlstemexpr, state.Urlstemexpr)
	markUnsetString("httpmethod", config.Httpmethod, state.Httpmethod)
	markUnsetList("headers", config.Headers, state.Headers)
	markUnsetList("parameters", config.Parameters, state.Parameters)

	if needUnset {
		tflog.Debug(ctx, "Unsetting mutually-exclusive request-shaping attributes removed from config")
		if err := r.client.ActOnResource(service.Policyhttpcallout.Type(), unsetPayload, "unset"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset policyhttpcallout attributes, got error: %s", err))
			return
		}
	}

	// Check if there are any changes in updateable attributes.
	// name and returntype are ForceNew/RequiresReplace and never reach Update.
	hasChange := false
	// attributesToUnset collects independent (non-mutually-exclusive) attributes
	// removed from config so they are reverted to NITRO defaults with a single
	// ?action=unset call after the update. The mutually-exclusive request-shaping
	// attributes above are handled by the pre-update unset block (order matters).
	attributesToUnset := []string{}
	if !data.Bodyexpr.Equal(state.Bodyexpr) {
		tflog.Debug(ctx, "bodyexpr has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Cacheforsecs.Equal(state.Cacheforsecs) {
		tflog.Debug(ctx, "cacheforsecs has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for policyhttpcallout")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Fullreqexpr.Equal(state.Fullreqexpr) {
		tflog.Debug(ctx, "fullreqexpr has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Headers.Equal(state.Headers) {
		tflog.Debug(ctx, "headers has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Hostexpr.Equal(state.Hostexpr) {
		tflog.Debug(ctx, "hostexpr has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Httpmethod.Equal(state.Httpmethod) {
		tflog.Debug(ctx, "httpmethod has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, "ipaddress has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Parameters.Equal(state.Parameters) {
		tflog.Debug(ctx, "parameters has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Port.Equal(state.Port) {
		tflog.Debug(ctx, "port has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Resultexpr.Equal(state.Resultexpr) {
		tflog.Debug(ctx, "resultexpr has changed for policyhttpcallout")
		if config.Resultexpr.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "resultexpr")
		} else {
			hasChange = true
		}
	}
	if !data.Scheme.Equal(state.Scheme) {
		tflog.Debug(ctx, "scheme has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Urlstemexpr.Equal(state.Urlstemexpr) {
		tflog.Debug(ctx, "urlstemexpr has changed for policyhttpcallout")
		hasChange = true
	}
	if !data.Vserver.Equal(state.Vserver) {
		tflog.Debug(ctx, "vserver has changed for policyhttpcallout")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model (updatable fields only)
		policyhttpcallout := policyhttpcalloutGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		policyhttpcalloutName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Policyhttpcallout.Type(), policyhttpcalloutName, &policyhttpcallout)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update policyhttpcallout, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated policyhttpcallout resource")
	} else {
		tflog.Debug(ctx, "No changes detected for policyhttpcallout resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Policyhttpcallout.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset policyhttpcallout attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readPolicyhttpcalloutFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "policyhttpcallout not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyhttpcalloutResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data PolicyhttpcalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting policyhttpcallout resource")

	// Named resource - delete using DeleteResource keyed off the ID (live name)
	policyhttpcalloutName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Policyhttpcallout.Type(), policyhttpcalloutName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete policyhttpcallout, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted policyhttpcallout resource")
}

// Helper function to read policyhttpcallout data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *PolicyhttpcalloutResource) readPolicyhttpcalloutFromApi(ctx context.Context, data *PolicyhttpcalloutResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	policyhttpcalloutName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Policyhttpcallout.Type(), policyhttpcalloutName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read policyhttpcallout, got error: %s", err))
		return false
	}

	policyhttpcalloutSetAttrFromGet(ctx, data, getResponseData)

	return true
}
