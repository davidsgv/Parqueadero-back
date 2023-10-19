FROM golang:1.20.2-bullseye

# install git
# RUN apt update && apt upgrade
# RUN apk add git

#Herramientas de desarrollo
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

#config git user and email
# RUN git config --global user.name "${git_user}"
# RUN git config --global user.email "${git_mail}"


# EXPOSE 80