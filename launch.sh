echo "Orchestrator 1 Started at [::]:9000"
echo "Orchestrator 2 Started at [::]:9001"
echo "Mock Data Service Started at [::]:10000"
gnome-terminal --tab --title="Orchestrator 1"   --command "sudo go run ./logic/orchestrator-1.go" --tab --title="Orchestrator 2" --command "sudo go run ./logic/orchestrator-2.go" --tab --title="Mock Data Service"  --command "sudo go run ./datamock/datamock.go"