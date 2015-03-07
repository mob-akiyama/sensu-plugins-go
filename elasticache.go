package main

import (
    "fmt"
    "time"
    "flag"
    "github.com/awslabs/aws-sdk-go/aws"
    "github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

func main() {
    var cluster_name       = flag.String("n", "", "RDS instance name")
    var aws_region     = flag.String("r", "us-east-1", "AWS Region (such as eu-west-1)")
    var aws_access_key = flag.String("a", "", "AWS API KEY")
    var aws_secret_key = flag.String("k", "", "AWS API SECRET")
    var fetch_age      = flag.Int("f", 300, "How long ago to fetch metrics for")
    flag.Parse()

    var creds = aws.Creds(*aws_access_key, *aws_secret_key, "")
    var cw    = cloudwatch.New(creds, *aws_region, nil)

    var metrics = [...]string{
        "CPUUtilization",
        "SwapUsage",
        "FreeableMemory",
        "NetworkBytesIn",
        "NetworkBytesOut",
        "CurrConnections",
        "Evictions",
        "Reclaimed",
        "NewConnections",
        "BytesUsedForCache",
        "CacheHits",
        "CacheMisses",
        "GetTypeCmds",
        "SetTypeCmds",
        "StringBasedCmds",
        "HashBasedCmds",
        "ListBasedCmds",
        "SetBasedCmds",
        "SortedSetBasedCmds",
        "CurrItems",
    }
    var dimensionParam = &cloudwatch.Dimension{
        Name:  aws.String("CacheClusterId"),
        Value: aws.String(*cluster_name),
    }

    var duration = -1 * time.Duration(*fetch_age) * time.Second
    var et = time.Now().Add(duration)
    var st = et.Add(duration)

    for index := range metrics {
        mt := &cloudwatch.GetMetricStatisticsInput{
            Dimensions: []cloudwatch.Dimension{*dimensionParam},
            StartTime:  st,
            EndTime:    et,
            MetricName: aws.String(metrics[index]),
            Namespace:  aws.String("AWS/ElastiCache"),
            Period:     aws.Integer(300),
            Statistics: []string{"Average"},
        }
        resp, _ := cw.GetMetricStatistics(mt)

        if len(resp.Datapoints) > 0 {
            var dp = resp.Datapoints[0]
            fmt.Print("aws.elasticache.", *cluster_name, ".", metrics[index], ".average")
            fmt.Print("\t", *dp.Average)
            fmt.Print("\t", dp.Timestamp.Unix(), "\n")
        }
    }
}
