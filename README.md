# hgrok - Introspected tunnels to localhost ([homepage](https://beta.hasura.io))
### ”I want to expose a local hasura project behind a NAT or firewall to the internet.”
![](https://dev.beta.hasura.io/rstatic/dist/d381a3c91a318589ac4b383222295d17.png)

## What is hgrok?
hgrok is a reverse proxy that creates a secure tunnel from a public endpoint to a locally running web service.
hgrok captures and analyzes all traffic over the tunnel for later inspection and replay.

## What can I do with hgrok?
- Expose any http service behind a NAT or firewall to the internet on a subdomain of hasura.me
- Inspect all http requests/responses that are transmitted over the tunnel
- Replay any request that was transmitted over the tunnel


## What is hgrok useful for?
- Temporarily sharing your hasura project that is only running on your development machine
- Demoing an app at a hackathon without deploying
- Developing any services which consume webhooks (HTTP callbacks) by allowing you to replay those requests
- Debugging and understanding any web service by inspecting the HTTP traffic
- Running networked services on machines that are firewalled off from the internet


## Usage
- Register an account in [hasura](https://beta.hasura.io)
- Configure your hasura.me subdomain in https://dev.beta.hasura.io/local-development
- Get your beta.hasura.io api token from https://dev.beta.hasura.io/settings
- Download hgrok binary based on your platform from [to-be-added](#)
- Execute

```
./ngrok -authtoken="<api token from beta.hasura.io>" -projectname="<projectName>" -proto=http 80 
```
