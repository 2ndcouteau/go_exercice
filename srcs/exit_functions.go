/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   exit_function.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 20:45:22 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/25 20:46:03 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

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
