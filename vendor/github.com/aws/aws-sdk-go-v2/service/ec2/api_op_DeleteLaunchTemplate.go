// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplateRequest
type DeleteLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateId *string `type:"string"`

	// The name of the launch template. You must specify either the launch template
	// ID or launch template name in the request.
	LaunchTemplateName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s DeleteLaunchTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLaunchTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLaunchTemplateInput"}
	if s.LaunchTemplateName != nil && len(*s.LaunchTemplateName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("LaunchTemplateName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplateResult
type DeleteLaunchTemplateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the launch template.
	LaunchTemplate *LaunchTemplate `locationName:"launchTemplate" type:"structure"`
}

// String returns the string representation
func (s DeleteLaunchTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLaunchTemplate = "DeleteLaunchTemplate"

// DeleteLaunchTemplateRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes a launch template. Deleting a launch template deletes all of its
// versions.
//
//    // Example sending a request using DeleteLaunchTemplateRequest.
//    req := client.DeleteLaunchTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteLaunchTemplate
func (c *Client) DeleteLaunchTemplateRequest(input *DeleteLaunchTemplateInput) DeleteLaunchTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteLaunchTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLaunchTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteLaunchTemplateOutput{})
	return DeleteLaunchTemplateRequest{Request: req, Input: input, Copy: c.DeleteLaunchTemplateRequest}
}

// DeleteLaunchTemplateRequest is the request type for the
// DeleteLaunchTemplate API operation.
type DeleteLaunchTemplateRequest struct {
	*aws.Request
	Input *DeleteLaunchTemplateInput
	Copy  func(*DeleteLaunchTemplateInput) DeleteLaunchTemplateRequest
}

// Send marshals and sends the DeleteLaunchTemplate API request.
func (r DeleteLaunchTemplateRequest) Send(ctx context.Context) (*DeleteLaunchTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLaunchTemplateResponse{
		DeleteLaunchTemplateOutput: r.Request.Data.(*DeleteLaunchTemplateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLaunchTemplateResponse is the response type for the
// DeleteLaunchTemplate API operation.
type DeleteLaunchTemplateResponse struct {
	*DeleteLaunchTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLaunchTemplate request.
func (r *DeleteLaunchTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
