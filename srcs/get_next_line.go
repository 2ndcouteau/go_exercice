/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_next_line.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: qrosa <qrosa@student.42.fr>                +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/05/25 20:34:05 by qrosa             #+#    #+#             */
/*   Updated: 2017/05/25 20:44:37 by qrosa            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "bufio"

/*
	get_next_line returns a single line (without the ending \n) from the input
	buffered reader.
	An error is returned if there is an error with the buffered reader.
	This function is usefull for bypass Scanner method len limitation ~65k char

	Try to return pointer == manual allocation ??
*/
func	get_next_line(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool = true
		err error = nil
		buf_line, line []byte
	)

	for isPrefix && err == nil {
		buf_line, isPrefix, err = r.ReadLine()
		line = append(line, buf_line...)		// Concat lines
	}
	return string(line), err				// Return (string(ln) for suppress \n)
}
