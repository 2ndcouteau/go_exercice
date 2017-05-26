/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:40:10 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 19:06:42 by qrosa            ###   ########.fr       */
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
	nb_map		uint
	nb_err		uint
	err			uint
	x_max		int
	y_max		int
	x_ship		int
	y_ship		int
	dir			string
	road		string
	save_pos	map[int][]int
}

func	check_len(args *[]string) {
	len := len(*args)
	if len != 1 {
		exit_bad_len(len)
	}
}

func	open_file(file_name string) *os.File {
	file, err := os.Open(file_name)
	if err != nil {
		exit_error("ERROR: Open failed.")
	}
	return file
}

func	close_file(file *os.File) {
	if err := file.Close(); err != nil {
		exit_error("ERROR: Close failed.")
	}
}

func	main_loop_process(rd *bufio.Reader) *S_course{
	var new_line	string
	var s_course	S_course
	var err			error

	s_course.save_pos = make(map[int][]int)
	for i := 0; err == nil ; i++ {
		if new_line, err = get_next_line(rd); err == io.EOF { break }
		switch i {
			case 0:
				save_map_coordinate(&s_course, &new_line)
			case 1:
				s_course.nb_map++
				save_ship_position(&s_course, &new_line)
			case 2:
				save_road_line(&s_course, &new_line)
				resolve_course(&s_course)
				i = 0
		}
	}
	return &s_course
}

func	main() {
	args := os.Args[1:]
	check_len(&args)

	open_file := open_file(args[0])
	rd := bufio.NewReader(open_file)

	s_course := main_loop_process(rd)

	close_file(open_file)
	exit_program(s_course)
}
