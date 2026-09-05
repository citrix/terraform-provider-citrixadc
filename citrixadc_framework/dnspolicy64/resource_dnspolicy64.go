package dnspolicy64

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
var _ resource.Resource = &Dnspolicy64Resource{}
var _ resource.ResourceWithConfigure = (*Dnspolicy64Resource)(nil)
var _ resource.ResourceWithImportState = (*Dnspolicy64Resource)(nil)

func NewDnspolicy64Resource() resource.Resource {
	return &Dnspolicy64Resource{}
}

// Dnspolicy64Resource defines the resource implementation.
type Dnspolicy64Resource struct {
	client *service.NitroClient
}

func (r *Dnspolicy64Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Dnspolicy64Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnspolicy64"
}

func (r *Dnspolicy64Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Dnspolicy64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnspolicy64 resource")

	// Create API request body from the model
	dnspolicy64 := dnspolicy64GetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	dnspolicy64Name := data.Name.ValueString()
	_, err := r.client.AddResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnspolicy64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnspolicy64 resource")

	// Set ID for the resource before reading state
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(dnspolicy64Name)

	// Read the updated state back
	if !r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnspolicy64 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnspolicy64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnspolicy64 resource")

	found := r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Dnspolicy64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Dnspolicy64ResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnspolicy64 resource")

	// Check if there are any changes in updateable attributes
	// name is RequiresReplace (ForceNew in SDK v2) - never reaches Update
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for dnspolicy64")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for dnspolicy64")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		dnspolicy64 := dnspolicy64GetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		dnspolicy64Name := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnspolicy64, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dnspolicy64 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dnspolicy64 resource, skipping update")
	}

	// Read the updated state back
	if !r.readDnspolicy64FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnspolicy64 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnspolicy64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Dnspolicy64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnspolicy64 resource")
	// Named resource - delete using DeleteResource
	dnspolicy64Name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnspolicy64, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnspolicy64 resource")
}

// Helper function to read dnspolicy64 data from API
func (r *Dnspolicy64Resource) readDnspolicy64FromApi(ctx context.Context, data *Dnspolicy64ResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	dnspolicy64Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnspolicy64, got error: %s", err))
		return false
	}

	dnspolicy64SetAttrFromGet(ctx, data, getResponseData)

	return true
}
