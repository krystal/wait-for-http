# Wait for HTTP

This tiny process will just wait for HTTP to respond on the given host and port before exiting. To be used in an init container when you want to wait for something to be online.

## Usage

```bash
# Wait 30 times with a 5 second sleep between each check.
$ wait-for-http http://localhost:3000

# Check 10 times before exiting with an error code (default sleep 5 seconds)
$ wait-for-http http://localhost:3000 --quantity 10

# Check 20 times with a check every 2 seconds
$ wait-for-http http://localhost:3000 --quantity 20 --sleep 2

# Check for 200 and 404 statuses
$ wait-for-http http://localhost:3000 --statuses 200,404

# Set the timeout
$ wait-for-http http://localhost:3000 --timeout 5
```
