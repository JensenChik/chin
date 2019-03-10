package tech.cuda.exceptions

/**
 * Created by Jensen on 19-1-26.
 */

class StringOutOfLengthException(column: String, maxLen: Int) :
        Exception("length of column `$column` must less than $maxLen") {
}