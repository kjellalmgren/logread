package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

//
func main() {

	verbose := false
	startTime := time.Now()
	//Check for command-line argument filename.
	//Ignore additional arguments.
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" {
			verbose = true
			fmt.Printf("Verbose is set to true\r\n")
		}
	}
	log.Println("Starting...")
	path := "./"
	fmt.Println("Starting log-analysis of lp.log...")
	//
	//var filenameList = []string{"lp1.log", "lp1_audit.log", "lp2.log", "lp2_audit.log",
	//	"2017.03.18.log", "2017.03.24.log", "2017.03.25.log", "lp3.log", "lp4.log"}
	var filenameList = []string{
		//"./LHBLP101/2017-04-05/lp.log", "./LHBLP102/2017-04-05/lp.log",
		//"./LHBLP101/2017-04-06/lp.log", "./LHBLP102/2017-04-06/lp.log",
		//"./LHBLP101/2017-04-07/lp.log", "./LHBLP102/2017-04-07/lp.log",
		//"./LHBLP101/2017-04-08/lp.log", "./LHBLP102/2017-04-08/lp.log",
		//"./LHBLP101/2017-04-09/lp.log", "./LHBLP102/2017-04-09/lp.log",
		//"./LHBLP101/2017-04-10/lp.log", "./LHBLP102/2017-04-10/lp.log",
		//"./LHBLP101/2017-04-11/lp.log", "./LHBLP102/2017-04-11/lp.log",
		//"./LHBLP101/2017-04-12/lp.log", "./LHBLP102/2017-04-12/lp.log",
		//"./LHBLP101/2017-04-13/lp.log", "./LHBLP102/2017-04-13/lp.log",
		//"./LHBLP101/2017-04-14/lp.log", "./LHBLP102/2017-04-14/lp.log",
		//"./LHBLP101/2017-04-15/lp.log", "./LHBLP102/2017-04-15/lp.log",
		//"./LHBLP101/2017-04-16/lp.log", "./LHBLP102/2017-04-16/lp.log",
		//"./LHBLP101/2017-04-17/lp.log", "./LHBLP102/2017-04-17/lp.log",
		//"./LHBLP101/2017-04-18/lp.log", "./LHBLP102/2017-04-18/lp.log",
		//"./LHBLP101/2017-04-19/lp.log", "./LHBLP102/2017-04-19/lp.log",
		//"./LHBLP101/2017-04-20/lp.log", "./LHBLP102/2017-04-20/lp.log",
		//"./LHBLP101/2017-04-21/lp.log", "./LHBLP102/2017-04-21/lp.log",
		//"./LHBLP101/2017-04-22/lp.log", "./LHBLP102/2017-04-22/lp.log",
		//"./LHBLP101/2017-04-23/lp.log", "./LHBLP102/2017-04-23/lp.log",
		//"./LHBLP101/2017-04-24/lp.log", "./LHBLP102/2017-04-24/lp.log",
		//"./LHBLP101/2017-04-25/lp.log", "./LHBLP102/2017-04-25/lp.log",
		//"./LHBLP101/2017-04-26/lp.log", "./LHBLP102/2017-04-26/lp.log",
		//"./LHBLP101/2017-04-27/lp.log", "./LHBLP102/2017-04-27/lp.log",
		//"./LHBLP101/2017-04-28/lp.log", "./LHBLP102/2017-04-28/lp.log",
		//"./LHBLP101/2017-04-29/lp.log", "./LHBLP102/2017-04-29/lp.log",
		//"./LHBLP101/2017-04-30/lp.log", "./LHBLP102/2017-04-30/lp.log",
		//"./LHBLP101/2017-05-01/lp.log", "./LHBLP102/2017-05-01/lp.log",
		//"./LHBLP101/2017-05-02/lp.log", "./LHBLP102/2017-05-02/lp.log",
		//"./LHBLP101/2017-05-03/lp.log", "./LHBLP102/2017-05-03/lp.log",
		//"./LHBLP101/2017-05-04/lp.log", "./LHBLP102/2017-05-04/lp.log",
		//"./LHBLP101/2017-05-05/lp.log", "./LHBLP102/2017-05-05/lp.log",
		//"./LHBLP101/2017-05-06/lp.log", "./LHBLP102/2017-05-06/lp.log",
		//"./LHBLP101/2017-05-07/lp.log", "./LHBLP102/2017-05-07/lp.log",
		//"./LHBLP101/2017-05-08/lp.log", "./LHBLP102/2017-05-08/lp.log",
		//"./LHBLP101/2017-05-09/lp.log", "./LHBLP102/2017-05-09/lp.log",
		//"./LHBLP101/2017-05-10/lp.log", "./LHBLP102/2017-05-10/lp.log",
		//"./LHBLP101/2017-05-11/lp.log", "./LHBLP102/2017-05-11/lp.log",
		//"./LHBLP101/2017-05-12/lp.log", "./LHBLP102/2017-05-12/lp.log",
		//"./LHBLP101/2017-05-12/lp-1.log", "./LHBLP102/2017-05-12/lp-1.log",
		//"./LHBLP101/2017-05-13/lp.log", "./LHBLP102/2017-05-13/lp.log",
		//"./LHBLP101/2017-05-14/lp.log", "./LHBLP102/2017-05-14/lp.log",
		//"./LHBLP101/2017-05-15/lp.log", "./LHBLP102/2017-05-15/lp.log",
		//"./LHBLP101/2017-05-16/lp.log", "./LHBLP102/2017-05-16/lp.log",
		//"./LHBLP101/2017-05-17/lp.log", "./LHBLP102/2017-05-17/lp.log",
		//"./LHBLP101/2017-05-18/lp.log", "./LHBLP102/2017-05-18/lp.log",
		//"./LHBLP101/2017-05-19/lp.log", "./LHBLP102/2017-05-19/lp.log",
		//"./LHBLP101/2017-05-20/lp.log", "./LHBLP102/2017-05-20/lp.log",
		//"./LHBLP101/2017-05-21/lp.log", "./LHBLP102/2017-05-21/lp.log",
		//"./LHBLP101/2017-05-22/lp.log", "./LHBLP102/2017-05-22/lp.log",
		//"./LHBLP101/2017-05-23/lp.log", "./LHBLP102/2017-05-23/lp.log",
		//"./LHBLP101/2017-05-24/lp.log", "./LHBLP102/2017-05-24/lp.log",
		//"./LHBLP101/2017-05-25/lp.log", "./LHBLP102/2017-05-25/lp.log",
		//"./LHBLP101/2017-05-26/lp.log", "./LHBLP102/2017-05-26/lp.log",
		//"./LHBLP101/2017-05-27/lp.log", "./LHBLP102/2017-05-27/lp.log",
		//"./LHBLP101/2017-05-28/lp.log", "./LHBLP102/2017-05-28/lp.log",
		//"./LHBLP101/2017-05-29/lp.log", "./LHBLP102/2017-05-29/lp.log",
		//"./LHBLP101/2017-05-30/lp.log", "./LHBLP102/2017-05-30/lp.log",
		//"./LHBLP101/2017-05-31/lp.log", "./LHBLP102/2017-05-31/lp.log",
		//"./LHBLP101/2017-06-01/lp.log", "./LHBLP102/2017-06-01/lp.log",
		//"./LHBLP101/2017-06-02/lp.log", "./LHBLP102/2017-06-02/lp.log",
		//"./LHBLP101/2017-06-03/lp.log", "./LHBLP102/2017-06-03/lp.log",
		//"./LHBLP101/2017-06-04/lp.log", "./LHBLP102/2017-06-04/lp.log",
		//"./LHBLP101/2017-06-05/lp.log", "./LHBLP102/2017-06-05/lp.log",
		//"./LHBLP101/2017-06-06/lp.log", "./LHBLP102/2017-06-06/lp.log",
		//"./LHBLP101/2017-06-07/lp.log", "./LHBLP102/2017-06-07/lp.log",
		//"./LHBLP101/2017-06-08/lp.log", "./LHBLP102/2017-06-08/lp.log",
		//"./LHBLP101/2017-06-09/lp.log", "./LHBLP102/2017-06-09/lp.log",
		//"./LHBLP101/2017-06-10/lp.log", "./LHBLP102/2017-06-10/lp.log",
		//"./LHBLP101/2017-06-11/lp.log", "./LHBLP102/2017-06-11/lp.log",
		//"./LHBLP101/2017-06-12/lp.log", "./LHBLP102/2017-06-12/lp.log",
		//"./LHBLP101/2017-06-13/lp.log", "./LHBLP102/2017-06-13/lp.log",
		//"./LHBLP101/2017-06-14/lp.log", "./LHBLP102/2017-06-14/lp.log",
		//"./LHBLP101/2017-06-15/lp.log", "./LHBLP102/2017-06-15/lp.log",
		//"./LHBLP101/2017-06-16/lp.log", "./LHBLP102/2017-06-16/lp.log",
		//"./LHBLP101/2017-06-17/lp.log", "./LHBLP102/2017-06-17/lp.log",
		//"./LHBLP101/2017-06-18/lp.log", "./LHBLP102/2017-06-18/lp.log",
		//"./LHBLP101/2017-06-19/lp.log", "./LHBLP102/2017-06-19/lp.log",
		//"./LHBLP101/2017-06-20/lp.log", "./LHBLP102/2017-06-20/lp.log",
		//"./LHBLP101/2017-06-21/lp.log", "./LHBLP102/2017-06-21/lp.log",
		//"./LHBLP101/2017-06-22/lp.log", "./LHBLP102/2017-06-22/lp.log",
		//"./LHBLP101/2017-06-23/lp.log", "./LHBLP102/2017-06-23/lp.log",
		//"./LHBLP101/2017-06-24/lp.log", "./LHBLP102/2017-06-24/lp.log",
		//"./LHBLP101/2017-06-25/lp.log", "./LHBLP102/2017-06-25/lp.log",
		//"./LHBLP101/2017-06-26/lp.log", "./LHBLP102/2017-06-26/lp.log",
		//"./LHBLP101/2017-06-27/lp.log", "./LHBLP102/2017-06-27/lp.log",
		//"./LHBLP101/2017-06-28/lp.log", "./LHBLP102/2017-06-28/lp.log",
		//"./LHBLP101/2017-06-29/lp.log", "./LHBLP102/2017-06-29/lp.log",
		//"./LHBLP101/2017-06-30/lp.log", "./LHBLP102/2017-06-30/lp.log",
		//"./LHBLP101/2017-07-01/lp.log", "./LHBLP102/2017-07-01/lp.log",
		//"./LHBLP101/2017-07-02/lp.log", "./LHBLP102/2017-07-02/lp.log",
		//"./LHBLP101/2017-07-03/lp.log", "./LHBLP102/2017-07-03/lp.log",
		//"./LHBLP101/2017-07-04/lp.log", "./LHBLP102/2017-07-04/lp.log",
		//"./LHBLP101/2017-07-05/lp.log", "./LHBLP102/2017-07-05/lp.log",
		//"./LHBLP101/2017-07-06/lp.log", "./LHBLP102/2017-07-06/lp.log",
		//"./LHBLP101/2017-07-07/lp.log", "./LHBLP102/2017-07-07/lp.log",
		//"./LHBLP101/2017-07-08/lp.log", "./LHBLP102/2017-07-08/lp.log",
		//"./LHBLP101/2017-07-09/lp.log", "./LHBLP102/2017-07-09/lp.log",
		//"./LHBLP101/2017-07-10/lp.log", "./LHBLP102/2017-07-10/lp.log",
		//"./LHBLP101/2017-07-11/lp.log", "./LHBLP102/2017-07-11/lp.log",
		//"./LHBLP101/2017-07-12/lp.log", "./LHBLP102/2017-07-12/lp.log",
		//"./LHBLP101/2017-07-13/lp.log", "./LHBLP102/2017-07-13/lp.log",
		//"./LHBLP101/2017-07-14/lp.log", "./LHBLP102/2017-07-14/lp.log",
		//"./LHBLP101/2017-07-15/lp.log", "./LHBLP102/2017-07-15/lp.log",
		//"./LHBLP101/2017-07-16/lp.log", "./LHBLP102/2017-07-16/lp.log",
		//"./LHBLP101/2017-07-17/lp.log", "./LHBLP102/2017-07-17/lp.log",
		//"./LHBLP101/2017-07-18/lp.log", "./LHBLP102/2017-07-18/lp.log",
		//"./LHBLP101/2017-07-19/lp.log", "./LHBLP102/2017-07-19/lp.log",
		//"./LHBLP101/2017-07-20/lp.log", "./LHBLP102/2017-07-20/lp.log",
		//"./LHBLP101/2017-07-21/lp.log", "./LHBLP102/2017-07-21/lp.log",
		//"./LHBLP101/2017-07-22/lp.log", "./LHBLP102/2017-07-22/lp.log",
		//"./LHBLP101/2017-07-23/lp.log", "./LHBLP102/2017-07-23/lp.log",
		//"./LHBLP101/2017-07-24/lp.log", "./LHBLP102/2017-07-24/lp.log",
		//"./LHBLP101/2017-09-11/lp.log", "./LHBLP102/2017-09-11/lp.log",
		//"./LHBLP101/2017-09-14/lp.log", "./LHBLP102/2017-09-14/lp.log",
		//"./LHBLP101/2017-09-15/lp.log", "./LHBLP102/2017-09-15/lp.log",
		"./LHBLP101/2017-09-29/lp.log", "./LHBLP102/2017-09-29/lp.log",
		// QRC
		//"./QRC/303/2017-04-20/lp.log", "./QRC/304/2017-04-20/lp.log",
		//"./QRC/303/2017-04-21/lp.log", "./QRC/304/2017-04-21/lp.log",
		//"./QRC/303/2017-04-22/lp.log", "./QRC/304/2017-04-22/lp.log",
		//"./QRC/303/2017-04-23/lp.log", "./QRC/304/2017-04-23/lp.log",
		//"./QRC/303/2017-04-24/lp.log", "./QRC/304/2017-04-24/lp.log",
		//"./QRC/303/2017-04-25/lp.log", "./QRC/304/2017-04-25/lp.log",
		//"./QRC/303/2017-04-26/lp.log", "./QRC/304/2017-04-26/lp.log",
		//"./QRC/303/2017-04-27/lp.log", "./QRC/304/2017-04-27/lp.log",
		//"./QRC/303/2017-04-28/lp.log", "./QRC/304/2017-04-28/lp.log",
		//"./QRC/303/2017-04-29/lp.log", "./QRC/304/2017-04-29/lp.log",
		//"./QRC/303/2017-04-30/lp.log", "./QRC/304/2017-04-30/lp.log",
		//"./QRC/303/2017-05-01/lp.log", "./QRC/304/2017-05-01/lp.log",
		//"./QRC/303/2017-05-02/lp.log", "./QRC/304/2017-05-02/lp.log",
		//"./QRC/303/2017-05-03/lp.log", "./QRC/304/2017-05-03/lp.log",
		//"./QRC/303/2017-05-04/lp.log", "./QRC/304/2017-05-04/lp.log",
		//"./QRC/303/2017-05-05/lp.log", "./QRC/304/2017-05-05/lp.log",
		//"./QRC/303/2017-05-06/lp.log", "./QRC/304/2017-05-06/lp.log",
		//"./QRC/303/2017-05-07/lp.log", "./QRC/304/2017-05-07/lp.log",
		//"./QRC/303/2017-05-08/lp.log", "./QRC/304/2017-05-08/lp.log",
		//"./QRC/303/2017-05-09/lp.log", "./QRC/304/2017-05-09/lp.log",
		//"./QRC/303/2017-05-10/lp.log", "./QRC/304/2017-05-10/lp.log",
		//"./QRC/303/2017-05-11/lp.log", "./QRC/304/2017-05-11/lp.log",
		//"./QRC/303/2017-05-12/lp.log", "./QRC/304/2017-05-12/lp.log",
		//"./QRC/303/2017-05-13/lp.log", "./QRC/304/2017-05-13/lp.log",
		//"./QRC/303/2017-05-14/lp.log", "./QRC/304/2017-05-14/lp.log",
		//"./QRC/303/2017-05-15/lp.log", "./QRC/304/2017-05-15/lp.log",
		//"./QRC/303/2017-05-16/lp.log", "./QRC/304/2017-05-16/lp.log",
		//"./QRC/303/2017-05-17/lp.log", "./QRC/304/2017-05-17/lp.log",
		//"./QRC/303/2017-05-18/lp.log", "./QRC/304/2017-05-18/lp.log",
		//"./QRC/303/2017-05-19/lp.log", "./QRC/304/2017-05-19/lp.log",
		//"./QRC/303/2017-05-20/lp.log", "./QRC/304/2017-05-20/lp.log",
		//"./QRC/303/2017-05-21/lp.log", "./QRC/304/2017-05-21/lp.log",
		//"./QRC/303/2017-05-22/lp.log", "./QRC/304/2017-05-22/lp.log",
		//"./QRC/303/2017-05-23/lp.log", "./QRC/304/2017-05-23/lp.log",
		//"./QRC/303/2017-05-24/lp.log", "./QRC/304/2017-05-24/lp.log",
		//"./QRC/303/2017-05-25/lp.log", "./QRC/304/2017-05-25/lp.log",
		//"./QRC/303/2017-05-26/lp.log", "./QRC/304/2017-05-26/lp.log",
		//"./QRC/303/2017-05-27/lp.log", "./QRC/304/2017-05-27/lp.log",
		//"./QRC/303/2017-05-28/lp.log", "./QRC/304/2017-05-28/lp.log",
		//"./QRC/303/2017-05-29/lp.log", "./QRC/304/2017-05-29/lp.log",
		//"./QRC/303/2017-05-30/lp.log", "./QRC/304/2017-05-30/lp.log",
		//"./QRC/303/2017-05-31/lp.log", "./QRC/304/2017-05-31/lp.log",
		//"./QRC/303/2017-06-01/lp.log", "./QRC/304/2017-06-01/lp.log",
		//"./QRC/303/2017-06-02/lp.log", "./QRC/304/2017-06-02/lp.log",
		//"./QRC/303/2017-06-03/lp.log", "./QRC/304/2017-06-03/lp.log",
		//"./QRC/303/2017-06-04/lp.log", "./QRC/304/2017-06-04/lp.log",
		//"./QRC/303/2017-06-05/lp.log", "./QRC/304/2017-06-05/lp.log",
		//"./QRC/303/2017-06-06/lp.log", "./QRC/304/2017-06-06/lp.log",
		//"./QRC/303/2017-06-07/lp.log", "./QRC/304/2017-06-07/lp.log",
		//"./QRC/303/2017-06-08/lp.log", "./QRC/304/2017-06-08/lp.log",
		//"./QRC/303/2017-06-09/lp.log", "./QRC/304/2017-06-09/lp.log",
		//"./QRC/303/2017-06-10/lp.log", "./QRC/304/2017-06-10/lp.log",
		//"./QRC/303/2017-06-11/lp.log", "./QRC/304/2017-06-11/lp.log",
		//"./QRC/303/2017-06-12/lp.log", "./QRC/304/2017-06-12/lp.log",
		//"./QRC/303/2017-06-13/lp.log", "./QRC/304/2017-06-13/lp.log",
		//"./QRC/303/2017-06-14/lp.log", "./QRC/304/2017-06-14/lp.log",
		//"./QRC/303/2017-06-15/lp.log", "./QRC/304/2017-06-15/lp.log",
		//"./QRC/303/2017-06-16/lp.log", "./QRC/304/2017-06-16/lp.log",
		//"./QRC/303/2017-06-17/lp.log", "./QRC/304/2017-06-17/lp.log",
		//"./QRC/303/2017-06-18/lp.log", "./QRC/304/2017-06-18/lp.log",
		//"./QRC/303/2017-06-19/lp.log", "./QRC/304/2017-06-19/lp.log",
		//"./QRC/303/2017-06-20/lp.log", "./QRC/304/2017-06-20/lp.log",
		//"./QRC/303/2017-06-21/lp.log", "./QRC/304/2017-06-21/lp.log",
		//"./QRC/303/2017-06-22/lp.log", "./QRC/304/2017-06-22/lp.log",
		//"./QRC/303/2017-06-23/lp.log", "./QRC/304/2017-06-23/lp.log",
		//"./QRC/303/2017-06-24/lp.log", "./QRC/304/2017-06-24/lp.log",
		//"./QRC/303/2017-06-25/lp.log", "./QRC/304/2017-06-25/lp.log",
		//"./QRC/303/2017-06-26/lp.log", "./QRC/304/2017-06-26/lp.log",
		//"./QRC/303/2017-06-27/lp.log", "./QRC/304/2017-06-27/lp.log",
		//"./QRC/303/2017-06-28/lp.log", "./QRC/304/2017-06-28/lp.log",
		//"./QRC/303/2017-06-29/lp.log", "./QRC/304/2017-06-29/lp.log",
		//"./QRC/303/2017-06-30/lp.log", "./QRC/304/2017-06-30/lp.log",
		//"./QRC/303/2017-07-01/lp.log", "./QRC/304/2017-07-01/lp.log",
		//"./QRC/303/2017-07-02/lp.log", "./QRC/304/2017-07-02/lp.log",
		//"./QRC/303/2017-07-03/lp.log", "./QRC/304/2017-07-03/lp.log",
		// QC
		//"./QC/303/2017-05-03/lp.log", "./QC/304/2017-05-03/lp.log",
		//"./QC/303/2017-05-04/lp.log", "./QC/304/2017-05-04/lp.log",
		//"./QC/303/2017-05-05/lp.log", "./QC/304/2017-05-05/lp.log",
		//"./QC/303/2017-05-03/lp.log", "./QC/304/2017-05-06/lp.log",
		//"./QC/303/2017-05-07/lp.log", "./QC/304/2017-05-07/lp.log",
		//"./QC/303/2017-05-08/lp.log", "./QC/304/2017-05-08/lp.log",
		//"./QC/303/2017-05-09/lp.log", "./QC/304/2017-05-09/lp.log",
		//"./QC/303/2017-05-10/lp.log", "./QC/304/2017-05-10/lp.log",
		//"./QC/303/2017-05-11/lp.log", "./QC/304/2017-05-11/lp.log",
		//"./QC/303/2017-05-12/lp.log", "./QC/304/2017-05-12/lp.log",
		//"./QC/303/2017-05-13/lp.log", "./QC/304/2017-05-13/lp.log",
		//"./QC/303/2017-05-14/lp.log", "./QC/304/2017-05-14/lp.log",
		//"./QC/303/2017-05-15/lp.log", "./QC/304/2017-05-15/lp.log",
		//"./QC/303/2017-05-16/lp.log", "./QC/304/2017-05-16/lp.log",
		//"./QC/303/2017-05-17/lp.log", "./QC/304/2017-05-17/lp.log",
		//"./QC/303/2017-05-18/lp.log", "./QC/304/2017-05-18/lp.log",
		//"./QC/303/2017-05-19/lp.log", "./QC/304/2017-05-19/lp.log",
		//"./QC/303/2017-05-20/lp.log", "./QC/304/2017-05-20/lp.log",
		//"./QC/303/2017-05-21/lp.log", "./QC/304/2017-05-21/lp.log",
		//"./QC/303/2017-05-22/lp.log", "./QC/304/2017-05-22/lp.log",
		//"./QC/303/2017-05-23/lp.log", "./QC/304/2017-05-23/lp.log",
		//"./QC/303/2017-05-24/lp.log", "./QC/304/2017-05-24/lp.log",
		//"./QC/303/2017-05-25/lp.log", "./QC/304/2017-05-25/lp.log",
		//"./QC/303/2017-05-26/lp.log", "./QC/304/2017-05-26/lp.log",
		//"./QC/303/2017-05-31/lp.log", "./QC/304/2017-05-31/lp.log",
		//"./QC/303/2017-06-01/lp.log", "./QC/304/2017-06-01/lp.log",
		//"./QC/303/2017-06-02/lp.log", "./QC/304/2017-06-02/lp.log",
		//"./QC/303/2017-06-03/lp.log", "./QC/304/2017-06-03/lp.log",
		//"./QC/303/2017-06-04/lp.log", "./QC/304/2017-06-04/lp.log",
		//"./QC/303/2017-06-05/lp.log", "./QC/304/2017-06-05/lp.log",
		//"./QC/303/2017-06-06/lp.log", "./QC/304/2017-06-06/lp.log",
		//"./QC/303/2017-06-07/lp.log", "./QC/304/2017-06-07/lp.log",
		//"./QC/303/2017-06-08/lp.log", "./QC/304/2017-06-08/lp.log",
		//"./QC/303/2017-06-09/lp.log", "./QC/304/2017-06-09/lp.log",
		//"./QC/303/2017-06-10/lp.log", "./QC/304/2017-06-10/lp.log",
		//"./QC/303/2017-06-11/lp.log", "./QC/304/2017-06-11/lp.log",
		//"./QC/303/2017-06-12/lp.log", "./QC/304/2017-06-12/lp.log",
		//"./QC/303/2017-06-13/lp.log", "./QC/304/2017-06-13/lp.log",
		//"./QC/303/2017-06-14/lp.log", "./QC/304/2017-06-14/lp.log",
		//"./QC/303/2017-06-15/lp.log", "./QC/304/2017-06-15/lp.log",
		//"./QC/303/2017-06-16/lp.log", "./QC/304/2017-06-16/lp.log",
		//"./QC/303/2017-06-17/lp.log", "./QC/304/2017-06-17/lp.log",
		//"./QC/303/2017-06-18/lp.log", "./QC/304/2017-06-18/lp.log",
		//"./QC/303/2017-06-19/lp.log", "./QC/304/2017-06-19/lp.log",
		//"./QC/303/2017-06-20/lp.log", "./QC/304/2017-06-20/lp.log",
		//"./QC/303/2017-06-21/lp.log", "./QC/304/2017-06-21/lp.log",
		//"./QC/303/2017-06-22/lp.log", "./QC/304/2017-06-22/lp.log",
		//"./QC/303/2017-06-23/lp.log", "./QC/304/2017-06-23/lp.log",
		//"./QC/303/2017-06-24/lp.log", "./QC/304/2017-06-24/lp.log",
		//"./QC/303/2017-06-25/lp.log", "./QC/304/2017-06-25/lp.log",
		//"./QC/303/2017-06-26/lp.log", "./QC/304/2017-06-26/lp.log",
		//"./QC/303/2017-06-27/lp.log", "./QC/304/2017-06-27/lp.log",
		//"./QC/303/2017-06-28/lp.log", "./QC/304/2017-06-28/lp.log",
		//"./QC/303/2017-06-29/lp.log", "./QC/304/2017-06-29/lp.log",
		//"./QC/303/2017-06-30/lp.log", "./QC/304/2017-06-30/lp.log",
		//"./QC/303/2017-07-01/lp.log", "./QC/304/2017-07-01/lp.log",
		//"./QC/303/2017-07-02/lp.log", "./QC/304/2017-07-02/lp.log",
		//"./QC/303/2017-07-03/lp.log", "./QC/304/2017-07-03/lp.log",
		// IB
		//"./IB/web-101-2017-04-05.log", "./IB/web-102-2017-04-05.log",
		//"./IB/web-101-2017-04-06.log", "./IB/web-102-2017-04-06.log",
		//"./IB/web-101-2017-04-07.log", "./IB/web-102-2017-04-07.log",
		//"./IB/web-101-2017-04-08.log", "./IB/web-102-2017-04-08.log",
		//"./IB/web-101-2017-04-09.log", "./IB/web-102-2017-04-09.log",
		//"./IB/web-101-2017-04-10.log", "./IB/web-102-2017-04-10.log",
		//"./IB/web-101-2017-04-11.log", "./IB/web-102-2017-04-11.log",
		//"./IB/web-101-2017-04-12.log", "./IB/web-102-2017-04-12.log",
		//"./IB/web-101-2017-04-13.log", "./IB/web-102-2017-04-13.log",
		//"./IB/web-101-2017-04-14.log", "./IB/web-102-2017-04-14.log",
		//"./IB/web-101-2017-04-15.log", "./IB/web-102-2017-04-15.log",
		//"./IB/web-101-2017-04-16.log", "./IB/web-102-2017-04-16.log",
		//"./IB/web-101-2017-04-17.log", "./IB/web-102-2017-04-17.log",
		//"./IB/web-101-2017-04-18.log", "./IB/web-102-2017-04-18.log",
		//"./IB/web-101-2017-04-19.log", "./IB/web-102-2017-04-19.log",
		//"./IB/web-101-2017-04-20.log", "./IB/web-102-2017-04-20.log",
	}
	//
	i := 0
	var wg sync.WaitGroup
	wg.Add(len(filenameList))
	//
	for _, file := range filenameList {
		//parseFile(path, file)
		go worker(path+filenameList[i], i+1, verbose, &wg)
		i++
		fmt.Printf("Processing file: %d - %s\r\n", i, file)
	}
	wg.Wait()
	//
	//endTime := time.Now()
	fmt.Println("Total processing time used: ", time.Since(startTime))
	//fmt.Printf("%T %T", 2.0, 2.0<<0)
	fmt.Println("Program ended...")
}

