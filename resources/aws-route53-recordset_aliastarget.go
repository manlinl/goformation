package resources

// AWS::Route53::RecordSet.AliasTarget AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type AWSRoute53RecordSetAliasTarget struct {

	// DNSName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname
	DNSName string `json:"DNSName"`

	// EvaluateTargetHealth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth
	EvaluateTargetHealth bool `json:"EvaluateTargetHealth"`

	// HostedZoneId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid
	HostedZoneId string `json:"HostedZoneId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetAliasTarget) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSet.AliasTarget"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetAliasTarget) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}