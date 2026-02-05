package ntpserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NtpserverResource{}
var _ resource.ResourceWithConfigure = (*NtpserverResource)(nil)
var _ resource.ResourceWithImportState = (*NtpserverResource)(nil)

func NewNtpserverResource() resource.Resource {
	return &NtpserverResource{}
}

// NtpserverResource defines the resource implementation.
type NtpserverResource struct {
	client *service.NitroClient
}

func (r *NtpserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NtpserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntpserver"
}

func (r *NtpserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NtpserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NtpserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ntpserver resource")

	ntpserver := ntpserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	_, err := r.client.AddResource(service.Ntpserver.Type(), "", &ntpserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ntpserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ntpserver resource")

	// Read the updated state back
	r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NtpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ntpserver resource")

	r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NtpserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ntpserver resource")

	// Create API request body from the model
	ntpserver := ntpserverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call - update the resource
	err := r.client.UpdateUnnamedResource(service.Ntpserver.Type(), &ntpserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ntpserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated ntpserver resource")

	// Read the updated state back
	r.readNtpserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NtpserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ntpserver resource")

	// Build the identifier for deletion
	var identifier string
	if !data.Serverip.IsNull() && data.Serverip.ValueString() != "" {
		identifier = data.Serverip.ValueString()
	} else if !data.Servername.IsNull() && data.Servername.ValueString() != "" {
		identifier = data.Servername.ValueString()
	}

	if identifier != "" {
		err := r.client.DeleteResource(service.Ntpserver.Type(), identifier)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ntpserver, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Deleted ntpserver resource")
}

// Helper function to read ntpserver data from API
func (r *NtpserverResource) readNtpserverFromApi(ctx context.Context, data *NtpserverResourceModel, diags *diag.Diagnostics) {
	// Build the identifier for reading
	var identifier string
	if !data.Serverip.IsNull() && data.Serverip.ValueString() != "" {
		identifier = data.Serverip.ValueString()
	} else if !data.Servername.IsNull() && data.Servername.ValueString() != "" {
		identifier = data.Servername.ValueString()
	}

	if identifier == "" {
		diags.AddError("Configuration Error", "At least one of serverip or servername must be specified")
		return
	}

	getResponseData, err := r.client.FindResource(service.Ntpserver.Type(), identifier)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ntpserver, got error: %s", err))
		return
	}

	ntpserverSetAttrFromGet(ctx, data, getResponseData)
}
