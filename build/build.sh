#/bin/bash
basepath=$(cd `dirname $0`; pwd)
ver=${1-"v1.0.0"}
rm -f cfgedit  
cd ..
CGO_ENABLED=0 go build  -v -a -ldflags '-extldflags "-static"' -o $basepath/cfgedit
mkdir -p $basepath/ca
cp -rf ca/ca.crt ca/ca.key $basepath/ca/
cd $basepath
sudo docker build -t editcfg:$ver .