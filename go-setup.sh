
#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail

GO_VERSION=${GO_VERSION:-"go1.20.2"}
GO_DOWNLOAD_URL=${GO_DOWNLOAD_URL:-"https://go.dev/dl/go1.20.2.linux-amd64.tar.gz"}
GO_BIN="/urs/local"

function install_go () {
    echo "Verify if Go is already installed"
    go version
    exitcode="${?}"

    if [ "$exitcode" -ne 0 ];  then
        echo "Go is not installed"
    else
        echo "Go is aready installed..."
        return 0
    fi 

}

if [ $1 == "install" ];  then
    install_go
elif [ $1 == "uninstall" ]; then 
    echo "uninstall go"
else
    echo "Invalid arguments provided. Usage: $0 [option..] {install {project} | uninstall}"
fi 
