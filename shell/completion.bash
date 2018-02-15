_kpanic_bash_complete()
{
    local cur_word prev_word type_list

    cur_word="${COMP_WORDS[COMP_CWORD]}"
    prev_word="${COMP_WORDS[COMP_CWORD-1]}"


    if [[ ${prev_word} == 'show' ]] ; then
	type_list=$(kpanic ls)
	COMPREPLY=( $(compgen -W "${type_list}" -- ${cur_word}) )
    else
	COMPREPLY=()
    fi
    return 0
}

complete -F _kpanic_bash_complete kpanic

