# hgrok - Introspected tunnels to localhost ([homepage](https://beta.hasura.io))
### ”I want to expose a local hasura project behind a NAT or firewall to the internet.”
![](https://raw.githubusercontent.com/hasura/ngrok/master/_docs/hgrok.png)

## What is hgrok?
hgrok is a fork of ngrok for developers to expose their local hasura projects to the Internet via a secure tunnel.

## What can I do with hgrok?
- Expose any http service behind a NAT or firewall to the internet on a subdomain of hasura.me. If you choose a subdomain called `john`, requests to `*.john.hasura.me` will come through the tunnel to your server.

## What is hgrok useful for?
- Temporarily sharing your hasura project that is only running on your development machine
- Demoing an app at a hackathon without deploying to a cloud provider


## Usage
- Register an account at [hasura](https://beta.hasura.io)
- Configure your hasura.me subdomain at https://dev.beta.hasura.io/local-development
- Get your hasura.io api token from https://dev.beta.hasura.io/settings
- Download the hgrok binary for your OS from [to-be-added](#)

- Run the following command if you are running a server on minikube (mac/linux):

```
./hgrok -authtoken="<api-token>" -subdomain="<subdomain>" -proto=http $(minikube ip):80 
```

- Get your minikube ip, and substitute \<minikube-ip> in the command below:

```
./hgrok -authtoken="<api-token>" -subdomain="<subdomain>" -proto=http <minikube-ip>:80 
```

- Run the following command if you are running a server on localhost:80

```
./hgrok -authtoken="<api-token>" -subdomain="<subdomain>" -proto=http 80 
```

## Thank you ngrok!
hgrok is a fork of the awesome [ngrok](https://github.com/inconshreveable/ngrok). That means you can do pretty much everything you do with ngrok, wth hgrok too!
