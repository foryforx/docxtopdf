package usecase

import (
	"fmt"
	"time"

	"github.com/karuppaiah/docxtopdf/pdfgenerate/model"
)

// CollectorJob is the function to be called for PDF generation with requestor name of work and filename with path for docx
func CollectorJob(name string, delay time.Duration, filename string) {
	// Now, we take the delay, and the person's name, and make a model.WorkRequest out of them.
	work := model.WorkRequest{Name: name, Delay: delay, File: filename}
	// Push the work onto the queue.
	model.WorkQueue <- work
	fmt.Println("Work request queued for PDF generation")
}
