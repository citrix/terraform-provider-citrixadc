package appfwprofile_sqlinjection_binding

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
var _ resource.Resource = &AppfwprofileSqlinjectionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileSqlinjectionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileSqlinjectionBindingResource)(nil)

func NewAppfwprofileSqlinjectionBindingResource() resource.Resource {
	return &AppfwprofileSqlinjectionBindingResource{}
}

// AppfwprofileSqlinjectionBindingResource defines the resource implementation.
type AppfwprofileSqlinjectionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileSqlinjectionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileSqlinjectionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_sqlinjection_binding"
}

func (r *AppfwprofileSqlinjectionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileSqlinjectionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_sqlinjection_binding resource")

	// appfwprofile_sqlinjection_binding := appfwprofile_sqlinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_sqlinjection_binding.Type(), &appfwprofile_sqlinjection_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_sqlinjection_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_sqlinjection_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_sqlinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_sqlinjection_binding resource")

	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_sqlinjection_binding resource")

	// Create API request body from the model
	// appfwprofile_sqlinjection_binding := appfwprofile_sqlinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_sqlinjection_binding.Type(), &appfwprofile_sqlinjection_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_sqlinjection_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_sqlinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileSqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileSqlinjectionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileSqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_sqlinjection_binding resource")

	// For appfwprofile_sqlinjection_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_sqlinjection_binding resource from state")
}

// Helper function to read appfwprofile_sqlinjection_binding data from API
func (r *AppfwprofileSqlinjectionBindingResource) readAppfwprofileSqlinjectionBindingFromApi(ctx context.Context, data *AppfwprofileSqlinjectionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_sqlinjection_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_sqlinjection_binding, got error: %s", err))
		return
	}

	appfwprofile_sqlinjection_bindingSetAttrFromGet(ctx, data, getResponseData)

}
