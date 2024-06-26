package comman_function

import (
	"log"
	"time"

	"github.com/Appkube-awsx/awsx-common/awsclient"
	"github.com/Appkube-awsx/awsx-common/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func GetMetricData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, dimensionsName string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
	input := &cloudwatch.GetMetricDataInput{
		EndTime:   endTime,
		StartTime: startTime,
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			{
				Id: aws.String("m1"),
				MetricStat: &cloudwatch.MetricStat{

					Metric: &cloudwatch.Metric{
						Dimensions: []*cloudwatch.Dimension{
							{
								Name:  aws.String(dimensionsName),
								Value: aws.String(instanceID),
							},
						},
						MetricName: aws.String(metricName),
						Namespace:  aws.String(elementType),
					},
					Period: aws.Int64(300),
					Stat:   aws.String(statistic),
				},
			},
		},
	}
	if cloudWatchClient == nil {
		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
	}

	result, err := cloudWatchClient.GetMetricData(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//
//func GetMetricClusterData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
//	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
//	input := &cloudwatch.GetMetricDataInput{
//		EndTime:   endTime,
//		StartTime: startTime,
//		MetricDataQueries: []*cloudwatch.MetricDataQuery{
//			{
//				Id: aws.String("m1"),
//				MetricStat: &cloudwatch.MetricStat{
//
//					Metric: &cloudwatch.Metric{
//						Dimensions: []*cloudwatch.Dimension{
//							{
//								Name:  aws.String("ClusterName"),
//								Value: aws.String(instanceID),
//							},
//						},
//						MetricName: aws.String(metricName),
//						Namespace:  aws.String(elementType),
//					},
//					Period: aws.Int64(300),
//					Stat:   aws.String(statistic),
//				},
//			},
//		},
//	}
//	if cloudWatchClient == nil {
//		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
//	}
//
//	result, err := cloudWatchClient.GetMetricData(input)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func GetMetricDatabaseData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
//	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
//	input := &cloudwatch.GetMetricDataInput{
//		EndTime:   endTime,
//		StartTime: startTime,
//		MetricDataQueries: []*cloudwatch.MetricDataQuery{
//			{
//				Id: aws.String("m1"),
//				MetricStat: &cloudwatch.MetricStat{
//
//					Metric: &cloudwatch.Metric{
//						Dimensions: []*cloudwatch.Dimension{
//							{
//								Name:  aws.String("DBInstanceIdentifier"),
//								Value: aws.String(instanceID),
//							},
//						},
//						MetricName: aws.String(metricName),
//						Namespace:  aws.String(elementType),
//					},
//					Period: aws.Int64(300),
//					Stat:   aws.String(statistic),
//				},
//			},
//		},
//	}
//	if cloudWatchClient == nil {
//		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
//	}
//
//	result, err := cloudWatchClient.GetMetricData(input)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func GetMetricLoadBalancerData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
//	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
//	input := &cloudwatch.GetMetricDataInput{
//		EndTime:   endTime,
//		StartTime: startTime,
//		MetricDataQueries: []*cloudwatch.MetricDataQuery{
//			{
//				Id: aws.String("m1"),
//				MetricStat: &cloudwatch.MetricStat{
//
//					Metric: &cloudwatch.Metric{
//						Dimensions: []*cloudwatch.Dimension{
//							{
//								Name:  aws.String("LoadBalancer"),
//								Value: aws.String(instanceID),
//							},
//						},
//						MetricName: aws.String(metricName),
//						Namespace:  aws.String(elementType),
//					},
//					Period: aws.Int64(300),
//					Stat:   aws.String(statistic),
//				},
//			},
//		},
//	}
//	if cloudWatchClient == nil {
//		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
//	}
//
//	result, err := cloudWatchClient.GetMetricData(input)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//
//func GetMetricFunctionNameData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
//	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
//	input := &cloudwatch.GetMetricDataInput{
//		EndTime:   endTime,
//		StartTime: startTime,
//		MetricDataQueries: []*cloudwatch.MetricDataQuery{
//			{
//				Id: aws.String("m1"),
//				MetricStat: &cloudwatch.MetricStat{
//
//					Metric: &cloudwatch.Metric{
//						Dimensions: []*cloudwatch.Dimension{
//							{
//								Name:  aws.String("FunctionName"),
//								Value: aws.String(instanceID),
//							},
//						},
//						MetricName: aws.String(metricName),
//						Namespace:  aws.String(elementType),
//					},
//					Period: aws.Int64(300),
//					Stat:   aws.String(statistic),
//				},
//			},
//		},
//	}
//	if cloudWatchClient == nil {
//		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
//	}
//
//	result, err := cloudWatchClient.GetMetricData(input)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
//func GetMetricAPIData(clientAuth *model.Auth, instanceID, elementType string, metricName string, startTime, endTime *time.Time, statistic string, cloudWatchClient *cloudwatch.CloudWatch) (*cloudwatch.GetMetricDataOutput, error) {
//	log.Printf("Getting metric data for instance %s in namespace %s from %v to %v", instanceID, elementType, startTime, endTime)
//	input := &cloudwatch.GetMetricDataInput{
//		EndTime:   endTime,
//		StartTime: startTime,
//		MetricDataQueries: []*cloudwatch.MetricDataQuery{
//			{
//				Id: aws.String("m1"),
//				MetricStat: &cloudwatch.MetricStat{
//
//					Metric: &cloudwatch.Metric{
//						Dimensions: []*cloudwatch.Dimension{
//							{
//								Name:  aws.String("ApiName"),
//								Value: aws.String(instanceID),
//							},
//						},
//						MetricName: aws.String(metricName),
//						Namespace:  aws.String(elementType),
//					},
//					Period: aws.Int64(300),
//					Stat:   aws.String(statistic),
//				},
//			},
//		},
//	}
//	if cloudWatchClient == nil {
//		cloudWatchClient = awsclient.GetClient(*clientAuth, awsclient.CLOUDWATCH).(*cloudwatch.CloudWatch)
//	}
//
//	result, err := cloudWatchClient.GetMetricData(input)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
