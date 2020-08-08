/*

Here is an example suite that illustrates an ability to apply contract testing of
quality assessment strategy. The contact implements a test cases as function of the form:

  func TestAbc() gurl.Arrow

where Abc is a unique name of test case. Each case declares cause-and-effect:

↣ "Given" specifies the communication context and the known state of the expected behavior;

↣ "When" executes key actions about the interaction with remote component;

↣ "Then" observes output of remote component, validates its correctness and outputs results.

The service evaluates suites and its test cases sequentially one after another.

*/

// Package sample is a standard Golang declaration. It groups set of logically related contracts.
package sample

/*

Standard Golang import declaration.

However, assay.it restricts usage of some package.
Please check doc.assay.it for details of allowed packages.
We are constantly looking for your feedback, please open an issue to us.
*/
import (

	//
	// gurl library is a class of High Order Component which can do http requests
	// with few interesting property such as composition and laziness.
	// It implements a human-friendly syntax of HTTP communication and
	// Behavior as a Code paradigm. It connect cause-and-effect with the networking
	// primitives. Usage of gurl is a preferred approach for networking I/O.
	"github.com/fogfish/gurl"
	ƒ "github.com/fogfish/gurl/http/recv"
	ø "github.com/fogfish/gurl/http/send"
)

/*

TestOk makes HTTP request and checks the quality of the response.
This is simple extample just validates that service is alive and
responding with HTML document.

Let's look in depth on the anatomy of the contract
*/
func TestOk() gurl.Arrow {
	/*
		The case is developed with Go URL library (gurl). This library defines
		a rich techniques to hide the networking complexity using higher-order-functions
		and its compositions. See https://assay.it/doc/core for details about api

		gurl.HTTP lifts primitive protocol functions to higher-order suite.
	*/
	return gurl.HTTP(
		// module ø (gurl/http/send) defines function to declare HTTP request.
		// See the https://assay.it/doc/core for details about module ø.

		// declares HTTP method and destination URL
		ø.GET("https://assay.it"),

		// module ƒ (gurl/http/recv) defines function to validate correctness of HTTP protocol.
		// Each ƒ constrain might terminate execution of consequent ƒ's if it expectation fails.
		// See the https://assay.it/doc/core for details about module ƒ.

		// requires HTTP Status Code to be 200 OK
		ƒ.Code(gurl.StatusCodeOK),

		// requires response content-type to be text/html
		ƒ.Header("Content-Type").Is("text/html"),
	)
}
