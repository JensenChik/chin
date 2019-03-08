package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.enums.ScheduleType
import tech.cuda.models.Task
import tech.cuda.models.TaskTable

/**
 * Created by Jensen on 19-3-5.
 */

object TaskService {

    fun getOneById(id: Int): Task? {
        return Task.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = TaskTable.select {
            TaskTable.removed.neq(true)
        }.orderBy(TaskTable.id to false).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getManyByUserId(userId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = TaskTable.select {
            TaskTable.removed.neq(true) and TaskTable.user.eq(userId)
        }.orderBy(TaskTable.id to false).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getManyBySchduleType(scheduleType: ScheduleType, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = TaskTable.select {
            TaskTable.removed.neq(true) and TaskTable.scheduleType.eq(scheduleType)
        }.orderBy(TaskTable.id to false).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun createOne() {

    }


}