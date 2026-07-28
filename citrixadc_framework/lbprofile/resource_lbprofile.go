package lbprofile

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
var _ resource.Resource = &LbprofileResource{}
var _ resource.ResourceWithConfigure = (*LbprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LbprofileResource)(nil)

func NewLbprofileResource() resource.Resource {
	return &LbprofileResource{}
}

// LbprofileResource defines the resource implementation.
type LbprofileResource struct {
	client *service.NitroClient
}

func (r *LbprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbprofile"
}

func (r *LbprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config LbprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbprofile resource")
	// Get payload from plan (regular attributes)
	lbprofile := lbprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	lbprofileGetThePayloadFromtheConfig(ctx, &config, &lbprofile)

	// Make API call
	// Named resource - use AddResource
	lbprofilename_value := data.Lbprofilename.ValueString()
	_, err := r.client.AddResource(service.Lbprofile.Type(), lbprofilename_value, &lbprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Lbprofilename.ValueString()))

	// Read the updated state back
	if !r.readLbprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbprofile resource")

	found := r.readLbprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LbprofileResourceModel

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

	tflog.Debug(ctx, "Updating lbprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Computedadccookieattribute.Equal(state.Computedadccookieattribute) {
		tflog.Debug(ctx, fmt.Sprintf("computedadccookieattribute has changed for lbprofile"))
		hasChange = true
	}
	// Check secret attribute cookiepassphrase or its version tracker
	if !data.Cookiepassphrase.Equal(state.Cookiepassphrase) {
		tflog.Debug(ctx, fmt.Sprintf("cookiepassphrase has changed for lbprofile"))
		hasChange = true
	} else if !data.CookiepassphraseWoVersion.Equal(state.CookiepassphraseWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("cookiepassphrase_wo_version has changed for lbprofile"))
		hasChange = true
	}
	if !data.Dbslb.Equal(state.Dbslb) {
		tflog.Debug(ctx, fmt.Sprintf("dbslb has changed for lbprofile"))
		if config.Dbslb.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dbslb")
		} else {
			hasChange = true
		}
	}
	if !data.Httponlycookieflag.Equal(state.Httponlycookieflag) {
		tflog.Debug(ctx, fmt.Sprintf("httponlycookieflag has changed for lbprofile"))
		if config.Httponlycookieflag.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "httponlycookieflag")
		} else {
			hasChange = true
		}
	}
	if !data.Lbhashalgorithm.Equal(state.Lbhashalgorithm) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashalgorithm has changed for lbprofile"))
		if config.Lbhashalgorithm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "lbhashalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Lbhashfingers.Equal(state.Lbhashfingers) {
		tflog.Debug(ctx, fmt.Sprintf("lbhashfingers has changed for lbprofile"))
		if config.Lbhashfingers.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "lbhashfingers")
		} else {
			hasChange = true
		}
	}
	if !data.Literaladccookieattribute.Equal(state.Literaladccookieattribute) {
		tflog.Debug(ctx, fmt.Sprintf("literaladccookieattribute has changed for lbprofile"))
		hasChange = true
	}
	if !data.Processlocal.Equal(state.Processlocal) {
		tflog.Debug(ctx, fmt.Sprintf("processlocal has changed for lbprofile"))
		if config.Processlocal.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "processlocal")
		} else {
			hasChange = true
		}
	}
	if !data.Proximityfromself.Equal(state.Proximityfromself) {
		tflog.Debug(ctx, fmt.Sprintf("proximityfromself has changed for lbprofile"))
		if config.Proximityfromself.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "proximityfromself")
		} else {
			hasChange = true
		}
	}
	if !data.Storemqttclientidandusername.Equal(state.Storemqttclientidandusername) {
		tflog.Debug(ctx, fmt.Sprintf("storemqttclientidandusername has changed for lbprofile"))
		if config.Storemqttclientidandusername.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "storemqttclientidandusername")
		} else {
			hasChange = true
		}
	}
	if !data.Useencryptedpersistencecookie.Equal(state.Useencryptedpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("useencryptedpersistencecookie has changed for lbprofile"))
		if config.Useencryptedpersistencecookie.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "useencryptedpersistencecookie")
		} else {
			hasChange = true
		}
	}
	if !data.Usesecuredpersistencecookie.Equal(state.Usesecuredpersistencecookie) {
		tflog.Debug(ctx, fmt.Sprintf("usesecuredpersistencecookie has changed for lbprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		lbprofile := lbprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		lbprofileGetThePayloadFromtheConfig(ctx, &config, &lbprofile)
		// Make API call
		// Named resource - use UpdateResource
		lbprofilename_value := data.Lbprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Lbprofile.Type(), lbprofilename_value, &lbprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbprofile resource, skipping update")
	}

	// Unset attributes that were removed from the configuration (revert to ADC default)
	unsetIdPayload := map[string]interface{}{
		"lbprofilename": data.Lbprofilename.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Lbprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lbprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLbprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbprofile resource")
	// Named resource - delete using DeleteResource
	lbprofilename_value := data.Lbprofilename.ValueString()
	err := r.client.DeleteResource(service.Lbprofile.Type(), lbprofilename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbprofile resource")
}

// Helper function to read lbprofile data from API
func (r *LbprofileResource) readLbprofileFromApi(ctx context.Context, data *LbprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	lbprofilename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Lbprofile.Type(), lbprofilename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbprofile, got error: %s", err))
		return false
	}

	lbprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
