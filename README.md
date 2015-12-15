# Cisco Spark Go Client

Testing library for running commands in Cisco Spark using Go.

## To use this: 

### 1.  Download code:

git clone https://github.com/vallard/gospark.git

### 2.  Define an environment variable

Go to the [https://developer.ciscospark.com/](Cisco Spark Developers homepage) and click on your profile 
picture.  From there you can copy your Auth Token.  In your ```.profile``` or ```.bash_profile``` add the line:

```
export SPARK_AUTH_TOKEN=2efal0msmeMWEzZsd....
```
Then re-export your environment or open a new window. 

### 3.  Run the code: 

```
go run main.go | json_pp
```

Look at the ```sparkClient/sparkClient.go``` file to see the available commands.

Check out [https://learninglabs.cisco.com/lab/collab-spark-message/step/4](Cisco DevNet) for more tutorials

