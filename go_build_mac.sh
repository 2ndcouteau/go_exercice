RED=$(printf "\e[31m")
CYAN=$(printf "\e[36m")
RESET=$(printf "\e[m")

go build -o "exec_sonda"	srcs/main.go \
							srcs/get_next_line.go \
							srcs/save_maps.go \
							srcs/resolve_course.go \
							srcs/exit_error_functions.go


[ `echo $?` == 0 ] && echo "Compilation ${CYAN}SUCCESS${RESET}\nUsage: ./exec_sonda FILE" || echo "Compilation ${RED}FAIL${RESET}"
