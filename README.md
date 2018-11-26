# gonpe

gonpe is video management server written by golang.

# Setup

installation

```
# language, library

# go, protobuf
brew install go (v1.11+)
brew install protobuf (v3+)

# node, npm, yarn
brew install nodebrew
nodebrew setup
export PATH=$HOME/.nodebrew/current/bin:$PATH
nodebrew install <v10+>
nodebrew use <installed_version>
npm update -g npm
npm i -g yarn

# Google Cloud SDK
curl https://sdk.cloud.google.com | bash
gcloud init
gcloud components install app-engine-go
gcloud config set project <your_project>

# envvars
export GO111MODULES=on
```

# Install dependency

`make dep`

# Serve

`make run`

# Gen proto

`make gen`

# Clean

`make clean`
