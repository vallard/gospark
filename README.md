# Cisco Spark Go Client

Testing library for running commands in Cisco Spark using Go.


## Spark integration service bot

Build it for testing: 
```
docker run --name gospark -v `pwd`:/go/src/github.com/vallard/gospark/ -w /go/src/github.com/vallard/gospark/ -p 8089:8089 --rm -it golang
go get
go run main.go
```

Testing

```
curl -X POST -H 'Content-Type: application/json' -d "{\"test\" : \"this\" }" localhost:8081/v1/commit/
```

## Spark Fortune

### 1.  Download code:

git clone https://github.com/vallard/gospark.git

### 2.  Define an environment variable

Go to the [Cisco Spark Developers homepage](https://developer.ciscospark.com)  and click on your profile 
picture.  From there you can copy your Auth Token.  In your ```.profile``` or ```.bash_profile``` add the line:

```
export SPARK_AUTH_TOKEN=2efal0msmeMWEzZsd....
```
Then re-export your environment or open a new window. 

### 3.  Run the code: 

```
go run main.go <email file> fortunes.txt
```

You'll need to supply a list of emails of people.  Just do one email per line.

Look at the ```sparkClient/sparkClient.go``` file to see the available commands.

Check out [Cisco DevNet](https://learninglabs.cisco.com/lab/collab-spark-message/step/4) for more tutorials

### Fortunes from: 
https://joshmadison.com/2008/04/20/fortune-cookie-fortunes/
