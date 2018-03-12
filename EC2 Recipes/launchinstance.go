package main

//Create an Amazon EC2 instance with tags and log into the instance

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	// Allows you to define the profile if not profile is given the default profile will be used

	profile := flag.String("p", "", "Default profile will be used")
	region := flag.String("r", "us-east-1", "Region defaults to us-east-2")
	imageid := flag.String("i", "ami-97785bed", "Image ID to launch")

	flag.Parse()

	// Create a new AWS Session with Options based on if a profile was given

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(*region)},
		Profile: *profile,
	}))

	//Get a handle on EC2 Service
	svc := ec2.New(sess)

	launchResult, err := svc.RunInstances(&ec2.RunInstancesInput{

		ImageId:      aws.String(*imageid),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		log.Println("Could not create instance", err)
		return
	}

	log.Println("Created instance", *launchResult.Instances[0].InstanceId)

}
