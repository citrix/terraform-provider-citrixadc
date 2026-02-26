package appfwprofile_xmlsqlinjection_binding

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
var _ resource.Resource = &AppfwprofileXmlsqlinjectionBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileXmlsqlinjectionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileXmlsqlinjectionBindingResource)(nil)

func NewAppfwprofileXmlsqlinjectionBindingResource() resource.Resource {
	return &AppfwprofileXmlsqlinjectionBindingResource{}
}

// AppfwprofileXmlsqlinjectionBindingResource defines the resource implementation.
type AppfwprofileXmlsqlinjectionBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlsqlinjection_binding"
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileXmlsqlinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_xmlsqlinjection_binding resource")

	// appfwprofile_xmlsqlinjection_binding := appfwprofile_xmlsqlinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), &appfwprofile_xmlsqlinjection_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_xmlsqlinjection_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_xmlsqlinjection_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_xmlsqlinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlsqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileXmlsqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_xmlsqlinjection_binding resource")

	r.readAppfwprofileXmlsqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileXmlsqlinjectionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_xmlsqlinjection_binding resource")

	// Create API request body from the model
	// appfwprofile_xmlsqlinjection_binding := appfwprofile_xmlsqlinjection_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), &appfwprofile_xmlsqlinjection_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_xmlsqlinjection_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_xmlsqlinjection_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlsqlinjectionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlsqlinjectionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileXmlsqlinjectionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_xmlsqlinjection_binding resource")

	// For appfwprofile_xmlsqlinjection_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_xmlsqlinjection_binding resource from state")
}

// Helper function to read appfwprofile_xmlsqlinjection_binding data from API
func (r *AppfwprofileXmlsqlinjectionBindingResource) readAppfwprofileXmlsqlinjectionBindingFromApi(ctx context.Context, data *AppfwprofileXmlsqlinjectionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlsqlinjection_binding, got error: %s", err))
		return
	}

	appfwprofile_xmlsqlinjection_bindingSetAttrFromGet(ctx, data, getResponseData)

}
