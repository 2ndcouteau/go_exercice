/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   resolve_course.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:33:27 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/27 19:56:40 by yoko             ###   ########.fr       */
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

func	value_exist(value int, coor []int) bool {
    for _, check_value := range coor {
        if check_value == value {
            return true
        }
    }
    return false
}

func	check_save_ship(s_course *S_course, ns int, we int) bool {
	var is_present bool

	_, is_present = (*s_course).save_pos[(*s_course).x_ship + we]
	if	is_present == true {
		is_present = value_exist((*s_course).y_ship + ns, (*s_course).save_pos[(*s_course).x_ship + we])
		if is_present == true {
			return false
		}
	}
	return true
}

func	move_one(s_course *S_course) {
	switch (*s_course).dir {
		case "N", "n":
			if ((*s_course).y_ship + 1) <= (*s_course).y_max && check_save_ship(s_course, 1, 0) {
				(*s_course).y_ship++
			} else { (*s_course).err++ }
		case "W", "w":
			if ((*s_course).x_ship - 1) >= 0 && check_save_ship(s_course, 0, -1){
				(*s_course).x_ship--
			} else { (*s_course).err++ }
		case "S", "s":
			if ((*s_course).y_ship - 1) >= 0 && check_save_ship(s_course, -1, 0){
				(*s_course).y_ship--
			} else { (*s_course).err++ }
		case "E", "e":
			if ((*s_course).x_ship + 1) <= (*s_course).x_max && check_save_ship(s_course, 0, 1){
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
			save_valid_probe(s_course)
//			fmt.Println("PROBE", (*s_course).nb_map, ":", (*s_course).x_ship, (*s_course).y_ship, (*s_course).dir) // DEBUG
			print_error("error: Collision detection activate, the probe is shut down.", s_course)
			return
		}
	}
	save_valid_probe(s_course)
	fmt.Println("PROBE", (*s_course).nb_map, ":", (*s_course).x_ship, (*s_course).y_ship, (*s_course).dir)
}
