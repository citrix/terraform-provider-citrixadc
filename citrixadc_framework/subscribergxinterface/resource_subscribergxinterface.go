package subscribergxinterface

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
var _ resource.Resource = &SubscribergxinterfaceResource{}
var _ resource.ResourceWithConfigure = (*SubscribergxinterfaceResource)(nil)
var _ resource.ResourceWithImportState = (*SubscribergxinterfaceResource)(nil)

func NewSubscribergxinterfaceResource() resource.Resource {
	return &SubscribergxinterfaceResource{}
}

// SubscribergxinterfaceResource defines the resource implementation.
type SubscribergxinterfaceResource struct {
	client *service.NitroClient
}

func (r *SubscribergxinterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SubscribergxinterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscribergxinterface"
}

func (r *SubscribergxinterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscribergxinterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscribergxinterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating subscribergxinterface resource")

	// subscribergxinterface := subscribergxinterfaceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Subscribergxinterface.Type(), &subscribergxinterface)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create subscribergxinterface, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("subscribergxinterface-config")

	tflog.Trace(ctx, "Created subscribergxinterface resource")

	// Read the updated state back
	r.readSubscribergxinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscribergxinterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscribergxinterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading subscribergxinterface resource")

	r.readSubscribergxinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscribergxinterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SubscribergxinterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating subscribergxinterface resource")

	// Create API request body from the model
	// subscribergxinterface := subscribergxinterfaceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Subscribergxinterface.Type(), &subscribergxinterface)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update subscribergxinterface, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated subscribergxinterface resource")

	// Read the updated state back
	r.readSubscribergxinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscribergxinterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SubscribergxinterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting subscribergxinterface resource")

	// For subscribergxinterface, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted subscribergxinterface resource from state")
}

// Helper function to read subscribergxinterface data from API
func (r *SubscribergxinterfaceResource) readSubscribergxinterfaceFromApi(ctx context.Context, data *SubscribergxinterfaceResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Subscribergxinterface.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read subscribergxinterface, got error: %s", err))
		return
	}

	subscribergxinterfaceSetAttrFromGet(ctx, data, getResponseData)

}
