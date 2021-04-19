README

1. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened. Put this in a section in
your README.md file.

- It is very likely that what happened was a race condition while updating the stock & creating the order, that was not prevented (by locking or anything). The race condition caused the stock to have negative value, and the invalid order to be created.

2. Based on your analysis, propose a solution that will prevent the incidents from occurring again. Put this in a section in your README.md
file

- To prevent the incident, do optimistic locking while updating the stock & creating the order. Return an error if a race condition happened, instead of updating the stock and creating an order.

To demonstrate, run the backend with `go run main.go`, and then run the functional test inside `functional test.zip`
 
