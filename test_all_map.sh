RED=$(printf "\e[31m")
GREEN=$(printf "\e[32m")
YELLOW=$(printf "\e[33m")
BLUE=$(printf "\e[34m")
MAGENTA=$(printf "\e[35m")
CYAN=$(printf "\e[36m")
WHITE=$(printf "\e[37m")
RESET=$(printf "\e[m")

for var in `ls map_test/*`
do
	echo ""
	echo "${BLUE}#${RESET} Test $var :"
	./exec_sonda $var

	[ `echo $?` == 0 ] && echo "${CYAN}SUCCESS${RESET}" || echo "${RED}ERROR${RESET}"
done
