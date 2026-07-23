package quicprofile

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
var _ resource.Resource = &QuicprofileResource{}
var _ resource.ResourceWithConfigure = (*QuicprofileResource)(nil)
var _ resource.ResourceWithImportState = (*QuicprofileResource)(nil)

func NewQuicprofileResource() resource.Resource {
	return &QuicprofileResource{}
}

// QuicprofileResource defines the resource implementation.
type QuicprofileResource struct {
	client *service.NitroClient
}

func (r *QuicprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *QuicprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quicprofile"
}

func (r *QuicprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *QuicprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data QuicprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating quicprofile resource")
	quicprofile := quicprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Quicprofile.Type(), name_value, &quicprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create quicprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created quicprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readQuicprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuicprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data QuicprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading quicprofile resource")

	r.readQuicprofileFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// If the resource was deleted out-of-band, remove it from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuicprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state, config QuicprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating quicprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Collect eligible attributes removed from config so they can be unset (reverted to ADC default)
	attributesToUnset := []string{}
	if !data.Ackdelayexponent.Equal(state.Ackdelayexponent) {
		tflog.Debug(ctx, fmt.Sprintf("ackdelayexponent has changed for quicprofile"))
		if config.Ackdelayexponent.IsNull() {
			attributesToUnset = append(attributesToUnset, "ackdelayexponent")
		} else {
			hasChange = true
		}
	}
	if !data.Activeconnectionidlimit.Equal(state.Activeconnectionidlimit) {
		tflog.Debug(ctx, fmt.Sprintf("activeconnectionidlimit has changed for quicprofile"))
		if config.Activeconnectionidlimit.IsNull() {
			attributesToUnset = append(attributesToUnset, "activeconnectionidlimit")
		} else {
			hasChange = true
		}
	}
	if !data.Activeconnectionmigration.Equal(state.Activeconnectionmigration) {
		tflog.Debug(ctx, fmt.Sprintf("activeconnectionmigration has changed for quicprofile"))
		if config.Activeconnectionmigration.IsNull() {
			attributesToUnset = append(attributesToUnset, "activeconnectionmigration")
		} else {
			hasChange = true
		}
	}
	if !data.Congestionctrlalgorithm.Equal(state.Congestionctrlalgorithm) {
		tflog.Debug(ctx, fmt.Sprintf("congestionctrlalgorithm has changed for quicprofile"))
		if config.Congestionctrlalgorithm.IsNull() {
			attributesToUnset = append(attributesToUnset, "congestionctrlalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxdata.Equal(state.Initialmaxdata) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxdata has changed for quicprofile"))
		if config.Initialmaxdata.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxdata")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxstreamdatabidilocal.Equal(state.Initialmaxstreamdatabidilocal) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxstreamdatabidilocal has changed for quicprofile"))
		if config.Initialmaxstreamdatabidilocal.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxstreamdatabidilocal")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxstreamdatabidiremote.Equal(state.Initialmaxstreamdatabidiremote) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxstreamdatabidiremote has changed for quicprofile"))
		if config.Initialmaxstreamdatabidiremote.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxstreamdatabidiremote")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxstreamdatauni.Equal(state.Initialmaxstreamdatauni) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxstreamdatauni has changed for quicprofile"))
		if config.Initialmaxstreamdatauni.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxstreamdatauni")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxstreamsbidi.Equal(state.Initialmaxstreamsbidi) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxstreamsbidi has changed for quicprofile"))
		if config.Initialmaxstreamsbidi.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxstreamsbidi")
		} else {
			hasChange = true
		}
	}
	if !data.Initialmaxstreamsuni.Equal(state.Initialmaxstreamsuni) {
		tflog.Debug(ctx, fmt.Sprintf("initialmaxstreamsuni has changed for quicprofile"))
		if config.Initialmaxstreamsuni.IsNull() {
			attributesToUnset = append(attributesToUnset, "initialmaxstreamsuni")
		} else {
			hasChange = true
		}
	}
	if !data.Maxackdelay.Equal(state.Maxackdelay) {
		tflog.Debug(ctx, fmt.Sprintf("maxackdelay has changed for quicprofile"))
		if config.Maxackdelay.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxackdelay")
		} else {
			hasChange = true
		}
	}
	if !data.Maxidletimeout.Equal(state.Maxidletimeout) {
		tflog.Debug(ctx, fmt.Sprintf("maxidletimeout has changed for quicprofile"))
		if config.Maxidletimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxidletimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Maxudpdatagramsperburst.Equal(state.Maxudpdatagramsperburst) {
		tflog.Debug(ctx, fmt.Sprintf("maxudpdatagramsperburst has changed for quicprofile"))
		if config.Maxudpdatagramsperburst.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxudpdatagramsperburst")
		} else {
			hasChange = true
		}
	}
	if !data.Maxudppayloadsize.Equal(state.Maxudppayloadsize) {
		tflog.Debug(ctx, fmt.Sprintf("maxudppayloadsize has changed for quicprofile"))
		if config.Maxudppayloadsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxudppayloadsize")
		} else {
			hasChange = true
		}
	}
	if !data.Newtokenvalidityperiod.Equal(state.Newtokenvalidityperiod) {
		tflog.Debug(ctx, fmt.Sprintf("newtokenvalidityperiod has changed for quicprofile"))
		if config.Newtokenvalidityperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "newtokenvalidityperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Retrytokenvalidityperiod.Equal(state.Retrytokenvalidityperiod) {
		tflog.Debug(ctx, fmt.Sprintf("retrytokenvalidityperiod has changed for quicprofile"))
		if config.Retrytokenvalidityperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "retrytokenvalidityperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Statelessaddressvalidation.Equal(state.Statelessaddressvalidation) {
		tflog.Debug(ctx, fmt.Sprintf("statelessaddressvalidation has changed for quicprofile"))
		if config.Statelessaddressvalidation.IsNull() {
			attributesToUnset = append(attributesToUnset, "statelessaddressvalidation")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		quicprofile := quicprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Quicprofile.Type(), name_value, &quicprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update quicprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated quicprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for quicprofile resource, skipping update")
	}

	// Unset (revert to ADC default) any eligible attributes removed from config.
	// Update-then-unset ordering ensures any default carried in the update payload is superseded.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Quicprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset quicprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readQuicprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuicprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data QuicprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting quicprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Quicprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete quicprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted quicprofile resource")
}

// Helper function to read quicprofile data from API
func (r *QuicprofileResource) readQuicprofileFromApi(ctx context.Context, data *QuicprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Quicprofile.Type(), name_Name)
	if err != nil {
		// Resource is missing (deleted out-of-band); signal removal to Read.
		if utils.IsNotFoundError(err) {
			tflog.Warn(ctx, fmt.Sprintf("quicprofile %s not found, removing from state", name_Name))
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read quicprofile, got error: %s", err))
		return
	}

	quicprofileSetAttrFromGet(ctx, data, getResponseData)

}
