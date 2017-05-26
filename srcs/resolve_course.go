/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   resolve_course.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:33:27 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/26 02:01:44 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func	resolve_course(s_course *S_course) {
	if (*s_course).ret != 0 { return }
	fmt.Println("COURSE", (*s_course).nb_map, ":", s_course) // DEBUG
}


// func resolution(current_road *road) bool {
//
// 	for elem in current_road->road {
//
// 		switch elem {
// 			case L
// 				turn_left(current_road)
// 			case R
// 				turn_right(current_road)
// 			case M
// 				move_straight(current_road)
// 			default
// 				return ERROR
// 			}
//
// 	}
//
// }
