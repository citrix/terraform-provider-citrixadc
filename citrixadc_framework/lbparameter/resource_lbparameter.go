package lbparameter

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
var _ resource.Resource = &LbparameterResource{}
var _ resource.ResourceWithConfigure = (*LbparameterResource)(nil)
var _ resource.ResourceWithImportState = (*LbparameterResource)(nil)

func NewLbparameterResource() resource.Resource {
	return &LbparameterResource{}
}

// LbparameterResource defines the resource implementation.
type LbparameterResource struct {
	client *service.NitroClient
}

func (r *LbparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbparameter"
}

func (r *LbparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config LbparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbparameter resource")
	// Get payload from plan (regular attributes)
	lbparameter := lbparameterGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	lbparameterGetThePayloadFromtheConfig(ctx, &config, &lbparameter)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("lbparameter-config")

	// Read the updated state back
	if !r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbparameter not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbparameter resource")

	found := r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LbparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Collect eligible attributes removed from config so they can be reverted
	// to their ADC default via a single ?action=unset call.
	attributesToUnset := []string{}
	if !data.Allowboundsvcremoval.Equal(state.Allowboundsvcremoval) {
		tflog.Debug(ctx, fmt.Sprintf("allowboundsvcremoval has changed for lbparameter"))
		if config.Allowboundsvcremoval.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "allowboundsvcremoval")
		} else {
			hasChange = true
		}
	}
	if !data.Computedadccookieattribute.Equal(state.Computedadccookieattribute) {
		tflog.Debug(ctx, fmt.Sprintf("computedadccookieattribute has changed for lbparameter"))
		hasChange = true
	}
	if !data.Consolidatedlconn.Equal(state.Consolidatedlconn) {
		tflog.Debug(ctx, fmt.Sprintf("consolidatedlconn has changed for lbparameter"))
		if config.Consolidatedlconn.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "consolidatedlconn")
		} else {
			hasChange = true
		}
	}
	// Check secret attribute cookiepassphrase or its version tracker
	if !data.Cookiepassphrase.Equal(state.Cookiepassphrase) {
		tflog.Debug(ctx, fmt.Sprintf("cookiepassphrase has changed for lbparameter"))
		hasChange = true
	} else if !data.CookiepassphraseWoVersion.Equal(state.CookiepassphraseWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("cookiepassphrase_wo_version has changed for lbparameter"))
		hasChange = true
	}
	if !data.Dbsttl.Equal(state.Dbsttl) {
		tflog.Debug(ctx, fmt.Sprintf("dbsttl has changed for lbparameter"))
		if config.Dbsttl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dbsttl")
		} else {
			hasChange = true
		}
	}
	if !data.Dropmqttjumbomessage.Equal(state.Dropmqttjumbomessage) {
		tflog.Debug(ctx, fmt.Sprintf("dropmqttjumbomessage has changed for lbparameter"))
		if config.Dropmqttjumbomessage.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dropmqttjumbomessage")
		} else {
			hasChange = true
		}
	}
	if !data.Httponlycookieflag.Equal(state.Httponlycookieflag) {
		tflog.Debug(ctx, fmt.Sprintf("httponlycookieflag has changed for lbparameter"))
		if config.Httponlycookieflag.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "httponlycookieflag")
		} else {
			hasChange = true
		}
	}
	if !data.Lbhashalgorithm.Equal(state.Lbhashalgorithm) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashalgorithm has changed for lbparameter"))
		if config.Lbhashalgorithm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "lbhashalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Lbhashfingers.Equal(state.Lbhashfingers) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashfingers has changed for lbparameter"))
		if config.Lbhashfingers.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "lbhashfingers")
		} else {
			hasChange = true
		}
	}
	if !data.Literaladccookieattribute.Equal(state.Literaladccookieattribute) {
		tflog.Debug(ctx, fmt.Sprintf("literaladccookieattribute has changed for lbparameter"))
		hasChange = true
	}
	if !data.Maxpipelinenat.Equal(state.Maxpipelinenat) {
		tflog.Debug(ctx, fmt.Sprintf("maxpipelinenat has changed for lbparameter"))
		hasChange = true
	}
	if !data.Monitorconnectionclose.Equal(state.Monitorconnectionclose) {
		tflog.Debug(ctx, fmt.Sprintf("monitorconnectionclose has changed for lbparameter"))
		if config.Monitorconnectionclose.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "monitorconnectionclose")
		} else {
			hasChange = true
		}
	}
	if !data.Monitorskipmaxclient.Equal(state.Monitorskipmaxclient) {
		tflog.Debug(ctx, fmt.Sprintf("monitorskipmaxclient has changed for lbparameter"))
		if config.Monitorskipmaxclient.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "monitorskipmaxclient")
		} else {
			hasChange = true
		}
	}
	if !data.Preferdirectroute.Equal(state.Preferdirectroute) {
		tflog.Debug(ctx, fmt.Sprintf("preferdirectroute has changed for lbparameter"))
		if config.Preferdirectroute.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preferdirectroute")
		} else {
			hasChange = true
		}
	}
	if !data.Proximityfromself.Equal(state.Proximityfromself) {
		tflog.Debug(ctx, fmt.Sprintf("proximityfromself has changed for lbparameter"))
		if config.Proximityfromself.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proximityfromself")
		} else {
			hasChange = true
		}
	}
	if !data.Retainservicestate.Equal(state.Retainservicestate) {
		tflog.Debug(ctx, fmt.Sprintf("retainservicestate has changed for lbparameter"))
		if config.Retainservicestate.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "retainservicestate")
		} else {
			hasChange = true
		}
	}
	if !data.Startuprrfactor.Equal(state.Startuprrfactor) {
		tflog.Debug(ctx, fmt.Sprintf("startuprrfactor has changed for lbparameter"))
		hasChange = true
	}
	if !data.Storemqttclientidandusername.Equal(state.Storemqttclientidandusername) {
		tflog.Debug(ctx, fmt.Sprintf("storemqttclientidandusername has changed for lbparameter"))
		if config.Storemqttclientidandusername.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "storemqttclientidandusername")
		} else {
			hasChange = true
		}
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, fmt.Sprintf("undefaction has changed for lbparameter"))
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}
	if !data.Useencryptedpersistencecookie.Equal(state.Useencryptedpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("useencryptedpersistencecookie has changed for lbparameter"))
		if config.Useencryptedpersistencecookie.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "useencryptedpersistencecookie")
		} else {
			hasChange = true
		}
	}
	if !data.Useportforhashlb.Equal(state.Useportforhashlb) {
		tflog.Debug(ctx, fmt.Sprintf("useportforhashlb has changed for lbparameter"))
		if config.Useportforhashlb.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "useportforhashlb")
		} else {
			hasChange = true
		}
	}
	if !data.Usesecuredpersistencecookie.Equal(state.Usesecuredpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("usesecuredpersistencecookie has changed for lbparameter"))
		hasChange = true
	}
	if !data.Vserverspecificmac.Equal(state.Vserverspecificmac) {
		tflog.Debug(ctx, fmt.Sprintf("vserverspecificmac has changed for lbparameter"))
		if config.Vserverspecificmac.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "vserverspecificmac")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		lbparameter := lbparameterGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		lbparameterGetThePayloadFromtheConfig(ctx, &config, &lbparameter)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbparameter resource, skipping update")
	}

	// Revert attributes removed from config to their ADC default via a single
	// batched unset. Done after the update so any default carried in the update
	// payload for a removed attribute is superseded by the unset.
	// lbparameter is a singleton resource, so no identity fields are required.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Lbparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lbparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbparameter not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbparameter resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed lbparameter from Terraform state")
}

// Helper function to read lbparameter data from API
func (r *LbparameterResource) readLbparameterFromApi(ctx context.Context, data *LbparameterResourceModel, diags *diag.Diagnostics) bool {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Lbparameter.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbparameter, got error: %s", err))
		return false
	}

	lbparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
