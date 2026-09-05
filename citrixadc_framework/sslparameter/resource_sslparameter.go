package sslparameter

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
var _ resource.Resource = &SslparameterResource{}
var _ resource.ResourceWithConfigure = (*SslparameterResource)(nil)
var _ resource.ResourceWithImportState = (*SslparameterResource)(nil)

func NewSslparameterResource() resource.Resource {
	return &SslparameterResource{}
}

// SslparameterResource defines the resource implementation.
type SslparameterResource struct {
	client *service.NitroClient
}

func (r *SslparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslparameter"
}

func (r *SslparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslparameter resource")

	sslparameter := sslparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslparameter.Type(), &sslparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslparameter, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue("sslparameter-config")

	tflog.Trace(ctx, "Created sslparameter resource")

	// Read the updated state back
	r.readSslparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslparameter resource")

	r.readSslparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslparameter resource")

	// Determine which attributes changed and which were removed from config (unset)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Crlmemorysizemb.Equal(state.Crlmemorysizemb) {
		if config.Crlmemorysizemb.IsNull() {
			attributesToUnset = append(attributesToUnset, "crlmemorysizemb")
		} else {
			hasChange = true
		}
	}
	if !data.Dropreqwithnohostheader.Equal(state.Dropreqwithnohostheader) {
		if config.Dropreqwithnohostheader.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropreqwithnohostheader")
		} else {
			hasChange = true
		}
	}
	if !data.Encrypttriggerpktcount.Equal(state.Encrypttriggerpktcount) {
		if config.Encrypttriggerpktcount.IsNull() {
			attributesToUnset = append(attributesToUnset, "encrypttriggerpktcount")
		} else {
			hasChange = true
		}
	}
	if !data.Insertcertspace.Equal(state.Insertcertspace) {
		if config.Insertcertspace.IsNull() {
			attributesToUnset = append(attributesToUnset, "insertcertspace")
		} else {
			hasChange = true
		}
	}
	if !data.Insertionencoding.Equal(state.Insertionencoding) {
		if config.Insertionencoding.IsNull() {
			attributesToUnset = append(attributesToUnset, "insertionencoding")
		} else {
			hasChange = true
		}
	}
	if !data.Ndcppcompliancecertcheck.Equal(state.Ndcppcompliancecertcheck) {
		if config.Ndcppcompliancecertcheck.IsNull() {
			attributesToUnset = append(attributesToUnset, "ndcppcompliancecertcheck")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspcachesize.Equal(state.Ocspcachesize) {
		if config.Ocspcachesize.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspcachesize")
		} else {
			hasChange = true
		}
	}
	if !data.Pushenctriggertimeout.Equal(state.Pushenctriggertimeout) {
		if config.Pushenctriggertimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "pushenctriggertimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Quantumsize.Equal(state.Quantumsize) {
		if config.Quantumsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "quantumsize")
		} else {
			hasChange = true
		}
	}
	if !data.Sendclosenotify.Equal(state.Sendclosenotify) {
		if config.Sendclosenotify.IsNull() {
			attributesToUnset = append(attributesToUnset, "sendclosenotify")
		} else {
			hasChange = true
		}
	}
	if !data.Snihttphostmatch.Equal(state.Snihttphostmatch) {
		if config.Snihttphostmatch.IsNull() {
			attributesToUnset = append(attributesToUnset, "snihttphostmatch")
		} else {
			hasChange = true
		}
	}
	if !data.Sslierrorcache.Equal(state.Sslierrorcache) {
		if config.Sslierrorcache.IsNull() {
			attributesToUnset = append(attributesToUnset, "sslierrorcache")
		} else {
			hasChange = true
		}
	}
	if !data.Ssltriggertimeout.Equal(state.Ssltriggertimeout) {
		if config.Ssltriggertimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "ssltriggertimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Strictcachecks.Equal(state.Strictcachecks) {
		if config.Strictcachecks.IsNull() {
			attributesToUnset = append(attributesToUnset, "strictcachecks")
		} else {
			hasChange = true
		}
	}
	if !data.Undefactioncontrol.Equal(state.Undefactioncontrol) {
		if config.Undefactioncontrol.IsNull() {
			attributesToUnset = append(attributesToUnset, "undefactioncontrol")
		} else {
			hasChange = true
		}
	}
	if !data.Undefactiondata.Equal(state.Undefactiondata) {
		if config.Undefactiondata.IsNull() {
			attributesToUnset = append(attributesToUnset, "undefactiondata")
		} else {
			hasChange = true
		}
	}
	// Attributes not wired for unset above still trigger a normal update on change.
	if !data.Cryptodevdisablelimit.Equal(state.Cryptodevdisablelimit) ||
		!data.Defaultprofile.Equal(state.Defaultprofile) ||
		!data.Denysslreneg.Equal(state.Denysslreneg) ||
		!data.Heterogeneoussslhw.Equal(state.Heterogeneoussslhw) ||
		!data.Hybridfipsmode.Equal(state.Hybridfipsmode) ||
		!data.Operationqueuelimit.Equal(state.Operationqueuelimit) ||
		!data.Pushflag.Equal(state.Pushflag) ||
		!data.Sigdigesttype.Equal(state.Sigdigesttype) ||
		!data.Softwarecryptothreshold.Equal(state.Softwarecryptothreshold) ||
		!data.Sslimaxerrorcachemem.Equal(state.Sslimaxerrorcachemem) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		sslparameter := sslparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslparameter.Type(), &sslparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts to defaults
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Sslparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readSslparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslparameter resource")

	// For sslparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted sslparameter resource from state")
}

// Helper function to read sslparameter data from API
func (r *SslparameterResource) readSslparameterFromApi(ctx context.Context, data *SslparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Sslparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslparameter, got error: %s", err))
		return
	}

	sslparameterSetAttrFromGet(ctx, data, getResponseData)

}
