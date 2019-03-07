package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.models.Job
import tech.cuda.models.Jobs
import tech.cuda.models.Task

/**
 * Created by Jensen on 19-3-5.
 */

object JobService {

    fun getOneById(id: Int): Job? {
        return Job.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Job> {
        val query = Jobs.select {
            Jobs.removed.neq(true)
        }.orderBy(Jobs.id to true).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }

    fun getManyByTaskId(taskId: Int, page: Int, pageSize: Int): List<Job> {
        val query = Jobs.select {
            Jobs.removed.neq(true) and Jobs.task.eq(taskId)
        }.orderBy(Jobs.id to true).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }

    fun createOneForTask(_task: Task) {
        val now = DateTime.now()
        val job = Job.new {
            task = _task
            createTime = now
            updateTime = now
        }
        _task.latestJobId = job.id.value
    }


}