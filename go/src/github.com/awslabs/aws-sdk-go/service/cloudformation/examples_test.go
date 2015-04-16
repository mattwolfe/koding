package cloudformation_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/cloudformation"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudFormation_CancelUpdateStack() {
	svc := cloudformation.New(nil)

	params := &cloudformation.CancelUpdateStackInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.CancelUpdateStack(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_CreateStack() {
	svc := cloudformation.New(nil)

	params := &cloudformation.CreateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		DisableRollback: aws.Boolean(true),
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		OnFailure: aws.String("OnFailure"),
		Parameters: []*cloudformation.Parameter{
			&cloudformation.Parameter{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Boolean(true),
			},
			// More values...
		},
		StackPolicyBody: aws.String("StackPolicyBody"),
		StackPolicyURL:  aws.String("StackPolicyURL"),
		Tags: []*cloudformation.Tag{
			&cloudformation.Tag{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:     aws.String("TemplateBody"),
		TemplateURL:      aws.String("TemplateURL"),
		TimeoutInMinutes: aws.Long(1),
	}
	resp, err := svc.CreateStack(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_DeleteStack() {
	svc := cloudformation.New(nil)

	params := &cloudformation.DeleteStackInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.DeleteStack(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_DescribeStackEvents() {
	svc := cloudformation.New(nil)

	params := &cloudformation.DescribeStackEventsInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String("StackName"),
	}
	resp, err := svc.DescribeStackEvents(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_DescribeStackResource() {
	svc := cloudformation.New(nil)

	params := &cloudformation.DescribeStackResourceInput{
		LogicalResourceID: aws.String("LogicalResourceId"), // Required
		StackName:         aws.String("StackName"),         // Required
	}
	resp, err := svc.DescribeStackResource(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_DescribeStackResources() {
	svc := cloudformation.New(nil)

	params := &cloudformation.DescribeStackResourcesInput{
		LogicalResourceID:  aws.String("LogicalResourceId"),
		PhysicalResourceID: aws.String("PhysicalResourceId"),
		StackName:          aws.String("StackName"),
	}
	resp, err := svc.DescribeStackResources(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_DescribeStacks() {
	svc := cloudformation.New(nil)

	params := &cloudformation.DescribeStacksInput{
		NextToken: aws.String("NextToken"),
		StackName: aws.String("StackName"),
	}
	resp, err := svc.DescribeStacks(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_EstimateTemplateCost() {
	svc := cloudformation.New(nil)

	params := &cloudformation.EstimateTemplateCostInput{
		Parameters: []*cloudformation.Parameter{
			&cloudformation.Parameter{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Boolean(true),
			},
			// More values...
		},
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.EstimateTemplateCost(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_GetStackPolicy() {
	svc := cloudformation.New(nil)

	params := &cloudformation.GetStackPolicyInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.GetStackPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_GetTemplate() {
	svc := cloudformation.New(nil)

	params := &cloudformation.GetTemplateInput{
		StackName: aws.String("StackName"), // Required
	}
	resp, err := svc.GetTemplate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_GetTemplateSummary() {
	svc := cloudformation.New(nil)

	params := &cloudformation.GetTemplateSummaryInput{
		StackName:    aws.String("StackNameOrId"),
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.GetTemplateSummary(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_ListStackResources() {
	svc := cloudformation.New(nil)

	params := &cloudformation.ListStackResourcesInput{
		StackName: aws.String("StackName"), // Required
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListStackResources(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_ListStacks() {
	svc := cloudformation.New(nil)

	params := &cloudformation.ListStacksInput{
		NextToken: aws.String("NextToken"),
		StackStatusFilter: []*string{
			aws.String("StackStatus"), // Required
			// More values...
		},
	}
	resp, err := svc.ListStacks(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_SetStackPolicy() {
	svc := cloudformation.New(nil)

	params := &cloudformation.SetStackPolicyInput{
		StackName:       aws.String("StackName"), // Required
		StackPolicyBody: aws.String("StackPolicyBody"),
		StackPolicyURL:  aws.String("StackPolicyURL"),
	}
	resp, err := svc.SetStackPolicy(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_SignalResource() {
	svc := cloudformation.New(nil)

	params := &cloudformation.SignalResourceInput{
		LogicalResourceID: aws.String("LogicalResourceId"),      // Required
		StackName:         aws.String("StackNameOrId"),          // Required
		Status:            aws.String("ResourceSignalStatus"),   // Required
		UniqueID:          aws.String("ResourceSignalUniqueId"), // Required
	}
	resp, err := svc.SignalResource(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_UpdateStack() {
	svc := cloudformation.New(nil)

	params := &cloudformation.UpdateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		Parameters: []*cloudformation.Parameter{
			&cloudformation.Parameter{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Boolean(true),
			},
			// More values...
		},
		StackPolicyBody:             aws.String("StackPolicyBody"),
		StackPolicyDuringUpdateBody: aws.String("StackPolicyDuringUpdateBody"),
		StackPolicyDuringUpdateURL:  aws.String("StackPolicyDuringUpdateURL"),
		StackPolicyURL:              aws.String("StackPolicyURL"),
		TemplateBody:                aws.String("TemplateBody"),
		TemplateURL:                 aws.String("TemplateURL"),
		UsePreviousTemplate:         aws.Boolean(true),
	}
	resp, err := svc.UpdateStack(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleCloudFormation_ValidateTemplate() {
	svc := cloudformation.New(nil)

	params := &cloudformation.ValidateTemplateInput{
		TemplateBody: aws.String("TemplateBody"),
		TemplateURL:  aws.String("TemplateURL"),
	}
	resp, err := svc.ValidateTemplate(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}