//
//	go routines
//
func worker(fileName string, id int, verbose bool, wgp *sync.WaitGroup) {

	startTime1 := time.Now()
	//fmt.Println("Starting go routine: %d", id)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: %s, %s", fileName, err)
		return
	}
	// deferred call to Close() at the end of current function
	defer file.Close()
	//
	//reader := csv.NewReader(file)
	//Configure reader options Ref http://golang.org/src/pkg/encoding/csv/reader.go?s=#L81
	//reader.Comma = ';'          //field delimiter
	//reader.Comment = '#'        //Comment character
	//reader.FieldsPerRecord = -1 //Number of records per record. Set to Negative value for variable
	//reader.TrimLeadingSpace = true
	//
	lineCount := 0
	errorLines := 0
	//
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 128*2048)
	scanner.Buffer(buf, 4096*4096)
	//scanner := bufio.Reader.ReadLine(line, err)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if strings.Contains(scanner.Text(), " ERROR ") {
			if (!strings.Contains(scanner.Text(), "updateDocumentStatus")) && (!strings.Contains(scanner.Text(), "Signing.Providers") && (!strings.Contains(scanner.Text(), "UpdateDocumentStatusException"))) {
				errorLines++
				if verbose {
					fmt.Printf("File: %s - Row %d: %s\r\n", fileName, lineCount, scanner.Text())
				}
			}
		}
		lineCount++
	}
	//
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}

	//result = fmt.Sprintf("go routine %d - File: %s, # of error lines: %d, processing time: %s \r\n", id, fileName, errorLines, time.Since(startTime1))
	fmt.Printf("go routine %d - File: %s, # of lines: %d, # of error lines: %d, processing time: %s \r\n", id, fileName, lineCount, errorLines, time.Since(startTime1))
	//
	wgp.Done()
}
