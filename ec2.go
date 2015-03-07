package main

import (
  "os"
  "fmt"
  "time"
  "flag"
  "net/http"
  "io/ioutil"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

func main() {
    var aws_access_key = flag.String("a", "", "AWS API KEY")
    var aws_secret_key = flag.String("k", "", "AWS API SECRET")
    var fetch_age      = flag.Int("f", 300, "How long ago to fetch metrics for")
    flag.Parse()

    var hostname, _ = os.Hostname()
    var metrics = [...]string{
        "CPUUtilization",
        "StatusCheckFailed",
        "StatusCheckFailed_Instance",
        "StatusCheckFailed_System",
    }

    var creds = aws.Creds(*aws_access_key, *aws_secret_key, "")
    var cw    = cloudwatch.New(creds, getRegion(), nil)

    var dimensionParam = &cloudwatch.Dimension{
        Name:  aws.String("InstanceId"),
        Value: aws.String(getInstanceID()),
    }

    var et = time.Now().Add(time.Duration(-*fetch_age) * time.Second)
    var st = et.Add(time.Duration(-*fetch_age) * time.Second)

    for index := range metrics {
        mt := &cloudwatch.GetMetricStatisticsInput{
            Dimensions: []cloudwatch.Dimension{*dimensionParam},
            StartTime:  st,
            EndTime:    et,
            MetricName: aws.String(metrics[index]),
            Namespace:  aws.String("AWS/EC2"),
            Period:     aws.Integer(300),
            Statistics: []string{"Average"},
        }
        resp, _ := cw.GetMetricStatistics(mt)
        fmt.Print(*resp.Datapoints[0].Average, "\n")
        fmt.Print("aws.ec2.", hostname, ".", metrics[index], "average\n")
    }
}

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
