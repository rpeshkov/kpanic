_kpanic_zsh_completion() {
    local state

    _arguments \
        '1: :->cmd'\
        '*: :->panic_name'

    case $state in
        (cmd) _arguments '1:commands:(ls help show last)' ;;
        (*) compadd "$@" $( [ $words[2] == 'show' ] && kpanic ls) ;;
    esac
}

compdef _kpanic_zsh_completion kpanic

