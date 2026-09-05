package transformaction

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
var _ resource.Resource = &TransformactionResource{}
var _ resource.ResourceWithConfigure = (*TransformactionResource)(nil)
var _ resource.ResourceWithImportState = (*TransformactionResource)(nil)

func NewTransformactionResource() resource.Resource {
	return &TransformactionResource{}
}

// TransformactionResource defines the resource implementation.
type TransformactionResource struct {
	client *service.NitroClient
}

func (r *TransformactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformaction"
}

func (r *TransformactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformaction resource")
	transformactionName := data.Name.ValueString()

	// Named resource. NITRO's add for transformaction does not accept all attributes,
	// so mirror SDK v2: add with the limited set (name/profilename/state/priority),
	// then update with the full set of transformation patterns.
	addPayload := transformactionGetTheAddPayloadFromthePlan(ctx, &data)
	_, err := r.client.AddResource(service.Transformaction.Type(), transformactionName, &addPayload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformaction, got error: %s", err))
		return
	}

	updatePayload := transformactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
	_, err = r.client.UpdateResource(service.Transformaction.Type(), transformactionName, &updatePayload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created transformaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(transformactionName)

	// Read the updated state back
	if !r.readTransformactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformaction resource")

	found := r.readTransformactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TransformactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TransformactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating transformaction resource")

	// name and profilename are ForceNew (RequiresReplace) and never reach Update as a
	// change. Detect changes only on the updateable attributes. For unsettable
	// attributes removed from config, collect them so the appliance reverts them.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		if config.Comment.IsNull() {
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Cookiedomainfrom.Equal(state.Cookiedomainfrom) {
		if config.Cookiedomainfrom.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookiedomainfrom")
		} else {
			hasChange = true
		}
	}
	if !data.Cookiedomaininto.Equal(state.Cookiedomaininto) {
		if config.Cookiedomaininto.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookiedomaininto")
		} else {
			hasChange = true
		}
	}
	if !data.Priority.Equal(state.Priority) {
		hasChange = true
	}
	if !data.Requrlfrom.Equal(state.Requrlfrom) {
		if config.Requrlfrom.IsNull() {
			attributesToUnset = append(attributesToUnset, "requrlfrom")
		} else {
			hasChange = true
		}
	}
	if !data.Requrlinto.Equal(state.Requrlinto) {
		if config.Requrlinto.IsNull() {
			attributesToUnset = append(attributesToUnset, "requrlinto")
		} else {
			hasChange = true
		}
	}
	if !data.Resurlfrom.Equal(state.Resurlfrom) {
		if config.Resurlfrom.IsNull() {
			attributesToUnset = append(attributesToUnset, "resurlfrom")
		} else {
			hasChange = true
		}
	}
	if !data.Resurlinto.Equal(state.Resurlinto) {
		if config.Resurlinto.IsNull() {
			attributesToUnset = append(attributesToUnset, "resurlinto")
		} else {
			hasChange = true
		}
	}
	if !data.State.Equal(state.State) {
		if config.State.IsNull() {
			attributesToUnset = append(attributesToUnset, "state")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		updatePayload := transformactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		_, err := r.client.UpdateResource(service.Transformaction.Type(), data.Id.ValueString(), &updatePayload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated transformaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for transformaction resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Done after any update so a default the update payload
	// carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Transformaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset transformaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readTransformactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformaction resource")

	// Named resource - delete using DeleteResource keyed by the ID (name).
	err := r.client.DeleteResource(service.Transformaction.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete transformaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted transformaction resource")
}

// Helper function to read transformaction data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *TransformactionResource) readTransformactionFromApi(ctx context.Context, data *TransformactionResourceModel, diags *diag.Diagnostics) bool {
	transformactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Transformaction.Type(), transformactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformaction, got error: %s", err))
		return false
	}

	transformactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
