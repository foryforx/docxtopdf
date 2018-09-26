package usecase

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/karuppaiah/docxtopdf/pdfgenerate/model"
)

///////WORKER
// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan model.WorkRequest) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan model.WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan model.WorkRequest
	WorkerQueue chan chan model.WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())

				//time.Sleep(work.Delay)
				fmt.Printf("Before Generating file:%s", work.File)
				err := RConvertUsingLibreOffice(work.File)
				if err != nil {
					fmt.Printf("Error Generating file:%s", work.File)
					// Failure recovery code here
				} else {
					fmt.Printf("Generated file:%s", work.File)
				}

				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

//ConvertUsingLibreOffice will use libreoffice to convert the docsx to pdf
func RConvertUsingLibreOffice(inputFile string) error {
	executable := "libreoffice"
	if os.Getenv("GOOS") == "darwin" || os.Getenv("GOOS") == "" {
		executable = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
	}
	fmt.Println(os.Getenv("GOOS"))

	fmt.Println(executable)
	docxToPDFCmd := exec.Command("bash", "-c", executable+" --headless --convert-to pdf "+inputFile)
	docxToPDFCmdErr := docxToPDFCmd.Run()
	if docxToPDFCmdErr != nil {
		fmt.Println("1.Conversion issue to pdf", docxToPDFCmdErr)
		return docxToPDFCmdErr
	}
	return nil

}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
