package tech.cuda.models

import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-15.
 */
object Instance : Table() {
    val id = integer(name = "id").autoIncrement().primaryKey()
    val jobId = integer(name = "job_id").index()

    val output = blob("output")


    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")

    val finished: Boolean
        get() {
            return false
        }

    val success: Boolean
        get() {
            return false
        }

    fun start() {
    }
}