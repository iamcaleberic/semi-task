minikube config set memory 4092
minikube start -p semi-task --nodes 2
minikube addons enable ingress -p semi-task
echo $(minikube ip -p semi-task) semi-task.local | sudo tee -a /etc/hosts
cat /etc/hosts

cd terraform
terraform init
terraform apply -auto-approve
