<h1><div><img alt="logo" src="./images/MQTTHook.svg" style="width: 60px; height: 60px;">MQTThook<div></h1>
is a simple Program that converts mqtt requests to Webhooks requests that <br>
are then used as a Trigger in another programm


### Configuration
To configure MqttHook create a **config.yaml** file with the following options:

```yaml

broker:                      #The Mqtt Broker
  host: "127.0.0.1"          #The Host domain
  port: 1883                 #The Host port
  username: "user"           #The Mqtt Username
  password: "s3cr3t"         #The Password
  topic: "hello"             #The Topic to subsribe to

hook:  
  host: "https://your-fancy-domain.de?token="         #The Domain used for the webhook
  method: "GET"                                       #The Method used for the webhook request
  #payload: string /  token                           #The Payload used that is apended to the host               
```
>[!WARNING]
>The Payload is used as an apendix for the the host an is used as a token
>for the host domain to authenticate

### Deployment
To deploy run this command
```shell
docker build -t mqtthook -f Dockerfile .
```
