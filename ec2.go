package main

import (
  "fmt"
  "time"
  "flag"
  "net/http"
  "io/ioutil"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

var aws_access_key = flag.String("a", "", "AWS API KEY")
var aws_secret_key = flag.String("k", "", "AWS API SECRET")
var f      = flag.Int("f", 300, "How long ago to fetch metrics for")
var fetch_age = time.Duration(*f) * time.Second

func getInstanceID() string {
    url := "http://169.254.169.254/latest/meta-data/instance-id"

    resp, _ := http.Get(url)
    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
    return string(byteArray)
}

func getRegion() string {
    url := "http://169.254.169.254/latest/meta-data/placement/availability-zone"

    resp, _ := http.Get(url)
    defer resp.Body.Close()

    byteArray, _ := ioutil.ReadAll(resp.Body)
    return string(byteArray[0:len(byteArray)-1])
}

var aws_region  = getRegion()
var instance_id = getInstanceID()

var creds = aws.Creds(*aws_access_key, *aws_secret_key, "")
var cw    = cloudwatch.New(creds, aws_region, nil)
var dimensionParam = &cloudwatch.Dimension{
    Name:  aws.String("InstanceId"),
    Value: aws.String(instance_id),
}

func main() {
    flag.Parse()
    fmt.Print(getCPUUtilization(), "\n")
    fmt.Print(getStatusCheckFailed(), "\n")
    fmt.Print(getStatusCheckFailed_Instance(), "\n")
    fmt.Print(getStatusCheckFailed_System(), "\n")
}

func getCPUUtilization() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(fetch_age),
        EndTime:    time.Now(),
        MetricName: aws.String("CPUUtilization"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cw.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(fetch_age),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cw.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed_Instance() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(fetch_age),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed_Instance"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cw.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}

func getStatusCheckFailed_System() float64 {
    mt := &cloudwatch.GetMetricStatisticsInput{
        Dimensions: []cloudwatch.Dimension{*dimensionParam},
        StartTime:  time.Now().Add(fetch_age),
        EndTime:    time.Now(),
        MetricName: aws.String("StatusCheckFailed_System"),
        Namespace:  aws.String("AWS/EC2"),
        Period:     aws.Integer(300),
        Statistics: []string{"Average"},
    }
    resp, _ := cw.GetMetricStatistics(mt)
    return *resp.Datapoints[0].Average
}
