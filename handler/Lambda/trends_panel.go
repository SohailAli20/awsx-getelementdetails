package Lambda

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/spf13/cobra"
)

type ErrorResult struct {
	Value            float64 `json:"Value"`
	PercentageChange float64 `json:"PercentageChange"`
	ChangeType       string  `json:"ChangeType"`
}

var AwsxLambdaTrendsCmd = &cobra.Command{
	Use:   "trends_panel",
	Short: "get trends metrics data",
	Long:  `command to get trends metrics data`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running from child command")
		var authFlag, clientAuth, err = authenticate.AuthenticateCommand(cmd)
		if err != nil {
			log.Printf("Error during authentication: %v\n", err)
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
		if authFlag {
			responseType, _ := cmd.PersistentFlags().GetString("responseType")
			jsonResp, cloudwatchMetricResp, err := GetLambdaTrendsData(cmd, clientAuth, nil)
			if err != nil {
				log.Println("Error getting lambda trends data : ", err)
				return
			}
			if responseType == "frame" {
				fmt.Println(cloudwatchMetricResp)
			} else {
				fmt.Println(jsonResp)
			}
		}

	},
}

func GetLambdaTrendsData(cmd *cobra.Command, clientAuth *model.Auth, cloudWatchClient *cloudwatch.CloudWatch) (string, map[string]interface{}, error) {
	startTimeStr, _ := cmd.PersistentFlags().GetString("startTime")
	endTimeStr, _ := cmd.PersistentFlags().GetString("endTime")

	var startTime, endTime *time.Time

	if startTimeStr != "" {
		parsedStartTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			log.Printf("Error parsing start time: %v", err)
			return "", nil, err
		}
		startTime = &parsedStartTime
	} else {
		defaultStartTime := time.Now().Add(-5 * time.Minute)
		startTime = &defaultStartTime
	}

	if endTimeStr != "" {
		parsedEndTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			log.Printf("Error parsing end time: %v", err)
			return "", nil, err
		}
		endTime = &parsedEndTime
	} else {
		defaultEndTime := time.Now()
		endTime = &defaultEndTime
	}

	// Debug prints
	log.Printf("StartTime: %v, EndTime: %v", startTime, endTime)

	cloudwatchMetricData := map[string]interface{}{}

	// Fetch raw data for last month and current month
	lastMonthStartTime := startTime.AddDate(0, -1, 0)
	lastMonthEndTime := endTime.AddDate(0, -1, 0)
	lastMonthMemory, err := GetTotalLambdaInvocations(clientAuth, &lastMonthStartTime, &lastMonthEndTime, cloudWatchClient)
	if err != nil {
		log.Println("Error in getting error metric value for last month: ", err)
		return "", nil, err
	}

	currentMonthMemory, err := GetTotalLambdaInvocations(clientAuth, startTime, endTime, cloudWatchClient)
	if err != nil {
		log.Println("Error in getting error metric value for current month: ", err)
		return "", nil, err
	}

	fmt.Println(lastMonthMemory, currentMonthMemory)
	// Calculate percentage change
	percentageChange := ((currentMonthMemory - lastMonthMemory) / lastMonthMemory) * 100

	// Determine if it's an increment or decrement
	changeType := "increment"
	if percentageChange < 0 {
		changeType = "decrement"
	}

	cloudwatchMetricData["LastMonthMemory"] = lastMonthMemory
	cloudwatchMetricData["CurrentMemory"] = currentMonthMemory
	cloudwatchMetricData["PercentageChange"] = fmt.Sprintf("%.2f%% %s", percentageChange, changeType)

	jsonString, err := json.Marshal(ErrorResult{Value: currentMonthMemory, PercentageChange: percentageChange, ChangeType: changeType})
	if err != nil {
		log.Println("Error in marshalling json in string: ", err)
		return "", nil, err
	}

	return string(jsonString), cloudwatchMetricData, nil
}

func GetTotalLambdaInvocations(clientAuth *model.Auth, startTime, endTime *time.Time, cloudWatchClient *cloudwatch.CloudWatch) (float64, error) {
	input := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String("AWS/Lambda"),
		MetricName: aws.String("Invocations"),
		StartTime:  startTime,
		EndTime:    endTime,
		Period:     aws.Int64(300), // Adjust period as needed (e.g., 5 minutes)
		Statistics: []*string{aws.String("Sum")},
	}

	if cloudWatchClient == nil {
		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
	}

	result, err := cloudWatchClient.GetMetricStatistics(input)
	if err != nil {
		return 0, err
	}

	if len(result.Datapoints) == 0 {
		return 0, fmt.Errorf("no data available for the specified time range")
	}

	// Sum up the values from all the datapoints
	totalInvocations := 0.0
	for _, dp := range result.Datapoints {
		totalInvocations += aws.Float64Value(dp.Sum)
	}

	return totalInvocations, nil
}

func init() {
	AwsxLambdaTrendsCmd.PersistentFlags().String("elementId", "", "element id")
	AwsxLambdaTrendsCmd.PersistentFlags().String("elementType", "", "element type")
	AwsxLambdaTrendsCmd.PersistentFlags().String("query", "", "query")
	AwsxLambdaTrendsCmd.PersistentFlags().String("cmdbApiUrl", "", "cmdb api")
	AwsxLambdaTrendsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxLambdaTrendsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxLambdaTrendsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxLambdaTrendsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxLambdaTrendsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxLambdaTrendsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxLambdaTrendsCmd.PersistentFlags().String("externalId", "", "aws external id")
	AwsxLambdaTrendsCmd.PersistentFlags().String("cloudWatchQueries", "", "aws cloudwatch metric queries")
	AwsxLambdaTrendsCmd.PersistentFlags().String("instanceId", "", "instance id")
	AwsxLambdaTrendsCmd.PersistentFlags().String("startTime", "", "start time")
	AwsxLambdaTrendsCmd.PersistentFlags().String("endTime", "", "end time")
	AwsxLambdaTrendsCmd.PersistentFlags().String("responseType", "", "response type. json/frame")
}
