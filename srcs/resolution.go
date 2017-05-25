/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   resolution.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 18:33:27 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/25 18:39:39 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */



func resolution(current_road *road) bool {

	for elem in current_road->road {

		switch elem {
			case L
				turn_left(current_road)
			case R
				turn_right(current_road)
			case M
				move_straight(current_road)
			default
				return ERROR
			}

	}

}
