package prepare

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func AwsConfig(ctx context.Context) (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-central-1"))
	if err != nil {
		return aws.Config{}, fmt.Errorf("unable to load SDK config, %v", err)
	}
	return awsCfg, nil
}

func Ec2(awsCfg aws.Config) *ec2.Client {
	return ec2.NewFromConfig(awsCfg)
}
