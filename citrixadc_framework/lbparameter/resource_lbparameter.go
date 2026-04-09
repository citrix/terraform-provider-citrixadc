package lbparameter

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
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
	r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics)

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
	if !data.Allowboundsvcremoval.Equal(state.Allowboundsvcremoval) {
		tflog.Debug(ctx, fmt.Sprintf("allowboundsvcremoval has changed for lbparameter"))
		hasChange = true
	}
	if !data.Computedadccookieattribute.Equal(state.Computedadccookieattribute) {
		tflog.Debug(ctx, fmt.Sprintf("computedadccookieattribute has changed for lbparameter"))
		hasChange = true
	}
	if !data.Consolidatedlconn.Equal(state.Consolidatedlconn) {
		tflog.Debug(ctx, fmt.Sprintf("consolidatedlconn has changed for lbparameter"))
		hasChange = true
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
		hasChange = true
	}
	if !data.Dropmqttjumbomessage.Equal(state.Dropmqttjumbomessage) {
		tflog.Debug(ctx, fmt.Sprintf("dropmqttjumbomessage has changed for lbparameter"))
		hasChange = true
	}
	if !data.Httponlycookieflag.Equal(state.Httponlycookieflag) {
		tflog.Debug(ctx, fmt.Sprintf("httponlycookieflag has changed for lbparameter"))
		hasChange = true
	}
	if !data.Lbhashalgorithm.Equal(state.Lbhashalgorithm) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashalgorithm has changed for lbparameter"))
		hasChange = true
	}
	if !data.Lbhashfingers.Equal(state.Lbhashfingers) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashfingers has changed for lbparameter"))
		hasChange = true
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
		hasChange = true
	}
	if !data.Monitorskipmaxclient.Equal(state.Monitorskipmaxclient) {
		tflog.Debug(ctx, fmt.Sprintf("monitorskipmaxclient has changed for lbparameter"))
		hasChange = true
	}
	if !data.Preferdirectroute.Equal(state.Preferdirectroute) {
		tflog.Debug(ctx, fmt.Sprintf("preferdirectroute has changed for lbparameter"))
		hasChange = true
	}
	if !data.Proximityfromself.Equal(state.Proximityfromself) {
		tflog.Debug(ctx, fmt.Sprintf("proximityfromself has changed for lbparameter"))
		hasChange = true
	}
	if !data.Retainservicestate.Equal(state.Retainservicestate) {
		tflog.Debug(ctx, fmt.Sprintf("retainservicestate has changed for lbparameter"))
		hasChange = true
	}
	if !data.Startuprrfactor.Equal(state.Startuprrfactor) {
		tflog.Debug(ctx, fmt.Sprintf("startuprrfactor has changed for lbparameter"))
		hasChange = true
	}
	if !data.Storemqttclientidandusername.Equal(state.Storemqttclientidandusername) {
		tflog.Debug(ctx, fmt.Sprintf("storemqttclientidandusername has changed for lbparameter"))
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, fmt.Sprintf("undefaction has changed for lbparameter"))
		hasChange = true
	}
	if !data.Useencryptedpersistencecookie.Equal(state.Useencryptedpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("useencryptedpersistencecookie has changed for lbparameter"))
		hasChange = true
	}
	if !data.Useportforhashlb.Equal(state.Useportforhashlb) {
		tflog.Debug(ctx, fmt.Sprintf("useportforhashlb has changed for lbparameter"))
		hasChange = true
	}
	if !data.Usesecuredpersistencecookie.Equal(state.Usesecuredpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("usesecuredpersistencecookie has changed for lbparameter"))
		hasChange = true
	}
	if !data.Vserverspecificmac.Equal(state.Vserverspecificmac) {
		tflog.Debug(ctx, fmt.Sprintf("vserverspecificmac has changed for lbparameter"))
		hasChange = true
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

	// Read the updated state back
	r.readLbparameterFromApi(ctx, &data, &resp.Diagnostics)

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
func (r *LbparameterResource) readLbparameterFromApi(ctx context.Context, data *LbparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Lbparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbparameter, got error: %s", err))
		return
	}

	lbparameterSetAttrFromGet(ctx, data, getResponseData)

}
