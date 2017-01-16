### Running

To run the example, simply use:

```
$ go run server.go
```

Then browse to `http://localhost:8080`. As-is, clicking the "Edit/Cancel" button
throws a `HierarchyRequestError` in the javascript console; it can be fixed by
modifying `index.html` to load `working.js` instead.
