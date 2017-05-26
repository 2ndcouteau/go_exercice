/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:40:10 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 02:13:35 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

//	"fmt"
import (
	"os"
	"bufio"
	"io"
)

type S_course struct {
	nb_map	int
	ret		int
	x_max	int
	y_max	int
	x_ship	int
	y_ship	int
	dir		string
	road	string // change to *string ??
}

func	check_len(args *[]string) {
	len := len(*args)
	if len != 1 {
		exit_bad_len(len)
	}
}
func main() {
	args := os.Args[1:]
	check_len(&args)

	// MAYBE simplified OPEN/CLOSE part
	// open input file check for its returned error
	file, err := os.Open(args[0])
	if err != nil {
		exit_error("ERROR: Open failed.") //       or panic(err) ??
	}
	// delay close file on exit and check for its returned error
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)  // exit_error("ERROR: Close failed.")
		}
	}()

	rd := bufio.NewReader(file)
	var new_line string
	var s_course S_course

	for i := 0; err == nil ; i++ {
//		fmt.Println(new_line)		// DEBUG
		if new_line, err = get_next_line(rd); err == io.EOF { break }
		switch i {
			case 0:
				s_course = save_map_coordinate(&new_line)
			case 1:
				s_course.nb_map++
				save_ship_position(&s_course, &new_line)
			case 2:
				save_road_line(&s_course, &new_line)
				resolve_course(&s_course)
				i = 0
		}
	}
	os.Exit(s_course.ret)
}
