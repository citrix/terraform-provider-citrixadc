package tmsessionaction

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
var _ resource.Resource = &TmsessionactionResource{}
var _ resource.ResourceWithConfigure = (*TmsessionactionResource)(nil)
var _ resource.ResourceWithImportState = (*TmsessionactionResource)(nil)

func NewTmsessionactionResource() resource.Resource {
	return &TmsessionactionResource{}
}

// TmsessionactionResource defines the resource implementation.
type TmsessionactionResource struct {
	client *service.NitroClient
}

func (r *TmsessionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmsessionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmsessionaction"
}

func (r *TmsessionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmsessionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmsessionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmsessionaction resource")

	tmsessionaction := tmsessionactionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (POST), matching SDK v2 behavior
	tmsessionactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmsessionaction.Type(), tmsessionactionName, &tmsessionaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmsessionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmsessionaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(tmsessionactionName)

	// Read the updated state back
	if !r.readTmsessionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsessionaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmsessionaction resource")

	found := r.readTmsessionactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmsessionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmsessionactionResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmsessionaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Defaultauthorizationaction.Equal(state.Defaultauthorizationaction) {
		tflog.Debug(ctx, "defaultauthorizationaction has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Homepage.Equal(state.Homepage) {
		tflog.Debug(ctx, "homepage has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Httponlycookie.Equal(state.Httponlycookie) {
		tflog.Debug(ctx, "httponlycookie has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		tflog.Debug(ctx, "kcdaccount has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Persistentcookie.Equal(state.Persistentcookie) {
		tflog.Debug(ctx, "persistentcookie has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Persistentcookievalidity.Equal(state.Persistentcookievalidity) {
		tflog.Debug(ctx, "persistentcookievalidity has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		tflog.Debug(ctx, "sesstimeout has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Sso.Equal(state.Sso) {
		tflog.Debug(ctx, "sso has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Ssocredential.Equal(state.Ssocredential) {
		tflog.Debug(ctx, "ssocredential has changed for tmsessionaction")
		hasChange = true
	}
	if !data.Ssodomain.Equal(state.Ssodomain) {
		tflog.Debug(ctx, "ssodomain has changed for tmsessionaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		tmsessionaction := tmsessionactionGetThePayloadFromtheConfig(ctx, &data)
		// Named resource with an unnamed (PUT) update endpoint - matching SDK v2 behavior
		err := r.client.UpdateUnnamedResource(service.Tmsessionaction.Type(), &tmsessionaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmsessionaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated tmsessionaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tmsessionaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readTmsessionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmsessionaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmsessionaction resource")

	// Named resource - delete using DeleteResource, matching SDK v2 behavior
	tmsessionactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmsessionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmsessionaction resource")
}

// Helper function to read tmsessionaction data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *TmsessionactionResource) readTmsessionactionFromApi(ctx context.Context, data *TmsessionactionResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain name value
	tmsessionactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tmsessionaction.Type(), tmsessionactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmsessionaction, got error: %s", err))
		return false
	}

	tmsessionactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
