package appfwprofile_creditcardnumber_binding

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
var _ resource.Resource = &AppfwprofileCreditcardnumberBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCreditcardnumberBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCreditcardnumberBindingResource)(nil)

func NewAppfwprofileCreditcardnumberBindingResource() resource.Resource {
	return &AppfwprofileCreditcardnumberBindingResource{}
}

// AppfwprofileCreditcardnumberBindingResource defines the resource implementation.
type AppfwprofileCreditcardnumberBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCreditcardnumberBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCreditcardnumberBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_creditcardnumber_binding"
}

func (r *AppfwprofileCreditcardnumberBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCreditcardnumberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCreditcardnumberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_creditcardnumber_binding resource")

	// appfwprofile_creditcardnumber_binding := appfwprofile_creditcardnumber_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_creditcardnumber_binding.Type(), &appfwprofile_creditcardnumber_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_creditcardnumber_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_creditcardnumber_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_creditcardnumber_binding resource")

	// Read the updated state back
	r.readAppfwprofileCreditcardnumberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCreditcardnumberBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCreditcardnumberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_creditcardnumber_binding resource")

	r.readAppfwprofileCreditcardnumberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCreditcardnumberBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileCreditcardnumberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_creditcardnumber_binding resource")

	// Create API request body from the model
	// appfwprofile_creditcardnumber_binding := appfwprofile_creditcardnumber_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_creditcardnumber_binding.Type(), &appfwprofile_creditcardnumber_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_creditcardnumber_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_creditcardnumber_binding resource")

	// Read the updated state back
	r.readAppfwprofileCreditcardnumberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCreditcardnumberBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCreditcardnumberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_creditcardnumber_binding resource")

	// For appfwprofile_creditcardnumber_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_creditcardnumber_binding resource from state")
}

// Helper function to read appfwprofile_creditcardnumber_binding data from API
func (r *AppfwprofileCreditcardnumberBindingResource) readAppfwprofileCreditcardnumberBindingFromApi(ctx context.Context, data *AppfwprofileCreditcardnumberBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_creditcardnumber_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_creditcardnumber_binding, got error: %s", err))
		return
	}

	appfwprofile_creditcardnumber_bindingSetAttrFromGet(ctx, data, getResponseData)

}
