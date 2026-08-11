package auditnslogparams

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
var _ resource.Resource = &AuditnslogparamsResource{}
var _ resource.ResourceWithConfigure = (*AuditnslogparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AuditnslogparamsResource)(nil)

func NewAuditnslogparamsResource() resource.Resource {
	return &AuditnslogparamsResource{}
}

// AuditnslogparamsResource defines the resource implementation.
type AuditnslogparamsResource struct {
	client *service.NitroClient
}

func (r *AuditnslogparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditnslogparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogparams"
}

func (r *AuditnslogparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditnslogparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditnslogparams resource")

	// Create API request body from the model
	auditnslogparams := auditnslogparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed (singleton) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditnslogparams, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditnslogparams-config")

	tflog.Trace(ctx, "Created auditnslogparams resource")

	// Read the updated state back
	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditnslogparams resource")

	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuditnslogparamsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read raw config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating auditnslogparams resource")

	// Determine whether an update is needed and which attributes must be unset
	// (removed from config -> revert to NITRO default).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acl.Equal(state.Acl) {
		if config.Acl.IsNull() {
			attributesToUnset = append(attributesToUnset, "acl")
		} else {
			hasChange = true
		}
	}
	if !data.Alg.Equal(state.Alg) {
		if config.Alg.IsNull() {
			attributesToUnset = append(attributesToUnset, "alg")
		} else {
			hasChange = true
		}
	}
	if !data.Appflowexport.Equal(state.Appflowexport) {
		if config.Appflowexport.IsNull() {
			attributesToUnset = append(attributesToUnset, "appflowexport")
		} else {
			hasChange = true
		}
	}
	if !data.Contentinspectionlog.Equal(state.Contentinspectionlog) {
		if config.Contentinspectionlog.IsNull() {
			attributesToUnset = append(attributesToUnset, "contentinspectionlog")
		} else {
			hasChange = true
		}
	}
	if !data.Dateformat.Equal(state.Dateformat) {
		if config.Dateformat.IsNull() {
			attributesToUnset = append(attributesToUnset, "dateformat")
		} else {
			hasChange = true
		}
	}
	if !data.Logfacility.Equal(state.Logfacility) {
		if config.Logfacility.IsNull() {
			attributesToUnset = append(attributesToUnset, "logfacility")
		} else {
			hasChange = true
		}
	}
	if !data.Lsn.Equal(state.Lsn) {
		if config.Lsn.IsNull() {
			attributesToUnset = append(attributesToUnset, "lsn")
		} else {
			hasChange = true
		}
	}
	if !data.Protocolviolations.Equal(state.Protocolviolations) {
		if config.Protocolviolations.IsNull() {
			attributesToUnset = append(attributesToUnset, "protocolviolations")
		} else {
			hasChange = true
		}
	}
	if !data.Sslinterception.Equal(state.Sslinterception) {
		if config.Sslinterception.IsNull() {
			attributesToUnset = append(attributesToUnset, "sslinterception")
		} else {
			hasChange = true
		}
	}
	if !data.Subscriberlog.Equal(state.Subscriberlog) {
		if config.Subscriberlog.IsNull() {
			attributesToUnset = append(attributesToUnset, "subscriberlog")
		} else {
			hasChange = true
		}
	}
	if !data.Tcp.Equal(state.Tcp) {
		if config.Tcp.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcp")
		} else {
			hasChange = true
		}
	}
	if !data.Timezone.Equal(state.Timezone) {
		if config.Timezone.IsNull() {
			attributesToUnset = append(attributesToUnset, "timezone")
		} else {
			hasChange = true
		}
	}
	if !data.Userdefinedauditlog.Equal(state.Userdefinedauditlog) {
		if config.Userdefinedauditlog.IsNull() {
			attributesToUnset = append(attributesToUnset, "userdefinedauditlog")
		} else {
			hasChange = true
		}
	}
	// Attributes without an unset-on-remove default still drive a normal update.
	if !data.Loglevel.Equal(state.Loglevel) {
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		hasChange = true
	}
	if !data.Urlfiltering.Equal(state.Urlfiltering) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		auditnslogparams := auditnslogparamsGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed (singleton) resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Auditnslogparams.Type(), &auditnslogparams)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditnslogparams, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated auditnslogparams resource")
	} else {
		tflog.Debug(ctx, "No changes detected for auditnslogparams resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Singleton resource -> empty identifying payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Auditnslogparams.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset auditnslogparams attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readAuditnslogparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditnslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditnslogparams resource")

	// For auditnslogparams, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditnslogparams resource from state")
}

// Helper function to read auditnslogparams data from API
func (r *AuditnslogparamsResource) readAuditnslogparamsFromApi(ctx context.Context, data *AuditnslogparamsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditnslogparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogparams, got error: %s", err))
		return
	}

	auditnslogparamsSetAttrFromGet(ctx, data, getResponseData)

}
