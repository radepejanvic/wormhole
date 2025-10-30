echo "Starting node: $(hostname)..."
cd /home/vagrant/tools

# Export configurable env vars
OS_NAME="$(uname -s)"
if [ "$OS_NAME" = "Darwin" ]; then
    echo "Running on MacOS, trying to use xargs"
    export $(grep -v '^#' node.env | gxargs -d '\n')
else
    echo "Running on linux, using xargs"
    export $(grep -v '^#' node.env | xargs -d '\n')
fi

sudo rm -rf /etc/c12s
sudo mkdir -p /etc/c12s
sudo chmod 0777 /etc/c12s

docker compose -f node.yml build
docker compose -f node.yml up -d

# Build and run star
cd ../star/cmd
sudo chown $USER:$USER .
chmod u+w .

echo "Building star binary..."
go build -buildvcs=false -o star

echo "Running ./star in $(pwd)"
nohup ./star > /vagrant/star.log 2>&1 &
STAR_PID=$!
echo $STAR_PID > /etc/c12s/star.pid

echo "Started star (PID $STAR_PID). Logs are in /vagrant/star.log"
