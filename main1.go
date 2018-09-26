package main

// import (
// 	"flag"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// //Convertor is the custom type to hold the docx to pdf conversion values
// type PDFConvertor struct {
// 	inputFile string
// 	method    string
// }

// func main1() {
// 	//Check if input file is docx

// 	//Check if libreoffice or textutil/cupsfilter is present

// 	//
// 	if len(os.Args) == 3 {
// 		pdfConv := PDFConvertor{inputFile: os.Args[1], method: os.Args[2]}
// 		//First solution
// 		if pdfConv.method == "1" {
// 			err := ConvertUsingTextUtil(pdfConv.inputFile)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			return
// 		} else if pdfConv.method == "2" {
// 			err := ConvertUsingLibreOffice(pdfConv.inputFile)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			return
// 		} else {
// 			fmt.Println("Please enter the method as 1 or 2")
// 		}

// 	} else {
// 		fmt.Println("Please enter the files as argument and method")
// 	}

// 	// } else if len(os.Args) == 2 {
// 	// 	//Second solution
// 	// 	//soffice --headless --convert-to pdf sample.docx
// 	// 	err := ConvertUsingLibreOffice(os.Args[1])
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}

// 	// }

// }

// func main2() {
// 	// ch := make(chan int, 100)
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("started ", i, " routine")
// 		go RConvertUsingLibreOffice("sample" + strconv.Itoa(i) + ".docx")
// 	}

// }

// // ConvertUsingTextUtil uses textutil and cupsfilter to convert docs to pdf
// // textutil -convert html -output "sample.html" sample.docx
// // cupsfilter sample.html > sample.pdf
// func ConvertUsingTextUtil(inputFile string) error {

// 	tempOutputHTMLFile := strings.Replace(inputFile, ".docx", ".html", -1)
// 	outputFile := strings.Replace(inputFile, ".docx", ".pdf", -1)

// 	docxToHTMLCmd := exec.Command("bash", "-c", "textutil -convert html -output "+tempOutputHTMLFile+" "+inputFile)
// 	docxToHTMLCmdErr := docxToHTMLCmd.Run()
// 	if docxToHTMLCmdErr != nil {
// 		fmt.Println("1.Conversion issue to pdf")
// 		return docxToHTMLCmdErr
// 	} else {
// 		htmlToPdfCmd := exec.Command("bash", "-c", "cupsfilter "+tempOutputHTMLFile+" > "+outputFile)
// 		htmlToPdfCmdErr := htmlToPdfCmd.Run()
// 		if htmlToPdfCmdErr != nil {
// 			fmt.Println("2.Conversion issue to pdf")
// 			return htmlToPdfCmdErr
// 		}
// 	}

// 	return nil
// }

// //ConvertUsingLibreOffice will use libreoffice to convert the docsx to pdf
// func ConvertUsingLibreOffice(inputFile string) error {
// 	executable := "libreoffice"
// 	if os.Getenv("GOOS") == "darwin" || os.Getenv("GOOS") == "" {

// 		executable = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
// 	}
// 	fmt.Println(os.Getenv("GOOS"))

// 	fmt.Println(executable)
// 	docxToPDFCmd := exec.Command("bash", "-c", executable+" --headless --convert-to pdf "+inputFile)
// 	docxToPDFCmdErr := docxToPDFCmd.Run()
// 	if docxToPDFCmdErr != nil {
// 		fmt.Println("1.Conversion issue to pdf", docxToPDFCmdErr)
// 		return docxToPDFCmdErr
// 	}
// 	return nil
// }

// //ValidateInputFile is to validate if the inputfile value is correct
// func (c *PDFConvertor) ValidateInputFile() (bool, error) {
// 	if c.inputFile == "" {
// 		err := fmt.Errorf("Convertor| InputFile Empty")
// 		return false, err
// 	}
// 	if strings.Index(c.inputFile, ".docx") == -1 {
// 		err := fmt.Errorf("Convertor| InputFile is not in docx format")
// 		return false, err
// 	}

// 	return true, nil

// }

// // GO ROUTINES

// type WorkRequest struct {
// 	Name  string
// 	Delay time.Duration
// 	File  string
// }

// //COLLECTOR
// // A buffered channel that we can send work requests on.
// var WorkQueue = make(chan WorkRequest, 100)

// func Collector(w http.ResponseWriter, r *http.Request) {
// 	// Make sure we can only be called with an HTTP POST request.
// 	if r.Method != "POST" {
// 		w.Header().Set("Allow", "POST")
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Parse the delay.
// 	delay, err := time.ParseDuration(r.FormValue("delay"))
// 	if err != nil {
// 		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Check to make sure the delay is anywhere from 1 to 10 seconds.
// 	if delay.Seconds() < 1 || delay.Seconds() > 10 {
// 		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
// 		return
// 	}

// 	// Now, we retrieve the person's name from the request.
// 	name := r.FormValue("name")

