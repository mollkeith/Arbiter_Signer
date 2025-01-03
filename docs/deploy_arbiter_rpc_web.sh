#!/bin/bash

echo_error()
{
    echo -e "\033[1;31mERROR:\033[0m $1"
}

echo_info()
{
    echo -e "\033[1;34mINFO:\033[0m $1"
}

echo_info_green() {
    echo -e "\033[1;32mINFO:\033[0m $1"
}

echo_info_red() {
    echo -e "\033[1;31mERROR:\033[0m $1"
}

#
# record update log
#
update_log()
{
  if [ ! -f $SCRIPT_PATH/update.log ]; then
    echo_error "$SCRIPT_PATH/update.log is not exist"
    echo_info "Create update.log"
    touch $SCRIPT_PATH/update.log
  fi

  local time=$(date "+%Y-%m-%d %H:%M:%S")
  echo_info "$time">>$SCRIPT_PATH/update.log
  echo_info "==========">>$SCRIPT_PATH/update.log
  echo_info "deploy arbiter">>$SCRIPT_PATH/update.log
  echo_info "">>$SCRIPT_PATH/update.log
  if [ $1 == "succeeded" ]; then
      echo_info "$time deploy arbiter succeeded!"
  else
      echo_error "$time deploy arbiter failed!"
  fi
  echo_info "Please check update log via command: cat $SCRIPT_PATH/update.log"
}

#
# check rpc status
#
check_rpc_status()
{
 PROCESS_NAME="arbiter_rpc"
 if pgrep -x "$PROCESS_NAME" > /dev/null
 then
     echo_info "$PROCESS_NAME is running."
     echo_info_green "Succeed!"
 else
     echo_info "$PROCESS_NAME is not running."
     echo_info_red "Failed!"
 fi
}

#
# check web status
#
check_rpc_status()
{
 PROCESS_NAME="arbiter_web"
 if pgrep -x "$PROCESS_NAME" > /dev/null
 then
     echo_info "$PROCESS_NAME is running."
     echo_info_green "Succeed!"
 else
     echo_info "$PROCESS_NAME is not running."
     echo_info_red "Failed!"
 fi
}

#
# deploy rpc
#
deploy_arbiter_rpc()
{
	echo_info $SCRIPT_PATH

  if [ ! -d "$SCRIPT_PATH" ]; then
		mkdir -p $SCRIPT_PATH/data/logs
		mkdir -p $SCRIPT_PATH/rpc
	fi
	cd $SCRIPT_PATH/rpc

	#prepare config.yaml
	wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/rpc_conf.tgz
	tar xf rpc_conf.tgz

	#prepare arbiter rpc
  if [ "$(uname -m)" == "armv6l" ] || [ "$(uname -m)" == "armv7l" ] || [ "$(uname -m)" == "aarch64" ]; then
    echo "The current system architecture is ARM"
    echo_info "Downloading loan arbiter rpc..."
    wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-rpc-linux-arm64.tgz
    tar xf loan-arbiter-rpc-linux-arm64.tgz
    echo_info "Replacing arbiter rpc.."
    cp -v loan-arbiter-rpc-linux-arm64/arbiter_rpc ~/loan_arbiter/rpc/
    echo_info "Starting arbtier rpc..."
    ./arbiter_rpc --gf.gcfg.file=config.yaml  > $SCRIPT_PATH/data/logs/arbiter_rpc.log 2>&1 &
  else
    echo "The current system architecture is x86"
    echo_info "Downloading loan arbiter rpc..."
    wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-rpc-linux-x86_64.tgz
    tar xf loan-arbiter-rpc-linux-x86_64.tgz
    echo_info "Replacing arbiter rpc.."
    cp -v loan-arbiter-rpc-linux-x86_64/arbiter_rpc ~/loan_arbiter/rpc/
    echo_info "Starting arbtier rpc..."
    ./arbiter_rpc --gf.gcfg.file=config.yaml > $SCRIPT_PATH/data/logs/arbiter_rpc.log 2>&1 &
  fi

  check_status
  echo_info "Please check arbiter log via command: cat $SCRIPT_PATH/data/logs/arbiter.log"
}


#
# deploy web
#
deploy_arbiter_web()
{
	echo_info $SCRIPT_PATH

  if [ ! -d "$SCRIPT_PATH" ]; then
		mkdir -p $SCRIPT_PATH/data/logs
		mkdir -p $SCRIPT_PATH/web
	fi
	cd $SCRIPT_PATH/web

	#prepare config.yaml
	wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/web_conf.tgz
	tar xf web_conf.tgz

	#prepare arbiter web
  if [ "$(uname -m)" == "armv6l" ] || [ "$(uname -m)" == "armv7l" ] || [ "$(uname -m)" == "aarch64" ]; then
    echo "The current system architecture is ARM"
    echo_info "Downloading loan arbiter web..."
    wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-web-linux-arm64.tgz
    tar xf loan-arbiter-web-linux-arm64.tgz
    echo_info "Replacing arbiter web.."
    cp -v loan-arbiter-web-linux-arm64/arbiter_web ~/loan_arbiter/web/
    echo_info "Starting arbtier web..."
    ./arbiter_web --gf.gcfg.file=config.yaml  > $SCRIPT_PATH/data/logs/arbiter_web.log 2>&1 &
  else
    echo "The current system architecture is x86"
    echo_info "Downloading loan arbiter web..."
    wget https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-web-linux-x86_64.tgz
    tar xf loan-arbiter-web-linux-x86_64.tgz
    echo_info "Replacing arbiter web.."
    cp -v loan-arbiter-web-linux-x86_64/arbiter_web ~/loan_arbiter/web/
    echo_info "Starting arbtier web..."
    ./arbiter_web --gf.gcfg.file=config.yaml > $SCRIPT_PATH/data/logs/arbiter_web.log 2>&1 &
  fi

  check_rpc_status
  check_web_status
  echo_info "Please check arbiter rpc log via command: cat $SCRIPT_PATH/data/logs/arbiter_rpc.log"
  echo_info "Please check arbiter web log via command: cat $SCRIPT_PATH/data/logs/arbiter_web.log"
}

SCRIPT_PATH=~/loan_arbiter
deploy_arbiter_rpc
deploy_arbiter_web
