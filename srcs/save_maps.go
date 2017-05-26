/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   save_maps.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/26 01:22:14 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 15:34:17 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"strings"
	"strconv"
	"regexp"
)

func	save_map_coordinate(line *string) S_course {
	var	s_course	S_course
	var err			error

	// fmt.Println("Save coordinate") // DEBUG
	elems := strings.Split(*line, " ") // CHECK return ??
	if len(elems) != 2 {
		exit_error("ERROR: Map line coordinate is not compliant.")
	}
	if s_course.x_max, err = strconv.Atoi(elems[0]); err != nil {
		exit_error("ERROR: X map value is not compliant.")
	} else if s_course.x_max < 0 {
		exit_error("ERROR: X map value can not be negative.")
	}
	if s_course.y_max, err = strconv.Atoi(elems[1]); err != nil {
		exit_error("ERROR: Y map value is not compliant.")
	} else if s_course.y_max < 0 {
		exit_error("ERROR: Y map value can not be negative.")
	}
	return s_course
}

func	save_ship_position(s_course *S_course, line *string) {
	var err			error

	// fmt.Println("Save ship positions") // DEBUG
	elems := strings.Split(*line, " ")
	(*s_course).err = 0
	if len(elems) != 3 {
		print_error("error: Ship position line is not compliant.", s_course)
		return
	}
	if (*s_course).x_ship, err = strconv.Atoi(elems[0]); err != nil {
		print_error("error: X ship position is not compliant.", s_course)
		return
	} else if (*s_course).x_ship < 0 {
		print_error("error: X ship position can not be negative.", s_course)
		return
	}
	if (*s_course).y_ship, err = strconv.Atoi(elems[1]); err != nil {
		print_error("error: Y ship position is not compliant.", s_course)
		return
	} else if (*s_course).y_ship < 0 {
		print_error("error: Y ship position can not be negative.", s_course)
		return
	}
	(*s_course).dir = elems[2]
	if match, _ := regexp.MatchString("[^NSEWnsew]{1}", (*s_course).dir); match == true {
		print_error("error: The ship direction does not exist.", s_course)
		return
	}
}

func	save_road_line(s_course *S_course, line *string) {
	// fmt.Println("Save road") // DEBUG
	if (*s_course).err != 0 { return }
	(*s_course).road = strings.TrimSpace(*line)
	if match, _ := regexp.MatchString("[^LRMrlm]+", (*s_course).road); match == true {
		print_error("error: The road is not valid", s_course)
		return
	}
}
