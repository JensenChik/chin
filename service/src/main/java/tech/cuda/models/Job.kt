package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-18.
 */
object Job : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val taskId = integer(name = "task_id").index()


    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")

    val shouldRunNow: Boolean
        get() {
            return false
        }

    fun createInstance() {

    }

}