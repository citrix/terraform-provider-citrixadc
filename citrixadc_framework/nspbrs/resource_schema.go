package nspbrs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NspbrsResourceModel describes the resource data model.
//
// nspbrs ("Configuration for Policy based routing") is a synthetic, action-only
// resource. NITRO exposes only the apply / clear / renumber POST actions for
// nspbrs and no GET / add / update / delete verbs. To preserve backward
// compatibility with the SDK v2 implementation
// (citrixadc/resource_citrixadc_nspbrs.go) the model carries the same single
// user-facing attribute:
//   - action : which one-shot action to perform (apply | clear | renumber)
type NspbrsResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
}

func (r *NspbrsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nspbrs resource.",
			},
			// SDK v2 parity: action was TypeString, Required, Computed:false,
			// ForceNew:true. ForceNew maps to RequiresReplace(); no Computed and
			// no Default (SDK v2 had none).
			"action": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The action to perform on the policy based routing configuration. Supported values of action are `apply`, `clear` or `renumber`.",
			},
		},
	}
}
