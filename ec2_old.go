package main

import (
  "fmt"
  "os"
  "time"
  "flag"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

func main() {
    var accessKey = flag.String("a", "", "AWS API KEY")
    var secretKey = flag.String("k", "", "AWS API SECRET")
    
    var defaultRegion  = os.Getenv("EC2_REGION")
    var instanceId     = os.Getenv("EC2_INSTANCE_ID")
    var creds          = aws.Creds(*accessKey, *secretKey, "")
    var cli            = cloudwatch.New(creds, defaultRegion, nil)
    var dimensionParam = &cloudwatch.Dimension{
        Name:  aws.String("InstanceId"),
        Value: aws.String(instanceId),
    }

    flag.Parse()
    fmt.Print(getCPUUtilization(), "\n")
    fmt.Print(getStatusCheckFailed(), "\n")
    fmt.Print(getStatusCheckFailed_Instance(), "\n")
    fmt.Print(getStatusCheckFailed_System(), "\n")
}

func getCPUUtilization(cloudwatch cli) float64 {
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

func getStatusCheckFailed(cloudwatch cli) float64 {
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

func getStatusCheckFailed_Instance(cloudwatch cli) float64 {
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

func getStatusCheckFailed_System(cloudwatch cli) float64 {
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
