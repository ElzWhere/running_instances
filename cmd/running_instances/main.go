package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type allInstances struct {
	InstanceID []string
	ImageID    []string
	LaunchTime []time.Time
}

func main() {
	//initailze struct
	/*	awsi := allInstances{
			InstanceID: nil,
			ImageID:    nil,
			LaunchTime: nil,
		}
	*/
	// Authentication Using Default Config
	/*
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	*/

	s := os.Args[1]
	// Authentication Using Shared Config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(s))

	if err != nil {
		log.Fatal(err)
	}
	client := ec2.NewFromConfig(cfg)

	parms := &ec2.DescribeInstancesInput{}

	//input := &ec2.DescribeInstancesInput{}

	//Get volumes
	result, err := client.DescribeInstances(context.TODO(), parms)
	if err != nil {
		fmt.Println("Error Getting Volumes")
	}

	//	var instId []string
	//	var imageId []string
	//	var launchTime []time.Time
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("%s\t %s\t %s\t", *instance.InstanceId, *instance.ImageId, *instance.LaunchTime)
			//			launchTime = append(launchTime, *instance.LaunchTime)
			//			instId = append(instId, *instance.InstanceId)
			//			imageId = append(imageId, *instance.ImageId)
			//fmt.Println("Id: ", *instance.InstanceId)
			//fmt.Println(*instance.ImageId)
			for _, p := range instance.Tags {
				if *p.Key == "Name" {
					fmt.Println("", *p.Value)
				}
			}

		}
	}

	//	awsi.InstanceID = instId
	//	awsi.LaunchTime = launchTime
	//	awsi.ImageID = imageId
	//	fmt.Println(awsi.ImageID, awsi.InstanceID, awsi.LaunchTime)
	//fmt.Println(awsi.InstanceID)
	//fmt.Println(awsi.LaunchTime)

}
