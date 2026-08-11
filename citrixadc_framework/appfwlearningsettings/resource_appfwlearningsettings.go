package appfwlearningsettings

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
var _ resource.Resource = &AppfwlearningsettingsResource{}
var _ resource.ResourceWithConfigure = (*AppfwlearningsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwlearningsettingsResource)(nil)

func NewAppfwlearningsettingsResource() resource.Resource {
	return &AppfwlearningsettingsResource{}
}

// AppfwlearningsettingsResource defines the resource implementation.
type AppfwlearningsettingsResource struct {
	client *service.NitroClient
}

func (r *AppfwlearningsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwlearningsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningsettings"
}

func (r *AppfwlearningsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwlearningsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwlearningsettings resource")
	// Get payload from plan
	appfwlearningsettings := appfwlearningsettingsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed resource - use UpdateUnnamedResource (NITRO only supports update/unset/get)
	err := r.client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwlearningsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwlearningsettings resource")

	// Set ID for the resource before reading state (ID is the profilename, matching SDK v2 d.SetId(profilename))
	data.Id = types.StringValue(data.Profilename.ValueString())

	// Read the updated state back
	if !r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwlearningsettings not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwlearningsettings resource")

	found := r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwlearningsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppfwlearningsettingsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset them)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwlearningsettings resource")

	// Determine which attributes were removed from config so they can be unset.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Contenttypeautodeploygraceperiod.Equal(state.Contenttypeautodeploygraceperiod) {
		if config.Contenttypeautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "contenttypeautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Contenttypeminthreshold.Equal(state.Contenttypeminthreshold) {
		if config.Contenttypeminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "contenttypeminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Contenttypepercentthreshold.Equal(state.Contenttypepercentthreshold) {
		if config.Contenttypepercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "contenttypepercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Cookieconsistencyautodeploygraceperiod.Equal(state.Cookieconsistencyautodeploygraceperiod) {
		if config.Cookieconsistencyautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookieconsistencyautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Cookieconsistencyminthreshold.Equal(state.Cookieconsistencyminthreshold) {
		if config.Cookieconsistencyminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookieconsistencyminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Cookieconsistencypercentthreshold.Equal(state.Cookieconsistencypercentthreshold) {
		if config.Cookieconsistencypercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "cookieconsistencypercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Creditcardnumberminthreshold.Equal(state.Creditcardnumberminthreshold) {
		if config.Creditcardnumberminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "creditcardnumberminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Creditcardnumberpercentthreshold.Equal(state.Creditcardnumberpercentthreshold) {
		if config.Creditcardnumberpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "creditcardnumberpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Crosssitescriptingautodeploygraceperiod.Equal(state.Crosssitescriptingautodeploygraceperiod) {
		if config.Crosssitescriptingautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "crosssitescriptingautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Crosssitescriptingminthreshold.Equal(state.Crosssitescriptingminthreshold) {
		if config.Crosssitescriptingminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "crosssitescriptingminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Crosssitescriptingpercentthreshold.Equal(state.Crosssitescriptingpercentthreshold) {
		if config.Crosssitescriptingpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "crosssitescriptingpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Csrftagautodeploygraceperiod.Equal(state.Csrftagautodeploygraceperiod) {
		if config.Csrftagautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "csrftagautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Csrftagminthreshold.Equal(state.Csrftagminthreshold) {
		if config.Csrftagminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "csrftagminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Csrftagpercentthreshold.Equal(state.Csrftagpercentthreshold) {
		if config.Csrftagpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "csrftagpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldconsistencyautodeploygraceperiod.Equal(state.Fieldconsistencyautodeploygraceperiod) {
		if config.Fieldconsistencyautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldconsistencyautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldconsistencyminthreshold.Equal(state.Fieldconsistencyminthreshold) {
		if config.Fieldconsistencyminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldconsistencyminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldconsistencypercentthreshold.Equal(state.Fieldconsistencypercentthreshold) {
		if config.Fieldconsistencypercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldconsistencypercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldformatautodeploygraceperiod.Equal(state.Fieldformatautodeploygraceperiod) {
		if config.Fieldformatautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldformatautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldformatminthreshold.Equal(state.Fieldformatminthreshold) {
		if config.Fieldformatminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldformatminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Fieldformatpercentthreshold.Equal(state.Fieldformatpercentthreshold) {
		if config.Fieldformatpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "fieldformatpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Sqlinjectionautodeploygraceperiod.Equal(state.Sqlinjectionautodeploygraceperiod) {
		if config.Sqlinjectionautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "sqlinjectionautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Sqlinjectionminthreshold.Equal(state.Sqlinjectionminthreshold) {
		if config.Sqlinjectionminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "sqlinjectionminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Sqlinjectionpercentthreshold.Equal(state.Sqlinjectionpercentthreshold) {
		if config.Sqlinjectionpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "sqlinjectionpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Starturlautodeploygraceperiod.Equal(state.Starturlautodeploygraceperiod) {
		if config.Starturlautodeploygraceperiod.IsNull() {
			attributesToUnset = append(attributesToUnset, "starturlautodeploygraceperiod")
		} else {
			hasChange = true
		}
	}
	if !data.Starturlminthreshold.Equal(state.Starturlminthreshold) {
		if config.Starturlminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "starturlminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Starturlpercentthreshold.Equal(state.Starturlpercentthreshold) {
		if config.Starturlpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "starturlpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Xmlattachmentminthreshold.Equal(state.Xmlattachmentminthreshold) {
		if config.Xmlattachmentminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "xmlattachmentminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Xmlattachmentpercentthreshold.Equal(state.Xmlattachmentpercentthreshold) {
		if config.Xmlattachmentpercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "xmlattachmentpercentthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Xmlwsiminthreshold.Equal(state.Xmlwsiminthreshold) {
		if config.Xmlwsiminthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "xmlwsiminthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Xmlwsipercentthreshold.Equal(state.Xmlwsipercentthreshold) {
		if config.Xmlwsipercentthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "xmlwsipercentthreshold")
		} else {
			hasChange = true
		}
	}

	// Create API request body from the model
	appfwlearningsettings := appfwlearningsettingsGetThePayloadFromtheConfig(ctx, &data)

	if hasChange {
		// Make API call
		// Unnamed resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwlearningsettings, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwlearningsettings resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwlearningsettings resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"profilename": data.Profilename.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Appfwlearningsettings.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset appfwlearningsettings attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwlearningsettings not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwlearningsettings resource")

	// appfwlearningsettings has no NITRO delete operation (only update/unset/get).
	// Matching SDK v2 behavior (d.SetId("")), we simply drop it from Terraform state.
	tflog.Trace(ctx, "Deleted appfwlearningsettings resource from state")
}

// Helper function to read appfwlearningsettings data from API
func (r *AppfwlearningsettingsResource) readAppfwlearningsettingsFromApi(ctx context.Context, data *AppfwlearningsettingsResourceModel, diags *diag.Diagnostics) bool {

	// The ID is the profilename (matching SDK v2 d.SetId(profilename)).
	appfwlearningsettingsName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwlearningsettings.Type(), appfwlearningsettingsName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwlearningsettings, got error: %s", err))
		return false
	}

	appfwlearningsettingsSetAttrFromGet(ctx, data, getResponseData)

	return true
}
