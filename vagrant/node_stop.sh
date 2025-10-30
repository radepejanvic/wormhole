echo "Stopping node: $(hostname)..."
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

# Stop star
STAR_PID=$(cat /etc/c12s/star.pid)

if ps -p $STAR_PID -o comm= | grep -q star; then
    echo "Killing star (PID $STAR_PID)..."
    kill $STAR_PID
    docker compose -f node.yml down -v
    echo "Star stopped."
else
    echo "PID $STAR_PID is not star, skipping."
fi

