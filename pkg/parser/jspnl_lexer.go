// Code generated from JSPNL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 43, 339,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 5, 34, 204, 10, 34, 3, 35, 3, 35, 6,
	35, 208, 10, 35, 13, 35, 14, 35, 209, 3, 35, 6, 35, 213, 10, 35, 13, 35,
	14, 35, 214, 3, 35, 7, 35, 218, 10, 35, 12, 35, 14, 35, 221, 11, 35, 3,
	36, 6, 36, 224, 10, 36, 13, 36, 14, 36, 225, 3, 37, 6, 37, 229, 10, 37,
	13, 37, 14, 37, 230, 3, 37, 3, 37, 6, 37, 235, 10, 37, 13, 37, 14, 37,
	236, 5, 37, 239, 10, 37, 3, 37, 5, 37, 242, 10, 37, 3, 37, 3, 37, 6, 37,
	246, 10, 37, 13, 37, 14, 37, 247, 5, 37, 250, 10, 37, 3, 37, 5, 37, 253,
	10, 37, 3, 37, 6, 37, 256, 10, 37, 13, 37, 14, 37, 257, 3, 37, 3, 37, 5,
	37, 262, 10, 37, 3, 38, 3, 38, 3, 38, 7, 38, 267, 10, 38, 12, 38, 14, 38,
	270, 11, 38, 3, 38, 3, 38, 3, 39, 6, 39, 275, 10, 39, 13, 39, 14, 39, 276,
	3, 40, 6, 40, 280, 10, 40, 13, 40, 14, 40, 281, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 41, 3, 41, 7, 41, 290, 10, 41, 12, 41, 14, 41, 293, 11, 41, 3, 41,
	3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 301, 10, 42, 12, 42, 14, 42,
	304, 11, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3,
	44, 5, 44, 315, 10, 44, 3, 44, 6, 44, 318, 10, 44, 13, 44, 14, 44, 319,
	5, 44, 322, 10, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 4, 268, 302,
	2, 49, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2,
	95, 2, 3, 2, 16, 4, 2, 12, 12, 15, 15, 7, 2, 12, 12, 15, 15, 61, 61, 71,
	72, 81, 81, 4, 2, 11, 11, 34, 34, 3, 2, 50, 59, 4, 2, 71, 71, 103, 103,
	4, 2, 45, 45, 47, 47, 5, 2, 67, 92, 97, 97, 99, 124, 4, 2, 86, 86, 118,
	118, 4, 2, 84, 84, 116, 116, 4, 2, 87, 87, 119, 119, 4, 2, 72, 72, 104,
	104, 4, 2, 67, 67, 99, 99, 4, 2, 78, 78, 110, 110, 4, 2, 85, 85, 117, 117,
	2, 358, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63,
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2,
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2,
	2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 3, 97, 3, 2, 2,
	2, 5, 103, 3, 2, 2, 2, 7, 109, 3, 2, 2, 2, 9, 113, 3, 2, 2, 2, 11, 118,
	3, 2, 2, 2, 13, 123, 3, 2, 2, 2, 15, 128, 3, 2, 2, 2, 17, 131, 3, 2, 2,
	2, 19, 138, 3, 2, 2, 2, 21, 140, 3, 2, 2, 2, 23, 142, 3, 2, 2, 2, 25, 144,
	3, 2, 2, 2, 27, 146, 3, 2, 2, 2, 29, 148, 3, 2, 2, 2, 31, 150, 3, 2, 2,
	2, 33, 152, 3, 2, 2, 2, 35, 154, 3, 2, 2, 2, 37, 156, 3, 2, 2, 2, 39, 158,
	3, 2, 2, 2, 41, 160, 3, 2, 2, 2, 43, 164, 3, 2, 2, 2, 45, 168, 3, 2, 2,
	2, 47, 170, 3, 2, 2, 2, 49, 173, 3, 2, 2, 2, 51, 175, 3, 2, 2, 2, 53, 178,
	3, 2, 2, 2, 55, 181, 3, 2, 2, 2, 57, 184, 3, 2, 2, 2, 59, 187, 3, 2, 2,
	2, 61, 190, 3, 2, 2, 2, 63, 197, 3, 2, 2, 2, 65, 199, 3, 2, 2, 2, 67, 203,
	3, 2, 2, 2, 69, 205, 3, 2, 2, 2, 71, 223, 3, 2, 2, 2, 73, 261, 3, 2, 2,
	2, 75, 263, 3, 2, 2, 2, 77, 274, 3, 2, 2, 2, 79, 279, 3, 2, 2, 2, 81, 285,
	3, 2, 2, 2, 83, 296, 3, 2, 2, 2, 85, 310, 3, 2, 2, 2, 87, 312, 3, 2, 2,
	2, 89, 323, 3, 2, 2, 2, 91, 325, 3, 2, 2, 2, 93, 330, 3, 2, 2, 2, 95, 336,
	3, 2, 2, 2, 97, 98, 7, 114, 2, 2, 98, 99, 7, 110, 2, 2, 99, 100, 7, 99,
	2, 2, 100, 101, 7, 101, 2, 2, 101, 102, 7, 103, 2, 2, 102, 4, 3, 2, 2,
	2, 103, 104, 7, 118, 2, 2, 104, 105, 7, 116, 2, 2, 105, 106, 7, 99, 2,
	2, 106, 107, 7, 112, 2, 2, 107, 108, 7, 117, 2, 2, 108, 6, 3, 2, 2, 2,
	109, 110, 7, 99, 2, 2, 110, 111, 7, 116, 2, 2, 111, 112, 7, 101, 2, 2,
	112, 8, 3, 2, 2, 2, 113, 114, 7, 107, 2, 2, 114, 115, 7, 99, 2, 2, 115,
	116, 7, 116, 2, 2, 116, 117, 7, 101, 2, 2, 117, 10, 3, 2, 2, 2, 118, 119,
	7, 113, 2, 2, 119, 120, 7, 99, 2, 2, 120, 121, 7, 116, 2, 2, 121, 122,
	7, 101, 2, 2, 122, 12, 3, 2, 2, 2, 123, 124, 7, 106, 2, 2, 124, 125, 7,
	99, 2, 2, 125, 126, 7, 116, 2, 2, 126, 127, 7, 101, 2, 2, 127, 14, 3, 2,
	2, 2, 128, 129, 7, 118, 2, 2, 129, 130, 7, 113, 2, 2, 130, 16, 3, 2, 2,
	2, 131, 132, 7, 116, 2, 2, 132, 133, 7, 103, 2, 2, 133, 134, 7, 121, 2,
	2, 134, 135, 7, 99, 2, 2, 135, 136, 7, 116, 2, 2, 136, 137, 7, 102, 2,
	2, 137, 18, 3, 2, 2, 2, 138, 139, 7, 42, 2, 2, 139, 20, 3, 2, 2, 2, 140,
	141, 7, 43, 2, 2, 141, 22, 3, 2, 2, 2, 142, 143, 7, 46, 2, 2, 143, 24,
	3, 2, 2, 2, 144, 145, 7, 63, 2, 2, 145, 26, 3, 2, 2, 2, 146, 147, 7, 125,
	2, 2, 147, 28, 3, 2, 2, 2, 148, 149, 7, 127, 2, 2, 149, 30, 3, 2, 2, 2,
	150, 151, 7, 35, 2, 2, 151, 32, 3, 2, 2, 2, 152, 153, 7, 45, 2, 2, 153,
	34, 3, 2, 2, 2, 154, 155, 7, 47, 2, 2, 155, 36, 3, 2, 2, 2, 156, 157, 7,
	44, 2, 2, 157, 38, 3, 2, 2, 2, 158, 159, 7, 49, 2, 2, 159, 40, 3, 2, 2,
	2, 160, 161, 7, 102, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7, 120, 2,
	2, 163, 42, 3, 2, 2, 2, 164, 165, 7, 111, 2, 2, 165, 166, 7, 113, 2, 2,
	166, 167, 7, 102, 2, 2, 167, 44, 3, 2, 2, 2, 168, 169, 7, 62, 2, 2, 169,
	46, 3, 2, 2, 2, 170, 171, 7, 62, 2, 2, 171, 172, 7, 63, 2, 2, 172, 48,
	3, 2, 2, 2, 173, 174, 7, 64, 2, 2, 174, 50, 3, 2, 2, 2, 175, 176, 7, 64,
	2, 2, 176, 177, 7, 63, 2, 2, 177, 52, 3, 2, 2, 2, 178, 179, 7, 63, 2, 2,
	179, 180, 7, 63, 2, 2, 180, 54, 3, 2, 2, 2, 181, 182, 7, 35, 2, 2, 182,
	183, 7, 63, 2, 2, 183, 56, 3, 2, 2, 2, 184, 185, 7, 40, 2, 2, 185, 186,
	7, 40, 2, 2, 186, 58, 3, 2, 2, 2, 187, 188, 7, 126, 2, 2, 188, 189, 7,
	126, 2, 2, 189, 60, 3, 2, 2, 2, 190, 191, 7, 107, 2, 2, 191, 192, 7, 104,
	2, 2, 192, 193, 7, 103, 2, 2, 193, 194, 7, 110, 2, 2, 194, 195, 7, 117,
	2, 2, 195, 196, 7, 103, 2, 2, 196, 62, 3, 2, 2, 2, 197, 198, 7, 37, 2,
	2, 198, 64, 3, 2, 2, 2, 199, 200, 7, 65, 2, 2, 200, 66, 3, 2, 2, 2, 201,
	204, 5, 91, 46, 2, 202, 204, 5, 93, 47, 2, 203, 201, 3, 2, 2, 2, 203, 202,
	3, 2, 2, 2, 204, 68, 3, 2, 2, 2, 205, 219, 5, 89, 45, 2, 206, 208, 5, 85,
	43, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2,
	209, 210, 3, 2, 2, 2, 210, 218, 3, 2, 2, 2, 211, 213, 5, 89, 45, 2, 212,
	211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215,
	3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 218, 7, 48, 2, 2, 217, 207, 3, 2,
	2, 2, 217, 212, 3, 2, 2, 2, 217, 216, 3, 2, 2, 2, 218, 221, 3, 2, 2, 2,
	219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 70, 3, 2, 2, 2, 221, 219,
	3, 2, 2, 2, 222, 224, 5, 85, 43, 2, 223, 222, 3, 2, 2, 2, 224, 225, 3,
	2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 72, 3, 2, 2,
	2, 227, 229, 5, 85, 43, 2, 228, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2,
	230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232,
	238, 7, 48, 2, 2, 233, 235, 5, 85, 43, 2, 234, 233, 3, 2, 2, 2, 235, 236,
	3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2,
	2, 2, 238, 234, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 241, 3, 2, 2, 2,
	240, 242, 5, 87, 44, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242,
	262, 3, 2, 2, 2, 243, 249, 7, 48, 2, 2, 244, 246, 5, 85, 43, 2, 245, 244,
	3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2,
	2, 2, 248, 250, 3, 2, 2, 2, 249, 245, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2,
	250, 252, 3, 2, 2, 2, 251, 253, 5, 87, 44, 2, 252, 251, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 262, 3, 2, 2, 2, 254, 256, 5, 85, 43, 2, 255, 254,
	3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2,
	2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 5, 87, 44, 2, 260, 262, 3, 2, 2,
	2, 261, 228, 3, 2, 2, 2, 261, 243, 3, 2, 2, 2, 261, 255, 3, 2, 2, 2, 262,
	74, 3, 2, 2, 2, 263, 268, 7, 36, 2, 2, 264, 267, 5, 95, 48, 2, 265, 267,
	10, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 270, 3, 2,
	2, 2, 268, 269, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2,
	270, 268, 3, 2, 2, 2, 271, 272, 7, 36, 2, 2, 272, 76, 3, 2, 2, 2, 273,
	275, 9, 3, 2, 2, 274, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 274,
	3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 78, 3, 2, 2, 2, 278, 280, 9, 4,
	2, 2, 279, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2,
	281, 282, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 8, 40, 2, 2, 284,
	80, 3, 2, 2, 2, 285, 286, 7, 49, 2, 2, 286, 287, 7, 49, 2, 2, 287, 291,
	3, 2, 2, 2, 288, 290, 10, 2, 2, 2, 289, 288, 3, 2, 2, 2, 290, 293, 3, 2,
	2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 294, 3, 2, 2, 2,
	293, 291, 3, 2, 2, 2, 294, 295, 8, 41, 3, 2, 295, 82, 3, 2, 2, 2, 296,
	297, 7, 49, 2, 2, 297, 298, 7, 44, 2, 2, 298, 302, 3, 2, 2, 2, 299, 301,
	11, 2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 303, 3, 2,
	2, 2, 302, 300, 3, 2, 2, 2, 303, 305, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2,
	305, 306, 7, 44, 2, 2, 306, 307, 7, 49, 2, 2, 307, 308, 3, 2, 2, 2, 308,
	309, 8, 42, 3, 2, 309, 84, 3, 2, 2, 2, 310, 311, 9, 5, 2, 2, 311, 86, 3,
	2, 2, 2, 312, 314, 9, 6, 2, 2, 313, 315, 9, 7, 2, 2, 314, 313, 3, 2, 2,
	2, 314, 315, 3, 2, 2, 2, 315, 321, 3, 2, 2, 2, 316, 318, 5, 85, 43, 2,
	317, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319,
	320, 3, 2, 2, 2, 320, 322, 3, 2, 2, 2, 321, 317, 3, 2, 2, 2, 321, 322,
	3, 2, 2, 2, 322, 88, 3, 2, 2, 2, 323, 324, 9, 8, 2, 2, 324, 90, 3, 2, 2,
	2, 325, 326, 9, 9, 2, 2, 326, 327, 9, 10, 2, 2, 327, 328, 9, 11, 2, 2,
	328, 329, 9, 6, 2, 2, 329, 92, 3, 2, 2, 2, 330, 331, 9, 12, 2, 2, 331,
	332, 9, 13, 2, 2, 332, 333, 9, 14, 2, 2, 333, 334, 9, 15, 2, 2, 334, 335,
	9, 6, 2, 2, 335, 94, 3, 2, 2, 2, 336, 337, 7, 94, 2, 2, 337, 338, 7, 36,
	2, 2, 338, 96, 3, 2, 2, 2, 27, 2, 203, 209, 214, 217, 219, 225, 230, 236,
	238, 241, 247, 249, 252, 257, 261, 266, 268, 276, 281, 291, 302, 314, 319,
	321, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'place'", "'trans'", "'arc'", "'iarc'", "'oarc'", "'harc'", "'to'",
	"'reward'", "'('", "')'", "','", "'='", "'{'", "'}'", "'!'", "'+'", "'-'",
	"'*'", "'/'", "'div'", "'mod'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='",
	"'&&'", "'||'", "'ifelse'", "'#'", "'?'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LOGICAL",
	"ID", "INT", "FLOAT", "STRING", "NEWLINE", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "LOGICAL",
	"ID", "INT", "FLOAT", "STRING", "NEWLINE", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
	"DIGIT", "EXPONENT", "CHAR", "TRUE", "FALSE", "ESCAPED_QUOTE",
}

type JSPNLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewJSPNLLexer(input antlr.CharStream) *JSPNLLexer {

	l := new(JSPNLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "JSPNL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JSPNLLexer tokens.
const (
	JSPNLLexerT__0          = 1
	JSPNLLexerT__1          = 2
	JSPNLLexerT__2          = 3
	JSPNLLexerT__3          = 4
	JSPNLLexerT__4          = 5
	JSPNLLexerT__5          = 6
	JSPNLLexerT__6          = 7
	JSPNLLexerT__7          = 8
	JSPNLLexerT__8          = 9
	JSPNLLexerT__9          = 10
	JSPNLLexerT__10         = 11
	JSPNLLexerT__11         = 12
	JSPNLLexerT__12         = 13
	JSPNLLexerT__13         = 14
	JSPNLLexerT__14         = 15
	JSPNLLexerT__15         = 16
	JSPNLLexerT__16         = 17
	JSPNLLexerT__17         = 18
	JSPNLLexerT__18         = 19
	JSPNLLexerT__19         = 20
	JSPNLLexerT__20         = 21
	JSPNLLexerT__21         = 22
	JSPNLLexerT__22         = 23
	JSPNLLexerT__23         = 24
	JSPNLLexerT__24         = 25
	JSPNLLexerT__25         = 26
	JSPNLLexerT__26         = 27
	JSPNLLexerT__27         = 28
	JSPNLLexerT__28         = 29
	JSPNLLexerT__29         = 30
	JSPNLLexerT__30         = 31
	JSPNLLexerT__31         = 32
	JSPNLLexerLOGICAL       = 33
	JSPNLLexerID            = 34
	JSPNLLexerINT           = 35
	JSPNLLexerFLOAT         = 36
	JSPNLLexerSTRING        = 37
	JSPNLLexerNEWLINE       = 38
	JSPNLLexerWS            = 39
	JSPNLLexerLINE_COMMENT  = 40
	JSPNLLexerBLOCK_COMMENT = 41
)