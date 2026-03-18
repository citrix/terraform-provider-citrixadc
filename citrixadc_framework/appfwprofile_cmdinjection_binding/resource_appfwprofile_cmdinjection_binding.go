package appfwprofile_cmdinjection_binding

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
var _ resource.Resource = &AppfwprofileCmdinjectionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCmdinjectionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCmdinjectionBindingResource)(nil)

func NewAppfwprofileCmdinjectionBindingResource() resource.Resource {
	return &AppfwprofileCmdinjectionBindingResource{}
}

// AppfwprofileCmdinjectionBindingResource defines the resource implementation.
type AppfwprofileCmdinjectionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCmdinjectionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCmdinjectionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_cmdinjection_binding"
}

func (r *AppfwprofileCmdinjectionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCmdinjectionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_cmdinjection_binding resource")

	// appfwprofile_cmdinjection_binding := appfwprofile_cmdinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_cmdinjection_binding.Type(), &appfwprofile_cmdinjection_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_cmdinjection_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_cmdinjection_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_cmdinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_cmdinjection_binding resource")

	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_cmdinjection_binding resource")

	// Create API request body from the model
	// appfwprofile_cmdinjection_binding := appfwprofile_cmdinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_cmdinjection_binding.Type(), &appfwprofile_cmdinjection_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_cmdinjection_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_cmdinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileCmdinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCmdinjectionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCmdinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_cmdinjection_binding resource")

	// For appfwprofile_cmdinjection_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_cmdinjection_binding resource from state")
}

// Helper function to read appfwprofile_cmdinjection_binding data from API
func (r *AppfwprofileCmdinjectionBindingResource) readAppfwprofileCmdinjectionBindingFromApi(ctx context.Context, data *AppfwprofileCmdinjectionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_cmdinjection_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_cmdinjection_binding, got error: %s", err))
		return
	}

	appfwprofile_cmdinjection_bindingSetAttrFromGet(ctx, data, getResponseData)

}
