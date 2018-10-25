package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-18.
 */
object User : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val groupId = integer(name = "group_id").index()
    val name = varchar(name = "name", length = 256).index()
    val password = varchar(name = "password", length = 256)
    val email = varchar(name = "email", length = 256)
    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}