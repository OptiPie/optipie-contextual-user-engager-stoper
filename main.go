package main

import (
	"context"
	"fmt"
	"github.com/OptiPie/optipie-contextual-user-engager-stoper/prepare"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func Handler(ctx context.Context) error {
	fmt.Println("stopping the ec2 instance of twitter-contextual-user-engager")

	awsCfg, err := prepare.AwsConfig(ctx)
	if err != nil {
		log.Fatalf("error on initializing aws config: %v", err)
	}

	ec2Client := prepare.Ec2(awsCfg)
	_, err = ec2Client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: []string{"i-023509645a6761561"},
	})

	if err != nil {
		log.Fatalf("error on ec2 stopInstances: %v", err)
	}

	return nil
}

func main() {
	if _, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); exists {
		lambda.Start(Handler)
	} else {
		err := Handler(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}
