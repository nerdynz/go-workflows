package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nerdynz/go-workflows/backend"
	"github.com/nerdynz/go-workflows/client"
	"github.com/nerdynz/go-workflows/diag"
	"github.com/nerdynz/go-workflows/samples"
	"github.com/nerdynz/go-workflows/worker"

	"github.com/google/uuid"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	b := samples.GetBackend("continue-as-new")

	db, ok := b.(diag.Backend)
	if !ok {
		panic("backend does not implement diag.Backend")
	}

	// Start diagnostic server under /diag
	m := http.NewServeMux()
	m.Handle("/diag/", http.StripPrefix("/diag", diag.NewServeMux(db)))
	go http.ListenAndServe(":3000", m)

	// Run worker
	w := RunWorker(ctx, b)

	// Start workflow via client
	c := client.New(b)

	runWorkflow(ctx, c)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	cancel()

	if err := w.WaitForCompletion(); err != nil {
		panic("could not stop worker" + err.Error())
	}
}

func runWorkflow(ctx context.Context, c *client.Client) {
	wf, err := c.CreateWorkflowInstance(ctx, client.WorkflowInstanceOptions{
		InstanceID: uuid.NewString(),
	}, Workflow1, 0, 4, Inputs{
		Msg:    "",
		Result: 0,
	})
	if err != nil {
		log.Fatal(err)
		panic("could not start workflow")
	}

	result, err := client.GetWorkflowResult[int](ctx, c, wf, time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Workflow finished. Result:", result)
}

func RunWorker(ctx context.Context, mb backend.Backend) *worker.Worker {
	w := worker.New(mb, nil)

	w.RegisterWorkflow(Workflow1)
	w.RegisterWorkflow(SubWorkflow)

	w.RegisterActivity(Activity1)

	if err := w.Start(ctx); err != nil {
		panic("could not start worker")
	}

	return w
}
