/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   exit_error_functions.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 20:45:22 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 19:02:31 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func	print_error(str string, s_course *S_course) {
	(*s_course).err = 1
	(*s_course).nb_err++
	fmt.Fprintln(os.Stderr, "PROBE", (*s_course).nb_map, ":", str)
}

func	exit_error(str string) {
	fmt.Fprintln(os.Stderr, str)
	os.Exit(1);
}

func	exit_bad_len(len int) {
	if len > 1 {
		exit_error("-- ERROR: Too many arguments\n-- Help : ./sonda FILE")
	} else if len == 0 {
		exit_error("-- ERROR: You need to specify a file\n-- Help : ./sonda FILE")
	}
}

func	exit_program(s_course *S_course) {
	ret := 0;

	fmt.Print("")
	if s_course.nb_err != 0 {
		ret = 1
		fmt.Printf("FAILED: [%d/%d]\n", (*s_course).nb_map - (*s_course).nb_err, (*s_course).nb_map)
	} else {
		fmt.Printf("SUCCESS: [%d/%d]\n", (*s_course).nb_map - (*s_course).nb_err, (*s_course).nb_map)
	}
	os.Exit(ret)
}
