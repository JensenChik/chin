package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-18.
 */
object Group : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val name = varchar(name = "name", length = 256).index()
    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}