package tech.cuda.models.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.enums.ScheduleType
import tech.cuda.models.mappers.Task
import tech.cuda.models.mappers.Tasks

/**
 * Created by Jensen on 19-3-5.
 */

object TaskService {
    fun getTasks(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getTasksByUserId(userId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true) and Tasks.user.eq(userId)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getTasksByScheduleType(scheduleType: ScheduleType, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true) and Tasks.scheduleType.eq(scheduleType)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }
}