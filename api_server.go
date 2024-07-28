package eks

import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/eks"
)

// APIServerURL returns the API server URL for the specified EKS cluster.
func APIServerURL(ctx context.Context, cfg *aws.Config, clusterName string) (string, error) {
    // Load default AWS configuration if cfg is nil
    if cfg == nil {
        var err error
        if cfg, err = config.LoadDefaultConfig(ctx); err != nil {
            return "", fmt.Errorf("unable to load default config: %w", err)
        }
    }

    // Describe the EKS cluster to get the API server endpoint
    client := eks.NewFromConfig(*cfg)
    result, err := client.DescribeCluster(ctx, &eks.DescribeClusterInput{
        Name: aws.String(clusterName),
    })
    if err != nil {
        return "", fmt.Errorf("failed to describe cluster: %w", err)
    }

    // Ensure the result and endpoint are not nil
    if result.Cluster == nil {
        return "", fmt.Errorf("bad result from client.DescribeCluster: cluster is nil")
    }
    if result.Cluster.Endpoint == nil {
        return "", fmt.Errorf("bad result from client.DescribeCluster: cluster endpoint is nil")
    }

    return *result.Cluster.Endpoint, nil
}
