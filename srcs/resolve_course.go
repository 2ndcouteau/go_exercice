/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   resolve_course.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:33:27 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 17:00:41 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func	turn_left(s_course *S_course) {
	switch (*s_course).dir {
		case "N", "n":
			(*s_course).dir = "W"
		case "W", "w":
			(*s_course).dir = "S"
		case "S", "s":
			(*s_course).dir = "E"
		case "E", "e":
			(*s_course).dir = "N"
	}
}

func	turn_right(s_course *S_course) {
	switch (*s_course).dir {
		case "N", "n":
			(*s_course).dir = "E"
		case "W", "w":
			(*s_course).dir = "N"
		case "S", "s":
			(*s_course).dir = "W"
		case "E", "e":
			(*s_course).dir = "S"
	}
}

func	move_one(s_course *S_course) {
	switch (*s_course).dir {
		case "N", "n":
			if ((*s_course).y_ship + 1) <= (*s_course).y_max {
				(*s_course).y_ship++
			} else { (*s_course).err++ }
		case "W", "w":
			if ((*s_course).x_ship - 1) >= 0 {
				(*s_course).x_ship--
			} else { (*s_course).err++ }
		case "S", "s":
			if ((*s_course).y_ship - 1) >= 0 {
				(*s_course).y_ship--
			} else { (*s_course).err++ }
		case "E", "e":
			if ((*s_course).x_ship + 1) <= (*s_course).x_max {
				(*s_course).x_ship++
			} else { (*s_course).err++ }
	}
}

func	resolve_course(s_course *S_course) {
	if (*s_course).err != 0 { return }

	for _, char := range (*s_course).road {
		switch char {
			case 'L', 'l':
				turn_left(s_course)
			case 'R', 'r':
				turn_right(s_course)
			case 'M', 'm':
				move_one(s_course)
		}

		if (*s_course).err != 0 {
			print_error("error: Sonda tries to go out of bound. The course is stopped.", s_course)
			return
		}
	}
	fmt.Println("SONDA", (*s_course).nb_map, ":", (*s_course).x_ship, (*s_course).y_ship, (*s_course).dir)
}
