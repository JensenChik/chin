package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-18.
 */
object Operation : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val userId = integer(name = "id")
    val action = varchar(name = "action", length = 16)


    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}