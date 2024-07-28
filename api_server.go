package eks

import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/eks"
)

func APIServerURL(ctx context.Context, cfg *aws.Config, clusterName string) (string, error) {
    // Describe the EKS cluster to get the API server endpoint
    result, err := eks.NewFromConfig(cfg).DescribeCluster(ctx, &eks.DescribeClusterInput{
        Name: aws.String(clusterName),
    })
    if err != nil {
        return "", fmt.Errorf("failed to describe cluster: %w", err)
    }
    // Get the cluster API server endpoint
    return *result.Cluster.Endpoint, nil
}
