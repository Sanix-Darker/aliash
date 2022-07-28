
#!/bin/bash

# The bash script that will get the key of the
# saved alias to the backend and then execute it
#
# By github.com/sanix-darker


HOST="http://127.0.0.1:5000/c"

# the colors list
if [[ -t 1 ]]; then
    COLOROFF='\033[0m'
    Red='\033[0;31m'
    Green='\033[0;32m'
    BGreen='\033[1;32m'
    BWhite='\033[1m'
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
        # echo ">>" $callback
        $callback
    fi
}

# To check if a command is available, if not raise an error
_check_command(){
    $(command -v $1 > /dev/null) &&\
        [[ $? == 1 ]] &&\
        echo "[x] $1 not available, please install it !" &&\
        exit 1
}

# make a HEAD request with curl to get status code
_status_code(){
    curl -I "$1" 2>/dev/null | head -n 1 | cut -d$' ' -f2
}

_get_content(){
    curl -sSL $1 | cat
}

# The run method with a simple eval in it
run(){
    content=$(_get_content $1)
    _confirm "[-] See the content" echo $content
    _confirm "[-] Execute the content" eval $"$content"
}

main(){
    # We check if curl is present or not
    _check_command curl

    # Proceed depending on status-code
    # We run the command
    [[ $(_status_code $1) -eq 200 ]] && run $"$1"
    # Show an error message in case of bad id provided
    [[ $? != 0 ]] && echo -e "[x] An error occured, please check your link again."
}
# We execute the main method
main "$HOST/$1"
