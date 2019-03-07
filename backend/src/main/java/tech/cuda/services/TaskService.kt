package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import tech.cuda.enums.ScheduleType
import tech.cuda.models.Task
import tech.cuda.models.Tasks

/**
 * Created by Jensen on 19-3-5.
 */

object TaskService {

    fun getOneById(id: Int): Task? {
        return Task.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getManyByUserId(userId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true) and Tasks.user.eq(userId)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun getManyBySchduleType(scheduleType: ScheduleType, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = Tasks.select {
            Tasks.removed.neq(true) and Tasks.scheduleType.eq(scheduleType)
        }.orderBy(Tasks.id to true).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun createOne() {

    }


}