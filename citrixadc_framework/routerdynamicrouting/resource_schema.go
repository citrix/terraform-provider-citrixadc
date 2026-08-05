package routerdynamicrouting

import (
	"context"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/router"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RouterdynamicroutingResourceModel describes the resource data model.
//
// routerdynamicrouting is an action-only resource. NITRO exposes only the
// "apply" action (no GET/DELETE for the applied configuration), so the model
// mirrors the SDK v2 contract exactly: a single list of command lines that are
// joined and pushed to the appliance via ?action=apply.
type RouterdynamicroutingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Commandlines types.List   `tfsdk:"commandlines"`
}

func (r *RouterdynamicroutingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the routerdynamicrouting resource.",
			},
			// SDK v2: TypeList of strings, Optional + Computed + ForceNew.
			// Optional+Computed+ForceNew maps to RequiresReplaceIfConfigured();
			// UseStateForUnknown is added before it because the attribute is
			// Computed (avoids spurious known-after-apply churn).
			"commandlines": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
					listplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The commands to be applied as dynamic routing configuration. Each element is a single command line; the lines are joined and pushed to the appliance.",
			},
		},
	}
}

// routerdynamicroutingGetThePayloadFromthePlan builds the NITRO payload for the
// "apply" action. The list of command lines is joined with newlines into the
// single Commandstring field, mirroring the SDK v2 implementation.
func routerdynamicroutingGetThePayloadFromthePlan(ctx context.Context, data *RouterdynamicroutingResourceModel) router.Routerdynamicrouting {
	tflog.Debug(ctx, "In routerdynamicroutingGetThePayloadFromthePlan Function")

	var lines []string
	if !data.Commandlines.IsNull() && !data.Commandlines.IsUnknown() {
		data.Commandlines.ElementsAs(ctx, &lines, false)
	}

	cmdString := strings.Join(lines, "\n")

	return router.Routerdynamicrouting{
		Commandstring: cmdString,
	}
}
