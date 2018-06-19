 # Overview
 
 In this  example / tutorial, I'm going to show you how to create
 a simple logstash filter pipeline. This tutorial targets logstash 5.
 
 In this tutorial, I am going to focus on the logstash filter configuration
 and the corresponding logstash-filter-verifier test. I'll try to be clear about
 my assumptions about the logstash inputs, but configuring filebeat and
 logstash are out of scope of this tutorial. That should be fine though, 
 because in my opinion, getting filebeat to throw some logs at logstash is
 the easy part about running an elastic stack.
 
 ## Setup
 
 In our setup, we're keeping our logstash configuration and the tests in a
 single logstash repository. The logstash config files are kept in a
 directory `config`, and the tests are kept in a directory `tests`. This makes
 it easy to keep the tests and the config in sync and to run the tests whenever
 the configuration is changed. 
 
 Furthermore, `logstash-filter-verifier` - and logstash testing in general has
 a problem: logstash-filter-verifier does not cooperate well with input or output
 configurations in logstash besides inputs and outputs provided on the fly. 
 At best this causes the test to hang. Even worse, missing secrets for e.g. the 
 filebeat input causes logstash to not start due to missing certificates and 
 everything becomes messy at a rapid pace.
 
 To circumvent this, we define all `input {}` and all `output {}` parts of the 
 configuration in a single file, `config/001_io.conf` and filters are defined
 in other files. Given this, we can delete the file `config/001_io.conf` just
 before we run `logstash-filter-verifier` and use `git checkout` later to 
 recover the file. With this, we can focus on processing logs. 
 
 ## Syslog
 
 Syslogs are usually a good start to start aggregating logs. Syslogs contain
 a lot of interesting information regarding security. And if you aggregate 
 syslogs, compatible tools can just log to the syslog and the aggregation 
 automatically works, or you can use tools like `logger -t mylog` to send
 logs to the syslog and thus, your elastic stack. I will assume you have setup
 a filebeat prospector or a similar forwarder to send logstash syslogs with
 a type of `syslog`. 
 
 From here, a good start is to just grab a log line. 
 
 - TODO: continue.
