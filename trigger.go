package meralambda

import (
	"context"
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	// Import the aws-lambda-go. Required for dep to pull on app create
	//	_ "github.com/aws/aws-lambda-go/lambda"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-lambda")
var singleton *LambdaTrigger

// LambdaTrigger AWS Lambda trigger struct
type LambdaTrigger struct {
	metadata *trigger.Metadata
	config   *trigger.Config
	handlers []*trigger.Handler
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &LambdaFactory{metadata: md}
}

// LambdaFactory AWS Lambda Trigger factory
type LambdaFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *LambdaFactory) New(config *trigger.Config) trigger.Trigger {
	singleton = &LambdaTrigger{metadata: t.metadata, config: config}
	return singleton
}

// Metadata implements trigger.Trigger.Metadata
func (t *LambdaTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *LambdaTrigger) Initialize(ctx trigger.InitContext) error {
	t.handlers = ctx.GetHandlers()
	return nil
}

func Invoke() {
	handler := singleton.handlers[0]
	fmt.Println("Invoke hua bhau")
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker bhi hua bhau")
			data := map[string]interface{}{
				"context": t,
				"evt":     time.Now(),
			}
			fmt.Println("Submitting Job to Handle")
			results, err := handler.Handle(context.Background(), data)
			fmt.Println(results, err)
		}

	}()
}

func (t *LambdaTrigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *LambdaTrigger) Stop() error {
	return nil
}
