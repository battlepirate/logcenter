#!/bin/bash 

_NORMAL="\033[0m"
_YELLOW="\033[0;33m"
_CYAN="\033[1;36m"
_GREEN="\033[1;32m"
_RED="\033[1;31m"
_PERPLE="\033[0;35m"

ulimit -c unlimited

usage()
{
printf "${_YELLOW}usage: ${_NORMAL}`basename $0` [start | restart | stop | check | logcenter]\n"
echo "    if there is no ARGS then the script will search the un-started"
echo "    server service and start it."
}

USER_NAME=`whoami`

###servers
LOGCENTER_SVR="logcenter"

### the configure file for each server
# single server (one configure file)
LOGCENTER_SVR_CFG="-config=config/logcenter.sample.config.xml"

### the logs
LOGCENTER_ERR="logs/logcenter.console.log"

#SVR_ARRAY and SVR_CFG_ARRAY 下标关联
SVR_ARRAY=([1]=${LOGCENTER_SVR})
SVR_CFG_ARRAY=([1]=${LOGCENTER_SVR_CFG})
SVR_ERR_ARRAY=([1]=${LOGCENTER_ERR})

CUR_DATE=`date +%Y%m%d%H%M`
VGD="valgrind --tool=memcheck --leak-check=full --leak-resolution=high --show-reachable=yes --log-file=valgrind.$CUR_DATE"

svr_pids=()

# need param: server_name
check_service_exist()
{
#printf "${_YELLOW}################ check if the service${_GREEN} $1 ${_YELLOW}exist ################${_NORMAL}\n"
for idx in ${!SVR_ARRAY[@]}; do
    if [ ${SVR_ARRAY[${idx}]} == ${1} ]; then
        local __svr_cfg_arr=(${SVR_CFG_ARRAY[${idx}]})
        local __svr_cfg_num=${#__svr_cfg_arr[@]}
        svr_pids=() 
        printf "${_YELLOW}check if server${_CYAN} $1 ${_YELLOW}running... \n"
        for cfg_file in ${__svr_cfg_arr[@]}; do
            #if [ $__svr_cfg_num -gt 1 ]; then printf "\n\t\t"; fi
            
            pid=`ps -ef | grep -w "$USER_NAME" | grep "./${1} ${cfg_file}" | grep -v grep | grep -v $0 | awk '{print $2}'`
            if [ "$pid" != "" ]
            then
                printf "\t${_YELLOW}with config file \"${_PERPLE}${cfg_file}${_YELLOW}\": pid=${_GREEN}${pid}\t${_NORMAL}[${_GREEN}RUNNING${_NORMAL}]\n"
                svr_pids[${#svr_pids[@]}]=$pid
            else
                printf "\t${_YELLOW}with config file \"${_PERPLE}${cfg_file}${_YELLOW}\": \t\t${_NORMAL}[${_RED} STOPED${_NORMAL}]\n"
            fi
        done
        break
    fi
done
}

echo_success()
{
	printf "${_NORMAL}[${_GREEN}SUCCESS${_NORMAL}]\n"
}
echo_failed()
{
	printf "${_NORMAL}[${_RED} FAILED${_NORMAL}]\n"
}

kill_service()
{
    check_service_exist $1
    if [ ${#svr_pids[@]} -gt 0 ]
    then
        printf "${_YELLOW}killing running server${_CYAN} $1 ${_YELLOW}with pid=${svr_pids[@]}...${_NORMAL}\t\t"
        kill -9 ${svr_pids[@]} || exit 0
        printf "[${_RED} KILLED${_NORMAL}]\n"
    fi

    sleep 1

}

kill_all_service()
{
    printf "${_YELLOW}################           stop${_CYAN} all ${_YELLOW}service           ################${_NORMAL}\n"
	
    kill_service $LOGCENTER_SVR

    echo "all server killed"
}

start_service_fail()
{
    printf "\n${_RED}start service failed, will stop all started service${_NORMAL}\n"
    kill_all_service
	echo "please check config"
    exit 0
}

start_service()
{
    for idx in ${!SVR_ARRAY[@]}; do
        if [ ${SVR_ARRAY[${idx}]} == ${1} ]; then
            local __svr_cfg_arr=(${SVR_CFG_ARRAY[${idx}]})
            local __svr_cfg_num=${#__svr_cfg_arr[@]}
            
            printf "${_YELLOW}starting server ${_CYAN}$1${_YELLOW} ...${_NORMAL}\n"
            for cfg_file in ${__svr_cfg_arr[@]}; do
               # if [ -f "${cfg_file}" ]; then
                    printf "\t${_YELLOW} executing \"./${1} ${cfg_file}\"...${_NORMAL}\t\t"
                    echo "executing \"./${1} ${cfg_file}\"...`date`" >> ${SVR_ERR_ARRAY[${idx}]}
                    ## start server command
                    #GODEBUG=gctrace=1 nohup ./${1} ${cfg_file} >/dev/null 2>> ${SVR_ERR_ARRAY[${idx}]} & > /dev/null  && echo_success || start_service_fail
                    nohup ./${1} ${cfg_file} >/dev/null 2>> ${SVR_ERR_ARRAY[${idx}]} & > /dev/null  && echo_success || start_service_fail
                    ## check memory
                    #nohup ${VGD} ./${1} ${cfg_file} >/dev/null 2>.errors.log & > /dev/null  && echo_success || start_service_fail
                # fi
            done
    
            break
        fi
    done
}


restart_service()
{
    check_service_exist $1

    if [ ${#svr_pids[@]} -gt 0 ]
    then
        printf "${_YELLOW}killing running server${_CYAN} $1 ${_YELLOW}with pid=${svr_pids[@]} ...${_NORMAL}\t\t"
        kill ${svr_pids[@]} || exit 0
        printf "[${_RED} KILLED${_NORMAL}]\n"
        sleep 1
    fi
    
    start_service $1
}

start_all_service()
{
    printf "${_YELLOW}################          start${_CYAN} all ${_YELLOW}service           ################${_NORMAL}\n" 

	check_service_exist $LOGCENTER_SVR
    if [ ${#svr_pids[@]} -lt 1 ]; then start_service $LOGCENTER_SVR ; fi
    sleep 1
}

restart_all_service()
{
    kill_all_service
    printf "${_YELLOW}################         restart${_CYAN} all ${_YELLOW}service          ################${_NORMAL}\n"
	
	start_service $LOGCENTER_SVR

    printf "${_YELLOW}all server restart succeed${_NORMAL}\n"
}

check_all_service()
{
printf "${_YELLOW}################     check if${_CYAN} all ${_YELLOW}service running     ################${_NORMAL}\n"

check_service_exist $LOGCENTER_SVR

printf "${_YELLOW}######################################################################${_NORMAL}\n"
}

if [ $# -lt 1 ];
then
  # check the un-started server and start it
  start_all_service
  exit 0
fi

case $1 in

    start)
        # check the un-started server and start it
        start_all_service
    ;;

    restart)
        # restart all server
        restart_all_service
    ;;
	
	logcenter | logcenter/)
        restart_service $LOGCENTER_SVR 
    ;;

    usage | help)
        usage
    ;;
	
	stop)
		kill_all_service
	;;	

    check)
        # check if all server running
        check_all_service
    ;;

    *)
        usage
    ;;

esac
