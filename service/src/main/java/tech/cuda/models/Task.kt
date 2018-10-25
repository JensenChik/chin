package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-18.
 */
object Task : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val groupId = integer(name = "group_id")
    val userId = integer(name = "user_id")
    val name = varchar(name = "name", length = 256)
    val command = text(name = "command")

    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")

    val shouldScheduledToday: Boolean
        get() {
            return false
        }

    fun createJob() {

    }
}