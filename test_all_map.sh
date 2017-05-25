for var in `ls map_test/*`
do
	echo ""
	echo "Test $var :"
	./exec_sonda $var

	[ `echo $?` == 0 ] && echo "SUCCESS" || echo "ERROR"
done
