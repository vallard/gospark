# Cisco Spark Go Client

Testing library for running commands in Cisco Spark using Go.

## To use this: 

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
