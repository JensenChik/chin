package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.exceptions.StringOutOfLengthException
import tech.cuda.models.*

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

    fun getManyByGroupId(groupId: Int, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val users = UserService.getManyByGroupId(groupId)
        val query = TaskTable.select {
            TaskTable.removed.neq(true) and TaskTable.user.inList(users.map { it.id })
        }.orderBy(TaskTable.id to false).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()

    }

    fun getManyBySchduleType(scheduleType: ScheduleType, page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Task> {
        val query = TaskTable.select {
            TaskTable.removed.neq(true) and TaskTable.scheduleType.eq(scheduleType)
        }.orderBy(TaskTable.id to false).limit(pageSize, offset = page * pageSize)
        return Task.wrapRows(query).toList()
    }

    fun createOne(
            user: User,
            name: String,
            scheduleType: ScheduleType,
            scheduleFormat: ScheduleFormat,
            command: String
    ): Task {
        val now = DateTime.now()
        return when {
            name.length > TaskTable.NAME_MAX_LEN ->
                throw StringOutOfLengthException("name", TaskTable.NAME_MAX_LEN)
            else -> Task.new {
                this.user = user
                this.name = name
                this.scheduleType = scheduleType
                this.scheduleFormat = scheduleFormat
                this.command = command
                latestJobId = null
                removed = false
                createTime = now
                updateTime = now
            }
        }


    }


}