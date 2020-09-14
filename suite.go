/*

Here is an example of Behavior as a Code suite for testing of Serverless,
Microservices and other systems that rely on interface syntaxes and its
behaviors. The behavior is implemented as function of the form:

  func TestAbc() assay.Arrow

where Abc is a unique name of test case. Each case declares cause-and-effect:

↣ "Given" specifies the communication context and the known state of the
expected behavior;

↣ "When" executes key actions about the interaction with remote component;

↣ "Then" observes output of remote component, validates its correctness and
outputs results.

assay.it evaluates suites and its functions sequentially one after another.

*/

/*

Package sample is a standard Golang declaration. It groups set of logically
related contracts.
*/
package sample

/*

Standard Golang import declaration.

However, assay.it restricts usage of some package.
Please check assay.it/doc/core for details of allowed packages.
We are constantly looking for your feedback, please open an issue to us.
*/
import (

	//
	// assay-it/sdk-go introduces pure functional and typesafe syntax to implement
	// HTTP communication and Behavior as a Code suites.
	"github.com/assay-it/sdk-go/assay"
	"github.com/assay-it/sdk-go/http"
	ƒ "github.com/assay-it/sdk-go/http/recv"
	ø "github.com/assay-it/sdk-go/http/send"
)

/*

TestOk makes HTTP request and checks the quality of the response. This is simple
extample just validates that service is alive and responding with HTML document.

Let's look in depth on the anatomy
*/
func TestOk() assay.Arrow {
	/*
		SDK defines a rich techniques to hide the networking complexity using
		higher-order-functions and its compositions. See https://assay.it/doc/core
		for details about api.

		http.Join lifts primitive protocol "modifiers" to higher-order suite.
	*/
	return http.Join(
		// module ø (http/send) defines function to declare HTTP request.
		// See the https://assay.it/doc/core for details about module ø.

		// declares HTTP method and destination URL
		ø.GET("https://assay.it"),

		// module ƒ (http/recv) defines function to validate correctness of HTTP
		// response. Each ƒ constrain might terminate execution of consequent ƒ's
		// if it expectation fails. See the https://assay.it/doc/core for details
		// about module ƒ.

		// requires HTTP Status Code to be 200 OK
		ƒ.Code(http.StatusCodeOK),

		// requires response content-type to be text/html
		ƒ.Header("Content-Type").Is("text/html"),
	)
}
