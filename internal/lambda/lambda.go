package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Lambda struct {
	Conf Config
}

// See https://docs.aws.amazon.com/lambda/latest/dg/golang-context.html for
// details on working with the `lambdacontext`
func (l Lambda) Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

}
