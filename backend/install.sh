#!/bin/bash

# The bash script that will get the key of the
# saved alias to the backend and then execute it
#
# By github.com/sanix-darker

# The host link of aliash
HOST="http://127.0.0.1:5000"

# THe temporary file where the content of bash snippets will be saved
TMP="./tmp"

TMP_HASH=""

# the colors list
if [[ -t 1 ]]; then
    COLOROFF='\033[0m'       # TEXT RESET
    BLACK='\033[0;30m'        # BLACK
    RED='\033[0;31m'          # RED
    GREEN='\033[0;32m'        # GREEN
    YELLOW='\033[0;33m'       # YELLOW
    BLUE='\033[0;34m'         # BLUE
    PURPLE='\033[0;35m'       # PURPLE
    CYAN='\033[0;36m'         # CYAN
    WHITE='\033[0;37m'        # WHITE

    # BOLD
    BBLACK='\033[1;30m'       # BLACK
    BRED='\033[1;31m'         # RED
    BGREEN='\033[1;32m'       # GREEN
    BYELLOW='\033[1;33m'      # YELLOw
    BBLUE='\033[1;34m'        # BLUE
    BPURPLE='\033[1;35m'      # PURPLE
    BCYAN='\033[1;36m'        # CYAN
    BWHITE='\033[1;37m'       # WHITE
fi

# With a given message as input, this function will execute anything
# after the second argument passed
# Ex : _confirm "Message" echo "test"
_confirm(){
    args=("${@}")
    read -p "${args[0]} (Y/y)? " -n 1 -r; echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        callback=${args[@]:1}
        echo ">>" $callback
        $callback
    fi
}

# To check if a command is available, if not raise an error
_check_command(){
    command -v $1 > /dev/null
    [[ $? == 1 ]] && echo "[x] $1 not available, please install it !" && exit 1
}

_count_down(){
    sec=15
    while [ $sec -ge 0 ]; do
        echo -ne "$GREEN[-] The script will be started in "$BRED$sec$COLOROFF" (Crtl+C to stop) \r"
        let "sec=sec-1"
        sleep 1
    done
    echo
}

# The run method with a simple eval in it
run_it(){
    echo -e "$GREEN[-] Content:$COLOROFF"
    echo "--------------------------------------------------------------------------------"
    echo -e $BWHITE
    cat $TMP
    echo -e $COLOROFF
    echo "--------------------------------------------------------------------------------"
    echo -e $BLUE**SHA512: ${TMP_HASH:0:70}$COLOROFF
    echo "--------------------------------------------------------------------------------"

    _count_down

    while read line; do
        [[  ${#line} > 1 ]] && echo -e "|>  $BYELLOW$line$COLOROFF"

        # we execute the corresponding line
        bash -c "$line"

        [[ $? != 0 ]] && echo -e "$RED[x] Error executing  : $BRED$line$COLOROFF"
        sleep 0.5
    done < $TMP

    # We clean the tmp
    rm -rf $TMP
}

main(){
    # We check if some required programs are availables
    for cc in "curl" "sha512sum" ;do
        _check_command $cc
    done

    # we save in a tmp file
    curl -sSL $1 | tr "\n" "\\n" > $TMP
    if [[ $? == 0 ]]; then
        TMP_HASH=$(cat $TMP | sha512sum)
        # We run the command
        run_it
    else
        # Show an error message in case of bad id provided
        echo -e "[x] An error occured, please check your link again."
    fi
}
# We execute the main method
main "$HOST/$1"
