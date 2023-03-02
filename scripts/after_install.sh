cd /home/ec2-user/chargepass/
go mod verify
go mod tidy
go version
cp /home/ec2-user/.env .env
sudo go build -o main main.go
ls
sudo systemctl daemon-reload && sudo systemctl enable chargepass && sudo systemctl start chargepass
sudo systemctl restart chargepass