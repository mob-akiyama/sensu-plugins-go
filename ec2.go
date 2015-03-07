package main

import (
  "fmt"
  "os"
  "time"
  "flag"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

var aws_access_key = flag.String("a", nil, "AWS API KEY")
var aws_secret_key = flag.String("k", nil, "AWS API SECRET")
var fetch_age      = flag.Int("f", nil, "How long ago to fetch metrics for")
flag.Parse()

var defaultRegion  = os.Getenv("EC2_REGION")
var instanceId     = os.Getenv("EC2_INSTANCE_ID")
var creds          = aws.Creds(aws_access_key, aws_secret_key, "")
var cli            = cloudwatch.New(creds, defaultRegion, nil)
var dimensionParam = &cloudwatch.Dimension{
    Name:  aws.String("InstanceId"),
    Value: aws.String(instanceId),
}

func main() {
    fmt.Print(getCPUUtilization(), "\n")
    fmt.Print(getStatusCheckFailed(), "\n")
    fmt.Print(getStatusCheckFailed_Instance(), "\n")
    fmt.Print(getStatusCheckFailed_System(), "\n")
}

func getCPUUtilization() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(-600 * time.Second),
        EndTime:    time.Now(),
        MetricName: aws.String("CPUUtilization"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cli.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(-600 * time.Second),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cli.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed_Instance() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(-600 * time.Second),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed_Instance"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cli.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed_System() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(-600 * time.Second),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed_System"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cli.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}
