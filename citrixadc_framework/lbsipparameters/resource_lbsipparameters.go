package lbsipparameters

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
var _ resource.Resource = &LbsipparametersResource{}
var _ resource.ResourceWithConfigure = (*LbsipparametersResource)(nil)
var _ resource.ResourceWithImportState = (*LbsipparametersResource)(nil)

func NewLbsipparametersResource() resource.Resource {
	return &LbsipparametersResource{}
}

// LbsipparametersResource defines the resource implementation.
type LbsipparametersResource struct {
	client *service.NitroClient
}

func (r *LbsipparametersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbsipparametersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbsipparameters"
}

func (r *LbsipparametersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbsipparametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbsipparametersResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbsipparameters resource")

	lbsipparameters := lbsipparametersGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbsipparameters, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbsipparameters-config")

	tflog.Trace(ctx, "Created lbsipparameters resource")

	// Read the updated state back
	r.readLbsipparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbsipparametersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbsipparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbsipparameters resource")

	r.readLbsipparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbsipparametersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LbsipparametersResourceModel

	// Read Terraform prior state to detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbsipparameters resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Addrportvip.Equal(state.Addrportvip) {
		tflog.Debug(ctx, "addrportvip has changed for lbsipparameters")
		if config.Addrportvip.IsNull() {
			attributesToUnset = append(attributesToUnset, "addrportvip")
		} else {
			hasChange = true
		}
	}
	if !data.Retrydur.Equal(state.Retrydur) {
		tflog.Debug(ctx, "retrydur has changed for lbsipparameters")
		if config.Retrydur.IsNull() {
			attributesToUnset = append(attributesToUnset, "retrydur")
		} else {
			hasChange = true
		}
	}
	if !data.Rnatdstport.Equal(state.Rnatdstport) {
		tflog.Debug(ctx, "rnatdstport has changed for lbsipparameters")
		if config.Rnatdstport.IsNull() {
			attributesToUnset = append(attributesToUnset, "rnatdstport")
		} else {
			hasChange = true
		}
	}
	if !data.Rnatsecuredstport.Equal(state.Rnatsecuredstport) {
		tflog.Debug(ctx, "rnatsecuredstport has changed for lbsipparameters")
		if config.Rnatsecuredstport.IsNull() {
			attributesToUnset = append(attributesToUnset, "rnatsecuredstport")
		} else {
			hasChange = true
		}
	}
	if !data.Rnatsecuresrcport.Equal(state.Rnatsecuresrcport) {
		tflog.Debug(ctx, "rnatsecuresrcport has changed for lbsipparameters")
		if config.Rnatsecuresrcport.IsNull() {
			attributesToUnset = append(attributesToUnset, "rnatsecuresrcport")
		} else {
			hasChange = true
		}
	}
	if !data.Rnatsrcport.Equal(state.Rnatsrcport) {
		tflog.Debug(ctx, "rnatsrcport has changed for lbsipparameters")
		if config.Rnatsrcport.IsNull() {
			attributesToUnset = append(attributesToUnset, "rnatsrcport")
		} else {
			hasChange = true
		}
	}
	if !data.Sip503ratethreshold.Equal(state.Sip503ratethreshold) {
		tflog.Debug(ctx, "sip503ratethreshold has changed for lbsipparameters")
		if config.Sip503ratethreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "sip503ratethreshold")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		lbsipparameters := lbsipparametersGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lbsipparameters.Type(), &lbsipparameters)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbsipparameters, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbsipparameters resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbsipparameters resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Lbsipparameters.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lbsipparameters attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readLbsipparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbsipparametersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbsipparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbsipparameters resource")

	// For lbsipparameters, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbsipparameters resource from state")
}

// Helper function to read lbsipparameters data from API
func (r *LbsipparametersResource) readLbsipparametersFromApi(ctx context.Context, data *LbsipparametersResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbsipparameters.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbsipparameters, got error: %s", err))
		return
	}

	lbsipparametersSetAttrFromGet(ctx, data, getResponseData)

}
