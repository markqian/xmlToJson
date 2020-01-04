# Senior Backend Engineer Takehome Project
- Your goal is to implement a simple HTTP server which exposes a single route to the user
- This route takes in an XML list and processes it *concurrently* ("processing" in this context means running `strings.ToUpper()` on the `data` field of each item), then returns the processed list as JSON
- Write at least 2 test cases for your server

## Some Clarifications
- Your server must be written in Go, and it should only have one route which takes in an XML list, performs processing, and returns JSON
- Your server should use Go routines and channels in the route to implement concurrent processing
	+ Although each item is being processed concurrently in your route, that does not mean that your response to the user should return early. Your response should wait until *all* items are processed, and return the complete processed list as JSON to the user. To clarify further; you are only exposing one route that does all processing, blocking the response to the user until the request has finished its concurrent processing. The user *should not* have to do things like call a separate route to poll statuses or get the results of their request.
- Assume that Order IDs are strings as in the example below
- You don't need a persistence layer in your solution, this should all be  done in memory
- Your code should include comments where you see fit. Assume the reader is a co-worker who understands Go but has vague understanding of what you're working on

## Things You'll Need
- Golang `strings.ToUpper` documentation (https://golang.org/pkg/strings/#ToUpper)
- Optional: [Golang Documentation](`https://golang.org`)
	+ You may need to install Go
- Optional: [Go Playground](`play.golang.org`)
	+ Go REPL, useful for quickly testing some Go code in the browser
- The API Documentation below

## API Documentation
We have a single resource, an `Order`, which has the following fields:
	-`id` (string)
	-`data` (string)
	-`createdAt` (string)
	-`updatedAt` (string)

### Full Example

#### Input Example (`input_example.xml`):
```
<orderList>	
	<order>
		<id>aeffb38f-a1a0-48e7-b7a8-2621a2678534</id>
		<data>first_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>beffb38f-b1a0-58e7-c7a8-3621a2678534</id>
		<data>second_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
</orderList>
```

#### Example Request:

	curl -X POST <YOUR_JSON_SERVER>/process -d '`cat ./input_example.xml`'

#### Example Response:
```
{
	{
		"id":"aeffb38f-a1a0-48e7-b7a8-2621a2678534",
		"data":"FIRST_ORDER_DATA",
		"createdAt":"0001-01-01T00:00:00Z",
		"updatedAt":"0001-01-01T00:00:00Z"
	},
	{
		"id":"beffb38f-b1a0-58e7-c7a8-3621a2678534",
		"data":"SECOND_ORDER_DATA",
		"createdAt":"0001-01-01T00:00:00Z",
		"updatedAt":"0001-01-01T00:00:00Z"
	}
}
```

Notice how *only* the `data` field is upper-cased, and not the other fields of the model. Your server should do the same.
 
## Submitting Your Code
- Your JSON server code should be uploaded to your public Git repository (Github, Gitlab, Bitbucket, etc) and a link to that repository should be sent to `thanasi@authenticiti.io` 
- Your repo should contain a `README.md` file describing how to run your code locally

- **Bonus Points**: In addition to pushing to Github, deploy your server to any cloud provider (ie AWS, Azure, GCP, etc) and provide the IP address and port to `thanasi@authenticiti.io` to interact with your server

## Additional Information
If you have any questions, don't hesitate to email `thanasi@authenticiti.io`