package tech.cuda.services

import org.jetbrains.exposed.sql.and
import org.jetbrains.exposed.sql.insertAndGetId
import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import tech.cuda.models.Job
import tech.cuda.models.JobTable
import tech.cuda.models.Task

/**
 * Created by Jensen on 19-3-5.
 */

object JobService {

    fun getOneById(id: Int): Job? {
        return Job.findById(id)
    }

    fun getMany(page: Int = 0, pageSize: Int = Int.MAX_VALUE): List<Job> {
        val query = JobTable.select {
            JobTable.removed.neq(true)
        }.orderBy(JobTable.id to false).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }

    fun getManyByTaskId(taskId: Int, page: Int, pageSize: Int): List<Job> {
        val query = JobTable.select {
            JobTable.removed.neq(true) and JobTable.task.eq(taskId)
        }.orderBy(JobTable.id to false).limit(pageSize, offset = page * pageSize)
        return Job.wrapRows(query).toList()
    }

    fun createOneForTask(task: Task): Job {
        val now = DateTime.now()
        val job = Job.new {
            this.task = task
            createTime = now
            updateTime = now
        }
        task.latestJobId = job.id.value
        return job
    }


}