cmd /c start powershell -Command {go run .\logic\orchestrator-1.go}
Write-Output "Orchestrator 1 Started at [::]:9000"
cmd /c start powershell -Command {go run .\logic\orchestrator-2.go}
Write-Output "Orchestrator 2 Started at [::]:9001"
cmd /c start powershell -Command {go run .\datamock\datamock.go}
Write-Output "Mock Data Service Started at [::]:10000"