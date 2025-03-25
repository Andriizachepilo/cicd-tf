package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type WorkDirValidator struct{}

func (v WorkDirValidator) Description(_ context.Context) string {
	return "Ensures that 'working_directory' is set when either 'build' or 'test' is provided."
}

func (v WorkDirValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v WorkDirValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.Config.Raw.IsNull() {
		return
	}

	var build, test string
	req.Config.GetAttribute(ctx, path.Root("build"), &build)
	req.Config.GetAttribute(ctx, path.Root("test"), &test)
	value := req.ConfigValue.ValueString()

	if (build != "" || test != "") && value == "" {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic(
			"Invalid Configuration",
			"'working_directory' must be set when either 'build' or 'test' is provided.",
		))
	} 
}
