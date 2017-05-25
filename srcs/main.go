/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:40:10 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/25 22:00:49 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)



type S_road struct {
	x_max	int
	y_max	int
	x_ship	int
	y_ship	int
	dir		byte
	road	string // change to *string
}

func	check_len(args *[]string) {
	len := len(*args)
	if len != 1 {
		exit_bad_len(len)
	}
}

func	save_map_coordinate(line *string) *S_road{
//	elems := strings.Split(line, " ")
	s_road := new(S_road)
	s_road.x_max = 32
	s_road.y_max = 42
	s_road.x_ship = 84
	s_road.y_ship = -12
	s_road.dir = 'K'
	s_road.road = "Bonjour Brasil"
	return s_road
}

func main() {
	args := os.Args[1:]
	check_len(&args)

	// open input file check for its returned error
	file, err := os.Open(args[0])
	if err != nil {
		exit_error("ERROR: Open failed") //       or panic(err) ??
	}
	// delay close file on exit and check for its returned error
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	rd := bufio.NewReader(file)
	new_line, err := get_next_line(rd)

	//	var road *S_road
//	road := save_map_coordinate(&new_line)			//save coodonate

	for err == nil {
		fmt.Println(new_line)		// DEBUG
		new_line, err = get_next_line(rd)
		// save ship position
	//	elem := strings.Split(new_line, " ")
		fmt.Println(new_line)		// DEBUG
		new_line, err = get_next_line(rd)
		// Trim line
		// save road

		// call resolution
	}
	if err == io.EOF {
		fmt.Println("var err ==", err) }
	os.Exit(0)
}