// 	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
// 	if name == "" {
// 		http.Error(w, "You must specify a name.", http.StatusBadRequest)
// 		return
// 	}
// 	// Now, we retrieve the person's name from the request.
// 	filename := r.FormValue("filename")
// 	// Just do a quick bit of sanity checking to make sure the client actually provided us with a name.
// 	if filename == "" {
// 		http.Error(w, "You must specify a docx filename.", http.StatusBadRequest)
// 		return
// 	}
// 	CollectorJob(name, delay, filename)
// 	// And let the user know their work request was created.
// 	w.WriteHeader(http.StatusCreated)
// 	return
// }

// // CollectorJob is the function to be called for PDF generation with requestor name of work and filename with path for docx
// func CollectorJob(name string, delay time.Duration, filename string) {
// 	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
// 	work := WorkRequest{Name: name, Delay: delay, File: filename}
// 	// Push the work onto the queue.
// 	WorkQueue <- work
// 	fmt.Println("Work request queued for PDF generation")
// }

// ///////WORKER
// // NewWorker creates, and returns a new Worker object. Its only argument
// // is a channel that the worker can add itself to whenever it is done its
// // work.
// func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
// 	// Create, and return the worker.
// 	worker := Worker{
// 		ID:          id,
// 		Work:        make(chan WorkRequest),
// 		WorkerQueue: workerQueue,
// 		QuitChan:    make(chan bool)}

// 	return worker
// }

// type Worker struct {
// 	ID          int
// 	Work        chan WorkRequest
// 	WorkerQueue chan chan WorkRequest
// 	QuitChan    chan bool
// }

// // This function "starts" the worker by starting a goroutine, that is
// // an infinite "for-select" loop.
// func (w *Worker) Start() {
// 	go func() {
// 		for {
// 			// Add ourselves into the worker queue.
// 			w.WorkerQueue <- w.Work

// 			select {
// 			case work := <-w.Work:
// 				// Receive a work request.
// 				fmt.Printf("worker%d: Received work request, delaying for %f seconds\n", w.ID, work.Delay.Seconds())

// 				//time.Sleep(work.Delay)
// 				fmt.Printf("Before Generating file:%s", work.File)
// 				err := RConvertUsingLibreOffice(work.File)
// 				if err != nil {
// 					fmt.Printf("Error Generating file:%s", work.File)
// 					// Failure recovery code here
// 				} else {
// 					fmt.Printf("Generated file:%s", work.File)
// 				}

// 				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

// 			case <-w.QuitChan:
// 				// We have been asked to stop.
// 				fmt.Printf("worker%d stopping\n", w.ID)
// 				return
// 			}
// 		}
// 	}()
// }

// //ConvertUsingLibreOffice will use libreoffice to convert the docsx to pdf
// func RConvertUsingLibreOffice(inputFile string) error {
// 	executable := "libreoffice"
// 	if os.Getenv("GOOS") == "darwin" || os.Getenv("GOOS") == "" {
// 		executable = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
// 	}
// 	fmt.Println(os.Getenv("GOOS"))

// 	fmt.Println(executable)
// 	docxToPDFCmd := exec.Command("bash", "-c", executable+" --headless --convert-to pdf "+inputFile)
// 	docxToPDFCmdErr := docxToPDFCmd.Run()
// 	if docxToPDFCmdErr != nil {
// 		fmt.Println("1.Conversion issue to pdf", docxToPDFCmdErr)
// 		return docxToPDFCmdErr
// 	}
// 	return nil

// }

// // Stop tells the worker to stop listening for work requests.
// //
// // Note that the worker will only stop *after* it has finished its work.
// func (w *Worker) Stop() {
// 	go func() {
// 		w.QuitChan <- true
// 	}()
// }

// //DISPATCHER

// var WorkerQueue chan chan WorkRequest

// func StartDispatcher(nworkers int) {
// 	// First, initialize the channel we are going to but the workers' work channels into.
// 	WorkerQueue = make(chan chan WorkRequest, nworkers)

// 	// Now, create all of our workers.
// 	for i := 0; i < nworkers; i++ {
// 		fmt.Println("Starting worker", i+1)
// 		worker := NewWorker(i+1, WorkerQueue)
// 		worker.Start()
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case work := <-WorkQueue:
// 				fmt.Println("Received work requeust")
// 				go func() {
// 					worker := <-WorkerQueue

// 					fmt.Println("Dispatching work request")
// 					worker <- work
// 				}()
// 			}
// 		}
// 	}()
// }

// var (
// 	NWorkers = flag.Int("n", 1, "The number of workers to start")
// 	HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")
// )

// func main3() {
// 	// Parse the command-line flags.
// 	flag.Parse()

// 	// Start the dispatcher.
// 	fmt.Println("Starting the dispatcher")
// 	StartDispatcher(*NWorkers)

// 	// Register our collector as an HTTP handler function.
// 	fmt.Println("Registering the collector")
// 	http.HandleFunc("/work", Collector)

// 	// Start the HTTP server!
// 	fmt.Println("HTTP server listening on", *HTTPAddr)
// 	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
