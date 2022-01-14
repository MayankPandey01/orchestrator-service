
 # Orchestrator Service

 The goal of this program is to build an orchestrator service that would read any request it receives and forwards it to other orchestrator services and data services.
The final working of the Orchestrator looks like this

- Request
```
client ---RPC--> orchestrator_1(:9000) ---RPC--> orchestrator_2(:9001) ---RPC--> mock_data_service(:10000)
```
- Response
```
mock_data_service(:10000) ---RPC--> orchestrator_2(:9001) ---RPC--> orchestrator_1(:9000) ---RPC--> client   
```

# Installation and Usage

- Use git to clone this repo 
	`git clone https://github.com/MayankPandey01/orchestrator-service`
 
 ### On Windows
 - Use Powershell to Launch all the Servers
 - Start a Powershell in the root of the Cloned Repo
 - Use `powershell launch.ps1`
 - This will Launch All the Servers [ Plz Wait A Few Seconds to Fully Load all the modules and Start the Server]
 - You will Be Prompted By Firewall to Open Ports `[9000,9001,10000]`
 - Now in a new Powershell window run the Client by using `go run .\client\client.go` 
 - You Can See the Output/Error depending on the input provided in the `client.go` file.


### On Linux
 - Use Terminal to Launch all the Servers
 - Start a bash Terminal in the root of the Cloned Repo
 - Use `sudo bash launch.ps1`
 - This will Launch All the Servers [ Plz Wait A Few Seconds to Fully Load all the modules and Start the Server]
 - Now in a new bash terminal run the Client by using `sudo go run .\client\client.go` 
 - You Can See the Output/Error depending on the input provided in the `client.go` file.

## File Structure

 ### `.\datamock\datamock.go`
 - It is the main service that receives a request and based on the request, returns a response based on simple logic.
 - Running as a separate microservice on port `10000`

 ### `.\logic\orchestrator-1.go`
 - This is the first Orchestrator Running on Port 9000
 -  It has a method `GetUserByName`  which Forwards the Requests to `orchestrator-2.go`

 ### `.\logic\orchestrator-2.go`
 - This is the Second Orchestrator Running on Port 9001
 -  It has a method `GetUser`  which recives request from `orchestrator-1.go` and Forwards the Requests to `datamock.go`

### `.\client\client.go`
- This is where the client logic Resides.
- An Input is passed to the `Orchestrator-1` and from there requested are passed to services running on different ports [as described above]
- The input is Hardcoded in this file, you can change this according to your need.

### `.\server\server.go`
- This is a different server, this was made in the initial steps of making this Orchestrator Service.
- It runs on POrt 50051 and returns a Hardcoded Response 
- It's not part of the actual working of the Program.

### `launch.ps1 and launch.sh`
- These files are made to automate the launching of the servers on Windows and Linux. It saves the effort to manually launch all the 3 servers

### Things to note

- This program was tested on both `Windows` and `Linux` and It works as expected
- If an error arises for `context timeout` then try to wait for a few seconds and relaunching the `client`, This issue only arises if the servers are not fully Online and the client is trying to make the Request
- Go Version Used: `1.14 - 1.17`
